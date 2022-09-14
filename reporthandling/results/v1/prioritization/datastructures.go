package prioritization

type PriorityVectorType string

const (
	ControlPriorityVectorType PriorityVectorType = "control"
)

type PriorityVectorControl struct {
	ControlID string   `json:"controlID"`
	Category  string   `json:"category"`
	Tags      []string `json:"tags"`
}

// ControlsVector - list of controls which represent a priority vector
type ControlsVector struct {
	AttackTrackName string                   `json:"attackTrackName"`
	Type            PriorityVectorType       `json:"type"`
	Vector          []*PriorityVectorControl `json:"vector"`
	Score           float64                  `json:"score"`
	Severity        int                      `json:"severity"`
}

// PrioritizedResource - resource with a score based on its priority vectors
type PrioritizedResource struct {
	ResourceID     string            `json:"resourceID"`
	PriorityVector []*ControlsVector `json:"priorityVector"`
	Score          float64           `json:"score"`
	Severity       int               `json:"severity"`
}

type ControlsVectorIterator struct {
	vector []*PriorityVectorControl
	size   int
	index  int
}
