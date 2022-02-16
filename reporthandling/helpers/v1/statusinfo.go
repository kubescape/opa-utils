package v1

import "github.com/armosec/opa-utils/reporthandling/apis"

type StatusInfo struct {
	status apis.ScanningStatus
	info   string
}

func NewStatusInfo(status apis.ScanningStatus, info string) *StatusInfo {
	return &StatusInfo{
		status: status,
		info:   info,
	}
}

func (s *StatusInfo) Status() apis.ScanningStatus {
	return s.status
}

func (s *StatusInfo) Info() string {
	return s.info
}

// IsPassed is the status pass
func (s *StatusInfo) IsPassed() bool {
	return s.status == apis.StatusPassed
}

// IsFailed is the status fail
func (s *StatusInfo) IsFailed() bool {
	return s.status == apis.StatusFailed
}

// IsExcluded is the status excluded
func (s *StatusInfo) IsExcluded() bool {
	return s.status == apis.StatusExcluded
}

// IsSkipped is the status skipped
func (s *StatusInfo) IsSkipped() bool {
	return s.status == apis.StatusSkipped
}
