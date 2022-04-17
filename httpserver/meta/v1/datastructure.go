package v1

import v1 "github.com/armosec/opa-utils/httpserver/apis/v1"

type PostScanRequest struct {
	Format             string                    `json:"format"`             // Format results (table, json, junit ...) - default json
	Account            string                    `json:"account"`            // account ID
	Logger             string                    `json:"-"`                  // logger level - debug/info/error - default is debug
	FailThreshold      float32                   `json:"failThreshold"`      // Failure score threshold
	ExcludedNamespaces []string                  `json:"excludedNamespaces"` // used for host scanner namespace
	IncludeNamespaces  []string                  `json:"includeNamespaces"`  // DEPRECATED?
	TargetNames        []string                  `json:"targetNames"`        // default is all
	TargetType         v1.NotificationPolicyKind `json:"targetType"`         // framework/control - default is framework
	Submit             *bool                     `json:"submit"`             // Submit results to Armo BE - default will
	HostScanner        *bool                     `json:"hostScanner"`        // Deploy ARMO K8s host scanner to collect data from certain controls
	KeepLocal          *bool                     `json:"keepLocal"`          // Do not submit results
	UseCachedArtifacts *bool                     `json:"useCachedArtifacts"` // Use the cached artifacts instead of downloading
	// UseExceptions      string      // Load file with exceptions configuration
	// ControlsInputs     string      // Load file with inputs for controls
	// VerboseMode        bool        // Display all of the input resources and not only failed resources
}

type ScanResponseType string

const (
	ErrorScanResponseType     ScanResponseType = "error"
	ResultsV1ScanResponseType ScanResponseType = "v1results"
	IDScanResponseType        ScanResponseType = "id"
	StatusScanResponseType    ScanResponseType = "status"
)

type Response struct {
	ID       string           `json:"id"`
	Type     ScanResponseType `json:"type"`
	Response []byte           `json:"response,omitempty"`
}
