package reporthandling

import (
	"fmt"
	"hash/fnv"

	"github.com/kubescape/k8s-interface/workloadinterface"
	"github.com/kubescape/opa-utils/objectsenvelopes"
	"github.com/kubescape/opa-utils/reporthandling/apis"
	"github.com/mitchellh/mapstructure"
)

const ActionRequiredAttribute string = "actionRequired"

// ==============================================================================================
// ========================== PostureReport =====================================================
// ==============================================================================================

// TODO - receive list full json paths
func (postureReport *PostureReport) RemoveData(keepFields, keepMetadataFields []string) {
	for i := range postureReport.FrameworkReports {
		postureReport.FrameworkReports[i].RemoveData(keepFields, keepMetadataFields)
	}
}

// ==============================================================================================
// ========================= FrameworkReport ====================================================
// ==============================================================================================

func (frameworkReport *FrameworkReport) RemoveData(keepFields, keepMetadataFields []string) {
	for i := range frameworkReport.ControlReports {
		frameworkReport.ControlReports[i].RemoveData(keepFields, keepMetadataFields)
	}
}

func (frameworkReport *FrameworkReport) SetResourcesCounters() {
	resourcesIDs := frameworkReport.ListResourcesIDs()
	frameworkReport.SetNumberOfResources(len(resourcesIDs.GetAllResources()))
	frameworkReport.SetNumberOfWarningResources(len(resourcesIDs.GetWarningResources()))
	frameworkReport.SetNumberOfFailedResources(len(resourcesIDs.GetFailedResources()))
}
func (frameworkReport *FrameworkReport) SetNumberOfResources(n int) {
	frameworkReport.TotalResources = n
}

func (frameworkReport *FrameworkReport) SetNumberOfWarningResources(n int) {
	frameworkReport.WarningResources = n
}

func (frameworkReport *FrameworkReport) SetNumberOfFailedResources(n int) {
	frameworkReport.FailedResources = n
}
func (frameworkReport *FrameworkReport) GetNumberOfWarningResources() int {
	return frameworkReport.WarningResources
}

func (frameworkReport *FrameworkReport) GetNumberOfResources() int {
	return frameworkReport.TotalResources
}

func (frameworkReport *FrameworkReport) GetNumberOfFailedResources() int {
	return frameworkReport.FailedResources
}

func (frameworkReport *FrameworkReport) GetStatus() string {
	if frameworkReport.Passed() {
		return StatusPassed
	}
	if frameworkReport.Warning() {
		return StatusWarning
	}
	return StatusFailed
}

// GetResourcesPerControl - return unique lists of resource IDs: all,warning,failed
func (frameworkReport *FrameworkReport) ListResourcesIDs() *ResourcesIDs {
	resourcesIDs := ResourcesIDs{}
	for c := range frameworkReport.ControlReports {
		resourcesIDs.append(frameworkReport.ControlReports[c].ListResourcesIDs())
	}
	return &resourcesIDs
}
func (frameworkReport *FrameworkReport) Passed() bool {
	for _, r := range frameworkReport.ControlReports {
		if r.Failed() || r.Warning() {
			return false
		}
	}
	return true
}

func (frameworkReport *FrameworkReport) Warning() bool {
	if frameworkReport.Passed() || frameworkReport.Failed() {
		return false
	}
	for _, r := range frameworkReport.ControlReports {
		if r.Warning() {
			return true
		}
	}
	return false
}

func (frameworkReport *FrameworkReport) Failed() bool {
	if frameworkReport.Passed() {
		return false
	}
	for _, r := range frameworkReport.ControlReports {
		if r.Failed() {
			return true
		}
	}
	return false
}
func (frameworkReport *FrameworkReport) SetDefaultScore() {
	frameworkReport.Score = float32(percentage(frameworkReport.GetNumberOfResources(), frameworkReport.GetNumberOfFailedResources()))
}

// ==============================================================================================
// ========================== ControlReport =====================================================
// ==============================================================================================

func (controlReport *ControlReport) SetResourcesCounters() {
	resourcesIDs := controlReport.ListResourcesIDs()
	controlReport.SetNumberOfResources(len(resourcesIDs.GetAllResources()))
	controlReport.SetNumberOfWarningResources(len(resourcesIDs.GetWarningResources()))
	controlReport.SetNumberOfFailedResources(len(resourcesIDs.GetFailedResources()))
}
func (controlReport *ControlReport) SetNumberOfResources(n int) {
	controlReport.TotalResources = n
}

func (controlReport *ControlReport) SetNumberOfWarningResources(n int) {
	controlReport.WarningResources = n
}

func (controlReport *ControlReport) SetNumberOfFailedResources(n int) {
	controlReport.FailedResources = n
}
func (controlReport *ControlReport) GetNumberOfWarningResources() int {
	return controlReport.WarningResources
}

func (controlReport *ControlReport) GetNumberOfResources() int {
	return controlReport.TotalResources
}

