package reporthandling

import (
	"encoding/json"

	"github.com/armosec/k8s-interface/workloadinterface"

	"github.com/armosec/opa-utils/objectsenvelopes"
)

// // Get
// GetNamespace() string
// GetName() string
// GetKind() string
// GetApiVersion() string
// GetWorkload() map[string]interface{} // DEPRECATED
// GetObject() map[string]interface{}
// GetID() string // Get K8S workload ID -> <api-group>/<api-version>/<kind>/<name>
// //isK8sObject()

func NewResource(obj map[string]interface{}) *Resource {
	return &Resource{
		Object: obj,
	}
}

func (r *Resource) middleware() workloadinterface.IMetadata {
	if r.IMetadata != nil {
		return r.IMetadata
	}

	if r.Object != nil {
		bObject, err := json.Marshal(r.Object)
		if err != nil {
			return r.IMetadata
		}
		mObject := map[string]interface{}{}
		if err := json.Unmarshal(bObject, &mObject); err != nil {
			return r.IMetadata
		}
		r.IMetadata = objectsenvelopes.NewObject(mObject)
	}
	return r.IMetadata
}

func (r *Resource) SetNamespace(s string) {
	mw := r.middleware()
	if mw != nil {
		mw.SetNamespace(s)
	}
}

func (r *Resource) SetName(s string) {
	mw := r.middleware()
	if mw != nil {
		mw.SetName(s)
	}
}

func (r *Resource) SetKind(s string) {
	mw := r.middleware()
	if mw != nil {
		mw.SetKind(s)
	}
}

func (r *Resource) SetWorkload(m map[string]interface{}) { // deprecated
	mw := r.middleware()
	if mw != nil {
		mw.SetWorkload(m)
	}
}

func (r *Resource) SetObject(m map[string]interface{}) {
	mw := r.middleware()
	if mw != nil {
		mw.SetObject(m)
	}
}

func (r *Resource) GetNamespace() string {
	mw := r.middleware()
	if mw == nil {
		return ""
	}
	return mw.GetNamespace()
}
func (r *Resource) GetName() string {
	mw := r.middleware()
	if mw == nil {
		return ""
	}
	return mw.GetName()
}
func (r *Resource) GetKind() string {
	mw := r.middleware()
	if mw == nil {
		return ""
	}
	return mw.GetKind()

}
func (r *Resource) GetApiVersion() string {
	mw := r.middleware()
	if mw == nil {
		return ""
	}
	return mw.GetApiVersion()

}
func (r *Resource) GetWorkload() map[string]interface{} {
	mw := r.middleware()
	if mw == nil {
		return nil
	}
	return mw.GetWorkload()

}
func (r *Resource) GetObject() map[string]interface{} {
	mw := r.middleware()
	if mw == nil {
		return nil
	}
	return mw.GetObject()
}
func (r *Resource) GetID() string {
	mw := r.middleware()
	if mw == nil {
		return ""
	}
	return mw.GetID()
}

func (r *Resource) GetObjectType() workloadinterface.ObjectType {
	mw := r.middleware()
	if mw == nil {
		return ""
	}

	return mw.GetObjectType()
}
