package reporthandling

import (
	"time"

	ik8s "github.com/kubescape/k8s-interface/workloadinterface"

	"github.com/armosec/armoapi-go/armotypes"
	rbacutils "github.com/kubescape/rbac-utils/rbacutils"
	"k8s.io/apimachinery/pkg/version"
)

type AlertScore float32

const (
	StatusPassed  string = "success"
	StatusWarning string = "warning"
	StatusIgnore  string = "ignore"
	StatusFailed  string = "failed"
)

// RegoResponse the expected response of single run of rego policy
type RuleResponse struct {
	AlertMessage string                            `json:"alertMessage"`
	FailedPaths  []string                          `json:"failedPaths"`          // path in yaml that led to failure of this resource
	FixPaths     []armotypes.FixPath               `json:"fixPaths"`             // path in yaml to be added to fix this resource
	FixCommand   string                            `json:"fixCommand,omitempty"` // command to fix this resource
	RuleStatus   string                            `json:"ruleStatus"`
	PackageName  string                            `json:"packagename"`
	AlertScore   AlertScore                        `json:"alertScore"`
	AlertObject  AlertObject                       `json:"alertObject"`
	Context      []string                          `json:"context,omitempty"`  // TODO - Remove
	Rulename     string                            `json:"rulename,omitempty"` // TODO - Remove
	Exception    *armotypes.PostureExceptionPolicy `json:"exception,omitempty"`
}

type AlertObject struct {
	K8SApiObjects   []map[string]interface{} `json:"k8sApiObjects,omitempty"`
	ExternalObjects map[string]interface{}   `json:"externalObjects,omitempty"`
}

type ResourceUniqueCounter struct {
	TotalResources   int `json:"totalResources"`
	FailedResources  int `json:"failedResources"`
	WarningResources int `json:"warningResources"`
}

type FrameworkReport struct {
	Name                  string          `json:"name"`
	ControlReports        []ControlReport `json:"controlReports"`
	Score                 float32         `json:"score,omitempty"`
	ARMOImprovement       float32         `json:"ARMOImprovement,omitempty"`
	WCSScore              float32         `json:"wcsScore,omitempty"`
	ResourceUniqueCounter `json:",inline"`
}
type ControlReport struct {
	armotypes.PortalBase  `json:",inline"`
	Control_ID            string       `json:"id,omitempty"` // to be Deprecated
	ControlID             string       `json:"controlID"`
	Name                  string       `json:"name"`
	RuleReports           []RuleReport `json:"ruleReports"`
	Remediation           string       `json:"remediation"`
	Description           string       `json:"description"`
	Score                 float32      `json:"score"`
	BaseScore             float32      `json:"baseScore,omitempty"`
	ARMOImprovement       float32      `json:"ARMOImprovement,omitempty"`
	ResourceUniqueCounter `json:",inline"`
}
type RuleReport struct {
	Name                  string         `json:"name"`
	Remediation           string         `json:"remediation"`
	RuleStatus            RuleStatus     `json:"ruleStatus"` // did we run the rule or not (if there where compile errors, the value will be failed)
	RuleResponses         []RuleResponse `json:"ruleResponses"`
	ListInputKinds        []string       `json:"listInputIDs"`
	ResourceUniqueCounter `json:",inline"`
}
type RuleStatus struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// PostureReport
type PostureReport struct {
	CustomerGUID         string                `json:"customerGUID"`
	ClusterName          string                `json:"clusterName"`
	ClusterAPIServerInfo *version.Info         `json:"clusterAPIServerInfo"`
	ClusterCloudProvider string                `json:"clusterCloudProvider"`
	ReportID             string                `json:"reportID"`
	JobID                string                `json:"jobID"`
	ReportGenerationTime time.Time             `json:"generationTime"`
	FrameworkReports     []FrameworkReport     `json:"frameworks"`            // DEPRECATED
	RBACObjects          rbacutils.RbacObjects `json:"rbacObjects,omitempty"` // all rbac objects in cluster - roles, clusterroles, rolebindings, clusterrolebindings
	Resources            []Resource            `json:"resource,omitempty"`
}

const (
	SourceTypeJson               string = "JSON"
	SourceTypeYaml               string = "YAML"
	SourceTypeHelmChart          string = "Helm Chart"
	SourceTypeKustomizeDirectory string = "Kustomize Directory"
)

// Source - File source metadata
type Source struct {
	Path                   string     `json:"path,omitempty"`                   // deprecated
	RelativePath           string     `json:"relativePath,omitempty"`           // relative path from the repo base
	FileType               string     `json:"fileType,omitempty"`               // file type
	HelmChartName          string     `json:"helmChartName,omitempty"`          // helm chart name (if FileType is "Helm Chart")
	KustomizeDirectoryName string     `json:"kustomizeDirectoryName,omitempty"` //Kustomize Directory name if File is from Kustomize Directory
	LastCommit             LastCommit `json:"lastCommit,omitempty"`
}

// LastCommit - file in git repo last commit metadata
type LastCommit struct {
	Hash           string    `json:"hash,omitempty"`           // commit hash
	Date           time.Time `json:"date,omitempty"`           // commit date
	CommitterName  string    `json:"committerName,omitempty"`  // committer name
	CommitterEmail string    `json:"committerEmail,omitempty"` // committer email
	Message        string    `json:"message,omitempty"`        // commit message
}

type Resource struct {
	ResourceID string         `json:"resourceID"`
	Object     interface{}    `json:"object"`
	IMetadata  ik8s.IMetadata `json:"-"`
	Source     *Source        `json:"source,omitempty"`
}
