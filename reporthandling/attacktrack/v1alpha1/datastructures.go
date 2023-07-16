package v1alpha1

const (
	ControlTypeTagDevops         string = "devops"
	ControlTypeTagSecurity       string = "security"
	ControlTypeTagCompliance     string = "compliance"
	ControlTypeTagSecurityImpact string = "security-impact"
)

type AttackTrack struct {
	ApiVersion string                   `json:"apiVersion"`
	Kind       string                   `json:"kind"`
	Metadata   map[string]interface{}   `json:"metadata"`
	Spec       AttackTrackSpecification `json:"spec"`
}

type AttackTrackSpecification struct {
	Version     string          `json:"version,omitempty"`
	Description string          `json:"description,omitempty"`
	Data        AttackTrackStep `json:"data"`
}

type AttackTrackStep struct {
	Name              string            `json:"name"`
	Description       string            `json:"description,omitempty"`
	VulnerabilityStep bool              `json:"isVulnerabilityStep,omitempty"`
	SubSteps          []AttackTrackStep `json:"subSteps,omitempty"`

	// failed controls which are related to this step
	Controls []IAttackTrackControl `json:"-"`
}

type AttackTrackIterator struct {
	stack *attackTrackStepStack
}

type AttackTrackAllPathsHandler struct {
	attackTrack     IAttackTrack
	inDegreeZero    map[string]bool               // A map of all nodes with in-degree 0
	outDegreeZero   map[string]bool               // A map of all nodes with out-degree 0
	adjacencyMatrix map[string][]IAttackTrackStep // A map of all nodes with their adjacent nodes
	visited         map[string]bool               // A map of all visited nodes
}

type AttackTrackControlsLookup map[string]map[string][]IAttackTrackControl
