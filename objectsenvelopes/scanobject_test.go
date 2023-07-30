package objectsenvelopes

import (
	"encoding/json"
	"testing"

	"github.com/kubescape/k8s-interface/workloadinterface"
	"github.com/stretchr/testify/assert"
)

func TestNewScanObject(t *testing.T) {
	// Test nil input
	assert.Nil(t, NewScanObject(nil))

	// Test valid input
	object := map[string]interface{}{
		"kind":       "Pod",
		"apiVersion": "v1",
		"metadata": map[string]interface{}{
			"name":      "test-pod",
			"namespace": "test-namespace",
		},
	}
	scanObj := NewScanObject(object)
	assert.NotNil(t, scanObj)
	assert.Equal(t, "Pod", scanObj.GetKind())
	assert.Equal(t, "v1", scanObj.GetApiVersion())
	assert.Equal(t, "test-pod", scanObj.GetName())
	assert.Equal(t, "test-namespace", scanObj.GetNamespace())
	assert.Equal(t, workloadinterface.TypeWorkloadObject, scanObj.GetObjectType())

	// Test unsupported input
	unsupportedObject := map[string]interface{}{
		"unknown": "unknown",
	}
	assert.Nil(t, NewScanObject(unsupportedObject))
}

func TestScanObjectMarshal(t *testing.T) {
	scanObj := ScanObject{
		Kind:       "Pod",
		ApiVersion: "v1",
		Metadata: ScanObjectMetadata{
			Name:      "test-pod",
			Namespace: "test-namespace",
		},
	}

	// test json marshal
	jsonEnc, err := json.Marshal(scanObj)
	expectedJsonStr := `{"apiVersion":"v1","kind":"Pod","metadata":{"name":"test-pod","namespace":"test-namespace"}}`
	assert.NoError(t, err)
	assert.Equal(t, expectedJsonStr, string(jsonEnc))
}

func TestScanObjectUnmarshal(t *testing.T) {
	var scanObj ScanObject
	jsonStr := `{"apiVersion":"v1","kind":"Node","metadata":{"name":"node-1"},"field":"value"}`
	err := json.Unmarshal([]byte(jsonStr), &scanObj)
	assert.NoError(t, err)
	assert.Equal(t, "Node", scanObj.GetKind())
	assert.Equal(t, "node-1", scanObj.GetName())
	assert.Equal(t, "", scanObj.GetNamespace())
}

func TestScanObjectSetters(t *testing.T) {
	scanObj := ScanObject{}

	scanObj.SetKind("Node")
	assert.Equal(t, "Node", scanObj.Kind)

	scanObj.SetApiVersion("v2")
	assert.Equal(t, "v2", scanObj.ApiVersion)

	scanObj.SetName("node-1")
	assert.Equal(t, "node-1", scanObj.Metadata.Name)

	scanObj.SetNamespace("namespace")
	assert.Equal(t, "namespace", scanObj.Metadata.Namespace)

	// test invalid objects - should remain unchanged
	scanObj.SetObject(map[string]interface{}{
		"invalid": "object",
	})
	scanObj.SetObject(map[string]interface{}{
		"apiVersion": "v3",
	})
	scanObj.SetObject(map[string]interface{}{
		"apiVersion": "v3",
		"kind":       "Secret",
	})
	scanObj.SetObject(map[string]interface{}{
		"apiVersion": "v3",
		"kind":       "Secret",
		"metadata":   "invalid_type",
	})
	scanObj.SetObject(map[string]interface{}{
		"apiVersion": "v3",
		"kind":       "Secret",
		"metadata": map[string]interface{}{
			"namespace": "ns-1",
		},
	})
	assert.Equal(t, "Node", scanObj.Kind)

	scanObj.SetObject(map[string]interface{}{
		"kind":       "Pod",
		"apiVersion": "v1",
		"metadata": map[string]interface{}{
			"name": "test-pod",
		},
	})
	assert.Equal(t, "Pod", scanObj.Kind)
	assert.Equal(t, "v1", scanObj.ApiVersion)
	assert.Equal(t, "test-pod", scanObj.Metadata.Name)
}

func TestScanObjectGetters(t *testing.T) {
	scanObj := ScanObject{
		Kind:       "Pod",
		ApiVersion: "v1",
		Metadata: ScanObjectMetadata{
			Name:      "test-pod",
			Namespace: "test-namespace",
		},
	}

	assert.Equal(t, "Pod", scanObj.GetKind())
	assert.Equal(t, "v1", scanObj.GetApiVersion())
	assert.Equal(t, "test-pod", scanObj.GetName())
	assert.Equal(t, "test-namespace", scanObj.GetNamespace())
	assert.Equal(t, workloadinterface.TypeWorkloadObject, scanObj.GetObjectType())
	assert.Equal(t, "/v1/test-namespace/Pod/test-pod", scanObj.GetID())
}
