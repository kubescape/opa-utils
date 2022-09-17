package v1

import (
	"testing"

	"github.com/kubescape/opa-utils/reporthandling/apis"
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

	// Updating list A to contain list B should contain all resources, even duplicates
	listB.Update(listA)
	assert.Equal(t, 24, listB.All().Len())

	// Enforcing unique resources should prune duplicate resources
	listB.ToUniqueResources()
	assert.Equal(t, 11, listB.All().Len())

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
	listA.ToUniqueResources()
	oldListA := mockAllListsA()

	assert.Equal(t, len(oldListA.Excluded()), len(listA.Excluded()))
}

func TestAllListsUniqueResources(t *testing.T) {
	listA := mockAllListsA()
	listA.Append(apis.StatusPassed, "b")
	listA.Append(apis.StatusPassed, "b")
	listA.Append(apis.StatusPassed, "b")
	listA.ToUniqueResources()

	oldListA := mockAllListsA()

	assert.Equal(t, len(oldListA.Passed()), len(listA.Passed()))

	listMock := AllLists{}

	listMock.Append(apis.StatusPassed, "a")
	listMock.ToUniqueResources()
	assert.Equal(t, 1, listMock.All().Len())
	assert.Equal(t, listMock.All().Len(), len(listMock.Passed()))
	assert.Equal(t, 0, len(listMock.Failed()))
	assert.Equal(t, 0, len(listMock.Excluded()))

	listMock.Append(apis.StatusExcluded, "a")
	listMock.Append(apis.StatusFailed, "a")
	listMock.ToUniqueResources()

	assert.Equal(t, 1, listMock.All().Len())
	assert.Equal(t, listMock.All().Len(), len(listMock.Failed()))
	assert.Equal(t, 0, len(listMock.Passed()))
}

func TestAllListsUniqueControls(t *testing.T) {
	listA := mockAllListsA()
	listA.Append(apis.StatusPassed, "b")
	listA.Append(apis.StatusPassed, "b")
	listA.Append(apis.StatusPassed, "b")
	listA.ToUniqueControls()

	oldListA := mockAllListsA()

	assert.Equal(t, len(oldListA.Passed()), len(listA.Passed()))

	listMock := AllLists{}

	listMock.Append(apis.StatusPassed, "a")
	listMock.ToUniqueControls()
	assert.Equal(t, 1, listMock.All().Len())
	assert.Equal(t, listMock.All().Len(), len(listMock.Passed()))
	assert.Equal(t, 0, len(listMock.Failed()))
	assert.Equal(t, 0, len(listMock.Excluded()))

	listMock.Append(apis.StatusExcluded, "b")
	listMock.Append(apis.StatusFailed, "b")
	listMock.ToUniqueControls()

	assert.Equal(t, 2, listMock.All().Len())
	assert.Equal(t, 1, len(listMock.Passed()))
	assert.Equal(t, 0, len(listMock.Excluded()))
}

func TestAppendSlice(t *testing.T) {
	type args struct {
		origin   []string
		expected []string
		appendTo []string
		index    *int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "append to empty slice",
			args: args{
				origin:   make([]string, 2),
				expected: []string{"a", "b"},
				appendTo: []string{"a", "b"},
				index:    intToPointer(0),
			},
		},
		{
			name: "append to non empty slice",
			args: args{
				origin:   []string{"a", "b", "", ""},
				expected: []string{"a", "b", "c", "d"},
				appendTo: []string{"c", "d"},
				index:    intToPointer(2),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			appendSlice(tt.args.origin, tt.args.appendTo, tt.args.index)
			assert.Equal(t, tt.args.expected, tt.args.origin)
			assert.Equal(t, len(tt.args.expected), *tt.args.index)
		})
	}
}
func intToPointer(i int) *int {
	return &i
}

func TestTrimUnique(t *testing.T) {
	type args struct {
		origin   []string
		trimFrom []string
		expected []string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "trim from begging of slice",
			args: args{
				origin:   []string{"a", "b", "c"},
				trimFrom: []string{"a"},
				expected: []string{"c", "b"},
			},
		},
		{
			name: "trim from middle of slice",
			args: args{
				origin:   []string{"a", "b", "c"},
				trimFrom: []string{"b"},
				expected: []string{"a", "c"},
			},
		},
		{
			name: "trim from end of slice",
			args: args{
				origin:   []string{"a", "b", "c"},
				trimFrom: []string{"c"},
				expected: []string{"a", "b"},
			},
		},
		{
			name: "do nothing",
			args: args{
				origin:   []string{"a", "b", "c"},
				trimFrom: []string{"d"},
				expected: []string{"a", "b", "c"},
			},
		},
		{
			name: "trim all",
			args: args{
				origin:   []string{"a", "b", "c"},
				trimFrom: []string{"a", "b", "c"},
				expected: []string{},
			},
		},
		{
			name: "trimFrom larger",
			args: args{
				origin:   []string{"a", "b", "c"},
				trimFrom: []string{"a", "b", "e", "d"},
				expected: []string{"c"},
			},
		},
		{
			name: "trim all not sorted",
			args: args{
				origin:   []string{"c", "a", "b"},
				trimFrom: []string{"a", "b", "c"},
				expected: []string{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dd := trimUnique(tt.args.origin, tt.args.trimFrom)
			assert.Equal(t, tt.args.expected, dd)
		})
	}
}
