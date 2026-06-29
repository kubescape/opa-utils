package exceptions

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/armosec/armoapi-go/identifiers"

	"github.com/kubescape/k8s-interface/workloadinterface"
	"github.com/kubescape/opa-utils/objectsenvelopes"
	"github.com/kubescape/opa-utils/reporthandling"

	"github.com/armosec/armoapi-go/armotypes"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

// rexContainerPath matches "containers[N]" and "initContainers[N]" in a
// FailedPath so we can resolve the container index to a name.
var rexContainerPath = regexp.MustCompile(`(initC|c)ontainers\[(\d+)\]`)

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
			// Resolve which containers actually produced the finding so that a
			// containerName exception is only applied when the excepted container
			// is the one that failed, not just any container in the pod.
			failingContainerNames := extractFailingContainerNames(results[i].FailedPaths, workloads[w])
			if exceptions := p.getResourceExceptions(ruleExceptions, workloads[w], clusterName, failingContainerNames); len(exceptions) > 0 {
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

// GetResourceExceptions returns the exception policies that match workload.
// It checks container membership across the whole workload; use
// SetRuleResponsExceptions when FailedPaths are available for precise matching.
func (p *Processor) GetResourceExceptions(ruleExceptions []armotypes.PostureExceptionPolicy, workload workloadinterface.IMetadata, clusterName string) []armotypes.PostureExceptionPolicy {
	return p.getResourceExceptions(ruleExceptions, workload, clusterName, nil)
}

func (p *Processor) getResourceExceptions(ruleExceptions []armotypes.PostureExceptionPolicy, workload workloadinterface.IMetadata, clusterName string, failingContainerNames []string) []armotypes.PostureExceptionPolicy {
	// no pre-allocation since most of the time it's empty or has only one element
	var postureExceptionPolicy []armotypes.PostureExceptionPolicy

	for _, ruleException := range ruleExceptions {
		exception := ruleException
		// objectSelector is an additional, policy-level workload-matching axis,
		// ANDed with the per-designator Resources matching below: when present it
		// must match the workload's labels for any designator match to count.
		if !p.matchesObjectSelector(&exception, workload) {
			continue
		}
		for _, resourceToPin := range exception.Resources {
			resource := resourceToPin
			if p.hasException(clusterName, &resource, workload, failingContainerNames) {
				postureExceptionPolicy = append(postureExceptionPolicy, exception)
			}
		}
	}

	return postureExceptionPolicy
}

// matchesObjectSelector reports whether the exception's ObjectSelector matches the
// workload's labels. A nil or empty selector imposes no label constraint and the
// label axis is skipped (returns true).
//
// This is the deliberate inverse of metav1.LabelSelectorAsSelector, whose conversion
// disagrees with the SecurityException intent: a nil selector there yields
// labels.Nothing() and an empty selector yields labels.Everything() — the latter would
// silently suppress findings on every workload. We therefore guard for both nil and
// empty before converting. A malformed selector matches nothing (returns false) rather
// than degrading into match-all.
func (p *Processor) matchesObjectSelector(exceptionPolicy *armotypes.PostureExceptionPolicy, workload workloadinterface.IMetadata) bool {
	sel := exceptionPolicy.ObjectSelector.ToMetaV1()
	if sel == nil || (len(sel.MatchLabels) == 0 && len(sel.MatchExpressions) == 0) {
		return true // no label constraint
	}

	selector, err := metav1.LabelSelectorAsSelector(sel)
	if err != nil {
		return false // malformed selector: do not match every workload
	}

	workloadLabels := workloadinterface.NewWorkloadObj(workload.GetObject()).GetLabels()
	return selector.Matches(labels.Set(workloadLabels))
}

// RegexCompareControlID reports whether pattern case-insensitively matches target.
func (p *Processor) RegexCompareControlID(pattern, target string) bool {
	return p.regexCompareI(pattern, target)
}

// MatchesCluster reports whether the designator's cluster constraint matches clusterName.
// A nil designator or empty cluster field matches any cluster.
func (p *Processor) MatchesCluster(designator *identifiers.PortalDesignator, clusterName string) bool {
	if designator == nil {
		return true
	}
	return p.matchesCluster(p.getAttributes(designator), clusterName)
}

// getAttributes returns digested attributes, using the cache when available.
func (p *Processor) getAttributes(designator *identifiers.PortalDesignator) identifiers.AttributesDesignators {
	if attrs, ok := p.designatorCache.Get(designator); ok {
		return attrs
	}
	attrs := designator.DigestPortalDesignator()
	p.designatorCache.Set(designator, attrs)
	return attrs
}

// matchesCluster checks the cluster constraint against pre-digested attributes.
func (p *Processor) matchesCluster(attributes identifiers.AttributesDesignators, clusterName string) bool {
	cluster := attributes.GetCluster()
	if cluster == "" {
		return true
	}
	return p.compareCluster(cluster, clusterName)
}

func (p *Processor) hasException(clusterName string, designator *identifiers.PortalDesignator, workload workloadinterface.IMetadata, failingContainerNames []string) bool {
	attributes := p.getAttributes(designator)

	if attributes.GetCluster() == "" && attributes.GetNamespace() == "" && attributes.GetKind() == "" && attributes.GetName() == "" && attributes.GetResourceID() == "" && attributes.GetPath() == "" && len(attributes.GetLabels()) == 0 {
		return false // if designators are empty
	}

	if !p.matchesCluster(attributes, clusterName) {
		return false // cluster name does not match
	}

	if isTypeRegoResponseVector(workload) {
		if p.iterateRegoResponseVector(workload, attributes, failingContainerNames) {
			return true
		}
		// If containerName is in the designator, stop here: the base
		// RegoResponseVector object is not a workload, so container membership
		// cannot be verified on it. Falling through would silently skip the
		// container check and produce false positives.
		if _, ok := attributes.GetLabels()[identifiers.AttributeContainerName]; ok {
			return false
		}
		// otherwise, continue to check the base object
	}
	return p.metadataHasException(workload, attributes, failingContainerNames)

}

func (p *Processor) metadataHasException(workload workloadinterface.IMetadata, attributes identifiers.AttributesDesignators, failingContainerNames []string) bool {

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

	if isTypeWorkload(workload) {
		allLabels := attributes.GetLabels()
		containerName, hasContainerName := allLabels[identifiers.AttributeContainerName]

		// Build a label map with containerName stripped out so it is not
		// treated as a Kubernetes label during label/annotation comparison.
		labelsWithoutContainer := allLabels
		if hasContainerName {
			labelsWithoutContainer = make(map[string]string, len(allLabels)-1)
			for k, v := range allLabels {
				if k != identifiers.AttributeContainerName {
					labelsWithoutContainer[k] = v
				}
			}
		}

		if len(labelsWithoutContainer) > 0 {
			if !p.compareLabels(workload, labelsWithoutContainer) && !p.compareAnnotations(workload, labelsWithoutContainer) {
				return false // labels nor annotations do not match
			}
		}

		if hasContainerName && !p.compareContainerName(workload, containerName, failingContainerNames) {
			return false // container name does not match
		}
	}

	return true
}

func (p *Processor) iterateRegoResponseVector(workload workloadinterface.IMetadata, attributes identifiers.AttributesDesignators, failingContainerNames []string) bool {
	v := objectsenvelopes.NewRegoResponseVectorObject(workload.GetObject())
	for _, r := range v.GetRelatedObjects() {
		if p.metadataHasException(r, attributes, failingContainerNames) {
			return true
		}
	}
	return false
}

// extractFailingContainerNames parses paths like "spec.containers[0].…" or
// "spec.template.spec.initContainers[1].…" to find which containers produced
// the finding, then returns their names from the workload spec. When the
// FailedPaths contain no container indices (e.g. pod-level findings) the
// returned slice is nil and compareContainerName falls back to checking all
// containers in the workload.
//
// For RegoResponseVector objects the vector itself carries no containers; the
// containers live in the related objects. We recurse into each related workload
// so that container-index resolution still works for vector-based findings.
func extractFailingContainerNames(paths []string, workload workloadinterface.IMetadata) []string {
	if len(paths) == 0 {
		return nil
	}

	if isTypeRegoResponseVector(workload) {
		v := objectsenvelopes.NewRegoResponseVectorObject(workload.GetObject())
		seen := make(map[string]struct{})
		for _, r := range v.GetRelatedObjects() {
			for _, name := range extractFailingContainerNames(paths, r) {
				seen[name] = struct{}{}
			}
		}
		if len(seen) == 0 {
			return nil
		}
		names := make([]string, 0, len(seen))
		for name := range seen {
			names = append(names, name)
		}
		return names
	}

	wl := workloadinterface.NewWorkloadObj(workload.GetObject())
	containers, _ := wl.GetContainers()
	initContainers, _ := wl.GetInitContainers()
	if len(containers)+len(initContainers) == 0 {
		return nil
	}

	seen := make(map[string]struct{})
	for _, path := range paths {
		for _, m := range rexContainerPath.FindAllStringSubmatch(path, -1) {
			idx, err := strconv.Atoi(m[2])
			if err != nil {
				continue
			}
			if m[1] == "initC" {
				if idx < len(initContainers) {
					seen[initContainers[idx].Name] = struct{}{}
				}
			} else {
				if idx < len(containers) {
					seen[containers[idx].Name] = struct{}{}
				}
			}
		}
	}

	if len(seen) == 0 {
		return nil
	}
	names := make([]string, 0, len(seen))
	for name := range seen {
		names = append(names, name)
	}
	return names
}
