package reportsummary

import (
	"github.com/armosec/opa-utils/reporthandling/apis"
	helpersv1 "github.com/armosec/opa-utils/reporthandling/helpers/v1"
	"github.com/armosec/opa-utils/shared"
)

// =================================== Status ============================================

// IsPassed did the scan pass
func (summaryDetails *SummaryDetails) IsPassed(f *helpersv1.Filters) bool {
	return summaryDetails.Status(f) == apis.StatusPassed
}

// IsFailed did the scan fail
func (summaryDetails *SummaryDetails) IsFailed(f *helpersv1.Filters) bool {
	return summaryDetails.Status(f) == apis.StatusFailed
}

// IsExcluded is the scan excluded
func (summaryDetails *SummaryDetails) IsExcluded(f *helpersv1.Filters) bool {
	return summaryDetails.Status(f) == apis.StatusExcluded
}

// IsSkipped was the scan skipped
func (summaryDetails *SummaryDetails) IsSkipped(f *helpersv1.Filters) bool {
	return summaryDetails.Status(f) == apis.StatusSkipped
}

// Status get the scan status. returns an apis.ScanningStatus object
func (summaryDetails *SummaryDetails) Status(f *helpersv1.Filters) apis.ScanningStatus {
	status := apis.StatusSkipped

	if len(summaryDetails.Frameworks) == 0 { // if there is only a list of controls
		for _, controlSummary := range summaryDetails.Controls {
			status = apis.Compare(controlSummary.Status(), status)
		}
	} else {
		for i := range summaryDetails.Frameworks {
			if len(f.FrameworkNames) == 0 || shared.StringInSlice(f.FrameworkNames, summaryDetails.Frameworks[i].Name) {
				status = apis.Compare(summaryDetails.Frameworks[i].Status(), status)
			}
		}
	}
	return status
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
		s := summaryDetails.Frameworks[i].Status()
		if status == "" || summaryDetails.Frameworks[i].Status() == status {
			frameworks.Append(s, summaryDetails.Frameworks[i].GetName())
		}
	}
	return frameworks
}

// =========================================== List Controls ====================================

// ListExcludedResources list all excluded resources IDs
func (summaryDetails *SummaryDetails) ListExcludedControls(f *helpersv1.Filters) []string {
	return summaryDetails.ListControls(f, apis.StatusExcluded).ListExcluded()
}

// ListPassedResources list all passed resources IDs
func (summaryDetails *SummaryDetails) ListPassedControls(f *helpersv1.Filters) []string {
	return summaryDetails.ListControls(f, apis.StatusPassed).ListPassed()
}

// ListSkippedResources list all skipped resources IDs
func (summaryDetails *SummaryDetails) ListSkippedControls(f *helpersv1.Filters) []string {
	return summaryDetails.ListControls(f, apis.StatusSkipped).ListSkipped()
}

// ListFailedResources list all failed resources IDs
func (summaryDetails *SummaryDetails) ListFailedControls(f *helpersv1.Filters) []string {
	return summaryDetails.ListControls(f, apis.StatusFailed).ListFailed()
}

// ListAllResources list all resources IDs. This function lists the resources IDs from the "results" and not from the "resources"
func (summaryDetails *SummaryDetails) ListAllControls(f *helpersv1.Filters) *helpersv1.AllLists {
	return summaryDetails.ListControls(f, "")
}

func (summaryDetails *SummaryDetails) ListControls(f *helpersv1.Filters, status apis.ScanningStatus) *helpersv1.AllLists {
	controls := &helpersv1.AllLists{}
	if len(f.FrameworkNames) == 0 { // no filters applied
		for controlID, controlSummary := range summaryDetails.Controls {
			s := controlSummary.Status()
			if status == "" || s == status {
				controls.Append(s, controlID)
			}
		}
	} else { // list based on filters
		for i := range summaryDetails.Frameworks {
			if shared.StringInSlice(f.FrameworkNames, summaryDetails.Frameworks[i].Name) {
				for controlID, controlSummary := range summaryDetails.Controls {
					s := controlSummary.Status()
					if status == "" || s == status {
						controls.Append(s, controlID)
					}
				}
			}
		}
		controls.ToUnique() // remove ID duplications
	}
	return controls
}
