package v1

import v1 "github.com/armosec/opa-utils/httpserver/apis/v1"

type PostScanRequest struct {
	Format             string                    `json:"format,omitempty"`             // Format results (table, json, junit ...) - default json
	Account            string                    `json:"account,omitempty"`            // account ID
	Logger             string                    `json:"-"`                            // logger level - debug/info/error - default is debug
	FailThreshold      float32                   `json:"failThreshold,omitempty"`      // Failure score threshold
	ExcludedNamespaces []string                  `json:"excludedNamespaces,omitempty"` // used for host scanner namespace
	IncludeNamespaces  []string                  `json:"includeNamespaces,omitempty"`  // DEPRECATED?
	TargetNames        []string                  `json:"targetNames,omitempty"`        // default is all
	TargetType         v1.NotificationPolicyKind `json:"targetType,omitempty"`         // framework/control - default is framework
	Submit             *bool                     `json:"submit,omitempty"`             // Submit results to Armo BE - default will
	HostScanner        *bool                     `json:"hostScanner,omitempty"`        // Deploy ARMO K8s host scanner to collect data from certain controls
	KeepLocal          *bool                     `json:"keepLocal,omitempty"`          // Do not submit results
	UseCachedArtifacts *bool                     `json:"useCachedArtifacts,omitempty"` // Use the cached artifacts instead of downloading
	// UseExceptions      string      // Load file with exceptions configuration
	// ControlsInputs     string      // Load file with inputs for controls
	// VerboseMode        bool        // Display all of the input resources and not only failed resources
}

type Response struct {
	ID       string              `json:"id"`
	Type     v1.ScanResponseType `json:"type"`
	Response interface{}         `json:"response,omitempty"`
}
