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

func (all *AllLists) ListFailed() []string   { return all.failed }
func (all *AllLists) ListPassed() []string   { return all.passed }
func (all *AllLists) ListExcluded() []string { return all.excluded }
func (all *AllLists) ListSkipped() []string  { return all.skipped }
func (all *AllLists) ListOther() []string    { return all.other }
func (all *AllLists) ListAll() []string {
	l := []string{}
	l = append(l, all.failed...)
	l = append(l, all.excluded...)
	l = append(l, all.passed...)
	l = append(l, all.skipped...)
	l = append(l, all.other...)
	return l
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
}

// Update AllLists objects with
func (all *AllLists) Update(all2 *AllLists) {
	all.passed = append(all.passed, all2.passed...)
	all.failed = append(all.failed, all2.failed...)
	all.excluded = append(all.excluded, all2.excluded...)
	all.skipped = append(all.skipped, all2.skipped...)
	all.other = append(all.other, all2.other...)
}

func (all *AllLists) ToUnique() {
	all.passed = shared.SliceStringToUnique(all.passed)
	all.failed = shared.SliceStringToUnique(all.failed)
	all.excluded = shared.SliceStringToUnique(all.excluded)
	all.skipped = shared.SliceStringToUnique(all.skipped)
	all.other = shared.SliceStringToUnique(all.other)
}
