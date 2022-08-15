package reporthandling

import (
	"encoding/json"
	"testing"

	"github.com/kubescape/k8s-interface/workloadinterface"
	"github.com/kubescape/opa-utils/objectsenvelopes"
	"github.com/stretchr/testify/assert"
)

var (
	// role        = `{"apiVersion": "rbac.authorization.k8s.io/v1", "kind": "Role", "metadata": {"creationTimestamp": "2021-06-13T13:17:24Z","managedFields": [{"apiVersion": "rbac.authorization.k8s.io/v1","fieldsType": "FieldsV1","fieldsV1": {"f:rules": {}},"manager": "kubectl-edit","operation": "Update","time": "2021-06-13T13:22:29Z"}],"name": "pod-reader","namespace": "default","resourceVersion": "40233","uid": "cea4a847-2f05-4a94-bf3f-a8d1907e60e0"},"rules": [{"apiGroups": [""],"resources": ["pods","secrets"],"verbs": ["get"]}]}`
	// rolebinding = `{"apiVersion":"rbac.authorization.k8s.io/v1", "kind":"RoleBinding", "metadata":{"annotations":{"kubectl.kubernetes.io/last-applied-configuration":"{\"apiVersion\":\"rbac.authorization.k8s.io/v1\",\"kind\":\"RoleBinding\",\"metadata\":{\"annotations\":{},\"name\":\"read-pods\",\"namespace\":\"default\"},\"roleRef\":{\"apiGroup\":\"rbac.authorization.k8s.io\",\"kind\":\"Role\",\"name\":\"pod-reader\"},\"subjects\":[{\"apiGroup\":\"rbac.authorization.k8s.io\",\"kind\":\"User\",\"name\":\"jane\"}]}\n"},"creationTimestamp":"2021-11-11T11:50:38Z","managedFields":[{"apiVersion":"rbac.authorization.k8s.io/v1","fieldsType":"FieldsV1","fieldsV1":{"f:metadata":{"f:annotations":{".":{},"f:kubectl.kubernetes.io/last-applied-configuration":{}}},"f:roleRef":{"f:apiGroup":{},"f:kind":{},"f:name":{}},"f:subjects":{}},"manager":"kubectl-client-side-apply","operation":"Update","time":"2021-11-11T11:50:38Z"}],"name":"read-pods","namespace":"default","resourceVersion":"650451","uid":"6038eca8-b13e-4557-bc92-8800a11197d3"},"roleRef":{"apiGroup":"rbac.authorization.k8s.io","kind":"Role","name":"pod-reader"},"subjects":[{"apiGroup":"rbac.authorization.k8s.io","kind":"User","name":"jane"}]}`
	requiredObjectFields = []string{"kind", "name"}
	role                 = `{"apiVersion": "rbac.authorization.k8s.io/v1","kind": "Role","metadata": {"creationTimestamp": "2021-06-13T13:17:24Z","managedFields": [{"apiVersion": "rbac.authorization.k8s.io/v1","fieldsType": "FieldsV1","fieldsV1": {"f:rules": {}},"manager": "kubectl-edit","operation": "Update","time": "2021-06-13T13:22:29Z"}],"name": "pod-reader","namespace": "default","resourceVersion": "40233","uid": "cea4a847-2f05-4a94-bf3f-a8d1907e60e0"},"rules": [{"apiGroups": [""],"resources": ["pods","secrets"],"verbs": ["get"]}]}`
	rolebinding          = `{"apiVersion":"rbac.authorization.k8s.io/v1","kind":"RoleBinding","metadata":{"annotations":{"kubectl.kubernetes.io/last-applied-configuration":"{\"apiVersion\":\"rbac.authorization.k8s.io/v1\",\"kind\":\"RoleBinding\",\"metadata\":{\"annotations\":{},\"name\":\"read-pods\",\"namespace\":\"default\"},\"roleRef\":{\"apiGroup\":\"rbac.authorization.k8s.io\",\"kind\":\"Role\",\"name\":\"pod-reader\"},\"subjects\":[{\"apiGroup\":\"rbac.authorization.k8s.io\",\"kind\":\"User\",\"name\":\"jane\"}]}\n"},"creationTimestamp":"2021-11-11T11:50:38Z","managedFields":[{"apiVersion":"rbac.authorization.k8s.io/v1","fieldsType":"FieldsV1","fieldsV1":{"f:metadata":{"f:annotations":{".":{},"f:kubectl.kubernetes.io/last-applied-configuration":{}}},"f:roleRef":{"f:apiGroup":{},"f:kind":{},"f:name":{}},"f:subjects":{}},"manager":"kubectl-client-side-apply","operation":"Update","time":"2021-11-11T11:50:38Z"}],"name":"read-pods","namespace":"default","resourceVersion":"650451","uid":"6038eca8-b13e-4557-bc92-8800a11197d3"},"roleRef":{"apiGroup":"rbac.authorization.k8s.io","kind":"Role","name":"pod-reader"},"subjects":[{"apiGroup":"rbac.authorization.k8s.io","kind":"User","name":"jane"}]}`
	// rolebindingmanysubjects = `{"apiVersion":"rbac.authorization.k8s.io/v1","kind":"RoleBinding","metadata":{"annotations":{"kubectl.kubernetes.io/last-applied-configuration":"{\"apiVersion\":\"rbac.authorization.k8s.io/v1\",\"kind\":\"RoleBinding\",\"metadata\":{\"annotations\":{},\"creationTimestamp\":\"2021-11-11T11:50:38Z\",\"name\":\"read-pods\",\"namespace\":\"default\",\"resourceVersion\":\"650451\",\"uid\":\"6038eca8-b13e-4557-bc92-8800a11197d3\"},\"roleRef\":{\"apiGroup\":\"rbac.authorization.k8s.io\",\"kind\":\"Role\",\"name\":\"pod-reader\"},\"subjects\":[{\"apiGroup\":\"rbac.authorization.k8s.io\",\"kind\":\"User\",\"name\":\"jane\"},{\"kind\":\"ServiceAccount\",\"name\":\"default\",\"namespace\":\"kube-system\"}]}\n"},"creationTimestamp":"2021-11-11T11:50:38Z","managedFields":[{"apiVersion":"rbac.authorization.k8s.io/v1","fieldsType":"FieldsV1","fieldsV1":{"f:metadata":{"f:annotations":{".":{},"f:kubectl.kubernetes.io/last-applied-configuration":{}}},"f:roleRef":{"f:apiGroup":{},"f:kind":{},"f:name":{}},"f:subjects":{}},"manager":"kubectl-client-side-apply","operation":"Update","time":"2021-11-11T11:50:38Z"}],"name":"read-pods","namespace":"default","resourceVersion":"689305","uid":"6038eca8-b13e-4557-bc92-8800a11197d3"},"roleRef":{"apiGroup":"rbac.authorization.k8s.io","kind":"Role","name":"pod-reader"},"subjects":[{"apiGroup":"rbac.authorization.k8s.io","kind":"User","name":"jane"},{"kind":"ServiceAccount","name":"default","namespace":"kube-system"}]}`
	// apiServerPod            = `{ "apiVersion": "v1", "kind": "Pod", "metadata": { "annotations": { "kubeadm.kubernetes.io/kube-apiserver.advertise-address.endpoint": "192.168.49.2:8443", "kubernetes.io/config.hash": "01d7e312da0f9c4176daa8464d4d1a50", "kubernetes.io/config.mirror": "01d7e312da0f9c4176daa8464d4d1a50", "kubernetes.io/config.seen": "2021-10-20T13:57:08.810837592Z", "kubernetes.io/config.source": "file" }, "creationTimestamp": "2021-10-20T13:57:15Z", "labels": { "component": "kube-apiserver", "tier": "control-plane" }, "name": "kube-apiserver-minikube", "namespace": "kube-system", "ownerReferences": [ { "apiVersion": "v1", "controller": true, "kind": "Node", "name": "minikube", "uid": "a7fb33a8-c44d-48cd-a000-d49727cc78e4" } ], "resourceVersion": "218397", "uid": "7a26335b-a3d4-458c-ad4a-d90ebe1f36ed" }, "spec": { "containers": [ { "command": [ "kube-apiserver", "--advertise-address=192.168.49.2", "--allow-privileged=true", "--authorization-mode=Node,RBAC", "--client-ca-file=/var/lib/minikube/certs/ca.crt", "--enable-admission-plugins=NamespaceLifecycle,LimitRanger,ServiceAccount,DefaultStorageClass,DefaultTolerationSeconds,NodeRestriction,MutatingAdmissionWebhook,ValidatingAdmissionWebhook,ResourceQuota", "--enable-bootstrap-token-auth=true", "--etcd-cafile=/var/lib/minikube/certs/etcd/ca.crt", "--etcd-certfile=/var/lib/minikube/certs/apiserver-etcd-client.crt", "--etcd-keyfile=/var/lib/minikube/certs/apiserver-etcd-client.key", "--etcd-servers=https://127.0.0.1:2379", "--insecure-port=0", "--kubelet-client-certificate=/var/lib/minikube/certs/apiserver-kubelet-client.crt", "--kubelet-client-key=/var/lib/minikube/certs/apiserver-kubelet-client.key", "--kubelet-preferred-address-types=InternalIP,ExternalIP,Hostname", "--proxy-client-cert-file=/var/lib/minikube/certs/front-proxy-client.crt", "--proxy-client-key-file=/var/lib/minikube/certs/front-proxy-client.key", "--requestheader-allowed-names=front-proxy-client", "--requestheader-client-ca-file=/var/lib/minikube/certs/front-proxy-ca.crt", "--requestheader-extra-headers-prefix=X-Remote-Extra-", "--requestheader-group-headers=X-Remote-Group", "--requestheader-username-headers=X-Remote-User", "--secure-port=8443", "--service-account-issuer=https://kubernetes.default.svc.cluster.local", "--service-account-key-file=/var/lib/minikube/certs/sa.pub", "--service-account-signing-key-file=/var/lib/minikube/certs/sa.key", "--service-cluster-ip-range=10.96.0.0/12", "--tls-cert-file=/var/lib/minikube/certs/apiserver.crt", "--tls-private-key-file=/var/lib/minikube/certs/apiserver.key" ], "image": "k8s.gcr.io/kube-apiserver:v1.20.7", "imagePullPolicy": "IfNotPresent", "livenessProbe": { "failureThreshold": 8, "httpGet": { "host": "192.168.49.2", "path": "/livez", "port": 8443, "scheme": "HTTPS" }, "initialDelaySeconds": 10, "periodSeconds": 10, "successThreshold": 1, "timeoutSeconds": 15 }, "name": "kube-apiserver", "readinessProbe": { "failureThreshold": 3, "httpGet": { "host": "192.168.49.2", "path": "/readyz", "port": 8443, "scheme": "HTTPS" }, "periodSeconds": 1, "successThreshold": 1, "timeoutSeconds": 15 }, "resources": { "requests": { "cpu": "250m" } }, "startupProbe": { "failureThreshold": 24, "httpGet": { "host": "192.168.49.2", "path": "/livez", "port": 8443, "scheme": "HTTPS" }, "initialDelaySeconds": 10, "periodSeconds": 10, "successThreshold": 1, "timeoutSeconds": 15 }, "terminationMessagePath": "/dev/termination-log", "terminationMessagePolicy": "File", "volumeMounts": [ { "mountPath": "/etc/ssl/certs", "name": "ca-certs", "readOnly": true }, { "mountPath": "/etc/ca-certificates", "name": "etc-ca-certificates", "readOnly": true }, { "mountPath": "/var/lib/minikube/certs", "name": "k8s-certs", "readOnly": true }, { "mountPath": "/usr/local/share/ca-certificates", "name": "usr-local-share-ca-certificates", "readOnly": true }, { "mountPath": "/usr/share/ca-certificates", "name": "usr-share-ca-certificates", "readOnly": true } ] } ], "dnsPolicy": "ClusterFirst", "enableServiceLinks": true, "hostNetwork": true, "nodeName": "minikube", "preemptionPolicy": "PreemptLowerPriority", "priority": 2000001000, "priorityClassName": "system-node-critical", "restartPolicy": "Always", "schedulerName": "default-scheduler", "securityContext": {}, "terminationGracePeriodSeconds": 30, "tolerations": [ { "effect": "NoExecute", "operator": "Exists" } ], "volumes": [ { "hostPath": { "path": "/etc/ssl/certs", "type": "DirectoryOrCreate" }, "name": "ca-certs" }, { "hostPath": { "path": "/etc/ca-certificates", "type": "DirectoryOrCreate" }, "name": "etc-ca-certificates" }, { "hostPath": { "path": "/var/lib/minikube/certs", "type": "DirectoryOrCreate" }, "name": "k8s-certs" }, { "hostPath": { "path": "/usr/local/share/ca-certificates", "type": "DirectoryOrCreate" }, "name": "usr-local-share-ca-certificates" }, { "hostPath": { "path": "/usr/share/ca-certificates", "type": "DirectoryOrCreate" }, "name": "usr-share-ca-certificates" } ] }, "status": { "conditions": [ { "lastProbeTime": null, "lastTransitionTime": "2021-11-17T06:58:52Z", "status": "True", "type": "Initialized" }, { "lastProbeTime": null, "lastTransitionTime": "2021-11-17T11:32:59Z", "status": "True", "type": "Ready" }, { "lastProbeTime": null, "lastTransitionTime": "2021-11-17T11:32:59Z", "status": "True", "type": "ContainersReady" }, { "lastProbeTime": null, "lastTransitionTime": "2021-11-17T06:58:52Z", "status": "True", "type": "PodScheduled" } ], "containerStatuses": [ { "containerID": "docker://9e698d29975f1685151fb5ef39881c9548cca4254ff8b849b90a3aba0b98a422", "image": "k8s.gcr.io/kube-apiserver:v1.20.7", "imageID": "docker-pullable://k8s.gcr.io/kube-apiserver@sha256:5ab3d676c426bfb272fb7605e6978b90d5676913636a6105688862849961386f", "lastState": { "terminated": { "containerID": "docker://8effb839686f18678aaef9611a78a2e99a197c2fb2f5aee10b52414b9a415cfb", "exitCode": 255, "finishedAt": "2021-11-17T06:58:27Z", "reason": "Error", "startedAt": "2021-11-16T06:55:06Z" } }, "name": "kube-apiserver", "ready": true, "restartCount": 23, "started": true, "state": { "running": { "startedAt": "2021-11-17T06:58:54Z" } } } ], "hostIP": "192.168.49.2", "phase": "Running", "podIP": "192.168.49.2", "podIPs": [ { "ip": "192.168.49.2" } ], "qosClass": "Burstable", "startTime": "2021-11-17T06:58:52Z" } }`
)

