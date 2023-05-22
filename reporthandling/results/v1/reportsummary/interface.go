package reportsummary

import (
	"time"

	"github.com/kubescape/opa-utils/reporthandling/apis"
	helpersv1 "github.com/kubescape/opa-utils/reporthandling/helpers/v1"
)

type ControlCriteria string

const (
	EControlCriteriaID   ControlCriteria = "ID"
	EControlCriteriaName ControlCriteria = "name" // DEPRECATED
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
	GetComplianceScore() float32
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

	// Get SubStatus() get control sub status
	GetSubStatus() apis.ScanningSubStatus

	StatusesCounters() (ICounters, ISubCounters)
}

type IControlsSummaries interface {
	GetControl(criteria ControlCriteria, value string) IControlSummary

	NumberOfControls() ICounters
	ListControlsIDs(*helpersv1.AllLists) *helpersv1.AllLists
	ListResourcesIDs(*helpersv1.AllLists) *helpersv1.AllLists
}

type IPolicies interface {
	GetStatus() apis.IStatus
	CalculateStatus()
	ListResourcesIDs(*helpersv1.AllLists) *helpersv1.AllLists

	// Counters
	NumberOfResources() ICounters

	// Score
	GetScore() float32

	// ComplianceScore
	GetComplianceScore() float32

	// Name
	GetName() string
}
