package v1

import v1 "github.com/armosec/opa-utils/httpserver/apis/v1"

// A request to trigger a Kubescape scan
type PostScanRequest struct {
	// Logger level (debug / info / error, default is "debug")
	Logger string `json:"-"`
	// Format of the results.
	//
	// Same as `kubescape scan --format`.
	//
	// Example: json
	Format string `json:"format,omitempty"`
	// A Kubescape account ID to use for scanning.
	//
	// Same as `kubescape scan --account`.
	//
	// Example: d13791eb-19b1-4222-867b-9a7c1799cfac
	// swagger:strfmt uuid4
	//
	Account string `json:"account,omitempty"`
	// Threshold for a failing score.
	//
	// Scores higher than the provided value will be considered failing.
	//
	// Example: 42
	FailThreshold float32 `json:"failThreshold,omitempty"`
	// Namespaces to exclude.
	//
	// Same as `kubescape scan --excluded-namespaces`.
	//
	// Example: ["armo-system", "kube-system"]
	ExcludedNamespaces []string `json:"excludedNamespaces,omitempty"`
	// Namespaces to include.
	//
	// Same as `kubescape scan --include-namespaces`.
	//
	// Example: ["litmus-tests", "known-bad"]
	IncludeNamespaces []string `json:"includeNamespaces,omitempty"`
	// Name of the scan targets.
	//
	// For example, if you select `targetType: "framework"`, you can trigger a scan using the NSA and MITRE ATT&CK Framework by passing `targetNames: ["nsa", "mitre"].
	//
	// Example: ["nsa", "mitre"]
	// Default: ["all"]
	TargetNames []string `json:"targetNames,omitempty"`
	// Type of the target. "framework" or "control".
	//
	// Example: "control"
	// Default: "framework"
	TargetType v1.NotificationPolicyKind `json:"targetType,omitempty"`
	// Submit results to Kubescape Cloud.
	//
	// Same as `kubescape scan --submit`.
	Submit *bool `json:"submit,omitempty"`
	// Deploy the Kubescape Kubernetes Host Scanner
	//
	// Deploys the Armo K8s Host Scanner DeamonSet in the scanned cluster to collect data from certain controls.
	//
	// Example: true
	HostScanner *bool `json:"hostScanner,omitempty"`
	// Do not submit results to Kubescape Cloud.
	//
	// Same as `kubescape scan --keep-local`
	//
	// Example: true
	KeepLocal *bool `json:"keepLocal,omitempty"`
	// Use the cached artifacts instead of downloading (offline support)
	//
	// Example: false
	UseCachedArtifacts *bool `json:"useCachedArtifacts,omitempty"`
	// UseExceptions      string      // Load file with exceptions configuration
	// ControlsInputs     string      // Load file with inputs for controls
	// VerboseMode        bool        // Display all of the input resources and not only failed resources
}

// A Scan Response object
type Response struct {
	// ID of the scan
	//
	// Example: d13791eb-19b1-4222-867b-9a7c1799cfac
	//
	// swagger:strfmt uuid4
	ID string `json:"id"`
	// Type of this response
	//
	// Example: busy
	Type v1.ScanResponseType `json:"type"`
	// The actual Response payload
	//
	// Example: d13791eb-19b1-4222-867b-9a7c1799cfac
	Response interface{} `json:"response,omitempty"`
}
