package resourcesresults

import (
	"github.com/kubescape/opa-utils/reporthandling/apis"
	helpersv1 "github.com/kubescape/opa-utils/reporthandling/helpers/v1"
)

// GetName get rule name
func (rule *ResourceAssociatedRule) GetName() string {
	return rule.Name
}

// SetName set rule name
func (rule *ResourceAssociatedRule) SetName(n string) {
	rule.Name = n
}

// =============================== Status ====================================

// GetStatus get rule status
func (rule *ResourceAssociatedRule) GetStatus(f *helpersv1.Filters) apis.IStatus {
	return helpersv1.NewStatus(rule.Status)
}

// GetSubStatus get rule sub status
func (rule *ResourceAssociatedRule) GetSubStatus() apis.IStatus {
	return helpersv1.NewStatus(rule.SubStatus)
}

// SetStatus set rule status and sub status
func (rule *ResourceAssociatedRule) SetStatus(s apis.ScanningStatus, f *helpersv1.Filters) {
	if s == apis.StatusFailed || s == apis.SubStatusException {
		if f != nil {
			if len(f.FilterExceptions(rule.Exception)) > 0 {
				rule.Status = apis.StatusPassed
				rule.SubStatus = apis.SubStatusException
				return
			} else {
				rule.Status = apis.StatusFailed
				return
			}
		} else {
			if len(rule.Exception) > 0 {
				rule.Status = apis.StatusPassed
				rule.SubStatus = apis.SubStatusException
				return
			} else {
				rule.Status = apis.StatusFailed
				return
			}
		}
	}
	rule.Status = apis.StatusPassed
}
