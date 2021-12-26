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
	assert.Equal(t, 2, len(listA.ListPassed()))
	assert.Equal(t, 3, len(listA.ListExcluded()))
	assert.Equal(t, 2, len(listA.ListFailed()))
	assert.Equal(t, 1, len(listA.ListSkipped()))
	assert.Equal(t, 4, len(listA.ListOther()))

}

func TestAllListsUpdate(t *testing.T) {
	listA := mockAllListsA()
	listB := mockAllListsB()
	listB.Update(listA)

	oldListB := mockAllListsB()
	assert.Equal(t, len(listA.ListPassed())+len(oldListB.ListPassed()), len(listB.ListPassed()))
	assert.Equal(t, len(listA.ListExcluded())+len(oldListB.ListExcluded()), len(listB.ListExcluded()))
	assert.Equal(t, len(listA.ListFailed())+len(oldListB.ListFailed()), len(listB.ListFailed()))
	assert.Equal(t, len(listA.ListSkipped())+len(oldListB.ListSkipped()), len(listB.ListSkipped()))
	assert.Equal(t, len(listA.ListOther())+len(oldListB.ListOther()), len(listB.ListOther()))
}

func TestAllListsAppend(t *testing.T) {
	listA := mockAllListsA()
	listA.Append(apis.StatusExcluded, "b")
	listA.Append(apis.StatusExcluded, "b")
	listA.Append(apis.StatusExcluded, "b")

	oldListA := mockAllListsA()

	assert.Equal(t, len(oldListA.ListExcluded())+3, len(listA.ListExcluded()))
}

func TestAllListsUnique(t *testing.T) {
	listA := mockAllListsA()
	listA.Append(apis.StatusPassed, "b")
	listA.Append(apis.StatusPassed, "b")
	listA.Append(apis.StatusPassed, "b")

	oldListA := mockAllListsA()

	listA.ToUnique()
	assert.Equal(t, len(oldListA.ListPassed()), len(listA.ListPassed()))
}
