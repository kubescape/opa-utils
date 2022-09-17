package reportsummary

import "github.com/kubescape/opa-utils/reporthandling/apis"

type ISeverityCounters interface {
	NumberOfResourcesWithCriticalSeverity() int
	NumberOfResourcesWithHighSeverity() int
	NumberOfResourcesWithMediumSeverity() int
	NumberOfResourcesWithLowSeverity() int
	Increase(severity string, amount int)
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

// Increase increments the counter of a given severity by a given amount
func (sc *SeverityCounters) Increase(severity string, amount int) {
	var counterToIncrement *int

	switch severity {
	case apis.SeverityCriticalString:
		counterToIncrement = &sc.ResourcesWithCriticalSeverityCounter
	case apis.SeverityHighString:
		counterToIncrement = &sc.ResourcesWithHighSeverityCounter
	case apis.SeverityMediumString:
		counterToIncrement = &sc.ResourcesWithMediumSeverityCounter
	case apis.SeverityLowString:
		counterToIncrement = &sc.ResourcesWithLowSeverityCounter
	// Return without incrementing on unrecognized severities
	default:
		return
	}

	*counterToIncrement++
}
