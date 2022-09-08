package prioritization

import "fmt"

func NewControlsVector() *ControlsVector {
	return &ControlsVector{
		Score:  0,
		Type:   ControlPriorityVectorType,
		Vector: []PriorityVectorControl{},
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

// ================================= Vector ==================================

// Add adds an item to the priority vector
func (cv *ControlsVector) Add(item interface{}) error {
	if control, ok := item.(PriorityVectorControl); ok {
		cv.Vector = append(cv.Vector, control)
		return nil
	}
	return fmt.Errorf("failed converting item to PriorityVectorControl")
}

// List returns the priority vector
func (cv *ControlsVector) List() interface{} {
	return cv.Vector
}

// Add adds an item (PriorityVectorControl) to the priority vector
func (cv *ControlsVector) AddControl(control PriorityVectorControl) {
	cv.Vector = append(cv.Vector, control)
}

// ListControls returns the priority vector
func (cv *ControlsVector) ListControls() []PriorityVectorControl {
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

// =============================== Iterator ====================================

// GetIterator initialize an iterator for the priority vector
func (cv *ControlsVector) GetIterator() IPriorityVectorIterator {
	return &ControlsVectorIterator{
		size:   len(cv.Vector),
		index:  0,
		vector: &cv.Vector,
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
		v := (*iter.vector)[iter.index]
		iter.index++
		return v
	}
	return nil
}

// NextControl returns the next PriorityVectorControl
func (iter *ControlsVectorIterator) NextControl() *PriorityVectorControl {
	if iter.HasNext() {
		v := (*iter.vector)[iter.index]
		iter.index++
		return &v
	}
	return nil
}
