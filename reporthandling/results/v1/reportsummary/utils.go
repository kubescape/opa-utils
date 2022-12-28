package reportsummary

import (
	"github.com/kubescape/opa-utils/reporthandling/apis"
)

func calculateStatus(counters *ResourceCounters) apis.ScanningStatus {

	if counters.Failed() != 0 {
		return apis.StatusFailed
	}
	if counters.SkippedIntegration()+counters.SkippedConfiguration()+counters.SkippedManualReview()+counters.SkippedRequiresReview() != 0 {
		return apis.StatusSkipped
	}

	return apis.StatusPassed
}

func calculateSubStatus(counters *ResourceCounters) apis.ScanningStatus {
	if counters.Failed() != 0 {
		return apis.ScanningStatus(apis.StatusUnknown)
	}
	if counters.SkippedConfiguration() != 0 {
		return apis.SubStatusConfiguration
	}
	if counters.SkippedIntegration() != 0 {
		return apis.SubStatusIntegration
	}
	if counters.SkippedRequiresReview() != 0 {
		return apis.SubStatusRequiresReview
	}
	if counters.SkippedManualReview() != 0 {
		return apis.SubStatusManualReview
	}
	if counters.PassedIrrelevant() != 0 {
		return apis.SubStatusIrrelevant
	}
	if counters.PassedExceptions() != 0 {
		return apis.SubStatusException
	}
	return apis.StatusPassed
}
