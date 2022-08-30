package prioritization

type PriorityVectorType string

const (
	ControlPriorityVectorType PriorityVectorType = "control"
)

// PriorityVector a list of items of a specific type and a score
type PriorityVector struct {
	Type   PriorityVectorType `json:"type"`
	Vector []string           `json:"vector"`
	Score  float64            `json:"score"`
}

// PrioritizedResource - resource with a score based on its priority vectors
type PrioritizedResource struct {
	ResourceID     string           `json:"resourceID"`
	PriorityVector []PriorityVector `json:"priorityVector"`
}
