package v2

import (
	"github.com/armosec/k8s-interface/workloadinterface"
	"github.com/armosec/opa-utils/objectsenvelopes"
	"github.com/armosec/opa-utils/reporthandling/apis"
	helpersv1 "github.com/armosec/opa-utils/reporthandling/helpers/v1"
	"github.com/armosec/opa-utils/reporthandling/results/v1/reportsummary"
	"github.com/armosec/opa-utils/reporthandling/results/v1/resourcesresults"
)

// Status get the overall scanning status
func (postureReport *PostureReport) GetStatus() *helpersv1.Status {
	return postureReport.SummaryDetails.GetStatus()
}

// =========================================== List resources ====================================

func (postureReport *PostureReport) ListResourcesIDs(f *helpersv1.Filters) *helpersv1.AllLists {
	resources := &helpersv1.AllLists{}
	for i := range postureReport.Results {
		resources.Append(postureReport.Results[i].GetStatus(f).Status(), postureReport.Results[i].GetResourceID())
	}
	return resources
}

// =========================================== List Frameworks ====================================

// ListFrameworksNames list all framework policies summary
func (postureReport *PostureReport) ListFrameworks() *reportsummary.ListPolicies {
	return postureReport.SummaryDetails.ListFrameworks()
}

// ListFrameworksNames list all frameworks names
func (postureReport *PostureReport) ListFrameworksNames() *helpersv1.AllLists {
	return postureReport.SummaryDetails.ListFrameworksNames()
}

// =========================================== List Controls ====================================
// ListControls list all controls policies summary
func (postureReport *PostureReport) ListControls() *reportsummary.ListPolicies {
	return postureReport.SummaryDetails.ListControls()
}

// ListControlsNames list all controls names
func (postureReport *PostureReport) ListControlsNames() *helpersv1.AllLists {
	return postureReport.SummaryDetails.ListControlsNames()
}

// ListControlsIDs list all controls names
func (postureReport *PostureReport) ListControlsIDs() *helpersv1.AllLists {
	return postureReport.SummaryDetails.ListControlsIDs()
}

// ==================================== Resource =============================================

// GetResource get single resource in IMetadata interface representation
func (postureReport *PostureReport) GetResource(resourceID string) workloadinterface.IMetadata {
	for i := range postureReport.Resources {
		if postureReport.Resources[i].ResourceID == resourceID {
			if m, ok := postureReport.Resources[i].Object.(map[string]interface{}); ok {
				return objectsenvelopes.NewObject(m)
			}
			break
		}
	}
	return nil
}

// ResourceStatus get single resource status. If resource not found will return an empty string
func (postureReport *PostureReport) ResourceStatus(resourceID string, f *helpersv1.Filters) apis.IStatus {
	for i := range postureReport.Results {
		if postureReport.Results[i].ResourceID == resourceID {
			return postureReport.Results[i].GetStatus(f)
		}
	}
	return helpersv1.NewStatus(apis.StatusUnknown)
}

// ResourceResult get the result of a single resource. If resource not found will return nil
func (postureReport *PostureReport) ResourceResult(resourceID string) *resourcesresults.Result {
	for i := range postureReport.Results {
		if postureReport.Results[i].ResourceID == resourceID {
			return &postureReport.Results[i]
		}
	}
	return nil
}

// UpdateSummary get the result of a single resource. If resource not found will return nil
func (postureReport *PostureReport) GenerateSummary() {
	for i := range postureReport.Results {
		postureReport.UpdateSummaryCounters(&postureReport.Results[i])
	}
	postureReport.SummaryDetails.CalculateStatus()
}

// UpdateSummary get the result of a single resource. If resource not found will return nil
func (postureReport *PostureReport) UpdateSummaryCounters(resourceResult *resourcesresults.Result) {

	// update full-summary counter
	updateControlsSummaryCounters(resourceResult, postureReport.SummaryDetails.Controls, nil)

	// update frameworks counters
	for _, framework := range postureReport.SummaryDetails.Frameworks {
		updateControlsSummaryCounters(resourceResult, framework.Controls, &helpersv1.Filters{FrameworkNames: []string{framework.Name}})
	}
}
