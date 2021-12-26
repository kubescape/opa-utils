package v1

// Filters fields that might take effect on the resource status. If this objects is empty or nil, the status will be as determined by pre-defined logic
type Filters struct {
	FilterExcluded
	FilterFailed
	FilterPassed
	FilterSkipped
}

// FilterPassed fields that might take effect on the resource status. If this objects is empty or nil, the status will be as determined by pre-defined logic
type FilterPassed struct {
}

// FilterFailed fields that might take effect on the resource status. If this objects is empty or nil, the status will be as determined by pre-defined logic
type FilterFailed struct {
	FrameworkName string // Framework name may effect the status
}

// FilterExcluded fields that might take effect on the resource status. If this objects is empty or nil, the status will be as determined by pre-defined logic
type FilterExcluded struct {
	FrameworkName string // Framework name may effect the status
}

// FilterSkipped fields that might take effect on the resource status. If this objects is empty or nil, the status will be as determined by pre-defined logic
type FilterSkipped struct {
	FrameworkName string // Framework name may effect the status
}
