package reportsummary

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/kubescape/opa-utils/reporthandling/apis"
	"github.com/kubescape/opa-utils/reporthandling/results/v1/resourcesresults"
	"github.com/stretchr/testify/assert"
)

var mockResultsPassed = resourcesresults.MockResults()[0]
var mockResultsFailed = resourcesresults.MockResults()[1]
var mockResultsException = resourcesresults.MockResults()[2]
var mockResultsConfigurqation = resourcesresults.MockResults()[3]
var mockResultRequiresReview = resourcesresults.MockResults()[4]
var mockResultManualReview = resourcesresults.MockResults()[5]

func TestRuleStatus(t *testing.T) {
	r := mockSummaryDetailsFailed()
	r.CalculateStatus()

	assert.Equal(t, apis.StatusFailed, r.GetStatus().Status())
	assert.True(t, r.GetStatus().IsFailed())
	assert.False(t, r.GetStatus().IsPassed())
	assert.False(t, r.GetStatus().IsSkipped())

	r1 := mockSummaryDetailsException()
	r1.CalculateStatus()

	assert.Equal(t, apis.StatusPassed, r1.GetStatus().Status())
	assert.False(t, r1.GetStatus().IsFailed())
	assert.True(t, r1.GetStatus().IsPassed())
	assert.False(t, r1.GetStatus().IsSkipped())

	r2 := mockSummaryDetailsPassed()
	r2.CalculateStatus()

	assert.Equal(t, apis.StatusPassed, r2.GetStatus().Status())
	assert.True(t, r2.GetStatus().IsPassed())
	assert.False(t, r2.GetStatus().IsFailed())
	assert.False(t, r2.GetStatus().IsSkipped())

}

func TestRuleListing(t *testing.T) {
	r := mockSummaryDetailsFailed()
	assert.NotEqual(t, 0, r.ListFrameworksNames().Len())
	assert.NotEqual(t, 0, r.ListFrameworksNames().Failed())
	assert.NotEqual(t, 0, r.ListControlsNames().Failed())
	assert.NotEqual(t, 0, r.ListControlsIDs().Failed())
}

func TestUpdateControlsSummaryCountersFailed(t *testing.T) {
	controls := map[string]ControlSummary{}

	allControls := mockResultsFailed.ListControlsIDs(nil).All()
	for ctrlId, status := range allControls {
		if status == apis.StatusFailed {
			controls[ctrlId] = ControlSummary{}
		}
	}

	// New control
	updateControlsSummaryCounters(&mockResultsFailed, controls, nil)
	for _, v := range controls {
		assert.Equal(t, 1, v.NumberOfResources().All())
		assert.Equal(t, 1, v.NumberOfResources().Failed())
		assert.Equal(t, 0, v.NumberOfResources().Passed())
		assert.Equal(t, 0, v.NumberOfResources().Skipped())
	}

	for _, v := range controls {
		statuses, subStatuses := v.StatusesCounters()
		assert.Equal(t, 1, statuses.All())
		assert.Equal(t, 1, statuses.Failed())
		assert.Equal(t, 0, statuses.Passed())
		assert.Equal(t, 0, statuses.Skipped())
		assert.Equal(t, 0, subStatuses.All())
		assert.Equal(t, 0, subStatuses.Ignored())
	}

}

func TestUpdateControlsSummaryCountersExcluded(t *testing.T) {
	controls := map[string]ControlSummary{}

	allControls := mockResultsFailed.ListControlsIDs(nil).All()
	for ctrlId, status := range allControls {
		if status == apis.StatusPassed {
			controls[ctrlId] = ControlSummary{}
		}
	}

	updateControlsSummaryCounters(&mockResultsException, controls, nil)
	for _, v := range controls {
		statuses, subStatuses := v.StatusesCounters()
		assert.Equal(t, 1, statuses.All())
		assert.Equal(t, 0, statuses.Failed())
		assert.Equal(t, 1, statuses.Passed())
		assert.Equal(t, 0, statuses.Skipped())
		assert.Equal(t, 1, subStatuses.All())
		assert.Equal(t, 1, subStatuses.Ignored())
	}

}
func TestUpdateControlsSummaryCountersPassed(t *testing.T) {
	controls := map[string]ControlSummary{}

	allControls := mockResultsFailed.ListControlsIDs(nil).All()
	for ctrlId, status := range allControls {
		if status == apis.StatusPassed {
			controls[ctrlId] = ControlSummary{}
		}
	}

	// New control
	updateControlsSummaryCounters(&mockResultsPassed, controls, nil)
	for _, v := range controls {
		assert.Equal(t, 1, v.NumberOfResources().All())
		assert.Equal(t, 0, v.NumberOfResources().Failed())
		assert.Equal(t, 1, v.NumberOfResources().Passed())
		assert.Equal(t, 0, v.NumberOfResources().Skipped())

	}
}

