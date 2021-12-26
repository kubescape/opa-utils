package resourcesresults

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetGetResourceID(t *testing.T) {
	r := Result{}
	id := "my/id"
	r.SetResourceID(id)
	assert.Equal(t, id, r.GetResourceID())
}
