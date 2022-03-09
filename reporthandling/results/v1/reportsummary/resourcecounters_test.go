package reportsummary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var resourcesCounter = ResourceCounters{}

// =================================== Counters ============================================

func setResourcesCountersMock() {

	resourcesCounter.ExcludedResources = 6
	resourcesCounter.FailedResources = 15
	resourcesCounter.PassedResources = 7
}

// Excluded get the number of excluded resources
func TestExcluded(t *testing.T) {
	setResourcesCountersMock()
	assert.Equal(t, 6, resourcesCounter.Excluded())
}

// Passed get the number of passed resources
func TestPassed(t *testing.T) {
	setResourcesCountersMock()
	assert.Equal(t, 7, resourcesCounter.Passed())
}

// Skipped get the number of skipped resources
func TestSkipped(t *testing.T) {
	setResourcesCountersMock()
	assert.Equal(t, 0, resourcesCounter.Skipped())
}

// Failed get the number of failed resources
func TestFailed(t *testing.T) {
	setResourcesCountersMock()
	assert.Equal(t, 15, resourcesCounter.Failed())
}

// NumberOfAll get the number of all resources
func TestNumberOfAll(t *testing.T) {
	setResourcesCountersMock()
	assert.Equal(t, 28, resourcesCounter.All())
}

// // IsPassed did this control pass
// func TestIsPassed(t *testing.T) {

// }

// // IsFailed did this control fail
// func TestIsFailed(t *testing.T) {
// }

// // IsExcluded is this control excluded
// func TestIsExcluded(t *testing.T) {
// }

// // IsSkipped was this control skipped
// func TestIsSkipped(t *testing.T) {
// }

// // Status get the control status. returns an ScanningStatus object
// func TestStatus(t *testing.T) {
// }
