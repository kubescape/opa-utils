package v1

import (
	"github.com/armosec/utils-go/str"
	"github.com/kubescape/opa-utils/reporthandling/apis"
	"k8s.io/apimachinery/pkg/util/sets"
)

// ReportObject any report object must be compliment with a map[string]interface{} structures
type ReportObject map[string]interface{}

// AllLists lists of resources/policies grouped by the status, this structure is meant for internal use of report handling and not an API
type AllLists struct {
	passed   sets.Set[string]
	failed   sets.Set[string]
	skipped  sets.Set[string]
	excluded sets.Set[string]
	other    sets.Set[string]
}

type Iterator interface {
	HasNext() bool
	Next() string
	Len() int
}

type AllListsIterator struct {
	passed        []string
	failed        []string
	skipped       []string
	other         []string
	size          int
	index         int
	failedIndex   int
	excludedIndex int
	passIndex     int
	skippedIndex  int
	otherIndex    int
}

func (all *AllLists) createIterator() Iterator {
	return &AllListsIterator{
		size:    len(all.failed) + len(all.excluded) + len(all.passed) + len(all.skipped) + len(all.other),
		passed:  all.Passed(),
		failed:  all.Failed(),
		skipped: all.Skipped(),
		other:   all.Other(),
	}
}

func (iter *AllListsIterator) Len() int {
	return iter.size
}

func (iter *AllListsIterator) HasNext() bool {
	return iter.index < iter.size
}

func (iter *AllListsIterator) Next() string {
	var item string
	if iter.HasNext() {
		if iter.failedIndex < len(iter.failed) {
			item = iter.failed[iter.failedIndex]
			iter.failedIndex++
		} else if iter.passIndex < len(iter.passed) {
			item = iter.passed[iter.passIndex]
			iter.passIndex++
		} else if iter.skippedIndex < len(iter.skipped) {
			item = iter.skipped[iter.skippedIndex]
			iter.skippedIndex++
		} else if iter.otherIndex < len(iter.other) {
			item = iter.other[iter.otherIndex]
			iter.otherIndex++
		}
		iter.index++
	}
	return item
}

// GetAllResources

func (all *AllLists) Failed() []string  { return all.failed.UnsortedList() }
func (all *AllLists) Passed() []string  { return all.passed.UnsortedList() }
func (all *AllLists) Skipped() []string { return all.skipped.UnsortedList() }
func (all *AllLists) Other() []string   { return all.other.UnsortedList() }
func (all *AllLists) All() Iterator {
	return all.createIterator()
}

// Append appends strings to matching status list
func (all *AllLists) Append(status apis.ScanningStatus, str ...string) {
	switch status {
	case apis.StatusPassed:
		if all.passed == nil {
			all.passed = sets.New(str...)
		} else {
			all.passed.Insert(str...)
		}
	case apis.StatusSkipped:
		if all.skipped == nil {
			all.skipped = sets.New(str...)
		} else {
			all.skipped.Insert(str...)
		}
	case apis.StatusFailed:
		if all.failed == nil {
			all.failed = sets.New(str...)
		} else {
			all.failed.Insert(str...)
		}
	default:
		if all.other == nil {
			all.other = sets.New(str...)
		} else {
			all.other.Insert(str...)
		}
	}
}

// Update AllLists objects with
func (all *AllLists) Update(all2 *AllLists) {
	all.passed.Insert(all2.Passed()...)
	all.skipped.Insert(all2.Skipped()...)
	all.failed.Insert(all2.Failed()...)
	all.other.Insert(all2.Other()...)
}

// ToUnique - Call this function only when setting the List
func (all *AllLists) ToUniqueControls() {
}

// ToUnique - Call this function only when setting the List
func (all *AllLists) ToUniqueResources() {
	// remove failed from passed list
	all.passed = all.passed.Difference(all.failed)
	// remove failed and passed from skipped list
	all.skipped = all.skipped.Difference(all.failed).Difference(all.passed)
	// remove failed, passed and skipped from other list
	all.other = all.other.Difference(all.failed).Difference(all.passed).Difference(all.skipped)
}

// trimUnique trim the list, return original list without the "trimFrom" list. the list is trimmed in place, so the original list is modified. Also, the list is not sorted
func trimUnique(origin, trimFrom []string) []string {
	if len(origin) == 0 || len(trimFrom) == 0 { // if there is nothing to trim
		return origin
	}
	toRemove := make(map[string]bool, len(trimFrom))

	for i := range trimFrom {
		toRemove[trimFrom[i]] = true
	}

	originLen := len(origin)
	for i := 0; i < originLen; {
		if _, ok := toRemove[origin[i]]; ok {
			str.RemoveIndexFromStringSlice(&origin, i)
			originLen--
		} else {
			i++
		}
	}
	return origin
}

// appendSlice append a slice to a slice the index indicates the position of the slice
func appendSlice(origin, appendTo []string, index *int) {
	for i := range appendTo {
		origin[*index] = appendTo[i]
		*index++
	}
}
