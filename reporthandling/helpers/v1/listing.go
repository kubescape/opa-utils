package v1

import (
	"github.com/armosec/utils-go/str"
	"github.com/kubescape/opa-utils/reporthandling/apis"
)

// ReportObject any report object must be compliment with a map[string]interface{} structures
type ReportObject map[string]interface{}

// AllLists lists of resources/policies grouped by the status, this structure is meant for internal use of report handling and not an API
type AllLists struct {
	passed                []string
	passedExceptions      []string
	passedIrrelevant      []string
	failed                []string
	skippedConfiguration  []string
	skippedIntegration    []string
	skippedRequiresReview []string
	skippedManualReview   []string
	other                 []string
}

type Iterator interface {
	HasNext() bool
	Next() string
	Len() int
}

type AllListsIterator struct {
	allLists                   *AllLists
	size                       int
	index                      int
	failedIndex                int
	passIndex                  int
	passExceptionIndex         int
	passIrrelevantIndex        int
	skippedConfigurationIndex  int
	skippedIntegrationIndex    int
	skippedRequiresReviewIndex int
	skippedManualReviewIndex   int
	otherIndex                 int
}

func (all *AllLists) createIterator() Iterator {
	return &AllListsIterator{
		size:     len(all.failed) + len(all.passed) + len(all.passedExceptions) + len(all.passedIrrelevant) + len(all.skippedConfiguration) + len(all.skippedIntegration) + len(all.skippedRequiresReview) + len(all.skippedManualReview) + len(all.other),
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
		} else if iter.passExceptionIndex < len(iter.allLists.passedExceptions) {
			item = iter.allLists.passedExceptions[iter.passExceptionIndex]
			iter.passExceptionIndex++
		} else if iter.passIrrelevantIndex < len(iter.allLists.passedIrrelevant) {
			item = iter.allLists.passedIrrelevant[iter.passIrrelevantIndex]
			iter.passIrrelevantIndex++
		} else if iter.skippedConfigurationIndex < len(iter.allLists.skippedConfiguration) {
			item = iter.allLists.skippedConfiguration[iter.skippedConfigurationIndex]
			iter.skippedConfigurationIndex++
		} else if iter.skippedIntegrationIndex < len(iter.allLists.skippedIntegration) {
			item = iter.allLists.skippedIntegration[iter.skippedIntegrationIndex]
			iter.skippedIntegrationIndex++
		} else if iter.skippedRequiresReviewIndex < len(iter.allLists.skippedRequiresReview) {
			item = iter.allLists.skippedRequiresReview[iter.skippedRequiresReviewIndex]
			iter.skippedRequiresReviewIndex++
		} else if iter.skippedManualReviewIndex < len(iter.allLists.skippedManualReview) {
			item = iter.allLists.skippedManualReview[iter.skippedManualReviewIndex]
			iter.skippedManualReviewIndex++
		} else if iter.otherIndex < len(iter.allLists.other) {
			item = iter.allLists.other[iter.otherIndex]
			iter.otherIndex++
		}
		iter.index++
	}
	return item
}

// GetAllResources

func (all *AllLists) Failed() []string                { return all.failed }
func (all *AllLists) Passed() []string                { return all.passed }
func (all *AllLists) PassedExceptions() []string      { return all.passedExceptions }
func (all *AllLists) PassedIrrelevant() []string      { return all.passedIrrelevant }
func (all *AllLists) SkippedConfiguration() []string  { return all.skippedConfiguration }
func (all *AllLists) SkippedIntegration() []string    { return all.skippedIntegration }
func (all *AllLists) SkippedRequiresReview() []string { return all.skippedRequiresReview }
func (all *AllLists) SkippedManualReview() []string   { return all.skippedManualReview }
func (all *AllLists) Other() []string                 { return all.other }
func (all *AllLists) All() Iterator {
	return all.createIterator()
}