func TestUpdateControlsSummaryCountersException(t *testing.T) {
	controls := map[string]ControlSummary{}

	allControls := mockResultsException.ListControlsIDs(nil).All()
	for ctrlId := range allControls {
		controls[ctrlId] = ControlSummary{}
	}

	// New control
	updateControlsSummaryCounters(&mockResultsException, controls, nil)
	for _, v := range controls {
		assert.Equal(t, 1, v.NumberOfResources().All())
		assert.Equal(t, 0, v.NumberOfResources().Failed())
		assert.Equal(t, 1, v.NumberOfResources().Passed())
		assert.Equal(t, 0, v.NumberOfResources().Skipped())
		assert.Equal(t, apis.SubStatusException, v.GetSubStatus())
	}
}

func TestUpdateControlsSummaryCountersConfiguration(t *testing.T) {
	controls := map[string]ControlSummary{}

	allControls := mockResultsConfigurqation.ListControlsIDs(nil).All()
	for ctrlId := range allControls {
		controls[ctrlId] = ControlSummary{}
	}

	// New control
	updateControlsSummaryCounters(&mockResultsConfigurqation, controls, nil)
	for _, v := range controls {
		assert.Equal(t, 1, v.NumberOfResources().All())
		assert.Equal(t, 0, v.NumberOfResources().Failed())
		assert.Equal(t, 0, v.NumberOfResources().Passed())
		assert.Equal(t, 1, v.NumberOfResources().Skipped())
		assert.Equal(t, apis.SubStatusConfiguration, v.GetSubStatus())
	}
}

func TestUpdateControlsSummaryCountersRequiresReview(t *testing.T) {
	controls := map[string]ControlSummary{}

	allControls := mockResultRequiresReview.ListControlsIDs(nil).All()
	for ctrlId := range allControls {
		controls[ctrlId] = ControlSummary{}
	}

	// New control
	updateControlsSummaryCounters(&mockResultRequiresReview, controls, nil)
	for _, v := range controls {
		assert.Equal(t, 1, v.NumberOfResources().All())
		assert.Equal(t, 0, v.NumberOfResources().Failed())
		assert.Equal(t, 0, v.NumberOfResources().Passed())
		assert.Equal(t, 1, v.NumberOfResources().Skipped())
		assert.Equal(t, apis.SubStatusRequiresReview, v.GetSubStatus())
	}
}

func TestUpdateControlsSummaryCountersManualReview(t *testing.T) {
	controls := map[string]ControlSummary{}

	allControls := mockResultManualReview.ListControlsIDs(nil).All()
	for ctrlId := range allControls {
		controls[ctrlId] = ControlSummary{}
	}

	// New control
	updateControlsSummaryCounters(&mockResultManualReview, controls, nil)
	for _, v := range controls {
		assert.Equal(t, 1, v.NumberOfResources().All())
		assert.Equal(t, 0, v.NumberOfResources().Failed())
		assert.Equal(t, 0, v.NumberOfResources().Passed())
		assert.Equal(t, 1, v.NumberOfResources().Skipped())
		assert.Equal(t, apis.SubStatusManualReview, v.GetSubStatus())
	}
}

