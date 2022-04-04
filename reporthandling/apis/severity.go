package apis

const (
	SeverityUnknown  = iota
	SeverityLow      = iota
	SeverityMedium   = iota
	SeverityHigh     = iota
	SeverityCritical = iota
)

func ControlSeverityToString(baseScore float32) string {
	/*
		9+	Critical
		7-8	High
		4-6	Medium
		1-3	Low
		0 Unknown
	*/
	switch ControlSeverityToInt(baseScore) {
	case SeverityCritical:
		return "Critical"
	case SeverityHigh:
		return "High"
	case SeverityMedium:
		return "Medium"
	case SeverityLow:
		return "Low"
	default:
		return "Unknown"
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
