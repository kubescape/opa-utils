package reportsummary

import "github.com/kubescape/opa-utils/reporthandling/apis"

type ISeverityCounters interface {
	NumberOfCriticalSeverity() int
	NumberOfHighSeverity() int
	NumberOfMediumSeverity() int
	NumberOfLowSeverity() int
	Increase(severity string, amount int)
}

func (sc *SeverityCounters) NumberOfCriticalSeverity() int {
	return sc.CriticalSeverityCounter
}

func (sc *SeverityCounters) NumberOfHighSeverity() int {
	return sc.HighSeverityCounter
}

func (sc *SeverityCounters) NumberOfMediumSeverity() int {
	return sc.MediumSeverityCounter
}

func (sc *SeverityCounters) NumberOfLowSeverity() int {
	return sc.LowSeverityCounter
}

// Increase increments the counter of a given severity by a given amount
func (sc *SeverityCounters) Increase(severity string, amount int) {
	var counterToIncrement *int

	switch severity {
	case apis.SeverityCriticalString:
		counterToIncrement = &sc.CriticalSeverityCounter
	case apis.SeverityHighString:
		counterToIncrement = &sc.HighSeverityCounter
	case apis.SeverityMediumString:
		counterToIncrement = &sc.MediumSeverityCounter
	case apis.SeverityLowString:
		counterToIncrement = &sc.LowSeverityCounter
	// Return without incrementing on unrecognized severities
	default:
		return
	}

	*counterToIncrement += amount
}
