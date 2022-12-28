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
	summaryDetails.Status = calculateStatus(&summaryDetails.ResourceCounters)
}

// =================================== Counters ============================================

// NumberOfExcluded get the number of excluded resources
func (summaryDetails *SummaryDetails) NumberOfResources() ICounters {
	return &summaryDetails.ResourceCounters
}

// Increase increases the counter based on the status
func (summaryDetails *SummaryDetails) Increase(status apis.IStatus) {
	summaryDetails.ResourceCounters.Increase(status)
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

	summaryDetails.ResourceCounters.Set(summaryDetails.ListResourcesIDs())
	summaryDetails.CalculateStatus()
}

// =========================================== List Frameworks ====================================

// ListFrameworksNames list all framework names
func (summaryDetails *SummaryDetails) ListFrameworksNames() *helpersv1.AllLists {
	frameworks := &helpersv1.AllLists{}
	for i := range summaryDetails.Frameworks {
		frameworks.Append(summaryDetails.Frameworks[i].GetStatus().Status(), summaryDetails.Frameworks[i].GetName())
	}
	frameworks.ToUniqueControls()
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
	for _, controlSummary := range summaryDetails.Controls {
		status := controlSummary.GetSubStatus().Status()
		if status == apis.StatusUnknown {
			status = controlSummary.GetStatus().Status()
		}
		controls.Append(status, controlSummary.Name)
	}
	controls.ToUniqueControls()
	return controls
}

func (summaryDetails *SummaryDetails) ListControlsIDs() *helpersv1.AllLists {
	controls := &helpersv1.AllLists{}
	for controlID, controlSummary := range summaryDetails.Controls {
		status := controlSummary.GetSubStatus().Status()
		if status == apis.StatusUnknown {
			status = controlSummary.GetStatus().Status()
		}
		controls.Append(status, controlID)
	}
	controls.ToUniqueControls()
	return controls
}

// ListControls list all controls
func (summaryDetails *SummaryDetails) ListControls() []IControlSummary {
	controls := make([]IControlSummary, len(summaryDetails.Controls))
	iter := summaryDetails.ListControlsIDs().All()
	i := 0
	for iter.HasNext() {
		controls[i] = summaryDetails.Controls.GetControl(EControlCriteriaID, iter.Next())
		i++
	}
	return controls
}

// NumberOfControls get number of controls
func (summaryDetails *SummaryDetails) NumberOfControls() ICounters {
	controlsIds := summaryDetails.ListControlsIDs()
	return &PostureCounters{
		PassedCounter:                len(controlsIds.Passed()),
		PassedExceptionCounter:       len(controlsIds.PassedExceptions()),
		PassedIrrelevantCounter:      len(controlsIds.PassedIrrelevant()),
		FailedCounter:                len(controlsIds.Failed()),
		SkippedConfigurationCounter:  len(controlsIds.SkippedConfiguration()),
		SkippedIntegrationCounter:    len(controlsIds.SkippedIntegration()),
		SkippedRequiresReviewCounter: len(controlsIds.SkippedRequiresReview()),
		SkippedManualReviewCounter:   len(controlsIds.SkippedManualReview()),
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
func (summaryDetails *SummaryDetails) ListResourcesIDs() *helpersv1.AllLists {
	return summaryDetails.Controls.ListResourcesIDs()
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
	for _, framework := range summaryDetails.Frameworks {
		updateControlsSummaryCounters(resourceResult, framework.Controls, &helpersv1.Filters{FrameworkNames: []string{framework.Name}})
		framework.CalculateStatus()
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
			status := resourceResult.AssociatedControls[i].GetSubStatus()
			if status.Status() == apis.StatusUnknown {
				status = resourceResult.AssociatedControls[i].GetStatus(f)
			}
			controlSummary.Append(status, resourceResult.ResourceID)
			controlSummary.CalculateStatus()
			controls[controlID] = controlSummary
		}
	}
}