func TestAggregateResourcesBySubjects(t *testing.T) {
	r, _ := objectsenvelopes.NewRegoResponseVectorObjectFromBytes([]byte(role))
	rb, _ := objectsenvelopes.NewRegoResponseVectorObjectFromBytes([]byte(rolebinding))
	inputList := []workloadinterface.IMetadata{r, rb}

	outputList, err := AggregateResourcesBySubjects(inputList)
	if err != nil {
		t.Errorf(err.Error())
	}

	assert.NotEqual(t, 1, len(outputList))
	assert.True(t, isObjectFields(outputList))

}

func TestAggregateResourcesBySubjects2(t *testing.T) {

	r := make(map[string]interface{})
	err := json.Unmarshal([]byte(role), &r)
	if err != nil {
		t.Errorf("error in unmarshal %s", err)
	}
	rb := make(map[string]interface{})
	err = json.Unmarshal([]byte(rolebinding), &rb)
	if err != nil {
		t.Errorf("error in unmarshal %s", err)
	}
	ro := objectsenvelopes.NewObject(r)
	rob := objectsenvelopes.NewObject(rb)
	inputList := []workloadinterface.IMetadata{ro, rob}

	outputList, err := AggregateResourcesBySubjects(inputList)
	if err != nil {
		t.Errorf(err.Error())
	}

	assert.Equal(t, 1, len(outputList))
	assert.True(t, isObjectFields(outputList))
}

func isObjectFields(objs []workloadinterface.IMetadata) bool {
	for _, obj := range objs {
		for _, field := range requiredObjectFields {
			if _, ok := obj.GetObject()[field]; !ok {
				return false
			}
		}
	}
	return true
}
