package reportsummary

import (
	"strings"

	"github.com/kubescape/opa-utils/reporthandling/apis"
	helpersv1 "github.com/kubescape/opa-utils/reporthandling/helpers/v1"
)

// =================================== Status ============================================

// GetStatus get the control status. returns an apis.StatusInfo object
func (controlSummary *ControlSummary) GetStatus() apis.IStatus {
	// Backward compatibility
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

func (controlSummary *ControlSummary) SetSubStatus(subStatus apis.ScanningSubStatus) {
	controlSummary.StatusInfo.SubStatus = subStatus
}

// GetSubStatus get the control sub status. returns an apis.StatusInfo object
func (controlSummary *ControlSummary) GetSubStatus() apis.ScanningSubStatus {
	return controlSummary.StatusInfo.SubStatus
}

func (controlSummary *ControlSummary) CalculateStatus() {
	controlSummary.calculateStatus(apis.SubStatusUnknown)
}

// calculateStatus set the control status based on the resource counters and the sub status based on the subStatus parameter
func (controlSummary *ControlSummary) calculateStatus(subStatus apis.ScanningSubStatus) {
	controlSummary.StatusInfo.InnerStatus = calculateStatus(&controlSummary.StatusCounters)

	// Statuses should be the same
	controlSummary.Status = controlSummary.StatusInfo.Status() // backward compatibility

	controlSummary.calculateNSetSubStatus(subStatus)
}

// CalculateSubStatus set the control sub status based on the resource associated control sub status
func (controlSummary *ControlSummary) calculateNSetSubStatus(subStatus apis.ScanningSubStatus) {
	switch controlSummary.Status {
	case apis.StatusPassed:
		if subStatus == apis.SubStatusIrrelevant || controlSummary.StatusInfo.SubStatus == apis.SubStatusIrrelevant || controlSummary.StatusCounters.All() == 0 {
			controlSummary.StatusInfo.SubStatus = apis.SubStatusIrrelevant
			controlSummary.StatusInfo.InnerInfo = ""
		} else if subStatus == apis.SubStatusException || controlSummary.StatusInfo.SubStatus == apis.SubStatusException {
			controlSummary.StatusInfo.SubStatus = apis.SubStatusException
			controlSummary.StatusInfo.InnerInfo = ""
		}
	case apis.StatusSkipped:
		if subStatus == apis.SubStatusConfiguration || controlSummary.StatusInfo.SubStatus == apis.SubStatusConfiguration {
			controlSummary.StatusInfo.SubStatus = apis.SubStatusConfiguration
			controlSummary.StatusInfo.InnerInfo = string(apis.SubStatusConfigurationInfo)
		} else if subStatus == apis.SubStatusManualReview || controlSummary.StatusInfo.SubStatus == apis.SubStatusManualReview {
			controlSummary.StatusInfo.SubStatus = apis.SubStatusManualReview
			controlSummary.StatusInfo.InnerInfo = string(apis.SubStatusManualReviewInfo)
		} else if subStatus == apis.SubStatusRequiresReview || controlSummary.StatusInfo.SubStatus == apis.SubStatusRequiresReview {
			controlSummary.StatusInfo.SubStatus = apis.SubStatusRequiresReview
			controlSummary.StatusInfo.InnerInfo = string(apis.SubStatusRequiresReviewInfo)
		}
	case apis.StatusFailed:
		controlSummary.StatusInfo.SubStatus = apis.SubStatusUnknown
		controlSummary.StatusInfo.InnerInfo = ""
	}
}

// =================================== Counters ============================================
func (controlSummary *ControlSummary) ListResourcesIDs() *helpersv1.AllLists {
	return &controlSummary.ResourceIDs
}

// Deprecated use 'ResourcesCounters' instead
// NumberOfResources get the status counters
func (controlSummary *ControlSummary) NumberOfResources() ICounters {
	return &controlSummary.StatusCounters
}

// NumberOfResources get the status counters
func (controlSummary *ControlSummary) ResourcesCounters() (ICounters, ISubCounters) {
	return &controlSummary.StatusCounters, &controlSummary.SubStatusCounters
}

// Increase increases the counter based on the status
func (controlSummary *ControlSummary) increase(status apis.IStatus) {
	controlSummary.StatusCounters.Increase(status)
	controlSummary.SubStatusCounters.Increase(status)
}

// Append increases the counter based on the status
func (controlSummary *ControlSummary) Append(status apis.IStatus, ids ...string) {
	for i := range ids {
		controlSummary.ResourceIDs.Append(status.Status(), ids[i])
		controlSummary.increase(status)
	}
	controlSummary.ResourceIDs.ToUniqueResources() // TODO: check if it is needed
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

// =============== ControlSummaries
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

	// TODO: remove the section once confirmed all system are not using EControlCriteriaName
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
	controls.ToUniqueControls()
	return controls
}

// might be redundant
func (controlSummaries *ControlSummaries) NumberOfControls() ICounters {
	l := controlSummaries.ListControlsIDs()
	return &PostureCounters{
		PassedCounter:  len(l.Passed()),
		FailedCounter:  len(l.Failed()),
		SkippedCounter: len(l.Skipped()),
	}
}

func (controlSummaries *ControlSummaries) ListResourcesIDs() *helpersv1.AllLists {
	allList := &helpersv1.AllLists{}

	//I've implemented it like this because i wanted to support future changes and access things only via interfaces
	ctrlIDsIter := controlSummaries.ListControlsIDs().All()
	for ctrlIDsIter.HasNext() {
		resourcesIDs := controlSummaries.GetControl(EControlCriteriaID, ctrlIDsIter.Next()).ListResourcesIDs()
		allList.Append(apis.StatusFailed, resourcesIDs.Failed()...)
		allList.Append(apis.StatusPassed, resourcesIDs.Passed()...)
		allList.Append(apis.StatusSkipped, resourcesIDs.Skipped()...)
		allList.Append(apis.StatusUnknown, resourcesIDs.Other()...)
	}

	// remove resources IDs duplications
	allList.ToUniqueResources()

	return allList
}
