package reporthandling

import (
	"encoding/json"
	"fmt"

	"github.com/armosec/k8s-interface/workloadinterface"

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
	uniqueAllFramework := []string{}
	uniqueWarningFramework := []string{}
	uniqueFailedFramework := []string{}
	for c := range frameworkReport.ControlReports {
		uniqueAllControls, uniqueWarningControls, uniqueFailedControls := GetIDsPerControl(&frameworkReport.ControlReports[c])

		// Set
		frameworkReport.ControlReports[c].SetNumberOfResources(len(uniqueAllControls))
		frameworkReport.ControlReports[c].SetNumberOfWarningResources(len(uniqueWarningControls))
		frameworkReport.ControlReports[c].SetNumberOfFailedResources(len(uniqueFailedControls))

		// Append
		uniqueAllFramework = append(uniqueAllFramework, uniqueAllControls...)
		uniqueWarningFramework = append(uniqueWarningFramework, uniqueWarningControls...)
		uniqueFailedFramework = append(uniqueFailedFramework, uniqueFailedControls...)
	}

	// Get
	uniqueAllFramework = GetUniqueResourcesIDs(uniqueAllFramework)
	uniqueWarningFramework = GetUniqueResourcesIDs(uniqueWarningFramework)
	uniqueFailedFramework = GetUniqueResourcesIDs(uniqueFailedFramework)
	uniqueWarningFramework = TrimUniqueIDs(uniqueWarningFramework, uniqueFailedFramework)

	// Set
	frameworkReport.SetNumberOfResources(len(uniqueAllFramework))
	frameworkReport.SetNumberOfWarningResources(len(uniqueWarningFramework))
	frameworkReport.SetNumberOfFailedResources(len(uniqueFailedFramework))
}

// GetResourcesPerControl - return unique lists of resource IDs: all,warning,failed
func GetIDsPerControl(ctrlReport *ControlReport) ([]string, []string, []string) {
	uniqueAllResources := []string{}
	uniqueWarningResources := []string{}
	uniqueFailedResources := []string{}
	for r := range ctrlReport.RuleReports {

		uniqueAll := ctrlReport.RuleReports[r].GetAllResourcesIDs()
		uniqueFailed := GetUniqueResourcesIDs(workloadinterface.ListMetaIDs(workloadinterface.ListMapToMeta(ctrlReport.RuleReports[r].GetFailedResources())))
		uniqueWarning := GetUniqueResourcesIDs(workloadinterface.ListMetaIDs(workloadinterface.ListMapToMeta(ctrlReport.RuleReports[r].GetWarnignResources())))
		uniqueWarning = TrimUniqueIDs(uniqueWarning, uniqueFailed)

		ctrlReport.RuleReports[r].SetNumberOfResources(len(uniqueAll))
		ctrlReport.RuleReports[r].SetNumberOfWarningResources(len(uniqueWarning))
		ctrlReport.RuleReports[r].SetNumberOfFailedResources(len(uniqueFailed))

		uniqueAllResources = append(uniqueAllResources, uniqueAll...)
		uniqueWarningResources = append(uniqueWarningResources, uniqueWarning...)
		uniqueFailedResources = append(uniqueFailedResources, uniqueFailed...)
	}
	uniqueAllResources = GetUniqueResourcesIDs(uniqueAllResources)
	uniqueFailedResources = GetUniqueResourcesIDs(uniqueFailedResources)
	uniqueWarningResources = GetUniqueResourcesIDs(uniqueWarningResources)
	uniqueWarningResources = TrimUniqueIDs(uniqueWarningResources, uniqueFailedResources)
	return uniqueAllResources, uniqueWarningResources, uniqueFailedResources
}

// GetUniqueResources the list of resources can contain duplications, this function removes the resource duplication based on workloadinterface.GetID
func GetUniqueResources(k8sResources []map[string]interface{}) []map[string]interface{} {
	uniqueRuleResponses := map[string]bool{}

	lenK8sResources := len(k8sResources)
	for i := 0; i < lenK8sResources; i++ {
		workload := workloadinterface.NewObject(k8sResources[i])
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
			k8sResourcesNewList = append(k8sResourcesNewList, k8sResourcesList[i])
			uniqueRuleResponses[k8sResourcesList[i]] = true
		}
	}
	return k8sResourcesNewList
}

// TrimUniqueResources trim the list, this wil trim in case the same resource appears in the warning list and in the failed list
func TrimUniqueResources(origin, trimFrom []map[string]interface{}) []map[string]interface{} {
	if len(origin) == 0 || len(trimFrom) == 0 { // if there is nothing to trim
		return origin
	}
	uniqueResources := map[string]bool{}

	for i := range trimFrom {
		workload := workloadinterface.NewObject(trimFrom[i])
		uniqueResources[workload.GetID()] = true
	}

	lenOrigin := len(origin)
	for i := 0; i < lenOrigin; i++ {
		workload := workloadinterface.NewObject(origin[i])
		if found := uniqueResources[workload.GetID()]; found {
			// resource found -> remove from slice
			origin = removeFromSlice(origin, i)
			lenOrigin -= 1
			i -= 1
		}
	}
	return origin
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
