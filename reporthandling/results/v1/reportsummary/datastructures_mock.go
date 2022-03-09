package reportsummary

import (
	"github.com/armosec/opa-utils/reporthandling/apis"
	helpersv1 "github.com/armosec/opa-utils/reporthandling/helpers/v1"
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
			"C-0001": *mockControlSummaryStatusIrelevant(),
		},
		ResourceCounters: *mockResourceCountersExcludeFailPass(),
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
		ResourceCounters: *mockResourceCountersExcludeFailPass(),
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
		ResourceCounters: *mockResourceCountersExcludeFailPass(),
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
		ResourceCounters: *mockResourceCountersExcludeFailPass(),
		Status:           apis.StatusFailed,
	}
}

func mockSummaryDetailsFailed() *SummaryDetails {
	return &SummaryDetails{
		Frameworks: []FrameworkSummary{
			*mockFrameworkSummaryFailPass(),
			*mockFrameworkSummaryFailExclude(),
			*mockFrameworkSummaryPassExclude(),
		},
		Controls: map[string]ControlSummary{
			"C-0001": *mockControlSummaryFailPass(),
			"C-0002": *mockControlSummaryFailPassExclude(),
			"C-0003": *mockControlSummaryExcludePass(),
			"C-0004": *mockControlSummaryPass(),
		},
		ResourceCounters: *mockResourceCountersExcludeFailPass(),
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
func mockSummaryDetailsExcluded() *SummaryDetails {
	return &SummaryDetails{
		Frameworks: []FrameworkSummary{
			*mockFrameworkSummaryPassExclude(),
		},
		Controls: map[string]ControlSummary{
			"C-0003": *mockControlSummaryExcludePass(),
		},
		ResourceCounters: *mockResourceCountersExclude(),
	}
}
func mockFrameworkSummaryPassExclude() *FrameworkSummary {
	return &FrameworkSummary{
		Name:             "fw-pass-exclude",
		Score:            0,
		Status:           apis.StatusExcluded,
		Version:          "utnitest",
		ResourceCounters: *mockResourceCountersExcludePass(),
		Controls: map[string]ControlSummary{
			"C-0003": *mockControlSummaryExcludePass(),
		},
	}
}
func mockFrameworkSummaryFailExclude() *FrameworkSummary {
	return &FrameworkSummary{
		Name:             "fw-failed",
		Score:            0,
		Status:           apis.StatusFailed,
		Version:          "utnitest",
		ResourceCounters: *mockResourceCountersExcludeFailPass(),
		Controls: map[string]ControlSummary{
			"C-0001": *mockControlSummaryFailPass(),
			"C-0002": *mockControlSummaryFailPassExclude(),
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
			"C-0003": *mockControlSummaryExcludePass(),
			"C-0004": *mockControlSummaryPass(),
		},
	}
}
func mockControlSummaryExcludePass() *ControlSummary {
	return &ControlSummary{
		Name:             "control-exclude-pass",
		Score:            0,
		Status:           apis.StatusExcluded,
		ResourceCounters: *mockResourceCountersExcludePass(),
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
func mockControlSummaryFailPassExclude() *ControlSummary {
	return &ControlSummary{
		Name:             "control-fail-pass-exclude",
		Status:           apis.StatusFailed,
		Score:            0,
		ResourceCounters: *mockResourceCountersExcludeFailPass(),
		ResourceIDs:      *helpersv1.MockAllListsForIntegration(),
	}
}
func mockControlSummaryStatusSkipped() *ControlSummary {
	return &ControlSummary{
		Name:   "control-skipped",
		Status: apis.StatusSkipped,
		StatusInfo: apis.StatusInfo{
			InnerStatus: apis.InfoStatusSkipped,
			InnerInfo:   "no host sensor flag",
		},
		Score: 0,
	}
}

func mockControlSummaryNoInnerStatus() *ControlSummary {
	return &ControlSummary{
		Name:       "control-irrelevant",
		Status:     apis.StatusSkipped,
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

func mockControlSummaryStatusIrelevant() *ControlSummary {
	return &ControlSummary{
		Name:   "control-irrelevant",
		Status: apis.StatusSkipped,
		StatusInfo: apis.StatusInfo{
			InnerStatus: apis.StatusIrrelevant,
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

func mockResourceCountersExcludePass() *ResourceCounters {
	return &ResourceCounters{
		PassedResources:   4,
		ExcludedResources: 3,
	}
}

func mockResourceCountersExcludeFailPass() *ResourceCounters {
	return &ResourceCounters{
		PassedResources:   4,
		ExcludedResources: 3,
		FailedResources:   5,
	}
}

func mockResourceCountersPass() *ResourceCounters {
	return &ResourceCounters{
		PassedResources: 4,
	}
}

func mockResourceCountersSkipped() *ResourceCounters {
	return &ResourceCounters{}
}

func mockResourceCountersExclude() *ResourceCounters {
	return &ResourceCounters{
		ExcludedResources: 4,
	}
}
