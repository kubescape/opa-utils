package resourcesresults

import (
	"testing"

	"github.com/armosec/opa-utils/reporthandling/apis"
	"github.com/stretchr/testify/assert"
)

func TestSetGetRuleName(t *testing.T) {
	r := ResourceAssociatedRule{}
	id := "my-rule"
	r.SetName(id)
	assert.Equal(t, id, r.GetName())
}

func TestRuleStatus(t *testing.T) {
	r := mockResourceAssociatedRuleA()
	assert.Equal(t, apis.StatusFailed, r.Status(nil))
	assert.True(t, r.IsFailed(nil))
	assert.False(t, r.IsPassed(nil))
	assert.False(t, r.IsExcluded(nil))
	assert.False(t, r.IsSkipped(nil))

	r2 := mockResourceAssociatedRuleB()
	assert.Equal(t, apis.StatusFailed, r2.Status(nil))
	assert.True(t, r2.IsFailed(nil))
	assert.False(t, r2.IsPassed(nil))
	assert.False(t, r2.IsExcluded(nil))
	assert.False(t, r2.IsSkipped(nil))

}
