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
			"name": "test-pod",
		},
	}
	scanObj := NewScanObject(object)
	assert.NotNil(t, scanObj)
	assert.Equal(t, "Pod", scanObj.GetKind())
	assert.Equal(t, "v1", scanObj.GetApiVersion())
	assert.Equal(t, "test-pod", scanObj.GetName())
	assert.Equal(t, "", scanObj.GetNamespace())
	scanObj.SetNamespace("test-namespace")
	assert.Equal(t, "test-namespace", scanObj.GetNamespace())
	assert.Equal(t, workloadinterface.TypeWorkloadObject, scanObj.GetObjectType())

	// test json marshal
	jsonEnc, err := json.Marshal(scanObj)
	expectedJsonStr := `{"apiVersion":"v1","kind":"Pod","metadata":{"name":"test-pod","namespace":"test-namespace"}}`
	assert.NoError(t, err)
	assert.Equal(t, expectedJsonStr, string(jsonEnc))

	// test json unmarshal
	var scanObj2 ScanObject
	jsonStr := `{"apiVersion":"v1","kind":"Node","metadata":{"name":"node-1"},"field":"value"}`
	err = json.Unmarshal([]byte(jsonStr), &scanObj2)
	assert.NoError(t, err)
	assert.Equal(t, "Node", scanObj2.GetKind())
	assert.Equal(t, "node-1", scanObj2.GetName())
	assert.Equal(t, "", scanObj2.GetNamespace())

	// Test unsupported input
	unsupportedObject := map[string]interface{}{
		"unknown": "unknown",
	}
	assert.Nil(t, NewScanObject(unsupportedObject))
}
