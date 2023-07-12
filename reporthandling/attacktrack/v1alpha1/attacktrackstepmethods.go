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

// Equal checks if the given attack track step is equal to the current one
// If compareControls is true, the controls are also compared
func (s *AttackTrackStep) Equal(other *AttackTrackStep, compareControls bool) bool {
	if s.Name != other.Name || s.Description != other.Description || len(s.SubSteps) != len(other.SubSteps) {
		return false
	}

	for i := range s.SubSteps {
		if !s.SubSteps[i].Equal(&other.SubSteps[i], compareControls) {
			return false
		}
	}

	if compareControls {

		if len(s.Controls) != len(other.Controls) {
			return false
		}

		for i := range s.Controls {

			if !(s.Controls[i] == other.Controls[i]) {
				return false
			}
		}
	}

	return true
}
