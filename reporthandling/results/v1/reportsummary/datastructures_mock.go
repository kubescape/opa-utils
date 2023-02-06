package reportsummary

import (
	"github.com/kubescape/opa-utils/reporthandling/apis"
	helpersv1 "github.com/kubescape/opa-utils/reporthandling/helpers/v1"
)

func MockSummaryDetails() *SummaryDetails {
	return mockSummaryDetailsFailed()
}
func mockSummaryDetailsStatusIrrelevant() *SummaryDetails {
	return &SummaryDetails{
		Frameworks: []FrameworkSummary{
			*mockFrameworkSummaryFailPass(),
		},
		Controls: map[string]ControlSummary{
			"C-0001": *mockControlSummaryStatusIrrelevant(),
		},
		StatusCounters: *mockStatusCountersSkippedFailPass(),
		Status:         apis.StatusFailed,
	}
}

func mockSummaryDetailsNoInnerStatus() *SummaryDetails {
	return &SummaryDetails{
		Frameworks: []FrameworkSummary{
			*mockFrameworkSummaryFailPass(),
		},
		Controls: map[string]ControlSummary{
			"C-0001": *mockControlSummaryNoInnerStatus(),
		},
		StatusCounters: *mockStatusCountersSkippedFailPass(),
		Status:         apis.StatusFailed,
	}
}

func mockSummaryDetailsStatusEmpty() *SummaryDetails {
	return &SummaryDetails{
		Frameworks: []FrameworkSummary{
			*mockFrameworkSummaryFailPass(),
		},
		Controls: map[string]ControlSummary{
			"C-0001": *mockControlSummaryStatusEmpty(),
		},
		StatusCounters: *mockStatusCountersSkippedFailPass(),
		Status:         apis.StatusFailed,
	}
}

func mockSummaryDetailsStatusSkipped() *SummaryDetails {
	return &SummaryDetails{
		Frameworks: []FrameworkSummary{
			*mockFrameworkSummaryFailPass(),
		},
		Controls: map[string]ControlSummary{
			"C-0001": *mockControlSummaryStatusSkipped(),
		},
		StatusCounters: *mockStatusCountersSkippedFailPass(),
		Status:         apis.StatusFailed,
	}
}

func mockSummaryDetailsFailed() *SummaryDetails {
	return &SummaryDetails{
		Frameworks: []FrameworkSummary{
			*mockFrameworkSummaryFailPass(),
			*mockFrameworkSummaryFailException(),
			*mockFrameworkSummaryPassException(),
		},
		Controls: map[string]ControlSummary{
			"C-0001": *mockControlSummaryFailPass(),
			"C-0002": *mockControlSummaryFailPassException(),
			"C-0003": *mockControlSummaryExceptionPass(),
			"C-0004": *mockControlSummaryPass(),
		},
		StatusCounters: *mockStatusCountersSkippedFailPass(),
		Status:         apis.StatusFailed,
	}
}
func mockSummaryDetailsPassed() *SummaryDetails {
	return &SummaryDetails{
		Controls: map[string]ControlSummary{
			"C-0004": *mockControlSummaryPass(),
		},
		StatusCounters: *mockStatusCountersPass(),
	}
}
func mockSummaryDetailsException() *SummaryDetails {
	return &SummaryDetails{
		Frameworks: []FrameworkSummary{
			*mockFrameworkSummaryPassException(),
		},
		Controls: map[string]ControlSummary{
			"C-0003": *mockControlSummaryExceptionPass(),
		},
		StatusCounters: *mockStatusCountersExceptions(),
	}
}
func mockFrameworkSummaryPassException() *FrameworkSummary {
	return &FrameworkSummary{
		Name:           "fw-pass-exclude",
		Score:          0,
		Status:         apis.StatusPassed,
		Version:        "utnitest",
		StatusCounters: *mockStatusCountersExceptionPass(),
		Controls: map[string]ControlSummary{
			"C-0003": *mockControlSummaryExceptionPass(),
		},
	}
}
func mockFrameworkSummaryFailException() *FrameworkSummary {
	return &FrameworkSummary{
		Name:           "fw-failed",
		Score:          0,
		Status:         apis.StatusFailed,
		Version:        "utnitest",
		StatusCounters: *mockStatusCountersSkippedFailPass(),
		Controls: map[string]ControlSummary{
			"C-0001": *mockControlSummaryFailPass(),
			"C-0002": *mockControlSummaryFailPassException(),
		},
	}
}
func mockFrameworkSummaryFailPass() *FrameworkSummary {
	return &FrameworkSummary{
		Name:           "fw-failed",
		Score:          0,
		Status:         apis.StatusFailed,
		Version:        "utnitest",
		StatusCounters: *mockStatusCountersFailPass(),
		Controls: map[string]ControlSummary{
			"C-0001": *mockControlSummaryFailPass(),
			"C-0003": *mockControlSummaryExceptionPass(),
			"C-0004": *mockControlSummaryPass(),
		},
	}
}
func mockControlSummaryExceptionPass() *ControlSummary {
	return &ControlSummary{
		Name:           "control-exclude-pass",
		Score:          0,
		Status:         apis.StatusPassed,
		StatusInfo:     *helpersv1.NewStatusInfo(apis.StatusPassed, apis.SubStatusException, ""),
		StatusCounters: *mockStatusCountersExceptionPass(),
		ResourceIDs:    *helpersv1.MockAllListsForIntegration(),
	}
}

