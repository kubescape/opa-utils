package v2

import (
	"time"

	ik8s "github.com/armosec/k8s-interface/workloadinterface"

	"github.com/armosec/armoapi-go/armotypes"
	"k8s.io/apimachinery/pkg/version"
)

// PostureReport posture scanning report structure
type PostureReport struct {
	CustomerGUID         string         `json:"customerGUID"`
	ClusterName          string         `json:"clusterName"`
	ClusterAPIServerInfo *version.Info  `json:"clusterAPIServerInfo"`
	ClusterCloudProvider string         `json:"clusterCloudProvider"`
	ReportID             string         `json:"reportID"`
	JobID                string         `json:"jobID"`
	ReportGenerationTime time.Time      `json:"generationTime"`
	SummaryDetails       SummaryDetails `json:"summaryDetails,omitempty"` // Developing
	Results              []Result       `json:"results,omitempty"`        // Developing
	Resources            []Resource     `json:"resource,omitempty"`
}

// SummaryDetails detailed summary of the scanning. will contain versions, counters, etc.
type SummaryDetails struct {
	Frameworks []FrameworkSummary        `json:"frameworks"`         // list of framework summary
	Controls   map[string]ControlSummary `json:"controls,omitempty"` // mapping of control - map[<control ID>]<control summary>
}

// FrameworkSummary summary of scanning from a single framework perspective
type FrameworkSummary struct {
	Score    float32                   `json:"score"`              // framework score
	Name     string                    `json:"name"`               // framework name
	Version  string                    `json:"version"`            // framework version
	Controls map[string]ControlSummary `json:"controls,omitempty"` // mapping of control - map[<control ID>]<control summary>
}

// FrameworkSummary summary of scanning from a single control perspective
type ControlSummary struct {
	Score            float32 `json:"score"`
	PassedResources  int     `json:"passedResources"`
	FailedResources  int     `json:"failedResources"`
	WarningResources int     `json:"warningResources"`
	SkippedResources int     `json:"skippedResources"`
	Status           string  `json:"status"`
}

// Result - resource result resourceID and the controls that where tested against the resource
type Result struct {
	ResourceID         string
	AssociatedControls []ResourceAssociatedControl
}

// ResourceAssociatedControl control that is associated to a resource
type ResourceAssociatedControl struct {
	ControlID               string
	ResourceAssociatedRules []ResourceAssociatedRule
}

// ResourceAssociatedRule failed rule that is associated to a resource
type ResourceAssociatedRule struct {
	RuleName    string                            `json:"ruleName"`
	FailedPaths []string                          `json:"failedPaths"`
	Status      string                            // "failed/passed/warning  ?skipped?"
	Exception   *armotypes.PostureExceptionPolicy `json:"exception,omitempty"`
	// TODO - add list of controls inputs
}

// Resource single resource representation from resource inventory
type Resource struct {
	ResourceID string         `json:"resourceID"`
	Object     interface{}    `json:"object"`
	IMetadata  ik8s.IMetadata `json:"-"`
}
