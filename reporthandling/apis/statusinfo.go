package apis

type StatusInfo struct {
	InnerStatus ScanningStatus    `json:"status,omitempty"`
	SubStatus   ScanningSubStatus `json:"subStatus,omitempty"`
	InnerInfo   string            `json:"info,omitempty"`
}

func (s *StatusInfo) GetSubStatus() ScanningSubStatus {
	return s.SubStatus
}

func (s *StatusInfo) Status() ScanningStatus {
	return s.InnerStatus
}

func (s *StatusInfo) Info() string {
	return s.InnerInfo
}

// IsPassed is the status pass
func (s *StatusInfo) IsPassed() bool {
	return s.InnerStatus == StatusPassed
}

// IsFailed is the status fail
func (s *StatusInfo) IsFailed() bool {
	return s.InnerStatus == StatusFailed
}

// IsSkipped is the status skipped
func (s *StatusInfo) IsSkipped() bool {
	return s.InnerStatus == StatusSkipped
}
