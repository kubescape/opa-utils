package reportsummary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var resourcesCounter = ResourceCounters{}

// =================================== Counters ============================================

func setResourcesCountersMock() {
	resourcesCounter.setNumberOfExcluded(6)
	resourcesCounter.setNumberOfFailed(15)
	resourcesCounter.setNumberOfPassed(7)
	resourcesCounter.setNumberOfSkipped(8)
}

// NumberOfExcluded get the number of excluded resources
func TestNumberOfExcluded(t *testing.T) {
	setResourcesCountersMock()
	assert.Equal(t, 6, resourcesCounter.NumberOfExcluded())
}

// NumberOfPassed get the number of passed resources
func TestNumberOfPassed(t *testing.T) {
	setResourcesCountersMock()
	assert.Equal(t, 7, resourcesCounter.NumberOfPassed())
}

// NumberOfSkipped get the number of skipped resources
func TestNumberOfSkipped(t *testing.T) {
	setResourcesCountersMock()
	assert.Equal(t, 8, resourcesCounter.NumberOfSkipped())
}

// NumberOfFailed get the number of failed resources
func TestNumberOfFailed(t *testing.T) {
	setResourcesCountersMock()
	assert.Equal(t, 15, resourcesCounter.NumberOfFailed())
}

// NumberOfAll get the number of all resources
func TestNumberOfAll(t *testing.T) {
	setResourcesCountersMock()
	assert.Equal(t, 36, resourcesCounter.NumberOfAll())
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
