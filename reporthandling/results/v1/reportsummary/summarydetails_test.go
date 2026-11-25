package reportsummary

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"maps"
	"slices"
	"testing"

	"github.com/armosec/armoapi-go/armotypes"
	"github.com/kubescape/k8s-interface/workloadinterface"
	"github.com/kubescape/opa-utils/reporthandling"
	"github.com/kubescape/opa-utils/reporthandling/apis"
	helpersv1 "github.com/kubescape/opa-utils/reporthandling/helpers/v1"
	"github.com/kubescape/opa-utils/reporthandling/results/v1/resourcesresults"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func loadSummaryDetails(tb testing.TB) *SummaryDetails {
	summaryDetails := &SummaryDetails{}
	require.NoError(tb, json.Unmarshal(summaryDetailsBytes, summaryDetails))
	return summaryDetails
}

// loadResults returns the fixture results sorted by ID, since sub-statuses depend on append order.
func loadResults(tb testing.TB) []*resourcesresults.Result {
	byID := map[string]resourcesresults.Result{}
	require.NoError(tb, json.Unmarshal(allResourcesResultsBytes, &byID))
	results := make([]*resourcesresults.Result, 0, len(byID))
	for _, id := range slices.Sorted(maps.Keys(byID)) {
		r := byID[id]
		results = append(results, &r)
	}
	return results
}

func TestSummaryDetails_AppendResourceResults_MatchesOneByOne(t *testing.T) {
	results := loadResults(t)

	oneByOne := loadSummaryDetails(t)
	for _, r := range results {
		oneByOne.AppendResourceResult(r)
	}
	batched := loadSummaryDetails(t)
	batched.AppendResourceResults(results...)

	assert.Equal(t, oneByOne, batched)
	assert.NotZero(t, batched.Frameworks[0].StatusCounters.All())

	empty := loadSummaryDetails(t)
	empty.AppendResourceResults()
	assert.Equal(t, loadSummaryDetails(t), empty)
}

// The frameworks share one pooled list, so a second, smaller framework must not count the first one's resources.
func TestSummaryDetails_AppendResourceResults_FrameworkCounters(t *testing.T) {
	summaryDetails := loadSummaryDetails(t)
	second := loadSummaryDetails(t).Frameworks[0]
	second.Name = "second"
	keep := slices.Sorted(maps.Keys(second.Controls))[0]
	maps.DeleteFunc(second.Controls, func(id string, _ ControlSummary) bool { return id != keep })
	summaryDetails.Frameworks = append(summaryDetails.Frameworks, second)

	summaryDetails.AppendResourceResults(loadResults(t)...)

	for _, fw := range summaryDetails.Frameworks {
		want := fw.StatusCounters
		want.Set(fw.ListResourcesIDs(nil))
		assert.Equal(t, want, fw.StatusCounters, fw.Name)
	}
	assert.NotEqual(t, summaryDetails.Frameworks[0].StatusCounters, summaryDetails.Frameworks[1].StatusCounters)
}

func BenchmarkSummaryDetails_AppendResourceResults(b *testing.B) {
	base := loadResults(b)
	for _, n := range []int{1000, 4000} {
		results := make([]*resourcesresults.Result, n)
		for i := range results {
			r := *base[i%len(base)]
			r.ResourceID = fmt.Sprintf("%s/%d", r.ResourceID, i)
			results[i] = &r
		}

		b.Run(fmt.Sprintf("one-by-one/%d", n), func(b *testing.B) {
			for range b.N {
				b.StopTimer()
				summaryDetails := loadSummaryDetails(b)
				b.StartTimer()
				for _, r := range results {
					summaryDetails.AppendResourceResult(r)
				}
			}
		})
		b.Run(fmt.Sprintf("batch/%d", n), func(b *testing.B) {
			for range b.N {
				b.StopTimer()
				summaryDetails := loadSummaryDetails(b)
				b.StartTimer()
				summaryDetails.AppendResourceResults(results...)
			}
		})
	}
}

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

