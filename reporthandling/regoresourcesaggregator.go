package reporthandling

import (
	"strings"

	"github.com/armosec/k8s-interface/workloadinterface"
)

var aggregatorAttribute = "resourcesAggregator"

func RegoResourcesAggregator(rule *PolicyRule, k8sObjects []map[string]interface{}) []map[string]interface{} {
	if aggregateBy, ok := rule.Attributes[aggregatorAttribute]; ok {
		switch aggregateBy {
		case "subject-role-rolebinding":
			return AggregateResourcesBySubjects(k8sObjects)
		case "apiserver-pod":
			return AggregateResourcesByAPIServerPod(k8sObjects)
		default:
			return k8sObjects
		}
	}
	return k8sObjects
}

func AggregateResourcesBySubjects(k8sObjects []map[string]interface{}) []map[string]interface{} {
	var aggregatedK8sObjects []map[string]interface{}
	for _, firstk8sObject := range k8sObjects {
		bindingWorkload := workloadinterface.NewWorkloadObj(firstk8sObject)
		if strings.HasSuffix(bindingWorkload.GetKind(), "Binding") { // types.Role
			for _, secondK8sObject := range k8sObjects {
				roleWorkload := workloadinterface.NewWorkloadObj(secondK8sObject)
				if strings.HasSuffix(roleWorkload.GetKind(), "Role") {
					bindingWorkloadObj := bindingWorkload.GetObject()
					if kind, ok := workloadinterface.InspectMap(bindingWorkloadObj, "roleRef", "kind"); ok {
						if name, ok := workloadinterface.InspectMap(bindingWorkloadObj, "roleRef", "name"); ok {
							if kind.(string) == roleWorkload.GetKind() && name.(string) == roleWorkload.GetName() {
								if subjects, ok := workloadinterface.InspectMap(bindingWorkloadObj, "subjects"); ok {
									if data, ok := subjects.([]interface{}); ok {
										for _, subject := range data {
											subjectAllFields := setSubjectFields(subject.(map[string]interface{}))
											subjectAllFields[workloadinterface.RelatedObjectsKey] = []map[string]interface{}{bindingWorkload.GetObject(), roleWorkload.GetObject()}
											newObj := workloadinterface.NewRegoResponseVectorObject(subjectAllFields)
											aggregatedK8sObjects = append(aggregatedK8sObjects, newObj.GetObject())
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return aggregatedK8sObjects
}

// Create custom object of apiserver pod. Has required fields + cmdline
func AggregateResourcesByAPIServerPod(k8sObjects []map[string]interface{}) []map[string]interface{} {
	apiServerPod := map[string]interface{}{}
	for _, obj := range k8sObjects {
		workload := workloadinterface.NewWorkloadObj(obj)
		if workload.GetKind() == "Pod" && workload.GetNamespace() == "kube-system" {
			if strings.Contains(workload.GetName(), "apiserver") || strings.Contains(workload.GetName(), "api-server") {
				apiServerPod["namespace"] = workload.GetNamespace()
				apiServerPod["name"] = workload.GetName()
				apiServerPod["kind"] = workload.GetKind()
				apiServerPod["apiVersion"] = workload.GetApiVersion()
				containers, err := workload.GetContainers()
				if err != nil || len(containers) == 0 {
					return nil
				}
				// apiServer has only one container
				apiServerPod["cmdline"] = containers[0].Command
				return []map[string]interface{}{apiServerPod}
			}
		}
	}
	return nil
}

func setSubjectFields(subject map[string]interface{}) map[string]interface{} {

	if _, ok := workloadinterface.InspectMap(subject, "name"); !ok {
		subject["name"] = ""
	}
	if _, ok := workloadinterface.InspectMap(subject, "namespace"); !ok {
		subject["namespace"] = ""
	}
	if _, ok := workloadinterface.InspectMap(subject, "kind"); !ok {
		subject["kind"] = ""
	}
	if _, ok := workloadinterface.InspectMap(subject, "apiVersion"); !ok {
		subject["apiVersion"] = ""
	}
	return subject
}
