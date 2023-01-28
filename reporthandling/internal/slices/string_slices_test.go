package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type (
	trimArgs struct {
		Origin         []string
		TrimFrom       []string
		Expected       []string
		ExpectedStable []string
	}

	trimFixture struct {
		Name string
		Args trimArgs
	}
)

// Clone the Origin slice to be used in parallel tests that modify their input in-place.
func (a trimArgs) Clone() trimArgs {
	if a.Origin == nil {
		return a
	}

	clone := trimArgs{
		Origin:         make([]string, len(a.Origin)),
		TrimFrom:       a.TrimFrom,
		Expected:       a.Expected,
		ExpectedStable: a.ExpectedStable,
	}
	copy(clone.Origin, a.Origin)

	return clone
}

func TestTrimAndTrimUnique(t *testing.T) {
	t.Parallel()

	for _, toPin := range trimTestCases() {
		tt := toPin
		ts := tt
		tu := tt
		tsu := tt
		// the Trim* functions modify their input slice: we need a deep-copy of the Origin argument
		ts.Args = tt.Args.Clone()
		tu.Args = tt.Args.Clone()
		tsu.Args = tt.Args.Clone()

		// exercise our variations around trimming slices...
		t.Run("Trim with "+tt.Name, func(t *testing.T) {
			t.Parallel()

			dd := Trim(tt.Args.Origin, tt.Args.TrimFrom)
			assert.Equal(t, tt.Args.Expected, dd)
			assert.Equalf(t, len(tt.Args.Expected), cap(dd),
				"expected capacity to be truncated but got %d", cap(dd),
			)
		})

		t.Run("TrimStable with "+ts.Name, func(t *testing.T) {
			t.Parallel()

			dd := TrimStable(ts.Args.Origin, ts.Args.TrimFrom)
			var expected []string
			if ts.Args.ExpectedStable != nil {
				// with the stable version, the expected slice comes in a different order
				expected = ts.Args.ExpectedStable
			} else {
				expected = ts.Args.Expected
			}

			assert.Equal(t, expected, dd)
			assert.Equalf(t, len(expected), cap(dd),
				"expected capacity to be truncated but got %d", cap(dd),
			)
		})

		t.Run("TrimUnique with "+tu.Name, func(t *testing.T) {
			t.Parallel()

			dd := TrimUnique(tu.Args.Origin, tu.Args.TrimFrom)
			assert.Equal(t, tu.Args.Expected, dd)
			assert.Equalf(t, len(tu.Args.Expected), cap(dd),
				"expected capacity to be truncated but got %d", cap(dd),
			)
		})

		t.Run("TrimStableUnique with "+tsu.Name, func(t *testing.T) {
			t.Parallel()

			dd := TrimStableUnique(tsu.Args.Origin, tsu.Args.TrimFrom)
			var expected []string
			if tsu.Args.ExpectedStable != nil {
				// with the stable version, the expected slice comes in a different order
				expected = tsu.Args.ExpectedStable
			} else {
				expected = tsu.Args.Expected
			}

			assert.Equal(t, expected, dd)
			assert.Equalf(t, len(expected), cap(dd),
				"expected capacity to be truncated but got %d", cap(dd),
			)
		})
	}

	t.Run("TrimStable should trim but not dedupe", func(t *testing.T) {
		assert.EqualValues(t,
			[]string{"c", "c"},
			TrimStable([]string{"c", "a", "b", "c", "b"}, []string{"a", "b", "e"}),
		)
	})

	t.Run("TrimStableUnique should dedupe and trim", func(t *testing.T) {
		assert.EqualValues(t,
			[]string{"c"},
			TrimStableUnique([]string{"c", "a", "b", "c", "b"}, []string{"a", "b", "e"}),
		)
	})

	t.Run("Trim should trim but not dedupe", func(t *testing.T) {
		assert.EqualValues(t,
			[]string{"c", "c"},
			Trim([]string{"c", "a", "b", "c", "b"}, []string{"a", "b", "e"}),
		)
	})

	t.Run("TrimUnique should dedupe and trim", func(t *testing.T) {
		assert.EqualValues(t,
			[]string{"c"},
			TrimStableUnique([]string{"c", "a", "b", "c", "b"}, []string{"a", "b", "e"}),
		)
	})
}

