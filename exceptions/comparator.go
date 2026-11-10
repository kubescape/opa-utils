package exceptions

import (
	"regexp"
	"strings"
	"sync"

	"github.com/kubescape/k8s-interface/k8sinterface"
	"github.com/kubescape/k8s-interface/workloadinterface"
	"github.com/kubescape/opa-utils/objectsenvelopes"
	"github.com/kubescape/opa-utils/objectsenvelopes/localworkload"

	"k8s.io/apimachinery/pkg/labels"
)

// newComparator builds a comparator for exceptions, with an inner cache for compiled regexps.
func newComparator() *comparator {
	return &comparator{}
}

// comparator exposes comparison methods, with a cache to reuse previously compiled regexps.
//
// NOTE(fredbi): is there a way to simplify here and adopt only case-insensitive regexps?
type comparator struct {
	rexCache  sync.Map // for case sensitive regexp
	irexCache sync.Map // for case insensitive regexps
}

func (c *comparator) compareNamespace(workload workloadinterface.IMetadata, namespace string) bool {
	if workload.GetKind() == "Namespace" {
		return c.regexCompare(namespace, workload.GetName())
	}

	return c.regexCompare(namespace, workload.GetNamespace())
}

func (c *comparator) compareKind(workload workloadinterface.IMetadata, kind string) bool {
	return c.regexCompare(kind, workload.GetKind())
}

// compareApiGroup matches against the resource's API group. A core-group resource reports an
// empty group, so it is matched only by a designator whose pattern accepts the empty string.
func (c *comparator) compareApiGroup(workload workloadinterface.IMetadata, apiGroup string) bool {
	return c.regexCompare(apiGroup, apiGroupOf(workload))
}

// apiGroupOf returns the resource's API group.
//
// Most objects carry a group/version apiVersion, and the group is its first half. A
// RegoResponseVector built from an RBAC subject instead carries a bare apiGroup such as
// "rbac.authorization.k8s.io" and no version at all, which GetApiVersion returns as-is;
// SplitApiVersion reads a value with no "/" as a version and reports no group, which would
// stop a correctly scoped subject exception from matching. The object's own apiGroup field
// is authoritative where it has one, so it is preferred.
func apiGroupOf(workload workloadinterface.IMetadata) string {
	if value, ok := workloadinterface.InspectMap(workload.GetObject(), "apiGroup"); ok {
		if group, isString := value.(string); isString {
			return group
		}
	}

	group, _ := k8sinterface.SplitApiVersion(workload.GetApiVersion())

	return group
}

func (c *comparator) compareName(workload workloadinterface.IMetadata, name string) bool {
	return c.regexCompare(name, workload.GetName())
}

func (c *comparator) compareResourceID(workload workloadinterface.IMetadata, resourceID string) bool {
	return c.regexCompare(resourceID, workload.GetID())
}

func (c *comparator) comparePath(workload workloadinterface.IMetadata, path string) bool {
	w := workload.GetObject()

	if val, ok := w["sourcePath"]; ok {
		if sourcePath, ok := val.(string); ok {
			return c.regexCompare(path, sourcePath)
		}
	}

	return false
}

func (c *comparator) compareLabels(workload workloadinterface.IMetadata, attributes map[string]string) bool {
	w := workload.GetObject()

	workloadLabels := labels.Set(workloadinterface.NewWorkloadObj(w).GetLabels())

	if len(workloadLabels) == 0 {
		return false
	}

	for key, val := range attributes {
		if !workloadLabels.Has(key) {
			return false
		}
		value := workloadLabels.Get(key)
		if !c.regexCompare(val, value) {
			return false
		}
	}

	return true // ignore labels
}

func (c *comparator) compareAnnotations(workload workloadinterface.IMetadata, attributes map[string]string) bool {

	w := workload.GetObject()
	workloadAnnotations := labels.Set(workloadinterface.NewWorkloadObj(w).GetAnnotations())
	if len(workloadAnnotations) == 0 {
		return false
	}

	for key, val := range attributes {
		if !workloadAnnotations.Has(key) {
			return false
		}
		value := workloadAnnotations.Get(key)
		if !c.regexCompare(val, value) {
			return false
		}
	}

	return true // ignore annotations
}

func (c *comparator) compareCluster(designatorCluster, clusterName string) bool {
	return designatorCluster != "" && c.regexCompare(designatorCluster, clusterName)
}

