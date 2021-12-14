package objectsenvelopes

import (
	"github.com/armosec/k8s-interface/cloudsupport"
	"github.com/armosec/k8s-interface/workloadinterface"
)

// Returns the currect object that supports the IMetadata interface
func NewObject(object map[string]interface{}) workloadinterface.IMetadata {
	if object == nil {
		return nil
	}
	switch GetObjectType(object) {
	case workloadinterface.TypeWorkloadObject:
		return workloadinterface.NewWorkloadObj(object)
	case TypeRegoResponseVectorObject:
		return NewRegoResponseVectorObject(object)
	case cloudsupport.TypeCloudProviderDescription:
		return cloudsupport.NewDescriptiveInfoFromCloudProvider(object)
		// TODO -
		// case TypeHostSensor:
		// 	return NewHostSensorDataEnvelope(object)

	}
	return nil
}

func GetObjectType(object map[string]interface{}) workloadinterface.ObjectType {
	if workloadinterface.IsTypeWorkload(object) {
		return workloadinterface.TypeWorkloadObject
	}
	if IsTypeRegoResponseVector(object) {
		return TypeRegoResponseVectorObject
	}
	if IsTypeTypeHostSensor(object) {
		return TypeHostSensor
	}
	if cloudsupport.IsTypeDescriptiveInfoFromCloudProvider(object) {
		return cloudsupport.TypeCloudProviderDescription
	}

	// TODO - support sensors
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
