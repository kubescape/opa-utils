package resourcesresults

import "github.com/armosec/armoapi-go/armotypes"

// Result - resource result resourceID and the controls that where tested against the resource
type Result struct {
	ResourceID         string                      `json:"resourceID"`
	AssociatedControls []ResourceAssociatedControl `json:"controls,omitempty"`
}

// ResourceAssociatedControl control that is associated to a resource
type ResourceAssociatedControl struct {
	ControlID               string                   `json:"controlID"`
	ResourceAssociatedRules []ResourceAssociatedRule `json:"rules,omitempty"`
	ControlConfigurations   map[string][]string      `json:"controlConfigurations,omitempty"`
}

// ResourceAssociatedRule failed rule that is associated to a resource
type ResourceAssociatedRule struct {
	Name        string                             `json:"name"` // rule name
	FailedPaths []string                           `json:"failedPaths"`
	Exception   []armotypes.PostureExceptionPolicy `json:"exception,omitempty"`
}
