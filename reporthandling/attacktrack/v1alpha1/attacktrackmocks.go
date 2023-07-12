package v1alpha1

import "github.com/kubescape/k8s-interface/workloadinterface"

type AttackTrackControlMock struct {
	ControlId  string
	Categories []string
	Tags       []string
	BaseScore  float64
	Severity   int
}

func (mock *AttackTrackControlMock) GetAttackTrackCategories(x string) []string {
	return mock.Categories
}

func (mock *AttackTrackControlMock) GetControlTypeTags() []string {
	return mock.Tags
}

func (mock *AttackTrackControlMock) GetControlId() string {
	return mock.ControlId
}

func (mock *AttackTrackControlMock) GetScore() float64 {
	return mock.BaseScore
}

func (mock *AttackTrackControlMock) GetSeverity() int {
	return mock.Severity
}

func GetAttackTrackMock(data AttackTrackStep) IAttackTrack {
	at := AttackTrack{}
	at.Metadata = make(map[string]interface{})
	at.Metadata["name"] = "TestAttackTrack"
	at.Spec.Version = "1.0"
	at.Spec.Data = data
	return &at
}

// Mocked AttackTrackStep implementation for testing
type AttackTrackStepMock struct {
	Name        string
	Description string
	SubSteps    []AttackTrackStepMock
	Controls    []IAttackTrackControl
}

// Mocked AttackTrackStep methods
func (s AttackTrackStepMock) GetName() string {
	return s.Name
}

func (s AttackTrackStepMock) GetDescription() string {
	return s.Description
}

func (s AttackTrackStepMock) GetControls() []IAttackTrackControl {
	return s.Controls
}

func (s AttackTrackStepMock) SubStepAt(index int) IAttackTrackStep {
	return s.SubSteps[index]
}

func (a AttackTrackStepMock) IsPartOfAttackTrackPath() bool {
	return len(a.Controls) > 0
}

func (s AttackTrackStepMock) SetControls(controls []IAttackTrackControl) {
	s.Controls = controls
}

func (s AttackTrackStepMock) Length() int {
	return len(s.SubSteps)
}

// Mocked AttackTrack implementation for testing
type AttackTrackMock struct {
	Kind       string                       `json:"kind"`
	ApiVersion string                       `json:"apiVersion"`
	Metadata   map[string]interface{}       `json:"metadata"`
	Spec       MockAttackTrackSpecification `json:"spec"`
}

// Mocked AttackTrack methods
func (a AttackTrackMock) GetData() IAttackTrackStep {
	return a.Spec.Data
}

func (at AttackTrackMock) GetApiVersion() string {
	return at.ApiVersion
}

func (at AttackTrackMock) GetKind() string {
	return at.Kind
}

func (at AttackTrackMock) GetName() string {
	if v, ok := workloadinterface.InspectMap(at.Metadata, "name"); ok {
		return v.(string)
	}
	return ""
}

func (at AttackTrackMock) GetDescription() string {
	return at.Spec.Description
}

func (at AttackTrackMock) GetVersion() string {
	return at.Spec.Version
}

// IsValid returns true if an attack track is valid
func (at AttackTrackMock) IsValid() bool {
	visited := make(map[string]bool)
	return directedDfs(at.GetData(), visited)
}

func (at AttackTrackMock) Iterator() IAttackTrackIterator {
	s := &attackTrackStepStack{}
	s.Push(at.GetData())

	return &AttackTrackIterator{
		stack: s,
	}
}

type MockAttackTrackSpecification struct {
	Version     string           `json:"version,omitempty"`
	Description string           `json:"description,omitempty"`
	Data        IAttackTrackStep `json:"data"`
}
