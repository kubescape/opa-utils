package reportsummary

import (
	"strings"

	"github.com/kubescape/opa-utils/reporthandling/apis"
	helpersv1 "github.com/kubescape/opa-utils/reporthandling/helpers/v1"
)

// =================================== Status ============================================

// GetStatus get the control status. returns an apis.ScanningStatus object
func (controlSummary *ControlSummary) GetStatus() apis.IStatus {
	// Backward compability
	if controlSummary.StatusInfo.Status() == apis.StatusUnknown {
		controlSummary.StatusInfo.InnerStatus = controlSummary.Status
	}
	return &controlSummary.StatusInfo
}

func (controlSummary *ControlSummary) SetStatus(statusInfo *apis.StatusInfo) {
	if statusInfo == nil || statusInfo.Status() == apis.StatusUnknown {
		controlSummary.CalculateStatus()
	} else {
		controlSummary.StatusInfo = *statusInfo
		controlSummary.Status = statusInfo.Status()
	}
}

// CalculateStatus set the control status based on the resource counters
func (controlSummary *ControlSummary) CalculateStatus() {
	controlSummary.StatusInfo.InnerStatus = calculateStatus(&controlSummary.ResourceCounters)
	// Statuses should be the same
	controlSummary.Status = controlSummary.StatusInfo.Status()
}

// =================================== Counters ============================================
func (controlSummary *ControlSummary) ListResourcesIDs() *helpersv1.AllLists {
	return &controlSummary.ResourceIDs
}

// NumberOf get the number of resources
func (controlSummary *ControlSummary) NumberOfResources() ICounters {
	return &controlSummary.ResourceCounters
}

// Increase increases the counter based on the status
func (controlSummary *ControlSummary) increase(status apis.IStatus) {
	controlSummary.ResourceCounters.Increase(status)
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

// GetScoreFactor return control score
func (controlSummary *ControlSummary) GetScoreFactor() float32 {
	return controlSummary.ScoreFactor
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

// GetRemediation get control remediation
func (controlSummary *ControlSummary) GetRemediation() string {
	return controlSummary.Remediation
}

// GetDescription get control description
func (controlSummary *ControlSummary) GetDescription() string {
	return controlSummary.Description

}

//=============== ControlSummaries
func (controlSummaries *ControlSummaries) GetIDs() []string {
	keys := make([]string, 0, len((*controlSummaries)))
	for k := range *controlSummaries {
		keys = append(keys, k)
	}
	return keys
}

// get control either by criteria = "ID" and value <controlID> or criteria = name and <controlName>
func (controlSummaries *ControlSummaries) GetControl(criteria ControlCriteria, value string) IControlSummary {
	switch criteria {
	case EControlCriteriaID:
		tmp, ok := (*controlSummaries)[value]
		//avoid handling empty objects
		if !ok {
			return nil
		}
		return &tmp
	case EControlCriteriaName:
		for ctrlID := range *controlSummaries {
			if strings.Contains((*controlSummaries)[ctrlID].Name, value) {
				tmp := (*controlSummaries)[ctrlID]
				return &tmp
			}
		}
	}

	return nil
}

func (controlSummaries *ControlSummaries) ListControlsIDs() *helpersv1.AllLists {
	controls := &helpersv1.AllLists{}
	for controlID, controlSummary := range *controlSummaries {
		controls.Append(controlSummary.GetStatus().Status(), controlID)
	}
	return controls
}

//might be redundant
func (controlSummaries *ControlSummaries) NumberOfControls() ICounters {

	return &PostureCounters{
		PassedCounter:   len(controlSummaries.ListControlsIDs().Passed()),
		FailedCounter:   len(controlSummaries.ListControlsIDs().Failed()),
		ExcludedCounter: len(controlSummaries.ListControlsIDs().Excluded()),
		SkippedCounter:  len(controlSummaries.ListControlsIDs().Skipped()),
		UnknownCounter:  len(controlSummaries.ListControlsIDs().Other()),
	}
}

func (controlSummaries *ControlSummaries) ListResourcesIDs() *helpersv1.AllLists {
	allList := &helpersv1.AllLists{}

	//I've implemented it like this because i wanted to support future changes and access things only via interfaces(Lior)
	ctrlIDs := controlSummaries.ListControlsIDs().All()
	for _, ctrlID := range ctrlIDs {
		resourcesIDs := controlSummaries.GetControl(EControlCriteriaID, ctrlID).ListResourcesIDs()
		allList.Append(apis.StatusFailed, resourcesIDs.Failed()...)
		allList.Append(apis.StatusExcluded, resourcesIDs.Excluded()...)
		allList.Append(apis.StatusPassed, resourcesIDs.Passed()...)
		allList.Append(apis.StatusSkipped, resourcesIDs.Skipped()...)
		allList.Append(apis.StatusUnknown, resourcesIDs.Other()...)
	}

	// initialize resources IDs
	allList.ToUnique()

	return allList
}
