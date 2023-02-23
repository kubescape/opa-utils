package helpers

import (
	"github.com/kubescape/opa-utils/reporthandling/apis"
	"github.com/kubescape/opa-utils/reporthandling/internal/slices"
)

// ReportObject any report object must be compliment with a map[string]interface{} structures
type ReportObject map[string]interface{}

// AllLists lists of resources/policies grouped by the status, this structure is meant for internal use of report handling and not an API
type AllLists struct {
	passed   []string
	failed   []string
	skipped  []string
	excluded []string
	other    []string
}

type Iterator interface {
	HasNext() bool
	Next() string
	Len() int
}

type AllListsIterator struct {
	allLists     *AllLists
	size         int
	index        int
	failedIndex  int
	passIndex    int
	skippedIndex int
	otherIndex   int
}

func (all *AllLists) createIterator() Iterator {
	return &AllListsIterator{
		size:     len(all.failed) + len(all.passed) + len(all.skipped) + len(all.other),
		allLists: all,
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
		if iter.failedIndex < len(iter.allLists.failed) {
			item = iter.allLists.failed[iter.failedIndex]
			iter.failedIndex++
		} else if iter.passIndex < len(iter.allLists.passed) {
			item = iter.allLists.passed[iter.passIndex]
			iter.passIndex++
		} else if iter.skippedIndex < len(iter.allLists.skipped) {
			item = iter.allLists.skipped[iter.skippedIndex]
			iter.skippedIndex++
		} else if iter.otherIndex < len(iter.allLists.other) {
			item = iter.allLists.other[iter.otherIndex]
			iter.otherIndex++
		}
		iter.index++
	}
	return item
}

// GetAllResources

func (all *AllLists) Failed() []string  { return all.failed }
func (all *AllLists) Passed() []string  { return append(all.passed, all.excluded...) }
func (all *AllLists) Skipped() []string { return all.skipped }
func (all *AllLists) Other() []string   { return all.other }
func (all *AllLists) All() Iterator {
	return all.createIterator()
}

// Append append single string to matching status list
func (all *AllLists) Append(status apis.ScanningStatus, str ...string) {
	switch status {
	case apis.StatusPassed:
		all.passed = append(all.passed, str...)
	case apis.StatusSkipped:
		all.skipped = append(all.skipped, str...)
	case apis.StatusFailed:
		all.failed = append(all.failed, str...)
	default:
		all.other = append(all.other, str...)
	}
}

// Update AllLists objects with
func (all *AllLists) Update(all2 *AllLists) {
	all.passed = append(all.passed, all2.passed...)
	all.skipped = append(all.skipped, all2.skipped...)
	all.failed = append(all.failed, all2.failed...)
	all.other = append(all.other, all2.other...)
}

// ToUnique - Call this function only when setting the List
func (all *AllLists) toUniqueBase() {
	// remove duplications from each resource list
	all.failed = slices.UniqueStrings(all.failed)
	all.passed = slices.UniqueStrings(all.passed)
	all.skipped = slices.UniqueStrings(all.skipped)
	all.other = slices.UniqueStrings(all.other)
}

// ToUnique - Call this function only when setting the List
func (all *AllLists) ToUniqueControls() {
	all.toUniqueBase()
}

// ToUnique - Call this function only when setting the List
func (all *AllLists) ToUniqueResources() {
	all.failed = slices.UniqueStrings(all.failed)

	const heuristicCapacity = 100 // alloc 100 slots to the stack. The rest would go to the heap - see https://github.com/golang/go/issues/58215

	trimmed := append(make([]string, 0, heuristicCapacity), make([]string, 0, max(len(all.failed)+len(all.excluded)+len(all.passed)+len(all.skipped), heuristicCapacity)-heuristicCapacity)...)

	// remove failed from excluded list
	trimmed = append(trimmed, all.failed...)
	all.skipped = slices.TrimStableUnique(all.skipped, trimmed)

	// remove failed and skipped from passed list
	trimmed = append(trimmed, all.skipped...)
	all.passed = slices.TrimStableUnique(all.passed, trimmed)

	// remove failed, skipped and passed from "other" list
	trimmed = append(trimmed, all.passed...)
	all.other = slices.TrimStableUnique(all.other, trimmed)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
