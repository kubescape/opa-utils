package reportsummary

import (
	"time"

	"github.com/armosec/opa-utils/reporthandling/apis"
	helpersv1 "github.com/armosec/opa-utils/reporthandling/helpers/v1"
)

type ControlCriteria string

const (
	EControlCriteriaID   ControlCriteria = "ID"
	EControlCriteriaName ControlCriteria = "name"
)

type IBasicPostureReport interface {
	GetCustomerGUID() string
	GetClusterName() string
	GetReportGUID() string
	GetJobID() string
	GetTimestamp() *time.Time

	//todo GetISummaryDetails
}

type IFrameworkSummary interface {
	IPolicies
	ListControls() []IControlSummary
	NumberOfControls() ICounters
}

type IControlSummary interface {
	IPolicies
	GetScoreFactor() float32

	// GetID get control ID
	GetID() string

	// GetRemediation get control remediation
	GetRemediation() string

	// GetDescription get control description
	GetDescription() string
}

type IControlsSummaries interface {
	GetControl(criteria ControlCriteria, value string) IControlSummary

	NumberOfControls() ICounters
	ListControlsIDs() *helpersv1.AllLists
	ListResourcesIDs() *helpersv1.AllLists //avoid using this  outside of kubescape
}
type IPolicies interface {
	GetStatus() apis.IStatus
	CalculateStatus()
	ListResourcesIDs() *helpersv1.AllLists //avoid using this outside of kubescape

	// Counters
	NumberOfResources() ICounters

	// Score
	GetScore() float32

	// Name
	GetName() string
}

type ListPolicies struct {
	passed   []IPolicies
	excluded []IPolicies
	failed   []IPolicies
	skipped  []IPolicies
	other    []IPolicies
}

func (all *ListPolicies) Failed() []IPolicies   { return all.failed }
func (all *ListPolicies) Passed() []IPolicies   { return all.passed }
func (all *ListPolicies) Excluded() []IPolicies { return all.excluded }
func (all *ListPolicies) Skipped() []IPolicies  { return all.skipped }
func (all *ListPolicies) Other() []IPolicies    { return all.other }
func (all *ListPolicies) All() []IPolicies {
	l := []IPolicies{}
	l = append(l, all.failed...)
	l = append(l, all.excluded...)
	l = append(l, all.passed...)
	l = append(l, all.skipped...)
	l = append(l, all.other...)
	return l
}

// Append append single string to matching status list
func (all *ListPolicies) Append(status apis.ScanningStatus, policy IPolicies) {
	switch status {
	case apis.StatusPassed:
		all.passed = append(all.passed, policy)
	case apis.StatusFailed:
		all.failed = append(all.failed, policy)
	case apis.StatusExcluded:
		all.excluded = append(all.excluded, policy)
	case apis.StatusSkipped:
		all.skipped = append(all.skipped, policy)
	default:
		all.other = append(all.other, policy)
	}
}

// Update AllLists objects with
func (all *ListPolicies) Update(all2 *ListPolicies) {
	all.passed = append(all.passed, all2.passed...)
	all.failed = append(all.failed, all2.failed...)
	all.excluded = append(all.excluded, all2.excluded...)
	all.skipped = append(all.skipped, all2.skipped...)
	all.other = append(all.other, all2.other...)
}
