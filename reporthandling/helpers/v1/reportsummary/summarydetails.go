package reportsummary

// =================================== Counters ============================================

// NumberOfExcluded get the number of excluded resources
func (summaryDetails *SummaryDetails) NumberOfExcluded() int {
	return summaryDetails.ResourceCounters.NumberOfExcluded()
}

// NumberOfPassed get the number of passed resources
func (summaryDetails *SummaryDetails) NumberOfPassed() int {
	return summaryDetails.ResourceCounters.NumberOfPassed()
}

// NumberOfSkipped get the number of skipped resources
func (summaryDetails *SummaryDetails) NumberOfSkipped() int {
	return summaryDetails.ResourceCounters.NumberOfSkipped()
}

// NumberOfFailed get the number of failed resources
func (summaryDetails *SummaryDetails) NumberOfFailed() int {
	return summaryDetails.ResourceCounters.NumberOfFailed()
}

// NumberOfAll get the number of all resources
func (summaryDetails *SummaryDetails) NumberOfAll() int {
	return summaryDetails.ResourceCounters.NumberOfAll()
}

// =================================== Setters ============================================

// setNumberOfFailed set the number of failed resources
func (summaryDetails *SummaryDetails) setNumberOfFailed(n int) {
	summaryDetails.FailedResources = n
}

// setNumberOfFailed set the number of passed resources
func (summaryDetails *SummaryDetails) setNumberOfPassed(n int) {
	summaryDetails.PassedResources = n
}

// setNumberOfFailed set the number of excluded resources
func (summaryDetails *SummaryDetails) setNumberOfExcluded(n int) {
	summaryDetails.ExcludedResources = n
}

// setNumberOfFailed set the number of skipped resources
func (summaryDetails *SummaryDetails) setNumberOfSkipped(n int) {
	summaryDetails.SkippedResources = n
}
