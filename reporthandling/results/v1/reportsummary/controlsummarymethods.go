package reportsummary

import (
	"github.com/armosec/opa-utils/reporthandling/apis"
	helpersv1 "github.com/armosec/opa-utils/reporthandling/helpers/v1"
)

// =================================== Status ============================================

// GetStatus get the control status. returns an apis.ScanningStatus object
func (controlSummary *ControlSummary) GetStatus() apis.IStatus {
	if controlSummary.Status == apis.StatusUnknown {
		controlSummary.CalculateStatus()
	}
	return helpersv1.NewStatus(controlSummary.Status)
}

// CalculateStatus set the control status based on the resource counters
func (controlSummary *ControlSummary) CalculateStatus() {
	controlSummary.Status = calculateStatus(&controlSummary.ResourceCounters)
}

// =================================== Counters ============================================

// NumberOf get the number of resources
func (controlSummary *ControlSummary) NumberOfResources() ICounters {
	return &controlSummary.ResourceCounters
}

// Increase increases the counter based on the status
func (controlSummary *ControlSummary) increase(status apis.IStatus) {
	controlSummary.ResourceCounters.Increase(status)
}

// List resources IDs
func (controlSummary *ControlSummary) List() *helpersv1.AllLists {
	return &controlSummary.ResourceIDs
}

// Append increases the counter based on the status
func (controlSummary *ControlSummary) Append(status apis.IStatus, ids ...string) {
	for i := range ids {
		controlSummary.ResourceIDs.Append(status.Status(), ids[i])
		controlSummary.increase(status)
	}
}

// =================================== Score ============================================

// GetScore return control score
func (controlSummary *ControlSummary) GetScore() float32 {
	return controlSummary.Score
}

// =================================== Name ============================================

// GetName return control name
func (controlSummary *ControlSummary) GetName() string {
	return controlSummary.Name
}

// GetID return control ID
func (controlSummary *ControlSummary) GetID() string {
	return controlSummary.ControlID
}

// initResourcesSummary must run this AFTER initializing the controls
func (controlSummary *ControlSummary) initResourcesSummary() {
	controlSummary.CalculateStatus()
}
