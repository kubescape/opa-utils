package reportsummary

import "github.com/kubescape/opa-utils/reporthandling/apis"

func calculateStatus(counters *ResourceCounters) apis.ScanningStatus {

	if counters.Failed() != 0 {
		return apis.StatusFailed
	}
	if counters.Excluded() != 0 {
		return apis.StatusExcluded
	}
	// No resources -> irrelevant
	if counters.All() == 0 {
		return apis.StatusIrrelevant
	}
	return apis.StatusPassed
}
