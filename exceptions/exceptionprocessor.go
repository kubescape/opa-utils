package exceptions

import (
	"fmt"
	"regexp"

	"github.com/armosec/k8s-interface/k8sinterface"
	"github.com/armosec/k8s-interface/workloadinterface"
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
		SetRuleExceptions(&controlReport.RuleReports[r], exceptionsPolicies, clusterName, frameworkName, controlReport.Name)
	}
}

// SetRuleExceptions add exceptions to rule report
func SetRuleExceptions(ruleReport *reporthandling.RuleReport, exceptionsPolicies []armotypes.PostureExceptionPolicy, clusterName, frameworkName, controlName string) {

	// adding exceptions to the rules
	ruleExceptions := ListRuleExceptions(exceptionsPolicies, frameworkName, controlName, ruleReport.Name)
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
func ListRuleExceptions(exceptionPolicies []armotypes.PostureExceptionPolicy, frameworkName, controlName, ruleName string) []armotypes.PostureExceptionPolicy {
	ruleExceptions := []armotypes.PostureExceptionPolicy{}
	for i := range exceptionPolicies {
		if ruleHasExceptions(&exceptionPolicies[i], frameworkName, controlName, ruleName) {
			ruleExceptions = append(ruleExceptions, exceptionPolicies[i])
		}
	}

	return ruleExceptions

}

func ruleHasExceptions(exceptionPolicy *armotypes.PostureExceptionPolicy, frameworkName, controlName, ruleName string) bool {
	for _, posturePolicy := range exceptionPolicy.PosturePolicies {
		if posturePolicy.FrameworkName == "" && posturePolicy.ControlName == "" && posturePolicy.RuleName == "" {
			continue // empty policy -> ignore
		}
		if posturePolicy.FrameworkName != "" && posturePolicy.FrameworkName != frameworkName {
			continue // policy does not match
		}
		if posturePolicy.ControlName != "" && posturePolicy.ControlName != controlName {
			continue // policy does not match
		}
		if posturePolicy.RuleName != "" && posturePolicy.RuleName != ruleName {
			continue // policy does not match
		}
		return true // policies match
	}

	return false

}

func alertObjectToWorkloads(obj *reporthandling.AlertObject) []k8sinterface.IWorkload {
	resource := []k8sinterface.IWorkload{}

	for i := range obj.K8SApiObjects {
		r := workloadinterface.NewWorkloadObj(obj.K8SApiObjects[i])
		if r == nil {
			continue
		}
		resource = append(resource, r)
		ns := r.GetNamespace()
		if ns != "" {

		}
	}

	return resource
}
func getException(ruleExceptions []armotypes.PostureExceptionPolicy, workload k8sinterface.IWorkload, clusterName string) *armotypes.PostureExceptionPolicy {
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
func hasException(clusterName string, designator *armotypes.PortalDesignator, workload k8sinterface.IWorkload) bool {
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

func compareNamespace(workload k8sinterface.IWorkload, namespace string) bool {
	if workload.GetKind() == "Namespace" {
		return regexCompare(namespace, workload.GetName())
	}
	return regexCompare(namespace, workload.GetNamespace())
}

func compareKind(workload k8sinterface.IWorkload, kind string) bool {
	return regexCompare(kind, workload.GetKind())
}

func compareName(workload k8sinterface.IWorkload, name string) bool {
	return regexCompare(workload.GetName(), name)
}

func compareLabels(workload k8sinterface.IWorkload, attributes map[string]string) bool {
	workloadLabels := labels.Set(workload.GetLabels())
	designators := labels.Set(attributes).AsSelector()

	return designators.Matches(workloadLabels)
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
