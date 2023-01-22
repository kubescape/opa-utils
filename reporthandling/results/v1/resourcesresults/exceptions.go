package resourcesresults

import (
	"github.com/armosec/armoapi-go/armotypes"
	"github.com/kubescape/k8s-interface/workloadinterface"
	"github.com/kubescape/opa-utils/exceptions"
)

// SetExceptions add exceptions to result
func (result *Result) SetExceptions(workload workloadinterface.IMetadata, exceptionsPolicies []armotypes.PostureExceptionPolicy, clusterName string) {
	if len(exceptionsPolicies) == 0 {
		return
	}

	for i := range result.AssociatedControls {
		result.AssociatedControls[i].setExceptions(workload, exceptionsPolicies, clusterName)
	}
}

// SetExceptions add exceptions to result
func (control *ResourceAssociatedControl) setExceptions(workload workloadinterface.IMetadata, exceptionsPolicies []armotypes.PostureExceptionPolicy, clusterName string) {
	// add exceptions only to failed controls
	if !control.GetStatus(nil).IsFailed() {
		return
	}

	for i := range control.ResourceAssociatedRules {
		exceptionsPolicies = exceptions.ListRuleExceptions(exceptionsPolicies, "", control.GetName(), control.GetID(), "")
		control.ResourceAssociatedRules[i].setExceptions(workload, exceptionsPolicies, clusterName)
	}
}

// SetExceptions add exceptions to result
func (rule *ResourceAssociatedRule) setExceptions(workload workloadinterface.IMetadata, exceptionsPolicies []armotypes.PostureExceptionPolicy, clusterName string) {
	// add exceptions only to failed rules
	if !rule.GetStatus(nil).IsFailed() {
		return
	}

	ruleExceptions := exceptions.ListRuleExceptions(exceptionsPolicies, "", "", "", rule.GetName())
	rule.Exception = exceptions.GetResourceExceptions(ruleExceptions, workload, clusterName)
}
