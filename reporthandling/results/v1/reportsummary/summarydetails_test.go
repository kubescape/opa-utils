package reportsummary

import (
	"testing"

	"github.com/kubescape/opa-utils/reporthandling/apis"
	"github.com/kubescape/opa-utils/reporthandling/results/v1/resourcesresults"
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
	assert.NotEqual(t, 0, r.ListFrameworksNames().All().Len())
	assert.NotEqual(t, 0, len(r.ListFrameworksNames().Failed()))
	assert.NotEqual(t, 0, len(r.ListControlsNames().Failed()))
	assert.NotEqual(t, 0, len(r.ListControlsIDs().Failed()))
}

func TestUpdateControlsSummaryCountersFailed(t *testing.T) {
	controls := map[string]ControlSummary{}

	failedControls := mockResultsFailed.ListControlsIDs(nil).Failed()
	for i := range failedControls {
		controls[failedControls[i]] = ControlSummary{}
	}

	// New control
	updateControlsSummaryCounters(&mockResultsFailed, controls, nil)
	for _, v := range controls {
		assert.Equal(t, 1, v.NumberOfResources().All())
		assert.Equal(t, 1, v.NumberOfResources().Failed())
		assert.Equal(t, 0, v.NumberOfResources().Passed())
		assert.Equal(t, 0, v.NumberOfResources().Skipped())
		assert.Equal(t, 0, v.NumberOfResources().Excluded())
	}

}
func TestUpdateControlsSummaryCountersPassed(t *testing.T) {
	controls := map[string]ControlSummary{}

	passedControls := mockResultsFailed.ListControlsIDs(nil).Passed()
	for i := range passedControls {
		controls[passedControls[i]] = ControlSummary{}
	}

	// New control
	updateControlsSummaryCounters(&mockResultsPassed, controls, nil)
	for _, v := range controls {
		assert.Equal(t, 1, v.NumberOfResources().All())
		assert.Equal(t, 1, v.NumberOfResources().Passed())
		assert.Equal(t, 0, v.NumberOfResources().Failed())
		assert.Equal(t, 0, v.NumberOfResources().Skipped())
		assert.Equal(t, 0, v.NumberOfResources().Excluded())
	}
}

func TestUpdateControlsSummaryCountersAll(t *testing.T) {
	controls := map[string]ControlSummary{}

	allControls := mockResultsFailed.ListControlsIDs(nil)
	tt := allControls.All()
	for tt.HasNext() {
		controls[tt.Next()] = ControlSummary{}
	}

	updateControlsSummaryCounters(&mockResultsFailed, controls, nil)
	for _, i := range allControls.Failed() {
		v, k := controls[i]
		assert.True(t, k)
		assert.NotEqual(t, 0, v.NumberOfResources().All())
		assert.NotEqual(t, 0, v.NumberOfResources().Failed())
		assert.Equal(t, 0, v.NumberOfResources().Passed())
		assert.Equal(t, 0, v.NumberOfResources().Skipped())
		assert.Equal(t, 0, v.NumberOfResources().Excluded())
	}
	for _, i := range allControls.Passed() {
		v, k := controls[i]
		assert.True(t, k)
		assert.NotEqual(t, 0, v.NumberOfResources().All())
		assert.NotEqual(t, 0, v.NumberOfResources().Passed())
		assert.Equal(t, 0, v.NumberOfResources().Failed())
		assert.Equal(t, 0, v.NumberOfResources().Skipped())
		assert.Equal(t, 0, v.NumberOfResources().Excluded())
	}
}

func TestSummaryDetails_GetResourcesSeverityCounters(t *testing.T) {
	type fields struct {
		SeverityCounters SeverityCounters
	}
	tests := []struct {
		name   string
		fields fields
		want   fields
	}{
		{
			name: "",
			fields: fields{
				SeverityCounters: SeverityCounters{
					ResourcesWithCriticalSeverityCounter: 1,
					ResourcesWithHighSeverityCounter:     2,
					ResourcesWithMediumSeverityCounter:   3,
					ResourcesWithLowSeverityCounter:      4,
				},
			},
			want: fields{
				SeverityCounters: SeverityCounters{
					ResourcesWithCriticalSeverityCounter: 1,
					ResourcesWithHighSeverityCounter:     2,
					ResourcesWithMediumSeverityCounter:   3,
					ResourcesWithLowSeverityCounter:      4,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sc := &SummaryDetails{
				SeverityCounters: tt.fields.SeverityCounters,
			}

			if got := sc.SeverityCounters.NumberOfResourcesWithCriticalSeverity(); got != tt.want.SeverityCounters.ResourcesWithCriticalSeverityCounter {
				t.Errorf("SeverityCounters.NumberOfResourcesWithCriticalSeverity() = %v, want %v", got, tt.want)
			}
			if got := sc.SeverityCounters.NumberOfResourcesWithHighSeverity(); got != tt.want.SeverityCounters.ResourcesWithHighSeverityCounter {
				t.Errorf("SeverityCounters.NumberOfResourcesWithCriticalSeverity() = %v, want %v", got, tt.want)
			}
			if got := sc.SeverityCounters.NumberOfResourcesWithMediumSeverity(); got != tt.want.SeverityCounters.ResourcesWithMediumSeverityCounter {
				t.Errorf("SeverityCounters.NumberOfResourcesWithCriticalSeverity() = %v, want %v", got, tt.want)
			}
			if got := sc.SeverityCounters.NumberOfResourcesWithLowSeverity(); got != tt.want.SeverityCounters.ResourcesWithLowSeverityCounter {
				t.Errorf("SeverityCounters.NumberOfResourcesWithCriticalSeverity() = %v, want %v", got, tt.want)
			}
		})
	}
}
