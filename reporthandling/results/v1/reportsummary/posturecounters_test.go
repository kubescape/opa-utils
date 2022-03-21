package reportsummary

import (
	"testing"

	"github.com/armosec/opa-utils/reporthandling/apis"
	"github.com/stretchr/testify/assert"
)

func TestNumberOfExcluded(t *testing.T) {
	pcounter := PostureCounters{
		PassedCounter:   4,
		ExcludedCounter: 5,
		FailedCounter:   9,
		IgnoredCounter:  1,
		SkippedCounter:  8,
	}

	assert.Equal(t, 4, pcounter.Passed())
	assert.Equal(t, 5, pcounter.Excluded())
	assert.Equal(t, 9, pcounter.Failed())
	assert.Equal(t, 1, pcounter.Ignored())
	assert.Equal(t, 8, pcounter.Skipped())
	assert.Equal(t, 27, pcounter.All())
}

func TestIncrease(t *testing.T) {
	pcounter := PostureCounters{
		PassedCounter:   0,
		SkippedCounter:  1,
		ExcludedCounter: 2,
		FailedCounter:   2,
	}
	assert.Equal(t, 5, pcounter.All())

	// adding 2 filed
	pcounter.Increase(&apis.StatusInfo{InnerStatus: apis.StatusFailed})
	pcounter.Increase(&apis.StatusInfo{InnerStatus: apis.StatusFailed})
	assert.Equal(t, 4, pcounter.Failed())

	// adding 2 excluded
	pcounter.Increase(&apis.StatusInfo{InnerStatus: apis.StatusExcluded})
	pcounter.Increase(&apis.StatusInfo{InnerStatus: apis.StatusExcluded})
	assert.Equal(t, 4, pcounter.Excluded())

	// adding 4 passed
	pcounter.Increase(&apis.StatusInfo{InnerStatus: apis.StatusPassed})
	pcounter.Increase(&apis.StatusInfo{InnerStatus: apis.StatusPassed})
	pcounter.Increase(&apis.StatusInfo{InnerStatus: apis.StatusPassed})
	pcounter.Increase(&apis.StatusInfo{InnerStatus: apis.StatusPassed})
	assert.Equal(t, 4, pcounter.Passed())

	// adding 3 skipped
	pcounter.Increase(&apis.StatusInfo{InnerStatus: apis.StatusSkipped})
	pcounter.Increase(&apis.StatusInfo{InnerStatus: apis.StatusError})
	pcounter.Increase(&apis.StatusInfo{InnerStatus: apis.StatusIrrelevant})
	assert.Equal(t, 4, pcounter.Skipped())

	assert.Equal(t, 4*4, pcounter.All())
}
