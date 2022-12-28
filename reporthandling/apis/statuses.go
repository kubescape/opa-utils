package apis

type ScanningStatus string

const (
	StatusPassed            ScanningStatus = "passed"
	StatusFailed            ScanningStatus = "failed"
	StatusSkipped           ScanningStatus = "skipped"
	SubStatusException      ScanningStatus = "w/exceptions"
	SubStatusIrrelevant     ScanningStatus = "irrelevant"
	SubStatusConfiguration  ScanningStatus = "configuration"
	SubStatusIntegration    ScanningStatus = "integration"
	SubStatusRequiresReview ScanningStatus = "requires review"
	SubStatusManualReview   ScanningStatus = "manual review"
	StatusUnknown           ScanningStatus = "" // keep this empty

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
func CompareStatusAndSubStatus(a, aSub, b, bSub ScanningStatus) (ScanningStatus, ScanningStatus) {
	status := Compare(a, b)
	if status == StatusFailed || status == StatusUnknown {
		return status, ScanningStatus(StatusUnknown)
	}

	if status == StatusPassed {
		if aSub == ScanningStatus(SubStatusException) || bSub == ScanningStatus(SubStatusException) {
			return status, ScanningStatus(SubStatusException)
		}
		if aSub == ScanningStatus(SubStatusIrrelevant) || bSub == ScanningStatus(SubStatusIrrelevant) {
			return status, ScanningStatus(SubStatusIrrelevant)
		}
	}

	if status == StatusSkipped {
		if aSub == ScanningStatus(SubStatusConfiguration) || bSub == ScanningStatus(SubStatusConfiguration) {
			return status, ScanningStatus(SubStatusConfiguration)
		}
		if aSub == ScanningStatus(SubStatusIntegration) || bSub == ScanningStatus(SubStatusIntegration) {
			return status, ScanningStatus(SubStatusIntegration)
		}
		if aSub == ScanningStatus(SubStatusRequiresReview) || bSub == ScanningStatus(SubStatusRequiresReview) {
			return status, ScanningStatus(SubStatusRequiresReview)
		}
		if aSub == ScanningStatus(SubStatusManualReview) || bSub == ScanningStatus(SubStatusManualReview) {
			return status, ScanningStatus(SubStatusManualReview)
		}

	}
	return status, ScanningStatus(StatusUnknown)
}
