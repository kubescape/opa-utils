package apis

type ScanningStatus string
type ScanningSubStatus string

const (
	StatusPassed            ScanningStatus    = "passed"
	StatusFailed            ScanningStatus    = "failed"
	StatusSkipped           ScanningStatus    = "skipped"
	SubStatusException      ScanningSubStatus = "w/exceptions"
	SubStatusIrrelevant     ScanningSubStatus = "irrelevant"
	SubStatusConfiguration  ScanningSubStatus = "configuration"
	SubStatusIntegration    ScanningSubStatus = "integration"
	SubStatusRequiresReview ScanningSubStatus = "requires review"
	SubStatusManualReview   ScanningSubStatus = "manual review"
	SubStatusUnknown        ScanningSubStatus = "" // keep this empty
	StatusUnknown           ScanningStatus    = "" // keep this empty

)
const (
	SubStatusConfigurationInfo  string = "Control missing configuration"
	SubStatusRequiresReviewInfo string = "Control type is requires-review"
	SubStatusManualReviewInfo   string = "Control type is manual-review"
)

// IStatus interface handling status
type IStatus interface {
	Status() ScanningStatus
	Info() string
	IsPassed() bool
	IsFailed() bool
	IsSkipped() bool
}

// Compare receive two statuses and returns the more significant one
/*

	status level:
		1. failed
		2. skipped
		3. passed

	e.g.:
	Compare(failed, skipped) -> failed
	Compare(passed, skipped) -> skipped
	Compare(failed, skipped) -> failed
	Compare(skipped, passed) -> skipped
*/
func Compare(a, b ScanningStatus) ScanningStatus {
	if a == StatusFailed || b == StatusFailed {
		return StatusFailed
	}
	if a == StatusSkipped || b == StatusSkipped {
		return StatusSkipped
	}
	if a == StatusUnknown && b == StatusUnknown {
		return StatusUnknown
	}
	return StatusPassed
}

// CompareStatusAndSubStatus receive two statuses + sub statuses and returns the more significant one
/*
	status/sub status level:
		1. status=failed or status=unknown:
			sub status = ""
		2. status=skipped:
			if aSub or bSub are configuration/integration/review:
				sub status = aSub or bSub
			else:
				sub status = status=unknown
		3. status=passed:
			if aSub or bSub are exception/irrelevant:
				sub status = aSub or bSub
			else:
				sub status = status=unknown
*/
func CompareStatusAndSubStatus(a, b ScanningStatus, aSub, bSub ScanningSubStatus) (ScanningStatus, ScanningSubStatus) {
	status := Compare(a, b)
	switch status {
	case StatusFailed, StatusUnknown:
		return status, SubStatusUnknown
	case StatusPassed:
		if aSub == SubStatusException || bSub == SubStatusException {
			return status, SubStatusException
		}
		if aSub == SubStatusIrrelevant || bSub == SubStatusIrrelevant {
			return status, SubStatusIrrelevant
		}

	case StatusSkipped:
		if aSub == SubStatusConfiguration || bSub == SubStatusConfiguration {
			return status, SubStatusConfiguration
		}
		if aSub == SubStatusIntegration || bSub == SubStatusIntegration {
			return status, SubStatusIntegration
		}
		if aSub == SubStatusRequiresReview || bSub == SubStatusRequiresReview {
			return status, SubStatusRequiresReview
		}
		if aSub == SubStatusManualReview || bSub == SubStatusManualReview {
			return status, SubStatusManualReview
		}

	}
	return status, SubStatusUnknown
}
