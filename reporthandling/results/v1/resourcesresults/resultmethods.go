package resourcesresults

import (
	"github.com/armosec/opa-utils/reporthandling/apis"
	helpersv1 "github.com/armosec/opa-utils/reporthandling/helpers/v1"
)

// SetResourceID set the resource ID
func (result *Result) SetResourceID(resourceID string) {
	result.ResourceID = resourceID
}

// GetResourceID get the resource ID
func (result *Result) GetResourceID() string {
	return result.ResourceID
}

// =============================== Status ====================================

// Status get resource status
func (result *Result) Status(f *helpersv1.Filters) apis.ScanningStatus {
	status := apis.StatusPassed
	for i := range result.AssociatedControls {
		status = apis.Compare(status, result.AssociatedControls[i].Status(f))
	}
	return status
}

// IsPassed did this resource pass
func (result *Result) IsPassed(f *helpersv1.Filters) bool {
	return result.Status(f) == apis.StatusPassed
}

// IsFailed did this resource fail
func (result *Result) IsFailed(f *helpersv1.Filters) bool {
	return result.Status(f) == apis.StatusFailed
}

// IsExcluded is this resource excluded
func (result *Result) IsExcluded(f *helpersv1.Filters) bool {
	return result.Status(f) == apis.StatusExcluded
}

// IsSkipped was this resource skipped
func (result *Result) IsSkipped(f *helpersv1.Filters) bool {
	return result.Status(f) == apis.StatusSkipped
}

// ================================= Listing ==================================

// ListFailedControls return list of failed controls IDs associated to this resource
func (result *Result) ListFailedControls(f *helpersv1.Filters) []string {
	return result.listControls(f, apis.StatusFailed).ListFailed()
}

// ListFailedControls return list of failed controls IDs associated to this resource
func (result *Result) ListPassedControls(f *helpersv1.Filters) []string {
	return result.listControls(f, apis.StatusPassed).ListPassed()
}

// ListExcludedControls return list of excluded controls IDs associated to this resource
func (result *Result) ListExcludedControls(f *helpersv1.Filters) []string {
	return result.listControls(f, apis.StatusExcluded).ListExcluded()
}

// ListAllControls return list of all controls IDs associated to this resource
func (result *Result) ListAllControls(f *helpersv1.Filters) *helpersv1.AllLists {
	return result.listControls(f, "")
}

// ListFailedControls return list of failed controls IDs
func (result *Result) listControls(f *helpersv1.Filters, status apis.ScanningStatus) *helpersv1.AllLists {
	controls := &helpersv1.AllLists{}
	for i := range result.AssociatedControls {
		s := result.AssociatedControls[i].Status(f)
		if status == "" || s == status {
			controls.Append(s, result.AssociatedControls[i].GetID())
		}
	}
	return controls
}
