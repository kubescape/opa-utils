package apis

const (
	SeverityCriticalString = "Critical"
	SeverityHighString     = "High"
	SeverityMediumString   = "Medium"
	SeverityLowString      = "Low"
	SeverityUnknownString  = "Unknown"
)

const (
	SeverityUnknown  = iota
	SeverityLow      = iota
	SeverityMedium   = iota
	SeverityHigh     = iota
	SeverityCritical = iota
)
const NumberOfSeverities = 5

// GetSupportedSeverities returns a slice of supported severities
func GetSupportedSeverities() []string {
	return []string{SeverityLowString, SeverityMediumString, SeverityHighString, SeverityCriticalString}
}

func ControlSeverityToString(baseScore float32) string {
	/*
		9+	Critical
		7-8	High
		4-6	Medium
		1-3	Low
		0 Unknown
	*/
	return SeverityNumberToString(ControlSeverityToInt(baseScore))
}

func SeverityNumberToString(severityNumber int) string {
	/*
		4 Critical
		3 High
		2 Medium
		1 Low
		0 Unknown
	*/
	switch severityNumber {
	case SeverityCritical:
		return SeverityCriticalString
	case SeverityHigh:
		return SeverityHighString
	case SeverityMedium:
		return SeverityMediumString
	case SeverityLow:
		return SeverityLowString
	default:
		return SeverityUnknownString
	}

}

func ControlSeverityToInt(baseScore float32) int {
	/*
		0   Unknown
		1	Low
		2	Medium
		3	High
		4	Critical
	*/
	if baseScore >= 9 {
		return SeverityCritical
	}
	if baseScore >= 7 {
		return SeverityHigh
	}
	if baseScore >= 4 {
		return SeverityMedium
	}
	if baseScore >= 1 {
		return SeverityLow
	}
	return SeverityUnknown
}
