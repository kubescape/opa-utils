package score

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/kubescape/k8s-interface/workloadinterface"
	armoupautils "github.com/kubescape/opa-utils/objectsenvelopes"
	"github.com/kubescape/opa-utils/reporthandling/results/v1/reportsummary"
	v2 "github.com/kubescape/opa-utils/reporthandling/v2"
	"go.uber.org/zap"
	appsv1 "k8s.io/api/apps/v1"

	// corev1 "k8s.io/api/core/v1"
	k8sinterface "github.com/kubescape/k8s-interface/k8sinterface"
	"github.com/kubescape/opa-utils/reporthandling"
)

const (
	replicaFactor = 1.1
)

type ControlScoreWeights struct {
	BaseScore                    float32 `json:"baseScore"`
	RuntimeImprovementMultiplier float32 `json:"improvementRatio"`
}

type ScoreUtil struct {
	// ResourceTypeScores map[string]float32
	// FrameworksScore    map[string]map[string]ControlScoreWeights
	K8SApoObj   *k8sinterface.KubernetesApi
	resources   map[string]workloadinterface.IMetadata
	isDebugMode bool
	// configPath  string
}

var postureScore *ScoreUtil

func (su *ScoreUtil) Calculate(frameworksReports []reporthandling.FrameworkReport) error {
	for i := range frameworksReports {
		su.CalculateFrameworkScore(&frameworksReports[i])
	}

	return nil
}

func (su *ScoreUtil) CalculateFrameworkScore(framework *reporthandling.FrameworkReport) error {
	framework.Score = 0
	for i := range framework.ControlReports {
		framework.ControlReports[i].Score = 0
		wcsCtrl, unormalizedScore := su.ControlScore(&framework.ControlReports[i], framework.Name)
		if su.isDebugMode {
			fmt.Printf("control %s(%s) failed %v wcs %v (baseScore: %v)\n", framework.ControlReports[i].Name, framework.ControlReports[i].ControlID, unormalizedScore, wcsCtrl, framework.ControlReports[i].BaseScore)

		}
		framework.WCSScore += wcsCtrl

		framework.Score += unormalizedScore
		framework.ARMOImprovement += framework.ControlReports[i].ARMOImprovement
	}
	if framework.WCSScore == 0 {
		framework.Score = 0
		return fmt.Errorf("unable to calculate score for framework %s due to bad wcs score\n", framework.Name)
	}
	framework.Score = (framework.Score * 100) / framework.WCSScore
	if su.isDebugMode {
		fmt.Printf("framework %s score %v\n", framework.Name, framework.Score)
	}
	framework.ARMOImprovement = (framework.ARMOImprovement * 100) / framework.WCSScore
	return nil
}

/*
daemonset: daemonsetscore*#nodes
workloads: if replicas:

	             replicascore*workloadkindscore*#replicas
	           else:
			     regular
*/
func (su *ScoreUtil) GetScore(v map[string]interface{}) float32 {
	var score float32 = 1.0
	if k8sinterface.IsTypeWorkload(v) {
		wl := workloadinterface.NewWorkloadObj(v)
		score = su.processWorkload(wl, score, v)
	} else if armoupautils.IsTypeRegoResponseVector(v) {
		if vec := armoupautils.NewRegoResponseVectorObject(v); vec != nil {
			related := vec.GetRelatedObjects()
			for i := range related {
				if k8sinterface.IsTypeWorkload(related[i].GetObject()) {
					wl := workloadinterface.NewWorkloadObj(v)
					score = float32(math.Max(float64(score), float64(su.processWorkload(wl, score, v))))

				}
			}
		}
	}

	return score
}

func (*ScoreUtil) processWorkload(wl *workloadinterface.Workload, score float32, v map[string]interface{}) float32 {
	if wl != nil {
		replicas := wl.GetReplicas()
		if replicas > 1 {
			score *= float32(replicas) * replicaFactor
		}

		if strings.ToLower(wl.GetKind()) == "daemonset" {
			/*
				if n, ok := workloadinterface.InspectMap(v, "status", "desiredNumberScheduled"); ok {
					if desiredNumberScheduled, ok := n.(int32); ok && desiredNumberScheduled > 0 {
						score *= float32(desiredNumberScheduled)
					}
				}
			*/

			//TODO - replace marshal and unmarshal by map inspection look code above
			b, err := json.Marshal(v)
			if err == nil {
				dmnset := appsv1.DaemonSet{}
				json.Unmarshal(b, &dmnset)

				if dmnset.Status.DesiredNumberScheduled > 0 {
					score *= float32(dmnset.Status.DesiredNumberScheduled)

				}
			}
		}

	}
	return score
}

