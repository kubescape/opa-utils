package reportsummary

import (
	"github.com/armosec/opa-utils/reporthandling/apis"
	helpersv1 "github.com/armosec/opa-utils/reporthandling/helpers/v1"
)

// =================================== Status ============================================

// Status get the scan status. returns an apis.ScanningStatus object
func (summaryDetails *SummaryDetails) GetStatus() *helpersv1.Status {
	if summaryDetails.Status == apis.StatusUnknown {
		summaryDetails.CalculateStatus()
	}
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
func (summaryDetails *SummaryDetails) NumberOf() ICounters {
	return &summaryDetails.ResourceCounters
}

// Increase increases the counter based on the status
func (summaryDetails *SummaryDetails) Increase(status apis.IStatus) {
	summaryDetails.ResourceCounters.Increase(status)
}

// =========================================== List Frameworks ====================================

// ListFrameworksNames list all framework names
func (summaryDetails *SummaryDetails) ListFrameworksNames() *helpersv1.AllLists {
	frameworks := &helpersv1.AllLists{}
	for i := range summaryDetails.Frameworks {
		frameworks.Append(summaryDetails.Frameworks[i].GetStatus().Status(), summaryDetails.Frameworks[i].GetName())
	}
	return frameworks
}

// ListFrameworks list all frameworks
func (summaryDetails *SummaryDetails) ListFrameworks() *ListPolicies {
	frameworks := ListPolicies{}
	for i := range summaryDetails.Frameworks {
		frameworks.Append(summaryDetails.Frameworks[i].GetStatus().Status(), &summaryDetails.Frameworks[i])
	}
	return &frameworks
}

// =========================================== List Controls ====================================

// ListControlsNames list all framework names
func (summaryDetails *SummaryDetails) ListControlsNames() *helpersv1.AllLists {
	controls := &helpersv1.AllLists{}
	for _, controlSummary := range summaryDetails.Controls {
		controls.Append(controlSummary.GetStatus().Status(), controlSummary.Name)
	}
	return controls
}

func (summaryDetails *SummaryDetails) ListControlsIDs() *helpersv1.AllLists {
	controls := &helpersv1.AllLists{}
	for controlID, controlSummary := range summaryDetails.Controls {
		controls.Append(controlSummary.GetStatus().Status(), controlID)
	}
	return controls
}

// ListFrameworks list all frameworks
func (summaryDetails *SummaryDetails) ListControls() *ListPolicies {
	controls := ListPolicies{}
	for i := range summaryDetails.Controls {
		control := summaryDetails.Controls[i]
		controls.Append(control.GetStatus().Status(), &control)
	}
	return &controls
}

// ================================================================================
func (summaryDetails *SummaryDetails) ControlName(controlID string) string {
	if c, ok := summaryDetails.Controls[controlID]; ok {
		return c.Name
	}
	return ""
}
