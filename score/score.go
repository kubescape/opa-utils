package score

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strings"
	"sync"

	k8sinterface "github.com/kubescape/k8s-interface/k8sinterface"
	"github.com/kubescape/k8s-interface/workloadinterface"
	armoupautils "github.com/kubescape/opa-utils/objectsenvelopes"
	"github.com/kubescape/opa-utils/reporthandling"
	"github.com/kubescape/opa-utils/reporthandling/apis"
	"github.com/kubescape/opa-utils/reporthandling/results/v1/reportsummary"
	v2 "github.com/kubescape/opa-utils/reporthandling/v2"
	"github.com/kubescape/opa-utils/shared"
	"go.uber.org/zap"
	appsv1 "k8s.io/api/apps/v1"
)

const (
	// replicaFactor defines how the score evolves whenever a workload defines multiple replicas.
	//
	// A value greater than 1 makes replicas an aggravating factor of the score.
	//
	// Example: with a factor of 1.1, any extra replica compounds an extra 10% to the score
	replicaFactor = 1.1
)

type (
	ControlScoreWeights struct {
		BaseScore                    float32 `json:"baseScore"`
		RuntimeImprovementMultiplier float32 `json:"improvementRatio"`
	}

	// ScoreUtil knows how to compute vulnerability risk scores for frameworks and the controls they define.
	ScoreUtil struct {
		K8SApoObj   *k8sinterface.KubernetesApi
		resources   map[string]workloadinterface.IMetadata
		isDebugMode bool

		// ResourceTypeScores map[string]float32
		// FrameworksScore    map[string]map[string]ControlScoreWeights
		// configPath  string
	}
)

var (
	postureScoreSingleton sync.Once
	postureScore          *ScoreUtil
)

// NewScore build a new ScoreUtil computer.
func NewScore(allResources map[string]workloadinterface.IMetadata) *ScoreUtil {
	postureScoreSingleton.Do(func() {
		// NOTE(fredbi): I don't really understand why we need this to be a singleton.
		// IMHO we should avoid this kind of package-level stickiness.
		//
		// Anyway for now, I am just fixing the data race on the initialization for now.
		postureScore = &ScoreUtil{
			resources:   allResources,
			isDebugMode: strings.EqualFold(os.Getenv("ARMO_DEBUG_MODE"), "true"),
		}
	})

	return postureScore
}

// Calculate scores from a list of framework reports.
//
// Each report is updated in place with the score.
func (su *ScoreUtil) Calculate(frameworksReports []reporthandling.FrameworkReport) error {
	for i := range frameworksReports {
		_ = su.CalculateFrameworkScore(&frameworksReports[i])
	}

	return nil
}

// CalculateFrameworkScore computes the score according to a given framework report.
//
// This method works against the data model v1.
//
// The report is updated in place.
func (su *ScoreUtil) CalculateFrameworkScore(framework *reporthandling.FrameworkReport) error {
	framework.Score = 0

	for i := range framework.ControlReports {
		framework.ControlReports[i].Score = 0
		wcsCtrl, unormalizedScore := su.ControlScore(&framework.ControlReports[i], framework.Name)
		su.debugf(
			"control %s(%s) failed %v wcs %v (baseScore: %v)",
			framework.ControlReports[i].Name,
			framework.ControlReports[i].ControlID,
			unormalizedScore,
			wcsCtrl,
			framework.ControlReports[i].BaseScore,
		)

		framework.WCSScore += wcsCtrl
		framework.Score += unormalizedScore
		framework.ARMOImprovement += framework.ControlReports[i].ARMOImprovement
	}

	if framework.WCSScore == 0 {
		framework.Score = 0

		return fmt.Errorf("unable to calculate score for framework %s due to bad wcs score\n", framework.Name)
	}

	framework.Score = (framework.Score * 100) / framework.WCSScore
	su.debugf("framework %s score %v", framework.Name, framework.Score)

	framework.ARMOImprovement = (framework.ARMOImprovement * 100) / framework.WCSScore

	return nil
}

