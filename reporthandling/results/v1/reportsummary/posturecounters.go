package reportsummary

import (
	"github.com/kubescape/opa-utils/reporthandling/apis"
)

// NumberOfPassed get the number of passed posture object
func (pcounter *PostureCounters) Passed() int {
	return pcounter.PassedCounter
}

// NumberOfPassedException get the number of passed exception posture object
func (pcounter *PostureCounters) PassedExceptions() int {
	return pcounter.PassedExceptionCounter
}

// NumberOfPassedIrrelevant get the number of passed irrelevant posture object
func (pcounter *PostureCounters) PassedIrrelevant() int {
	return pcounter.PassedIrrelevantCounter
}

// NumberOfSkippedIntegration get the number of skipped integration posture object
func (pcounter *PostureCounters) SkippedIntegration() int {
	return pcounter.SkippedIntegrationCounter
}

// NumberOfSkippedConfiguration get the number of skipped configuration posture object
func (pcounter *PostureCounters) SkippedConfiguration() int {
	return pcounter.SkippedConfigurationCounter
}

// NumberOfSkippedRequiresReview get the number of skipped requires review posture object
func (pcounter *PostureCounters) SkippedRequiresReview() int {
	return pcounter.SkippedRequiresReviewCounter
}

// NumberOfSkippedManualReview get the number of skipped manual review posture object
func (pcounter *PostureCounters) SkippedManualReview() int {
	return pcounter.SkippedManualReviewCounter
}

// Failed get the number of failed posture object
func (pcounter *PostureCounters) Failed() int {
	return pcounter.FailedCounter
}

func (pcounter *PostureCounters) All() int {
	// return pcounter.Passed() + pcounter.Excluded() + pcounter.Failed() + pcounter.Ignored() + pcounter.Skipped()
	return pcounter.Passed() +
		pcounter.PassedExceptions() +
		pcounter.PassedIrrelevant() +
		pcounter.SkippedIntegration() +
		pcounter.SkippedConfiguration() +
		pcounter.SkippedRequiresReview() +
		pcounter.SkippedManualReview() +
		pcounter.Failed()
}

func (pcounter *PostureCounters) Increase(status apis.IStatus) {
	switch status.Status() {
	case apis.StatusFailed:
		pcounter.FailedCounter++
	case apis.StatusPassed:
		pcounter.PassedCounter++
	case apis.SubStatusException:
		pcounter.PassedExceptionCounter++
	case apis.SubStatusIrrelevant:
		pcounter.PassedIrrelevantCounter++
	case apis.SubStatusIntegration:
		pcounter.SkippedIntegrationCounter++
	case apis.SubStatusConfiguration:
		pcounter.SkippedConfigurationCounter++
	case apis.SubStatusRequiresReview:
		pcounter.SkippedRequiresReviewCounter++
	case apis.SubStatusManualReview:
		pcounter.SkippedManualReviewCounter++
	}
}
