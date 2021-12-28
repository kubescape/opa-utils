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
	assert.Equal(t, apis.StatusFailed, r.GetStatus(nil).Status())
	assert.True(t, r.GetStatus(nil).IsFailed())
	assert.False(t, r.GetStatus(nil).IsPassed())
	assert.False(t, r.GetStatus(nil).IsExcluded())
	assert.False(t, r.GetStatus(nil).IsSkipped())

	r2 := mockResultPassed()
	assert.Equal(t, apis.StatusPassed, r2.GetStatus(nil).Status())
	assert.True(t, r2.GetStatus(nil).IsPassed())
	assert.False(t, r2.GetStatus(nil).IsFailed())
	assert.False(t, r2.GetStatus(nil).IsExcluded())
	assert.False(t, r2.GetStatus(nil).IsSkipped())

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