func TestFrameworkScopedExceptionEndToEnd(t *testing.T) {
	newSummary := func() *SummaryDetails {
		newControl := func() ControlSummary {
			return ControlSummary{ControlID: "C-0034", Name: "shared control"}
		}
		return &SummaryDetails{
			Controls: ControlSummaries{"C-0034": newControl()},
			Frameworks: []FrameworkSummary{
				{Name: "NSA", Controls: ControlSummaries{"C-0034": newControl()}},
				{Name: "MITRE", Controls: ControlSummaries{"C-0034": newControl()}},
			},
		}
	}

	exception := armotypes.PostureExceptionPolicy{
		Actions: []armotypes.PostureExceptionPolicyActions{armotypes.Disable},
		PosturePolicies: []armotypes.PosturePolicy{{
			FrameworkName: "NSA", ControlID: "C-0034", RuleName: "R1",
		}},
	}
	result := resourcesresults.Result{
		ResourceID: "apps/v1/default/deployment/example",
		AssociatedControls: []resourcesresults.ResourceAssociatedControl{{
			ControlID: "C-0034",
			Status:    apis.StatusInfo{InnerStatus: apis.StatusFailed},
			ResourceAssociatedRules: []resourcesresults.ResourceAssociatedRule{{
				Name: "R1", Status: apis.StatusFailed,
			}},
		}},
	}
	result.SetExceptions(
		workloadinterface.NewWorkloadMock(nil),
		[]armotypes.PostureExceptionPolicy{exception},
		"",
		map[string]reporthandling.Control{"C-0034": {}},
	)

	assert.Equal(t, apis.StatusFailed, result.AssociatedControls[0].ResourceAssociatedRules[0].Status)
	assertResultView := func(t *testing.T, result *resourcesresults.Result, filters *helpersv1.Filters, expectedStatus apis.ScanningStatus, expectedSubStatus apis.ScanningSubStatus) {
		t.Helper()
		ruleStatus := result.AssociatedControls[0].ResourceAssociatedRules[0].GetStatus(filters)
		controlStatus := result.AssociatedControls[0].GetStatus(filters)
		resourceStatus := result.GetStatus(filters)
		for level, status := range map[string]apis.IStatus{
			"rule": ruleStatus, "control": controlStatus, "resource": resourceStatus,
		} {
			assert.Equal(t, expectedStatus, status.Status(), level)
			assert.Equal(t, expectedSubStatus, status.GetSubStatus(), level)
		}
	}

	nsaFilter := &helpersv1.Filters{FrameworkNames: []string{"NSA"}}
	mitreFilter := &helpersv1.Filters{FrameworkNames: []string{"MITRE"}}
	aggregateFilter := &helpersv1.Filters{FrameworkNames: []string{"NSA", "MITRE"}}
	assertResultView(t, &result, nsaFilter, apis.StatusPassed, apis.SubStatusException)
	assertResultView(t, &result, mitreFilter, apis.StatusFailed, apis.SubStatusUnknown)
	assertResultView(t, &result, aggregateFilter, apis.StatusFailed, apis.SubStatusException)

	payload, err := json.Marshal(result)
	assert.NoError(t, err)
	var roundTripped resourcesresults.Result
	assert.NoError(t, json.Unmarshal(payload, &roundTripped))
	assertResultView(t, &roundTripped, nsaFilter, apis.StatusPassed, apis.SubStatusException)
	assertResultView(t, &roundTripped, mitreFilter, apis.StatusFailed, apis.SubStatusUnknown)
	assertResultView(t, &roundTripped, aggregateFilter, apis.StatusFailed, apis.SubStatusException)

	summary := newSummary()
	summary.AppendResourceResult(&roundTripped)
	summary.InitResourcesSummary(nil)

	assertSummary := func(t *testing.T, summary *SummaryDetails) {
		t.Helper()
		aggregateControl := summary.Controls["C-0034"]
		aggregate := aggregateControl.GetStatus()
		assert.Equal(t, apis.StatusFailed, aggregate.Status())
		assert.Equal(t, apis.SubStatusException, aggregate.GetSubStatus())
		assert.Equal(t, apis.StatusFailed, summary.GetStatus().Status())
		assert.Equal(t, apis.SubStatusException, summary.GetStatus().GetSubStatus())

		nsaControl := summary.Frameworks[0].Controls["C-0034"]
		nsa := nsaControl.GetStatus()
		assert.Equal(t, apis.StatusPassed, nsa.Status())
		assert.Equal(t, apis.SubStatusException, nsa.GetSubStatus())
		assert.Equal(t, apis.StatusPassed, summary.Frameworks[0].GetStatus().Status())
		assert.Equal(t, apis.SubStatusException, summary.Frameworks[0].GetStatus().GetSubStatus())

		mitreControl := summary.Frameworks[1].Controls["C-0034"]
		mitre := mitreControl.GetStatus()
		assert.Equal(t, apis.StatusFailed, mitre.Status())
		assert.Equal(t, apis.SubStatusUnknown, mitre.GetSubStatus())
		assert.Equal(t, apis.StatusFailed, summary.Frameworks[1].GetStatus().Status())
		assert.Equal(t, apis.SubStatusUnknown, summary.Frameworks[1].GetStatus().GetSubStatus())
	}
	assertSummary(t, summary)

	summaryPayload, err := json.Marshal(summary)
	assert.NoError(t, err)
	var roundTrippedSummary SummaryDetails
	assert.NoError(t, json.Unmarshal(summaryPayload, &roundTrippedSummary))
	roundTrippedSummary.CalculateStatus()
	for i := range roundTrippedSummary.Frameworks {
		roundTrippedSummary.Frameworks[i].CalculateStatus()
	}
	assertSummary(t, &roundTrippedSummary)
}

