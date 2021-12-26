package v2

import (
	"time"

	ik8s "github.com/armosec/k8s-interface/workloadinterface"
	"github.com/armosec/opa-utils/reporthandling/helpers/v1/reportsummary"

	"github.com/armosec/armoapi-go/armotypes"
	"k8s.io/apimachinery/pkg/version"
)

// PostureReport posture scanning report structure
type PostureReport struct {
	CustomerGUID         string                       `json:"customerGUID"`
	ClusterName          string                       `json:"clusterName"`
	ClusterAPIServerInfo *version.Info                `json:"clusterAPIServerInfo"`
	ClusterCloudProvider string                       `json:"clusterCloudProvider"`
	ReportID             string                       `json:"reportID"`
	JobID                string                       `json:"jobID"`
	ReportGenerationTime time.Time                    `json:"generationTime"`
	SummaryDetails       reportsummary.SummaryDetails `json:"summaryDetails,omitempty"` // Developing
	Results              []Result                     `json:"results,omitempty"`        // Developing
	Resources            []Resource                   `json:"resource,omitempty"`
}

// Result - resource result resourceID and the controls that where tested against the resource
type Result struct {
	ResourceID         string
	AssociatedControls []ResourceAssociatedControl
}

// ResourceAssociatedControl control that is associated to a resource
type ResourceAssociatedControl struct {
	ControlID string
	// TODO - add list of controls inputs
	ResourceAssociatedRules []ResourceAssociatedRule
}

// ResourceAssociatedRule failed rule that is associated to a resource
type ResourceAssociatedRule struct {
	RuleName    string                             `json:"ruleName"`
	FailedPaths []string                           `json:"failedPaths"`
	Exception   []armotypes.PostureExceptionPolicy `json:"exception,omitempty"`
}

// Resource single resource representation from resource inventory
type Resource struct {
	ResourceID string         `json:"resourceID"`
	Object     interface{}    `json:"object"`
	IMetadata  ik8s.IMetadata `json:"-"`
}
