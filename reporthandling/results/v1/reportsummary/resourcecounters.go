package reportsummary

import "github.com/armosec/opa-utils/reporthandling/apis"

// =================================== Counters ============================================

// NumberOfExcluded get the number of excluded resources
func (resourceCounters *ResourceCounters) NumberOfExcluded() int {
	return resourceCounters.ExcludedResources
}

// NumberOfPassed get the number of passed resources
func (resourceCounters *ResourceCounters) NumberOfPassed() int {
	return resourceCounters.PassedResources
}

// NumberOfSkipped get the number of skipped resources
func (resourceCounters *ResourceCounters) NumberOfSkipped() int {
	return resourceCounters.SkippedResources
}

// NumberOfFailed get the number of failed resources
func (resourceCounters *ResourceCounters) NumberOfFailed() int {
	return resourceCounters.FailedResources
}

// NumberOfAll get the number of all resources
func (resourceCounters *ResourceCounters) NumberOfAll() int {
	return resourceCounters.NumberOfExcluded() + resourceCounters.NumberOfFailed() + resourceCounters.NumberOfSkipped() + resourceCounters.NumberOfPassed()
}

// =================================== Setters ============================================

// Increase increases the counter based on the status
func (resourceCounters *ResourceCounters) Increase(status apis.IStatus) {
	switch status.Status() {
	case apis.StatusExcluded:
		resourceCounters.FailedResources++
	case apis.StatusFailed:
		resourceCounters.FailedResources++
	case apis.StatusSkipped:
		resourceCounters.SkippedResources++
	case apis.StatusPassed:
		resourceCounters.PassedResources++
	}
}

// setNumberOfFailed set the number of failed resources
func (resourceCounters *ResourceCounters) setNumberOfFailed(n int) {
	resourceCounters.FailedResources = n
}

// setNumberOfFailed set the number of passed resources
func (resourceCounters *ResourceCounters) setNumberOfPassed(n int) {
	resourceCounters.PassedResources = n
}

// setNumberOfFailed set the number of excluded resources
func (resourceCounters *ResourceCounters) setNumberOfExcluded(n int) {
	resourceCounters.ExcludedResources = n
}

// setNumberOfFailed set the number of skipped resources
func (resourceCounters *ResourceCounters) setNumberOfSkipped(n int) {
	resourceCounters.SkippedResources = n
}