func TestAggregateUsesOnlyFrameworksContainingControl(t *testing.T) {
	controlSummary := ControlSummary{ControlID: "C-0034", Name: "NSA-only control"}
	summary := &SummaryDetails{
		Controls: ControlSummaries{"C-0034": controlSummary},
		Frameworks: []FrameworkSummary{
			{Name: "NSA", Controls: ControlSummaries{"C-0034": controlSummary}},
			{Name: "MITRE", Controls: ControlSummaries{}},
		},
	}
	result := resourcesresults.Result{
		ResourceID: "apps/v1/default/deployment/example",
		AssociatedControls: []resourcesresults.ResourceAssociatedControl{{
			ControlID: "C-0034",
			Status:    apis.StatusInfo{InnerStatus: apis.StatusFailed},
			ResourceAssociatedRules: []resourcesresults.ResourceAssociatedRule{{
				Name:   "R1",
				Status: apis.StatusFailed,
				Exception: []armotypes.PostureExceptionPolicy{{
					Actions: []armotypes.PostureExceptionPolicyActions{armotypes.Disable},
					PosturePolicies: []armotypes.PosturePolicy{{
						FrameworkName: "NSA", ControlID: "C-0034", RuleName: "R1",
					}},
				}},
			}},
		}},
	}

	summary.AppendResourceResult(&result)
	summary.InitResourcesSummary(nil)

	aggregateControl := summary.Controls["C-0034"]
	assert.Equal(t, apis.StatusPassed, aggregateControl.GetStatus().Status())
	assert.Equal(t, apis.SubStatusException, aggregateControl.GetStatus().GetSubStatus())
	assert.Equal(t, apis.StatusPassed, summary.GetStatus().Status())
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
				InnerInfo:   "Control configurations are empty (docs: https://kubescape.io/docs/frameworks-and-controls/configuring-controls)",
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
