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

// CalculateScore calculates resource score, which is a sum of its priority vectors scores
func (pr *PrioritizedResource) CalculateScore() float64 {
	var score float64 = 0
	for i := range pr.PriorityVector {
		score += pr.PriorityVector[i].GetScore()
	}
	return score
}

// GetScore returns the score of the prioritized resource
func (pr *PrioritizedResource) GetScore() float64 {
	return pr.Score
}

// SetScore updates the score of the prioritized resource
func (pr *PrioritizedResource) SetScore(score float64) {
	pr.Score = score
}

// =============================== Severity ====================================

// CalculateSeverity calculates resource severity, which is the max severity of its priority vectors scores
func (pr *PrioritizedResource) CalculateSeverity() int {
	var severity int = 0
	for i := range pr.PriorityVector {
		if pr.PriorityVector[i].GetSeverity() > severity {
			severity = pr.PriorityVector[i].GetSeverity()
		}
	}
	return severity
}

func (pr *PrioritizedResource) GetSeverity() int {
	return pr.Severity
}

func (pr *PrioritizedResource) SetSeverity(severity int) {
	pr.Severity = severity
}

// ================================= Listing ==================================

// ListControlsIDs return a list of controls IDs from all priority vectors of a given resource
func (pr *PrioritizedResource) ListControlsIDs() []string {
	ids := make([]string, 0)
	for i := range pr.PriorityVector {
		if pr.PriorityVector[i].GetType() == ControlPriorityVectorType {
			iter := pr.PriorityVector[i].GetIterator()
			for iter.HasNext() {
				control := iter.Next().(*PriorityVectorControl)
				ids = append(ids, control.ControlID)
			}
		}
	}
	return ids
}
