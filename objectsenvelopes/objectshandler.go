package objectsenvelopes

import (
	"github.com/armosec/k8s-interface/cloudsupport/apis"
	cloudsupportv1 "github.com/armosec/k8s-interface/cloudsupport/v1"
	"github.com/armosec/k8s-interface/k8sinterface"
	"github.com/armosec/k8s-interface/workloadinterface"
	"github.com/armosec/opa-utils/objectsenvelopes/hostsensor"
)

// Returns the currect object that supports the IMetadata interface
func NewObject(object map[string]interface{}) workloadinterface.IMetadata {
	if object == nil {
		return nil
	}
	switch GetObjectType(object) {
	case TypeRegoResponseVectorObject:
		return NewRegoResponseVectorObject(object)
	case cloudsupportv1.TypeCloudProviderDescribe:
		return cloudsupportv1.NewDescriptiveInfoFromCloudProvider(object)
	case hostsensor.TypeHostSensor:
		return hostsensor.NewHostSensorDataEnvelope(object)
	case workloadinterface.TypeWorkloadObject:
		return workloadinterface.NewWorkloadObj(object)
	default: // objects should follow the basic k8s structure
		return workloadinterface.NewBaseObject(object)
	}
}

func GetObjectType(object map[string]interface{}) workloadinterface.ObjectType {
	if IsTypeRegoResponseVector(object) {
		return TypeRegoResponseVectorObject
	}
	if hostsensor.IsTypeTypeHostSensor(object) {
		return hostsensor.TypeHostSensor
	}
	if apis.IsTypeDescriptiveInfoFromCloudProvider(object) {
		return cloudsupportv1.TypeCloudProviderDescribe
	}
	if k8sinterface.IsTypeWorkload(object) {
		return workloadinterface.TypeWorkloadObject
	}

	// Test if basic object only after testing the rest
	if workloadinterface.IsBaseObject(object) {
		return workloadinterface.TypeBaseObject
	}
	return workloadinterface.TypeUnknown
}

func ListMapToMeta(resourceMap []map[string]interface{}) []workloadinterface.IMetadata {
	workloads := []workloadinterface.IMetadata{}
	for i := range resourceMap {
		if w := NewObject(resourceMap[i]); w != nil {
			workloads = append(workloads, w)
		}
	}
	return workloads
}
