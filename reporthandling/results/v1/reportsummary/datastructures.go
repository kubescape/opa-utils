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
	Controls         ControlSummaries    `json:"controls,omitempty"`
	Status           apis.ScanningStatus `json:"status"`
	Frameworks       []FrameworkSummary  `json:"frameworks"`
	SeverityCounters SeverityCounters    `json:"severityCounters,omitempty"`
	ResourceCounters ResourceCounters    `json:",inline"`
	Score            float32             `json:"score"`
}

// FrameworkSummary summary of scanning from a single framework perspective
type FrameworkSummary struct {
	Controls         ControlSummaries    `json:"controls,omitempty"` // mapping of control - map[<control ID>]<control summary>
	Name             string              `json:"name"`               // framework name
	Status           apis.ScanningStatus `json:"status"`
	Version          string              `json:"version"`
	ResourceCounters ResourceCounters    `json:",inline"`
	Score            float32             `json:"score"`
}

// ControlSummary summary of scanning from a single control perspective
type ControlSummary struct {
	StatusInfo       apis.StatusInfo     `json:"statusInfo,omitempty"`
	ControlID        string              `json:"controlID"`
	Name             string              `json:"name"`
	Status           apis.ScanningStatus `json:"status"`
	Description      string              `json:"-"`
	Remediation      string              `json:"-"`
	ResourceIDs      helpersv1.AllLists  `json:"resourceIDs"`
	ResourceCounters ResourceCounters    `json:",inline"`
	Score            float32             `json:"score"`
	ScoreFactor      float32             `json:"scoreFactor"`
}

type ResourceCounters struct {
	PassedResources   int `json:"passedResources"`
	FailedResources   int `json:"failedResources"`
	ExcludedResources int `json:"excludedResources"`
}

type SeverityCounters struct {
	ResourcesWithCriticalSeverityCounter int `json:"criticalSeverity"`
	ResourcesWithHighSeverityCounter     int `json:"highSeverity"`
	ResourcesWithMediumSeverityCounter   int `json:"mediumSeverity"`
	ResourcesWithLowSeverityCounter      int `json:"lowSeverity"`
}

type PostureCounters struct {
	PassedCounter   int `json:"passed"`
	FailedCounter   int `json:"failed"`
	ExcludedCounter int `json:"excluded"`
	SkippedCounter  int `json:"skipped"`
	IgnoredCounter  int `json:"ignored"`
	UnknownCounter  int `json:"unknown"`
}
