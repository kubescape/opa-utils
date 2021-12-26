package resourcesresults

import (
	"testing"

	"github.com/armosec/opa-utils/reporthandling/apis"
	"github.com/stretchr/testify/assert"
)

func TestSetGetResourceID(t *testing.T) {
	r := Result{}
	id := "my/id"
	r.SetResourceID(id)
	assert.Equal(t, id, r.GetResourceID())
}

func TestResultStatus(t *testing.T) {
	r := mockResultFailed()
	assert.Equal(t, apis.StatusFailed, r.Status(nil))
	assert.True(t, r.IsFailed(nil))
	assert.False(t, r.IsPassed(nil))
	assert.False(t, r.IsExcluded(nil))
	assert.False(t, r.IsSkipped(nil))

	r2 := mockResultPassed()
	assert.Equal(t, apis.StatusPassed, r2.Status(nil))
	assert.True(t, r2.IsPassed(nil))
	assert.False(t, r2.IsFailed(nil))
	assert.False(t, r2.IsExcluded(nil))
	assert.False(t, r2.IsSkipped(nil))

}

func TestResultList(t *testing.T) {
	r := mockResultFailed()
	assert.NotEqual(t, 0, len(r.ListAllControls(nil).ListAll()))
	assert.NotEqual(t, 0, len(r.ListFailedControls(nil)))
	assert.NotEqual(t, 0, len(r.ListPassedControls(nil)))
	assert.Equal(t, 0, len(r.ListExcludedControls(nil)))

	r3 := mockResultPassed()
	assert.NotEqual(t, 0, len(r3.ListAllControls(nil).ListAll()))
	assert.NotEqual(t, 0, len(r3.ListPassedControls(nil)))
	assert.Equal(t, 0, len(r3.ListExcludedControls(nil)))
	assert.Equal(t, 0, len(r3.ListFailedControls(nil)))
}
