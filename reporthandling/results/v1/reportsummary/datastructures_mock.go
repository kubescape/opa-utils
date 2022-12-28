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
		ResourceCounters: *mockResourceCountersExceptionFailPass(),
		Status:           apis.StatusFailed,
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
		ResourceCounters: *mockResourceCountersExceptionFailPass(),
		Status:           apis.StatusFailed,
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
		ResourceCounters: *mockResourceCountersExceptionFailPass(),
		Status:           apis.StatusFailed,
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
		ResourceCounters: *mockResourceCountersExceptionFailPass(),
		Status:           apis.StatusFailed,
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
		ResourceCounters: *mockResourceCountersExceptionFailPass(),
		Status:           apis.StatusFailed,
	}
}
func mockSummaryDetailsPassed() *SummaryDetails {
	return &SummaryDetails{
		Controls: map[string]ControlSummary{
			"C-0004": *mockControlSummaryPass(),
		},
		ResourceCounters: *mockResourceCountersPass(),
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
		ResourceCounters: *mockResourceCountersExceptions(),
	}
}
func mockFrameworkSummaryPassException() *FrameworkSummary {
	return &FrameworkSummary{
		Name:             "fw-pass-exclude",
		Score:            0,
		Status:           apis.StatusPassed,
		Version:          "utnitest",
		ResourceCounters: *mockResourceCountersExceptionPass(),
		Controls: map[string]ControlSummary{
			"C-0003": *mockControlSummaryExceptionPass(),
		},
	}
}
func mockFrameworkSummaryFailException() *FrameworkSummary {
	return &FrameworkSummary{
		Name:             "fw-failed",
		Score:            0,
		Status:           apis.StatusFailed,
		Version:          "utnitest",
		ResourceCounters: *mockResourceCountersExceptionFailPass(),
		Controls: map[string]ControlSummary{
			"C-0001": *mockControlSummaryFailPass(),
			"C-0002": *mockControlSummaryFailPassException(),
		},
	}
}
func mockFrameworkSummaryFailPass() *FrameworkSummary {
	return &FrameworkSummary{
		Name:             "fw-failed",
		Score:            0,
		Status:           apis.StatusFailed,
		Version:          "utnitest",
		ResourceCounters: *mockResourceCountersFailPass(),
		Controls: map[string]ControlSummary{
			"C-0001": *mockControlSummaryFailPass(),
			"C-0003": *mockControlSummaryExceptionPass(),
			"C-0004": *mockControlSummaryPass(),
		},
	}
}
func mockControlSummaryExceptionPass() *ControlSummary {
	return &ControlSummary{
		Name:             "control-exclude-pass",
		Score:            0,
		Status:           apis.StatusPassed,
		SubStatus:        apis.SubStatusException,
		ResourceCounters: *mockResourceCountersExceptionPass(),
		ResourceIDs:      *helpersv1.MockAllListsForIntegration(),
	}
}

func mockControlSummaryPass() *ControlSummary {
	return &ControlSummary{
		Name:             "control-pass",
		Score:            0,
		Status:           apis.StatusPassed,
		ResourceCounters: *mockResourceCountersPass(),
		ResourceIDs:      *helpersv1.MockAllListsForIntegration(),
	}
}

func mockControlSummaryFailPass() *ControlSummary {
	return &ControlSummary{
		Name:             "control-fail-pass",
		Score:            0,
		Status:           apis.StatusFailed,
		ResourceCounters: *mockResourceCountersFailPass(),
		ResourceIDs:      *helpersv1.MockAllListsForIntegration(),
	}
}
func mockControlSummaryFailPassException() *ControlSummary {
	return &ControlSummary{
		Name:             "control-fail-pass-exclude",
		Status:           apis.StatusFailed,
		Score:            0,
		ResourceCounters: *mockResourceCountersExceptionFailPass(),
		ResourceIDs:      *helpersv1.MockAllListsForIntegration(),
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
		Status:     apis.StatusSkipped,
		SubStatus:  apis.SubStatusIntegration,
		StatusInfo: apis.StatusInfo{},
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
		Name:      "control-irrelevant",
		Status:    apis.StatusPassed,
		SubStatus: apis.SubStatusIrrelevant,
		StatusInfo: apis.StatusInfo{
			InnerStatus: apis.StatusPassed,
			InnerInfo:   "no k8s dashboard in cluster",
		},
		Score: 0,
	}
}

func mockResourceCountersFailPass() *ResourceCounters {
	return &ResourceCounters{
		PassedResources: 5,
		FailedResources: 6,
	}
}

func mockResourceCountersExceptionPass() *ResourceCounters {
	return &ResourceCounters{
		PassedResources:          4,
		PassedExceptionResources: 3,
	}
}

func mockResourceCountersExceptionFailPass() *ResourceCounters {
	return &ResourceCounters{
		PassedResources:          4,
		PassedExceptionResources: 3,
		FailedResources:          5,
	}
}

func mockResourceCountersPass() *ResourceCounters {
	return &ResourceCounters{
		PassedResources: 4,
	}
}

func mockResourceCountersSkipped() *ResourceCounters {
	return &ResourceCounters{
		SkippedIntegrationResources: 4,
	}
}

func mockResourceCountersExceptions() *ResourceCounters {
	return &ResourceCounters{
		PassedExceptionResources: 4,
	}
}
