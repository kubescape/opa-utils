package apis

import "testing"

func TestIsPassed(t *testing.T) {
	s := &StatusInfo{InnerStatus: StatusPassed}
	if !s.IsPassed() {
		t.Fatalf("Expected status to be passed, got %v", s.InnerStatus)
	}
}

func TestIsFailed(t *testing.T) {
	s := &StatusInfo{InnerStatus: StatusFailed}
	if !s.IsFailed() {
		t.Fatalf("Expected status to be failed, got %v", s.InnerStatus)
	}
}

func TestIsSkipped(t *testing.T) {
	s := &StatusInfo{InnerStatus: StatusSkipped}
	if !s.IsSkipped() {
		t.Fatalf("Expected status to be skipped, got %v", s.InnerStatus)
	}
}

func TestGetSubStatus(t *testing.T) {
	s := &StatusInfo{SubStatus: SubStatusException}
	if s.GetSubStatus() != SubStatusException {
		t.Fatalf("Expected sub-status to be %v, got %v", SubStatusException, s.GetSubStatus())
	}
}

func TestStatus(t *testing.T) {
	s := &StatusInfo{InnerStatus: StatusPassed}
	if s.Status() != StatusPassed {
		t.Fatalf("Expected status to be %v, got %v", StatusPassed, s.Status())
	}
}

func TestInfo(t *testing.T) {
	info := "Test info"
	s := &StatusInfo{InnerInfo: info}
	if s.Info() != info {
		t.Fatalf("Expected info to be %v, got %v", info, s.Info())
	}
}