func (controlReport *ControlReport) GetNumberOfFailedResources() int {
	return controlReport.FailedResources
}

// GetResourcesPerControl - return unique lists of resource IDs: all,warning,failed
func (controlReport *ControlReport) ListResourcesIDs() *ResourcesIDs {
	resourcesIDs := ResourcesIDs{}
	for r := range controlReport.RuleReports {
		resourcesIDs.append(controlReport.RuleReports[r].ListResourcesIDs())
	}
	return &resourcesIDs
}

func (controlReport *ControlReport) ListControlsInputKinds() []string {
	listControlsInputKinds := []string{}
	for i := range controlReport.RuleReports {
		listControlsInputKinds = append(listControlsInputKinds, controlReport.RuleReports[i].ListInputKinds...)
	}
	return listControlsInputKinds
}

func (controlReport *ControlReport) GetStatus() string {
	if controlReport.Passed() {
		return StatusPassed
	}
	if controlReport.Warning() {
		return StatusWarning
	}
	return StatusFailed
}
func (controlReport *ControlReport) Passed() bool {
	for _, r := range controlReport.RuleReports {
		if r.Failed() || r.Warning() {
			return false
		}
	}
	return true
}

func (controlReport *ControlReport) Warning() bool {
	if controlReport.Passed() || controlReport.Failed() {
		return false
	}
	for _, r := range controlReport.RuleReports {
		if r.Warning() {
			return true
		}
	}
	return false
}

func (controlReport *ControlReport) Failed() bool {
	if controlReport.Passed() {
		return false
	}
	for _, r := range controlReport.RuleReports {
		if r.Failed() {
			return true
		}
	}
	return false
}

func (controlReport *ControlReport) GetID() string {
	h := fnv.New32a()
	h.Write([]byte(controlReport.Name))
	s := fmt.Sprintf("%d", h.Sum32())

	return "C-" + s
}

func (controlReport *ControlReport) RemoveData(keepFields, keepMetadataFields []string) {
	for i := range controlReport.RuleReports {
		controlReport.RuleReports[i].RemoveData(keepFields, keepMetadataFields)
	}
}

func (controlReport *ControlReport) SetDefaultScore() {
	controlReport.Score = float32(percentage(controlReport.GetNumberOfResources(), controlReport.GetNumberOfFailedResources()))
}

// ==============================================================================================
// ============================ RuleReport ======================================================
// ==============================================================================================

func (ruleReport *RuleReport) GetStatus() string {
	if ruleReport.Passed() {
		return StatusPassed
	}
	if ruleReport.Warning() {
		return StatusWarning
	}
	return StatusFailed
}
func (ruleReport *RuleReport) Passed() bool {
	return len(ruleReport.RuleResponses) == 0
}

func (ruleReport *RuleReport) Warning() bool {
	if ruleReport.Passed() {
		return false
	}
	for _, r := range ruleReport.RuleResponses {
		if r.Failed() {
			return false
		}
	}
	return true
}
func (ruleReport *RuleReport) Failed() bool {
	if ruleReport.Passed() {
		return false
	}
	for _, r := range ruleReport.RuleResponses {
		if r.Failed() {
			return true
		}
	}
	return false
}

func (ruleReport *RuleReport) SetResourcesCounters() {
	resourcesIDs := ruleReport.ListResourcesIDs()
	ruleReport.SetNumberOfResources(len(resourcesIDs.GetAllResources()))
	ruleReport.SetNumberOfWarningResources(len(resourcesIDs.GetWarningResources()))
	ruleReport.SetNumberOfFailedResources(len(resourcesIDs.GetFailedResources()))
}

func (ruleReport *RuleReport) SetNumberOfResources(n int) {
	ruleReport.TotalResources = n
}

func (ruleReport *RuleReport) SetNumberOfWarningResources(n int) {
	ruleReport.WarningResources = n
}

func (ruleReport *RuleReport) SetNumberOfFailedResources(n int) {
	ruleReport.FailedResources = n
}

func (ruleReport *RuleReport) ListResourcesIDs() *ResourcesIDs {
	resourcesIDs := ResourcesIDs{}
	resourcesIDs.setFailedResources(GetUniqueResourcesIDs(workloadinterface.ListMetaIDs(objectsenvelopes.ListMapToMeta(ruleReport.GetFailedResources()))))
	resourcesIDs.setWarningResources(GetUniqueResourcesIDs(workloadinterface.ListMetaIDs(objectsenvelopes.ListMapToMeta(ruleReport.GetWarnignResources())))) // needs to be initialized after failed
	resourcesIDs.setPassedResources(ruleReport.GetAllResourcesIDs())                                                                                         // needs to be initialized after warning
	return &resourcesIDs
}
func (ruleReport *RuleReport) GetNumberOfWarningResources() int {
	return ruleReport.WarningResources
}

func (ruleReport *RuleReport) GetNumberOfResources() int {
	return ruleReport.TotalResources
}

