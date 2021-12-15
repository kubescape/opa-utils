package v2

import (
	"time"

	ik8s "github.com/armosec/k8s-interface/workloadinterface"

	"github.com/armosec/armoapi-go/armotypes"
	"k8s.io/apimachinery/pkg/version"
)

type RuleStatus struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// PostureReport
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
type SummaryDetails struct {
	Frameworks []FrameworkSummary        `json:"frameworks"`
	Controls   map[string]ControlSummary `json:"controls"`
}
type FrameworkSummary struct {
	Score     float32 `json:"score"`
	Framework string  `json:"framework"`
}
type ControlSummary struct {
	Score            float32 `json:"score"`
	PassedResources  int     `json:"passedResources"`
	FailedResources  int     `json:"failedResources"`
	WarningResources int     `json:"warningResources"`
	SkippedResources int     `json:"skippedResources"`
	Status           string  `json:"status"`
}

type ResourceAssociatedRule struct {
	RuleName    string                            `json:"ruleName"`
	FailedPaths []string                          `json:"failedPaths"`
	Exception   *armotypes.PostureExceptionPolicy `json:"exception,omitempty"`
}
type ResourceAssociatedControl struct {
	ControlID               string
	ResourceAssociatedRules ResourceAssociatedRule
}
type Result struct {
	ResourceID         string
	Status             string // "failed/passed/warning  ?skipped?"
	AssociatedControls []ResourceAssociatedControl
}

type Resource struct {
	ResourceID string         `json:"resourceID"`
	Object     interface{}    `json:"object"`
	IMetadata  ik8s.IMetadata `json:"-"`
}
