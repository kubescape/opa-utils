package reportsummary

import (
	"github.com/kubescape/opa-utils/reporthandling/apis"
	helpersv1 "github.com/kubescape/opa-utils/reporthandling/helpers/v1"
)

type PostureAttributes struct {
	Attribute string   `json:"attributeName"`
	Values    []string `json:"values"`
}

type ControlSummaries map[string]ControlSummary

// SummaryDetails detailed summary of the scanning. will contain versions, counters, etc.
type SummaryDetails struct {
	Controls                  ControlSummaries    `json:"controls,omitempty"`
	Status                    apis.ScanningStatus `json:"status"`
	Frameworks                []FrameworkSummary  `json:"frameworks"`
	ResourcesSeverityCounters SeverityCounters    `json:"resourcesSeverityCounters,omitempty"`
	ControlsSeverityCounters  SeverityCounters    `json:"controlsSeverityCounters,omitempty"`
	StatusCounters            StatusCounters      `json:"ResourceCounters"` // Backward compatibility
	Score                     float32             `json:"score"`
}

// FrameworkSummary summary of scanning from a single framework perspective
type FrameworkSummary struct {
	Controls        ControlSummaries    `json:"controls,omitempty"` // mapping of control - map[<control ID>]<control summary>
	Name            string              `json:"name"`               // framework name
	Status          apis.ScanningStatus `json:"status"`
	Version         string              `json:"version"`
	StatusCounters  StatusCounters      `json:"ResourceCounters"` // Backward compatibility
	Score           float32             `json:"score"`
	ComplianceScore float32             `json:"complianceScore"`
}

// ControlSummary summary of scanning from a single control perspective
type ControlSummary struct {
	StatusInfo        apis.StatusInfo     `json:"statusInfo,omitempty"`
	ControlID         string              `json:"controlID"`
	Name              string              `json:"name"`
	Status            apis.ScanningStatus `json:"status"` // backward compatibility
	Description       string              `json:"-"`
	Remediation       string              `json:"-"`
	ResourceIDs       helpersv1.AllLists  `json:"resourceIDs"`
	StatusCounters    StatusCounters      `json:"ResourceCounters"` // Backward compatibility
	SubStatusCounters SubStatusCounters   `json:"subStatusCounters"`
	Score             float32             `json:"score"`
	ScoreFactor       float32             `json:"scoreFactor"`
}

type StatusCounters struct {
	PassedResources   int `json:"passedResources"`
	FailedResources   int `json:"failedResources"`
	SkippedResources  int `json:"skippedResources"`
	ExcludedResources int `json:"excludedResources"` // Deprecated
}

type SubStatusCounters struct {
	IgnoredResources int `json:"ignoredResources"`
}

type SeverityCounters struct {
	CriticalSeverityCounter int `json:"criticalSeverity"`
	HighSeverityCounter     int `json:"highSeverity"`
	MediumSeverityCounter   int `json:"mediumSeverity"`
	LowSeverityCounter      int `json:"lowSeverity"`
}

type PostureCounters struct {
	PassedCounter   int `json:"passed"`
	FailedCounter   int `json:"failed"`
	SkippedCounter  int `json:"skipped"`
	ExcludedCounter int `json:"excluded"` // Deprecated
}
