package helpers

import (
	"testing"

	"github.com/kubescape/opa-utils/reporthandling/apis"
	"github.com/stretchr/testify/assert"
)

// mockAllListsA - DO NOT CHANGE MOCK FUNCTION RETURN
func mockAllListsA() *AllLists {
	mock := &AllLists{}
	mock.Append(apis.StatusPassed, "a")
	mock.Append(apis.StatusPassed, "b")
	mock.Append(apis.StatusFailed, "e")
	mock.Append(apis.StatusFailed, "g")
	mock.Append(apis.StatusSkipped, "h")
	mock.Append(apis.StatusUnknown, "i")
	mock.Append(apis.StatusUnknown, "l")
	mock.Append(apis.StatusUnknown, "m")
	mock.Append(apis.StatusUnknown, "n")
	return mock
}

// mockAllListsB - DO NOT CHANGE MOCK FUNCTION RETURN
func mockAllListsB() *AllLists {
	mock := &AllLists{}
	mock.Append(apis.StatusPassed, "c")
	mock.Append(apis.StatusPassed, "d")
	mock.Append(apis.StatusPassed, "e")
	mock.Append(apis.StatusUnknown, "e")
	mock.Append(apis.StatusUnknown, "g")
	mock.Append(apis.StatusFailed, "h")
	mock.Append(apis.StatusSkipped, "i")
	mock.Append(apis.StatusSkipped, "l")
	mock.Append(apis.StatusSkipped, "m")
	mock.Append(apis.StatusSkipped, "n")

	return mock
}

func TestAllLists(t *testing.T) {
	listA := mockAllListsA()
	assert.Equal(t, 2, listA.Passed())
	assert.Equal(t, 2, listA.Failed())
	assert.Equal(t, 1, listA.Skipped())
	assert.Equal(t, 4, listA.Other())
}

func TestAllListsUpdate(t *testing.T) {
	listA := mockAllListsA()
	listB := mockAllListsB()

	// Updating list A to contain list B should contain all resources, even duplicates
	// Enforcing unique resources should prune duplicate resources

	listB.Update(listA)
	assert.Equal(t, 11, len(listB.All()))
	assert.Equal(t, 4, listB.Passed())
	assert.Equal(t, 3, listB.Failed())
	assert.Equal(t, 4, listB.Skipped())
	assert.Equal(t, 0, listB.Other())
}

func TestAllListsAppend(t *testing.T) {
	listA := mockAllListsA()
	listA.Append(apis.StatusPassed, "b")
	listA.Append(apis.StatusPassed, "b")
	listA.Append(apis.StatusPassed, "b")

	oldListA := mockAllListsA()

	assert.Equal(t, oldListA.Passed(), listA.Passed())
}

func TestAllListsUniqueResources(t *testing.T) {
	listA := mockAllListsA()
	listA.Append(apis.StatusPassed, "a")
	listA.Append(apis.StatusPassed, "a")
	listA.Append(apis.StatusPassed, "a")

	oldListA := mockAllListsA()

	assert.Equal(t, oldListA.Passed(), listA.Passed())

	listMock := AllLists{}

	listMock.Append(apis.StatusPassed, "a")
	assert.Equal(t, 1, listMock.Len())
	assert.Equal(t, listMock.Len(), listMock.Passed())
	assert.Equal(t, 0, listMock.Failed())
	assert.Equal(t, 0, listMock.Skipped())

	listMock.Append(apis.StatusSkipped, "a")
	listMock.Append(apis.StatusFailed, "a")

	assert.Equal(t, 1, listMock.Len())
	assert.Equal(t, listMock.Len(), listMock.Failed())
	assert.Equal(t, 0, listMock.Passed())
}

func TestAllListsUniqueControls(t *testing.T) {
	listA := mockAllListsA()
	listA.Append(apis.StatusPassed, "b")
	listA.Append(apis.StatusPassed, "b")
	listA.Append(apis.StatusPassed, "b")

	oldListA := mockAllListsA()

	assert.Equal(t, oldListA.Passed(), listA.Passed())

	listMock := AllLists{}

	listMock.Append(apis.StatusPassed, "a")
	assert.Equal(t, 1, listMock.Len())
	assert.Equal(t, listMock.Len(), listMock.Passed())
	assert.Equal(t, 0, listMock.Failed())
	assert.Equal(t, 0, listMock.Skipped())

	listMock.Append(apis.StatusSkipped, "b")
	assert.Equal(t, 1, listMock.Skipped())

	listMock.Append(apis.StatusFailed, "b")
	assert.Equal(t, 2, listMock.Len())
	assert.Equal(t, 1, listMock.Passed())
	assert.Equal(t, 0, listMock.Skipped())
	assert.Equal(t, 1, listMock.Failed())
}

func TestAllListsClear(t *testing.T) {
	listA := mockAllListsA()
	assert.NotEqual(t, 0, listA.Len())
	assert.NotEqual(t, 0, listA.Passed())
	assert.NotEqual(t, 0, listA.Failed())
	assert.NotEqual(t, 0, listA.Skipped())
	assert.NotEqual(t, 0, listA.Other())
	assert.NotEqual(t, 0, len(listA.itemToStatus))

	listA.Clear()

	assert.Equal(t, 0, len(listA.itemToStatus))
	assert.Equal(t, 0, listA.Len())
	assert.Equal(t, 0, listA.Passed())
	assert.Equal(t, 0, listA.Failed())
	assert.Equal(t, 0, listA.Skipped())
	assert.Equal(t, 0, listA.Other())
}

func TestAllListsGetItems(t *testing.T) {
	listA := mockAllListsA()
	assert.ElementsMatch(t, []string{"a", "b"}, listA.GetItems(apis.StatusPassed))
	assert.ElementsMatch(t, []string{"e", "g"}, listA.GetItems(apis.StatusFailed))
	assert.ElementsMatch(t, []string{"h"}, listA.GetItems(apis.StatusSkipped))
	assert.ElementsMatch(t, []string{"l", "n", "i", "m"}, listA.GetItems(apis.StatusUnknown))
}
