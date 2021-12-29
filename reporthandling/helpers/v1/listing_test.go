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

	oldListB := mockAllListsB()
	assert.Equal(t, len(listA.Passed())+len(oldListB.Passed()), len(listB.Passed()))
	assert.Equal(t, len(listA.Excluded())+len(oldListB.Excluded()), len(listB.Excluded()))
	assert.Equal(t, len(listA.Failed())+len(oldListB.Failed()), len(listB.Failed()))
	assert.Equal(t, len(listA.Skipped())+len(oldListB.Skipped()), len(listB.Skipped()))
	assert.Equal(t, len(listA.Other())+len(oldListB.Other()), len(listB.Other()))
}

func TestAllListsAppend(t *testing.T) {
	listA := mockAllListsA()
	listA.Append(apis.StatusExcluded, "b")
	listA.Append(apis.StatusExcluded, "b")
	listA.Append(apis.StatusExcluded, "b")

	oldListA := mockAllListsA()

	assert.Equal(t, len(oldListA.Excluded())+3, len(listA.Excluded()))
}

func TestAllListsUnique(t *testing.T) {
	listA := mockAllListsA()
	listA.Append(apis.StatusPassed, "b")
	listA.Append(apis.StatusPassed, "b")
	listA.Append(apis.StatusPassed, "b")

	oldListA := mockAllListsA()

	listA.ToUnique()
	assert.Equal(t, len(oldListA.Passed()), len(listA.Passed()))
}
