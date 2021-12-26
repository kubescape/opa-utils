package resourcesresults

import "github.com/armosec/armoapi-go/armotypes"

// Result - resource result resourceID and the controls that where tested against the resource
type Result struct {
	ResourceID         string
	AssociatedControls []ResourceAssociatedControl
}

// ResourceAssociatedControl control that is associated to a resource
type ResourceAssociatedControl struct {
	ControlID string
	// TODO - add list of controls inputs
	ResourceAssociatedRules []ResourceAssociatedRule
}

// ResourceAssociatedRule failed rule that is associated to a resource
type ResourceAssociatedRule struct {
	Name        string                             `json:"name"`
	FailedPaths []string                           `json:"failedPaths"`
	Exception   []armotypes.PostureExceptionPolicy `json:"exception,omitempty"`
}
