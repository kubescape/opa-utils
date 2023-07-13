package v1alpha1

type IAttackTrack interface {
	GetApiVersion() string
	GetKind() string
	GetName() string
	GetDescription() string
	GetVersion() string
	GetData() IAttackTrackStep
	Iterator() IAttackTrackIterator
	IsValid() bool
}

// A step in an attack track
type IAttackTrackStep interface {
	GetName() string                            // returns the name of the step
	GetDescription() string                     // returns the description of the step
	GetControls() []IAttackTrackControl         // returns the list of controls which failed on this step
	SetControls(controls []IAttackTrackControl) // sets the list of controls which failed on this step
	Length() int                                // returns the number of sub steps
	SubStepAt(index int) IAttackTrackStep       // returns a sub step at the given index
	IsPartOfAttackTrackPath() bool              // checks if the step can be a part of an attack track path
	IsLeaf() bool                               // checks if the step is a leaf node
}

// A control related to an attack track step
type IAttackTrackControl interface {
	GetAttackTrackCategories(attackTrack string) []string
	GetControlTypeTags() []string
	GetControlId() string
	GetScore() float64
	GetSeverity() int
}

// Iterator interface for iterating over the attack track's steps
type IAttackTrackIterator interface {
	HasNext() bool
	Next() IAttackTrackStep
}

type IAttackTrackControlsLookup interface {
	// returns a list of controls associated with the given attack track and category
	GetAssociatedControls(attackTrack, category string) []IAttackTrackControl

	// returns true if attack track as any associated controls
	HasAssociatedControls(attackTrack string) bool
}
