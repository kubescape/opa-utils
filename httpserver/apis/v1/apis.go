package v1

type NotificationPolicyKind string

// Supported NotificationKinds
const (
	KindFramework NotificationPolicyKind = "Framework"
	KindControl   NotificationPolicyKind = "Control"
	KindRule      NotificationPolicyKind = "Rule"
)

type ScanResponseType string

const (
	IDScanResponseType        ScanResponseType = "id"        // DEPRECATED - will return busy/notBusy instead
	ErrorScanResponseType     ScanResponseType = "error"     // error accrued, returning error message
	ResultsV1ScanResponseType ScanResponseType = "v1results" // returning v1 results object
	BusyScanResponseType      ScanResponseType = "busy"      // Server is busy with previous request
	NotBusyScanResponseType   ScanResponseType = "notBusy"   // Server is not busy with previous request
	ReadyScanResponseType     ScanResponseType = "ready"     // Server successfully completed request
)
