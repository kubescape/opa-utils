package prioritization

// SetResourceID set the resource ID
func (pr *PrioritizedResource) SetResourceID(resourceID string) {
	pr.ResourceID = resourceID
}

// GetResourceID get the resource ID
func (pr *PrioritizedResource) GetResourceID() string {
	return pr.ResourceID
}

// =============================== Score ====================================

// GetScore get resource score, which is a sum of its priority vectors scores
func (pr *PrioritizedResource) GetScore() float64 {
	var score float64 = 0
	for i := range pr.PriorityVector {
		score += pr.PriorityVector[i].GetScore()
	}
	return score
}

// ================================= Listing ==================================

// ListControlsIDs return a list of controls IDs from all priority vectors of a given resource
func (pr *PrioritizedResource) ListControlsIDs() []string {
	ids := make([]string, 0)
	for i := range pr.PriorityVector {
		if pr.PriorityVector[i].Type == ControlPriorityVectorType {
			ids = append(ids, pr.PriorityVector[i].Vector...)
		}
	}
	return ids
}
