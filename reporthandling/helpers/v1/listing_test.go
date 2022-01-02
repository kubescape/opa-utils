package v1

import (
	"testing"

	"github.com/armosec/opa-utils/reporthandling/apis"
	"github.com/stretchr/testify/assert"
)

// mockAllListsA - DO NOT CHANGE MOCK FUNCTION RETURN
func mockAllListsA() *AllLists {
	return &AllLists{
		passed:   []string{"a", "b"},
		excluded: []string{"c", "d", "e"},
		failed:   []string{"e", "g"},
		skipped:  []string{"h"},
		other:    []string{"i", "l", "m", "n"},
	}
}

// mockAllListsB - DO NOT CHANGE MOCK FUNCTION RETURN
func mockAllListsB() *AllLists {
	return &AllLists{
		excluded: []string{"a", "b"},
		passed:   []string{"c", "d", "e"},
		other:    []string{"e", "g"},
		failed:   []string{"h"},
		skipped:  []string{"i", "l", "m", "n"},
	}
}

func TestAllLists(t *testing.T) {
	listA := mockAllListsA()
	assert.Equal(t, 2, len(listA.Passed()))
	assert.Equal(t, 3, len(listA.Excluded()))
	assert.Equal(t, 2, len(listA.Failed()))
	assert.Equal(t, 1, len(listA.Skipped()))
	assert.Equal(t, 4, len(listA.Other()))

}

func TestAllListsUpdate(t *testing.T) {
	listA := mockAllListsA()
	listB := mockAllListsB()
	listB.Update(listA)

	assert.Equal(t, 11, len(listB.All()))
	assert.Equal(t, 0, len(listB.Passed()))
	assert.Equal(t, 4, len(listB.Excluded()))
	assert.Equal(t, 3, len(listB.Failed()))
	assert.Equal(t, 4, len(listB.Skipped()))
	assert.Equal(t, 0, len(listB.Other()))
}

func TestAllListsAppend(t *testing.T) {
	listA := mockAllListsA()
	listA.Append(apis.StatusExcluded, "b")
	listA.Append(apis.StatusExcluded, "b")
	listA.Append(apis.StatusExcluded, "b")

	oldListA := mockAllListsA()

	assert.Equal(t, len(oldListA.Excluded()), len(listA.Excluded()))
}

func TestAllListsUnique(t *testing.T) {
	listA := mockAllListsA()
	listA.Append(apis.StatusPassed, "b")
	listA.Append(apis.StatusPassed, "b")
	listA.Append(apis.StatusPassed, "b")

	oldListA := mockAllListsA()

	assert.Equal(t, len(oldListA.Passed()), len(listA.Passed()))

	listMock := AllLists{}

	listMock.Append(apis.StatusPassed, "a")
	assert.Equal(t, 1, len(listMock.All()))
	assert.Equal(t, len(listMock.All()), len(listMock.Passed()))
	assert.Equal(t, 0, len(listMock.Failed()))
	assert.Equal(t, 0, len(listMock.Excluded()))

	listMock.Append(apis.StatusExcluded, "a")
	listMock.Append(apis.StatusFailed, "a")

	assert.Equal(t, 1, len(listMock.All()))
	assert.Equal(t, len(listMock.All()), len(listMock.Failed()))
	assert.Equal(t, 0, len(listMock.Passed()))
}
