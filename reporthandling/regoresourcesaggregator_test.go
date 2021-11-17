package reporthandling

import (
	"encoding/json"
	"testing"
)

var (
	role                    = `{"apiVersion": "rbac.authorization.k8s.io/v1","kind": "Role","metadata": {"creationTimestamp": "2021-06-13T13:17:24Z","managedFields": [{"apiVersion": "rbac.authorization.k8s.io/v1","fieldsType": "FieldsV1","fieldsV1": {"f:rules": {}},"manager": "kubectl-edit","operation": "Update","time": "2021-06-13T13:22:29Z"}],"name": "pod-reader","namespace": "default","resourceVersion": "40233","uid": "cea4a847-2f05-4a94-bf3f-a8d1907e60e0"},"rules": [{"apiGroups": [""],"resources": ["pods","secrets"],"verbs": ["get"]}]}`
	rolebinding             = `{"apiVersion":"rbac.authorization.k8s.io/v1","kind":"RoleBinding","metadata":{"annotations":{"kubectl.kubernetes.io/last-applied-configuration":"{\"apiVersion\":\"rbac.authorization.k8s.io/v1\",\"kind\":\"RoleBinding\",\"metadata\":{\"annotations\":{},\"name\":\"read-pods\",\"namespace\":\"default\"},\"roleRef\":{\"apiGroup\":\"rbac.authorization.k8s.io\",\"kind\":\"Role\",\"name\":\"pod-reader\"},\"subjects\":[{\"apiGroup\":\"rbac.authorization.k8s.io\",\"kind\":\"User\",\"name\":\"jane\"}]}\n"},"creationTimestamp":"2021-11-11T11:50:38Z","managedFields":[{"apiVersion":"rbac.authorization.k8s.io/v1","fieldsType":"FieldsV1","fieldsV1":{"f:metadata":{"f:annotations":{".":{},"f:kubectl.kubernetes.io/last-applied-configuration":{}}},"f:roleRef":{"f:apiGroup":{},"f:kind":{},"f:name":{}},"f:subjects":{}},"manager":"kubectl-client-side-apply","operation":"Update","time":"2021-11-11T11:50:38Z"}],"name":"read-pods","namespace":"default","resourceVersion":"650451","uid":"6038eca8-b13e-4557-bc92-8800a11197d3"},"roleRef":{"apiGroup":"rbac.authorization.k8s.io","kind":"Role","name":"pod-reader"},"subjects":[{"apiGroup":"rbac.authorization.k8s.io","kind":"User","name":"jane"}]}`
	rolebindingmanysubjects = `{"apiVersion":"rbac.authorization.k8s.io/v1","kind":"RoleBinding","metadata":{"annotations":{"kubectl.kubernetes.io/last-applied-configuration":"{\"apiVersion\":\"rbac.authorization.k8s.io/v1\",\"kind\":\"RoleBinding\",\"metadata\":{\"annotations\":{},\"creationTimestamp\":\"2021-11-11T11:50:38Z\",\"name\":\"read-pods\",\"namespace\":\"default\",\"resourceVersion\":\"650451\",\"uid\":\"6038eca8-b13e-4557-bc92-8800a11197d3\"},\"roleRef\":{\"apiGroup\":\"rbac.authorization.k8s.io\",\"kind\":\"Role\",\"name\":\"pod-reader\"},\"subjects\":[{\"apiGroup\":\"rbac.authorization.k8s.io\",\"kind\":\"User\",\"name\":\"jane\"},{\"kind\":\"ServiceAccount\",\"name\":\"default\",\"namespace\":\"kube-system\"}]}\n"},"creationTimestamp":"2021-11-11T11:50:38Z","managedFields":[{"apiVersion":"rbac.authorization.k8s.io/v1","fieldsType":"FieldsV1","fieldsV1":{"f:metadata":{"f:annotations":{".":{},"f:kubectl.kubernetes.io/last-applied-configuration":{}}},"f:roleRef":{"f:apiGroup":{},"f:kind":{},"f:name":{}},"f:subjects":{}},"manager":"kubectl-client-side-apply","operation":"Update","time":"2021-11-11T11:50:38Z"}],"name":"read-pods","namespace":"default","resourceVersion":"689305","uid":"6038eca8-b13e-4557-bc92-8800a11197d3"},"roleRef":{"apiGroup":"rbac.authorization.k8s.io","kind":"Role","name":"pod-reader"},"subjects":[{"apiGroup":"rbac.authorization.k8s.io","kind":"User","name":"jane"},{"kind":"ServiceAccount","name":"default","namespace":"kube-system"}]}`
)

func TestAggregateResourcesBySubjects(t *testing.T) {
	r := make(map[string]interface{})
	err := json.Unmarshal([]byte(role), &r)
	if err != nil {
		t.Errorf("error unmarshaling %s", err)
	}
	rb := make(map[string]interface{})
	err = json.Unmarshal([]byte(rolebinding), &rb)
	if err != nil {
		t.Errorf("error unmarshaling %s", err)
	}
	// r := make(map[string]interface{}, []byte(role))
	inputList := []map[string]interface{}{r, rb}
	outputList := AggregateResourcesBySubjects(inputList)
	if len(outputList) != 1 {
		t.Errorf("error in AggregateResourcesBySubjects, len should be 1, got len = %d", len(outputList))
	}
}

func TestAggregateResourcesBySubjects2(t *testing.T) {
	r := make(map[string]interface{})
	err := json.Unmarshal([]byte(role), &r)
	if err != nil {
		t.Errorf("error unmarshaling %s", err)
	}
	rb := make(map[string]interface{})
	err = json.Unmarshal([]byte(rolebindingmanysubjects), &rb)
	if err != nil {
		t.Errorf("error unmarshaling %s", err)
	}
	// r := make(map[string]interface{}, []byte(role))
	inputList := []map[string]interface{}{r, rb}
	outputList := AggregateResourcesBySubjects(inputList)
	if len(outputList) != 2 {
		t.Errorf("error in AggregateResourcesBySubjects, len should be 2, got len = %d", len(outputList))
	}
}
