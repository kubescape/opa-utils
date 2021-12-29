package reportsummary

import (
	"testing"

	"github.com/armosec/opa-utils/reporthandling/apis"
	"github.com/stretchr/testify/assert"
)

func TestRuleStatus(t *testing.T) {
	r := mockSummaryDetailsFailed()
	r.CalculateStatus()

	assert.Equal(t, apis.StatusFailed, r.GetStatus().Status())
	assert.True(t, r.GetStatus().IsFailed())
	assert.False(t, r.GetStatus().IsPassed())
	assert.False(t, r.GetStatus().IsExcluded())
	assert.False(t, r.GetStatus().IsSkipped())

	r1 := mockSummaryDetailsExcluded()
	r1.CalculateStatus()

	assert.Equal(t, apis.StatusExcluded, r1.GetStatus().Status())
	assert.True(t, r1.GetStatus().IsExcluded())
	assert.False(t, r1.GetStatus().IsFailed())
	assert.False(t, r1.GetStatus().IsPassed())
	assert.False(t, r1.GetStatus().IsSkipped())

	r2 := mockSummaryDetailsPassed()
	r2.CalculateStatus()

	assert.Equal(t, apis.StatusPassed, r2.GetStatus().Status())
	assert.True(t, r2.GetStatus().IsPassed())
	assert.False(t, r2.GetStatus().IsFailed())
	assert.False(t, r2.GetStatus().IsExcluded())
	assert.False(t, r2.GetStatus().IsSkipped())

}

func TestRuleListing(t *testing.T) {
	r := mockSummaryDetailsFailed()
	assert.NotEqual(t, 0, len(r.ListFrameworksNames().All()))
	assert.NotEqual(t, 0, len(r.ListFrameworksNames().Failed()))
	assert.NotEqual(t, 0, len(r.ListControlsNames().Failed()))
}
