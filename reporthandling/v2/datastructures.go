package v2

import (
	"time"

	ik8s "github.com/armosec/k8s-interface/workloadinterface"
	"github.com/armosec/opa-utils/reporthandling/results/v1/reportsummary"
	"github.com/armosec/opa-utils/reporthandling/results/v1/resourcesresults"

	"k8s.io/apimachinery/pkg/version"
)

type PostureAttributes struct {
	Attribute string   `json:"attributeName"`
	Values    []string `json:"values"`
}

// PostureReport posture scanning report structure
type PostureReport struct {
	Attributes           []PostureAttributes          `json:"attributes"` //allow flexible properties for posture reports
	CustomerGUID         string                       `json:"customerGUID"`
	ClusterName          string                       `json:"clusterName"`
	ClusterCloudProvider string                       `json:"clusterCloudProvider"`
	ReportID             string                       `json:"reportGUID"`
	JobID                string                       `json:"jobID"`
	ClusterAPIServerInfo *version.Info                `json:"clusterAPIServerInfo"`
	ReportGenerationTime time.Time                    `json:"generationTime"`
	SummaryDetails       reportsummary.SummaryDetails `json:"summaryDetails,omitempty"` // Developing
	Results              []resourcesresults.Result    `json:"results,omitempty"`        // Developing
	Resources            []Resource                   `json:"resources,omitempty"`
}

// Resource single resource representation from resource inventory
type Resource struct {
	ResourceID string         `json:"resourceID"`
	Object     interface{}    `json:"object"`
	IMetadata  ik8s.IMetadata `json:"-"`
}
