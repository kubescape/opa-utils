package reportsummary

import (
	"github.com/kubescape/opa-utils/reporthandling/apis"
)

// NumberOfExcluded get the number of excluded posture object
func (pcounter *PostureCounters) Excluded() int {
	return pcounter.ExcludedCounter
}

// NumberOfPassed get the number of passed posture object
func (pcounter *PostureCounters) Passed() int {
	return pcounter.PassedCounter
}

// NumberOfSkipped get the number of skipped posture object
func (pcounter *PostureCounters) Skipped() int {
	return pcounter.SkippedCounter
}

// Failed get the number of failed posture object
func (pcounter *PostureCounters) Failed() int {
	return pcounter.FailedCounter
}

func (pcounter *PostureCounters) Ignored() int {
	return pcounter.IgnoredCounter
}

func (pcounter *PostureCounters) Unknown() int {
	return pcounter.UnknownCounter
}

func (pcounter *PostureCounters) All() int {
	return pcounter.Passed() + pcounter.Excluded() + pcounter.Failed() + pcounter.Ignored() + pcounter.Skipped()
}

func (pcounter *PostureCounters) Increase(status apis.IStatus) {
	switch status.Status() {
	case apis.StatusExcluded:
		pcounter.ExcludedCounter++
	case apis.StatusFailed:
		pcounter.FailedCounter++
	case apis.StatusPassed:
		pcounter.PassedCounter++
	case apis.StatusSkipped, apis.StatusIrrelevant, apis.StatusError:
		pcounter.SkippedCounter++
	}
}
