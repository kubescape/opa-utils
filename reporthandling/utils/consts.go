package utils

const (
	StatusPassed  string = "passed"
	StatusWarning string = "warning"
	StatusIgnore  string = "ignore"
	StatusFailed  string = "failed"
	StatusSkipped string = "skipped"
)

// Supported NotificationKinds
const (
	KindFramework NotificationPolicyKind = "Framework"
	KindControl   NotificationPolicyKind = "Control"
	KindRule      NotificationPolicyKind = "Rule"
)
