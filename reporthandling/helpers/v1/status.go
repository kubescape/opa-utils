package v1

import "github.com/armosec/opa-utils/reporthandling/apis"

type Status struct {
	status apis.ScanningStatus
}

func NewStatus(status apis.ScanningStatus) *Status {
	return &Status{status: status}
}

func (s *Status) Status() apis.ScanningStatus {
	return s.status
}

func (s *Status) Info() string {
	return ""
}

// IsPassed is the status pass
func (s *Status) IsPassed() bool {
	return s.status == apis.StatusPassed
}

// IsFailed is the status fail
func (s *Status) IsFailed() bool {
	return s.status == apis.StatusFailed
}

// IsExcluded is the status excluded
func (s *Status) IsExcluded() bool {
	return s.status == apis.StatusExcluded
}

// IsSkipped is the status skipped
func (s *Status) IsSkipped() bool {
	return s.status == apis.StatusSkipped
}