/*
ControlScore:
@input:
ctrlReport - reporthandling.ControlReport object, must contain down the line the Input resources and the output resources
frameworkName - calculate this control according to a given framework weights

ctrl.score = baseScore * SUM_resource (resourceWeight*min(#replicas*replicaweight,1)(nodes if daemonset)

returns wcsscore,ctrlscore(unnormalized)
*/
func (su *ScoreUtil) ControlScore(ctrlReport *reporthandling.ControlReport, frameworkName string) (float32, float32) {

	resourcesIDs := ctrlReport.ListResourcesIDs()
	failedResourceIDS := resourcesIDs.GetFailedResources()
	allResourcesIDS := resourcesIDs.GetAllResources()
	for i := range failedResourceIDS {
		if failedResourceIDS, ok := su.resources[failedResourceIDS[i]]; ok {
			ctrlReport.Score += su.GetScore(failedResourceIDS.GetObject())
		}
	}
	ctrlReport.Score *= ctrlReport.BaseScore

	var wcsScore float32 = 0
	for i := range allResourcesIDS {
		if allResourcesIDS, ok := su.resources[allResourcesIDS[i]]; ok {
			wcsScore += su.GetScore(allResourcesIDS.GetObject())
		}
	}
	if ctrlReport.Score != 0 {
		wcsScore *= ctrlReport.BaseScore
	} else {
		wcsScore = ctrlReport.BaseScore
	}

	//x
	unormalizedScore := ctrlReport.Score
	ctrlReport.ARMOImprovement = unormalizedScore * ctrlReport.ARMOImprovement
	if wcsScore > 0 {
		ctrlReport.Score = (ctrlReport.Score * 100) / wcsScore
	} else {
		//ctrlReport.Score = 0
		zap.L().Error("worst case scenario was 0, meaning no resources input were given - score is not available(will appear as > 1)")
	}
	return wcsScore, unormalizedScore

}

func NewScore(allResources map[string]workloadinterface.IMetadata) *ScoreUtil {
	if postureScore == nil {

		isDebugBool := false
		if strings.ToLower(os.Getenv("ARMO_DEBUG_MODE")) == "true" {
			isDebugBool = true
		}
		postureScore = &ScoreUtil{
			resources:   allResources,
			isDebugMode: isDebugBool,
		}

	}

	return postureScore
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
		if su.isDebugMode {
			fmt.Printf("framework %s score %v\n", report.SummaryDetails.Frameworks[i].GetName(), report.SummaryDetails.Frameworks[i].GetScore())
		}
	}

	totalUnormalizedScore, totalWcsScore := su.ControlsSummariesScore(&report.SummaryDetails.Controls, "")
	report.SummaryDetails.Score = (totalUnormalizedScore * 100) / totalWcsScore

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

// /*
// ControlScore:
// @input:
// IControlSummary - //assuming ListResourcesIDs() is functional @ this scope
// frameworkName - calculate this control according to a given framework weights

// ctrl.score = baseScore * SUM_resource (resourceWeight*min(#replicas*replicaweight,1)(nodes if daemonset)

// returns ctrlscore(normalized),ctrlscore(unnormalized),wcsscore,

// */
func (su *ScoreUtil) ControlV2Score(ctrl reportsummary.IControlSummary, frameworkName string) (ctrlScore float32, unormalizedScore float32, wcsScore float32) {
	ctrlScore = 0
	unormalizedScore = 0
	wcsScore = 0
	resourcesIDs := ctrl.ListResourcesIDs()
	failedResourceIDS := resourcesIDs.Failed()
	allResourcesIDSIter := resourcesIDs.All()

	for i := range failedResourceIDS {
		if _, ok := su.resources[failedResourceIDS[i]]; ok {
			unormalizedScore += su.GetScore(su.resources[failedResourceIDS[i]].GetObject())
		}

	}
	unormalizedScore *= ctrl.GetScoreFactor()

	for allResourcesIDSIter.HasNext() {
		resourceId := allResourcesIDSIter.Next()
		if _, ok := su.resources[resourceId]; ok {
			wcsScore += su.GetScore(su.resources[resourceId].GetObject())
		}
	}

	wcsScore *= ctrl.GetScoreFactor()
	// //x
	// ctrlReport.ARMOImprovement = unormalizedScore * ctrlReport.ARMOImprovement
	if wcsScore > 0 {
		ctrlScore = (unormalizedScore * 100) / wcsScore
	} else {
		//ctrlReport.Score = 0
		zap.L().Error("worst case scenario was 0, meaning no resources input were given - score is not available(will appear as > 1)")
	}
	return ctrlScore, unormalizedScore, wcsScore

}
