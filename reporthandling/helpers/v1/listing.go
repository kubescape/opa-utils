package helpers

import (
	"sync"

	"github.com/kubescape/opa-utils/reporthandling/apis"
	"golang.org/x/exp/maps"
)

var allListsPool = &sync.Pool{
	New: func() interface{} {
		return &AllLists{}
	},
}

// GetAllListsFromPool get the AllLists object from the pool
func GetAllListsFromPool() *AllLists {
	l, ok := allListsPool.Get().(*AllLists)
	if !ok {
		return nil
	}
	// reset the object before returning it as it might be dirty
	l.Clear()
	return l
}

// PutAllListsToPool put the AllLists object back to the pool
func PutAllListsToPool(l *AllLists) {
	allListsPool.Put(l)
}

// ReportObject any report object must be compliment with a map[string]interface{} structures
type ReportObject map[string]interface{}

// AllLists lists of resources/policies grouped by the status, this structure is meant for internal use of report handling and not an API
type AllLists struct {
	itemToStatus map[string]apis.ScanningStatus
	passed       int
	failed       int
	skipped      int
	other        int
}

func (all *AllLists) Failed() int  { return all.failed }
func (all *AllLists) Passed() int  { return all.passed }
func (all *AllLists) Skipped() int { return all.skipped }
func (all *AllLists) Other() int   { return all.other }
func (all *AllLists) Len() int {
	return all.failed + all.passed + all.skipped + all.other
}
func (all *AllLists) All() map[string]apis.ScanningStatus {
	return all.itemToStatus
}

// Initialize initialize the AllLists object map with the given size - this is an optimization for the map
func (all *AllLists) Initialize(size int) {
	if all.itemToStatus == nil {
		all.itemToStatus = make(map[string]apis.ScanningStatus, size)
	}
}

// Clear remove all items and reset the counters
func (all *AllLists) Clear() {
	if all.itemToStatus != nil {
		maps.Clear(all.itemToStatus)
		all.passed = 0
		all.failed = 0
		all.skipped = 0
		all.other = 0
	}
}

// Append append single string to matching status list
func (all *AllLists) Append(status apis.ScanningStatus, str ...string) {
	if all.itemToStatus == nil {
		all.itemToStatus = make(map[string]apis.ScanningStatus, len(str))
	}

	for _, s := range str {
		oldStatus, exist := all.itemToStatus[s]
		if !exist {
			all.itemToStatus[s] = status
			all.updateCounters(status, true)
			// element exist with different status
		} else if oldStatus != status {
			// check if the new status is more significant
			if result := apis.Compare(oldStatus, status); result == status {
				all.itemToStatus[s] = status
				all.updateCounters(status, true)
				all.updateCounters(oldStatus, false)
			}
		}
	}
}

// Helper function to update status counters
func (all *AllLists) updateCounters(status apis.ScanningStatus, increment bool) {
	var delta int
	if increment {
		delta = 1
	} else {
		delta = -1
	}

	switch status {
	case apis.StatusPassed:
		all.passed += delta
	case apis.StatusFailed:
		all.failed += delta
	case apis.StatusSkipped:
		all.skipped += delta
	default:
		all.other += delta
	}
}

// Update AllLists objects with
func (all *AllLists) Update(all2 *AllLists) {
	for item, status := range all2.itemToStatus {
		all.Append(apis.ScanningStatus(status), item)
	}
}

func (all *AllLists) GetItems(status apis.ScanningStatus) []string {
	var amount int
	switch status {
	case apis.StatusPassed:
		amount = all.passed
	case apis.StatusFailed:
		amount = all.failed
	case apis.StatusSkipped:
		amount = all.skipped
	default:
		amount = all.other
	}

	if amount == 0 {
		return []string{}
	}

	items := make([]string, 0, amount)
	for item, itemStatus := range all.itemToStatus {
		if itemStatus == status {
			items = append(items, item)
		}
	}

	return items
}
