package score

import (
	"testing"

	"github.com/armosec/k8s-interface/workloadinterface"
	"github.com/armosec/opa-utils/reporthandling"
	"github.com/armosec/opa-utils/reporthandling/results/v1/reportsummary"
	"github.com/armosec/opa-utils/reporthandling/results/v1/resourcesresults"
	v2 "github.com/armosec/opa-utils/reporthandling/v2"
	"github.com/stretchr/testify/assert"
)

func TestReplicaScore(t *testing.T) {
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

func TestEmptyFrameworV2kMock(t *testing.T) {
	s := NewScore(map[string]workloadinterface.IMetadata{})
	report := &v2.PostureReport{
		Attributes:           []reportsummary.PostureAttributes{},
		CustomerGUID:         "",
		ClusterName:          "",
		ClusterCloudProvider: "",
		ReportID:             "",
		JobID:                "",
		PaginationInfo:       v2.PaginationMarks{},
		SummaryDetails:       reportsummary.SummaryDetails{Frameworks: []reportsummary.FrameworkSummary{{Name: "empty", Controls: reportsummary.ControlSummaries{}}}},
		Results:              []resourcesresults.Result{},
		Resources:            []reporthandling.Resource{},
	}
	err := s.CalculatePostureReportV2(report)

	if err == nil || report.SummaryDetails.Frameworks[0].Score != 0.0 {
		t.Errorf("empty framework should return an error and have score equals 0")
	}
}

func TestEmptyFrameworV1kMock(t *testing.T) {
	s := NewScore(map[string]workloadinterface.IMetadata{})
	report := reporthandling.PostureReport{FrameworkReports: []reporthandling.FrameworkReport{{
		Name:           "empty",
		ControlReports: []reporthandling.ControlReport{},
	}}}
	s.Calculate(report.FrameworkReports)
	assert.Equal(t, float32(0.0), report.FrameworkReports[0].Score, "empty framework should have score equals 0")

}
