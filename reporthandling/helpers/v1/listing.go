package v1

import (
	"github.com/armosec/opa-utils/reporthandling/apis"
	"github.com/armosec/opa-utils/shared"
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

func (all *AllLists) Failed() []string   { return all.failed }
func (all *AllLists) Passed() []string   { return all.passed }
func (all *AllLists) Excluded() []string { return all.excluded }
func (all *AllLists) Skipped() []string  { return all.skipped }
func (all *AllLists) Other() []string    { return all.other }
func (all *AllLists) All() []string {
	l := []string{}
	l = append(l, all.failed...)
	l = append(l, all.excluded...)
	l = append(l, all.passed...)
	l = append(l, all.skipped...)
	l = append(l, all.other...)
	return shared.SliceStringToUnique(l)
}

// Append append single string to matching status list
func (all *AllLists) Append(status apis.ScanningStatus, str string) {

	switch status {
	case apis.StatusPassed:
		all.passed = append(all.passed, str)
	case apis.StatusFailed:
		all.failed = append(all.failed, str)
	case apis.StatusExcluded:
		all.excluded = append(all.excluded, str)
	case apis.StatusSkipped:
		all.skipped = append(all.skipped, str)
	default:
		all.other = append(all.other, str)
	}

	// make sure the lists are unique
	all.toUnique()
}

// Update AllLists objects with
func (all *AllLists) Update(all2 *AllLists) {
	all.passed = append(all.passed, all2.passed...)
	all.failed = append(all.failed, all2.failed...)
	all.excluded = append(all.excluded, all2.excluded...)
	all.skipped = append(all.skipped, all2.skipped...)
	all.other = append(all.other, all2.other...)

	// make sure the lists are unique
	all.toUnique()
}

func (all *AllLists) toUnique() {
	// remove duplications from each resource list
	all.failed = shared.SliceStringToUnique(all.failed)
	all.excluded = shared.SliceStringToUnique(all.excluded)
	all.passed = shared.SliceStringToUnique(all.passed)
	all.skipped = shared.SliceStringToUnique(all.skipped)
	all.other = shared.SliceStringToUnique(all.other)

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

// trimUnique trim the list, return original list without the "trimFrom" list
func trimUnique(origin, trimFrom []string) []string {
	if len(origin) == 0 || len(trimFrom) == 0 { // if there is nothing to trim
		return origin
	}
	unique := map[string]bool{}
	originList := []string{}

	for i := range trimFrom {
		unique[trimFrom[i]] = true
	}

	for i := range origin {
		if found := unique[origin[i]]; !found {
			originList = append(originList, origin[i])
		}
	}
	return originList
}
