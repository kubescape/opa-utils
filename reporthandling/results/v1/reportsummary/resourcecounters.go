package reportsummary

import (
	"github.com/kubescape/opa-utils/reporthandling/apis"
	helpersv1 "github.com/kubescape/opa-utils/reporthandling/helpers/v1"
)

type ICounters interface {
	Passed() int
	Skipped() int
	Failed() int
	All() int
}

// =================================== Counters ============================================

// NumberOfExcluded get the number of excluded resources
func (resourceCounters *ResourceCounters) Set(allLists *helpersv1.AllLists) {
	resourceCounters.FailedResources = len(allLists.Failed())
	resourceCounters.PassedResources = len(allLists.Passed())
	resourceCounters.SkippedResources = len(allLists.Skipped())
}

// NumberOfPassed get the number of passed resources
func (resourceCounters *ResourceCounters) Passed() int {
	return resourceCounters.PassedResources
}

// NumberOfSkipped get the number of skipped resources
func (resourceCounters *ResourceCounters) Skipped() int {
	return resourceCounters.SkippedResources
}

// NumberOfFailed get the number of failed resources
func (resourceCounters *ResourceCounters) Failed() int {
	return resourceCounters.FailedResources
}

// NumberOfAll get the number of all resources
func (resourceCounters *ResourceCounters) All() int {
	return resourceCounters.Failed() + resourceCounters.Passed() + resourceCounters.Skipped()
}

// =================================== Setters ============================================

// Increase increases the counter based on the status
func (resourceCounters *ResourceCounters) Increase(status apis.IStatus) {
	switch status.Status() {
	case apis.StatusFailed:
		resourceCounters.FailedResources++
	case apis.StatusPassed:
		resourceCounters.PassedResources++
	case apis.StatusSkipped:
		resourceCounters.SkippedResources++
	}
}
