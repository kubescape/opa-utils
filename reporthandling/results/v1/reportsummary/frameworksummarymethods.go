package reportsummary

import (
	"github.com/armosec/opa-utils/reporthandling/apis"
	helpersv1 "github.com/armosec/opa-utils/reporthandling/helpers/v1"
)

// =================================== Status ============================================

// Status get the framework status. returns an apis.ScanningStatus object
func (frameworkSummary *FrameworkSummary) GetStatus() apis.IStatus {
	return helpersv1.NewStatus(frameworkSummary.Status)
}

// SetStatus set the framework status based on the resource counters
func (frameworkSummary *FrameworkSummary) CalculateStatus() {
	frameworkSummary.Status = calculateStatus(&frameworkSummary.ResourceCounters)
	for k, v := range frameworkSummary.Controls {
		v.CalculateStatus()
		frameworkSummary.Controls[k] = v
	}
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

// Increase increases the counter based on the status
func (frameworkSummary *FrameworkSummary) Increase(status apis.IStatus) {
	frameworkSummary.ResourceCounters.Increase(status)
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
