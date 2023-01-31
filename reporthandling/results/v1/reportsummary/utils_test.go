package reportsummary

import (
	"testing"

	"github.com/kubescape/opa-utils/reporthandling/apis"
	"github.com/stretchr/testify/assert"
)

func TestCalculateStatus(t *testing.T) {
	passed := mockResourceCountersPass()
	failed := mockResourceCountersFailPass()
	skipped := mockResourceCountersSkipped()

	assert.Equal(t, apis.StatusPassed, calculateStatus(passed))
	assert.Equal(t, apis.StatusFailed, calculateStatus(failed))
	assert.Equal(t, apis.StatusSkipped, calculateStatus(skipped))

}
