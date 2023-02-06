package reportsummary

import (
	"github.com/kubescape/opa-utils/reporthandling/apis"
)

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

// deprecated
func (pcounter *PostureCounters) Excluded() int {
	return pcounter.ExcludedCounter
}

func (pcounter *PostureCounters) All() int {
	return pcounter.Passed() + pcounter.Skipped() + pcounter.Failed()
}

func (pcounter *PostureCounters) Increase(status apis.IStatus) {
	switch status.Status() {
	case apis.StatusFailed:
		pcounter.FailedCounter++
	case apis.StatusPassed:
		pcounter.PassedCounter++
	case apis.StatusSkipped:
		pcounter.SkippedCounter++
	}
}
