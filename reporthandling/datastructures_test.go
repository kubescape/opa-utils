package reporthandling

import (
	"encoding/json"
	"testing"

	"github.com/kubescape/k8s-interface/workloadinterface"
	"github.com/stretchr/testify/assert"
)

func TestPostureReportWithK8SResource(t *testing.T) {

	expectedID := "apps/v1/default/Deployment/demoservice-server"
	report := MockPostureReportA()
	report.Resources = append(report.Resources, Resource{
		ResourceID: "test-id",
		Object:     workloadinterface.NewWorkloadMock(nil).GetObject(),
	})

	// t.Errorf(report.Resources[0].Object.GetID())

	a, e := json.Marshal(report)
	if e != nil {
		t.Errorf("failed to marshal the report: %v", e.Error())
	}

	report2 := PostureReport{}
	json.Unmarshal(a, &report2)

	id := report2.Resources[0].GetID()

	if id != expectedID {
		t.Errorf("unexpected id from custom object, given id: %s expected: %s", id, expectedID)
	}
}

func TestPostureReportWithExternalResource(t *testing.T) {
	expectedID := "//Subject/MySubject"
	report := MockPostureReportA()
	report.Resources = append(report.Resources, Resource{
		ResourceID: "test-id",
		Object: map[string]interface{}{
			"namespace":      "",
			"group":          "",
			"name":           "MySubject",
			"kind":           "Subject",
			"relatedObjects": nil,
			"failedCreteria": "RBAC",
		},
	})

	// t.Errorf(report.Resources[0].Object.GetID())

	a, e := json.Marshal(report)
	if e != nil {
		t.Errorf("failed to marshal the report: %v", e.Error())
	}

	report2 := PostureReport{}
	json.Unmarshal(a, &report2)

	assert.Equal(t, expectedID, report2.Resources[0].GetID())
}
func TestMockFrameworkA(t *testing.T) {
	policy := MockFrameworkA()
	bp, err := json.Marshal(policy)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%s\n", string(bp))
		// t.Errorf("%s\n", string(bp))
	}

}

func TestMockPostureReportA(t *testing.T) {
	policy := MockPostureReportA()
	bp, err := json.Marshal(policy)
	if err != nil {
		t.Error(err)
	} else {
		// t.Errorf("%s\n", string(bp))
		t.Logf("%s\n", string(bp))
	}

}