// compareContainerName reports whether containerName matches any of the
// failing containers. If failingNames is provided, only those containers are
// checked; otherwise every container, init container and ephemeral container
// in the workload is inspected (used by callers that have no FailedPaths context).
func (c *comparator) compareContainerName(workload workloadinterface.IMetadata, containerName string, failingNames []string) bool {
	if len(failingNames) > 0 {
		for _, name := range failingNames {
			if c.regexCompare(containerName, name) {
				return true
			}
		}
		return false
	}

	wl := workloadinterface.NewWorkloadObj(workload.GetObject())

	if containers, err := wl.GetContainers(); err == nil {
		for i := range containers {
			if c.regexCompare(containerName, containers[i].Name) {
				return true
			}
		}
	}

	if initContainers, err := wl.GetInitContainers(); err == nil {
		for i := range initContainers {
			if c.regexCompare(containerName, initContainers[i].Name) {
				return true
			}
		}
	}

	for _, name := range ephemeralContainerNames(workload.GetObject()) {
		if c.regexCompare(containerName, name) {
			return true
		}
	}

	return false
}

// ephemeralContainerNames returns the names of a Pod's ephemeral containers, in
// spec order. The workload interface has no getter for them, so they are read
// from the object.
func ephemeralContainerNames(obj map[string]interface{}) []string {
	raw, ok := workloadinterface.InspectMap(obj, "spec", "ephemeralContainers")
	if !ok {
		return nil
	}
	list, ok := raw.([]interface{})
	if !ok {
		return nil
	}

	names := make([]string, 0, len(list))
	for _, item := range list {
		container, _ := item.(map[string]interface{})
		name, _ := container["name"].(string)
		names = append(names, name)
	}

	return names
}

// regexpCompareI performs a case-insensitive regexp match
func (c *comparator) regexCompareI(reg, name string) bool {
	var (
		val interface{}
		ok  bool
	)

	val, ok = c.irexCache.Load(reg)
	if ok {
		// we've already compiled this regexp
		rex := val.(*regexp.Regexp)

		return rex.MatchString(name)
	}

	// builds a case-insensitive regexp: more efficient than calling ToLower on both operands
	r, err := compileAnchored("(?i)", reg)
	if err != nil {
		return false
	}

	c.irexCache.Store(reg, r) // keep the compiled regexp in cache

	return r.MatchString(name)
}

// regexCompare performs a case-sensitive regexp match
func (c *comparator) regexCompare(reg, name string) bool {
	var (
		val interface{}
		ok  bool
	)

	val, ok = c.rexCache.Load(reg)
	if ok {
		// we've already compiled this regexp
		rex := val.(*regexp.Regexp)

		return rex.MatchString(name)
	}

	r, err := compileAnchored("", reg)
	if err != nil {
		return false
	}

	c.rexCache.Store(reg, r) // keep the compiled regexp in cache

	return r.MatchString(name)
}

// compileAnchored compiles reg so that it has to match the whole value.
//
// The pattern is grouped before anchoring: "^a|b$" would match any value that
// starts with a or ends with b, so an exception written for "kube-system|kube-public"
// would also apply to "kube-system-tenant".
//
// A pattern that only compiles without the group, such as the "*" wildcard
// (which "^*$" accepts for any value), keeps its previous meaning.
func compileAnchored(flags, reg string) (*regexp.Regexp, error) {
	var rexBuilder strings.Builder // builds the regexp with minimal alloc
	rexBuilder.Grow(len(flags) + len(reg) + 6)
	_, _ = rexBuilder.WriteString(flags)
	_, _ = rexBuilder.WriteString("^(?:")
	_, _ = rexBuilder.WriteString(reg)
	_, _ = rexBuilder.WriteString(")$")

	if r, err := regexp.Compile(rexBuilder.String()); err == nil {
		return r, nil
	}

	return regexp.Compile(flags + "^" + reg + "$")
}

func isTypeWorkload(workload workloadinterface.IMetadata) bool {
	switch objectsenvelopes.GetObjectType(workload.GetObject()) {
	case workloadinterface.TypeBaseObject:
		return true
	case workloadinterface.TypeWorkloadObject:
		return true
	case workloadinterface.TypeListWorkloads:
		return true
	case localworkload.TypeLocalWorkload:
		return true
	default:
		return false
	}
}
func isTypeRegoResponseVector(workload workloadinterface.IMetadata) bool {
	switch objectsenvelopes.GetObjectType(workload.GetObject()) {
	case objectsenvelopes.TypeRegoResponseVectorObject:
		return true
	default:
		return false
	}
}
