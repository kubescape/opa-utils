package v1alpha1

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

func AttackTrackMock(data AttackTrackStep) *AttackTrack {
	at := AttackTrack{}
	at.Metadata = make(map[string]interface{})
	at.Metadata["name"] = "TestAttackTrack"
	at.Spec.Version = "1.0"
	at.Spec.Data = data
	return &at
}
