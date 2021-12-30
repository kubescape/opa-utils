package reportsummary

// type IDs interface {
// 	Excluded() []string
// 	Passed() []string
// 	Skipped() []string
// 	Failed() []string
// 	All() []string

// 	Append(status apis.IStatus, ids ...string)
// }

// // =================================== Counters ============================================

// // NumberOfExcluded get the number of excluded resources
// func (rid *resourceIDs) Excluded() []string {
// 	return rid.excludedResources
// }

// // NumberOfPassed get the number of passed resources
// func (rid *resourceIDs) Passed() []string {
// 	return rid.passedResources
// }

// // NumberOfSkipped get the number of skipped resources
// func (rid *resourceIDs) Skipped() []string {
// 	return rid.skippedResources
// }

// // NumberOfFailed get the number of failed resources
// func (rid *resourceIDs) Failed() []string {
// 	return rid.failedResources
// }

// // NumberOfAll get the number of all resources
// func (rid *resourceIDs) All() []string {
// 	l := []string{}
// 	l = append(l, rid.failedResources...)
// 	l = append(l, rid.excludedResources...)
// 	l = append(l, rid.passedResources...)
// 	l = append(l, rid.skippedResources...)
// 	return l
// }

// // =================================== Setters ============================================

// // Append append single resource
// func (rid *resourceIDs) Append(status apis.IStatus, ids ...string) {
// 	switch status.Status() {
// 	case apis.StatusExcluded:
// 		rid.excludedResources = append(rid.excludedResources, ids...)
// 	case apis.StatusFailed:
// 		rid.failedResources = append(rid.failedResources, ids...)
// 	case apis.StatusSkipped:
// 		rid.skippedResources = append(rid.skippedResources, ids...)
// 	case apis.StatusPassed:
// 		rid.passedResources = append(rid.passedResources, ids...)
// 	}
// }

// func (rid *resourceIDs) ListResourcesIDs(f *helpersv1.Filters) *helpersv1.AllLists {
// 	resources := &helpersv1.AllLists{}
// 	for i := range postureReport.Results {
// 		resources.Append(postureReport.Results[i].GetStatus(f).Status(), postureReport.Results[i].GetResourceID())
// 	}
// 	return rid
// }
