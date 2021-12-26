package resourcesresults

import (
	"github.com/armosec/opa-utils/reporthandling/apis"
	helpersv1 "github.com/armosec/opa-utils/reporthandling/helpers/v1"
)

// GetID get control ID
func (control *ResourceAssociatedControl) GetID() string {
	return control.ControlID
}

// SetID set control ID
func (control *ResourceAssociatedControl) SetID(id string) {
	control.ControlID = id
}

// =============================== Status ====================================

// Status get control status
func (control *ResourceAssociatedControl) Status(f *helpersv1.Filters) apis.ScanningStatus {

	status := apis.StatusPassed
	for i := range control.ResourceAssociatedRules {
		status = apis.Compare(status, control.ResourceAssociatedRules[i].Status(f))
	}
	return status
}

// IsPassed did this control pass
func (control *ResourceAssociatedControl) IsPassed(f *helpersv1.Filters) bool {
	return control.Status(f) == apis.StatusPassed
}

// IsFailed did this control fail
func (control *ResourceAssociatedControl) IsFailed(f *helpersv1.Filters) bool {
	return control.Status(f) == apis.StatusFailed
}

// IsExcluded is this rule excluded
func (control *ResourceAssociatedControl) IsExcluded(f *helpersv1.Filters) bool {
	return control.Status(f) == apis.StatusExcluded
}

// IsSkipped was this rule skipped
func (control *ResourceAssociatedControl) IsSkipped(f *helpersv1.Filters) bool {
	return control.Status(f) == apis.StatusSkipped
}
