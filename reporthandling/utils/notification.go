package utils

import "github.com/armosec/armoapi-go/armotypes"

type NotificationPolicyType string
type NotificationPolicyKind string

type PolicyNotification struct {
	NotificationType NotificationPolicyType     `json:"notificationType"`
	Rules            []PolicyIdentifier         `json:"rules"`
	ReportID         string                     `json:"reportID"`
	JobID            string                     `json:"jobID"`
	Designators      armotypes.PortalDesignator `json:"designators"`
}

type PolicyIdentifier struct {
	Kind NotificationPolicyKind `json:"kind"`
	Name string                 `json:"name"`
}
