package v2

import (
	"time"

	"github.com/armosec/opa-utils/reporthandling"
	"github.com/armosec/opa-utils/reporthandling/results/v1/reportsummary"
	"github.com/armosec/opa-utils/reporthandling/results/v1/resourcesresults"

	"k8s.io/apimachinery/pkg/version"
)

type PaginationMarks struct {
	ReportNumber int  `json:"chunkNumber"` // serial number of report, used in pagination
	IsLastReport bool `json:"isLastChunk"` //specify this is the last report, used in pagination
}

// PostureReport posture scanning report structure
type PostureReport struct {
	Attributes           []reportsummary.PostureAttributes `json:"attributes"` //allow flexible properties for posture reports
	CustomerGUID         string                            `json:"customerGUID"`
	ClusterName          string                            `json:"clusterName"`
	ClusterCloudProvider string                            `json:"clusterCloudProvider"` // Deprecated
	ReportID             string                            `json:"reportGUID"`
	JobID                string                            `json:"jobID"`
	PaginationInfo       PaginationMarks                   `json:"paginationInfo"`
	ClusterAPIServerInfo *version.Info                     `json:"clusterAPIServerInfo"`
	ReportGenerationTime time.Time                         `json:"generationTime"`
	SummaryDetails       reportsummary.SummaryDetails      `json:"summaryDetails,omitempty"`
	Results              []resourcesresults.Result         `json:"results,omitempty"`
	Resources            []reporthandling.Resource         `json:"resources,omitempty"`
	Metadata             Metadata                          `json:"metadata,omitempty"`
}

type ClusterMetadata struct {
	NumberOfWorkerNodes int    `json:"numberOfWorkerNodes,omitempty"`
	CloudProvider       string `json:"cloudProvider,omitempty"`
	ContextName         string `json:"contextName,omitempty"`
}

type RepoContextMetadata struct {
	Repo   string `json:"repo,omitempty"`
	Owner  string `json:"owner,omitempty"`
	Branch string `json:"branch,omitempty"`
}

type FileContextMetadata struct {
	FilePath string `json:"filePath,omitempty"` // like "hostpath"
	HostName string `json:"hostName,omitempty"` // like "hostpath"
}
type DirectoryContextMetadata struct {
	BasePath string `json:"basePath,omitempty"` // the scanning request base path
	HostName string `json:"hostName,omitempty"` // like "hostpath"
}

type HelmContextMetadata struct {
	ChartName string `json:"chartName,omitempty"`
}
type ContextMetadata struct {
	ClusterContextMetadata   *ClusterMetadata          `json:"clusterContextMetadata,omitempty"`
	RepoContextMetadata      *RepoContextMetadata      `json:"gitRepoContextMetadata,omitempty"`
	FileContextMetadata      *FileContextMetadata      `json:"fileContextMetadata,omitempty"`
	HelmContextMetadata      *HelmContextMetadata      `json:"helmContextMetadata,omitempty"`
	DirectoryContextMetadata *DirectoryContextMetadata `json:"directoryContextMetadata,omitempty"`
}

type Metadata struct {
	ScanMetadata    ScanMetadata    `json:"scanMetadata,omitempty"`
	ContextMetadata ContextMetadata `json:"targetMetadata,omitempty"`
	ClusterMetadata ClusterMetadata `json:"clusterMetadata,omitempty"` // DEPRECATED
}

type ScanningTarget uint16

const (
	Cluster ScanningTarget = 0
	File    ScanningTarget = 1
)

type ScanMetadata struct {
	Format             string         `json:"format,omitempty"`             // Format results (table, json, junit ...)
	KubescapeVersion   string         `json:"kubescapeVersion,omitempty"`   // Kubescape version
	FormatVersion      string         `json:"formatVersion,omitempty"`      // Format version (v1, v2)
	ScanningTarget     ScanningTarget `json:"scanningTarget,omitempty"`     // scanning target
	ExcludedNamespaces []string       `json:"excludedNamespaces,omitempty"` // used for host sensor namespace
	IncludeNamespaces  []string       `json:"includeNamespaces,omitempty"`
	FailThreshold      float32        `json:"failThreshold,omitempty"`  // Failure score threshold
	Submit             bool           `json:"submit,omitempty"`         // Submit results to Armo BE
	HostScanner        bool           `json:"hostScanner,omitempty"`    // Deploy ARMO K8s host sensor to collect data from certain controls
	Logger             string         `json:"logger,omitempty"`         // logger level - debug/info/error
	TargetType         string         `json:"targetType,omitempty"`     // framework/control
	TargetNames        []string       `json:"targetNames,omitempty"`    // list of frameworks/controls
	UseExceptions      string         `json:"useExceptions,omitempty"`  // Load file with exceptions configuration
	ControlsInputs     string         `json:"controlsInputs,omitempty"` // Load file with inputs for controls
	VerboseMode        bool           `json:"verboseMode,omitempty"`    // Display all of the input resources and not only failed resources
}

const (
	GKE = "GKE"
	GCP = "GCP"
	EKS = "EKS"
)

func (st *ScanningTarget) String() string {
	switch *st {
	case 0:
		return "Cluster"
	case 1:
		return "File"
	default:
		return ""
	}
}
