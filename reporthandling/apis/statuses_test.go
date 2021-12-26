package apis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompare(t *testing.T) {
	assert.Equal(t, StatusFailed, Compare(StatusFailed, StatusFailed))
	assert.Equal(t, StatusFailed, Compare(StatusFailed, StatusExcluded))
	assert.Equal(t, StatusFailed, Compare(StatusExcluded, StatusFailed))
	assert.Equal(t, StatusFailed, Compare(StatusPassed, StatusFailed))
	assert.Equal(t, StatusExcluded, Compare(StatusExcluded, StatusPassed))
	assert.Equal(t, StatusPassed, Compare(StatusPassed, StatusPassed))
	assert.Equal(t, StatusPassed, Compare(StatusSkipped, StatusPassed))
}
