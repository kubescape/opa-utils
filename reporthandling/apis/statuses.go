package apis

type ScanningStatus string

const (
	StatusPassed   ScanningStatus = "passed"
	StatusExcluded ScanningStatus = "excluded"
	StatusIgnore   ScanningStatus = "ignored"
	StatusFailed   ScanningStatus = "failed"
	StatusSkipped  ScanningStatus = "skipped"
)
