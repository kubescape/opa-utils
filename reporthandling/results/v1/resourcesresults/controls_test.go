package resourcesresults

import (
	"testing"

	"github.com/armosec/opa-utils/reporthandling/apis"
	"github.com/stretchr/testify/assert"
)

func TestSetGetControlID(t *testing.T) {
	r := ResourceAssociatedControl{}
	id := "C-0078"
	r.SetID(id)
	assert.Equal(t, id, r.GetID())
}

func TestControlStatus(t *testing.T) {
	r := mockResourceAssociatedControl0087Failed()
	assert.Equal(t, apis.StatusFailed, r.Status(nil))
	assert.True(t, r.IsFailed(nil))
	assert.False(t, r.IsPassed(nil))
	assert.False(t, r.IsExcluded(nil))
	assert.False(t, r.IsSkipped(nil))

	r2 := mockResourceAssociatedControl0089Passed()
	assert.Equal(t, apis.StatusPassed, r2.Status(nil))
	assert.True(t, r2.IsPassed(nil))
	assert.False(t, r2.IsFailed(nil))
	assert.False(t, r2.IsExcluded(nil))
	assert.False(t, r2.IsSkipped(nil))

}
