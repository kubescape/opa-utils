package score

import (
	"testing"

	k8sinterface "github.com/armosec/k8s-interface/k8sinterface"
	"github.com/armosec/k8s-interface/workloadinterface"
)

func TestReplicaScore(t *testing.T) {
	k8sinterface.InitializeMapResourcesMock()

	deployment := getResourceByType("deployment")
	if wl := workloadinterface.NewWorkloadObj(deployment); wl == nil || wl.GetReplicas() != 3 {
		t.Errorf("invalid wl was put into the test, should have 3 replicas %v", deployment)
	}

	s := ScoreUtil{}
	score := s.GetScore(deployment)
	if score > 3.3001 || score < 3.3000001 {
		t.Errorf("invalid score: %v should be 3.3~(numerical errrors considered) ", score)
	}
}

func TestDaemonScore(t *testing.T) {
	ds := getResourceByType("daemonset")
	s := ScoreUtil{}
	score := s.GetScore(ds)
	if score != 13 {
		t.Errorf("invalid score: %v should be 13 ", score)
	}
}

func TestInactiveDaemonScore(t *testing.T) {
	ds := getResourceByType("daemonset")
	tmp := ds["status"].(map[string]interface{})
	tmp["desiredNumberScheduled"] = 0
	ds["status"] = tmp
	s := ScoreUtil{}
	score := s.GetScore(ds)
	if score != 1 {
		t.Errorf("invalid score: %v should be 1 ", score)
	}
}

func TestFrameworkMock(t *testing.T) {
}
