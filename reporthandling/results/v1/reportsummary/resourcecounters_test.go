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
	resourcesCounter.PassedExceptionResources = 6
	resourcesCounter.PassedIrrelevantResources = 1
	resourcesCounter.SkippedConfigurationResources = 2
	resourcesCounter.SkippedIntegrationResources = 3
	resourcesCounter.SkippedManualReviewResources = 4
	resourcesCounter.SkippedRequiresReviewResources = 5
}

// Excluded get the number of excluded resources
func TestExcluded(t *testing.T) {
	setResourcesCountersMock()
	assert.Equal(t, 6, resourcesCounter.PassedExceptions())
}

// Passed get the number of passed resources
func TestPassed(t *testing.T) {
	setResourcesCountersMock()
	assert.Equal(t, 7, resourcesCounter.Passed())
}

// PassedExceptions get the number of passed exception resources
func TestPassedExceptions(t *testing.T) {
	setResourcesCountersMock()
	assert.Equal(t, 6, resourcesCounter.PassedExceptions())
}

// PassedIrrelevant get the number of passed irrelevant resources
func TestPassedIrrelevant(t *testing.T) {
	setResourcesCountersMock()
	assert.Equal(t, 1, resourcesCounter.PassedIrrelevant())
}

// SkippedIntegration get the number of skipped integration resources
func TestSkippedIntegration(t *testing.T) {
	setResourcesCountersMock()
	assert.Equal(t, 3, resourcesCounter.SkippedIntegration())
}

// SkippedConfiguration get the number of skipped configuration resources
func TestSkippedConfiguration(t *testing.T) {
	setResourcesCountersMock()
	assert.Equal(t, 2, resourcesCounter.SkippedConfiguration())
}

// SkippedRequiresReview get the number of skipped requires review resources
func TestSkippedRequiresReview(t *testing.T) {
	setResourcesCountersMock()
	assert.Equal(t, 5, resourcesCounter.SkippedRequiresReview())
}

// SkippedManualReview get the number of skipped manual review resources
func TestSkippedManualReview(t *testing.T) {
	setResourcesCountersMock()
	assert.Equal(t, 4, resourcesCounter.SkippedManualReview())
}

// Failed get the number of failed resources
func TestFailed(t *testing.T) {
	setResourcesCountersMock()
	assert.Equal(t, 15, resourcesCounter.Failed())
}

// NumberOfAll get the number of all resources
func TestNumberOfAll(t *testing.T) {
	setResourcesCountersMock()
	assert.Equal(t, 43, resourcesCounter.All())
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
