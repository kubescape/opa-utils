package localworkload

import (
	"fmt"

	"github.com/armosec/k8s-interface/workloadinterface"
	"github.com/armosec/utils-go/str"
)

const TypeLocalWorkload workloadinterface.ObjectType = "LocalWorkload"
const PathKey = "sourcePath"

type LocalWorkload struct {
	*workloadinterface.BaseObject
}

// NewLocalWorkload construct a NewLocalWorkload from map[string]interface{}. If the map does not match the object, will return nil
func NewLocalWorkload(object map[string]interface{}) *LocalWorkload {
	b := workloadinterface.NewBaseObject(object)
	if b == nil {
		return nil
	}
	localWorkload := &LocalWorkload{BaseObject: b}
	return localWorkload
}
func (localWorkload *LocalWorkload) GetID() string {
	return fmt.Sprintf("path=%s/api=%s", str.AsFNVHash(localWorkload.GetPath()), localWorkload.BaseObject.GetID())
}
func (localWorkload *LocalWorkload) SetPath(p string) {
	workloadinterface.SetInMap(localWorkload.GetObject(), []string{}, PathKey, p)
}

func (localWorkload *LocalWorkload) GetPath() string {
	if p, ok := workloadinterface.InspectMap(localWorkload.GetObject(), PathKey); ok {
		return p.(string)
	}
	return ""
}

func (localWorkload *LocalWorkload) DeletePathEntry() {
	workloadinterface.RemoveFromMap(localWorkload.GetObject(), PathKey)
}

func IsTypeLocalWorkload(object map[string]interface{}) bool {
	if object == nil {
		return false
	}

	if _, ok := object[PathKey]; ok {
		return true
	}
	return false
}
