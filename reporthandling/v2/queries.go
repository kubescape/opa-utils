package v2

import (
	"github.com/armosec/k8s-interface/workloadinterface"
	"github.com/armosec/utils-go/str"
	"github.com/kubescape/opa-utils/objectsenvelopes"
	"github.com/kubescape/opa-utils/reporthandling/results/v1/resourcesresults"
)

// GetResource get single resource in IMetadata interface representation
func (postureReport *PostureReport) GetResource(resourceID string) workloadinterface.IMetadata {
	for i := range postureReport.Resources {
		if postureReport.Resources[i].ResourceID == resourceID {
			if m, ok := postureReport.Resources[i].Object.(map[string]interface{}); ok {
				return objectsenvelopes.NewObject(m)
			}
			break
		}
	}
	return nil
}

func (postureReport *PostureReport) ListResultsWithControlID(controlID string) []resourcesresults.Result {
	results := []resourcesresults.Result{}
	for i := range postureReport.Results {
		if str.StringInSliceCaseInsensitive(postureReport.Results[i].ListControlsIDs(nil).All(), controlID) {
			results = append(results, postureReport.Results[i])
		}
	}
	return results
}

func (postureReport *PostureReport) ListResultsWithControlName(name string) []resourcesresults.Result {
	results := []resourcesresults.Result{}
	for i := range postureReport.Results {
		if str.StringInSliceCaseInsensitive(postureReport.Results[i].ListControlsNames(nil).All(), name) {
			results = append(results, postureReport.Results[i])
		}
	}
	return results
}

func (postureReport *PostureReport) ListResultsWithRuleName(ruleName string) []resourcesresults.Result {
	results := []resourcesresults.Result{}
	for i := range postureReport.Results {
		if str.StringInSliceCaseInsensitive(postureReport.Results[i].ListRulesNames(nil).All(), ruleName) {
			results = append(results, postureReport.Results[i])
		}
	}
	return results
}
