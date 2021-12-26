package reportsummary

import (
	"testing"

	"github.com/armosec/opa-utils/reporthandling/apis"
	"github.com/stretchr/testify/assert"
)

func TestRuleStatus(t *testing.T) {
	r := mockSummaryDetailsFailed()
	assert.Equal(t, apis.StatusFailed, r.Status(nil))
	assert.True(t, r.IsFailed(nil))
	assert.False(t, r.IsPassed(nil))
	assert.False(t, r.IsExcluded(nil))
	assert.False(t, r.IsSkipped(nil))

	r1 := mockSummaryDetailsExcluded()
	assert.Equal(t, apis.StatusExcluded, r1.Status(nil))
	assert.True(t, r1.IsExcluded(nil))
	assert.False(t, r1.IsFailed(nil))
	assert.False(t, r1.IsPassed(nil))
	assert.False(t, r1.IsSkipped(nil))

	r2 := mockSummaryDetailsPassed()
	assert.Equal(t, apis.StatusPassed, r2.Status(nil))
	assert.True(t, r2.IsPassed(nil))
	assert.False(t, r2.IsFailed(nil))
	assert.False(t, r2.IsExcluded(nil))
	assert.False(t, r2.IsSkipped(nil))

}

func TestRuleListing(t *testing.T) {
	r := mockSummaryDetailsFailed()
	assert.NotEqual(t, 0, len(r.ListAllFrameworks().ListAll()))
	assert.NotEqual(t, 0, len(r.ListFailedFrameworks()))
	assert.NotEqual(t, 0, len(r.ListFailedControls(nil)))
}
