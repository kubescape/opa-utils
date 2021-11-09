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
	uniqueAllFramework := []map[string]interface{}{}
	uniqueWarningFramework := []map[string]interface{}{}
	uniqueFailedFramework := []map[string]interface{}{}
	for c := range frameworkReport.ControlReports {
		uniqueAllControls, uniqueWarningControls, uniqueFailedControls := GetResourcesPerControl(&frameworkReport.ControlReports[c])

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
	uniqueAllFramework = GetUniqueResources(uniqueAllFramework)
	uniqueWarningFramework = GetUniqueResources(uniqueWarningFramework)
	uniqueFailedFramework = GetUniqueResources(uniqueFailedFramework)
	uniqueWarningFramework = TrimUniqueResources(uniqueWarningFramework, uniqueFailedFramework)

	// Set
	frameworkReport.SetNumberOfResources(len(uniqueAllFramework))
	frameworkReport.SetNumberOfWarningResources(len(uniqueWarningFramework))
	frameworkReport.SetNumberOfFailedResources(len(uniqueFailedFramework))
}

// GetResourcesPerControl - return unique lists of resources: all,warning,failed
func GetResourcesPerControl(ctrlReport *ControlReport) ([]map[string]interface{}, []map[string]interface{}, []map[string]interface{}) {
	uniqueAllResources := []map[string]interface{}{}
	uniqueWarningResources := []map[string]interface{}{}
	uniqueFailedResources := []map[string]interface{}{}
	for r := range ctrlReport.RuleReports {

		uniqueAll := GetUniqueResources(ctrlReport.RuleReports[r].GetAllResources())
		uniqueFailed := GetUniqueResources(ctrlReport.RuleReports[r].GetFailedResources())
		uniqueWarning := GetUniqueResources(ctrlReport.RuleReports[r].GetWarnignResources())
		uniqueWarning = TrimUniqueResources(uniqueWarning, uniqueFailed)

		ctrlReport.RuleReports[r].SetNumberOfResources(len(uniqueAll))
		ctrlReport.RuleReports[r].SetNumberOfWarningResources(len(uniqueWarning))
		ctrlReport.RuleReports[r].SetNumberOfFailedResources(len(uniqueFailed))

		uniqueAllResources = append(uniqueAllResources, uniqueAll...)
		uniqueWarningResources = append(uniqueWarningResources, uniqueWarning...)
		uniqueFailedResources = append(uniqueFailedResources, uniqueFailed...)
	}
	uniqueAllResources = GetUniqueResources(uniqueAllResources)
	uniqueFailedResources = GetUniqueResources(uniqueFailedResources)
	uniqueWarningResources = GetUniqueResources(uniqueWarningResources)
	uniqueWarningResources = TrimUniqueResources(uniqueWarningResources, uniqueFailedResources)
	return uniqueAllResources, uniqueWarningResources, uniqueFailedResources
}

// GetUniqueResources the list of resources can contain duplications, this function removes the resource duplication based on workloadinterface.GetID
func GetUniqueResources(k8sResources []map[string]interface{}) []map[string]interface{} {
	uniqueRuleResponses := map[string]bool{}

	lenK8sResources := len(k8sResources)
	for i := 0; i < lenK8sResources; i++ {
		workload := workloadinterface.NewWorkloadObj(k8sResources[i])
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

// TrimUniqueResources trim the list, this wil trim in case the same resource appears in the warning list and in the failed list
func TrimUniqueResources(origin, trimFrom []map[string]interface{}) []map[string]interface{} {
	if len(origin) == 0 || len(trimFrom) == 0 { // if there is nothing to trim
		return origin
	}
	uniqueResources := map[string]bool{}

	for i := range trimFrom {
		workload := workloadinterface.NewWorkloadObj(trimFrom[i])
		workload.GetVersion()
		uniqueResources[workload.GetID()] = true
	}

	lenOrigin := len(origin)
	for i := 0; i < lenOrigin; i++ {
		workload := workloadinterface.NewWorkloadObj(origin[i])
		if found := uniqueResources[workload.GetID()]; found {
			// resource found -> remove from slice
			origin = removeFromSlice(origin, i)
			lenOrigin -= 1
			i -= 1
		}
	}
	return origin
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

// type uniqueResources struct {
// 	allResources     []map[string]interface{}
// 	failedResources  []map[string]interface{}
// 	warningResources []map[string]interface{}
// }
