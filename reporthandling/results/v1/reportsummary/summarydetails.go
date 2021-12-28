package reportsummary

import (
	"github.com/armosec/opa-utils/reporthandling/apis"
	helpersv1 "github.com/armosec/opa-utils/reporthandling/helpers/v1"
)

// =================================== Status ============================================

// Status get the scan status. returns an apis.ScanningStatus object
func (summaryDetails *SummaryDetails) GetStatus() *helpersv1.Status {
	return helpersv1.NewStatus(summaryDetails.Status)
}

// SetStatus set the framework status based on the resource counters
func (summaryDetails *SummaryDetails) CalculateStatus() {
	summaryDetails.Status = calculateStatus(&summaryDetails.ResourceCounters)
	for i := range summaryDetails.Frameworks {
		summaryDetails.Frameworks[i].CalculateStatus()
	}
	for k, c := range summaryDetails.Controls {
		c.CalculateStatus()
		summaryDetails.Controls[k] = c
	}
}

// =================================== Counters ============================================

// NumberOfExcluded get the number of excluded resources
func (summaryDetails *SummaryDetails) NumberOfExcluded() int {
	return summaryDetails.ResourceCounters.NumberOfExcluded()
}

// NumberOfPassed get the number of passed resources
func (summaryDetails *SummaryDetails) NumberOfPassed() int {
	return summaryDetails.ResourceCounters.NumberOfPassed()
}

// NumberOfSkipped get the number of skipped resources
func (summaryDetails *SummaryDetails) NumberOfSkipped() int {
	return summaryDetails.ResourceCounters.NumberOfSkipped()
}

// NumberOfFailed get the number of failed resources
func (summaryDetails *SummaryDetails) NumberOfFailed() int {
	return summaryDetails.ResourceCounters.NumberOfFailed()
}

// NumberOfAll get the number of all resources
func (summaryDetails *SummaryDetails) NumberOfAll() int {
	return summaryDetails.ResourceCounters.NumberOfAll()
}

// Increase increases the counter based on the status
func (summaryDetails *SummaryDetails) Increase(status apis.IStatus) {
	summaryDetails.ResourceCounters.Increase(status)
}

// =========================================== List Frameworks ====================================

// ListExcludedResources list all excluded resources IDs
func (summaryDetails *SummaryDetails) ListExcludedFrameworks() []string {
	return summaryDetails.ListFrameworks(apis.StatusExcluded).ListExcluded()
}

// ListPassedResources list all passed resources IDs
func (summaryDetails *SummaryDetails) ListPassedFrameworks() []string {
	return summaryDetails.ListFrameworks(apis.StatusPassed).ListPassed()
}

// ListSkippedResources list all skipped resources IDs
func (summaryDetails *SummaryDetails) ListSkippedFrameworks() []string {
	return summaryDetails.ListFrameworks(apis.StatusSkipped).ListSkipped()
}

// ListFailedResources list all failed resources IDs
func (summaryDetails *SummaryDetails) ListFailedFrameworks() []string {
	return summaryDetails.ListFrameworks(apis.StatusFailed).ListFailed()
}

// ListAllResources list all resources IDs. This function lists the resources IDs from the "results" and not from the "resources"
func (summaryDetails *SummaryDetails) ListAllFrameworks() *helpersv1.AllLists {
	return summaryDetails.ListFrameworks("")
}

func (summaryDetails *SummaryDetails) ListFrameworks(status apis.ScanningStatus) *helpersv1.AllLists {
	frameworks := &helpersv1.AllLists{}
	for i := range summaryDetails.Frameworks {
		s := summaryDetails.Frameworks[i].GetStatus().Status()
		if status == "" || summaryDetails.Frameworks[i].GetStatus().Status() == status {
			frameworks.Append(s, summaryDetails.Frameworks[i].GetName())
		}
	}
	return frameworks
}

// =========================================== List Controls ====================================

// ListExcludedResources list all excluded resources IDs
func (summaryDetails *SummaryDetails) ListExcludedControls() []string {
	return summaryDetails.ListControls(apis.StatusExcluded).ListExcluded()
}

// ListPassedResources list all passed resources IDs
func (summaryDetails *SummaryDetails) ListPassedControls(f *helpersv1.Filters) []string {
	return summaryDetails.ListControls(apis.StatusPassed).ListPassed()
}

// ListSkippedResources list all skipped resources IDs
func (summaryDetails *SummaryDetails) ListSkippedControls(f *helpersv1.Filters) []string {
	return summaryDetails.ListControls(apis.StatusSkipped).ListSkipped()
}

// ListFailedResources list all failed resources IDs
func (summaryDetails *SummaryDetails) ListFailedControls(f *helpersv1.Filters) []string {
	return summaryDetails.ListControls(apis.StatusFailed).ListFailed()
}

// ListAllResources list all resources IDs. This function lists the resources IDs from the "results" and not from the "resources"
func (summaryDetails *SummaryDetails) ListAllControls() *helpersv1.AllLists {
	return summaryDetails.ListControls("")
}

func (summaryDetails *SummaryDetails) ListControls(status apis.ScanningStatus) *helpersv1.AllLists {
	controls := &helpersv1.AllLists{}
	for controlID, controlSummary := range summaryDetails.Controls {
		s := controlSummary.GetStatus().Status()
		if status == apis.StatusUnknown || status == s {
			controls.Append(s, controlID)
		}
	}
	return controls
}
