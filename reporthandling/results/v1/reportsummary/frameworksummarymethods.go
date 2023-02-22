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
func (frameworkSummary *FrameworkSummary) ListResourcesIDs() *helpersv1.AllLists {
	return frameworkSummary.Controls.ListResourcesIDs()
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

	frameworkSummary.StatusCounters.Set(frameworkSummary.ListResourcesIDs())
	frameworkSummary.CalculateStatus()
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
	controls.ToUniqueControls()
	return controls
}

func (frameworkSummary *FrameworkSummary) ListControlsIDs() *helpersv1.AllLists {
	controls := &helpersv1.AllLists{}
	for controlID, controlSummary := range frameworkSummary.Controls {
		controls.Append(controlSummary.GetStatus().Status(), controlID)
	}
	controls.ToUniqueControls()
	return controls
}

// ListControls list all controls
func (frameworkSummary *FrameworkSummary) ListControls() []IControlSummary {
	controls := make([]IControlSummary, len(frameworkSummary.Controls))
	iter := frameworkSummary.Controls.ListControlsIDs().All()
	i := 0
	for iter.HasNext() {
		controls[i] = frameworkSummary.Controls.GetControl(EControlCriteriaID, iter.Next())
	}
	return controls
}

// Controls return the controls
func (frameworkSummary *FrameworkSummary) GetControls() IControlsSummaries {
	return &frameworkSummary.Controls
}