func (ruleReport *RuleReport) GetNumberOfFailedResources() int {
	return ruleReport.FailedResources
}
func (ruleReport *RuleReport) RemoveData(keepFields, keepMetadataFields []string) {
	for i := range ruleReport.RuleResponses {
		ruleReport.RuleResponses[i].RemoveData(keepFields, keepMetadataFields)
	}
}

func (ruleReport *RuleReport) GetAllResourcesIDs() []string {
	return ruleReport.ListInputKinds
}

// DO NOT USE!
//
//	func (ruleReport *RuleReport) GetAllResources() []map[string]interface{} {
//		return ruleReport.ListInputResources
//	}
func (ruleReport *RuleReport) GetFailedResources() []map[string]interface{} {

	failedResources := []map[string]interface{}{}
	for _, ruleResponse := range ruleReport.RuleResponses {
		failedResources = append(failedResources, ruleResponse.GetFailedResources()...)
	}
	return failedResources
}

func (ruleReport *RuleReport) GetWarnignResources() []map[string]interface{} {

	failedResources := []map[string]interface{}{}
	for _, ruleResponse := range ruleReport.RuleResponses {
		failedResources = append(failedResources, ruleResponse.GetWarnignResources()...)
	}
	return failedResources
}

// ==============================================================================================
// =========================== RuleResponse =====================================================
// ==============================================================================================

func (ruleResponse *RuleResponse) GetFailedResources() []map[string]interface{} {

	failedResources := []map[string]interface{}{}
	if ruleResponse.Failed() {
		failedResources = append(failedResources, ruleResponse.AlertObject.K8SApiObjects...)
		if ruleResponse.AlertObject.ExternalObjects != nil {
			failedResources = append(failedResources, ruleResponse.AlertObject.ExternalObjects)
		}
	}
	return failedResources
}

func (ruleResponse *RuleResponse) GetWarnignResources() []map[string]interface{} {

	failedResources := []map[string]interface{}{}
	if ruleResponse.Warning() {
		failedResources = append(failedResources, ruleResponse.AlertObject.K8SApiObjects...)
		if ruleResponse.AlertObject.ExternalObjects != nil {
			failedResources = append(failedResources, ruleResponse.AlertObject.ExternalObjects)
		}
	}
	return failedResources
}
func (ruleResponse *RuleResponse) Passed() bool {
	return false
}

func (ruleResponse *RuleResponse) Warning() bool {
	if ruleResponse.Exception != nil {
		if ruleResponse.Exception.IsAlertOnly() {
			return true
		}
	}
	return false
}
func (ruleResponse *RuleResponse) Failed() bool {
	return ruleResponse.Exception == nil
}

func (ruleResponse *RuleResponse) GetStatus() string {
	if ruleResponse.Warning() {
		return StatusWarning
	}
	if ruleResponse.Passed() {
		return StatusPassed
	}
	return StatusFailed
}
func (ruleResponse *RuleResponse) RemoveData(keepFields, keepMetadataFields []string) {

	for i := range ruleResponse.AlertObject.K8SApiObjects {
		deleteFromMap(ruleResponse.AlertObject.K8SApiObjects[i], keepFields)
		for k := range ruleResponse.AlertObject.K8SApiObjects[i] {
			if k == "metadata" {
				if b, ok := ruleResponse.AlertObject.K8SApiObjects[i][k].(map[string]interface{}); ok {
					deleteFromMap(b, keepMetadataFields)
					ruleResponse.AlertObject.K8SApiObjects[i][k] = b
				}
			}
		}
	}
}

func (control *Control) GetAttackTrackCategories(attackTrackName string) []string {
	if v, exist := control.Attributes[ControlAttributeKeyAttackTracks]; exist {
		var attackTrackToCategories []AttackTrackCategories
		if err := mapstructure.Decode(v, &attackTrackToCategories); err == nil {
			for _, attackTrackToCategory := range attackTrackToCategories {
				if attackTrackToCategory.AttackTrack == attackTrackName {
					return attackTrackToCategory.Categories
				}
			}
		}
	}
	return []string{}
}

func (control *Control) GetControlTypeTags() []string {
	if v, exist := control.Attributes[ControlAttributeKeyTypeTag]; exist {
		tags := []string{}
		if err := mapstructure.Decode(v, &tags); err == nil {
			return tags
		}
	}
	return []string{}
}

func (control *Control) GetControlId() string {
	return control.ControlID
}

func (control *Control) GetScore() float64 {
	return float64(control.BaseScore)
}

func (control *Control) GetSeverity() int {
	return apis.ControlSeverityToInt(control.BaseScore)
}

func (control *Control) GetActionRequiredAttribute() string {
	if control.Attributes == nil {
		return ""
	}
	if v, ok := control.Attributes[ActionRequiredAttribute]; ok {
		if actionRequired, ok := v.(string); ok {
			return actionRequired
		}
	}
	return ""
}
