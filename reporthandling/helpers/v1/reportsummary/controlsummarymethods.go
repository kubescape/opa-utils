package reportsummary

import (
	"github.com/armosec/opa-utils/reporthandling/apis"
)

// =================================== Status ============================================

// IsPassed did this control pass
func (controlSummary *ControlSummary) IsPassed() bool {
	return controlSummary.ScanStatus == apis.StatusPassed
}

// IsFailed did this control fail
func (controlSummary *ControlSummary) IsFailed() bool {
	return controlSummary.ScanStatus == apis.StatusFailed
}

// IsExcluded is this control excluded
func (controlSummary *ControlSummary) IsExcluded() bool {
	return controlSummary.ScanStatus == apis.StatusExcluded
}

// IsSkipped was this control skipped
func (controlSummary *ControlSummary) IsSkipped() bool {
	return controlSummary.ScanStatus == apis.StatusSkipped
}

// Status get the control status. returns an apis.ScanningStatus object
func (controlSummary *ControlSummary) Status() apis.ScanningStatus {
	return controlSummary.ScanStatus
}

// setStatus set the control status. returns an apis.ScanningStatus object
func (controlSummary *ControlSummary) setStatus(s apis.ScanningStatus) {
	controlSummary.ScanStatus = s
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

// =================================== Setters ============================================

// setNumberOfFailed set the number of failed resources
func (controlSummary *ControlSummary) setNumberOfFailed(n int) {
	controlSummary.setNumberOfFailed(n)
}

// setNumberOfFailed set the number of passed resources
func (controlSummary *ControlSummary) setNumberOfPassed(n int) {
	controlSummary.setNumberOfPassed(n)
}

// setNumberOfFailed set the number of excluded resources
func (controlSummary *ControlSummary) setNumberOfExcluded(n int) {
	controlSummary.setNumberOfExcluded(n)
}

// setNumberOfFailed set the number of skipped resources
func (controlSummary *ControlSummary) setNumberOfSkipped(n int) {
	controlSummary.setNumberOfSkipped(n)
}

// =================================== Score ============================================

// GetScore return control score
func (controlSummary *ControlSummary) GetScore() float32 {
	return controlSummary.Score
}

// setScore set control score
func (controlSummary *ControlSummary) setScore(n float32) {
	controlSummary.Score = n
}
