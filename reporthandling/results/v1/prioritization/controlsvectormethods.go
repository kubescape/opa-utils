package prioritization

import (
	"fmt"

	"github.com/kubescape/opa-utils/reporthandling/attacktrack/v1alpha1"
)

func NewControlsVector(attackTrackName string) *ControlsVector {
	return &ControlsVector{
		AttackTrackName: attackTrackName,
		Score:           0,
		Type:            ControlPriorityVectorType,
		Vector:          []*PriorityVectorControl{},
	}
}

// =============================== Score ====================================

// GetScore get priority vector score
func (cv *ControlsVector) GetScore() float64 {
	return cv.Score
}

// SetScore set priority vector score
func (cv *ControlsVector) SetScore(score float64) {
	cv.Score = score
}

// CalculateScore calculates the priority vector score
func (cv *ControlsVector) CalculateScore(allControls map[string]v1alpha1.IAttackTrackControl, replicaCount int) (float64, error) {
	if len(cv.Vector) == 0 {
		return 0, nil
	}

	var totalScore float64 = 1
	for i := range cv.Vector {
		if control, ok := allControls[cv.Vector[i].ControlID]; ok {
			totalScore *= control.GetScore()
		} else {
			return 0, fmt.Errorf("failed finding control %s in map", cv.Vector[i].ControlID)
		}
	}

	totalScore *= 1 + float64(replicaCount)/10

	return totalScore, nil
}

// ================================= Vector ==================================

// Add adds an item to the priority vector
func (cv *ControlsVector) Add(item interface{}) error {
	if control, ok := item.(PriorityVectorControl); ok {
		cv.Vector = append(cv.Vector, &control)
		return nil
	}
	return fmt.Errorf("failed converting item to PriorityVectorControl")
}

// List returns the priority vector
func (cv *ControlsVector) List() interface{} {
	return cv.Vector
}

// Add adds an item (PriorityVectorControl) to the priority vector
func (cv *ControlsVector) AddControl(control *PriorityVectorControl) {
	cv.Vector = append(cv.Vector, control)
}

// ListControls returns the priority vector
func (cv *ControlsVector) ListControls() []*PriorityVectorControl {
	return cv.Vector
}

// ================================= Type ==================================

// GetType returns the priority vector type
func (cv *ControlsVector) GetType() PriorityVectorType {
	return cv.Type
}

// =============================== Severity ====================================

// GetSeverity returns the severity of the controls vector
func (cv *ControlsVector) GetSeverity() int {
	return cv.Severity
}

func (cv *ControlsVector) SetSeverity(severity int) {
	cv.Severity = severity
}

// CalculateSeverity calculates the priority vector severity
func (cv *ControlsVector) CalculateSeverity(allControls map[string]v1alpha1.IAttackTrackControl) (int, error) {
	maxSeverity := 0
	for i := range cv.Vector {
		if control, ok := allControls[cv.Vector[i].ControlID]; ok {
			controlSeverity := control.GetSeverity()
			if controlSeverity > maxSeverity {
				maxSeverity = controlSeverity
			}
		} else {
			return 0, fmt.Errorf("failed finding control %s in map", cv.Vector[i].ControlID)
		}
	}
	return maxSeverity, nil
}

// =============================== Iterator ====================================

// GetIterator initialize an iterator for the priority vector
func (cv *ControlsVector) GetIterator() IPriorityVectorIterator {
	return &ControlsVectorIterator{
		size:   len(cv.Vector),
		index:  0,
		vector: cv.Vector,
	}
}

// Len returns the length of the vector
func (iter *ControlsVectorIterator) Len() int {
	return iter.size
}

// HasNext returns true if vector has more items to iterate on
func (iter *ControlsVectorIterator) HasNext() bool {
	return iter.index < iter.size
}

// Next returns the next item in the vector
func (iter *ControlsVectorIterator) Next() interface{} {
	if iter.HasNext() {
		v := (iter.vector)[iter.index]
		iter.index++
		return v
	}
	return nil
}

// NextControl returns the next PriorityVectorControl
func (iter *ControlsVectorIterator) NextControl() *PriorityVectorControl {
	if iter.HasNext() {
		v := (iter.vector)[iter.index]
		iter.index++
		return v
	}
	return nil
}

// A valid vector must have at least one control which is not a Security Impact-only Control
func (v *ControlsVector) IsValid() bool {
	iter := v.GetIterator().(*ControlsVectorIterator)
	for iter.HasNext() {
		control := iter.NextControl()
		for _, tag := range control.Tags {
			if tag != v1alpha1.ControlTypeTagSecurityImpact {
				return true
			}
		}
	}
	return false
}