func TestUpdateControlsSummaryCountersAll(t *testing.T) {
	controls := map[string]ControlSummary{}

	allControls := mockResultsFailed.ListControlsIDs(nil).All()
	for ctrlId := range allControls {
		controls[ctrlId] = ControlSummary{}
	}

	updateControlsSummaryCounters(&mockResultsFailed, controls, nil)
	for ctrlId, status := range allControls {
		if status == apis.StatusFailed {
			v, k := controls[ctrlId]
			assert.True(t, k)
			assert.NotEqual(t, 0, v.NumberOfResources().All())
			assert.NotEqual(t, 0, v.NumberOfResources().Failed())
			assert.Equal(t, 0, v.NumberOfResources().Passed())
			assert.Equal(t, 0, v.NumberOfResources().Skipped())
		}

	}
	for ctrlId, status := range allControls {
		if status == apis.StatusPassed {
			v, k := controls[ctrlId]
			assert.True(t, k)
			assert.NotEqual(t, 0, v.NumberOfResources().All())
			assert.NotEqual(t, 0, v.NumberOfResources().Passed())
			assert.Equal(t, 0, v.NumberOfResources().Failed())
			assert.Equal(t, 0, v.NumberOfResources().Skipped())
		}
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
					CriticalSeverityCounter: 1,
					HighSeverityCounter:     2,
					MediumSeverityCounter:   3,
					LowSeverityCounter:      4,
				},
			},
			want: fields{
				SeverityCounters: SeverityCounters{
					CriticalSeverityCounter: 1,
					HighSeverityCounter:     2,
					MediumSeverityCounter:   3,
					LowSeverityCounter:      4,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sc := &SummaryDetails{
				ResourcesSeverityCounters: tt.fields.SeverityCounters,
			}

			if got := sc.ResourcesSeverityCounters.NumberOfCriticalSeverity(); got != tt.want.SeverityCounters.CriticalSeverityCounter {
				t.Errorf("SeverityCounters.NumberOfResourcesWithCriticalSeverity() = %v, want %v", got, tt.want)
			}
			if got := sc.ResourcesSeverityCounters.NumberOfHighSeverity(); got != tt.want.SeverityCounters.HighSeverityCounter {
				t.Errorf("SeverityCounters.NumberOfResourcesWithCriticalSeverity() = %v, want %v", got, tt.want)
			}
			if got := sc.ResourcesSeverityCounters.NumberOfMediumSeverity(); got != tt.want.SeverityCounters.MediumSeverityCounter {
				t.Errorf("SeverityCounters.NumberOfResourcesWithCriticalSeverity() = %v, want %v", got, tt.want)
			}
			if got := sc.ResourcesSeverityCounters.NumberOfLowSeverity(); got != tt.want.SeverityCounters.LowSeverityCounter {
				t.Errorf("SeverityCounters.NumberOfResourcesWithCriticalSeverity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSummaryDetails_GetControlsSeverityCounters(t *testing.T) {
	type fields struct {
		ControlsSeverityCounters SeverityCounters
	}
	tests := []struct {
		want   ISeverityCounters
		name   string
		fields fields
	}{
		{
			name: "Controls severities",
			want: &SeverityCounters{
				CriticalSeverityCounter: 1,
				HighSeverityCounter:     2,
				MediumSeverityCounter:   3,
				LowSeverityCounter:      4,
			},
			fields: fields{
				ControlsSeverityCounters: SeverityCounters{
					CriticalSeverityCounter: 1,
					HighSeverityCounter:     2,
					MediumSeverityCounter:   3,
					LowSeverityCounter:      4,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			summaryDetails := &SummaryDetails{
				ControlsSeverityCounters: tt.fields.ControlsSeverityCounters,
			}
			if got := summaryDetails.GetControlsSeverityCounters().NumberOfCriticalSeverity(); got != tt.want.NumberOfCriticalSeverity() {
				t.Errorf("NumberOfCriticalSeverity() = %v, want %v", got, tt.want.NumberOfCriticalSeverity())
			}
			if got := summaryDetails.GetControlsSeverityCounters().NumberOfHighSeverity(); got != tt.want.NumberOfHighSeverity() {
				t.Errorf("NumberOfHighSeverity() = %v, want %v", got, tt.want)
			}
			if got := summaryDetails.GetControlsSeverityCounters().NumberOfMediumSeverity(); got != tt.want.NumberOfMediumSeverity() {
				t.Errorf("NumberOfMediumSeverity() = %v, want %v", got, tt.want)
			}
			if got := summaryDetails.GetControlsSeverityCounters().NumberOfLowSeverity(); got != tt.want.NumberOfLowSeverity() {
				t.Errorf("NumberOfLowSeverity() = %v, want %v", got, tt.want)
			}
		})
	}
}

//go:embed testdata/summaryDetails.json
var summaryDetailsBytes []byte

//go:embed testdata/allResourcesResults.json
var allResourcesResultsBytes []byte

func setUpSummaryDetails() (*SummaryDetails, error) {
	summaryDetails := &SummaryDetails{}
	if err := json.Unmarshal(summaryDetailsBytes, summaryDetails); err != nil {
		return nil, fmt.Errorf("failed to unmarshal summaryDetailsBytes: %v", err)
	}

	allResourcesResults := map[string]resourcesresults.Result{}
	if err := json.Unmarshal(allResourcesResultsBytes, &allResourcesResults); err != nil {
		return nil, fmt.Errorf("failed to unmarshal allResourcesResults: %v", err)
	}

	for i := range allResourcesResults {
		t := allResourcesResults[i]
		summaryDetails.AppendResourceResult(&t)
	}

	summaryDetails.InitResourcesSummary(nil)

	return summaryDetails, nil
}
func TestSummaryDetails_Counters(t *testing.T) {

	summaryDetails, err := setUpSummaryDetails()
	if err != nil {
		t.Fatalf("failed to unmarshal allResourcesResults: %v", err)
	}

	// testing counters
	assert.Equal(t, 93, summaryDetails.StatusCounters.All())
	assert.Equal(t, 4, summaryDetails.StatusCounters.Passed())
	assert.Equal(t, 9, summaryDetails.StatusCounters.Failed())
	assert.Equal(t, 80, summaryDetails.StatusCounters.Skipped())

	assert.Equal(t, 93, summaryDetails.NumberOfResources().All())
	assert.Equal(t, 4, summaryDetails.NumberOfResources().Passed())
	assert.Equal(t, 9, summaryDetails.NumberOfResources().Failed())
	assert.Equal(t, 80, summaryDetails.NumberOfResources().Skipped())

	assert.Equal(t, 0, summaryDetails.GetControlsSeverityCounters().NumberOfCriticalSeverity())
	assert.Equal(t, 3, summaryDetails.GetControlsSeverityCounters().NumberOfHighSeverity())
	assert.Equal(t, 1, summaryDetails.GetControlsSeverityCounters().NumberOfMediumSeverity())
	assert.Equal(t, 0, summaryDetails.GetControlsSeverityCounters().NumberOfLowSeverity())

	assert.Equal(t, 0, summaryDetails.GetResourcesSeverityCounters().NumberOfCriticalSeverity())
	assert.Equal(t, 20, summaryDetails.GetResourcesSeverityCounters().NumberOfHighSeverity())
	assert.Equal(t, 8, summaryDetails.GetResourcesSeverityCounters().NumberOfMediumSeverity())
	assert.Equal(t, 0, summaryDetails.GetResourcesSeverityCounters().NumberOfLowSeverity())

	assert.Equal(t, 27, summaryDetails.NumberOfControls().All())
	assert.Equal(t, 22, summaryDetails.NumberOfControls().Passed())
	assert.Equal(t, 4, summaryDetails.NumberOfControls().Failed())
	assert.Equal(t, 1, summaryDetails.NumberOfControls().Skipped())
}

func TestSummaryDetails_UniqueControls(t *testing.T) {

	summaryDetails, err := setUpSummaryDetails()
	if err != nil {
		t.Fatalf("failed to unmarshal allResourcesResults: %v", err)
	}
	m := map[string]interface{}{}
	for _, c := range summaryDetails.ListControls() {
		m[c.GetID()] = nil
	}

	assert.Equal(t, len(summaryDetails.ListControls()), len(m))

}

func TestSummaryDetails_UniqueFrameworks(t *testing.T) {

	summaryDetails, err := setUpSummaryDetails()
	if err != nil {
		t.Fatalf("failed to unmarshal allResourcesResults: %v", err)
	}
	m := map[string]interface{}{}
	for _, c := range summaryDetails.ListFrameworks() {
		m[c.GetName()] = nil
	}

	assert.Equal(t, len(summaryDetails.ListFrameworks()), len(m))

}

func TestSummaryDetails_UniqueResources(t *testing.T) {

	summaryDetails, err := setUpSummaryDetails()
	if err != nil {
		t.Fatalf("failed to unmarshal allResourcesResults: %v", err)
	}

	m := map[string]interface{}{}
	for rId := range summaryDetails.ListResourcesIDs(nil).All() {
		m[rId] = nil
	}

	assert.Equal(t, summaryDetails.ListResourcesIDs(nil).Len(), len(m))

}

//go:embed testdata/initSummaryDetails.json
var initSummaryDetailsBytes []byte

//go:embed testdata/resourcesResult.json
var resourcesResultBytes []byte

func TestSummaryDetails_AppendResourceResult(t *testing.T) {

	summaryDetails := &SummaryDetails{}
	if err := json.Unmarshal(initSummaryDetailsBytes, summaryDetails); err != nil {
		t.Fatalf("failed to unmarshal initSummaryDetailsBytes: %v", err)
	}

	resourcesResult := &resourcesresults.Result{}
	if err := json.Unmarshal(resourcesResultBytes, resourcesResult); err != nil {
		t.Fatalf("failed to unmarshal resourcesResultBytes: %v", err)
	}
	summaryDetails.AppendResourceResult(resourcesResult)

	// Test framework status
	fw := summaryDetails.Frameworks[0]

	assert.Equal(t, 1, fw.StatusCounters.All())
	assert.Equal(t, 0, fw.StatusCounters.Passed())
	assert.Equal(t, 0, fw.StatusCounters.Failed())
	assert.Equal(t, 1, fw.StatusCounters.Skipped())

	assert.Truef(t, fw.GetStatus().IsSkipped(), "framework status is \"%s\"", fw.GetStatus().Status())
}

func TestUpdateControlsSummaryCounters(t *testing.T) {

	tests := []struct {
		want      apis.IStatus
		controlID string
		name      string
	}{
		{
			name:      "Skipped control",
			controlID: "C-0012",
			want: &apis.StatusInfo{
				InnerStatus: apis.StatusSkipped,
				SubStatus:   apis.SubStatusConfiguration,
				InnerInfo:   "Control configurations are empty",
			},
		},
		{
			name:      "Passed control",
			controlID: "C-0057",
			want: &apis.StatusInfo{
				InnerStatus: apis.StatusPassed,
				SubStatus:   "",
				InnerInfo:   "",
			},
		},
	}

	summaryDetails := &SummaryDetails{}
	if err := json.Unmarshal(initSummaryDetailsBytes, summaryDetails); err != nil {
		t.Fatalf("failed to unmarshal initSummaryDetailsBytes: %v", err)
	}

	resourcesResult := &resourcesresults.Result{}
	if err := json.Unmarshal(resourcesResultBytes, resourcesResult); err != nil {
		t.Fatalf("failed to unmarshal resourcesResultBytes: %v", err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			summaryDetails := &SummaryDetails{}
			if err := json.Unmarshal(initSummaryDetailsBytes, summaryDetails); err != nil {
				t.Fatalf("failed to unmarshal initSummaryDetailsBytes: %v", err)
			}

			resourcesResult := &resourcesresults.Result{}
			if err := json.Unmarshal(resourcesResultBytes, resourcesResult); err != nil {
				t.Fatalf("failed to unmarshal resourcesResultBytes: %v", err)
			}

			updateControlsSummaryCounters(resourcesResult, summaryDetails.Controls, nil)

			if summaryDetails.Controls.GetControl(EControlCriteriaID, tt.controlID).GetStatus().Status() != tt.want.Status() {
				t.Errorf("Status() = %v, want %v", summaryDetails.Controls.GetControl(EControlCriteriaID, tt.controlID).GetStatus().Status(), tt.want.Status())
			}
			if summaryDetails.Controls.GetControl(EControlCriteriaID, tt.controlID).GetStatus().GetSubStatus() != tt.want.GetSubStatus() {
				t.Errorf("GetSubStatus() = %v, want %v", summaryDetails.Controls.GetControl(EControlCriteriaID, tt.controlID).GetStatus().GetSubStatus(), tt.want.GetSubStatus())
			}
			if summaryDetails.Controls.GetControl(EControlCriteriaID, tt.controlID).GetStatus().Info() != tt.want.Info() {
				t.Errorf("Info() = %v, want %v", summaryDetails.Controls.GetControl(EControlCriteriaID, tt.controlID).GetStatus().Info(), tt.want.Info())
			}
		})
	}

}
