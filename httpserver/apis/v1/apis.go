package v1

// A kind of a Notification Policy
//
// swagger:enum NotificationPolicyKind
type NotificationPolicyKind string

// Supported NotificationKinds
const (
	KindFramework NotificationPolicyKind = "Framework"
	KindControl   NotificationPolicyKind = "Control"
	KindRule      NotificationPolicyKind = "Rule"
)

// Type of a Scan Response
//
// swagger:enum ScanResponseType
type ScanResponseType string

const (
	// Deprecated: will return busy / notBusy instead
	IDScanResponseType ScanResponseType = "id"
	// ErrorScanResponseType indicates a response that reports an error
	ErrorScanResponseType ScanResponseType = "error"
	// ResultsV1ScanResponseType indicates a response that carries a v1 Results object as payload
	ResultsV1ScanResponseType ScanResponseType = "v1results"
	// BusyScanResponseType indicates that a server is busy with a previous request
	BusyScanResponseType ScanResponseType = "busy"
	// NotBusyScanResponseType indicates that a server is not busy with a previous request
	NotBusyScanResponseType ScanResponseType = "notBusy"
	// ReadyScanResponseType indicates that a server has successfully completed a request
	ReadyScanResponseType ScanResponseType = "ready"
)

// A WorkloadScan contains the identifiers of a workload to be scanned
type WorkloadScan struct {
	// The ApiVersion of the workload
	ApiVersion string `json:"apiVersion"`
	// The kind of the workload
	Kind string `json:"kind"`
	// The name of the workload
	Name string `json:"name"`
	// The namespace of the workload
	Namespace string `json:"namespace"`
}
