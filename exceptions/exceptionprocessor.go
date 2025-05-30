package exceptions

import (
	"strings"

	"github.com/armosec/armoapi-go/identifiers"

	"github.com/kubescape/k8s-interface/workloadinterface"
	"github.com/kubescape/opa-utils/objectsenvelopes"
	"github.com/kubescape/opa-utils/reporthandling"

	"github.com/armosec/armoapi-go/armotypes"
)

// Processor processes exceptions.
type Processor struct {
	*comparator
	designatorCache *designatorCache
}

func NewProcessor() *Processor {
	return &Processor{
		comparator:      newComparator(),
		designatorCache: newDesignatorCache(),
	}
}

// SetFrameworkExceptions add exceptions to framework report
func (p *Processor) SetFrameworkExceptions(frameworkReport *reporthandling.FrameworkReport, exceptionsPolicies []armotypes.PostureExceptionPolicy, clusterName string) {
	for c := range frameworkReport.ControlReports {
		p.SetControlExceptions(&frameworkReport.ControlReports[c], exceptionsPolicies, clusterName, frameworkReport.Name)
	}
}

// SetControlExceptions add exceptions to control report
func (p *Processor) SetControlExceptions(controlReport *reporthandling.ControlReport, exceptionsPolicies []armotypes.PostureExceptionPolicy, clusterName, frameworkName string) {
	for r := range controlReport.RuleReports {
		p.SetRuleExceptions(&controlReport.RuleReports[r], exceptionsPolicies, clusterName, frameworkName, controlReport.ControlID)
	}
}

// SetRuleExceptions add exceptions to rule report
func (p *Processor) SetRuleExceptions(ruleReport *reporthandling.RuleReport, exceptionsPolicies []armotypes.PostureExceptionPolicy, clusterName, frameworkName, controlID string) {
	// adding exceptions to the rules
	ruleExceptions := p.ListRuleExceptions(exceptionsPolicies, frameworkName, controlID, ruleReport.Name)
	p.SetRuleResponsExceptions(ruleReport.RuleResponses, ruleExceptions, clusterName)
}

// SetRuleExceptions add exceptions to rule respons structure
func (p *Processor) SetRuleResponsExceptions(results []reporthandling.RuleResponse, ruleExceptions []armotypes.PostureExceptionPolicy, clusterName string) {
	if len(ruleExceptions) == 0 {
		return
	}

	for i := range results {
		workloads := alertObjectToWorkloads(&results[i].AlertObject)
		if len(workloads) == 0 {
			continue
		}

		for w := range workloads {
			if exceptions := p.GetResourceExceptions(ruleExceptions, workloads[w], clusterName); len(exceptions) > 0 {
				results[i].Exception = &exceptions[0]
			}
		}

		results[i].RuleStatus = results[i].GetStatus()
	}
}

func (p *Processor) ListRuleExceptions(exceptionPolicies []armotypes.PostureExceptionPolicy, frameworkName, controlID, ruleName string) []armotypes.PostureExceptionPolicy {
	ruleExceptions := make([]armotypes.PostureExceptionPolicy, 0, len(exceptionPolicies))

	for i := range exceptionPolicies {
		if p.ruleHasExceptions(&exceptionPolicies[i], frameworkName, controlID, ruleName) {
			ruleExceptions = append(ruleExceptions, exceptionPolicies[i])
		}
	}

	return ruleExceptions[:len(ruleExceptions):len(ruleExceptions)]

}

func (p *Processor) ruleHasExceptions(exceptionPolicy *armotypes.PostureExceptionPolicy, frameworkName, controlID, ruleName string) bool {
	if len(exceptionPolicy.PosturePolicies) == 0 {
		return true // empty policy -> apply all
	}

	for _, posturePolicy := range exceptionPolicy.PosturePolicies {
		if posturePolicy.FrameworkName == "" && posturePolicy.ControlID == "" && posturePolicy.RuleName == "" {
			return true // empty policy -> apply all
		}
		if posturePolicy.FrameworkName != "" && frameworkName != "" && !(strings.EqualFold(posturePolicy.FrameworkName, frameworkName) || p.regexCompareI(posturePolicy.FrameworkName, frameworkName)) {
			continue // policy does not match
		}
		if posturePolicy.ControlID != "" && controlID != "" && !(strings.EqualFold(posturePolicy.ControlID, controlID) || p.regexCompareI(posturePolicy.ControlID, controlID)) {
			continue // policy does not match
		}
		if posturePolicy.RuleName != "" && ruleName != "" && !(strings.EqualFold(posturePolicy.RuleName, ruleName) || p.regexCompareI(posturePolicy.RuleName, ruleName)) {
			continue // policy does not match
		}

		return true // policies match
	}

	return false

}

