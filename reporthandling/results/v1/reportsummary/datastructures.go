package reportsummary

import (
	"github.com/armosec/opa-utils/reporthandling/apis"
	helpersv1 "github.com/armosec/opa-utils/reporthandling/helpers/v1"
)

// SummaryDetails detailed summary of the scanning. will contain versions, counters, etc.
type SummaryDetails struct {
	Score            float32             `json:"score"`              // overall score
	Status           apis.ScanningStatus `json:"status"`             // overall status
	Frameworks       []FrameworkSummary  `json:"frameworks"`         // list of framework summary
	Controls         ControlSummaries    `json:"controls,omitempty"` // mapping of control - map[<control ID>]<control summary>
	ResourceCounters ResourceCounters    `json:",inline"`
	resourceIDs      helpersv1.AllLists  `json:"-"`
}

// FrameworkSummary summary of scanning from a single framework perspective
type FrameworkSummary struct {
	Name             string              `json:"name"` // framework name
	Status           apis.ScanningStatus `json:"status"`
	Score            float32             `json:"score"`              // framework score
	Version          string              `json:"version"`            // framework version
	Controls         ControlSummaries    `json:"controls,omitempty"` // mapping of control - map[<control ID>]<control summary>
	ResourceCounters ResourceCounters    `json:",inline"`
	resourceIDs      helpersv1.AllLists  `json:"-"`
}

type ControlSummaries map[string]ControlSummary

// ControlSummary summary of scanning from a single control perspective
type ControlSummary struct {
	ControlID        string              `json:"controlID"`
	Name             string              `json:"name"`
	Status           apis.ScanningStatus `json:"status"`
	Score            float32             `json:"score"`
	ResourceCounters ResourceCounters    `json:",inline"`
	ResourceIDs      helpersv1.AllLists  `json:"resourceIDs"`
	Description      string              `json:"-"`
	Remediation      string              `json:"-"`
}

type ResourceCounters struct {
	PassedResources   int `json:"passedResources"`
	FailedResources   int `json:"failedResources"`
	ExcludedResources int `json:"excludedResources"`
	SkippedResources  int `json:"skippedResources"`
}

type PostureCounters struct {
	PassedCounter   int `json:"passed"`
	FailedCounter   int `json:"failed"`
	ExcludedCounter int `json:"excluded"`
	SkippedCounter  int `json:"skipped"`
	IgnoredCounter  int `json:"ignored"`
	UnknownCounter  int `json:"unknown"`
}
