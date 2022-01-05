package reportsummary

import (
	"github.com/armosec/opa-utils/reporthandling/apis"
	helpersv1 "github.com/armosec/opa-utils/reporthandling/helpers/v1"
)

// =================================== Status ============================================

// Status get the framework status. returns an apis.ScanningStatus object
func (frameworkSummary *FrameworkSummary) GetStatus() apis.IStatus {
	if frameworkSummary.Status == apis.StatusUnknown {
		frameworkSummary.CalculateStatus()
	}
	return helpersv1.NewStatus(frameworkSummary.Status)
}

// SetStatus set the framework status based on the resource counters
func (frameworkSummary *FrameworkSummary) CalculateStatus() {
	frameworkSummary.Status = calculateStatus(&frameworkSummary.ResourceCounters)
}

// =================================== Counters ============================================

// NumberOfExcluded get the number of excluded resources
func (frameworkSummary *FrameworkSummary) NumberOfResources() ICounters {
	return &frameworkSummary.ResourceCounters
}

// Increase increases the counter based on the status
func (frameworkSummary *FrameworkSummary) Increase(status apis.IStatus) {
	frameworkSummary.ResourceCounters.Increase(status)
}

// List resources IDs
func (frameworkSummary *FrameworkSummary) List() *helpersv1.AllLists {
	return &frameworkSummary.resourceIDs
}

// initResourcesSummary must run this AFTER initializing the controls
func (frameworkSummary *FrameworkSummary) initResourcesSummary() {
	frameworkSummary.resourceIDs = helpersv1.AllLists{}
	for k, control := range frameworkSummary.Controls {
		control.initResourcesSummary()
		frameworkSummary.Controls[k] = control
		frameworkSummary.resourceIDs.Update(control.List())
	}

	frameworkSummary.ResourceCounters.Set(&frameworkSummary.resourceIDs)
	frameworkSummary.CalculateStatus()
}

// Append increases the counter based on the status
func (frameworkSummary *FrameworkSummary) Append(status apis.IStatus, ids ...string) {
	for i := range ids {
		frameworkSummary.resourceIDs.Append(status.Status(), ids[i])
	}
}

// =================================== Score ============================================

// GetScore return framework score
func (frameworkSummary *FrameworkSummary) GetScore() float32 {
	return frameworkSummary.Score
}

// =================================== Name ============================================

// GetName return framework name
func (frameworkSummary *FrameworkSummary) GetName() string {
	return frameworkSummary.Name
}

// =========================================== List Controls ====================================

// ListControlsNames list all framework names
func (frameworkSummary *FrameworkSummary) ListControlsNames() *helpersv1.AllLists {
	controls := &helpersv1.AllLists{}
	for _, controlSummary := range frameworkSummary.Controls {
		controls.Append(controlSummary.GetStatus().Status(), controlSummary.Name)
	}
	return controls
}

func (frameworkSummary *FrameworkSummary) ListControlsIDs() *helpersv1.AllLists {
	controls := &helpersv1.AllLists{}
	for controlID, controlSummary := range frameworkSummary.Controls {
		controls.Append(controlSummary.GetStatus().Status(), controlID)
	}
	return controls
}

// ListFrameworks list all frameworks
func (frameworkSummary *FrameworkSummary) ListControls() *ListPolicies {
	controls := ListPolicies{}
	for i := range frameworkSummary.Controls {
		control := frameworkSummary.Controls[i]
		controls.Append(control.GetStatus().Status(), &control)
	}
	return &controls
}