func alertObjectToWorkloads(obj *reporthandling.AlertObject) []workloadinterface.IMetadata {
	resources := make([]workloadinterface.IMetadata, 0, len(obj.K8SApiObjects)+1)

	for i := range obj.K8SApiObjects {
		r := objectsenvelopes.NewObject(obj.K8SApiObjects[i])
		if r == nil {
			continue
		}

		resources = append(resources, r)
		/*
			ns : = r.GetNamespace()
			if ns != "" {
				// TODO - handle empty namespace
			}
		*/
	}

	if obj.ExternalObjects != nil {
		if r := objectsenvelopes.NewObject(obj.ExternalObjects); r != nil {
			// TODO - What about linked objects?
			resources = append(resources, r)
		}
	}

	return resources[:len(resources):len(resources)]
}

// GetResourceException get exceptions of single resource
func (p *Processor) GetResourceExceptions(ruleExceptions []armotypes.PostureExceptionPolicy, workload workloadinterface.IMetadata, clusterName string) []armotypes.PostureExceptionPolicy {
	// no pre-allocation since most of the time it's empty or has only one element
	var postureExceptionPolicy []armotypes.PostureExceptionPolicy

	for _, ruleException := range ruleExceptions {
		for _, resourceToPin := range ruleException.Resources {
			resource := resourceToPin
			if p.hasException(clusterName, &resource, workload) {
				postureExceptionPolicy = append(postureExceptionPolicy, ruleException)
			}
		}
	}

	return postureExceptionPolicy
}

// compareMetadata - compare namespace and kind
func (p *Processor) hasException(clusterName string, designator *identifiers.PortalDesignator, workload workloadinterface.IMetadata) bool {
	var attributes identifiers.AttributesDesignators
	if attrs, ok := p.designatorCache.Get(designator); ok {
		attributes = attrs
	} else {
		attrs := designator.DigestPortalDesignator()
		attributes = attrs
		p.designatorCache.Set(designator, attributes)
	}

	if attributes.GetCluster() == "" && attributes.GetNamespace() == "" && attributes.GetKind() == "" && attributes.GetName() == "" && attributes.GetResourceID() == "" && attributes.GetPath() == "" && len(attributes.GetLabels()) == 0 {
		return false // if designators are empty
	}

	if attributes.GetCluster() != "" && !p.compareCluster(attributes.GetCluster(), clusterName) { // TODO - where do we receive cluster name from?
		return false // cluster name does not match
	}

	if isTypeRegoResponseVector(workload) {
		if p.iterateRegoResponseVector(workload, attributes) {
			return true
		}
		// otherwise, continue to check the base object
	}
	return p.metadataHasException(workload, attributes)

}

func (p *Processor) metadataHasException(workload workloadinterface.IMetadata, attributes identifiers.AttributesDesignators) bool {

	if attributes.GetNamespace() != "" && !p.compareNamespace(workload, attributes.GetNamespace()) {
		return false // namespaces do not match
	}

	if attributes.GetKind() != "" && !p.compareKind(workload, attributes.GetKind()) {
		return false // kinds do not match
	}

	if attributes.GetName() != "" && !p.compareName(workload, attributes.GetName()) {
		return false // names do not match
	}

	if attributes.GetResourceID() != "" && !p.compareResourceID(workload, attributes.GetResourceID()) {
		return false // names do not match
	}

	if attributes.GetPath() != "" && !p.comparePath(workload, attributes.GetPath()) {
		return false // paths do not match
	}

	if isTypeWorkload(workload) && len(attributes.GetLabels()) > 0 {
		if !p.compareLabels(workload, attributes.GetLabels()) && !p.compareAnnotations(workload, attributes.GetLabels()) {
			return false // labels nor annotations do not match
		}
	}
	return true
}

func (p *Processor) iterateRegoResponseVector(workload workloadinterface.IMetadata, attributes identifiers.AttributesDesignators) bool {
	v := objectsenvelopes.NewRegoResponseVectorObject(workload.GetObject())
	for _, r := range v.GetRelatedObjects() {
		if p.metadataHasException(r, attributes) {
			return true
		}
	}
	return false
}
