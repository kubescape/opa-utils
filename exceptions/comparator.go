package exceptions

import (
	"regexp"
	"strings"
	"sync"

	"github.com/kubescape/k8s-interface/k8sinterface"
	"github.com/kubescape/k8s-interface/workloadinterface"

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

func (c *comparator) compareName(workload workloadinterface.IMetadata, name string) bool {
	return c.regexCompare(name, workload.GetName())
}

func (c *comparator) comparePath(workload workloadinterface.IMetadata, path string) bool {
	w := workload.GetObject()
	if !k8sinterface.IsTypeWorkload(w) {
		return false
	}

	if val, ok := w["sourcePath"]; ok {
		if sourcePath, ok := val.(string); ok {
			return c.regexCompare(path, sourcePath)
		}
	}

	return false
}

func (c *comparator) compareLabels(workload workloadinterface.IMetadata, attributes map[string]string) bool {
	w := workload.GetObject()
	if !k8sinterface.IsTypeWorkload(w) {
		return true
	}

	workloadLabels := labels.Set(workloadinterface.NewWorkloadObj(w).GetLabels())

	if len(workloadLabels) == 0 {
		return false
	}

	for key, val := range attributes {
		for _, annotation := range workloadLabels {
			if !workloadLabels.Has(key) {
				return false
			}
			if !c.regexCompare(val, annotation) {
				return false
			}
		}
	}

	return true // ignore labels
}

func (c *comparator) compareAnnotations(workload workloadinterface.IMetadata, attributes map[string]string) bool {
	w := workload.GetObject()
	if !k8sinterface.IsTypeWorkload(w) {
		return true
	}

	workloadAnnotations := labels.Set(workloadinterface.NewWorkloadObj(w).GetAnnotations())
	if len(workloadAnnotations) == 0 {
		return false
	}

	for key, val := range attributes {
		for _, annotation := range workloadAnnotations {
			if !workloadAnnotations.Has(key) {
				return false
			}
			if !c.regexCompare(val, annotation) {
				return false
			}
		}
	}

	return true // ignore annotations
}

func (c *comparator) compareCluster(designatorCluster, clusterName string) bool {
	return designatorCluster != "" && c.regexCompare(designatorCluster, clusterName)
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

	var rexBuilder strings.Builder // builds the regexp with minimal alloc
	rexBuilder.Grow(len(reg) + 6)
	_, _ = rexBuilder.Write([]byte("(?i)")) // builds a case-insensitive regexp: more efficient than calling ToLower on both operands
	_, _ = rexBuilder.Write([]byte("^"))
	_, _ = rexBuilder.WriteString(reg)
	_, _ = rexBuilder.Write([]byte("$"))

	r, err := regexp.Compile(rexBuilder.String())
	if err != nil {
		return false
	}

	c.irexCache.Store(reg, r) // keep the compiled regexp in cache

	return r.MatchString(name)
}

// regexpCompareI performs a case-sensitive regexp match
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

	var rexBuilder strings.Builder // builds the regexp with minimal alloc
	rexBuilder.Grow(len(reg) + 2)

	_, _ = rexBuilder.Write([]byte("^"))
	_, _ = rexBuilder.WriteString(reg)
	_, _ = rexBuilder.Write([]byte("$"))

	r, err := regexp.Compile(rexBuilder.String())
	if err != nil {
		return false
	}

	c.rexCache.Store(reg, r) // keep the compiled regexp in cache

	return r.MatchString(name)
}
