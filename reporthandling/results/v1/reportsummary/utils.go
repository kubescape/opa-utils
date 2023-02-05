package reportsummary

import (
	"github.com/kubescape/opa-utils/reporthandling/apis"
)

func calculateStatus(counters *StatusCounters) apis.ScanningStatus {

	if counters.Failed() != 0 {
		return apis.StatusFailed
	}
	if counters.Skipped() != 0 {
		return apis.StatusSkipped
	}

	return apis.StatusPassed
}