// Append append single string to matching status list
func (all *AllLists) Append(status apis.ScanningStatus, str ...string) {
	switch status {
	case apis.StatusPassed:
		all.passed = append(all.passed, str...)
	case apis.SubStatusException:
		all.passedExceptions = append(all.passedExceptions, str...)
	case apis.SubStatusIrrelevant:
		all.passedIrrelevant = append(all.passedIrrelevant, str...)
	case apis.SubStatusConfiguration:
		all.skippedConfiguration = append(all.skippedConfiguration, str...)
	case apis.SubStatusIntegration:
		all.skippedIntegration = append(all.skippedIntegration, str...)
	case apis.SubStatusRequiresReview:
		all.skippedRequiresReview = append(all.skippedRequiresReview, str...)
	case apis.SubStatusManualReview:
		all.skippedManualReview = append(all.skippedManualReview, str...)
	case apis.StatusFailed:
		all.failed = append(all.failed, str...)
	default:
		all.other = append(all.other, str...)
	}
}

// Update AllLists objects with
func (all *AllLists) Update(all2 *AllLists) {
	all.passed = append(all.passed, all2.passed...)
	all.passedExceptions = append(all.passedExceptions, all2.passedExceptions...)
	all.passedIrrelevant = append(all.passedIrrelevant, all2.passedIrrelevant...)
	all.skippedConfiguration = append(all.skippedConfiguration, all2.skippedConfiguration...)
	all.skippedIntegration = append(all.skippedIntegration, all2.skippedIntegration...)
	all.skippedRequiresReview = append(all.skippedRequiresReview, all2.skippedRequiresReview...)
	all.skippedManualReview = append(all.skippedManualReview, all2.skippedManualReview...)
	all.failed = append(all.failed, all2.failed...)
	all.other = append(all.other, all2.other...)
}

// ToUnique - Call this function only when setting the List
func (all *AllLists) toUniqueBase() {
	// remove duplications from each resource list
	all.failed = str.SliceStringToUnique(all.failed)
	all.passed = str.SliceStringToUnique(all.passed)
	all.passedExceptions = str.SliceStringToUnique(all.passedExceptions)
	all.passedIrrelevant = str.SliceStringToUnique(all.passedIrrelevant)
	all.skippedConfiguration = str.SliceStringToUnique(all.skippedConfiguration)
	all.skippedIntegration = str.SliceStringToUnique(all.skippedIntegration)
	all.skippedRequiresReview = str.SliceStringToUnique(all.skippedRequiresReview)
	all.skippedManualReview = str.SliceStringToUnique(all.skippedManualReview)
	all.other = str.SliceStringToUnique(all.other)
}

// ToUnique - Call this function only when setting the List
func (all *AllLists) ToUniqueControls() {
	all.toUniqueBase()
}

// ToUnique - Call this function only when setting the List
func (all *AllLists) ToUniqueResources() {
	all.toUniqueBase()
	// remove failed from passedExceptions list
	all.passedExceptions = trimUnique(all.passedExceptions, all.failed)
	// remove passedIrrelevant from passed list
	all.passed = trimUnique(all.passed, all.passedIrrelevant)
	// remove failed, passedExceptions from passed list
	trimmed := append(all.failed, all.passedExceptions...)
	all.passed = trimUnique(all.passed, trimmed)

	// remove failed, and all kind of passed from skipped list
	trimmed = append(trimmed, all.passed...)
	all.skippedConfiguration = trimUnique(all.skippedConfiguration, trimmed)
	all.skippedIntegration = trimUnique(all.skippedIntegration, trimmed)
	all.skippedManualReview = trimUnique(all.skippedManualReview, trimmed)
	all.skippedRequiresReview = trimUnique(all.skippedRequiresReview, trimmed)

	// remove failed, all passed and all skipped from other list
	trimmed = append(trimmed, all.skippedConfiguration...)
	trimmed = append(trimmed, all.skippedIntegration...)
	trimmed = append(trimmed, all.skippedManualReview...)
	trimmed = append(trimmed, all.skippedRequiresReview...)
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
