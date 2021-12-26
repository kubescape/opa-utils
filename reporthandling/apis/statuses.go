package apis

type ScanningStatus string

const (
	StatusPassed   ScanningStatus = "passed"
	StatusExcluded ScanningStatus = "excluded"
	StatusIgnored  ScanningStatus = "ignored"
	StatusFailed   ScanningStatus = "failed"
	StatusSkipped  ScanningStatus = "skipped"
)

// Compare receive two statuses and returns the more significant one
/*

	status level:
		1. failed
		2. excludes
		3. passed
		4. skipped/ignore

	e.g.:
	Compare(failed, excludes) -> failed
	Compare(passed, excludes) -> excludes
	Compare(skipped, excludes) -> excludes
	Compare(failed, passed) -> failed
*/
func Compare(a, b ScanningStatus) ScanningStatus {
	if a == StatusFailed || b == StatusFailed {
		return StatusFailed
	}
	if a == StatusExcluded || b == StatusExcluded {
		return StatusExcluded
	}
	if a != StatusPassed && b != StatusPassed {
		return StatusSkipped
	}
	return StatusPassed
}
