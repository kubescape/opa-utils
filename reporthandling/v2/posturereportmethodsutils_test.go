package v2

import (
	"testing"

	"github.com/armosec/opa-utils/reporthandling/results/v1/reportsummary"
	"github.com/armosec/opa-utils/reporthandling/results/v1/resourcesresults"
	"github.com/stretchr/testify/assert"
)

var mockResultsPassed = resourcesresults.MockResults()[0]
var mockResultsFailed = resourcesresults.MockResults()[1]

func TestUpdateControlsSummaryCountersFailed(t *testing.T) {
	controls := map[string]reportsummary.ControlSummary{}

	failedControls := mockResultsFailed.ListControls(nil).Failed()
	for i := range failedControls {
		controls[failedControls[i]] = reportsummary.ControlSummary{}
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
	controls := map[string]reportsummary.ControlSummary{}

	passedControls := mockResultsFailed.ListControls(nil).Passed()
	for i := range passedControls {
		controls[passedControls[i]] = reportsummary.ControlSummary{}
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
	controls := map[string]reportsummary.ControlSummary{}

	allControls := mockResultsFailed.ListControls(nil)
	tt := allControls.All()
	for i := range tt {
		controls[tt[i]] = reportsummary.ControlSummary{}
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
