package reportsummary

import (
	"testing"

	"github.com/kubescape/opa-utils/reporthandling/apis"
	"github.com/stretchr/testify/assert"
)

func TestNumberOfExcluded(t *testing.T) {
	pcounter := PostureCounters{
		PassedCounter:  1,
		SkippedCounter: 7,
		FailedCounter:  9,
	}

	assert.Equal(t, 1, pcounter.Passed())
	assert.Equal(t, 9, pcounter.Failed())
	assert.Equal(t, 7, pcounter.Skipped())
	assert.Equal(t, 17, pcounter.All())
}

func TestIncrease(t *testing.T) {
	pcounter := PostureCounters{
		PassedCounter:  1,
		SkippedCounter: 7,
		FailedCounter:  9,
	}
	assert.Equal(t, 17, pcounter.All())

	// adding 2 filed
	pcounter.Increase(&apis.StatusInfo{InnerStatus: apis.StatusFailed})
	pcounter.Increase(&apis.StatusInfo{InnerStatus: apis.StatusFailed})
	assert.Equal(t, 11, pcounter.Failed())

	// adding 2 skipped
	pcounter.Increase(&apis.StatusInfo{InnerStatus: apis.StatusSkipped})
	pcounter.Increase(&apis.StatusInfo{InnerStatus: apis.StatusSkipped})
	assert.Equal(t, 9, pcounter.Skipped())

	// adding 4 passed
	pcounter.Increase(&apis.StatusInfo{InnerStatus: apis.StatusPassed})
	pcounter.Increase(&apis.StatusInfo{InnerStatus: apis.StatusPassed})
	pcounter.Increase(&apis.StatusInfo{InnerStatus: apis.StatusPassed})
	pcounter.Increase(&apis.StatusInfo{InnerStatus: apis.StatusPassed})
	assert.Equal(t, 5, pcounter.Passed())

	assert.Equal(t, 25, pcounter.All())
}
