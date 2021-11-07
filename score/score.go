package score

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/armosec/k8s-interface/workloadinterface"
	"go.uber.org/zap"
	appsv1 "k8s.io/api/apps/v1"

	// corev1 "k8s.io/api/core/v1"
	k8sinterface "github.com/armosec/k8s-interface/k8sinterface"
	"github.com/armosec/opa-utils/reporthandling"
)

type ControlScoreWeights struct {
	BaseScore                    float32 `json:"baseScore"`
	RuntimeImprovementMultiplier float32 `json:"improvementRatio"`
}

type ScoreUtil struct {
	// ResourceTypeScores map[string]float32
	// FrameworksScore    map[string]map[string]ControlScoreWeights
	K8SApoObj  *k8sinterface.KubernetesApi
	configPath string
}

var postureScore *ScoreUtil

func (su *ScoreUtil) Calculate(frameworksReports []reporthandling.FrameworkReport) error {
	for i := range frameworksReports {
		su.CalculateFrameworkScore(&frameworksReports[i])
	}

	return nil
}

func (su *ScoreUtil) CalculateFrameworkScore(framework *reporthandling.FrameworkReport) error {
	for i := range framework.ControlReports {
		wcsCtrl, unormalizedScore := su.ControlScore(&framework.ControlReports[i], framework.Name)
		framework.WCSScore += wcsCtrl

		framework.Score += unormalizedScore
		framework.ARMOImprovement += framework.ControlReports[i].ARMOImprovement
	}
	if framework.WCSScore == 0 {
		framework.Score = 0
		return fmt.Errorf("unable to calculate score for framework %s due to bad wcs score", framework.Name)
	}
	framework.Score = (framework.Score * 100) / framework.WCSScore
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

	var score float32 = 1
	wl := workloadinterface.NewWorkloadObj(v)
	kind := ""
	if wl != nil {
		kind = strings.ToLower(wl.GetKind())
		replicas := wl.GetReplicas()
		if replicas > 1 {
			score *= float32(replicas)
		}

	} else {
		//TODO: external object
	}

	if kind == "daemonset" {
		b, err := json.Marshal(v)
		if err == nil {
			dmnset := appsv1.DaemonSet{}
			json.Unmarshal(b, &dmnset)

			if dmnset.Status.DesiredNumberScheduled > 0 {
				score *= float32(dmnset.Status.DesiredNumberScheduled)

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
	all, failed, _ := reporthandling.GetResourcesPerControl(ctrlReport)
	for i := range failed {
		ctrlReport.Score += ctrlReport.BaseScore * su.GetScore(failed[i])
	}
	var wcsScore float32 = 0
	for i := range all {
		wcsScore += ctrlReport.BaseScore * su.GetScore(all[i])
	}

	unormalizedScore := ctrlReport.Score
	ctrlReport.ARMOImprovement = unormalizedScore * ctrlReport.ARMOImprovement
	if wcsScore > 0 {
		ctrlReport.Score /= wcsScore
	} else {
		// ctrlReport.Score = 0
		zap.L().Error("worst case scenario was 0, meaning no resources input were given - score is not available(will appear as 0)")
	}
	return ctrlReport.BaseScore * wcsScore, unormalizedScore

}

func NewScore() *ScoreUtil {
	if postureScore == nil {

		postureScore = &ScoreUtil{}

	}

	return postureScore
}
