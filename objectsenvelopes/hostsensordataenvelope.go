package objectsenvelopes

import (
	"encoding/json"
	"fmt"

	"github.com/armosec/k8s-interface/workloadinterface"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const TypeHostSensor workloadinterface.ObjectType = "hostSensor"

type HostSensorDataEnvelope struct {
	schema.GroupVersionResource
	NodeName string          `json:"nodeName"`
	Data     json.RawMessage `json:"data"`
}

func NewHostSensorDataEnvelope(object map[string]interface{}) *HostSensorDataEnvelope {
	return &HostSensorDataEnvelope{} // TODO - convert object to HostSensorDataEnvelope
}

func (hsde *HostSensorDataEnvelope) SetNamespace(string) {

}

func (hsde *HostSensorDataEnvelope) SetName(val string) {
	hsde.NodeName = val
}

func (hsde *HostSensorDataEnvelope) SetKind(val string) {
	hsde.Resource = val

}

func (hsde *HostSensorDataEnvelope) SetWorkload(val map[string]interface{}) { //deprecated
	hsde.Data, _ = json.Marshal(val)
}

func (hsde *HostSensorDataEnvelope) SetObject(val map[string]interface{}) {
	hsde.Data, _ = json.Marshal(val)
}

func (hsde *HostSensorDataEnvelope) GetNamespace() string {
	return ""
}

func (hsde *HostSensorDataEnvelope) GetName() string {
	return hsde.NodeName
}

func (hsde *HostSensorDataEnvelope) GetKind() string {
	return hsde.Resource
}

func (hsde *HostSensorDataEnvelope) GetApiVersion() string {
	return hsde.Version
}

func (hsde *HostSensorDataEnvelope) GetGroup() string {
	return hsde.Group
}

func (hsde *HostSensorDataEnvelope) GetWorkload() map[string]interface{} { // DEPRECATED
	res := map[string]interface{}{}
	json.Unmarshal(hsde.Data, &res)
	return res
}

func (hsde *HostSensorDataEnvelope) GetObject() map[string]interface{} {
	res := map[string]interface{}{}
	json.Unmarshal(hsde.Data, &res)
	return res
}

func (hsde *HostSensorDataEnvelope) GetObjectType() workloadinterface.ObjectType {
	return TypeHostSensor
}
func (hsde *HostSensorDataEnvelope) GetID() string { // -> <api-group>/<api-version>/<kind>/<name>
	return fmt.Sprintf("%s/%s/%s/%s", hsde.Group, hsde.GetApiVersion(), hsde.GetKind(), hsde.GetName())
}

func IsTypeTypeHostSensor(object map[string]interface{}) bool {
	if object == nil {
		return false
	}

	if _, ok := object["nodeName"]; !ok {
		return false
	}
	if _, ok := object["data"]; !ok {
		return false
	}
	return true
}
