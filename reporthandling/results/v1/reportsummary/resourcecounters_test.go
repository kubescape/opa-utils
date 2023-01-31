package reportsummary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var resourcesCounter = ResourceCounters{}

// =================================== Counters ============================================

func setResourcesCountersMock() {
	resourcesCounter.PassedResources = 7
	resourcesCounter.FailedResources = 15
	resourcesCounter.SkippedResources = 5
}

func TestSet(t *testing.T) {
	rc := ResourceCounters{}
	summaryDetails := MockSummaryDetails()
	rc.Set(summaryDetails.ListResourcesIDs())
	assert.Equal(t, 1, rc.SkippedResources)
	assert.Equal(t, 1, rc.PassedResources)
	assert.Equal(t, 2, rc.FailedResources)
}

// Excluded get the number of skipped resources
func TestExcluded(t *testing.T) {
	setResourcesCountersMock()
	assert.Equal(t, 5, resourcesCounter.Skipped())
}

// Passed get the number of passed resources
func TestPassed(t *testing.T) {
	setResourcesCountersMock()
	assert.Equal(t, 7, resourcesCounter.Passed())
}

// Failed get the number of failed resources
func TestFailed(t *testing.T) {
	setResourcesCountersMock()
	assert.Equal(t, 15, resourcesCounter.Failed())
}

// NumberOfAll get the number of all resources
func TestNumberOfAll(t *testing.T) {
	setResourcesCountersMock()
	assert.Equal(t, 27, resourcesCounter.All())
}
