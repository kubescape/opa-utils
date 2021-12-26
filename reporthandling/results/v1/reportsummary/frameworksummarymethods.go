package reportsummary

import (
	"github.com/armosec/opa-utils/reporthandling/apis"
)

// =================================== Status ============================================

// IsPassed did this framework pass
func (frameworkSummary *FrameworkSummary) IsPassed() bool {
	return frameworkSummary.Status() == apis.StatusPassed
}

// IsFailed did this framework fail
func (frameworkSummary *FrameworkSummary) IsFailed() bool {
	return frameworkSummary.Status() == apis.StatusFailed
}

// IsExcluded is this framework excluded
func (frameworkSummary *FrameworkSummary) IsExcluded() bool {
	return frameworkSummary.Status() == apis.StatusExcluded
}

// IsSkipped was this framework skipped
func (frameworkSummary *FrameworkSummary) IsSkipped() bool {
	return frameworkSummary.Status() == apis.StatusSkipped
}

// Status get the framework status. returns an apis.ScanningStatus object
func (frameworkSummary *FrameworkSummary) Status() apis.ScanningStatus {
	return calculateStatus(&frameworkSummary.ResourceCounters)
}

// =================================== Counters ============================================

// NumberOfExcluded get the number of excluded resources
func (frameworkSummary *FrameworkSummary) NumberOfExcluded() int {
	return frameworkSummary.ResourceCounters.NumberOfExcluded()
}

// NumberOfPassed get the number of passed resources
func (frameworkSummary *FrameworkSummary) NumberOfPassed() int {
	return frameworkSummary.ResourceCounters.NumberOfPassed()
}

// NumberOfSkipped get the number of skipped resources
func (frameworkSummary *FrameworkSummary) NumberOfSkipped() int {
	return frameworkSummary.ResourceCounters.NumberOfSkipped()
}

// NumberOfFailed get the number of failed resources
func (frameworkSummary *FrameworkSummary) NumberOfFailed() int {
	return frameworkSummary.ResourceCounters.NumberOfFailed()
}

// NumberOfAll get the number of all resources
func (frameworkSummary *FrameworkSummary) NumberOfAll() int {
	return frameworkSummary.ResourceCounters.NumberOfAll()
}

// =================================== Score ============================================

// GetScore return framework score
func (frameworkSummary *FrameworkSummary) GetScore() float32 {
	return frameworkSummary.Score
}

// =================================== Name ============================================

// GetName return framework name
func (frameworkSummary *FrameworkSummary) GetName() string {
	return frameworkSummary.Name
}
