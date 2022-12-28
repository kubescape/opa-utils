package reportsummary

import (
	"github.com/kubescape/opa-utils/reporthandling/apis"
	helpersv1 "github.com/kubescape/opa-utils/reporthandling/helpers/v1"
)

type ICounters interface {
	Passed() int
	PassedExceptions() int
	PassedIrrelevant() int
	SkippedIntegration() int
	SkippedConfiguration() int
	SkippedRequiresReview() int
	SkippedManualReview() int
	Failed() int
	All() int
}

// =================================== Counters ============================================

// NumberOfExcluded get the number of excluded resources
func (resourceCounters *ResourceCounters) Set(allLists *helpersv1.AllLists) {
	resourceCounters.FailedResources = len(allLists.Failed())
	resourceCounters.PassedResources = len(allLists.Passed())
	resourceCounters.PassedExceptionResources = len(allLists.PassedExceptions())
	resourceCounters.PassedIrrelevantResources = len(allLists.PassedIrrelevant())
	resourceCounters.SkippedIntegrationResources = len(allLists.SkippedIntegration())
	resourceCounters.SkippedConfigurationResources = len(allLists.SkippedConfiguration())
	resourceCounters.SkippedRequiresReviewResources = len(allLists.SkippedRequiresReview())
	resourceCounters.SkippedManualReviewResources = len(allLists.SkippedManualReview())
}

// NumberOfPassed get the number of passed resources
func (resourceCounters *ResourceCounters) Passed() int {
	return resourceCounters.PassedResources
}

// NumberOfPassedException get the number of passed exception resources
func (resourceCounters *ResourceCounters) PassedExceptions() int {
	return resourceCounters.PassedExceptionResources
}

// NumberOfPassedIrrelevant get the number of passed irrelevant resources
func (resourceCounters *ResourceCounters) PassedIrrelevant() int {
	return resourceCounters.PassedIrrelevantResources
}

// NumberOfSkippedIntegration get the number of skipped integration resources
func (resourceCounters *ResourceCounters) SkippedIntegration() int {
	return resourceCounters.SkippedIntegrationResources
}

// NumberOfSkippedConfiguration get the number of skipped configuration resources
func (resourceCounters *ResourceCounters) SkippedConfiguration() int {
	return resourceCounters.SkippedConfigurationResources
}

// NumberOfSkippedRequiresReview get the number of skipped requires review resources
func (resourceCounters *ResourceCounters) SkippedRequiresReview() int {
	return resourceCounters.SkippedRequiresReviewResources
}

// NumberOfSkippedManualReview get the number of skipped manual review resources
func (resourceCounters *ResourceCounters) SkippedManualReview() int {
	return resourceCounters.SkippedManualReviewResources
}

// NumberOfFailed get the number of failed resources
func (resourceCounters *ResourceCounters) Failed() int {
	return resourceCounters.FailedResources
}

// NumberOfAll get the number of all resources
func (resourceCounters *ResourceCounters) All() int {
	return resourceCounters.Failed() +
		resourceCounters.Passed() +
		resourceCounters.PassedExceptions() +
		resourceCounters.PassedIrrelevant() +
		resourceCounters.SkippedIntegration() +
		resourceCounters.SkippedConfiguration() +
		resourceCounters.SkippedRequiresReview() +
		resourceCounters.SkippedManualReview()
}

// =================================== Setters ============================================

// Increase increases the counter based on the status
func (resourceCounters *ResourceCounters) Increase(status apis.IStatus) {
	switch status.Status() {
	case apis.StatusFailed:
		resourceCounters.FailedResources++
	case apis.StatusPassed:
		resourceCounters.PassedResources++
	case apis.SubStatusException:
		resourceCounters.PassedExceptionResources++
	case apis.SubStatusIrrelevant:
		resourceCounters.PassedIrrelevantResources++
	case apis.SubStatusIntegration:
		resourceCounters.SkippedIntegrationResources++
	case apis.SubStatusConfiguration:
		resourceCounters.SkippedConfigurationResources++
	case apis.SubStatusRequiresReview:
		resourceCounters.SkippedRequiresReviewResources++
	case apis.SubStatusManualReview:
		resourceCounters.SkippedManualReviewResources++
	}
}
