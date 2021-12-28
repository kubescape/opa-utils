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

	failedControls := mockResultsFailed.ListFailedControls(nil)
	for i := range failedControls {
		controls[failedControls[i]] = reportsummary.ControlSummary{}
	}

	// New control
	updateControlsSummaryCounters(&mockResultsFailed, controls, nil)
	for _, v := range controls {
		assert.Equal(t, 1, v.NumberOfAll())
		assert.Equal(t, 1, v.NumberOfFailed())
		assert.Equal(t, 0, v.NumberOfPassed())
		assert.Equal(t, 0, v.NumberOfSkipped())
		assert.Equal(t, 0, v.NumberOfExcluded())
	}

}
func TestUpdateControlsSummaryCountersPassed(t *testing.T) {
	controls := map[string]reportsummary.ControlSummary{}

	passedControls := mockResultsFailed.ListPassedControls(nil)
	for i := range passedControls {
		controls[passedControls[i]] = reportsummary.ControlSummary{}
	}

	// New control
	updateControlsSummaryCounters(&mockResultsPassed, controls, nil)
	for _, v := range controls {
		assert.Equal(t, 1, v.NumberOfAll())
		assert.Equal(t, 1, v.NumberOfPassed())
		assert.Equal(t, 0, v.NumberOfFailed())
		assert.Equal(t, 0, v.NumberOfSkipped())
		assert.Equal(t, 0, v.NumberOfExcluded())
	}
}

func TestUpdateControlsSummaryCountersAll(t *testing.T) {
	controls := map[string]reportsummary.ControlSummary{}

	allControls := mockResultsFailed.ListAllControls(nil)
	tt := allControls.ListAll()
	for i := range tt {
		controls[tt[i]] = reportsummary.ControlSummary{}
	}

	updateControlsSummaryCounters(&mockResultsFailed, controls, nil)
	for _, i := range allControls.ListFailed() {
		v, k := controls[i]
		assert.True(t, k)
		assert.NotEqual(t, 0, v.NumberOfAll())
		assert.NotEqual(t, 0, v.NumberOfFailed())
		assert.Equal(t, 0, v.NumberOfPassed())
		assert.Equal(t, 0, v.NumberOfSkipped())
		assert.Equal(t, 0, v.NumberOfExcluded())
	}
	for _, i := range allControls.ListPassed() {
		v, k := controls[i]
		assert.True(t, k)
		assert.NotEqual(t, 0, v.NumberOfAll())
		assert.NotEqual(t, 0, v.NumberOfPassed())
		assert.Equal(t, 0, v.NumberOfFailed())
		assert.Equal(t, 0, v.NumberOfSkipped())
		assert.Equal(t, 0, v.NumberOfExcluded())
	}
}
