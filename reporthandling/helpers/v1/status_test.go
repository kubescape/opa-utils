package v1

import (
	"testing"

	"github.com/kubescape/opa-utils/reporthandling/apis"
)

func TestNewStatus(t *testing.T) {
	tests := []struct {
		name   string
		status apis.ScanningStatus
		want   apis.ScanningStatus
	}{
		{
			name:   "Test passed status",
			status: apis.StatusPassed,
			want:   apis.StatusPassed,
		},
		{
			name:   "Test failed status",
			status: apis.StatusFailed,
			want:   apis.StatusFailed,
		},
		{
			name:   "Test skipped status",
			status: apis.StatusSkipped,
			want:   apis.StatusSkipped,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStatus(tt.status)
			if got := s.Status(); got != tt.want {
				t.Errorf("NewStatus().Status() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewStatusInfo(t *testing.T) {
	tests := []struct {
		name          string
		status        apis.ScanningStatus
		subStatus     apis.ScanningSubStatus
		info          string
		wantStatus    apis.ScanningStatus
		wantSubStatus apis.ScanningSubStatus
		wantInfo      string
	}{
		{
			name:          "Test passed status with ignore sub status and info",
			status:        apis.StatusPassed,
			subStatus:     apis.SubStatusException,
			info:          string(apis.SubStatusConfigurationInfo),
			wantStatus:    apis.StatusPassed,
			wantSubStatus: apis.SubStatusException,
			wantInfo:      string(apis.SubStatusConfigurationInfo),
		},
		{
			name:          "Test failed status and info",
			status:        apis.StatusFailed,
			subStatus:     "",
			info:          string(apis.SubStatusManualReviewInfo),
			wantStatus:    apis.StatusFailed,
			wantSubStatus: "",
			wantInfo:      string(apis.SubStatusManualReviewInfo),
		},
		{
			name:          "Test skipped status with irrelevant sub status without info",
			status:        apis.StatusSkipped,
			subStatus:     apis.SubStatusConfiguration,
			info:          "",
			wantStatus:    apis.StatusSkipped,
			wantSubStatus: apis.SubStatusConfiguration,
			wantInfo:      "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStatusInfo(tt.status, tt.subStatus, tt.info)
			if got := s.Status(); got != tt.wantStatus {
				t.Errorf("NewStatusInfo().Status() = %v, want %v", got, tt.wantStatus)
			}
			if got := s.GetSubStatus(); got != tt.wantSubStatus {
				t.Errorf("NewStatusInfo().GetSubStatus() = %v, want %v", got, tt.wantSubStatus)
			}
			if got := s.Info(); got != tt.wantInfo {
				t.Errorf("NewStatusInfo().Info() = %v, want %v", got, tt.wantInfo)
			}
		})
	}
}

func TestGetSubStatus(t *testing.T) {
	status := &Status{subStatus: apis.SubStatusIntegration}
	if status.GetSubStatus() != apis.SubStatusIntegration {
		t.Errorf("Expected subStatus to be %s, got %s", apis.SubStatusIntegration, status.GetSubStatus())
	}
}

func TestStatus(t *testing.T) {
	status := &Status{status: apis.StatusFailed}
	if status.Status() != apis.StatusFailed {
		t.Errorf("Expected status to be %s, got %s", apis.StatusFailed, status.Status())
	}
}

func TestInfo(t *testing.T) {
	status := &Status{}
	if status.Info() != "" {
		t.Errorf("Expected Info to be empty string, got %s", status.Info())
	}
}

func TestIsPassed(t *testing.T) {
	status := &Status{status: apis.StatusPassed}
	if !status.IsPassed() {
		t.Errorf("Expected IsPassed to be true, got false")
	}
}

func TestIsFailed(t *testing.T) {
	status := &Status{status: apis.StatusFailed}
	if !status.IsFailed() {
		t.Errorf("Expected IsFailed to be true, got false")
	}
}

func TestIsSkipped(t *testing.T) {
	status := &Status{status: apis.StatusSkipped}
	if !status.IsSkipped() {
		t.Errorf("Expected IsSkipped to be true, got false")
	}
}
