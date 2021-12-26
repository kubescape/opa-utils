package reportsummary

import "github.com/armosec/opa-utils/reporthandling/apis"

func calculateStatus(counters *ResourceCounters) apis.ScanningStatus {

	if counters.NumberOfFailed() != 0 {
		return apis.StatusFailed
	}
	if counters.NumberOfExcluded() != 0 {
		return apis.StatusExcluded
	}
	if counters.NumberOfSkipped() == counters.NumberOfAll() {
		return apis.StatusSkipped
	}
	return apis.StatusPassed
}
