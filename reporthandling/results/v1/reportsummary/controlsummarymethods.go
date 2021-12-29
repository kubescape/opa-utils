package reportsummary

import (
	"github.com/armosec/opa-utils/reporthandling/apis"
	helpersv1 "github.com/armosec/opa-utils/reporthandling/helpers/v1"
)

// =================================== Status ============================================

// GetStatus get the control status. returns an apis.ScanningStatus object
func (controlSummary *ControlSummary) GetStatus() apis.IStatus {
	if controlSummary.Status == apis.StatusUnknown {
		controlSummary.CalculateStatus()
	}
	return helpersv1.NewStatus(controlSummary.Status)
}

// CalculateStatus set the control status based on the resource counters
func (controlSummary *ControlSummary) CalculateStatus() {
	controlSummary.Status = calculateStatus(&controlSummary.ResourceCounters)
}

// =================================== Counters ============================================

// NumberOfExcluded get the number of excluded resources
func (controlSummary *ControlSummary) NumberOfExcluded() int {
	return controlSummary.ResourceCounters.NumberOfExcluded()
}

// NumberOfPassed get the number of passed resources
func (controlSummary *ControlSummary) NumberOfPassed() int {
	return controlSummary.ResourceCounters.NumberOfPassed()
}

// NumberOfSkipped get the number of skipped resources
func (controlSummary *ControlSummary) NumberOfSkipped() int {
	return controlSummary.ResourceCounters.NumberOfSkipped()
}

// NumberOfFailed get the number of failed resources
func (controlSummary *ControlSummary) NumberOfFailed() int {
	return controlSummary.ResourceCounters.NumberOfFailed()
}

// NumberOfAll get the number of all resources
func (controlSummary *ControlSummary) NumberOfAll() int {
	return controlSummary.ResourceCounters.NumberOfAll()
}

// Increase increases the counter based on the status
func (controlSummary *ControlSummary) Increase(status apis.IStatus) {
	controlSummary.ResourceCounters.Increase(status)
}

// =================================== Score ============================================

// GetScore return control score
func (controlSummary *ControlSummary) GetScore() float32 {
	return controlSummary.Score
}
