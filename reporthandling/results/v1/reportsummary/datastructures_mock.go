package reportsummary

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
	}
}
func mockFrameworkSummaryPassExclude() *FrameworkSummary {
	return &FrameworkSummary{
		Name:             "fw-pass-exclude",
		Score:            0,
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
		ResourceCounters: *mockResourceCountersExcludePass(),
	}
}

func mockControlSummaryPass() *ControlSummary {
	return &ControlSummary{
		Name:             "control-pass",
		Score:            0,
		ResourceCounters: *mockResourceCountersPass(),
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

func mockResourceCountersPass() *ResourceCounters {
	return &ResourceCounters{
		PassedResources: 4,
	}
}

func mockResourceCountersExclude() *ResourceCounters {
	return &ResourceCounters{
		ExcludedResources: 4,
	}
}
