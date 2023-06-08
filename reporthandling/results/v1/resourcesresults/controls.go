package resourcesresults

import (
	"github.com/kubescape/opa-utils/reporthandling"
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

// isOldControl checks if a control has no status set - this is for backward compatibility with old reports
func (control *ResourceAssociatedControl) isOldControl() bool {
	return control.Status.InnerStatus == "" && control.Status.InnerInfo == ""
}

// Status get control status
func (control *ResourceAssociatedControl) GetStatus(f *helpersv1.Filters) apis.IStatus {
	if control.isOldControl() {
		status := apis.StatusPassed
		for i := range control.ResourceAssociatedRules {
			// we must set the status since this is an old flow and the calculation of the status wasn't done
			control.ResourceAssociatedRules[i].SetStatus(control.ResourceAssociatedRules[i].GetStatus(f).Status(), f)
			status = apis.Compare(status, control.ResourceAssociatedRules[i].GetStatus(f).Status())
		}
		return helpersv1.NewStatus(status)
	}

	return &control.Status
}

// GetSubStatus get control sub status
func (control *ResourceAssociatedControl) GetSubStatus() apis.ScanningSubStatus {
	if control.isOldControl() {
		status := apis.StatusPassed
		subStatus := apis.SubStatusUnknown

		for i := range control.ResourceAssociatedRules {
			control.ResourceAssociatedRules[i].SetStatus(control.ResourceAssociatedRules[i].GetStatus(nil).Status(), nil)
			status, subStatus = apis.CompareStatusAndSubStatus(status, control.ResourceAssociatedRules[i].GetStatus(nil).Status(), subStatus, control.ResourceAssociatedRules[i].GetSubStatus())
		}
		return subStatus
	}
	return control.Status.SubStatus
}

// SetStatus set control status and sub status
/*
	SetStatus set control status and sub status according to the following logic:
	1. Calculate control status with all the resource associated rules:
		1.1 if the status is failed and the control contains attributes of actionRequired: requires/manual review,
			the status is skipped and the sub status is requires/manual review
		1.2 if the control contains attributes of actionRequired: configuration and the configuration is not set,
			the status is skipped and the sub status is configuration
*/
func (control *ResourceAssociatedControl) SetStatus(c reporthandling.Control) {
	// calculate the status with all the resource associated rules
	status := apis.StatusPassed
	subStatus := apis.SubStatusUnknown
	statusInfo := ""
	for i := range control.ResourceAssociatedRules {
		status, subStatus = apis.CompareStatusAndSubStatus(status, control.ResourceAssociatedRules[i].GetStatus(nil).Status(), subStatus, control.ResourceAssociatedRules[i].GetSubStatus())
	}
	actionRequiredStr := c.GetActionRequiredAttribute()
	if actionRequiredStr == "" {
		control.Status.InnerStatus = status
		control.Status.SubStatus = subStatus
		return
	}

	// If the control type is requires review, the status is failed and the sub status is requires review
	actionRequired := apis.ScanningSubStatus(actionRequiredStr)
	if status == apis.StatusFailed && actionRequired == apis.SubStatusRequiresReview {
		status = apis.StatusFailed
		subStatus = apis.SubStatusRequiresReview
		statusInfo = string(apis.SubStatusRequiresReviewInfo)
	}

	// If the control type is manual review, the status is skipped and the sub status is manual review
	if status == apis.StatusFailed && actionRequired == apis.SubStatusManualReview {
		status = apis.StatusSkipped
		subStatus = apis.SubStatusManualReview
		statusInfo = string(apis.SubStatusManualReviewInfo)
	}

	// If the control type is configuration and the configuration is not set, the status is skipped and the sub status is configuration
	if actionRequired == apis.SubStatusConfiguration && controlMissingAllConfigurations(control) {
		status = apis.StatusSkipped
		subStatus = apis.SubStatusConfiguration
		statusInfo = string(apis.SubStatusConfigurationInfo)
	}

	control.Status.InnerStatus = status
	control.Status.InnerInfo = statusInfo
	control.Status.SubStatus = subStatus

}

// ListRules return list of rules
func (control *ResourceAssociatedControl) ListRules() []ResourceAssociatedRule {
	return control.ResourceAssociatedRules
}

// controlMissingConfiguration returns true if all control configurations are missing
func controlMissingAllConfigurations(control *ResourceAssociatedControl) bool {
	for _, rule := range control.ResourceAssociatedRules {
		if len(rule.ControlConfigurations) == 0 {
			return true
		}
		for _, configuration := range rule.ControlConfigurations {
			if len(configuration) != 0 {
				return false
			}
		}
	}
	return true
}
