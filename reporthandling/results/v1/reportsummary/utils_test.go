package reportsummary

import (
	"testing"

	"github.com/armosec/opa-utils/reporthandling/apis"
	"github.com/stretchr/testify/assert"
)

func TestCalculateStatus(t *testing.T) {
	exclude := mockResourceCountersExcludePass()
	passed := mockResourceCountersPass()
	failed := mockResourceCountersExcludeFailPass()
	skipped := mockResourceCountersSkipped()

	assert.Equal(t, apis.StatusExcluded, calculateStatus(exclude))
	assert.Equal(t, apis.StatusPassed, calculateStatus(passed))
	assert.Equal(t, apis.StatusFailed, calculateStatus(failed))
	assert.Equal(t, apis.StatusIrrelevant, calculateStatus(skipped))

}
