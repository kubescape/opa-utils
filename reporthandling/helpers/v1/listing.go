package v1

import (
	"github.com/armosec/utils-go/str"
	"github.com/kubescape/opa-utils/reporthandling/apis"
)

// ReportObject any report object must be compliment with a map[string]interface{} structures
type ReportObject map[string]interface{}

// AllLists lists of resources/policies grouped by the status, this structure is meant for internal use of report handling and not an API
type AllLists struct {
	passed   []string
	excluded []string
	failed   []string
	skipped  []string
	other    []string
}

type Iterator interface {
	HasNext() bool
	Next() string
	Len() int
}

type AllListsIterator struct {
	size          int
	index         int
	failedIndex   int
	excludedIndex int
	passIndex     int
	skippedIndex  int
	otherIndex    int
	allLists      *AllLists
}

func (all *AllLists) createIterator() Iterator {
	return &AllListsIterator{
		size:     len(all.failed) + len(all.excluded) + len(all.passed) + len(all.skipped) + len(all.other),
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
		} else if iter.excludedIndex < len(iter.allLists.excluded) {
			item = iter.allLists.excluded[iter.excludedIndex]
			iter.excludedIndex++
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

func (all *AllLists) Failed() []string   { return all.failed }
func (all *AllLists) Passed() []string   { return all.passed }
func (all *AllLists) Excluded() []string { return all.excluded }
func (all *AllLists) Skipped() []string  { return all.skipped }
func (all *AllLists) Other() []string    { return all.other }
func (all *AllLists) All() Iterator {
	return all.createIterator()
}

// Append append single string to matching status list
func (all *AllLists) Append(status apis.ScanningStatus, str ...string) {

	switch status {
	case apis.StatusPassed:
		all.passed = append(all.passed, str...)
	case apis.StatusFailed:
		all.failed = append(all.failed, str...)
	case apis.StatusExcluded:
		all.excluded = append(all.excluded, str...)
	case apis.StatusSkipped:
		all.skipped = append(all.skipped, str...)
	default:
		all.other = append(all.other, str...)
	}
}

// Update AllLists objects with
func (all *AllLists) Update(all2 *AllLists) {
	all.passed = append(all.passed, all2.passed...)
	all.failed = append(all.failed, all2.failed...)
	all.excluded = append(all.excluded, all2.excluded...)
	all.skipped = append(all.skipped, all2.skipped...)
	all.other = append(all.other, all2.other...)
}

// ToUnique - Call this function only when setting the List
func (all *AllLists) toUniqueBase() {
	// remove duplications from each resource list
	all.failed = str.SliceStringToUnique(all.failed)
	all.excluded = str.SliceStringToUnique(all.excluded)
	all.passed = str.SliceStringToUnique(all.passed)
	all.skipped = str.SliceStringToUnique(all.skipped)
	all.other = str.SliceStringToUnique(all.other)
}

// ToUnique - Call this function only when setting the List
func (all *AllLists) ToUniqueControls() {
	all.toUniqueBase()

	// remove failed from excluded list
	all.excluded = trimUnique(all.excluded, all.failed)
}

// ToUnique - Call this function only when setting the List
func (all *AllLists) ToUniqueResources() {
	all.toUniqueBase()

	// remove failed from excluded list
	all.excluded = trimUnique(all.excluded, all.failed)
	// remove failed and excluded from passed list
	trimmed := append(all.failed, all.excluded...)
	all.passed = trimUnique(all.passed, trimmed)

	// remove failed, excluded and passed from skipped list
	trimmed = append(trimmed, all.passed...)
	all.skipped = trimUnique(all.skipped, trimmed)

	// remove failed, excluded, passed and skipped from other list
	trimmed = append(trimmed, all.skipped...)
	all.other = trimUnique(all.other, trimmed)
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
