package reporthandling

import (
	"reflect"
	"testing"

	"github.com/kubescape/k8s-interface/workloadinterface"
	"github.com/stretchr/testify/assert"
)

func TestResourceImplementsIMetadata(t *testing.T) {
	imetadata := reflect.TypeOf((*workloadinterface.IMetadata)(nil)).Elem()
	resource := NewResource(map[string]interface{}{})
	assert.True(t, reflect.TypeOf(resource).Implements(imetadata), "Resource does not implement IMetadata")
}
