package apis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// makeIS will convert any number of parameters to a []interface{}
func makeIS(v ...interface{}) []interface{} {
	return v
}
func TestCompare(t *testing.T) {
	assert.Equal(t, StatusFailed, Compare(StatusFailed, StatusFailed))
	assert.Equal(t, StatusFailed, Compare(StatusFailed, StatusSkipped))
	assert.Equal(t, StatusFailed, Compare(StatusSkipped, StatusFailed))
	assert.Equal(t, StatusFailed, Compare(StatusPassed, StatusFailed))
	assert.Equal(t, StatusSkipped, Compare(StatusSkipped, StatusPassed))
	assert.Equal(t, StatusPassed, Compare(StatusPassed, StatusPassed))
	assert.Equal(t, StatusUnknown, Compare(StatusUnknown, StatusUnknown))
}

func TestCompareStatusAndSubStatus(t *testing.T) {
	assert.Equal(t, makeIS(StatusFailed, SubStatusUnknown), makeIS(CompareStatusAndSubStatus(StatusFailed, StatusPassed, SubStatusUnknown, SubStatusUnknown)))
	assert.Equal(t, makeIS(StatusFailed, SubStatusUnknown), makeIS(CompareStatusAndSubStatus(StatusFailed, StatusSkipped, SubStatusUnknown, SubStatusConfiguration)))
	assert.Equal(t, makeIS(StatusPassed, SubStatusIrrelevant), makeIS(CompareStatusAndSubStatus(StatusPassed, StatusPassed, SubStatusUnknown, SubStatusIrrelevant)))
	assert.Equal(t, makeIS(StatusPassed, SubStatusException), makeIS(CompareStatusAndSubStatus(StatusPassed, StatusPassed, SubStatusException, SubStatusUnknown)))
	assert.Equal(t, makeIS(StatusSkipped, SubStatusConfiguration), makeIS(CompareStatusAndSubStatus(StatusSkipped, StatusPassed, SubStatusConfiguration, SubStatusUnknown)))
	assert.Equal(t, makeIS(StatusSkipped, SubStatusIntegration), makeIS(CompareStatusAndSubStatus(StatusSkipped, StatusPassed, SubStatusIntegration, SubStatusUnknown)))
	assert.Equal(t, makeIS(StatusSkipped, SubStatusManualReview), makeIS(CompareStatusAndSubStatus(StatusPassed, StatusSkipped, SubStatusUnknown, SubStatusManualReview)))
	assert.Equal(t, makeIS(StatusSkipped, SubStatusRequiresReview), makeIS(CompareStatusAndSubStatus(StatusPassed, StatusSkipped, SubStatusUnknown, SubStatusRequiresReview)))
}

func TestConvertStatusToNewStatus(t *testing.T) {
	tests := []struct {
		name     string
		status   ScanningStatus
		expected ScanningStatus
		sub      ScanningSubStatus
	}{
		{
			name:     "StatusExcluded",
			status:   StatusExcluded,
			expected: StatusPassed,
			sub:      SubStatusException,
		},
		{
			name:     "StatusIrrelevant",
			status:   StatusIrrelevant,
			expected: StatusPassed,
			sub:      SubStatusIrrelevant,
		},
		{
			name:     "StatusPassed",
			status:   StatusPassed,
			expected: StatusPassed,
			sub:      SubStatusUnknown,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actualStatus, actualSub := ConvertStatusToNewStatus(test.status)
			if actualStatus != test.expected {
				t.Errorf("Expected status %s, but got %s", test.expected, actualStatus)
			}
			if actualSub != test.sub {
				t.Errorf("Expected sub status %s, but got %s", test.sub, actualSub)
			}
		})
	}
}
