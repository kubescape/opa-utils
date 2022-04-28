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
	ErrorScanResponseType     ScanResponseType = "error"
	ResultsV1ScanResponseType ScanResponseType = "v1results"
	IDScanResponseType        ScanResponseType = "id"
	StatusScanResponseType    ScanResponseType = "status"
)
