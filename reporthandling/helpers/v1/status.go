package v1

import "github.com/kubescape/opa-utils/reporthandling/apis"

type Status struct {
	status apis.ScanningStatus
}

func NewStatus(status apis.ScanningStatus) *Status {
	return &Status{status: status}
}

func NewStatusInfo(status apis.ScanningStatus, info string) *apis.StatusInfo {
	return &apis.StatusInfo{
		InnerStatus: status,
		InnerInfo:   info,
	}
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

// IsSkipped is the status skipped
func (s *Status) IsSkipped() bool {
	return s.status == apis.StatusSkipped
}
