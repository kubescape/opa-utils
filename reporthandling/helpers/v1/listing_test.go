package v1

import (
	"testing"

	"github.com/kubescape/opa-utils/reporthandling/apis"
	"github.com/stretchr/testify/assert"
)

// mockAllListsA - DO NOT CHANGE MOCK FUNCTION RETURN
func mockAllListsA() *AllLists {
	return &AllLists{
		passed:                []string{"a", "b"},
		passedExceptions:      []string{"c"},
		passedIrrelevant:      []string{"d"},
		skippedConfiguration:  []string{"e", "a"},
		skippedIntegration:    []string{"g"},
		skippedRequiresReview: []string{"h", "i"},
		skippedManualReview:   []string{"j", "k"},
		failed:                []string{"e", "g"},
		other:                 []string{"i", "l", "m", "n"},
	}
}

// mockAllListsB - DO NOT CHANGE MOCK FUNCTION RETURN
func mockAllListsB() *AllLists {
	return &AllLists{
		passedExceptions:      []string{"a", "b"},
		passedIrrelevant:      []string{"c"},
		passed:                []string{"d"},
		skippedIntegration:    []string{"e", "a"},
		failed:                []string{"g"},
		skippedManualReview:   []string{"h", "i"},
		other:                 []string{"j", "k"},
		skippedConfiguration:  []string{"e", "g"},
		skippedRequiresReview: []string{"i", "l", "m", "n"},
	}
}

func TestAllLists(t *testing.T) {
	listA := mockAllListsA()
	assert.Equal(t, 2, len(listA.Passed()))
	assert.Equal(t, 1, len(listA.PassedExceptions()))
	assert.Equal(t, 1, len(listA.PassedIrrelevant()))
	assert.Equal(t, 2, len(listA.Failed()))
	assert.Equal(t, 2, len(listA.SkippedConfiguration()))
	assert.Equal(t, 1, len(listA.SkippedIntegration()))
	assert.Equal(t, 2, len(listA.SkippedManualReview()))
	assert.Equal(t, 2, len(listA.SkippedRequiresReview()))
	assert.Equal(t, 4, len(listA.Other()))

}

func TestAllListsUpdate(t *testing.T) {
	listA := mockAllListsA()
	listB := mockAllListsB()

	// Updating list A to contain list B should contain all resources, even duplicates
	listB.Update(listA)
	assert.Equal(t, 34, listB.All().Len())

	// Enforcing unique resources should prune duplicate resources
	listB.ToUniqueResources()
	assert.Equal(t, 16, listB.All().Len())

	assert.Equal(t, 0, len(listB.Passed()))
	assert.Equal(t, 3, len(listB.PassedExceptions()))
	assert.Equal(t, 2, len(listB.PassedIrrelevant()))
	assert.Equal(t, 2, len(listB.Failed()))
	assert.Equal(t, 0, len(listB.SkippedConfiguration()))
	assert.Equal(t, 0, len(listB.SkippedIntegration()))
	assert.Equal(t, 4, len(listB.SkippedManualReview()))
	assert.Equal(t, 5, len(listB.SkippedRequiresReview()))
	assert.Equal(t, 0, len(listB.Other()))
}

func TestAllListsAppend(t *testing.T) {
	listA := mockAllListsA()
	listA.Append(apis.SubStatusException, "e")
	listA.Append(apis.SubStatusException, "e")
	listA.Append(apis.SubStatusException, "e")
	listA.ToUniqueResources()
	oldListA := mockAllListsA()

	assert.Equal(t, len(oldListA.PassedExceptions()), len(listA.PassedExceptions()))
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
	assert.Equal(t, 0, len(listMock.PassedExceptions()))
	assert.Equal(t, 0, len(listMock.PassedIrrelevant()))
	assert.Equal(t, 0, len(listMock.SkippedConfiguration()))
	assert.Equal(t, 0, len(listMock.SkippedIntegration()))
	assert.Equal(t, 0, len(listMock.SkippedManualReview()))
	assert.Equal(t, 0, len(listMock.SkippedRequiresReview()))

	listMock.Append(apis.SubStatusException, "a")
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
	assert.Equal(t, 0, len(listMock.PassedExceptions()))
	assert.Equal(t, 0, len(listMock.PassedIrrelevant()))
	assert.Equal(t, 0, len(listMock.SkippedConfiguration()))
	assert.Equal(t, 0, len(listMock.SkippedIntegration()))
	assert.Equal(t, 0, len(listMock.SkippedManualReview()))
	assert.Equal(t, 0, len(listMock.SkippedRequiresReview()))

	listMock.Append(apis.SubStatusException, "b")
	listMock.Append(apis.StatusFailed, "b")
	listMock.ToUniqueControls()

	assert.Equal(t, 3, listMock.All().Len())
	assert.Equal(t, 1, len(listMock.Passed()))
	assert.Equal(t, 1, len(listMock.PassedExceptions()))
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
