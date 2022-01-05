package reportsummary

import "github.com/armosec/opa-utils/reporthandling/apis"

func calculateStatus(counters *ResourceCounters) apis.ScanningStatus {

	if counters.Failed() != 0 {
		return apis.StatusFailed
	}
	if counters.Excluded() != 0 {
		return apis.StatusExcluded
	}
	if counters.All() == counters.Skipped() {
		return apis.StatusSkipped
	}
	return apis.StatusPassed
}
