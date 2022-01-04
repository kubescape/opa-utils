package resourcesresults

import (
	"github.com/armosec/opa-utils/reporthandling/apis"
	helpersv1 "github.com/armosec/opa-utils/reporthandling/helpers/v1"
)

// SetResourceID set the resource ID
func (result *Result) SetResourceID(resourceID string) {
	result.ResourceID = resourceID
}

// GetResourceID get the resource ID
func (result *Result) GetResourceID() string {
	return result.ResourceID
}

// =============================== Status ====================================

// Status get resource status
func (result *Result) GetStatus(f *helpersv1.Filters) apis.IStatus {
	status := apis.StatusUnknown // Resource was not tested
	for i := range result.AssociatedControls {
		status = apis.Compare(status, result.AssociatedControls[i].GetStatus(f).Status())
	}
	return helpersv1.NewStatus(status)
}

// ================================= Listing ==================================

// // ListControls return list of controls
func (result *Result) ListControls() []ResourceAssociatedControl {
	return result.AssociatedControls
}

// ListFailedControls return list of failed controls IDs
func (result *Result) ListControlsIDs(f *helpersv1.Filters) *helpersv1.AllLists {
	controls := &helpersv1.AllLists{}
	for i := range result.AssociatedControls {
		s := result.AssociatedControls[i].GetStatus(f).Status()
		controls.Append(s, result.AssociatedControls[i].GetID())
	}
	return controls
}

// ListFailedControls return list of controls IDs
func (result *Result) ListControlsNames(f *helpersv1.Filters) *helpersv1.AllLists {
	controls := &helpersv1.AllLists{}
	for i := range result.AssociatedControls {
		s := result.AssociatedControls[i].GetStatus(f).Status()
		controls.Append(s, result.AssociatedControls[i].GetName())
	}
	return controls
}

// ListRulesNames return list of rules names
func (result *Result) ListRulesNames(f *helpersv1.Filters) *helpersv1.AllLists {
	rules := &helpersv1.AllLists{}
	for i := range result.AssociatedControls {
		for j := range result.AssociatedControls[i].ResourceAssociatedRules {
			s := result.AssociatedControls[i].GetStatus(f).Status()
			rules.Append(s, result.AssociatedControls[i].ResourceAssociatedRules[j].GetName())
		}
	}
	return rules
}

// ListRules return list of rules
func (result *Result) ListRules() []ResourceAssociatedRule {
	rules := []ResourceAssociatedRule{}
	ruleNames := map[string]bool{}
	for i := range result.AssociatedControls {
		for j := range result.AssociatedControls[i].ResourceAssociatedRules {
			if _, ok := ruleNames[result.AssociatedControls[i].ResourceAssociatedRules[j].GetName()]; !ok {
				rules = append(rules, result.AssociatedControls[i].ResourceAssociatedRules[j])
				ruleNames[result.AssociatedControls[i].ResourceAssociatedRules[j].GetName()] = true
			}
		}
	}
	return rules
}

// ListRulesOfControl return list of rules related to a controlID or controlName
func (result *Result) ListRulesOfControl(controlID, controlName string) []ResourceAssociatedRule {
	rules := []ResourceAssociatedRule{}
	ruleNames := map[string]bool{}
	for i := range result.AssociatedControls {
		if (controlID != "" && result.AssociatedControls[i].ControlID != controlID) || (controlName != "" && result.AssociatedControls[i].Name != controlName) {
			continue
		}
		for j := range result.AssociatedControls[i].ResourceAssociatedRules {
			if _, ok := ruleNames[result.AssociatedControls[i].ResourceAssociatedRules[j].GetName()]; !ok {
				rules = append(rules, result.AssociatedControls[i].ResourceAssociatedRules[j])
				ruleNames[result.AssociatedControls[i].ResourceAssociatedRules[j].GetName()] = true
			}
		}
	}
	return rules
}

// // ListRulesNames return list of rules names
// func (result *Result) ListRules(ls *helpersv1.ListingFilters) []ResourceAssociatedRule {
// 	rules := []ResourceAssociatedRule{}
// 	ruleNames := map[string]bool{}
// 	for i := range result.AssociatedControls {
// 		for j := range result.AssociatedControls[i].ResourceAssociatedRules {
// 			if _, ok := ruleNames[result.AssociatedControls[i].ResourceAssociatedRules[j].GetName()]; !ok { // check for uniqueness
// 				rules = append(rules, result.AssociatedControls[i].ResourceAssociatedRules[j])
// 				ruleNames[result.AssociatedControls[i].ResourceAssociatedRules[j].GetName()] = true
// 			}
// 		}
// 	}
// 	return rules
// }
