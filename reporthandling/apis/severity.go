package apis

func ControlSeverityToString(baseScore float32) string {
	/*
		9+	Critical
		7-8	High
		4-6	Medium
		1-3	Low
		0 Unknown
	*/
	if baseScore >= 9 {
		return "Critical"
	}
	if baseScore >= 7 {
		return "High"
	}
	if baseScore >= 4 {
		return "Medium"
	}
	if baseScore >= 1 {
		return "Low"
	}
	return "Unknown"
}
