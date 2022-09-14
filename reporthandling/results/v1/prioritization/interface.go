package prioritization

type IPriorityVectorIterator interface {
	HasNext() bool
	Next() interface{}
	Len() int
}

type IPriorityVector interface {
	GetType() PriorityVectorType
	GetScore() float64
	SetScore(float64)
	List() interface{}
	Add(interface{}) error
	SetSeverity(int)
	GetSeverity() int
	GetIterator() IPriorityVectorIterator
}
