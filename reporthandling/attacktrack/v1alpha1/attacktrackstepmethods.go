package v1alpha1

// IsPartOfAttackTrackPath checks if the step can be a part of an attack track path (i.e. step has failed controls)
func (step *AttackTrackStep) IsPartOfAttackTrackPath() bool {
	return len(step.Controls) > 0
}

func (step *AttackTrackStep) GetControls() []IAttackTrackControl {
	return step.Controls
}

func (step *AttackTrackStep) GetDescription() string {
	return step.Description
}

func (step *AttackTrackStep) GetName() string {
	return step.Name
}

func (step *AttackTrackStep) SetControls(controls []IAttackTrackControl) {
	step.Controls = controls
}

func (step *AttackTrackStep) Length() int {
	return len(step.SubSteps)
}

func (step *AttackTrackStep) SubStepAt(index int) IAttackTrackStep {
	return &step.SubSteps[index]
}
