package apis

import (
	"reflect"
	"testing"
)

// GetSupportedSeverities should return a list of supported severities
func TestGetSupportedSeverities(t *testing.T) {
	want := []string{"Low", "Medium", "High", "Critical"}
	got := GetSupportedSeverities()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