// GetScore decodes a workload from the input map and yields the associated score.
//
// The default score is 1.00 for any object that is not recognized.
//
// Special rules:
//   - daemonset:
//     daemonsetScore * #desired nodes
//   - workloads with replicas:
//     replicaFactor * workloadkindscore * #replicas
//
// This method is used with all versions of the reporthandling model.
func (su *ScoreUtil) GetScore(v map[string]interface{}) float32 {
	const defaultScore = float32(1.00)

	switch {
	case k8sinterface.IsTypeWorkload(v):
		wl := workloadinterface.NewWorkloadObj(v)

		return su.processWorkload(wl, defaultScore, v)

	case armoupautils.IsTypeRegoResponseVector(v):
		vec := armoupautils.NewRegoResponseVectorObject(v)
		related := vec.GetRelatedObjects()
		score := defaultScore

		for i := range related {
			if !k8sinterface.IsTypeWorkload(related[i].GetObject()) {
				continue
			}

			wl := workloadinterface.NewWorkloadObj(v)
			score = max32(score, su.processWorkload(wl, score, v))
		}

		return score

	default:
		return defaultScore
	}
}

// processWorkload handles special scoring rules for workloads with replicas (e.g. deployments, statefulsets) and DaemonSets.
func (su *ScoreUtil) processWorkload(wl *workloadinterface.Workload, score float32, v map[string]interface{}) float32 {
	replicas := wl.GetReplicas()
	if replicas > 1 {
		score *= float32(replicas) * replicaFactor
	}

	if !strings.EqualFold(wl.GetKind(), "daemonset") {
		return score
	}

	/* TODO - replace marshal and unmarshal by map inspection, like so:
	if n, ok := workloadinterface.InspectMap(v, "status", "desiredNumberScheduled"); ok {
		if desiredNumberScheduled, ok := n.(int32); ok && desiredNumberScheduled > 0 {
			score *= float32(desiredNumberScheduled)
		}
	}
	*/
	b, err := json.Marshal(v)
	if err != nil {
		return score
	}

	// special rule for DaemonSets
	dmnset := appsv1.DaemonSet{}
	_ = json.Unmarshal(b, &dmnset)

	if dmnset.Status.DesiredNumberScheduled > 0 {
		score *= float32(dmnset.Status.DesiredNumberScheduled)
	}

	return score
}

// ControlScore yields the unnormalized score contribution of a framework's control, as well as the weight used to normalize.
//
// This method works against the data model v1.
//
// ctrlReport: reporthandling.ControlReport object, must contain down the line the input resources as well as the output resources
// frameworkName: calculates this control according to a given framework weights (currently unused)
//
// ctrl.score = baseScore * SUM_resource (resourceWeight*min(#replicas*replicaweight,1)(nodes if daemonset)
//
// The input ctrlReport is updated with the new (normalized) score, that is the percentage: controlScore*100/wssscore.
//
// Returns wcsscore,ctrlscore(unnormalized)
//
// The wcsscore is evaluated over all resources, whereas the control score is evaluated only on resources that have failed this control.
func (su *ScoreUtil) ControlScore(ctrlReport *reporthandling.ControlReport, _ /* frameworkName */ string) (float32, float32) {
	resourceIDs := ctrlReport.ListResourcesIDs()
	failedResourceIDs := resourceIDs.GetFailedResources()
	allResourceIDs := resourceIDs.GetAllResources()

	for i := range failedResourceIDs {
		if failedResourceIDs, ok := su.resources[failedResourceIDs[i]]; ok {
			ctrlReport.Score += su.GetScore(failedResourceIDs.GetObject())
		}
	}
	ctrlReport.Score *= ctrlReport.BaseScore

	var wcsScore float32
	for i := range allResourceIDs {
		if resource, ok := su.resources[allResourceIDs[i]]; ok {
			wcsScore += su.GetScore(resource.GetObject())
		}
	}

	// NOTE(fredbi): in V2, wcs weights are computed differently for items with a zero score ("passed")
	if ctrlReport.Score != 0 {
		wcsScore *= ctrlReport.BaseScore
	} else {
		wcsScore = ctrlReport.BaseScore
	}

	unormalizedScore := ctrlReport.Score
	ctrlReport.ARMOImprovement = unormalizedScore * ctrlReport.ARMOImprovement
	if wcsScore > 0 {
		ctrlReport.Score = (ctrlReport.Score * 100) / wcsScore
	} else {
		zap.L().Error("worst case scenario was 0, meaning no resources input were given - score is not available(will appear as > 1)")
	}
	su.debugf("control %q un-normalized score: %v, wcs: %v, improvement: %v", ctrlReport.ControlID, unormalizedScore, wcsScore, ctrlReport.ARMOImprovement)

	return wcsScore, unormalizedScore

}

