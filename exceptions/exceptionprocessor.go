package exceptions

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/armosec/k8s-interface/workloadinterface"
	"github.com/armosec/opa-utils/objectsenvelopes"
	"github.com/armosec/opa-utils/reporthandling"

	"github.com/armosec/armoapi-go/armotypes"
	"k8s.io/apimachinery/pkg/labels"
)

// SetFrameworkExceptions add exceptions to framework report
func SetFrameworkExceptions(frameworkReport *reporthandling.FrameworkReport, exceptionsPolicies []armotypes.PostureExceptionPolicy, clusterName string) {
	for c := range frameworkReport.ControlReports {
		SetControlExceptions(&frameworkReport.ControlReports[c], exceptionsPolicies, clusterName, frameworkReport.Name)
	}
}

// SetControlExceptions add exceptions to control report
func SetControlExceptions(controlReport *reporthandling.ControlReport, exceptionsPolicies []armotypes.PostureExceptionPolicy, clusterName, frameworkName string) {
	for r := range controlReport.RuleReports {
		SetRuleExceptions(&controlReport.RuleReports[r], exceptionsPolicies, clusterName, frameworkName, controlReport.Name, controlReport.ControlID)
	}
}

// SetRuleExceptions add exceptions to rule report
func SetRuleExceptions(ruleReport *reporthandling.RuleReport, exceptionsPolicies []armotypes.PostureExceptionPolicy, clusterName, frameworkName, controlName, controlID string) {

	// adding exceptions to the rules
	ruleExceptions := ListRuleExceptions(exceptionsPolicies, frameworkName, controlName, controlID, ruleReport.Name)
	SetRuleResponsExceptions(ruleReport.RuleResponses, ruleExceptions, clusterName)
}

// SetRuleExceptions add exceptions to rule respons structure
func SetRuleResponsExceptions(results []reporthandling.RuleResponse, ruleExceptions []armotypes.PostureExceptionPolicy, clusterName string) {
	if len(ruleExceptions) == 0 {
		return
	}
	for i := range results {
		workloads := alertObjectToWorkloads(&results[i].AlertObject)
		if len(workloads) == 0 {
			continue
		}
		for w := range workloads {
			if exception := getException(ruleExceptions, workloads[w], clusterName); exception != nil {
				results[i].Exception = exception
			}
		}
		results[i].RuleStatus = results[i].GetStatus()
	}
}
func ListRuleExceptions(exceptionPolicies []armotypes.PostureExceptionPolicy, frameworkName, controlName, controlID, ruleName string) []armotypes.PostureExceptionPolicy {
	ruleExceptions := []armotypes.PostureExceptionPolicy{}
	for i := range exceptionPolicies {
		if ruleHasExceptions(&exceptionPolicies[i], frameworkName, controlName, controlID, ruleName) {
			ruleExceptions = append(ruleExceptions, exceptionPolicies[i])
		}
	}

	return ruleExceptions

}

func ruleHasExceptions(exceptionPolicy *armotypes.PostureExceptionPolicy, frameworkName, controlName, controlID, ruleName string) bool {
	for _, posturePolicy := range exceptionPolicy.PosturePolicies {
		if posturePolicy.FrameworkName == "" && posturePolicy.ControlName == "" && posturePolicy.ControlID == "" && posturePolicy.RuleName == "" {
			continue // empty policy -> ignore
		}
		if posturePolicy.FrameworkName != "" && !strings.EqualFold(posturePolicy.FrameworkName, frameworkName) {
			continue // policy does not match
		}
		if posturePolicy.ControlName != "" && !strings.EqualFold(posturePolicy.ControlName, controlName) {
			continue // policy does not match
		}
		if posturePolicy.ControlID != "" && !strings.EqualFold(posturePolicy.ControlID, controlID) {
			continue // policy does not match
		}
		if posturePolicy.RuleName != "" && !strings.EqualFold(posturePolicy.RuleName, ruleName) {
			continue // policy does not match
		}
		return true // policies match
	}

	return false

}

func alertObjectToWorkloads(obj *reporthandling.AlertObject) []workloadinterface.IMetadata {
	resource := []workloadinterface.IMetadata{}

	for i := range obj.K8SApiObjects {
		r := objectsenvelopes.NewObject(obj.K8SApiObjects[i])
		if r == nil {
			continue
		}
		resource = append(resource, r)
		ns := r.GetNamespace()
		if ns != "" {
			// TODO - handle empty namespace
		}
	}

	if obj.ExternalObjects != nil {
		if r := objectsenvelopes.NewObject(obj.ExternalObjects); r != nil {
			// TODO - What about linked objects?
			resource = append(resource, r)
		}
	}

	return resource
}
func getException(ruleExceptions []armotypes.PostureExceptionPolicy, workload workloadinterface.IMetadata, clusterName string) *armotypes.PostureExceptionPolicy {
	for e := range ruleExceptions {
		for _, resource := range ruleExceptions[e].Resources {
			if hasException(clusterName, &resource, workload) {
				return &ruleExceptions[e] // TODO - return disable exception out of all exceptions
			}
		}
	}
	return nil
}

// compareMetadata - compare namespace and kind
func hasException(clusterName string, designator *armotypes.PortalDesignator, workload workloadinterface.IMetadata) bool {
	cluster, namespace, kind, name, labels := designator.DigestPortalDesignator()

	if cluster == "" && namespace == "" && kind == "" && name == "" && len(labels) == 0 {
		return false // if designators are empty
	}

	if cluster != "" && !compareCluster(cluster, clusterName) { // TODO - where do we receive cluster name from?
		return false // cluster name does not match
	}

	if namespace != "" && !compareNamespace(workload, namespace) {
		return false // namespaces do not match
	}

	if kind != "" && !compareKind(workload, kind) {
		return false // kinds do not match
	}

	if name != "" && !compareName(workload, name) {
		return false // names do not match
	}
	if len(labels) > 0 && !compareLabels(workload, labels) {
		return false // labels do not match
	}

	return true // no mismatch found -> the workload has an exception
}

func compareNamespace(workload workloadinterface.IMetadata, namespace string) bool {
	if workload.GetKind() == "Namespace" {
		return regexCompare(namespace, workload.GetName())
	}
	return regexCompare(namespace, workload.GetNamespace())
}

func compareKind(workload workloadinterface.IMetadata, kind string) bool {
	return regexCompare(kind, workload.GetKind())
}

func compareName(workload workloadinterface.IMetadata, name string) bool {
	return regexCompare(workload.GetName(), name)
}

func compareLabels(workload workloadinterface.IMetadata, attributes map[string]string) bool {
	w := workload.GetObject()
	if workloadinterface.IsTypeWorkload(w) {
		workloadLabels := labels.Set(workloadinterface.NewWorkloadObj(w).GetLabels())
		designators := labels.Set(attributes).AsSelector()

		return designators.Matches(workloadLabels)
	}
	return true // ignore labels
}

func compareCluster(designatorCluster, clusterName string) bool {
	return designatorCluster != "" && regexCompare(designatorCluster, clusterName)
}

func regexCompare(reg, name string) bool {
	r, _ := regexp.MatchString(fmt.Sprintf("^%s$", reg), name)
	// if err != nil {
	// 	return false
	// }
	return r
}
