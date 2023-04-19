package objectsenvelopes

import (
	"testing"

	cloudsupportv1 "github.com/kubescape/k8s-interface/cloudsupport/v1"
	"github.com/kubescape/k8s-interface/workloadinterface"
	"github.com/kubescape/opa-utils/objectsenvelopes/hostsensor"
	"github.com/kubescape/opa-utils/objectsenvelopes/localworkload"
	"github.com/stretchr/testify/assert"
)

func TestNewObject(t *testing.T) {
	// Test nil input
	assert.Nil(t, NewObject(nil))

	// Test valid input
	object := map[string]interface{}{
		"kind":       "Pod",
		"apiVersion": "v1",
		"metadata": map[string]interface{}{
			"name": "test-pod",
		},
	}
	workloadObj := NewObject(object)
	assert.NotNil(t, workloadObj)
	assert.Equal(t, "Pod", workloadObj.GetKind())

	// Test unsupported input
	unsupportedObject := map[string]interface{}{
		"unknown": "unknown",
	}
	assert.Nil(t, NewObject(unsupportedObject))
}

func TestGetObjectType(t *testing.T) {
	// Test unsupported input
	unsupportedObject := map[string]interface{}{
		"unknown": "unknown",
	}
	assert.Equal(t, workloadinterface.TypeUnknown, GetObjectType(unsupportedObject))

	// Test RegoResponseVectorObject
	relatedObjects := []map[string]interface{}{}
	relatedObject := getMock(role)
	relatedObject2 := getMock(rolebinding)
	relatedObjects = append(relatedObjects, relatedObject)
	relatedObjects = append(relatedObjects, relatedObject2)
	subject := map[string]interface{}{"name": "user@example.com", "kind": "User", "namespace": "default", "group": "rbac.authorization.k8s.io", RelatedObjectsKey: relatedObjects}
	assert.Equal(t, TypeRegoResponseVectorObject, GetObjectType(subject))

	// Test CloudProviderDescribe
	cloudProviderDescribe := map[string]interface{}{
		"apiVersion": "container.googleapis.com/v1",
		"kind":       "ClusterDescribe",
	}
	assert.Equal(t, cloudsupportv1.TypeCloudProviderDescribe, GetObjectType(cloudProviderDescribe))

	// Test HostSensor
	hostSensor := map[string]interface{}{
		"apiVersion": "hostdata.kubescape.cloud/v1",
		"metadata": map[string]interface{}{
			"name": "test-pod",
		},
	}
	assert.Equal(t, hostsensor.TypeHostSensor, GetObjectType(hostSensor))

	// Test LocalWorkload
	localWorkload := map[string]interface{}{
		"kind":       "b",
		"sourcePath": "/path/file",
	}
	assert.Equal(t, localworkload.TypeLocalWorkload, GetObjectType(localWorkload))

	// Test WorkloadObject
	workloadObject := map[string]interface{}{
		"kind":       "Pod",
		"apiVersion": "v1",
		"metadata": map[string]interface{}{
			"name": "test-pod",
		},
	}
	assert.Equal(t, workloadinterface.TypeWorkloadObject, GetObjectType(workloadObject))

	// Test ListWorkloads
	listWorkloads := map[string]interface{}{
		"kind": "List",
		"items": []interface{}{
			map[string]interface{}{
				"kind": "Pod",
				"metadata": map[string]interface{}{
					"name": "test-pod",
				},
			},
		},
	}
	assert.Equal(t, workloadinterface.TypeListWorkloads, GetObjectType(listWorkloads))
}

// // Test BaseObject
// baseObject := map[string]interface{}{
// 	"kind": "Pod",
// 	"metadata": map[string]interface{}{
// 		"name": "test-pod",
// 	},
// }
// assert.Equal(t
