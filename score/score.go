package score

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/armosec/k8s-interface/workloadinterface"
	armoupautils "github.com/armosec/opa-utils/objectsenvelopes"
	"go.uber.org/zap"
	appsv1 "k8s.io/api/apps/v1"

	// corev1 "k8s.io/api/core/v1"
	k8sinterface "github.com/armosec/k8s-interface/k8sinterface"
	"github.com/armosec/opa-utils/reporthandling"
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
	configPath  string
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
	if workloadinterface.IsTypeWorkload(v) {
		wl := workloadinterface.NewWorkloadObj(v)
		score = su.processWorkload(wl, score, v)
	} else if armoupautils.IsTypeRegoResponseVector(v) {
		if vec := armoupautils.NewRegoResponseVectorObject(v); vec != nil {
			related := vec.GetRelatedObjects()
			for i := range related {
				if workloadinterface.IsTypeWorkload(related[i].GetObject()) {
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

	failedResourceIDS := ctrlReport.ListResourcesIDs().GetFailedResources()
	allResourcesIDS := ctrlReport.ListResourcesIDs().GetAllResources()
	for i := range failedResourceIDS {
		ctrlReport.Score += su.GetScore(su.resources[failedResourceIDS[i]].GetObject())
	}
	ctrlReport.Score *= ctrlReport.BaseScore

	var wcsScore float32 = 0
	for i := range allResourcesIDS {
		wcsScore += su.GetScore(su.resources[allResourcesIDS[i]].GetObject())
	}

	wcsScore *= ctrlReport.BaseScore
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