// CalculatePostureReportV2 calculates controls by framework score.
func (su *ScoreUtil) CalculatePostureReportV2(report *v2.PostureReport) error {
	for i := range report.SummaryDetails.Frameworks {
		report.SummaryDetails.Frameworks[i].Score = 0
		var wcsFwork float32
		fwUnormalizedScore, wcsFwork := su.ControlsSummariesScore(&report.SummaryDetails.Frameworks[i].Controls, report.SummaryDetails.Frameworks[i].GetName())

		if wcsFwork == 0 { // NOTE(fred): since this is a float32, perhaps we should use a tolerance here
			report.SummaryDetails.Frameworks[i].Score = 0

			return fmt.Errorf(
				"unable to calculate score for framework %s due to bad wcs score\n",
				report.SummaryDetails.Frameworks[i].GetName(),
			)
		}

		report.SummaryDetails.Frameworks[i].Score = (fwUnormalizedScore * 100) / wcsFwork
		su.debugf("framework %s score %v", report.SummaryDetails.Frameworks[i].GetName(), report.SummaryDetails.Frameworks[i].GetScore())
	}

	totalUnormalizedScore, totalWcsScore := su.ControlsSummariesScore(&report.SummaryDetails.Controls, "") // populate summaries per control
	su.debugf("total un-normalized score %v, (total wcs: %v)", totalUnormalizedScore, totalWcsScore)
	report.SummaryDetails.Score = (totalUnormalizedScore * 100) / totalWcsScore // populate final summarized score

	return nil
}

func (su *ScoreUtil) ControlsSummariesScore(ctrls *reportsummary.ControlSummaries, frameworkName string) (totalUnormalizedScore float32, totalWcsScore float32) {
	totalUnormalizedScore = 0
	totalWcsScore = 0

	for ctrlID := range *ctrls {
		ctrl := (*ctrls)[ctrlID]
		ctrl.Score = 0
		ctrlScore, unormScore, wcs := su.ControlV2Score(&ctrl, frameworkName)
		ctrl.Score = ctrlScore
		(*ctrls)[ctrlID] = ctrl
		totalUnormalizedScore += unormScore
		totalWcsScore += wcs
	}

	return totalUnormalizedScore, totalWcsScore
}

// ControlV2Score returns the score for a given control (as a percentage), the unnormalized score and the weight.
//
// This method works against the data model v2.
//
// IControlSummary: requires a fully populated set of controls (with resources, statuses and score factor hydrated).
// We assume that ListResourcesIDs() operates the same as when scanning controls declared by frameworks.
//
// frameworkName - calculate this control according to a given framework weights (unused for now)
//
// ctrl.score = baseScore * SUM_resource (resourceWeight*min(#replicas*replicaweight,1)(nodes if daemonset)
//
// Returns: ctrlscore(normalized),ctrlscore(unnormalized),wcsscore,
func (su *ScoreUtil) ControlV2Score(ctrl reportsummary.IControlSummary, _ /*frameworkName*/ string) (ctrlScore float32, unormalizedScore float32, wcsScore float32) {
	resourcesIDs := ctrl.ListResourcesIDs(nil) // TODO: use pool?

	for resourceId, status := range resourcesIDs.All() {
		if status == apis.StatusFailed {
			if _, ok := su.resources[resourceId]; ok {
				unormalizedScore += su.GetScore(su.resources[resourceId].GetObject())
			}
		}
	}

	unormalizedScore *= ctrl.GetScoreFactor()

	for resourceId := range resourcesIDs.All() {
		if _, ok := su.resources[resourceId]; ok {
			wcsScore += su.GetScore(su.resources[resourceId].GetObject())
		}
	}

	wcsScore *= ctrl.GetScoreFactor()

	// ctrlReport.ARMOImprovement = unormalizedScore * ctrlReport.ARMOImprovement
	if wcsScore > 0 {
		ctrlScore = (unormalizedScore * 100) / wcsScore
	} else {
		zap.L().Error("worst case scenario was 0, meaning no resources input were given - score is not available(will appear as > 1)")
	}
	su.debugf("control %q score:%v, unnormalized:%v, wcs:%v)", ctrl.GetID(), ctrlScore, unormalizedScore, wcsScore)

	return ctrlScore, unormalizedScore, wcsScore

}

