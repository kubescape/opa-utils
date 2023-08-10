package objectsenvelopes

import (
	"encoding/json"
	"fmt"

	"github.com/kubescape/k8s-interface/k8sinterface"
	"github.com/kubescape/k8s-interface/workloadinterface"
)

var _ workloadinterface.IMetadata = (*ScanObject)(nil)

type ScanObjectMetadata struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace,omitempty"`
}

// A ScanObject represents a K8S object to be scanned
type ScanObject struct {
	ApiVersion string             `json:"apiVersion"`
	Kind       string             `json:"kind"`
	Metadata   ScanObjectMetadata `json:"metadata"`
}

func IsTypeScanObject(object map[string]interface{}) bool {
	if object == nil {
		return false
	}

	if _, ok := object["apiVersion"]; !ok {
		return false
	}
	if _, ok := object["kind"]; !ok {
		return false
	}

	if _, ok := object["metadata"]; !ok {
		return false
	}

	metadata, ok := object["metadata"].(map[string]interface{})
	if !ok {
		return false
	}

	if _, ok = metadata["name"]; !ok {
		return false
	}

	return true
}

// NewScanObject construct a ScanObject from map[string]interface{}. If the map does not match the object, will return nil
func NewScanObject(object map[string]interface{}) *ScanObject {
	if !IsTypeScanObject(object) {
		return nil
	}

	scanObject := &ScanObject{}
	if b := workloadinterface.MapToBytes(object); b != nil {
		if err := json.Unmarshal(b, scanObject); err != nil {
			return nil
		}
	} else {
		return nil
	}
	return scanObject
}

func (scanObject *ScanObject) GetNamespace() string {
	return scanObject.Metadata.GetNamespace()
}

func (scanObjectMetadata *ScanObjectMetadata) GetName() string {
	return scanObjectMetadata.Name
}

func (scanObjectMetadata *ScanObjectMetadata) GetNamespace() string {
	return scanObjectMetadata.Namespace
}

func (scanObjectMetadata *ScanObjectMetadata) SetName(name string) {
	scanObjectMetadata.Name = name
}

func (scanObjectMetadata *ScanObjectMetadata) SetNamespace(namespace string) {
	scanObjectMetadata.Namespace = namespace
}

func (scanObject *ScanObject) GetName() string {
	return scanObject.Metadata.GetName()
}

func (scanObject *ScanObject) GetKind() string {
	return scanObject.Kind
}

func (scanObject *ScanObject) GetApiVersion() string {
	return scanObject.ApiVersion
}

func (scanObject *ScanObject) GetWorkload() map[string]interface{} {
	return scanObject.GetObject()
}

func (scanObject *ScanObject) GetObject() map[string]interface{} {
	m := map[string]interface{}{}
	b, err := json.Marshal(*scanObject)
	if err != nil {
		return m
	}
	return workloadinterface.BytesToMap(b)
}

func (scanObject *ScanObject) GetID() string {
	return fmt.Sprintf("%s/%s/%s/%s", k8sinterface.JoinGroupVersion(k8sinterface.SplitApiVersion(scanObject.GetApiVersion())), scanObject.GetNamespace(), scanObject.GetKind(), scanObject.GetName())
}

func (scanObject *ScanObject) GetObjectType() workloadinterface.ObjectType {
	return GetObjectType(scanObject.GetObject())
}

func (scanObject *ScanObject) SetNamespace(namespace string) {
	scanObject.Metadata.SetNamespace(namespace)
}

func (scanObject *ScanObject) SetName(name string) {
	scanObject.Metadata.SetName(name)
}

func (scanObject *ScanObject) SetKind(kind string) {
	scanObject.Kind = kind
}

func (scanObject *ScanObject) SetWorkload(object map[string]interface{}) {
	scanObject.SetObject(object)
}

func (scanObject *ScanObject) SetObject(object map[string]interface{}) {
	if !IsTypeScanObject(object) {
		return
	}
	if b := workloadinterface.MapToBytes(object); len(b) > 0 {
		obj := &ScanObject{}
		if err := json.Unmarshal(b, obj); err == nil {
			scanObject.SetApiVersion(obj.GetApiVersion())
			scanObject.SetKind(obj.GetKind())
			scanObject.SetName(obj.GetName())
			scanObject.SetNamespace(obj.GetNamespace())
		}
	}
}

func (scanObject *ScanObject) SetApiVersion(apiVersion string) {
	scanObject.ApiVersion = apiVersion
}
