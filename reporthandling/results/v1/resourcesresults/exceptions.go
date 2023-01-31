package resourcesresults

import (
	"github.com/armosec/armoapi-go/armotypes"
	"github.com/kubescape/k8s-interface/workloadinterface"
	"github.com/kubescape/opa-utils/exceptions"
	"github.com/kubescape/opa-utils/reporthandling"
)

type (
	// ResultOption allows for fine-grained tuning of the Result methods.
	ResultOption func(*resultOptions)

	resultOptions struct {
		processor *exceptions.Processor
	}
)

func resultOptionsWithDefaults(opts []ResultOption) *resultOptions {
	o := &resultOptions{}
	for _, apply := range opts {
		apply(o)
	}

	return o
}

// WithExceptionsProcessor allows the SetExceptions method to reuse an already allocated exceptions.Processor.
func WithExceptionsProcessor(processor *exceptions.Processor) ResultOption {
	return func(o *resultOptions) {
		o.processor = processor
	}
}

// SetExceptions add exceptions to result.
//
// If the caller has already instanciated an exceptions processor, the latter may be reused with option "WithExceptionsProcessor(processor)".
func (result *Result) SetExceptions(workload workloadinterface.IMetadata, exceptionsPolicies []armotypes.PostureExceptionPolicy, clusterName string, controls map[string]reporthandling.Control, opts ...ResultOption) {
	if len(exceptionsPolicies) == 0 {
		return
	}

	result.resultOptions = resultOptionsWithDefaults(opts)
	if result.processor == nil {
		// when no processor is specified by the providd options, create a new one
		result.processor = exceptions.NewProcessor()
	}

	for i := range result.AssociatedControls {
		result.AssociatedControls[i].setExceptions(workload, exceptionsPolicies, clusterName, controls[result.AssociatedControls[i].GetID()], result.processor)
	}
}

// SetExceptions add exceptions to result
func (control *ResourceAssociatedControl) setExceptions(workload workloadinterface.IMetadata, exceptionsPolicies []armotypes.PostureExceptionPolicy, clusterName string, c reporthandling.Control, processor *exceptions.Processor) {
	// add exceptions only to failed controls
	// don't add exceptions to passed controls
	if control.GetStatus(nil).IsPassed() {
		return
	}

	for i := range control.ResourceAssociatedRules {
		exceptionsPolicies = processor.ListRuleExceptions(exceptionsPolicies, "", control.GetName(), control.GetID(), "")
		control.ResourceAssociatedRules[i].setExceptions(workload, exceptionsPolicies, clusterName, processor)
		// Update rule status according to exceptions
		control.ResourceAssociatedRules[i].SetStatus(control.Status.Status(), nil)
	}
	// Update control status according to rules status
	control.SetStatus(c)
}

// SetExceptions add exceptions to result
func (rule *ResourceAssociatedRule) setExceptions(workload workloadinterface.IMetadata, exceptionsPolicies []armotypes.PostureExceptionPolicy, clusterName string, processor *exceptions.Processor) {
	// add exceptions only to failed rules
	if !rule.GetStatus(nil).IsFailed() {
		return
	}

	ruleExceptions := processor.ListRuleExceptions(exceptionsPolicies, "", "", "", rule.GetName())
	rule.Exception = processor.GetResourceExceptions(ruleExceptions, workload, clusterName)
}
