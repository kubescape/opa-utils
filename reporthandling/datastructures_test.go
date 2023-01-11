package reporthandling

import (
	"encoding/json"
	"testing"

	"github.com/kubescape/k8s-interface/workloadinterface"
	"github.com/kubescape/opa-utils/reporthandling/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPostureReportWithK8SResource(t *testing.T) {

	expectedID := "apps/v1/default/Deployment/demoservice-server"
	report := MockPostureReportA()
	report.Resources = append(report.Resources, Resource{
		ResourceID: "test-id",
		Object:     workloadinterface.NewWorkloadMock(nil).GetObject(),
	})

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
	}

}

func TestMockPostureReportA(t *testing.T) {
	policy := MockPostureReportA()
	bp, err := json.Marshal(policy)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%s\n", string(bp))
	}

}

const randomizedTests = 25

func TestMarshalStruct(t *testing.T) {
	t.Parallel()

t.Run("should marshal/unmarshal JSON",testMarshalDataStructure[PolicyRule]()
	t.Parallel()

	for n := 0; n < randomizedTests; n++ {
		testMarshalDataStructure[PolicyRule](t)
	}
}

func TestMarshalControl(t *testing.T) {
	t.Parallel()

	for n := 0; n < randomizedTests; n++ {
		testMarshalDataStructure[Control](t)
	}
}

func TestMarshalFramework(t *testing.T) {
	t.Parallel()

	for n := 0; n < randomizedTests; n++ {
		testMarshalDataStructure[Control](t)
	}
}

func TestMarshalAttackTrackCategories(t *testing.T) {
	t.Parallel()

	for n := 0; n < randomizedTests; n++ {
		testMarshalDataStructure[AttackTrackCategories](t)
	}
}

func testMarshalDataStructure[T any]() func ( *testing.T) {
	return func(t *testing.T) {
	m:= mock.MockData[T]()

	buf, err := json.Marshal(m)
	require.NoError(t, err)

	var target T
	require.NoError(t, json.Unmarshal(buf, &target))
	require.EqualValues(t, m, target)
	}
}