func (su ScoreUtil) debugf(format string, args ...any) {
	if !su.isDebugMode {
		return
	}

	fmt.Printf(format+"\n", args...)
}

func max32(a, b float32) float32 {
	return float32(math.Max(float64(a), float64(b)))
}

// ==================================== Compliance Score ====================================

// SetPostureReportComplianceScores calculates and populates scores for all controls, frameworks and whole scan.
func (su *ScoreUtil) SetPostureReportComplianceScores(report *v2.PostureReport) error {
	// call CalculatePostureReportV2 to set frameworks.score and summaryDetails.score for backward compatibility
	// afterwards we set complianceScore for frameworks, controls and summaryDetails
	// TODO: remove CalculatePostureReportV2 call after we deprecate frameworks.score and summaryDetails.score
	if err := su.CalculatePostureReportV2(report); err != nil {
		return err
	}
	// set compliance score for each framework
	for i := range report.SummaryDetails.Frameworks {
		// set compliance score for framework and all controls in framework
		report.SummaryDetails.Frameworks[i].ComplianceScore = su.GetFrameworkComplianceScore(&report.SummaryDetails.Frameworks[i])
	}
	// set compliance score per control
	sumScore := su.ControlsSummariesComplianceScore(&report.SummaryDetails.Controls, "")
	// set compliance score for whole scan
	summaryScore := float32(0)
	if len(report.SummaryDetails.Controls) > 0 {
		summaryScore = sumScore / float32(len(report.SummaryDetails.Controls))
	}
	report.SummaryDetails.ComplianceScore = summaryScore
	return nil
}

// ControlsSummariesComplianceScore sets the controls compliance score
// and returns the sum of all controls scores
func (su *ScoreUtil) ControlsSummariesComplianceScore(ctrls *reportsummary.ControlSummaries, frameworkName string) (sumScore float32) {
	for ctrlID := range *ctrls {
		ctrl := (*ctrls)[ctrlID]
		ctrl.ComplianceScore = shared.Ptr(su.GetControlComplianceScore(&ctrl, frameworkName))
		(*ctrls)[ctrlID] = ctrl
		sumScore += ctrl.GetComplianceScore()
	}
	return sumScore
}

// GetFrameworkComplianceScore returns the compliance score for a given framework (as a percentage)
// The framework compliance score is the average of all controls scores in that framework
func (su *ScoreUtil) GetFrameworkComplianceScore(framework *reportsummary.FrameworkSummary) float32 {
	frameworkScore := float32(0)
	sumScore := su.ControlsSummariesComplianceScore(&framework.Controls, framework.GetName())
	if len(framework.Controls) > 0 {
		frameworkScore = sumScore / float32(len(framework.Controls))
	}
	return frameworkScore
}

// GetControlComplianceScore returns the compliance score for a given control (as a percentage).
func (su *ScoreUtil) GetControlComplianceScore(ctrl reportsummary.IControlSummary, _ /*frameworkName*/ string) float32 {
	// If a control has status passed it should always be considered as having 100% compliance score
	if ctrl.GetStatus().IsPassed() {
		return 100
	}

	resourcesIDs := ctrl.ListResourcesIDs(nil)
	numOfPassedResources := resourcesIDs.Passed()
	numOfAllResources := resourcesIDs.Len()

	if numOfAllResources > 0 {
		return (float32(numOfPassedResources) / float32(numOfAllResources)) * 100
	}
	return 0
}
