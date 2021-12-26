package v2

// func (postureReport *PostureReport) GenarateSummary(frameworkNames []string) *reportsummary.SummaryDetails {
// 	reportSummary := reportsummary.SummaryDetails{}
// 	allControls := map[string]*helpersv1.AllLists{}
// 	for i := range postureReport.Results {
// 		if _, ok := allControls[""]; !ok {
// 			allControls[""] = &helpersv1.AllLists{}
// 		}
// 		allControls[""].Update(postureReport.Results[i].ListAllControls(nil))
// 		for _, frameworkName := range frameworkNames {
// 			if _, ok := allControls[frameworkName]; !ok {
// 				allControls[frameworkName] = &helpersv1.AllLists{}
// 			}
// 			filter := &helpersv1.Filters{FrameworkNames: []string{frameworkName}}
// 			allControls[frameworkName].Update(postureReport.Results[i].ListAllControls(filter))
// 		}
// 	}

// 	for fwName := range allControls {
// 		reportSummary.ExcludedResources = len(allControls[fwName].ListExcluded())
// 		reportSummary.FailedResources = len(allControls[fwName].ListFailed())
// 		reportSummary.PassedResources = len(allControls[fwName].ListPassed())
// 		reportSummary.SkippedResources = len(allControls[fwName].ListSkipped())
// 		for c := range allControls[fwName].ListAll() {
// 			reportSummary.Controls = map[string]reportsummary.ControlSummary{}
// 		}
// 	}

// 	return &reportSummary
// }
