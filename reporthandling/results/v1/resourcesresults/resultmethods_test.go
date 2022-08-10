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
	assert.False(t, r.GetStatus(nil).IsExcluded())
	assert.False(t, r.GetStatus(nil).IsSkipped())

	r2 := mockResultPassed()
	assert.Equal(t, apis.StatusPassed, r2.GetStatus(nil).Status())
	assert.True(t, r2.GetStatus(nil).IsPassed())
	assert.False(t, r2.GetStatus(nil).IsFailed())
	assert.False(t, r2.GetStatus(nil).IsExcluded())
	assert.False(t, r2.GetStatus(nil).IsSkipped())

}

func TestResultList(t *testing.T) {
	r := mockResultFailed()
	assert.NotEqual(t, 0, len(r.ListControlsIDs(nil).All()))
	assert.NotEqual(t, 0, len(r.ListControlsIDs(nil).Failed()))
	assert.NotEqual(t, 0, len(r.ListControlsIDs(nil).Passed()))
	assert.Equal(t, 0, len(r.ListControlsIDs(nil).Excluded()))

	r3 := mockResultPassed()
	assert.NotEqual(t, 0, len(r3.ListControlsIDs(nil).All()))
	assert.NotEqual(t, 0, len(r3.ListControlsIDs(nil).Passed()))
	assert.Equal(t, 0, len(r3.ListControlsIDs(nil).Excluded()))
	assert.Equal(t, 0, len(r3.ListControlsIDs(nil).Failed()))
}

func TestListRulesOfControl(t *testing.T) {
	r := mockResultFailed()
	assert.Equal(t, 3, len(r.ListRulesOfControl("", "")))
	assert.NotEqual(t, 0, len(r.ListRulesOfControl("", r.ListControlsNames(nil).All()[0])))
	assert.NotEqual(t, 3, len(r.ListRulesOfControl("", r.ListControlsNames(nil).All()[0])))
	assert.NotEqual(t, 0, len(r.ListRulesOfControl(r.ListControlsIDs(nil).All()[0], "")))
	assert.NotEqual(t, 3, len(r.ListRulesOfControl(r.ListControlsIDs(nil).All()[0], "")))
}
