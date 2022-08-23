package listworkloads

import (
	"testing"

	"github.com/kubescape/k8s-interface/workloadinterface"
	"github.com/stretchr/testify/assert"
)

var (
	list       = `{"apiVersion": "v1","items": [{"apiVersion": "apps/v1","kind": "Deployment","metadata": {"name": "nginx-deployment","labels": {"app": "nginx"}},"spec": {"replicas": 3,"selector": {"matchLabels": {"app": "nginx"}},"template": {"metadata": {"labels": {"app": "nginx"}},"spec": {"containers": [{"name": "nginx","image": "nginx:1.14.2","ports": [{"containerPort": 80}]}]}}}}],"kind": "List","metadata": {"resourceVersion": ""}}`
	deployment = `{"apiVersion": "apps/v1","kind": "Deployment","metadata": {"name": "nginx-deployment","labels": {"app": "nginx"}},"spec": {"replicas": 3,"selector": {"matchLabels": {"app": "nginx"}},"template": {"metadata": {"labels": {"app": "nginx"}},"spec": {"containers": [{"name": "nginx","image": "nginx:1.14.2","ports": [{"containerPort": 80}]}]}}}}`
)

func getListMock(r string) *ListWorkloads {
	relatedObject, err := NewListWorkloads([]byte(r))
	if err != nil {
		panic(err)
	}
	return relatedObject
}

func getWorkloadMock(r string) map[string]interface{} {
	relatedObject, err := workloadinterface.NewWorkload([]byte(r))
	if err != nil {
		panic(err)
	}
	return relatedObject.GetObject()
}

func assertObjectFields(t *testing.T, l *ListWorkloads) {
	assert.Equal(t, "v1", l.GetApiVersion())
	assert.Equal(t, "List", l.GetKind())
	assert.Equal(t, TypeListWorkloads, l.GetObjectType())
}

func TestGetItems(t *testing.T) {
	listObj := getListMock(list)
	deployment := getWorkloadMock(deployment)
	assertObjectFields(t, listObj)
	items := listObj.GetItems()
	assert.ElementsMatch(t, []interface{}{deployment}, items)
}
