package v2

import (
	helpersv1 "github.com/armosec/opa-utils/reporthandling/helpers/v1"
	"github.com/armosec/opa-utils/reporthandling/results/v1/reportsummary"
	"github.com/armosec/opa-utils/reporthandling/results/v1/resourcesresults"
)

func updateControlsSummaryCounters(resourceResult *resourcesresults.Result, controls map[string]reportsummary.ControlSummary, f *helpersv1.Filters) {
	// update controls counters
	for i := range resourceResult.AssociatedControls {
		controlID := resourceResult.AssociatedControls[i].ControlID
		if controlSummary, ok := controls[controlID]; ok {
			controlSummary.Increase(resourceResult.AssociatedControls[i].GetStatus(f))
			controls[controlID] = controlSummary
		}
	}
}
