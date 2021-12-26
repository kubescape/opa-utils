package resourcesresults

import (
	"github.com/armosec/opa-utils/reporthandling/apis"
	helpersv1 "github.com/armosec/opa-utils/reporthandling/helpers/v1"
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
func (rule *ResourceAssociatedRule) Status(f *helpersv1.Filters) apis.ScanningStatus {
	if f != nil {
		if len(f.FilterExceptions(rule.Exception)) > 0 {
			return apis.StatusExcluded
		}
	}
	return apis.StatusFailed
}

// IsPassed did this rule pass
func (rule *ResourceAssociatedRule) IsPassed(f *helpersv1.Filters) bool {
	return rule.Status(f) == apis.StatusPassed
}

// IsFailed did this rule fail
func (rule *ResourceAssociatedRule) IsFailed(f *helpersv1.Filters) bool {
	return rule.Status(f) == apis.StatusFailed
}

// IsExcluded is this rule excluded
func (rule *ResourceAssociatedRule) IsExcluded(f *helpersv1.Filters) bool {
	return rule.Status(f) == apis.StatusExcluded
}

// IsSkipped was this rule skipped
func (rule *ResourceAssociatedRule) IsSkipped(f *helpersv1.Filters) bool {
	return rule.Status(f) == apis.StatusSkipped
}
