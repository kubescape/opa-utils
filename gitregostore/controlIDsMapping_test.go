package gitregostore

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

func TestGetNewControlID(t *testing.T) {
	controlIDS_tests := []struct {
		name        string
		controlID   string
		expectedRes string
	}{
		{
			name:        "ControlID_exists_in_mapping_uppercase",
			controlID:   "CIS-1.1.1",
			expectedRes: "C-0091",
		},
		{
			name:        "ControlID_exists_in_mapping_lowercase",
			controlID:   "cis-1.1.1",
			expectedRes: "C-0091",
		},
		{
			name:        "ControlID_doestexists_in_mapping",
			controlID:   "IDontExist",
			expectedRes: "IDontExist",
		},
	}

	for _, tt := range controlIDS_tests {
		t.Run(tt.name, func(t *testing.T) {
			res := getNewControlID(tt.controlID)
			assert.Equal(t, tt.expectedRes, res)
		})
	}
}
