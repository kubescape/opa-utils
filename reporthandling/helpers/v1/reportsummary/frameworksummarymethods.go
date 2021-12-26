package reportsummary

import (
	"github.com/armosec/opa-utils/reporthandling/apis"
)

// =================================== Status ============================================

// IsPassed did this framework pass
func (frameworkSummary *FrameworkSummary) IsPassed() bool {
	return frameworkSummary.ScanStatus == apis.StatusPassed
}

// IsFailed did this framework fail
func (frameworkSummary *FrameworkSummary) IsFailed() bool {
	return frameworkSummary.ScanStatus == apis.StatusFailed
}

// IsExcluded is this framework excluded
func (frameworkSummary *FrameworkSummary) IsExcluded() bool {
	return frameworkSummary.ScanStatus == apis.StatusExcluded
}

// IsSkipped was this framework skipped
func (frameworkSummary *FrameworkSummary) IsSkipped() bool {
	return frameworkSummary.ScanStatus == apis.StatusSkipped
}

// Status get the framework status. returns an apis.ScanningStatus object
func (frameworkSummary *FrameworkSummary) Status() apis.ScanningStatus {
	return frameworkSummary.ScanStatus
}

// setStatus set the framework status. returns an apis.ScanningStatus object
func (frameworkSummary *FrameworkSummary) setStatus(s apis.ScanningStatus) {
	frameworkSummary.ScanStatus = s
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

// =================================== Setters ============================================

// setNumberOfFailed set the number of failed resources
func (frameworkSummary *FrameworkSummary) setNumberOfFailed(n int) {
	frameworkSummary.FailedResources = n
}

// setNumberOfFailed set the number of passed resources
func (frameworkSummary *FrameworkSummary) setNumberOfPassed(n int) {
	frameworkSummary.PassedResources = n
}

// setNumberOfFailed set the number of excluded resources
func (frameworkSummary *FrameworkSummary) setNumberOfExcluded(n int) {
	frameworkSummary.ExcludedResources = n
}

// setNumberOfFailed set the number of skipped resources
func (frameworkSummary *FrameworkSummary) setNumberOfSkipped(n int) {
	frameworkSummary.SkippedResources = n
}

// =================================== Score ============================================

// GetScore return framework score
func (frameworkSummary *FrameworkSummary) GetScore() float32 {
	return frameworkSummary.Score
}

// setScore set framework score
func (frameworkSummary *FrameworkSummary) setScore(n float32) {
	frameworkSummary.Score = n
}
