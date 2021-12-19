package hostsensor

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/armosec/k8s-interface/workloadinterface"
)

const TypeHostSensor workloadinterface.ObjectType = "HostSensor"

const (
	GroupHostSensor = "hostdata.armo.cloud"
	Version         = "v1beta0"
	// KindOsReleaseFile = "OsReleaseFile"

)

type HostSensorMetadata struct {
	Name string `json:"name"` // nodeName
}
type HostSensorDataEnvelope struct {
	ApiVersion string             `json:"apiVersion"`
	Kind       string             `json:"kind"`
	Metadata   HostSensorMetadata `json:"Metadata"`
	Data       json.RawMessage    `json:"data"`
}

// NewHostSensorDataEnvelope construct a HostSensorDataEnvelope from map[string]interface{}. If the map does not match the object, will return nil
func NewHostSensorDataEnvelope(object map[string]interface{}) *HostSensorDataEnvelope {
	if !IsTypeTypeHostSensor(object) {
		return nil
	}

	hostSensorDataEnvelope := &HostSensorDataEnvelope{}
	if b := workloadinterface.MapToBytes(object); b != nil {
		if err := json.Unmarshal(b, hostSensorDataEnvelope); err != nil {
			return nil
		}
	} else {
		return nil
	}
	return hostSensorDataEnvelope
}

func (hostSensorMetadata *HostSensorMetadata) GetName() string {
	return hostSensorMetadata.Name
}

func (hostSensorMetadata *HostSensorMetadata) SetName(name string) {
	hostSensorMetadata.Name = name
}

// SetNamespace kept for compatibility with the IMetdata interface
func (hsde *HostSensorDataEnvelope) SetNamespace(string) {
	// not namespaced object
}

func (hsde *HostSensorDataEnvelope) SetName(name string) {
	hsde.Metadata.SetName(name)
}

func (hsde *HostSensorDataEnvelope) SetKind(kind string) {
	hsde.Kind = kind

}

func (hsde *HostSensorDataEnvelope) SetData(data json.RawMessage) {
	hsde.Data = data
}

func (hsde *HostSensorDataEnvelope) SetApiVersion(apiVersion string) {
	hsde.ApiVersion = apiVersion
}

func (hsde *HostSensorDataEnvelope) SetWorkload(object map[string]interface{}) { //deprecated
	hsde.SetObject(object)
}

func (hsde *HostSensorDataEnvelope) SetObject(object map[string]interface{}) {
	if !IsTypeTypeHostSensor(object) {
		return
	}
	if b := workloadinterface.MapToBytes(object); len(b) > 0 {
		hostSensorDataEnvelope := &HostSensorDataEnvelope{}
		if err := json.Unmarshal(b, hostSensorDataEnvelope); err == nil {
			hsde.SetApiVersion(hostSensorDataEnvelope.GetApiVersion())
			hsde.SetKind(hostSensorDataEnvelope.GetKind())
			hsde.SetData(hostSensorDataEnvelope.GetData())
			hsde.Metadata = hostSensorDataEnvelope.Metadata
		}
	}
}

// GetNamespace kept for compatibility with the IMetdata interface
func (hsde *HostSensorDataEnvelope) GetNamespace() string {
	return ""
}

func (hsde *HostSensorDataEnvelope) GetData() json.RawMessage {
	return hsde.Data
}

func (hsde *HostSensorDataEnvelope) GetName() string {
	return hsde.Metadata.GetName()
}

func (hsde *HostSensorDataEnvelope) GetKind() string {
	return hsde.Kind
}

func (hsde *HostSensorDataEnvelope) GetApiVersion() string {
	return hsde.ApiVersion
}

// GetWorkload - DEPRECATED - kept for compatibility with the IMetdata interface
func (hsde *HostSensorDataEnvelope) GetWorkload() map[string]interface{} {
	return hsde.GetObject()
}

func (hsde *HostSensorDataEnvelope) GetObject() map[string]interface{} {
	m := map[string]interface{}{}
	b, err := json.Marshal(*hsde)
	if err != nil {
		return m
	}
	return workloadinterface.BytesToMap(b)
}

func (hsde *HostSensorDataEnvelope) GetObjectType() workloadinterface.ObjectType {
	return TypeHostSensor
}
func (hsde *HostSensorDataEnvelope) GetID() string { // ->  <api-version>/<kind>/<name>
	return fmt.Sprintf("%s/%s/%s", hsde.GetApiVersion(), hsde.GetKind(), hsde.GetName())
}

func IsTypeTypeHostSensor(object map[string]interface{}) bool {
	if object == nil {
		return false
	}

	if apiVersion, ok := object["apiVersion"]; ok {
		if group := strings.Split(apiVersion.(string), "/"); group[0] == GroupHostSensor {
			return true
		}
	}
	return false
}
