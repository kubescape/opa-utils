package reporthandling

import (
	"encoding/json"

	"github.com/kubescape/k8s-interface/workloadinterface"

	"github.com/kubescape/opa-utils/objectsenvelopes"
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

// IResource is an extension for IMetadata so we can include the source and other file metadata
type IResource interface {
	workloadinterface.IMetadata //
	GetSource() *Source
	SetSource(s *Source)
}

func NewResource(obj map[string]interface{}) *Resource {
	return &Resource{
		Object: obj,
	}
}
func NewResourceIMetadata(obj workloadinterface.IMetadata) *Resource {
	return &Resource{
		Object:     obj.GetObject(),
		IMetadata:  obj,
		ResourceID: obj.GetID(),
		Source:     nil,
	}
}

func (r *Resource) middleware() workloadinterface.IMetadata {
	if r.IMetadata != nil {
		return r.IMetadata
	}

	if r.Object != nil {
		mObject, ok := r.Object.(map[string]interface{})
		if !ok {
			// can we get rid of this marshal / unmarshal totally?
			// NOTE(fred): if the intent is to perform a deep-copy, encoding/gob would be more efficient.
			bObject, err := json.Marshal(r.Object)
			if err != nil {
				return r.IMetadata
			}

			if err = json.Unmarshal(bObject, &mObject); err != nil {
				return r.IMetadata
			}
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

func (r *Resource) SetApiVersion(s string) {
	mw := r.middleware()
	if mw != nil {
		mw.SetApiVersion(s)
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

func (r *Resource) SetSource(s *Source) { r.Source = s }
func (r *Resource) GetSource() *Source  { return r.Source }

func (r *Resource) GetObjectType() workloadinterface.ObjectType {
	mw := r.middleware()
	if mw == nil {
		return ""
	}

	return mw.GetObjectType()
}
