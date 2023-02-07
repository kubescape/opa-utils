package reporthandling

import (
	"encoding/json"
	"fmt"

	"github.com/kubescape/opa-utils/objectsenvelopes"
	"github.com/kubescape/opa-utils/reporthandling/internal/slices"
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

// GetUniqueResourcesIDs yields the list of unique resource IDs. Duplicates are removed, based on the workload.GetID() interface method.
//
// NOTE: the input slice is modified in-place.
func GetUniqueResourcesIDs(k8sResourcesList []string) []string {
	return slices.UniqueStrings(k8sResourcesList)
}

// TrimUniqueResources trims the origin list to contain only elements that are NOT already present in the trimFrom list.
//
// # This is used to cover the case when the same resource appears in the warning list and in the failed list
//
// NOTE: the origin slice is modified in-place.
func TrimUniqueIDs(origin, trimFrom []string) []string {
	return slices.TrimStable(origin, trimFrom)
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
