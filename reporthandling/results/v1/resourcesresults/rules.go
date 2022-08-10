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

// SetName set rule name
func (rule *ResourceAssociatedRule) GetStatus(f *helpersv1.Filters) apis.IStatus {
	if rule.Status == apis.StatusFailed || rule.Status == apis.StatusExcluded {
		if f != nil {
			if len(f.FilterExceptions(rule.Exception)) > 0 {
				return helpersv1.NewStatus(apis.StatusExcluded)
			}
		} else {
			if len(rule.Exception) > 0 {
				return helpersv1.NewStatus(apis.StatusExcluded)
			}
		}
	}
	if rule.Status == apis.StatusUnknown {
		return helpersv1.NewStatus(apis.StatusIgnored)
	}
	return helpersv1.NewStatus(rule.Status)
}
