package reporthandling

import (
	"encoding/json"

	ik8s "github.com/armosec/k8s-interface/workloadinterface"
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
func (r *Resource) SetNamespace(s string) {

}

func (r *Resource) middleware() ik8s.IMetadata {
	if r.IMetadata != nil {
		return r.IMetadata
	}
	b, err := json.Marshal(r.Object)
	if err != nil {
		return nil
	}

	if vc, err := ik8s.NewRegoResponseVectorObjectFromBytes(b, nil); err == nil && vc.GetID() != "" && vc.GetName() != "" {
		r.IMetadata = vc
		return vc
	}

	if wl, err := ik8s.NewWorkload(b); err == nil && wl.GetID() != "" && wl.GetName() != "" {
		r.IMetadata = wl
		return r.IMetadata
	}

	return nil

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

func (r *Resource) SetWorkload(m map[string]interface{}) { // deprectaed
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
	return mw.GetKind()

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
