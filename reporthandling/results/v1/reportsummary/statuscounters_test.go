package reportsummary

import (
	"testing"

	"github.com/kubescape/opa-utils/reporthandling/apis"
	"github.com/stretchr/testify/assert"
)

var resourcesCounter = StatusCounters{}

// =================================== Counters ============================================

func setResourcesCountersMock() {
	resourcesCounter.PassedResources = 7
	resourcesCounter.FailedResources = 15
	resourcesCounter.SkippedResources = 5
}

func TestSet(t *testing.T) {
	rc := StatusCounters{}
	summaryDetails := MockSummaryDetails()
	rc.Set(summaryDetails.ListResourcesIDs(nil))
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

func TestStatusCounters_Increase(t *testing.T) {
	tests := []struct {
		resourceCounters *StatusCounters
		name             string
		status           apis.ScanningStatus
		expectedPassed   int
		expectedFailed   int
		expectedSkipped  int
		expectedExcluded int
	}{
		{
			name:   "Test passed status",
			status: apis.StatusPassed,
			resourceCounters: &StatusCounters{
				FailedResources:  1,
				SkippedResources: 2,
				PassedResources:  3,
			},
			expectedFailed:  1,
			expectedSkipped: 2,
			expectedPassed:  4,
		},
		{
			name:   "Test failed status",
			status: apis.StatusFailed,
			resourceCounters: &StatusCounters{
				FailedResources:  1,
				SkippedResources: 2,
				PassedResources:  3,
			},
			expectedFailed:  2,
			expectedSkipped: 2,
			expectedPassed:  3,
		},
		{
			name:   "Test skipped status",
			status: apis.StatusSkipped,
			resourceCounters: &StatusCounters{
				FailedResources:  1,
				SkippedResources: 2,
				PassedResources:  3,
			},
			expectedFailed:  1,
			expectedSkipped: 3,
			expectedPassed:  3,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			status := &apis.StatusInfo{InnerStatus: test.status}
			test.resourceCounters.Increase(status)
			if test.resourceCounters.PassedResources != test.expectedPassed {
				t.Errorf("Expected PassedResources to be %d, but got %d", test.expectedPassed, test.resourceCounters.PassedResources)
			}
			if test.resourceCounters.FailedResources != test.expectedFailed {
				t.Errorf("Expected FailedResources to be %d, but got %d", test.expectedFailed, test.resourceCounters.FailedResources)
			}
			if test.resourceCounters.SkippedResources != test.expectedSkipped {
				t.Errorf("Expected SkippedResources to be %d, but got %d", test.expectedSkipped, test.resourceCounters.SkippedResources)
			}
		})
	}
}

func TestSubStatusCounters_Increase(t *testing.T) {
	tests := []struct {
		subStatusCounters *SubStatusCounters
		name              string
		status            apis.ScanningStatus
		subStatus         apis.ScanningSubStatus
		expectedIgnored   int
	}{
		{
			name:      "Test ignored and passed status",
			status:    apis.StatusPassed,
			subStatus: apis.SubStatusException,
			subStatusCounters: &SubStatusCounters{
				IgnoredResources: 1,
			},
			expectedIgnored: 2,
		},
		{
			name:      "Test ignored and failed status",
			status:    apis.StatusFailed,
			subStatus: apis.SubStatusException,
			subStatusCounters: &SubStatusCounters{
				IgnoredResources: 1,
			},
			expectedIgnored: 1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			status := &apis.StatusInfo{InnerStatus: test.status, SubStatus: test.subStatus}
			test.subStatusCounters.Increase(status)
			if test.subStatusCounters.IgnoredResources != test.expectedIgnored {
				t.Errorf("Expected IgnoredResources to be %d, but got %d", test.expectedIgnored, test.subStatusCounters.IgnoredResources)
			}
		})
	}
}
