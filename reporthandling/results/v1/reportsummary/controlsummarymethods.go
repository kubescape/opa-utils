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

// GetSubStatus get the control sub status. returns an apis.StatusInfo object
func (controlSummary *ControlSummary) GetSubStatus() apis.IStatus {
	return &apis.StatusInfo{InnerStatus: controlSummary.SubStatus}
}

// CalculateStatus set the control status and subStatus based on the resource counters
func (controlSummary *ControlSummary) CalculateStatus() {
	controlSummary.StatusInfo.InnerStatus = calculateStatus(&controlSummary.ResourceCounters)
	// Statuses should be the same
	controlSummary.Status = controlSummary.StatusInfo.Status()

	controlSummary.SubStatus = calculateSubStatus(&controlSummary.ResourceCounters)
}

// CalculateSubStatus set the control sub status based on the resource associated control sub status
func (controlSummary *ControlSummary) CalculateSubStatus(subStatus apis.ScanningStatus) {
	switch controlSummary.Status {
	case apis.StatusPassed:
		if subStatus == apis.SubStatusIrrelevant || controlSummary.SubStatus == apis.SubStatusIrrelevant || controlSummary.ResourceCounters.All() == 0 {
			controlSummary.SubStatus = apis.SubStatusIrrelevant
		} else if subStatus == apis.SubStatusException || controlSummary.SubStatus == apis.SubStatusException {
			controlSummary.SubStatus = apis.SubStatusException
		}
	case apis.StatusSkipped:
		if subStatus == apis.SubStatusConfiguration || controlSummary.SubStatus == apis.SubStatusConfiguration {
			controlSummary.SubStatus = apis.SubStatusConfiguration
		} else if subStatus == apis.SubStatusManualReview || controlSummary.SubStatus == apis.SubStatusManualReview {
			controlSummary.SubStatus = apis.SubStatusManualReview
		} else if subStatus == apis.SubStatusRequiresReview || controlSummary.SubStatus == apis.SubStatusRequiresReview {
			controlSummary.SubStatus = apis.SubStatusRequiresReview
		}
	}
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
	controlSummary.ResourceIDs.ToUniqueResources()
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
		status := controlSummary.GetSubStatus().Status()
		if status == apis.StatusUnknown {
			status = controlSummary.GetStatus().Status()
		}
		controls.Append(status, controlID)
	}
	controls.ToUniqueControls()
	return controls
}

// might be redundant
func (controlSummaries *ControlSummaries) NumberOfControls() ICounters {
	l := controlSummaries.ListControlsIDs()
	return &PostureCounters{
		PassedCounter:                len(l.Passed()),
		PassedExceptionCounter:       len(l.PassedExceptions()),
		PassedIrrelevantCounter:      len(l.PassedIrrelevant()),
		FailedCounter:                len(l.Failed()),
		SkippedConfigurationCounter:  len(l.SkippedConfiguration()),
		SkippedManualReviewCounter:   len(l.SkippedManualReview()),
		SkippedRequiresReviewCounter: len(l.SkippedRequiresReview()),
		SkippedIntegrationCounter:    len(l.SkippedIntegration()),
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
		allList.Append(apis.SubStatusException, resourcesIDs.PassedExceptions()...)
		allList.Append(apis.SubStatusIrrelevant, resourcesIDs.PassedIrrelevant()...)
		allList.Append(apis.SubStatusConfiguration, resourcesIDs.SkippedConfiguration()...)
		allList.Append(apis.SubStatusManualReview, resourcesIDs.SkippedManualReview()...)
		allList.Append(apis.SubStatusRequiresReview, resourcesIDs.SkippedRequiresReview()...)
		allList.Append(apis.SubStatusIntegration, resourcesIDs.SkippedIntegration()...)
		allList.Append(apis.StatusUnknown, resourcesIDs.Other()...)
	}

	// remove resources IDs duplications
	allList.ToUniqueResources()

	return allList
}
