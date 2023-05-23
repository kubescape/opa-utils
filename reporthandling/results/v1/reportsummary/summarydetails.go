package reportsummary

import (
	"github.com/kubescape/opa-utils/reporthandling/apis"
	helpersv1 "github.com/kubescape/opa-utils/reporthandling/helpers/v1"
	"github.com/kubescape/opa-utils/reporthandling/results/v1/resourcesresults"
)

// =================================== Status ============================================

// Status get the scan status. returns an apis.ScanningStatus object
func (summaryDetails *SummaryDetails) GetStatus() *helpersv1.Status {
	if summaryDetails.Status == apis.StatusUnknown {
		summaryDetails.CalculateStatus()
	}
	return helpersv1.NewStatus(summaryDetails.Status)
}

// GetScore return score
func (summaryDetails *SummaryDetails) GetScore() float32 {
	return summaryDetails.Score
}

// SetStatus set the framework status based on the resource counters
func (summaryDetails *SummaryDetails) CalculateStatus() {
	summaryDetails.Status = calculateStatus(&summaryDetails.StatusCounters)
}

// =================================== Counters ============================================

// NumberOfExcluded get the number of excluded resources
func (summaryDetails *SummaryDetails) NumberOfResources() ICounters {
	return &summaryDetails.StatusCounters
}

// Increase increases the counter based on the status
func (summaryDetails *SummaryDetails) Increase(status apis.IStatus) {
	summaryDetails.StatusCounters.Increase(status)
}

// InitResourcesSummary must run this AFTER initializing the controls
func (summaryDetails *SummaryDetails) InitResourcesSummary(controlInfoMap map[string]apis.StatusInfo) {
	for i := range summaryDetails.Frameworks {
		summaryDetails.Frameworks[i].initResourcesSummary(controlInfoMap)
	}

	for k, control := range summaryDetails.Controls {
		if statusInfo, ok := controlInfoMap[control.ControlID]; ok && statusInfo.InnerStatus != apis.StatusUnknown {
			control.SetStatus(&statusInfo)
		} else if control.GetStatus().Status() == apis.StatusUnknown {
			control.CalculateStatus()
		}
		summaryDetails.Controls[k] = control

		// Summarize the failed controls severity
		if control.GetStatus().IsFailed() {
			summaryDetails.ControlsSeverityCounters.Increase(apis.ControlSeverityToString(control.GetScoreFactor()), 1)
		}
	}

	l := helpersv1.GetAllListsFromPool()
	defer helpersv1.PutAllListsToPool(l)

	summaryDetails.StatusCounters.Set(summaryDetails.ListResourcesIDs(l))
	summaryDetails.CalculateStatus()
}

// =========================================== List Frameworks ====================================

// ListFrameworksNames list all framework names
func (summaryDetails *SummaryDetails) ListFrameworksNames() *helpersv1.AllLists {
	frameworks := &helpersv1.AllLists{}
	frameworks.Initialize(len(summaryDetails.Frameworks))
	for i := range summaryDetails.Frameworks {
		frameworks.Append(summaryDetails.Frameworks[i].GetStatus().Status(), summaryDetails.Frameworks[i].GetName())
	}
	return frameworks
}

// ListFrameworks list all frameworks
func (summaryDetails *SummaryDetails) ListFrameworks() []IFrameworkSummary {
	frameworks := []IFrameworkSummary{}
	for i := range summaryDetails.Frameworks {
		frameworks = append(frameworks, &summaryDetails.Frameworks[i])
	}
	return frameworks
}

// =========================================== List Controls ====================================

// ListControlsNames list all framework names
func (summaryDetails *SummaryDetails) ListControlsNames() *helpersv1.AllLists {
	controls := &helpersv1.AllLists{}
	controls.Initialize(len(summaryDetails.Controls))
	for _, controlSummary := range summaryDetails.Controls {
		controls.Append(controlSummary.GetStatus().Status(), controlSummary.Name)
	}
	return controls
}

func (summaryDetails *SummaryDetails) ListControlsIDs() *helpersv1.AllLists {
	controls := &helpersv1.AllLists{}
	controls.Initialize(len(summaryDetails.Controls))
	for controlID, controlSummary := range summaryDetails.Controls {
		controls.Append(controlSummary.GetStatus().Status(), controlID)
	}
	return controls
}

