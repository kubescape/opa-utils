package reporthandling

import (
	"encoding/json"
	"fmt"

	"github.com/kubescape/opa-utils/objectsenvelopes"
	"github.com/open-policy-agent/opa/rego"
)

// SetDefaultScore sets the framework,control default score
func SetDefaultScore(frameworkReport *FrameworkReport) {
	for c := range frameworkReport.ControlReports {
		frameworkReport.ControlReports[c].SetDefaultScore()
	}
	frameworkReport.SetDefaultScore()
}

// SetDefaultScore sets the framework,control,rule resource counter
func SetUniqueResourcesCounter(frameworkReport *FrameworkReport) {

	for c := range frameworkReport.ControlReports {
		for r := range frameworkReport.ControlReports[c].RuleReports {
			frameworkReport.ControlReports[c].RuleReports[r].SetResourcesCounters()
		}
		frameworkReport.ControlReports[c].SetResourcesCounters()
	}
	frameworkReport.SetResourcesCounters()
}

// GetUniqueResources the list of resources can contain duplications, this function removes the resource duplication based on workloadinterface.GetID
func GetUniqueResources(k8sResources []map[string]interface{}) []map[string]interface{} {
	uniqueRuleResponses := map[string]bool{}

	lenK8sResources := len(k8sResources)
	for i := 0; i < lenK8sResources; i++ {
		workload := objectsenvelopes.NewObject(k8sResources[i])
		if workload == nil { // remove none supported types
			k8sResources = removeFromSlice(k8sResources, i)
			lenK8sResources -= 1
			i -= 1
			continue
		}
		resourceID := workload.GetID()
		if found := uniqueRuleResponses[resourceID]; found {
			// resource found -> remove from slice
			k8sResources = removeFromSlice(k8sResources, i)
			lenK8sResources -= 1
			i -= 1
		} else {
			uniqueRuleResponses[resourceID] = true
		}
	}
	return k8sResources
}

// GetUniqueResources the list of resources can contain duplications, this function removes the resource duplication based on workloadinterface.GetID
func GetUniqueResourcesIDs(k8sResourcesList []string) []string {
	uniqueRuleResponses := map[string]bool{}
	k8sResourcesNewList := []string{}

	for i := range k8sResourcesList {
		if found := uniqueRuleResponses[k8sResourcesList[i]]; !found {
			uniqueRuleResponses[k8sResourcesList[i]] = true
			k8sResourcesNewList = append(k8sResourcesNewList, k8sResourcesList[i])
		}
	}
	return k8sResourcesNewList
}

// TrimUniqueResources trim the list, this wil trim in case the same resource appears in the warning list and in the failed list
func TrimUniqueIDs(origin, trimFrom []string) []string {
	if len(origin) == 0 || len(trimFrom) == 0 { // if there is nothing to trim
		return origin
	}
	uniqueResources := map[string]bool{}
	listResources := []string{}

	for i := range trimFrom {
		uniqueResources[trimFrom[i]] = true
	}

	for i := range origin {
		if found := uniqueResources[origin[i]]; !found {
			listResources = append(listResources, origin[i])
		}
	}
	return listResources
}

func removeFromSlice(k8sResources []map[string]interface{}, i int) []map[string]interface{} {
	if i != len(k8sResources)-1 {
		k8sResources[i] = k8sResources[len(k8sResources)-1]
	}

	return k8sResources[:len(k8sResources)-1]
}

func ParseRegoResult(regoResult *rego.ResultSet) ([]RuleResponse, error) {
	var errs error
	ruleResponses := []RuleResponse{}
	for _, result := range *regoResult {
		for desicionIdx := range result.Expressions {
			if resMap, ok := result.Expressions[desicionIdx].Value.(map[string]interface{}); ok {
				for objName := range resMap {
					jsonBytes, err := json.Marshal(resMap[objName])
					if err != nil {
						err = fmt.Errorf("in parseRegoResult, json.Marshal failed. name: %s, obj: %v, reason: %s", objName, resMap[objName], err)
						errs = fmt.Errorf("%s\n%s", errs, err)
						continue
					}
					desObj := make([]RuleResponse, 0)
					if err := json.Unmarshal(jsonBytes, &desObj); err != nil {
						err = fmt.Errorf("in parseRegoResult, json.Unmarshal failed. name: %s, obj: %v, reason: %s", objName, resMap[objName], err)
						errs = fmt.Errorf("%s\n%s", errs, err)
						continue
					}
					ruleResponses = append(ruleResponses, desObj...)
				}
			}
		}
	}
	return ruleResponses, errs
}
