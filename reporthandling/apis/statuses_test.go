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
}

func TestCompareStatusAndSubStatus(t *testing.T) {
	assert.Equal(t, makeIS(StatusFailed, StatusUnknown), makeIS(CompareStatusAndSubStatus(StatusFailed, StatusUnknown, StatusPassed, StatusUnknown)))
	assert.Equal(t, makeIS(StatusFailed, StatusUnknown), makeIS(CompareStatusAndSubStatus(StatusFailed, StatusUnknown, StatusSkipped, SubStatusConfiguration)))
	assert.Equal(t, makeIS(StatusPassed, SubStatusIrrelevant), makeIS(CompareStatusAndSubStatus(StatusPassed, StatusUnknown, StatusPassed, SubStatusIrrelevant)))
	assert.Equal(t, makeIS(StatusPassed, SubStatusException), makeIS(CompareStatusAndSubStatus(StatusPassed, SubStatusException, StatusPassed, StatusUnknown)))
	assert.Equal(t, makeIS(StatusSkipped, SubStatusConfiguration), makeIS(CompareStatusAndSubStatus(StatusSkipped, SubStatusConfiguration, StatusPassed, StatusUnknown)))
	assert.Equal(t, makeIS(StatusSkipped, SubStatusIntegration), makeIS(CompareStatusAndSubStatus(StatusSkipped, SubStatusIntegration, StatusPassed, StatusUnknown)))
	assert.Equal(t, makeIS(StatusSkipped, SubStatusManualReview), makeIS(CompareStatusAndSubStatus(StatusPassed, StatusUnknown, StatusSkipped, SubStatusManualReview)))
	assert.Equal(t, makeIS(StatusSkipped, SubStatusRequiresReview), makeIS(CompareStatusAndSubStatus(StatusPassed, StatusUnknown, StatusSkipped, SubStatusRequiresReview)))
}
