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
		uniqueAllControls := []map[string]interface{}{}
		uniqueWarningControls := []map[string]interface{}{}
		uniqueFailedControls := []map[string]interface{}{}
		for r := range frameworkReport.ControlReports[c].RuleReports {

			// Get
			uniqueAll := GetUniqueResources(frameworkReport.ControlReports[c].RuleReports[r].GetAllResources())
			uniqueFailed := GetUniqueResources(frameworkReport.ControlReports[c].RuleReports[r].GetFailedResources())
			uniqueWarning := GetUniqueResources(frameworkReport.ControlReports[c].RuleReports[r].GetWarnignResources())
			uniqueWarning = TrimUniqueResources(uniqueWarning, uniqueFailed)

			// Set
			frameworkReport.ControlReports[c].RuleReports[r].SetNumberOfResources(len(uniqueAll))
			frameworkReport.ControlReports[c].RuleReports[r].SetNumberOfWarningResources(len(uniqueWarning))
			frameworkReport.ControlReports[c].RuleReports[r].SetNumberOfFailedResources(len(uniqueFailed))

			// Append
			uniqueAllControls = append(uniqueAllControls, uniqueAll...)
			uniqueWarningControls = append(uniqueWarningControls, uniqueWarning...)
			uniqueFailedControls = append(uniqueFailedControls, uniqueFailed...)
		}
		uniqueAllControls = GetUniqueResources(uniqueAllControls)
		uniqueFailedControls = GetUniqueResources(uniqueFailedControls)
		uniqueWarningControls = GetUniqueResources(uniqueWarningControls)
		uniqueWarningControls = TrimUniqueResources(uniqueWarningControls, uniqueFailedControls)

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
