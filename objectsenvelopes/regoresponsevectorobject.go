package objectsenvelopes

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/armosec/k8s-interface/workloadinterface"
)

// This object is the structure of externalObject responses from regos
// that are ru as part of kubescape/ posture scan
// i.e any object that isn't an exact k8s object
// such as - subjects.

// expected fields:
// name string, namespace string, kind string, apiVersion string (can be empty str if not relevant)
// relatedObjects []IMetadata - includes related objects that need to be shown together with failed object
// e.g subjects will have in relatedObjects - role + rolebinding

const RelatedObjectsKey string = "relatedObjects"
const TypeRegoResponseVectorObject workloadinterface.ObjectType = "regoResponse"

type RegoResponseVectorObject struct {
	object map[string]interface{}
}

func NewRegoResponseVectorObject(object map[string]interface{}) *RegoResponseVectorObject {
	return &RegoResponseVectorObject{
		object: object,
	}
}

func NewRegoResponseVectorObjectFromBytes(object []byte) (*RegoResponseVectorObject, error) {
	obj := make(map[string]interface{})
	if object != nil {
		if err := json.Unmarshal(object, &obj); err != nil {
			return nil, err
		}
	}
	return NewRegoResponseVectorObject(obj), nil

}
func (obj *RegoResponseVectorObject) ToString() string {
	o := obj.GetObject()
	if o == nil {
		return ""
	}
	bWorkload, err := json.Marshal(o)
	if err != nil {
		return err.Error()
	}
	return string(bWorkload)
}

// =================== Set ================================
func (obj *RegoResponseVectorObject) SetNamespace(namespace string) {
	obj.object["namespace"] = namespace
}

func (obj *RegoResponseVectorObject) SetName(name string) {
	obj.object["name"] = name
}

func (obj *RegoResponseVectorObject) SetKind(kind string) {
	obj.object["kind"] = kind
}

func (obj *RegoResponseVectorObject) SetWorkload(object map[string]interface{}) { // DEPRECATED
	obj.SetObject(object)
}

func (obj *RegoResponseVectorObject) SetObject(object map[string]interface{}) {
	obj.object = object
}

func (obj *RegoResponseVectorObject) SetRelatedObjects(relatedObjects []map[string]interface{}) {
	obj.object[RelatedObjectsKey] = relatedObjects
}

// =================== Get ================================
func (obj *RegoResponseVectorObject) GetApiVersion() string {
	if v, ok := workloadinterface.InspectMap(obj.object, "apiVersion"); ok {
		return v.(string)
	} else if v, ok := workloadinterface.InspectMap(obj.object, "apiGroup"); ok {
		return v.(string)
	}
	return ""
}

func (obj *RegoResponseVectorObject) GetNamespace() string {
	if v, ok := workloadinterface.InspectMap(obj.object, "namespace"); ok {
		return v.(string)
	}
	return ""
}

func (obj *RegoResponseVectorObject) GetName() string {
	if v, ok := workloadinterface.InspectMap(obj.object, "name"); ok {
		return v.(string)
	}
	return ""
}

func (obj *RegoResponseVectorObject) GetKind() string {
	if v, ok := workloadinterface.InspectMap(obj.object, "kind"); ok {
		return v.(string)
	}
	return ""
}

func (obj *RegoResponseVectorObject) GetWorkload() map[string]interface{} { // DEPRECATED
	return obj.GetObject()
}

func (obj *RegoResponseVectorObject) GetObject() map[string]interface{} {
	return obj.object
}

func (obj *RegoResponseVectorObject) GetObjectType() workloadinterface.ObjectType {
	return TypeRegoResponseVectorObject
}

func (obj *RegoResponseVectorObject) GetRelatedObjects() []workloadinterface.IMetadata {
	relatedObjects := []workloadinterface.IMetadata{}
	if r, ok := obj.object[RelatedObjectsKey]; ok {
		switch l := r.(type) {
		case []map[string]interface{}:
			for _, obj := range l {
				if o := NewObject(obj); o != nil {
					relatedObjects = append(relatedObjects, o)
				}
			}
		case []interface{}:
			for _, obj := range l {
				if m, ok := obj.(map[string]interface{}); ok {
					if o := NewObject(m); o != nil {
						relatedObjects = append(relatedObjects, o)
					}
				}
			}
		}
	}
	return relatedObjects
}

func (obj *RegoResponseVectorObject) GetID() string {
	relatedObjectsIDs := []string{}
	rr := obj.GetRelatedObjects()
	for _, o := range rr {
		if o != nil {
			relatedObjectsIDs = append(relatedObjectsIDs, o.GetID())
		}
	}
	relatedObjectsIDs = append(relatedObjectsIDs, fmt.Sprintf("%s/%s/%s/%s", obj.GetApiVersion(), obj.GetNamespace(), obj.GetKind(), obj.GetName()))
	sort.Strings(relatedObjectsIDs)
	return strings.Join(relatedObjectsIDs, "/")
}

// ===============================================================

func IsTypeRegoResponseVector(object map[string]interface{}) bool {
	if object == nil {
		return false
	}

	if _, ok := object["kind"]; !ok {
		return false
	}
	if _, ok := object["name"]; !ok {
		return false
	}
	if _, ok := object[RelatedObjectsKey]; !ok {
		return false
	}
	// DO NOT TEST "GROUP" - Not all objects have a group
	// DO NOT TEST "Namespace" - Not all objects have a namespace

	return true
}
