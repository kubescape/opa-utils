package reportsummary

import (
	"testing"

	"github.com/armosec/opa-utils/reporthandling/apis"
	"github.com/armosec/opa-utils/reporthandling/results/v1/resourcesresults"
	"github.com/stretchr/testify/assert"
)

var mockResultsPassed = resourcesresults.MockResults()[0]
var mockResultsFailed = resourcesresults.MockResults()[1]

func TestRuleStatus(t *testing.T) {
	r := mockSummaryDetailsFailed()
	r.CalculateStatus()

	assert.Equal(t, apis.StatusFailed, r.GetStatus().Status())
	assert.True(t, r.GetStatus().IsFailed())
	assert.False(t, r.GetStatus().IsPassed())
	assert.False(t, r.GetStatus().IsExcluded())
	assert.False(t, r.GetStatus().IsSkipped())

	r1 := mockSummaryDetailsExcluded()
	r1.CalculateStatus()

	assert.Equal(t, apis.StatusExcluded, r1.GetStatus().Status())
	assert.True(t, r1.GetStatus().IsExcluded())
	assert.False(t, r1.GetStatus().IsFailed())
	assert.False(t, r1.GetStatus().IsPassed())
	assert.False(t, r1.GetStatus().IsSkipped())

	r2 := mockSummaryDetailsPassed()
	r2.CalculateStatus()

	assert.Equal(t, apis.StatusPassed, r2.GetStatus().Status())
	assert.True(t, r2.GetStatus().IsPassed())
	assert.False(t, r2.GetStatus().IsFailed())
	assert.False(t, r2.GetStatus().IsExcluded())
	assert.False(t, r2.GetStatus().IsSkipped())

}

func TestRuleListing(t *testing.T) {
	r := mockSummaryDetailsFailed()
	assert.NotEqual(t, 0, len(r.ListFrameworksNames().All()))
	assert.NotEqual(t, 0, len(r.ListFrameworksNames().Failed()))
	assert.NotEqual(t, 0, len(r.ListControlsNames().Failed()))
	assert.NotEqual(t, 0, len(r.ListControlsIDs().Failed()))
}

func TestUpdateControlsSummaryCountersFailed(t *testing.T) {
	controls := map[string]ControlSummary{}

	failedControls := mockResultsFailed.ListControls(nil).Failed()
	for i := range failedControls {
		controls[failedControls[i]] = ControlSummary{}
	}

	// New control
	updateControlsSummaryCounters(&mockResultsFailed, controls, nil)
	for _, v := range controls {
		assert.Equal(t, 1, v.NumberOf().All())
		assert.Equal(t, 1, v.NumberOf().Failed())
		assert.Equal(t, 0, v.NumberOf().Passed())
		assert.Equal(t, 0, v.NumberOf().Skipped())
		assert.Equal(t, 0, v.NumberOf().Excluded())
	}

}
func TestUpdateControlsSummaryCountersPassed(t *testing.T) {
	controls := map[string]ControlSummary{}

	passedControls := mockResultsFailed.ListControls(nil).Passed()
	for i := range passedControls {
		controls[passedControls[i]] = ControlSummary{}
	}

	// New control
	updateControlsSummaryCounters(&mockResultsPassed, controls, nil)
	for _, v := range controls {
		assert.Equal(t, 1, v.NumberOf().All())
		assert.Equal(t, 1, v.NumberOf().Passed())
		assert.Equal(t, 0, v.NumberOf().Failed())
		assert.Equal(t, 0, v.NumberOf().Skipped())
		assert.Equal(t, 0, v.NumberOf().Excluded())
	}
}

func TestUpdateControlsSummaryCountersAll(t *testing.T) {
	controls := map[string]ControlSummary{}

	allControls := mockResultsFailed.ListControls(nil)
	tt := allControls.All()
	for i := range tt {
		controls[tt[i]] = ControlSummary{}
	}

	updateControlsSummaryCounters(&mockResultsFailed, controls, nil)
	for _, i := range allControls.Failed() {
		v, k := controls[i]
		assert.True(t, k)
		assert.NotEqual(t, 0, v.NumberOf().All())
		assert.NotEqual(t, 0, v.NumberOf().Failed())
		assert.Equal(t, 0, v.NumberOf().Passed())
		assert.Equal(t, 0, v.NumberOf().Skipped())
		assert.Equal(t, 0, v.NumberOf().Excluded())
	}
	for _, i := range allControls.Passed() {
		v, k := controls[i]
		assert.True(t, k)
		assert.NotEqual(t, 0, v.NumberOf().All())
		assert.NotEqual(t, 0, v.NumberOf().Passed())
		assert.Equal(t, 0, v.NumberOf().Failed())
		assert.Equal(t, 0, v.NumberOf().Skipped())
		assert.Equal(t, 0, v.NumberOf().Excluded())
	}
}
