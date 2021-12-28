package resourcesresults

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetGetRuleName(t *testing.T) {
	r := ResourceAssociatedRule{}
	id := "my-rule"
	r.SetName(id)
	assert.Equal(t, id, r.GetName())
}