// ListControls list all controls
func (summaryDetails *SummaryDetails) ListControls() []IControlSummary {
	controls := make([]IControlSummary, len(summaryDetails.Controls))
	i := 0
	for ctrlId := range summaryDetails.ListControlsIDs().All() {
		controls[i] = summaryDetails.Controls.GetControl(EControlCriteriaID, ctrlId)
		i++
	}
	return controls
}

// NumberOfControls get number of controls
func (summaryDetails *SummaryDetails) NumberOfControls() ICounters {
	controlsIds := summaryDetails.ListControlsIDs()
	return &PostureCounters{
		PassedCounter:  controlsIds.Passed(),
		FailedCounter:  controlsIds.Failed(),
		SkippedCounter: controlsIds.Skipped(),
	}
}

// ================================================================================
func (summaryDetails *SummaryDetails) ControlName(controlID string) string {
	if c, ok := summaryDetails.Controls[controlID]; ok {
		return c.Name
	}
	return ""
}

// ListResourcesIDs list all resources IDs
//
// If an optional pointer to an AllLists object is provided as a parameter, it will be used to store the results,
// avoiding unnecessary memory allocations. If the parameter is nil, a new AllLists object will be created and returned.
func (summaryDetails *SummaryDetails) ListResourcesIDs(l *helpersv1.AllLists) *helpersv1.AllLists {
	return summaryDetails.Controls.ListResourcesIDs(l)
}

// AppendResourceResult appends the given resource result to the summary
//
// Updates any necessary info accordingly
func (summaryDetails *SummaryDetails) AppendResourceResult(resourceResult *resourcesresults.Result) {

	// update full-summary counter
	updateControlsSummaryCounters(resourceResult, summaryDetails.Controls, nil)

	// update the summary’s severity counters
	if resourceResult.GetStatus(nil).IsFailed() {
		for _, resourceControl := range resourceResult.ListControls() {
			if resourceControl.GetStatus(nil).IsFailed() {
				control := summaryDetails.Controls.GetControl(EControlCriteriaID, resourceControl.GetID())
				severityScore := control.GetScoreFactor()
				severity := apis.ControlSeverityToString(severityScore)
				summaryDetails.ResourcesSeverityCounters.Increase(severity, 1)
			}
		}
	}

	// update frameworks counters
	for i := range summaryDetails.Frameworks {
		updateControlsSummaryCounters(resourceResult, summaryDetails.Frameworks[i].Controls, &helpersv1.Filters{FrameworkNames: []string{summaryDetails.Frameworks[i].GetName()}})

		l := helpersv1.GetAllListsFromPool()
		defer helpersv1.PutAllListsToPool(l)

		summaryDetails.Frameworks[i].StatusCounters.Set(summaryDetails.Frameworks[i].ListResourcesIDs(l))
		summaryDetails.Frameworks[i].CalculateStatus()
	}
}

// GetResourcesSeverityCounters get the resources severity counters
func (summaryDetails *SummaryDetails) GetResourcesSeverityCounters() ISeverityCounters {
	return &summaryDetails.ResourcesSeverityCounters
}

// GetResourcesSeverityCounters get the resources severity counters
func (summaryDetails *SummaryDetails) GetControlsSeverityCounters() ISeverityCounters {
	return &summaryDetails.ControlsSeverityCounters
}

func updateControlsSummaryCounters(resourceResult *resourcesresults.Result, controls map[string]ControlSummary, f *helpersv1.Filters) {
	// update controls counters
	for i := range resourceResult.AssociatedControls {
		controlID := resourceResult.AssociatedControls[i].ControlID
		if controlSummary, ok := controls[controlID]; ok {
			subStatus := resourceResult.AssociatedControls[i].GetSubStatus()
			status := resourceResult.AssociatedControls[i].GetStatus(f)
			controlSummary.Append(status, resourceResult.ResourceID)
			controlSummary.calculateStatus(subStatus)
			controls[controlID] = controlSummary
		}
	}
}
