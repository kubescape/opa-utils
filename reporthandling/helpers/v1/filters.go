package helpers

import (
	"github.com/armosec/armoapi-go/armotypes"
	"github.com/armosec/utils-go/str"
	"github.com/kubescape/opa-utils/reporthandling/apis"
)

// Filters fields that might take effect on the resource status. If this objects is empty or nil, the status will be as determined by pre-defined logic
type Filters struct {
	FrameworkNames []string // Framework name may effect the status
}

// ListFrameworkNames list the framework name in filter object. Removes empty names
func (f *Filters) ListFrameworkNames() []string {
	fn := []string{}
	for i := range f.FrameworkNames {
		if f.FrameworkNames[i] != "" {
			fn = append(fn, f.FrameworkNames[i])
		}
	}
	return fn
}

// // FilterPassed fields that might take effect on the resource status. If this objects is empty or nil, the status will be as determined by pre-defined logic
// type FilterPassed struct {
// }

// // FilterFailed fields that might take effect on the resource status. If this objects is empty or nil, the status will be as determined by pre-defined logic
// type FilterFailed struct {
// 	FrameworkName string // Framework name may effect the status
// }

// // FilterExcluded fields that might take effect on the resource status. If this objects is empty or nil, the status will be as determined by pre-defined logic
// type FilterExcluded struct {
// 	FrameworkName string // Framework name may effect the status
// }

// // FilterSkipped fields that might take effect on the resource status. If this objects is empty or nil, the status will be as determined by pre-defined logic
// type FilterSkipped struct {
// 	FrameworkName string // Framework name may effect the status
// }

// FilterExceptions get list of exceptions and return the list of filtered exceptions
func (f *Filters) FilterExceptions(exceptions []armotypes.PostureExceptionPolicy) []armotypes.PostureExceptionPolicy {
	if len(f.ListFrameworkNames()) == 0 || len(exceptions) == 0 { // there is nothing to filter
		return exceptions
	}
	filteredExceptions := []armotypes.PostureExceptionPolicy{}
	for i := range exceptions {
		for j := range exceptions[i].PosturePolicies {
			if exceptions[i].PosturePolicies[j].FrameworkName == "" ||
				str.StringInSliceCaseInsensitive(f.ListFrameworkNames(), exceptions[i].PosturePolicies[j].FrameworkName) {
				filteredExceptions = append(filteredExceptions, exceptions[i])
			}
		}
	}
	return filteredExceptions
}

// ListingFilters filter list based on filters. If nil of empty list, the list will be ignored
type ListingFilters struct {
	FrameworkNames []string              // Framework name may effect the status
	ControlsNames  []string              // Framework name may effect the status
	ControlsIDs    []string              // Framework name may effect the status
	Statuses       []apis.ScanningStatus // Framework name may effect the status
}
