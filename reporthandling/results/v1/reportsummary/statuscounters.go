package reportsummary

import (
	"github.com/kubescape/opa-utils/reporthandling/apis"
	helpersv1 "github.com/kubescape/opa-utils/reporthandling/helpers/v1"
)

type ICounters interface {
	Passed() int
	Skipped() int
	Failed() int
	Excluded() int // deprecated
	All() int
}

type ISubCounters interface {
	All() int
	Ignored() int
}

// =================================== Counters ============================================

// NumberOfExcluded get the number of excluded resources
func (resourceCounters *StatusCounters) Set(allLists *helpersv1.AllLists) {
	resourceCounters.FailedResources = allLists.Failed()
	resourceCounters.PassedResources = allLists.Passed()
	resourceCounters.SkippedResources = allLists.Skipped()
}

// NumberOfPassed get the number of passed resources
func (resourceCounters *StatusCounters) Passed() int {
	return resourceCounters.PassedResources + resourceCounters.ExcludedResources
}

// NumberOfSkipped get the number of skipped resources
func (resourceCounters *StatusCounters) Skipped() int {
	return resourceCounters.SkippedResources
}

// NumberOfFailed get the number of failed resources
func (resourceCounters *StatusCounters) Failed() int {
	return resourceCounters.FailedResources
}

// deprecated
func (resourceCounters *StatusCounters) Excluded() int {
	return resourceCounters.ExcludedResources
}

// NumberOfAll get the number of all resources
func (resourceCounters *StatusCounters) All() int {
	return resourceCounters.Failed() + resourceCounters.Passed() + resourceCounters.Skipped()
}

// =================================== SubCounters ============================================

func (subStatusCounters *SubStatusCounters) All() int {
	return subStatusCounters.Ignored()
}

func (subStatusCounters *SubStatusCounters) Ignored() int {
	return subStatusCounters.IgnoredResources
}

// =================================== Setters ============================================

// Increase increases the counter based on the status
func (resourceCounters *StatusCounters) Increase(status apis.IStatus) {
	switch status.Status() {
	case apis.StatusFailed:
		resourceCounters.FailedResources++
	case apis.StatusPassed:
		resourceCounters.PassedResources++
	case apis.StatusSkipped:
		resourceCounters.SkippedResources++
	}
}

// Increase increases the counter based on the status
func (subStatusCounters *SubStatusCounters) Increase(status apis.IStatus) {
	if status.IsPassed() && status.GetSubStatus() == apis.SubStatusException {
		subStatusCounters.IgnoredResources++
	}
}
