package v2

import (
	"time"

	armoapi "github.com/armosec/armoapi-go/apis"
	"github.com/kubescape/opa-utils/reporthandling"
	"github.com/kubescape/opa-utils/reporthandling/apis"
	"github.com/kubescape/opa-utils/reporthandling/results/v1/reportsummary"
	"github.com/kubescape/opa-utils/reporthandling/results/v1/resourcesresults"

	"k8s.io/apimachinery/pkg/version"
)

// PostureReport posture scanning report structure
type PostureReport struct {
	ReportGenerationTime  time.Time                         `json:"generationTime"`
	ClusterAPIServerInfo  *version.Info                     `json:"clusterAPIServerInfo"`
	ClusterCloudProvider  string                            `json:"clusterCloudProvider"`
	CustomerGUID          string                            `json:"customerGUID"`
	ClusterName           string                            `json:"clusterName"`
	ReportID              string                            `json:"reportGUID"`
	JobID                 string                            `json:"jobID"`
	SummaryDetails        reportsummary.SummaryDetails      `json:"summaryDetails,omitempty"`
	Resources             []reporthandling.Resource         `json:"resources,omitempty"`
	Attributes            []reportsummary.PostureAttributes `json:"attributes"`
	Results               []resourcesresults.Result         `json:"results,omitempty"`
	Metadata              Metadata                          `json:"metadata,omitempty"`
	PaginationInfo        armoapi.PaginationMarks           `json:"paginationInfo"`
	CustomerGUIDGenerated bool                              `json:"customerGUIDGenerated"`
}

type ClusterMetadata struct {
	MapNamespaceToNumberOfResources map[string]int `json:"namespaceToNumberOfResources,omitempty"`
	CloudMetadata                   *CloudMetadata `json:"cloudMetadata,omitempty"`
	CloudProvider                   string         `json:"cloudProvider,omitempty"` // Deprecated - info should be in cloudMetadata
	ContextName                     string         `json:"contextName,omitempty"`
	NumberOfWorkerNodes             int            `json:"numberOfWorkerNodes,omitempty"`
}

// CloudMetadata metadata of the cloud the cluster is running on. Compatible with the reporthandling.ICloudMetadata interface
type CloudMetadata struct {
	CloudProvider apis.CloudProviderName `json:"cloudProvider,omitempty"`
	ShortName     string                 `json:"shortName,omitempty"`
	FullName      string                 `json:"fullName,omitempty"`
	PrefixName    string                 `json:"prefixName,omitempty"`
}

type RepoContextMetadata struct {
	Provider      string                    `json:"provider,omitempty"` // repo provider name. e.g. github, gitlab
	Repo          string                    `json:"repo,omitempty"`
	Owner         string                    `json:"owner,omitempty"`
	Branch        string                    `json:"branch,omitempty"`
	DefaultBranch string                    `json:"defaultBranch,omitempty"`
	RemoteURL     string                    `json:"remoteURL,omitempty"`
	LastCommit    reporthandling.LastCommit `json:"lastCommit,omitempty"`
	LocalRootPath string                    `json:"localRootPath,omitempty"` // repo root path (local)
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
	ContextMetadata ContextMetadata `json:"targetMetadata,omitempty"`
	ClusterMetadata ClusterMetadata `json:"clusterMetadata,omitempty"`
	ScanMetadata    ScanMetadata    `json:"scanMetadata,omitempty"`
}

type ScanningTarget uint16

const (
	Cluster   ScanningTarget = 0
	File      ScanningTarget = 1
	Repo      ScanningTarget = 2
	GitLocal  ScanningTarget = 3
	Directory ScanningTarget = 4
)

type ScanMetadata struct {
	TargetType       string `json:"targetType,omitempty"`
	KubescapeVersion string `json:"kubescapeVersion,omitempty"`
	FormatVersion    string `json:"formatVersion,omitempty"`
	ControlsInputs   string `json:"controlsInputs,omitempty"`
	// Format that has been requested for the output results.
	//
	// Since Kubescape added support for multiple outputs, might be not a
	// single format, but a comma-separated string of the multiple
	// requested formats.
	//
	// Deprecated: Since Kubescape added support for multiple outputs,
	// `Format` exists only for backward compatibility. Please use the
	// `Formats` field instead.
	Format string `json:"format,omitempty"`
	// Formats that have been requested for the output results.
	Formats             []string       `json:"formats,omitempty"`
	UseExceptions       string         `json:"useExceptions,omitempty"`
	Logger              string         `json:"logger,omitempty"`
	ExcludedNamespaces  []string       `json:"excludedNamespaces,omitempty"`
	IncludeNamespaces   []string       `json:"includeNamespaces,omitempty"`
	TargetNames         []string       `json:"targetNames,omitempty"`
	FailThreshold       float32        `json:"failThreshold,omitempty"`
	ComplianceThreshold float32        `json:"complianceThreshold,omitempty"`
	ScanningTarget      ScanningTarget `json:"scanningTarget,omitempty"`
	HostScanner         bool           `json:"hostScanner,omitempty"`
	Submit              bool           `json:"submit,omitempty"`
	VerboseMode         bool           `json:"verboseMode,omitempty"`
}

// Moved to apis/cloudmetadata.go
// const (
// 	GKE = "GKE"
// 	GCP = "GCP"
// 	EKS = "EKS"
// )

func (st *ScanningTarget) String() string {
	switch *st {
	case 0:
		return "Cluster"
	case 1:
		return "File"
	case 2:
		return "Repo"
	case 3:
		return "GitLocal"
	case 4:
		return "Directory"
	default:
		return ""
	}
}
