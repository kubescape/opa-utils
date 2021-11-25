package reporthandling

import (
	"bytes"
	"encoding/gob"
	"strings"

	"github.com/armosec/k8s-interface/workloadinterface"
)

var aggregatorAttribute = "resourcesAggregator"

func RegoResourcesAggregator(rule *PolicyRule, k8sObjects []workloadinterface.IMetadata) ([]workloadinterface.IMetadata, error) {
	if aggregateBy, ok := rule.Attributes[aggregatorAttribute]; ok {
		switch aggregateBy {
		case "subject-role-rolebinding":
			return AggregateResourcesBySubjects(k8sObjects)
		case "apiserver-pod":
			return AggregateResourcesByAPIServerPod(k8sObjects)
		default:
			return k8sObjects, nil
		}
	}
	return k8sObjects, nil
}

func AggregateResourcesBySubjects(k8sObjects []workloadinterface.IMetadata) ([]workloadinterface.IMetadata, error) {
	aggregatedK8sObjects := []workloadinterface.IMetadata{}
	for _, bindingWorkload := range k8sObjects {
		if strings.HasSuffix(bindingWorkload.GetKind(), "Binding") { // types.Role
			for _, roleWorkload := range k8sObjects {
				if strings.HasSuffix(roleWorkload.GetKind(), "Role") {
					bindingWorkloadObj := bindingWorkload.GetObject()
					if kind, ok := workloadinterface.InspectMap(bindingWorkloadObj, "roleRef", "kind"); ok {
						if name, ok := workloadinterface.InspectMap(bindingWorkloadObj, "roleRef", "name"); ok {
							if kind.(string) == roleWorkload.GetKind() && name.(string) == roleWorkload.GetName() {
								if subjects, ok := workloadinterface.InspectMap(bindingWorkloadObj, "subjects"); ok {
									if data, ok := subjects.([]interface{}); ok {
										for _, subject := range data {
											subjectAllFields, err := setSubjectFields(subject.(map[string]interface{}))
											if err != nil {
												return aggregatedK8sObjects, err
											}
											subjectAllFields[workloadinterface.RelatedObjectsKey] = []map[string]interface{}{bindingWorkload.GetObject(), roleWorkload.GetObject()}
											newObj := workloadinterface.NewRegoResponseVectorObject(subjectAllFields)
											aggregatedK8sObjects = append(aggregatedK8sObjects, newObj)
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
	return aggregatedK8sObjects, nil
}

// Create custom object of apiserver pod. Has required fields + cmdline
func AggregateResourcesByAPIServerPod(k8sObjects []workloadinterface.IMetadata) ([]workloadinterface.IMetadata, error) {
	return k8sObjects, nil
	// apiServerPod := []workloadinterface.IMetadata{}
	// for _, obj := range k8sObjects {
	// 	if !workloadinterface.IsTypeWorkload(obj.GetObject()) {
	// 		continue
	// 	}
	// 	workload := workloadinterface.NewWorkloadObj(obj.GetObject())
	// 	if workload.GetKind() == "Pod" && workload.GetNamespace() == "kube-system" {
	// 		if strings.Contains(workload.GetName(), "apiserver") || strings.Contains(workload.GetName(), "api-server") {
	//
	/*

		TODO
		====
		Create a new supported IMetadata object

	*/
	// 			apiServerPod["namespace"] = workload.GetNamespace()
	// 			apiServerPod["name"] = workload.GetName()
	// 			apiServerPod["kind"] = workload.GetKind()
	// 			apiServerPod["apiVersion"] = workload.GetApiVersion()
	// 			containers, err := workload.GetContainers()
	// 			if err != nil || len(containers) == 0 {
	// 				return nil, err
	// 			}
	// 			// apiServer has only one container
	// 			apiServerPod["cmdline"] = containers[0].Command
	// 			return []map[string]interface{}{apiServerPod}, nil
	// 		}
	// 	}
	// }
	// return nil, nil
}

func setSubjectFields(subject map[string]interface{}) (map[string]interface{}, error) {
	newSubject, err := DeepCopyMap(subject)
	if err != nil {
		return nil, err
	}
	if _, ok := workloadinterface.InspectMap(newSubject, "name"); !ok {
		newSubject["name"] = ""
	}
	if _, ok := workloadinterface.InspectMap(newSubject, "namespace"); !ok {
		newSubject["namespace"] = ""
	}
	if _, ok := workloadinterface.InspectMap(newSubject, "kind"); !ok {
		newSubject["kind"] = ""
	}
	if _, ok := workloadinterface.InspectMap(newSubject, "apiVersion"); !ok {
		newSubject["apiVersion"] = ""
	}
	return newSubject, nil
}

// DeepCopyMap performs a deep copy of the given map m.
func DeepCopyMap(m map[string]interface{}) (map[string]interface{}, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	dec := gob.NewDecoder(&buf)
	err := enc.Encode(m)
	if err != nil {
		return nil, err
	}
	var copy map[string]interface{}
	err = dec.Decode(&copy)
	if err != nil {
		return nil, err
	}
	return copy, nil
}
