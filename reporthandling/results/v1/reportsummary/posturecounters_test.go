package reportsummary

import (
	"testing"

	"github.com/kubescape/opa-utils/reporthandling/apis"
	"github.com/stretchr/testify/assert"
)

func TestNumberOfExcluded(t *testing.T) {
	pcounter := PostureCounters{
		PassedCounter:                1,
		PassedExceptionCounter:       2,
		PassedIrrelevantCounter:      3,
		SkippedConfigurationCounter:  4,
		SkippedIntegrationCounter:    5,
		SkippedRequiresReviewCounter: 6,
		SkippedManualReviewCounter:   7,
		FailedCounter:                9,
	}

	assert.Equal(t, 1, pcounter.Passed())
	assert.Equal(t, 2, pcounter.PassedExceptions())
	assert.Equal(t, 3, pcounter.PassedIrrelevant())
	assert.Equal(t, 9, pcounter.Failed())
	assert.Equal(t, 4, pcounter.SkippedConfiguration())
	assert.Equal(t, 5, pcounter.SkippedIntegration())
	assert.Equal(t, 7, pcounter.SkippedManualReview())
	assert.Equal(t, 6, pcounter.SkippedRequiresReview())
	assert.Equal(t, 37, pcounter.All())
}

func TestIncrease(t *testing.T) {
	pcounter := PostureCounters{
		PassedCounter:                1,
		PassedExceptionCounter:       2,
		PassedIrrelevantCounter:      3,
		SkippedConfigurationCounter:  4,
		SkippedIntegrationCounter:    5,
		SkippedRequiresReviewCounter: 6,
		SkippedManualReviewCounter:   7,
		FailedCounter:                9,
	}
	assert.Equal(t, 37, pcounter.All())

	// adding 2 filed
	pcounter.Increase(&apis.StatusInfo{InnerStatus: apis.StatusFailed})
	pcounter.Increase(&apis.StatusInfo{InnerStatus: apis.StatusFailed})
	assert.Equal(t, 11, pcounter.Failed())

	// adding 2 excluded
	pcounter.Increase(&apis.StatusInfo{InnerStatus: apis.SubStatusException})
	pcounter.Increase(&apis.StatusInfo{InnerStatus: apis.SubStatusException})
	assert.Equal(t, 4, pcounter.PassedExceptions())

	// adding 4 passed
	pcounter.Increase(&apis.StatusInfo{InnerStatus: apis.StatusPassed})
	pcounter.Increase(&apis.StatusInfo{InnerStatus: apis.StatusPassed})
	pcounter.Increase(&apis.StatusInfo{InnerStatus: apis.StatusPassed})
	pcounter.Increase(&apis.StatusInfo{InnerStatus: apis.StatusPassed})
	assert.Equal(t, 5, pcounter.Passed())

	// adding 3 skippedConfiguration
	pcounter.Increase(&apis.StatusInfo{InnerStatus: apis.SubStatusConfiguration})
	pcounter.Increase(&apis.StatusInfo{InnerStatus: apis.SubStatusConfiguration})
	pcounter.Increase(&apis.StatusInfo{InnerStatus: apis.SubStatusConfiguration})
	assert.Equal(t, 7, pcounter.SkippedConfiguration())

	assert.Equal(t, 48, pcounter.All())
}
