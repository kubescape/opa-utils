package exceptions

import (
	"regexp"
	"slices"
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

// rexContainerPath matches "containers[N]", "initContainers[N]" and
// "ephemeralContainers[N]" in a FailedPath so we can resolve the container
// index to a name.
var rexContainerPath = regexp.MustCompile(`(initC|ephemeralC|c)ontainers\[(\d+)\]`)

// Processor processes exceptions.
type Processor struct {
	*comparator
	designatorCache *designatorCache
}

// posturePolicyProcessor owns the concurrency-safe regexp cache used by status
// views. It deliberately has no designator cache because posture-policy tuple
// filtering does not inspect resource designators.
var posturePolicyProcessor = &Processor{comparator: newComparator()}

type posturePolicyMatch struct {
	frameworkNames []string
	matchFramework bool
	controlID      string
	ruleName       string
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
	match := posturePolicyMatch{
		matchFramework: frameworkName != "",
		controlID:      controlID,
		ruleName:       ruleName,
	}
	if match.matchFramework {
		match.frameworkNames = []string{frameworkName}
	}
	return p.filterPostureExceptions(exceptionPolicies, match)
}

// FilterExceptionsByFrameworks uses the shared posture-policy matcher without
// allocating a full resource exception processor for each result status view.
// It returns copies containing only tuples applicable to frameworkNames; an
// empty frameworkNames slice selects framework-agnostic policies.
func FilterExceptionsByFrameworks(exceptionPolicies []armotypes.PostureExceptionPolicy, frameworkNames []string, controlID, ruleName string) []armotypes.PostureExceptionPolicy {
	return posturePolicyProcessor.filterPostureExceptions(exceptionPolicies, posturePolicyMatch{
		frameworkNames: frameworkNames,
		matchFramework: true,
		controlID:      controlID,
		ruleName:       ruleName,
	})
}

func (p *Processor) filterPostureExceptions(exceptionPolicies []armotypes.PostureExceptionPolicy, match posturePolicyMatch) []armotypes.PostureExceptionPolicy {
	ruleExceptions := make([]armotypes.PostureExceptionPolicy, 0, len(exceptionPolicies))

	for i := range exceptionPolicies {
		if len(exceptionPolicies[i].PosturePolicies) == 0 {
			ruleExceptions = append(ruleExceptions, exceptionPolicies[i])
			continue
		}

		matchingPolicies := make([]armotypes.PosturePolicy, 0, len(exceptionPolicies[i].PosturePolicies))
		for j := range exceptionPolicies[i].PosturePolicies {
			if p.posturePolicyMatches(&exceptionPolicies[i].PosturePolicies[j], match) {
				matchingPolicies = append(matchingPolicies, exceptionPolicies[i].PosturePolicies[j])
			}
		}

		if len(matchingPolicies) > 0 {
			filteredException := exceptionPolicies[i]
			filteredException.PosturePolicies = matchingPolicies
			ruleExceptions = append(ruleExceptions, filteredException)
		}
	}

	return ruleExceptions[:len(ruleExceptions):len(ruleExceptions)]

}

func (p *Processor) posturePolicyMatches(posturePolicy *armotypes.PosturePolicy, match posturePolicyMatch) bool {
	if match.matchFramework && posturePolicy.FrameworkName != "" {
		frameworkMatches := false
		for _, frameworkName := range match.frameworkNames {
			if strings.EqualFold(posturePolicy.FrameworkName, frameworkName) || p.regexCompareI(posturePolicy.FrameworkName, frameworkName) {
				frameworkMatches = true
				break
			}
		}
		if !frameworkMatches {
			return false
		}
	}
	if posturePolicy.ControlID != "" && match.controlID != "" && !(strings.EqualFold(posturePolicy.ControlID, match.controlID) || p.regexCompareI(posturePolicy.ControlID, match.controlID)) {
		return false
	}
	if posturePolicy.RuleName != "" && match.ruleName != "" && !(strings.EqualFold(posturePolicy.RuleName, match.ruleName) || p.regexCompareI(posturePolicy.RuleName, match.ruleName)) {
		return false
	}

	return true
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

// GetResourceExceptionsForPaths is GetResourceExceptions for a finding reported on
// the given paths: a containerName exception only matches when it covers the
// containers those paths point at. Paths without a container index keep the
// workload-wide container check.
func (p *Processor) GetResourceExceptionsForPaths(ruleExceptions []armotypes.PostureExceptionPolicy, workload workloadinterface.IMetadata, clusterName string, paths []string) []armotypes.PostureExceptionPolicy {
	return p.getResourceExceptions(ruleExceptions, workload, clusterName, extractFailingContainerNames(paths, workload))
}

func (p *Processor) getResourceExceptions(ruleExceptions []armotypes.PostureExceptionPolicy, workload workloadinterface.IMetadata, clusterName string, failingContainerNames []string) []armotypes.PostureExceptionPolicy {
	// no pre-allocation since most of the time it's empty or has only one element
	var postureExceptionPolicy []armotypes.PostureExceptionPolicy

	for _, ruleException := range ruleExceptions {
		// objectSelector is an additional, policy-level workload-matching axis. It is
		// parsed once per exception here and threaded down into metadataHasException so
		// it is evaluated against the *same* object as the designators — including each
		// related object of a RegoResponseVector — and ANDed with them. A single object
		// must satisfy both the designator and the selector for the exception to apply.
		selector, ok := parseObjectSelector(&ruleException)
		if !ok {
			continue // malformed selector: the exception matches nothing
		}

		// No Resources at all = no scope constraint, matches everywhere - the same
		// "no resources = no scope constraint" convention kubescape's own
		// hasExplicitControlException already applies for manual controls (which have
		// no workload to match a designator against at all). An ObjectSelector, if
		// present, still constrains which workloads this applies to; only the
		// designator axis (cluster/namespace/name/kind/labels/path) is left
		// unconstrained. This intentionally bypasses hasException's guard against an
		// all-empty-fielded PortalDesignator: that guard protects against a populated
		// Resources entry whose fields all happen to be empty (a different footgun),
		// not against Resources itself being an empty list.
		if len(ruleException.Resources) == 0 {
			emptyAttributes := p.getAttributes(&identifiers.PortalDesignator{})
			if p.hasExceptionForAttributes(clusterName, emptyAttributes, workload, failingContainerNames, selector) {
				postureExceptionPolicy = append(postureExceptionPolicy, ruleException)
			}
			continue
		}

		for _, resourceToPin := range ruleException.Resources {
			resource := resourceToPin
			if p.hasException(clusterName, &resource, workload, failingContainerNames, selector) {
				postureExceptionPolicy = append(postureExceptionPolicy, ruleException)
			}
		}
	}

	return postureExceptionPolicy
}

// parseObjectSelector converts an exception's ObjectSelector into a labels.Selector
// for evaluation against a workload's labels. It returns:
//
//   - (nil, true)      when there is no label constraint — a nil OR a non-nil but
//     empty selector. A nil result tells callers to skip the label axis entirely.
//   - (selector, true) for a valid, non-empty constraint.
//   - (nil, false)     for a malformed selector; the exception then matches nothing,
//     never match-all.
//
// The nil/empty guard is deliberate, and the nil case is load-bearing:
// metav1.LabelSelectorAsSelector(nil) yields labels.Nothing(), which would silently
// disable every selector-less exception (the common case — cloud exceptions and
// posture-only CRDs). An empty selector yields labels.Everything(); collapsing it to
// the same "no constraint" nil is equivalent under the AND with the designators, and
// keeps the intent explicit rather than relying on Everything() matching.
//
// A malformed selector is unreachable from a CRD (the apiserver/CEL validate the
// LabelSelector shape) and unset on cloud exceptions, so it is treated as a defensive
// match-nothing here rather than logged; the consumer that decodes the CRD is the
// right layer to surface a bad selector, once per resource instead of once per workload.
func parseObjectSelector(exceptionPolicy *armotypes.PostureExceptionPolicy) (labels.Selector, bool) {
	sel := exceptionPolicy.ObjectSelector.ToMetaV1()
	if sel == nil || (len(sel.MatchLabels) == 0 && len(sel.MatchExpressions) == 0) {
		return nil, true // no label constraint
	}

	selector, err := metav1.LabelSelectorAsSelector(sel)
	if err != nil {
		return nil, false // malformed selector: do not degrade into match-all
	}
	return selector, true
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

func (p *Processor) hasException(clusterName string, designator *identifiers.PortalDesignator, workload workloadinterface.IMetadata, failingContainerNames []string, selector labels.Selector) bool {
	attributes := p.getAttributes(designator)

	if attributes.GetCluster() == "" && attributes.GetNamespace() == "" && attributes.GetKind() == "" && attributes.GetName() == "" && attributes.GetResourceID() == "" && attributes.GetPath() == "" && len(attributes.GetLabels()) == 0 {
		return false // if designators are empty
	}

	return p.hasExceptionForAttributes(clusterName, attributes, workload, failingContainerNames, selector)
}

// hasExceptionForAttributes is hasException without the leading guard against an
// all-empty-fielded PortalDesignator. That guard exists to reject a populated
// Resources entry whose fields all happen to be empty (a likely misconfiguration);
// it has nothing to do with the Resources list itself being empty, which
// getResourceExceptions already handles as its own, deliberate "matches
// everywhere" case and calls this directly for.
func (p *Processor) hasExceptionForAttributes(clusterName string, attributes identifiers.AttributesDesignators, workload workloadinterface.IMetadata, failingContainerNames []string, selector labels.Selector) bool {
	if !p.matchesCluster(attributes, clusterName) {
		return false // cluster name does not match
	}

	if isTypeRegoResponseVector(workload) {
		if p.iterateRegoResponseVector(workload, attributes, failingContainerNames, selector) {
			return true
		}
		// If containerName is in the designator, stop here: the base
		// RegoResponseVector object is not a workload, so container membership
		// cannot be verified on it. Falling through would silently skip the
		// container check and produce false positives.
		if _, ok := attributes.GetLabels()[identifiers.AttributeContainerName]; ok {
			return false
		}
		// A non-empty objectSelector is a workload-label axis, and the base
		// RegoResponseVector envelope is not a real labeled workload. Evaluating the
		// selector against its empty label set is unsafe: negative operators
		// (DoesNotExist / NotIn) match an empty label set, so the fall-through would
		// wrongly apply an exception even though no related workload satisfied the
		// selector. Skip the base object entirely when a selector is present — the
		// selector was already ANDed with the designator per related object above.
		if selector != nil {
			return false
		}
		// otherwise (no selector), continue to check the base object against the designator.
	}
	return p.metadataHasException(workload, attributes, failingContainerNames, selector)

}

func (p *Processor) metadataHasException(workload workloadinterface.IMetadata, attributes identifiers.AttributesDesignators, failingContainerNames []string, selector labels.Selector) bool {

	if attributes.GetNamespace() != "" && !p.compareNamespace(workload, attributes.GetNamespace()) {
		return false // namespaces do not match
	}

	if attributes.GetKind() != "" && !p.compareKind(workload, attributes.GetKind()) {
		return false // kinds do not match
	}

	// apiGroup has no case in armoapi-go's DigestAttributesDesignator, so it arrives in the
	// labels map even though it is a schema field rather than a Kubernetes label. Match it
	// here, next to kind, and keep it out of the label comparison below: no resource carries
	// a literal apiGroup label, so leaving it there makes the whole designator unmatchable.
	//
	// Presence decides whether to compare, not emptiness. An explicit apiGroup: "" names the
	// core group, and is a constraint like any other: gating on a non-empty value would drop
	// it while still stripping the key, widening a core-only exception to every named group.
	if apiGroup, ok := attributes.GetLabels()[identifiers.AttributeApiGroup]; ok && !p.compareApiGroup(workload, apiGroup) {
		return false // api groups do not match
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

	// objectSelector (when present) is ANDed with the designator and evaluated against
	// this exact object's labels — so for a RegoResponseVector it is checked per related
	// object, the same object the designator above is checked against. Unlike the regex
	// label path below, the selector matches labels only (not annotations), per the
	// SecurityException spec ("selects workloads by their labels").
	if selector != nil {
		objLabels := labels.Set(workloadinterface.NewWorkloadObj(workload.GetObject()).GetLabels())
		if !selector.Matches(objLabels) {
			return false // objectSelector does not match this object's labels
		}
	}

	allLabels := attributes.GetLabels()
	containerName, hasContainerName := allLabels[identifiers.AttributeContainerName]

	// Build a label map with the attributes that are not Kubernetes labels stripped out,
	// so they are not compared as such.
	realLabels := stripNonLabelAttributes(allLabels)

	if !isTypeWorkload(workload) {
		// RegoResponseVector, host-sensor and cloud-provider objects have no labels,
		// annotations or containers to compare, so a designator that constrains any of
		// them cannot match. Skipping the comparison instead would widen a label-scoped
		// exception to every such object.
		return len(realLabels) == 0 && !hasContainerName
	}

	if len(realLabels) > 0 {
		if !p.compareLabels(workload, realLabels) && !p.compareAnnotations(workload, realLabels) {
			return false // labels nor annotations do not match
		}
	}

	if hasContainerName && !p.compareContainerName(workload, containerName, failingContainerNames) {
		return false // container name does not match
	}

	return true
}

// nonLabelAttributes are designator attribute keys that armoapi-go's
// DigestAttributesDesignator has no case for, so they are collected into the labels map
// even though they are schema fields rather than Kubernetes labels. Each one is matched
// explicitly, and has to be removed before the label and annotation comparison: no resource
// carries a literal containerName or apiGroup label, so comparing them as labels can only
// fail, and would make the whole designator unmatchable.
var nonLabelAttributes = []string{
	identifiers.AttributeContainerName,
	identifiers.AttributeApiGroup,
}

// stripNonLabelAttributes returns attributes without those keys, leaving the map untouched
// when it holds none of them.
func stripNonLabelAttributes(attributes map[string]string) map[string]string {
	held := slices.ContainsFunc(nonLabelAttributes, func(key string) bool {
		_, ok := attributes[key]
		return ok
	})
	if !held {
		return attributes
	}

	stripped := make(map[string]string, len(attributes))
	for key, value := range attributes {
		if !slices.Contains(nonLabelAttributes, key) {
			stripped[key] = value
		}
	}

	return stripped
}

func (p *Processor) iterateRegoResponseVector(workload workloadinterface.IMetadata, attributes identifiers.AttributesDesignators, failingContainerNames []string, selector labels.Selector) bool {
	v := objectsenvelopes.NewRegoResponseVectorObject(workload.GetObject())
	for _, r := range v.GetRelatedObjects() {
		if p.metadataHasException(r, attributes, failingContainerNames, selector) {
			return true
		}
	}
	return false
}

// extractFailingContainerNames parses paths like "spec.containers[0].…",
// "spec.template.spec.initContainers[1].…" or "spec.ephemeralContainers[0].…" to find which containers produced
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
	ephemeralContainers := ephemeralContainerNames(workload.GetObject())
	if len(containers)+len(initContainers)+len(ephemeralContainers) == 0 {
		return nil
	}

	seen := make(map[string]struct{})
	for _, path := range paths {
		for _, m := range rexContainerPath.FindAllStringSubmatch(path, -1) {
			idx, err := strconv.Atoi(m[2])
			if err != nil {
				continue
			}
			switch m[1] {
			case "initC":
				if idx < len(initContainers) {
					seen[initContainers[idx].Name] = struct{}{}
				}
			case "ephemeralC":
				if idx < len(ephemeralContainers) {
					seen[ephemeralContainers[idx]] = struct{}{}
				}
			default:
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
