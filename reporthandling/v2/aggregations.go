package v2

import (
	helpersv1 "github.com/armosec/opa-utils/reporthandling/helpers/v1"
	"github.com/armosec/opa-utils/reporthandling/results/v1/reportsummary"
)

func (postureReport *PostureReport) GenarateSummary(f *helpersv1.Filters) *reportsummary.SummaryDetails {
	reportSummary := reportsummary.SummaryDetails{}
	for i := range postureReport.Results {
		allControls := postureReport.Results[i].ListAllControls(nil)
		// for i
		// reportSummary
	}
}