func TestUniqueSlice(t *testing.T) {
	t.Parallel()

	t.Run("should yield unique strings", func(t *testing.T) {
		t.Parallel()

		uniques := UniqueStrings([]string{"B", "B", "A", "A", "B", "C", "B", "A"})

		require.EqualValues(t, []string{"B", "A", "C"}, uniques)
		require.Equal(t, 3, cap(uniques),
			"expected capacity to be truncated but got %d", cap(uniques),
		)
	})

	t.Run("should yield empty", func(t *testing.T) {
		t.Parallel()

		require.Empty(t, UniqueStrings([]string{}))
	})

	t.Run("should yield nil", func(t *testing.T) {
		t.Parallel()

		require.Empty(t, UniqueStrings(nil))
	})
}

func TestStringInslice(t *testing.T) {
	t.Parallel()

	require.True(t, StringInSlice([]string{"A", "B", "C"}, "B"))
	require.False(t, StringInSlice([]string{"A", "B", "C"}, "D"))
	require.False(t, StringInSlice([]string{}, "D"))
	require.False(t, StringInSlice(nil, "D"))
}

func trimTestCases() []trimFixture {
	return []trimFixture{
		{
			Name: "trim from beginning of slice (exhibits swapped elements)",
			Args: trimArgs{
				Origin:         []string{"a", "b", "c"},
				TrimFrom:       []string{"a"},
				Expected:       []string{"c", "b"},
				ExpectedStable: []string{"b", "c"},
			},
		},
		{
			Name: "trim from middle of slice",
			Args: trimArgs{
				Origin:   []string{"a", "b", "c"},
				TrimFrom: []string{"b"},
				Expected: []string{"a", "c"},
			},
		},
		{
			Name: "trim from end of slice",
			Args: trimArgs{
				Origin:   []string{"a", "b", "c"},
				TrimFrom: []string{"c"},
				Expected: []string{"a", "b"},
			},
		},
		{
			Name: "do nothing",
			Args: trimArgs{
				Origin:   []string{"a", "b", "c"},
				TrimFrom: []string{"d"},
				Expected: []string{"a", "b", "c"},
			},
		},
		{
			Name: "trim all",
			Args: trimArgs{
				Origin:   []string{"a", "b", "c"},
				TrimFrom: []string{"a", "b", "c"},
				Expected: []string{},
			},
		},
		{
			Name: "trimFrom larger",
			Args: trimArgs{
				Origin:   []string{"a", "b", "c"},
				TrimFrom: []string{"a", "b", "e", "d"},
				Expected: []string{"c"},
			},
		},
		{
			Name: "trim all not sorted",
			Args: trimArgs{
				Origin:   []string{"c", "a", "b"},
				TrimFrom: []string{"a", "b", "c"},
				Expected: []string{},
			},
		},
		{
			Name: "nothing to do (1)",
			Args: trimArgs{
				Origin:   []string{},
				TrimFrom: []string{"d"},
				Expected: []string{},
			},
		},
		{
			Name: "nothing to do (2)",
			Args: trimArgs{
				Origin:   []string{"a"},
				TrimFrom: []string{},
				Expected: []string{"a"},
			},
		},
		{
			Name: "with nil origin",
			Args: trimArgs{
				Origin:   nil,
				TrimFrom: []string{},
				Expected: nil,
			},
		},
		{
			Name: "with nil trim list",
			Args: trimArgs{
				Origin:   []string{"a", "b"},
				TrimFrom: []string{},
				Expected: []string{"a", "b"},
			},
		},
	}
}
