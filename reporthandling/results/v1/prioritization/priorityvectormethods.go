package prioritization

func NewControlPriorityVector() *PriorityVector {
	return &PriorityVector{
		Score:  0,
		Type:   ControlPriorityVectorType,
		Vector: []string{},
	}
}

// =============================== Score ====================================

// GetScore get priority vector score
func (pv *PriorityVector) GetScore() float64 {
	return pv.Score
}

// SetScore set priority vector score
func (pr *PriorityVector) SetScore(score float64) {
	pr.Score = score
}

// ================================= Vector ==================================

// Add adds an item to the priority vector
func (pv *PriorityVector) Add(item string) {
	pv.Vector = append(pv.Vector, item)
}

// List returns the list of items in the priority vector
func (pv *PriorityVector) List() []string {
	return pv.Vector
}

// ================================= Type ==================================

// GetType returns the priority vector type
func (pv *PriorityVector) GetType() PriorityVectorType {
	return pv.Type
}
