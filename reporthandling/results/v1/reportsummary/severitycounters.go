package reportsummary

type ISeverityCounters interface {
	NumberOfResourcesWithCriticalSeverity() int
	NumberOfResourcesWithHighSeverity() int
	NumberOfResourcesWithMediumSeverity() int
	NumberOfResourcesWithLowSeverity() int
}

func (sc *SeverityCounters) NumberOfResourcesWithCriticalSeverity() int {
	return sc.ResourcesWithCriticalSeverityCounter
}

func (sc *SeverityCounters) NumberOfResourcesWithHighSeverity() int {
	return sc.ResourcesWithHighSeverityCounter
}

func (sc *SeverityCounters) NumberOfResourcesWithMediumSeverity() int {
	return sc.ResourcesWithMediumSeverityCounter
}

func (sc *SeverityCounters) NumberOfResourcesWithLowSeverity() int {
	return sc.ResourcesWithLowSeverityCounter
}
