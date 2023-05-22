package reportsummary

import (
	"github.com/kubescape/opa-utils/reporthandling/apis"
	helpersv1 "github.com/kubescape/opa-utils/reporthandling/helpers/v1"
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
	frameworkSummary.Status = calculateStatus(&frameworkSummary.StatusCounters)
}

// =================================== Counters ============================================

// NumberOfExcluded get the number of excluded resources
func (frameworkSummary *FrameworkSummary) NumberOfResources() ICounters {
	return &frameworkSummary.StatusCounters
}

// Increase increases the counter based on the status
func (frameworkSummary *FrameworkSummary) Increase(status apis.IStatus) {
	frameworkSummary.StatusCounters.Increase(status)
}

// List resources IDs
//
// If an optional pointer to an AllLists object is provided as a parameter, it will be used to store the results,
// avoiding unnecessary memory allocations. If the parameter is nil, a new AllLists object will be created and returned.
func (frameworkSummary *FrameworkSummary) ListResourcesIDs(l *helpersv1.AllLists) *helpersv1.AllLists {
	return frameworkSummary.Controls.ListResourcesIDs(l)
}

func (frameworkSummary *FrameworkSummary) NumberOfControls() ICounters {
	controlsCounters := &PostureCounters{}
	for _, ctrlSummary := range frameworkSummary.Controls {
		controlsCounters.Increase(ctrlSummary.GetStatus())
	}

	return controlsCounters
}

// initResourcesSummary must run this AFTER initializing the controls
func (frameworkSummary *FrameworkSummary) initResourcesSummary(controlInfoMap map[string]apis.StatusInfo) {
	for k, control := range frameworkSummary.Controls {
		if statusInfo, ok := controlInfoMap[control.ControlID]; ok && statusInfo.InnerStatus != apis.StatusUnknown {
			control.SetStatus(&statusInfo)
		} else if control.GetStatus().Status() == apis.StatusUnknown {
			control.CalculateStatus()
		}
		frameworkSummary.Controls[k] = control
	}

	l := helpersv1.GetAllListsFromPool()
	defer helpersv1.PutAllListsToPool(l)
	frameworkSummary.StatusCounters.Set(frameworkSummary.ListResourcesIDs(l))
	frameworkSummary.CalculateStatus()
}

// =================================== ComplianceScore ============================================

// GetComplianceScore returns framework ComplianceScore
func (frameworkSummary *FrameworkSummary) GetComplianceScore() float32 {
	return frameworkSummary.ComplianceScore
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
//
// If an optional pointer to an AllLists object is provided as a parameter, it will be used to store the results,
// avoiding unnecessary memory allocations. If the parameter is nil, a new AllLists object will be created and returned.
func (frameworkSummary *FrameworkSummary) ListControlsNames(controls *helpersv1.AllLists) *helpersv1.AllLists {
	if controls == nil {
		controls = &helpersv1.AllLists{}
	}

	controls.Initialize(len(frameworkSummary.Controls))
	for _, controlSummary := range frameworkSummary.Controls {
		controls.Append(controlSummary.GetStatus().Status(), controlSummary.Name)
	}
	return controls
}

// ListControlsIDs list all controls IDs in the framework
//
// If an optional pointer to an AllLists object is provided as a parameter, it will be used to store the results,
// avoiding unnecessary memory allocations. If the parameter is nil, a new AllLists object will be created and returned.
func (frameworkSummary *FrameworkSummary) ListControlsIDs(controls *helpersv1.AllLists) *helpersv1.AllLists {
	if controls == nil {
		controls = &helpersv1.AllLists{}
	}
	controls.Initialize(len(frameworkSummary.Controls))
	for controlID, controlSummary := range frameworkSummary.Controls {
		controls.Append(controlSummary.GetStatus().Status(), controlID)
	}
	return controls
}

// ListControls list all controls
func (frameworkSummary *FrameworkSummary) ListControls() []IControlSummary {
	controls := make([]IControlSummary, len(frameworkSummary.Controls))
	i := 0

	l := helpersv1.GetAllListsFromPool()
	defer helpersv1.PutAllListsToPool(l)

	for ctrlId := range frameworkSummary.Controls.ListControlsIDs(l).All() {
		controls[i] = frameworkSummary.Controls.GetControl(EControlCriteriaID, ctrlId)
		i++
	}
	return controls
}

// Controls return the controls
func (frameworkSummary *FrameworkSummary) GetControls() IControlsSummaries {
	return &frameworkSummary.Controls
}
