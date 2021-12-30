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
func (result *Result) GetStatus(f *helpersv1.Filters) apis.IStatus {
	status := apis.StatusUnknown // Resource was not tested
	for i := range result.AssociatedControls {
		status = apis.Compare(status, result.AssociatedControls[i].GetStatus(f).Status())
	}
	return helpersv1.NewStatus(status)
}

// ================================= Listing ==================================

// ListFailedControls return list of failed controls IDs
func (result *Result) ListControls(f *helpersv1.Filters) *helpersv1.AllLists {
	controls := &helpersv1.AllLists{}
	for i := range result.AssociatedControls {
		s := result.AssociatedControls[i].GetStatus(f).Status()
		controls.Append(s, result.AssociatedControls[i].GetID())
	}
	return controls
}

// ListFailedControls return list of failed controls IDs
func (result *Result) ListControlsIDs(f *helpersv1.Filters) *helpersv1.AllLists {
	controls := &helpersv1.AllLists{}
	for i := range result.AssociatedControls {
		s := result.AssociatedControls[i].GetStatus(f).Status()
		controls.Append(s, result.AssociatedControls[i].GetID())
	}
	return controls
}

// ListFailedControls return list of controls IDs
func (result *Result) ListControlsNames(f *helpersv1.Filters) *helpersv1.AllLists {
	controls := &helpersv1.AllLists{}
	for i := range result.AssociatedControls {
		s := result.AssociatedControls[i].GetStatus(f).Status()
		controls.Append(s, result.AssociatedControls[i].GetName())
	}
	return controls
}

// ListRulesNames return list of rules names
func (result *Result) ListRulesNames(f *helpersv1.Filters) *helpersv1.AllLists {
	rules := &helpersv1.AllLists{}
	for i := range result.AssociatedControls {
		for j := range result.AssociatedControls[i].ResourceAssociatedRules {
			s := result.AssociatedControls[i].GetStatus(f).Status()
			rules.Append(s, result.AssociatedControls[i].ResourceAssociatedRules[j].GetName())
		}
	}
	return rules
}
