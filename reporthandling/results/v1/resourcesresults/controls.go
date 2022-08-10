package resourcesresults

import (
	"github.com/kubescape/opa-utils/reporthandling/apis"
	helpersv1 "github.com/kubescape/opa-utils/reporthandling/helpers/v1"
)

// GetID get control ID
func (control *ResourceAssociatedControl) GetID() string {
	return control.ControlID
}

// SetID set control ID
func (control *ResourceAssociatedControl) SetID(id string) {
	control.ControlID = id
}

// GetID get control ID
func (control *ResourceAssociatedControl) GetName() string {
	return control.Name
}

// SetID set control ID
func (control *ResourceAssociatedControl) SetName(name string) {
	control.Name = name
}

// =============================== Status ====================================

// Status get control status
func (control *ResourceAssociatedControl) GetStatus(f *helpersv1.Filters) apis.IStatus {
	status := apis.StatusPassed // if len(control.ResourceAssociatedRules) == 0 the resource passed
	for i := range control.ResourceAssociatedRules {
		status = apis.Compare(status, control.ResourceAssociatedRules[i].GetStatus(f).Status())
	}
	return helpersv1.NewStatus(status)
}

// ListRules return list of rules
func (control *ResourceAssociatedControl) ListRules() []ResourceAssociatedRule {
	return control.ResourceAssociatedRules
}
