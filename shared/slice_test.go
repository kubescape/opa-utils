package shared

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringInSlice(t *testing.T) {
	assert.True(t, StringInSlice([]string{"a"}, "a"))
	assert.True(t, StringInSlice([]string{"a", "b", "c"}, "a"))
	assert.True(t, StringInSlice([]string{"a", "b", "c"}, "b"))
	assert.True(t, StringInSlice([]string{"a", "b", "c"}, "c"))
	assert.True(t, StringInSlice([]string{"a", "a"}, "a"))
	assert.False(t, StringInSlice([]string{"a", "b", "c"}, "d"))
	assert.False(t, StringInSlice([]string{""}, "a"))
	assert.False(t, StringInSlice([]string{"a"}, ""))
}

func TestMapStringToSlice(t *testing.T) {
	assert.ElementsMatch(t, MapStringToSlice(map[string]interface{}{"a": nil}), []string{"a"})
	assert.ElementsMatch(t, MapStringToSlice(map[string]interface{}{"a": nil, "b": nil}), []string{"a", "b"})
	assert.ElementsMatch(t, MapStringToSlice(nil), []string{})
	assert.ElementsMatch(t, MapStringToSlice(map[string]interface{}{}), []string{})
}

func TestSliceStringToUnique(t *testing.T) {
	assert.ElementsMatch(t, SliceStringToUnique([]string{"a"}), []string{"a"})
	assert.ElementsMatch(t, SliceStringToUnique([]string{}), []string{})
	assert.ElementsMatch(t, SliceStringToUnique([]string{"a", "b", "b", "a"}), []string{"a", "b"})
}
