package reportsummary

func MockSummaryDetails() *SummaryDetails {
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
		},
	}
}
func mockFrameworkSummaryPassExclude() *FrameworkSummary {
	return &FrameworkSummary{
		Name:    "fw-pass-exclude",
		Score:   0,
		Version: "utnitest",
		Controls: map[string]ControlSummary{
			"C-0003": *mockControlSummaryExcludePass(),
		},
	}
}
func mockFrameworkSummaryFailExclude() *FrameworkSummary {
	return &FrameworkSummary{
		Name:    "fw-failed",
		Score:   0,
		Version: "utnitest",
		Controls: map[string]ControlSummary{
			"C-0001": *mockControlSummaryFailPass(),
			"C-0002": *mockControlSummaryFailPassExclude(),
		},
	}
}
func mockFrameworkSummaryFailPass() *FrameworkSummary {
	return &FrameworkSummary{
		Name:    "fw-failed",
		Score:   0,
		Version: "utnitest",
		Controls: map[string]ControlSummary{
			"C-0001": *mockControlSummaryFailPass(),
			"C-0003": *mockControlSummaryExcludePass(),
		},
	}
}
func mockControlSummaryExcludePass() *ControlSummary {
	return &ControlSummary{
		Name:             "control-exclude-pass",
		Score:            0,
		ResourceCounters: *mockResourceCountersExcludePass(),
	}
}

func mockControlSummaryFailPass() *ControlSummary {
	return &ControlSummary{
		Name:             "control-fail-pass",
		Score:            0,
		ResourceCounters: *mockResourceCountersFailPass(),
	}
}
func mockControlSummaryFailPassExclude() *ControlSummary {
	return &ControlSummary{
		Name:             "control-fail-pass-exclude",
		Score:            0,
		ResourceCounters: *mockResourceCountersExcludeFailPass(),
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
