package reportsummary

import (
	"testing"

	"github.com/kubescape/opa-utils/reporthandling/apis"
	"github.com/stretchr/testify/assert"
)

func TestCalculateStatus(t *testing.T) {
	passed := mockStatusCountersPass()
	failed := mockStatusCountersFailPass()
	skipped := mockStatusCountersSkipped()

	assert.Equal(t, apis.StatusPassed, calculateStatus(passed))
	assert.Equal(t, apis.StatusFailed, calculateStatus(failed))
	assert.Equal(t, apis.StatusSkipped, calculateStatus(skipped))

}
