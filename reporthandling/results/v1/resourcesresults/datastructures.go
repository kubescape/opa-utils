package resourcesresults

import (
	"github.com/armosec/armoapi-go/armotypes"
	"github.com/kubescape/opa-utils/reporthandling"
	"github.com/kubescape/opa-utils/reporthandling/apis"
	"github.com/kubescape/opa-utils/reporthandling/results/v1/prioritization"
)

// Result - resource result resourceID and the controls that where tested against the resource, along with the raw resource and prioritization information
type Result struct {
	ResourceID string `json:"resourceID"` // <apigroup>/<namespace>/<kind>/<name>

	AssociatedControls  []ResourceAssociatedControl         `json:"controls,omitempty"`
	RawResource         reporthandling.Resource             `json:"rawResource,omitempty"`
	PrioritizedResource *prioritization.PrioritizedResource `json:"prioritizedResource,omitempty"`
}

// ResourceAssociatedControl control that is associated to a resource
type ResourceAssociatedControl struct {
	ControlID               string                   `json:"controlID"`
	Name                    string                   `json:"name"`
	ResourceAssociatedRules []ResourceAssociatedRule `json:"rules,omitempty"`
}

// ResourceAssociatedRule failed rule that is associated to a resource
type ResourceAssociatedRule struct {
	Name                  string                             `json:"name"` // rule name
	Status                apis.ScanningStatus                `json:"status"`
	Paths                 []armotypes.PosturePaths           `json:"paths,omitempty"`
	Exception             []armotypes.PostureExceptionPolicy `json:"exception,omitempty"`
	ControlConfigurations map[string][]string                `json:"controlConfigurations,omitempty"`
}
