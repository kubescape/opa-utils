package reportsummary

import (
	"strings"

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

//===============ControlSummaries
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
		tmp := controlSummaries.GetControl(EControlCriteriaID, ctrlID).ListResourcesIDs()
		allList.Append(apis.StatusFailed, controlSummaries.GetControl(EControlCriteriaID, ctrlID).ListResourcesIDs().Failed()...)
		allList.Append(apis.StatusExcluded, controlSummaries.GetControl(EControlCriteriaID, ctrlID).ListResourcesIDs().Excluded()...)
		allList.Append(apis.StatusPassed, tmp.Passed()...)
		allList.Append(apis.StatusSkipped, controlSummaries.GetControl(EControlCriteriaID, ctrlID).ListResourcesIDs().Skipped()...)
		allList.Append(apis.StatusUnknown, controlSummaries.GetControl(EControlCriteriaID, ctrlID).ListResourcesIDs().Other()...)
	}

	// allList.ToUnique()
	return allList
}
