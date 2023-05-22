package resourcesresults

import (
	"testing"

	"github.com/kubescape/opa-utils/reporthandling/apis"
	"github.com/stretchr/testify/assert"
)

func TestSetGetResourceID(t *testing.T) {
	r := Result{}
	id := "my/id"
	r.SetResourceID(id)
	assert.Equal(t, id, r.GetResourceID())
}

func TestResultStatus(t *testing.T) {
	r := mockResultFailed()
	assert.Equal(t, apis.StatusFailed, r.GetStatus(nil).Status())
	assert.True(t, r.GetStatus(nil).IsFailed())
	assert.False(t, r.GetStatus(nil).IsPassed())
	assert.False(t, r.GetStatus(nil).IsSkipped())

	r2 := mockResultPassed()
	assert.Equal(t, apis.StatusPassed, r2.GetStatus(nil).Status())
	assert.True(t, r2.GetStatus(nil).IsPassed())
	assert.False(t, r2.GetStatus(nil).IsFailed())
	assert.False(t, r2.GetStatus(nil).IsSkipped())

}

func TestResultList(t *testing.T) {
	r := mockResultFailed()
	assert.NotEqual(t, 0, r.ListControlsIDs(nil).Len())
	assert.NotEqual(t, 0, r.ListControlsIDs(nil).Failed())
	assert.NotEqual(t, 0, r.ListControlsIDs(nil).Passed())

	r3 := mockResultPassed()
	assert.NotEqual(t, 0, r3.ListControlsIDs(nil).Len())
	assert.NotEqual(t, 0, r3.ListControlsIDs(nil).Passed())
	assert.Equal(t, 0, r3.ListControlsIDs(nil).Failed())
}

func TestListRulesOfControl(t *testing.T) {
	r := mockResultFailed()
	assert.Equal(t, 3, len(r.ListRulesOfControl("", "")))
	controlNames := r.ListControlsNames(nil).All()
	assert.Containsf(t, controlNames, "0087", "expected to find 0087 in %v", controlNames)
	assert.Containsf(t, controlNames, "0088", "expected to find 0088 in %v", controlNames)
	assert.Containsf(t, controlNames, "0089", "expected to find 0089 in %v", controlNames)
	assert.Equalf(t, 2, len(r.ListRulesOfControl("", "0087")), "expected to find 2 rules for control named 0087")
	assert.Equalf(t, 1, len(r.ListRulesOfControl("", "0088")), "expected to find 1 rule for control named 0088")
	assert.Equalf(t, 1, len(r.ListRulesOfControl("", "0089")), "expected to find 1 rule for control named 0089")

	controlIds := r.ListControlsIDs(nil).All()
	assert.Containsf(t, controlIds, "C-0087", "expected to find C-0087 in %v", controlIds)
	assert.Containsf(t, controlIds, "C-0088", "expected to find C-0088 in %v", controlIds)
	assert.Containsf(t, controlIds, "C-0089", "expected to find C-0089 in %v", controlIds)
	assert.Equalf(t, 2, len(r.ListRulesOfControl("C-0087", "")), "expected to find 2 rules for control id C-0087")
	assert.Equalf(t, 1, len(r.ListRulesOfControl("C-0088", "")), "expected to find 1 rule for control id C-0088")
	assert.Equalf(t, 1, len(r.ListRulesOfControl("C-0089", "")), "expected to find 1 rule for control id C-0089")

}
