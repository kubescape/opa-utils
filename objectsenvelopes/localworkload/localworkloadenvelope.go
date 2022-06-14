package localworkload

import (
	"github.com/armosec/k8s-interface/workloadinterface"
)

const TypeLocalWorkload workloadinterface.ObjectType = "LocalWorkload"

type LocalWorkload struct {
	*workloadinterface.BaseObject
}

// NewHostSensorDataEnvelope construct a HostSensorDataEnvelope from map[string]interface{}. If the map does not match the object, will return nil
func NewLocalWorkload(object map[string]interface{}) *LocalWorkload {
	b := workloadinterface.NewBaseObject(object)
	if b == nil {
		return nil
	}
	localWorkload := &LocalWorkload{BaseObject: b}
	return localWorkload
}

func (localWorkload *LocalWorkload) SetPath(p string) {
	workloadinterface.SetInMap(localWorkload.GetObject(), []string{}, "path", p)
}

func (localWorkload *LocalWorkload) GetPath() string {
	if p, ok := workloadinterface.InspectMap(localWorkload.GetObject(), "path"); ok {
		return p.(string)
	}
	return ""
}

func IsTypeLocalWorkload(object map[string]interface{}) bool {
	if object == nil {
		return false
	}

	if _, ok := object["path"]; ok {
		return true
	}
	return false
}