func mockControlSummaryPass() *ControlSummary {
	return &ControlSummary{
		Name:           "control-pass",
		Score:          0,
		Status:         apis.StatusPassed,
		StatusCounters: *mockStatusCountersPass(),
		ResourceIDs:    *helpersv1.MockAllListsForIntegration(),
	}
}

func mockControlSummaryFailPass() *ControlSummary {
	return &ControlSummary{
		Name:           "control-fail-pass",
		Score:          0,
		Status:         apis.StatusFailed,
		StatusCounters: *mockStatusCountersFailPass(),
		ResourceIDs:    *helpersv1.MockAllListsForIntegration(),
	}
}
func mockControlSummaryFailPassException() *ControlSummary {
	return &ControlSummary{
		Name:           "control-fail-pass-exclude",
		Status:         apis.StatusFailed,
		Score:          0,
		StatusCounters: *mockStatusCountersSkippedFailPass(),
		ResourceIDs:    *helpersv1.MockAllListsForIntegration(),
	}
}
func mockControlSummaryStatusSkipped() *ControlSummary {
	return &ControlSummary{
		Name:   "control-skipped",
		Status: apis.StatusSkipped,
		StatusInfo: apis.StatusInfo{
			InnerStatus: apis.StatusSkipped,
			InnerInfo:   "no host sensor flag",
		},
		Score: 0,
	}
}

func mockControlSummaryNoInnerStatus() *ControlSummary {
	return &ControlSummary{
		Name:       "control-irrelevant",
		StatusInfo: *helpersv1.NewStatusInfo(apis.StatusSkipped, apis.SubStatusIntegration, ""),
		Status:     apis.StatusSkipped,
		Score:      0,
	}
}

func mockControlSummaryStatusEmpty() *ControlSummary {
	return &ControlSummary{
		Name:  "control-irrelevant",
		Score: 0,
	}
}

func mockControlSummaryStatusIrrelevant() *ControlSummary {
	return &ControlSummary{
		Name:       "control-irrelevant",
		Status:     apis.StatusPassed,
		StatusInfo: *helpersv1.NewStatusInfo(apis.StatusPassed, apis.SubStatusIrrelevant, "no k8s dashboard in cluster"),
		Score:      0,
	}
}

func mockStatusCountersFailPass() *StatusCounters {
	return &StatusCounters{
		PassedResources: 5,
		FailedResources: 6,
	}
}

func mockStatusCountersExceptionPass() *StatusCounters {
	return &StatusCounters{
		PassedResources: 4,
	}
}

func mockStatusCountersSkippedFailPass() *StatusCounters {
	return &StatusCounters{
		PassedResources:  4,
		SkippedResources: 3,
		FailedResources:  5,
	}
}

func mockStatusCountersPass() *StatusCounters {
	return &StatusCounters{
		PassedResources: 4,
	}
}

func mockStatusCountersSkipped() *StatusCounters {
	return &StatusCounters{
		SkippedResources: 4,
	}
}

func mockStatusCountersExceptions() *StatusCounters {
	return &StatusCounters{
		PassedResources: 4,
	}
}
