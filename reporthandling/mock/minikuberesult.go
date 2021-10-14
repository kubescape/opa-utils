package mock

var NSAScanV10119 = `
{
    "name": "NSA",
    "controlReports": [
        {
            "guid": "",
            "attributes": {
                "armoBuiltin": true
            },
            "id": "C-0005",
            "controlID": "C-0005",
            "name": "Control plane hardening",
            "ruleReports": [
                {
                    "name": "insecure-port-flag",
                    "remediation": "",
                    "ruleStatus": {
                        "status": "success",
                        "message": ""
                    },
                    "ruleResponses": null
                }
            ],
            "remediation": "Set the insecure-port flag of the API server to zero.",
            "description": "Kubernetes control plane API is running with non-secure port enabled which allows attackers to gain unprotected access to the cluster.",
            "score": 100
        },
        {
            "guid": "",
            "attributes": {
                "armoBuiltin": true
            },
            "id": "C-0038",
            "controlID": "C-0038",
            "name": "Host PID/IPC privileges",
            "ruleReports": [
                {
                    "name": "host-pid-ipc-privileges",
                    "remediation": "",
                    "ruleStatus": {
                        "status": "success",
                        "message": ""
                    },
                    "ruleResponses": null
                }
            ],
            "remediation": "Remove hostPID and hostIPC privileges unless they are absolutely necessary.",
            "description": "Containers should be as isolated as possible from the host machine. The hostPID and hostIPC fields in Kubernetes may excessively expose the host to potentially malicious actions.",
            "score": 100
        },
        {
            "guid": "",
            "attributes": {
                "armoBuiltin": true
            },
            "id": "C-0017",
            "controlID": "C-0017",
            "name": "Immutable container filesystem",
            "ruleReports": [
                {
                    "name": "immutable-container-filesystem",
                    "remediation": "",
                    "ruleStatus": {
                        "status": "success",
                        "message": ""
                    },
                    "ruleResponses": [
                        {
                            "alertMessage": "container: etcd in pod: etcd-david-virtualbox  has  mutable filesystem",
                            "ruleStatus": "warning",
                            "packagename": "armo_builtins",
                            "alertScore": 7,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "Pod",
                                        "metadata": {
                                            "annotations": {
                                                "kubernetes.io/config.hash": "e0fcc6e4323055b5880f8aac4c950836",
                                                "kubernetes.io/config.mirror": "e0fcc6e4323055b5880f8aac4c950836",
                                                "kubernetes.io/config.seen": "2021-06-20T12:06:52.495386281+03:00",
                                                "kubernetes.io/config.source": "file"
                                            },
                                            "creationTimestamp": "2021-06-20T09:08:22Z",
                                            "labels": {
                                                "component": "etcd",
                                                "tier": "control-plane"
                                            },
                                            "name": "etcd-david-virtualbox",
                                            "namespace": "kube-system",
                                            "resourceVersion": "2110909",
                                            "selfLink": "/api/v1/namespaces/kube-system/pods/etcd-david-virtualbox",
                                            "uid": "154e7f87-907f-4edb-a73c-26e965d4fe02"
                                        },
                                        "spec": {
                                            "containers": [
                                                {
                                                    "command": [
                                                        "etcd",
                                                        "--advertise-client-urls=https://10.0.2.15:2379",
                                                        "--cert-file=/var/lib/minikube/certs/etcd/server.crt",
                                                        "--client-cert-auth=true",
                                                        "--data-dir=/var/lib/minikube/etcd",
                                                        "--initial-advertise-peer-urls=https://10.0.2.15:2380",
                                                        "--initial-cluster=david-virtualbox=https://10.0.2.15:2380",
                                                        "--key-file=/var/lib/minikube/certs/etcd/server.key",
                                                        "--listen-client-urls=https://127.0.0.1:2379,https://10.0.2.15:2379",
                                                        "--listen-metrics-urls=http://127.0.0.1:2381,http://10.0.2.15:2381",
                                                        "--listen-peer-urls=https://10.0.2.15:2380",
                                                        "--name=david-virtualbox",
                                                        "--peer-cert-file=/var/lib/minikube/certs/etcd/peer.crt",
                                                        "--peer-client-cert-auth=true",
                                                        "--peer-key-file=/var/lib/minikube/certs/etcd/peer.key",
                                                        "--peer-trusted-ca-file=/var/lib/minikube/certs/etcd/ca.crt",
                                                        "--snapshot-count=10000",
                                                        "--trusted-ca-file=/var/lib/minikube/certs/etcd/ca.crt"
                                                    ],
                                                    "image": "k8s.gcr.io/etcd:3.3.15-0",
                                                    "imagePullPolicy": "IfNotPresent",
                                                    "livenessProbe": {
                                                        "failureThreshold": 8,
                                                        "httpGet": {
                                                            "host": "127.0.0.1",
                                                            "path": "/health",
                                                            "port": 2381,
                                                            "scheme": "HTTP"
                                                        },
                                                        "initialDelaySeconds": 15,
                                                        "periodSeconds": 10,
                                                        "successThreshold": 1,
                                                        "timeoutSeconds": 15
                                                    },
                                                    "name": "etcd",
                                                    "resources": {},
                                                    "terminationMessagePath": "/dev/termination-log",
                                                    "terminationMessagePolicy": "File",
                                                    "volumeMounts": [
                                                        {
                                                            "mountPath": "/var/lib/minikube/etcd",
                                                            "name": "etcd-data"
                                                        },
                                                        {
                                                            "mountPath": "/var/lib/minikube/certs/etcd",
                                                            "name": "etcd-certs"
                                                        }
                                                    ]
                                                }
                                            ],
                                            "dnsPolicy": "ClusterFirst",
                                            "enableServiceLinks": true,
                                            "hostNetwork": true,
                                            "nodeName": "david-virtualbox",
                                            "priority": 2000000000,
                                            "priorityClassName": "system-cluster-critical",
                                            "restartPolicy": "Always",
                                            "schedulerName": "default-scheduler",
                                            "securityContext": {},
                                            "terminationGracePeriodSeconds": 30,
                                            "tolerations": [
                                                {
                                                    "effect": "NoExecute",
                                                    "operator": "Exists"
                                                }
                                            ],
                                            "volumes": [
                                                {
                                                    "hostPath": {
                                                        "path": "/var/lib/minikube/certs/etcd",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "etcd-certs"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/var/lib/minikube/etcd",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "etcd-data"
                                                }
                                            ]
                                        },
                                        "status": {
                                            "conditions": [
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-03T06:28:57Z",
                                                    "status": "True",
                                                    "type": "Initialized"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-06T05:23:22Z",
                                                    "status": "True",
                                                    "type": "Ready"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-06T05:23:22Z",
                                                    "status": "True",
                                                    "type": "ContainersReady"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-03T06:28:57Z",
                                                    "status": "True",
                                                    "type": "PodScheduled"
                                                }
                                            ],
                                            "containerStatuses": [
                                                {
                                                    "containerID": "docker://0cc62102444cc74e3dddb2b6cd7550036feaa0e61f4297c85fc72e3be2905ec9",
                                                    "image": "k8s.gcr.io/etcd:3.3.15-0",
                                                    "imageID": "docker-pullable://k8s.gcr.io/etcd@sha256:12c2c5e5731c3bcd56e6f1c05c0f9198b6f06793fa7fca2fb43aab9622dc4afa",
                                                    "lastState": {
                                                        "terminated": {
                                                            "containerID": "docker://61de0abaa35617dafed9274c5b03f99f269412bdea2f15d903b4293d620f4b9e",
                                                            "exitCode": 255,
                                                            "finishedAt": "2021-10-14T05:35:02Z",
                                                            "reason": "Error",
                                                            "startedAt": "2021-10-12T10:28:38Z"
                                                        }
                                                    },
                                                    "name": "etcd",
                                                    "ready": true,
                                                    "restartCount": 51,
                                                    "started": true,
                                                    "state": {
                                                        "running": {
                                                            "startedAt": "2021-10-14T05:35:12Z"
                                                        }
                                                    }
                                                }
                                            ],
                                            "hostIP": "10.0.2.15",
                                            "phase": "Running",
                                            "podIP": "10.0.2.15",
                                            "podIPs": [
                                                {
                                                    "ip": "10.0.2.15"
                                                }
                                            ],
                                            "qosClass": "BestEffort",
                                            "startTime": "2021-10-03T06:28:57Z"
                                        }
                                    }
                                ]
                            },
                            "exception": {
                                "guid": "0e71632b-d377-4d54-a6d5-67643bcb6047",
                                "name": "exception_C-0017_kube-system_cc21d0ba28b70809f341e4809ecd51d2",
                                "attributes": {
                                    "namespaceOnly": "true"
                                },
                                "policyType": "postureExceptionPolicy",
                                "creationTime": "2021-10-14T12:13:00.213686",
                                "actions": [
                                    "alertOnly"
                                ],
                                "resources": [
                                    {
                                        "designatorType": "Attributes",
                                        "wlid": "",
                                        "wildwlid": "",
                                        "sid": "",
                                        "attributes": {
                                            "cluster": "minikube",
                                            "namespace": "kube-system"
                                        }
                                    }
                                ],
                                "posturePolicies": [
                                    {
                                        "frameworkName": "NSA",
                                        "controlName": "Immutable container filesystem",
                                        "ruleName": ""
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "container: kube-apiserver in pod: kube-apiserver-david-virtualbox  has  mutable filesystem",
                            "ruleStatus": "warning",
                            "packagename": "armo_builtins",
                            "alertScore": 7,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "Pod",
                                        "metadata": {
                                            "annotations": {
                                                "kubernetes.io/config.hash": "668e9396bc7b8c495d39f4a3bc479397",
                                                "kubernetes.io/config.mirror": "668e9396bc7b8c495d39f4a3bc479397",
                                                "kubernetes.io/config.seen": "2021-06-20T12:06:52.495387809+03:00",
                                                "kubernetes.io/config.source": "file"
                                            },
                                            "creationTimestamp": "2021-06-20T09:08:05Z",
                                            "labels": {
                                                "component": "kube-apiserver",
                                                "tier": "control-plane"
                                            },
                                            "name": "kube-apiserver-david-virtualbox",
                                            "namespace": "kube-system",
                                            "resourceVersion": "2110920",
                                            "selfLink": "/api/v1/namespaces/kube-system/pods/kube-apiserver-david-virtualbox",
                                            "uid": "327cbf13-97d6-42a3-8469-5acfdbe1be09"
                                        },
                                        "spec": {
                                            "containers": [
                                                {
                                                    "command": [
                                                        "kube-apiserver",
                                                        "--advertise-address=10.0.2.15",
                                                        "--allow-privileged=true",
                                                        "--authorization-mode=Node,RBAC",
                                                        "--client-ca-file=/var/lib/minikube/certs/ca.crt",
                                                        "--enable-admission-plugins=NamespaceLifecycle,LimitRanger,ServiceAccount,DefaultStorageClass,DefaultTolerationSeconds,NodeRestriction,MutatingAdmissionWebhook,ValidatingAdmissionWebhook,ResourceQuota",
                                                        "--enable-bootstrap-token-auth=true",
                                                        "--etcd-cafile=/var/lib/minikube/certs/etcd/ca.crt",
                                                        "--etcd-certfile=/var/lib/minikube/certs/apiserver-etcd-client.crt",
                                                        "--etcd-keyfile=/var/lib/minikube/certs/apiserver-etcd-client.key",
                                                        "--etcd-servers=https://127.0.0.1:2379",
                                                        "--insecure-port=0",
                                                        "--kubelet-client-certificate=/var/lib/minikube/certs/apiserver-kubelet-client.crt",
                                                        "--kubelet-client-key=/var/lib/minikube/certs/apiserver-kubelet-client.key",
                                                        "--kubelet-preferred-address-types=InternalIP,ExternalIP,Hostname",
                                                        "--proxy-client-cert-file=/var/lib/minikube/certs/front-proxy-client.crt",
                                                        "--proxy-client-key-file=/var/lib/minikube/certs/front-proxy-client.key",
                                                        "--requestheader-allowed-names=front-proxy-client",
                                                        "--requestheader-client-ca-file=/var/lib/minikube/certs/front-proxy-ca.crt",
                                                        "--requestheader-extra-headers-prefix=X-Remote-Extra-",
                                                        "--requestheader-group-headers=X-Remote-Group",
                                                        "--requestheader-username-headers=X-Remote-User",
                                                        "--secure-port=8443",
                                                        "--service-account-key-file=/var/lib/minikube/certs/sa.pub",
                                                        "--service-cluster-ip-range=10.96.0.0/12",
                                                        "--tls-cert-file=/var/lib/minikube/certs/apiserver.crt",
                                                        "--tls-private-key-file=/var/lib/minikube/certs/apiserver.key"
                                                    ],
                                                    "image": "k8s.gcr.io/kube-apiserver:v1.16.0",
                                                    "imagePullPolicy": "IfNotPresent",
                                                    "livenessProbe": {
                                                        "failureThreshold": 8,
                                                        "httpGet": {
                                                            "host": "10.0.2.15",
                                                            "path": "/healthz",
                                                            "port": 8443,
                                                            "scheme": "HTTPS"
                                                        },
                                                        "initialDelaySeconds": 15,
                                                        "periodSeconds": 10,
                                                        "successThreshold": 1,
                                                        "timeoutSeconds": 15
                                                    },
                                                    "name": "kube-apiserver",
                                                    "resources": {
                                                        "requests": {
                                                            "cpu": "250m"
                                                        }
                                                    },
                                                    "terminationMessagePath": "/dev/termination-log",
                                                    "terminationMessagePolicy": "File",
                                                    "volumeMounts": [
                                                        {
                                                            "mountPath": "/etc/ssl/certs",
                                                            "name": "ca-certs",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/etc/ca-certificates",
                                                            "name": "etc-ca-certificates",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/etc/pki",
                                                            "name": "etc-pki",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/var/lib/minikube/certs",
                                                            "name": "k8s-certs",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/usr/local/share/ca-certificates",
                                                            "name": "usr-local-share-ca-certificates",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/usr/share/ca-certificates",
                                                            "name": "usr-share-ca-certificates",
                                                            "readOnly": true
                                                        }
                                                    ]
                                                }
                                            ],
                                            "dnsPolicy": "ClusterFirst",
                                            "enableServiceLinks": true,
                                            "hostNetwork": true,
                                            "nodeName": "david-virtualbox",
                                            "priority": 2000000000,
                                            "priorityClassName": "system-cluster-critical",
                                            "restartPolicy": "Always",
                                            "schedulerName": "default-scheduler",
                                            "securityContext": {},
                                            "terminationGracePeriodSeconds": 30,
                                            "tolerations": [
                                                {
                                                    "effect": "NoExecute",
                                                    "operator": "Exists"
                                                }
                                            ],
                                            "volumes": [
                                                {
                                                    "hostPath": {
                                                        "path": "/etc/ssl/certs",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "ca-certs"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/etc/ca-certificates",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "etc-ca-certificates"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/etc/pki",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "etc-pki"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/var/lib/minikube/certs",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "k8s-certs"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/usr/local/share/ca-certificates",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "usr-local-share-ca-certificates"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/usr/share/ca-certificates",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "usr-share-ca-certificates"
                                                }
                                            ]
                                        },
                                        "status": {
                                            "conditions": [
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-03T06:28:57Z",
                                                    "status": "True",
                                                    "type": "Initialized"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-06T05:23:25Z",
                                                    "status": "True",
                                                    "type": "Ready"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-06T05:23:25Z",
                                                    "status": "True",
                                                    "type": "ContainersReady"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-03T06:28:57Z",
                                                    "status": "True",
                                                    "type": "PodScheduled"
                                                }
                                            ],
                                            "containerStatuses": [
                                                {
                                                    "containerID": "docker://c6d5933ac2752c18942e488d2d162fcc3141d1ae131be2abc3a50d2cb91e4e22",
                                                    "image": "k8s.gcr.io/kube-apiserver:v1.16.0",
                                                    "imageID": "docker-pullable://k8s.gcr.io/kube-apiserver@sha256:f4168527c91289da2708f62ae729fdde5fb484167dd05ffbb7ab666f60de96cd",
                                                    "lastState": {
                                                        "terminated": {
                                                            "containerID": "docker://3714c8295fe8b0ff97082305a44f08764e06d1a599ab3ee3ca6db15ab498a276",
                                                            "exitCode": 255,
                                                            "finishedAt": "2021-10-14T05:35:02Z",
                                                            "reason": "Error",
                                                            "startedAt": "2021-10-12T10:28:38Z"
                                                        }
                                                    },
                                                    "name": "kube-apiserver",
                                                    "ready": true,
                                                    "restartCount": 51,
                                                    "started": true,
                                                    "state": {
                                                        "running": {
                                                            "startedAt": "2021-10-14T05:35:12Z"
                                                        }
                                                    }
                                                }
                                            ],
                                            "hostIP": "10.0.2.15",
                                            "phase": "Running",
                                            "podIP": "10.0.2.15",
                                            "podIPs": [
                                                {
                                                    "ip": "10.0.2.15"
                                                }
                                            ],
                                            "qosClass": "Burstable",
                                            "startTime": "2021-10-03T06:28:57Z"
                                        }
                                    }
                                ]
                            },
                            "exception": {
                                "guid": "0e71632b-d377-4d54-a6d5-67643bcb6047",
                                "name": "exception_C-0017_kube-system_cc21d0ba28b70809f341e4809ecd51d2",
                                "attributes": {
                                    "namespaceOnly": "true"
                                },
                                "policyType": "postureExceptionPolicy",
                                "creationTime": "2021-10-14T12:13:00.213686",
                                "actions": [
                                    "alertOnly"
                                ],
                                "resources": [
                                    {
                                        "designatorType": "Attributes",
                                        "wlid": "",
                                        "wildwlid": "",
                                        "sid": "",
                                        "attributes": {
                                            "cluster": "minikube",
                                            "namespace": "kube-system"
                                        }
                                    }
                                ],
                                "posturePolicies": [
                                    {
                                        "frameworkName": "NSA",
                                        "controlName": "Immutable container filesystem",
                                        "ruleName": ""
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "container: kube-controller-manager in pod: kube-controller-manager-david-virtualbox  has  mutable filesystem",
                            "ruleStatus": "warning",
                            "packagename": "armo_builtins",
                            "alertScore": 7,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "Pod",
                                        "metadata": {
                                            "annotations": {
                                                "kubernetes.io/config.hash": "a16b2d5766eae37796e4a8ed7f8ce12a",
                                                "kubernetes.io/config.mirror": "a16b2d5766eae37796e4a8ed7f8ce12a",
                                                "kubernetes.io/config.seen": "2021-06-20T12:06:52.495389283+03:00",
                                                "kubernetes.io/config.source": "file"
                                            },
                                            "creationTimestamp": "2021-06-20T09:08:00Z",
                                            "labels": {
                                                "component": "kube-controller-manager",
                                                "tier": "control-plane"
                                            },
                                            "name": "kube-controller-manager-david-virtualbox",
                                            "namespace": "kube-system",
                                            "resourceVersion": "2110899",
                                            "selfLink": "/api/v1/namespaces/kube-system/pods/kube-controller-manager-david-virtualbox",
                                            "uid": "6ca9d32c-21c3-4c0e-8087-5445c80a2bcc"
                                        },
                                        "spec": {
                                            "containers": [
                                                {
                                                    "command": [
                                                        "kube-controller-manager",
                                                        "--allocate-node-cidrs=true",
                                                        "--authentication-kubeconfig=/etc/kubernetes/controller-manager.conf",
                                                        "--authorization-kubeconfig=/etc/kubernetes/controller-manager.conf",
                                                        "--bind-address=127.0.0.1",
                                                        "--client-ca-file=/var/lib/minikube/certs/ca.crt",
                                                        "--cluster-cidr=10.244.0.0/16",
                                                        "--cluster-signing-cert-file=/var/lib/minikube/certs/ca.crt",
                                                        "--cluster-signing-key-file=/var/lib/minikube/certs/ca.key",
                                                        "--controllers=*,bootstrapsigner,tokencleaner",
                                                        "--kubeconfig=/etc/kubernetes/controller-manager.conf",
                                                        "--leader-elect=false",
                                                        "--node-cidr-mask-size=24",
                                                        "--requestheader-client-ca-file=/var/lib/minikube/certs/front-proxy-ca.crt",
                                                        "--root-ca-file=/var/lib/minikube/certs/ca.crt",
                                                        "--service-account-private-key-file=/var/lib/minikube/certs/sa.key",
                                                        "--service-cluster-ip-range=10.96.0.0/12",
                                                        "--use-service-account-credentials=true"
                                                    ],
                                                    "image": "k8s.gcr.io/kube-controller-manager:v1.16.0",
                                                    "imagePullPolicy": "IfNotPresent",
                                                    "livenessProbe": {
                                                        "failureThreshold": 8,
                                                        "httpGet": {
                                                            "host": "127.0.0.1",
                                                            "path": "/healthz",
                                                            "port": 10252,
                                                            "scheme": "HTTP"
                                                        },
                                                        "initialDelaySeconds": 15,
                                                        "periodSeconds": 10,
                                                        "successThreshold": 1,
                                                        "timeoutSeconds": 15
                                                    },
                                                    "name": "kube-controller-manager",
                                                    "resources": {
                                                        "requests": {
                                                            "cpu": "200m"
                                                        }
                                                    },
                                                    "terminationMessagePath": "/dev/termination-log",
                                                    "terminationMessagePolicy": "File",
                                                    "volumeMounts": [
                                                        {
                                                            "mountPath": "/etc/ssl/certs",
                                                            "name": "ca-certs",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/etc/ca-certificates",
                                                            "name": "etc-ca-certificates",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/etc/pki",
                                                            "name": "etc-pki",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/usr/libexec/kubernetes/kubelet-plugins/volume/exec",
                                                            "name": "flexvolume-dir"
                                                        },
                                                        {
                                                            "mountPath": "/var/lib/minikube/certs",
                                                            "name": "k8s-certs",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/etc/kubernetes/controller-manager.conf",
                                                            "name": "kubeconfig",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/usr/local/share/ca-certificates",
                                                            "name": "usr-local-share-ca-certificates",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/usr/share/ca-certificates",
                                                            "name": "usr-share-ca-certificates",
                                                            "readOnly": true
                                                        }
                                                    ]
                                                }
                                            ],
                                            "dnsPolicy": "ClusterFirst",
                                            "enableServiceLinks": true,
                                            "hostNetwork": true,
                                            "nodeName": "david-virtualbox",
                                            "priority": 2000000000,
                                            "priorityClassName": "system-cluster-critical",
                                            "restartPolicy": "Always",
                                            "schedulerName": "default-scheduler",
                                            "securityContext": {},
                                            "terminationGracePeriodSeconds": 30,
                                            "tolerations": [
                                                {
                                                    "effect": "NoExecute",
                                                    "operator": "Exists"
                                                }
                                            ],
                                            "volumes": [
                                                {
                                                    "hostPath": {
                                                        "path": "/etc/ssl/certs",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "ca-certs"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/etc/ca-certificates",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "etc-ca-certificates"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/etc/pki",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "etc-pki"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/usr/libexec/kubernetes/kubelet-plugins/volume/exec",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "flexvolume-dir"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/var/lib/minikube/certs",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "k8s-certs"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/etc/kubernetes/controller-manager.conf",
                                                        "type": "FileOrCreate"
                                                    },
                                                    "name": "kubeconfig"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/usr/local/share/ca-certificates",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "usr-local-share-ca-certificates"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/usr/share/ca-certificates",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "usr-share-ca-certificates"
                                                }
                                            ]
                                        },
                                        "status": {
                                            "conditions": [
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-14T05:35:11Z",
                                                    "status": "True",
                                                    "type": "Initialized"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-14T05:35:13Z",
                                                    "status": "True",
                                                    "type": "Ready"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-14T05:35:13Z",
                                                    "status": "True",
                                                    "type": "ContainersReady"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-14T05:35:11Z",
                                                    "status": "True",
                                                    "type": "PodScheduled"
                                                }
                                            ],
                                            "containerStatuses": [
                                                {
                                                    "containerID": "docker://eefb83fb1b81a497703013f609a54ffc03b57157395cb5f9a128dfeffea54b58",
                                                    "image": "k8s.gcr.io/kube-controller-manager:v1.16.0",
                                                    "imageID": "docker-pullable://k8s.gcr.io/kube-controller-manager@sha256:c156a05ee9d40e3ca2ebf9337f38a10558c1fc6c9124006f128a82e6c38cdf3e",
                                                    "lastState": {
                                                        "terminated": {
                                                            "containerID": "docker://81d8dca7f5b87e8b9bd8f515e9bcf86f4614306a9c729c30aad92560f520ed54",
                                                            "exitCode": 255,
                                                            "finishedAt": "2021-10-14T05:35:02Z",
                                                            "reason": "Error",
                                                            "startedAt": "2021-10-12T10:28:38Z"
                                                        }
                                                    },
                                                    "name": "kube-controller-manager",
                                                    "ready": true,
                                                    "restartCount": 55,
                                                    "started": true,
                                                    "state": {
                                                        "running": {
                                                            "startedAt": "2021-10-14T05:35:12Z"
                                                        }
                                                    }
                                                }
                                            ],
                                            "hostIP": "10.0.2.15",
                                            "phase": "Running",
                                            "podIP": "10.0.2.15",
                                            "podIPs": [
                                                {
                                                    "ip": "10.0.2.15"
                                                }
                                            ],
                                            "qosClass": "Burstable",
                                            "startTime": "2021-10-14T05:35:11Z"
                                        }
                                    }
                                ]
                            },
                            "exception": {
                                "guid": "0e71632b-d377-4d54-a6d5-67643bcb6047",
                                "name": "exception_C-0017_kube-system_cc21d0ba28b70809f341e4809ecd51d2",
                                "attributes": {
                                    "namespaceOnly": "true"
                                },
                                "policyType": "postureExceptionPolicy",
                                "creationTime": "2021-10-14T12:13:00.213686",
                                "actions": [
                                    "alertOnly"
                                ],
                                "resources": [
                                    {
                                        "designatorType": "Attributes",
                                        "wlid": "",
                                        "wildwlid": "",
                                        "sid": "",
                                        "attributes": {
                                            "cluster": "minikube",
                                            "namespace": "kube-system"
                                        }
                                    }
                                ],
                                "posturePolicies": [
                                    {
                                        "frameworkName": "NSA",
                                        "controlName": "Immutable container filesystem",
                                        "ruleName": ""
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "container: kube-scheduler in pod: kube-scheduler-david-virtualbox  has  mutable filesystem",
                            "ruleStatus": "warning",
                            "packagename": "armo_builtins",
                            "alertScore": 7,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "Pod",
                                        "metadata": {
                                            "annotations": {
                                                "kubernetes.io/config.hash": "b3d303074fe0ca1d42a8bd9ed248df09",
                                                "kubernetes.io/config.mirror": "b3d303074fe0ca1d42a8bd9ed248df09",
                                                "kubernetes.io/config.seen": "2021-06-20T12:06:52.495381685+03:00",
                                                "kubernetes.io/config.source": "file"
                                            },
                                            "creationTimestamp": "2021-06-20T09:08:11Z",
                                            "labels": {
                                                "component": "kube-scheduler",
                                                "tier": "control-plane"
                                            },
                                            "name": "kube-scheduler-david-virtualbox",
                                            "namespace": "kube-system",
                                            "resourceVersion": "2110888",
                                            "selfLink": "/api/v1/namespaces/kube-system/pods/kube-scheduler-david-virtualbox",
                                            "uid": "226e285e-b2b3-423c-8bd4-04cfb775cbc6"
                                        },
                                        "spec": {
                                            "containers": [
                                                {
                                                    "command": [
                                                        "kube-scheduler",
                                                        "--authentication-kubeconfig=/etc/kubernetes/scheduler.conf",
                                                        "--authorization-kubeconfig=/etc/kubernetes/scheduler.conf",
                                                        "--bind-address=127.0.0.1",
                                                        "--kubeconfig=/etc/kubernetes/scheduler.conf",
                                                        "--leader-elect=false"
                                                    ],
                                                    "image": "k8s.gcr.io/kube-scheduler:v1.16.0",
                                                    "imagePullPolicy": "IfNotPresent",
                                                    "livenessProbe": {
                                                        "failureThreshold": 8,
                                                        "httpGet": {
                                                            "host": "127.0.0.1",
                                                            "path": "/healthz",
                                                            "port": 10251,
                                                            "scheme": "HTTP"
                                                        },
                                                        "initialDelaySeconds": 15,
                                                        "periodSeconds": 10,
                                                        "successThreshold": 1,
                                                        "timeoutSeconds": 15
                                                    },
                                                    "name": "kube-scheduler",
                                                    "resources": {
                                                        "requests": {
                                                            "cpu": "100m"
                                                        }
                                                    },
                                                    "terminationMessagePath": "/dev/termination-log",
                                                    "terminationMessagePolicy": "File",
                                                    "volumeMounts": [
                                                        {
                                                            "mountPath": "/etc/kubernetes/scheduler.conf",
                                                            "name": "kubeconfig",
                                                            "readOnly": true
                                                        }
                                                    ]
                                                }
                                            ],
                                            "dnsPolicy": "ClusterFirst",
                                            "enableServiceLinks": true,
                                            "hostNetwork": true,
                                            "nodeName": "david-virtualbox",
                                            "priority": 2000000000,
                                            "priorityClassName": "system-cluster-critical",
                                            "restartPolicy": "Always",
                                            "schedulerName": "default-scheduler",
                                            "securityContext": {},
                                            "terminationGracePeriodSeconds": 30,
                                            "tolerations": [
                                                {
                                                    "effect": "NoExecute",
                                                    "operator": "Exists"
                                                }
                                            ],
                                            "volumes": [
                                                {
                                                    "hostPath": {
                                                        "path": "/etc/kubernetes/scheduler.conf",
                                                        "type": "FileOrCreate"
                                                    },
                                                    "name": "kubeconfig"
                                                }
                                            ]
                                        },
                                        "status": {
                                            "conditions": [
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-14T05:35:11Z",
                                                    "status": "True",
                                                    "type": "Initialized"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-14T05:35:13Z",
                                                    "status": "True",
                                                    "type": "Ready"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-14T05:35:13Z",
                                                    "status": "True",
                                                    "type": "ContainersReady"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-14T05:35:11Z",
                                                    "status": "True",
                                                    "type": "PodScheduled"
                                                }
                                            ],
                                            "containerStatuses": [
                                                {
                                                    "containerID": "docker://53c67aa4e9ac29bf19ebaa451fe9de11d2bad3442dc6ece2dc3e03a0dc266526",
                                                    "image": "k8s.gcr.io/kube-scheduler:v1.16.0",
                                                    "imageID": "docker-pullable://k8s.gcr.io/kube-scheduler@sha256:094023ab9cd02059eb0295d234ff9ea321e0e22e4813986d7f1a1ac4dc1990d0",
                                                    "lastState": {
                                                        "terminated": {
                                                            "containerID": "docker://fbe8092ea4bbc12f8ed65da47d80b24528557bc81a7fc54948da26e76d704b4a",
                                                            "exitCode": 255,
                                                            "finishedAt": "2021-10-14T05:35:02Z",
                                                            "reason": "Error",
                                                            "startedAt": "2021-10-12T10:28:38Z"
                                                        }
                                                    },
                                                    "name": "kube-scheduler",
                                                    "ready": true,
                                                    "restartCount": 52,
                                                    "started": true,
                                                    "state": {
                                                        "running": {
                                                            "startedAt": "2021-10-14T05:35:12Z"
                                                        }
                                                    }
                                                }
                                            ],
                                            "hostIP": "10.0.2.15",
                                            "phase": "Running",
                                            "podIP": "10.0.2.15",
                                            "podIPs": [
                                                {
                                                    "ip": "10.0.2.15"
                                                }
                                            ],
                                            "qosClass": "Burstable",
                                            "startTime": "2021-10-14T05:35:11Z"
                                        }
                                    }
                                ]
                            },
                            "exception": {
                                "guid": "0e71632b-d377-4d54-a6d5-67643bcb6047",
                                "name": "exception_C-0017_kube-system_cc21d0ba28b70809f341e4809ecd51d2",
                                "attributes": {
                                    "namespaceOnly": "true"
                                },
                                "policyType": "postureExceptionPolicy",
                                "creationTime": "2021-10-14T12:13:00.213686",
                                "actions": [
                                    "alertOnly"
                                ],
                                "resources": [
                                    {
                                        "designatorType": "Attributes",
                                        "wlid": "",
                                        "wildwlid": "",
                                        "sid": "",
                                        "attributes": {
                                            "cluster": "minikube",
                                            "namespace": "kube-system"
                                        }
                                    }
                                ],
                                "posturePolicies": [
                                    {
                                        "frameworkName": "NSA",
                                        "controlName": "Immutable container filesystem",
                                        "ruleName": ""
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "container: storage-provisioner in pod: storage-provisioner  has  mutable filesystem",
                            "ruleStatus": "warning",
                            "packagename": "armo_builtins",
                            "alertScore": 7,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "Pod",
                                        "metadata": {
                                            "annotations": {
                                                "kubectl.kubernetes.io/last-applied-configuration": "{\"apiVersion\":\"v1\",\"kind\":\"Pod\",\"metadata\":{\"annotations\":{},\"labels\":{\"addonmanager.kubernetes.io/mode\":\"Reconcile\",\"integration-test\":\"storage-provisioner\"},\"name\":\"storage-provisioner\",\"namespace\":\"kube-system\"},\"spec\":{\"containers\":[{\"command\":[\"/storage-provisioner\"],\"image\":\"gcr.io/k8s-minikube/storage-provisioner:v4\",\"imagePullPolicy\":\"IfNotPresent\",\"name\":\"storage-provisioner\",\"volumeMounts\":[{\"mountPath\":\"/tmp\",\"name\":\"tmp\"}]}],\"hostNetwork\":true,\"serviceAccountName\":\"storage-provisioner\",\"volumes\":[{\"hostPath\":{\"path\":\"/tmp\",\"type\":\"Directory\"},\"name\":\"tmp\"}]}}\n"
                                            },
                                            "creationTimestamp": "2021-06-20T09:07:09Z",
                                            "labels": {
                                                "addonmanager.kubernetes.io/mode": "Reconcile",
                                                "integration-test": "storage-provisioner"
                                            },
                                            "name": "storage-provisioner",
                                            "namespace": "kube-system",
                                            "resourceVersion": "2110982",
                                            "selfLink": "/api/v1/namespaces/kube-system/pods/storage-provisioner",
                                            "uid": "ea5dc2e2-4f7a-49f4-9e88-37e8e2d741a5"
                                        },
                                        "spec": {
                                            "containers": [
                                                {
                                                    "command": [
                                                        "/storage-provisioner"
                                                    ],
                                                    "image": "gcr.io/k8s-minikube/storage-provisioner:v4",
                                                    "imagePullPolicy": "IfNotPresent",
                                                    "name": "storage-provisioner",
                                                    "resources": {},
                                                    "terminationMessagePath": "/dev/termination-log",
                                                    "terminationMessagePolicy": "File",
                                                    "volumeMounts": [
                                                        {
                                                            "mountPath": "/tmp",
                                                            "name": "tmp"
                                                        },
                                                        {
                                                            "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount",
                                                            "name": "storage-provisioner-token-bbjlq",
                                                            "readOnly": true
                                                        }
                                                    ]
                                                }
                                            ],
                                            "dnsPolicy": "ClusterFirst",
                                            "enableServiceLinks": true,
                                            "hostNetwork": true,
                                            "nodeName": "david-virtualbox",
                                            "priority": 0,
                                            "restartPolicy": "Always",
                                            "schedulerName": "default-scheduler",
                                            "securityContext": {},
                                            "serviceAccount": "storage-provisioner",
                                            "serviceAccountName": "storage-provisioner",
                                            "terminationGracePeriodSeconds": 30,
                                            "tolerations": [
                                                {
                                                    "effect": "NoExecute",
                                                    "key": "node.kubernetes.io/not-ready",
                                                    "operator": "Exists",
                                                    "tolerationSeconds": 300
                                                },
                                                {
                                                    "effect": "NoExecute",
                                                    "key": "node.kubernetes.io/unreachable",
                                                    "operator": "Exists",
                                                    "tolerationSeconds": 300
                                                }
                                            ],
                                            "volumes": [
                                                {
                                                    "hostPath": {
                                                        "path": "/tmp",
                                                        "type": "Directory"
                                                    },
                                                    "name": "tmp"
                                                },
                                                {
                                                    "name": "storage-provisioner-token-bbjlq",
                                                    "secret": {
                                                        "defaultMode": 420,
                                                        "secretName": "storage-provisioner-token-bbjlq"
                                                    }
                                                }
                                            ]
                                        },
                                        "status": {
                                            "conditions": [
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-06-20T09:07:23Z",
                                                    "status": "True",
                                                    "type": "Initialized"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-14T05:36:17Z",
                                                    "status": "True",
                                                    "type": "Ready"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-14T05:36:17Z",
                                                    "status": "True",
                                                    "type": "ContainersReady"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-06-20T09:07:23Z",
                                                    "status": "True",
                                                    "type": "PodScheduled"
                                                }
                                            ],
                                            "containerStatuses": [
                                                {
                                                    "containerID": "docker://78d5935d6367da4a887877b19e472885edcfa17c8beb3425058f517002ccac4b",
                                                    "image": "gcr.io/k8s-minikube/storage-provisioner:v4",
                                                    "imageID": "docker-pullable://gcr.io/k8s-minikube/storage-provisioner@sha256:06f83c679a723d938b8776510d979c69549ad7df516279981e23554b3e68572f",
                                                    "lastState": {
                                                        "terminated": {
                                                            "containerID": "docker://4ac8e5fb478cd62a8edeb8a14ea0e63989df016d15f48a21279eee9a19e631a3",
                                                            "exitCode": 1,
                                                            "finishedAt": "2021-10-14T05:36:01Z",
                                                            "reason": "Error",
                                                            "startedAt": "2021-10-14T05:35:30Z"
                                                        }
                                                    },
                                                    "name": "storage-provisioner",
                                                    "ready": true,
                                                    "restartCount": 94,
                                                    "started": true,
                                                    "state": {
                                                        "running": {
                                                            "startedAt": "2021-10-14T05:36:16Z"
                                                        }
                                                    }
                                                }
                                            ],
                                            "hostIP": "10.0.2.15",
                                            "phase": "Running",
                                            "podIP": "10.0.2.15",
                                            "podIPs": [
                                                {
                                                    "ip": "10.0.2.15"
                                                }
                                            ],
                                            "qosClass": "BestEffort",
                                            "startTime": "2021-06-20T09:07:23Z"
                                        }
                                    }
                                ]
                            },
                            "exception": {
                                "guid": "0e71632b-d377-4d54-a6d5-67643bcb6047",
                                "name": "exception_C-0017_kube-system_cc21d0ba28b70809f341e4809ecd51d2",
                                "attributes": {
                                    "namespaceOnly": "true"
                                },
                                "policyType": "postureExceptionPolicy",
                                "creationTime": "2021-10-14T12:13:00.213686",
                                "actions": [
                                    "alertOnly"
                                ],
                                "resources": [
                                    {
                                        "designatorType": "Attributes",
                                        "wlid": "",
                                        "wildwlid": "",
                                        "sid": "",
                                        "attributes": {
                                            "cluster": "minikube",
                                            "namespace": "kube-system"
                                        }
                                    }
                                ],
                                "posturePolicies": [
                                    {
                                        "frameworkName": "NSA",
                                        "controlName": "Immutable container filesystem",
                                        "ruleName": ""
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "container :kube-proxy in DaemonSet: kube-proxy has  mutable filesystem",
                            "ruleStatus": "warning",
                            "packagename": "armo_builtins",
                            "alertScore": 7,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "apps/v1",
                                        "kind": "DaemonSet",
                                        "metadata": {
                                            "annotations": {
                                                "deprecated.daemonset.template.generation": "1"
                                            },
                                            "creationTimestamp": "2021-06-20T09:07:08Z",
                                            "generation": 1,
                                            "labels": {
                                                "k8s-app": "kube-proxy"
                                            },
                                            "name": "kube-proxy",
                                            "namespace": "kube-system",
                                            "resourceVersion": "1450902",
                                            "selfLink": "/apis/apps/v1/namespaces/kube-system/daemonsets/kube-proxy",
                                            "uid": "dd1ba553-66da-47bc-8bc1-79c4b2f47dab"
                                        },
                                        "spec": {
                                            "revisionHistoryLimit": 10,
                                            "selector": {
                                                "matchLabels": {
                                                    "k8s-app": "kube-proxy"
                                                }
                                            },
                                            "template": {
                                                "metadata": {
                                                    "creationTimestamp": null,
                                                    "labels": {
                                                        "k8s-app": "kube-proxy"
                                                    }
                                                },
                                                "spec": {
                                                    "containers": [
                                                        {
                                                            "command": [
                                                                "/usr/local/bin/kube-proxy",
                                                                "--config=/var/lib/kube-proxy/config.conf",
                                                                "--hostname-override=$(NODE_NAME)"
                                                            ],
                                                            "env": [
                                                                {
                                                                    "name": "NODE_NAME",
                                                                    "valueFrom": {
                                                                        "fieldRef": {
                                                                            "apiVersion": "v1",
                                                                            "fieldPath": "spec.nodeName"
                                                                        }
                                                                    }
                                                                }
                                                            ],
                                                            "image": "k8s.gcr.io/kube-proxy:v1.16.0",
                                                            "imagePullPolicy": "IfNotPresent",
                                                            "name": "kube-proxy",
                                                            "resources": {},
                                                            "securityContext": {
                                                                "privileged": true
                                                            },
                                                            "terminationMessagePath": "/dev/termination-log",
                                                            "terminationMessagePolicy": "File",
                                                            "volumeMounts": [
                                                                {
                                                                    "mountPath": "/var/lib/kube-proxy",
                                                                    "name": "kube-proxy"
                                                                },
                                                                {
                                                                    "mountPath": "/run/xtables.lock",
                                                                    "name": "xtables-lock"
                                                                },
                                                                {
                                                                    "mountPath": "/lib/modules",
                                                                    "name": "lib-modules",
                                                                    "readOnly": true
                                                                }
                                                            ]
                                                        }
                                                    ],
                                                    "dnsPolicy": "ClusterFirst",
                                                    "hostNetwork": true,
                                                    "nodeSelector": {
                                                        "beta.kubernetes.io/os": "linux"
                                                    },
                                                    "priorityClassName": "system-node-critical",
                                                    "restartPolicy": "Always",
                                                    "schedulerName": "default-scheduler",
                                                    "securityContext": {},
                                                    "serviceAccount": "kube-proxy",
                                                    "serviceAccountName": "kube-proxy",
                                                    "terminationGracePeriodSeconds": 30,
                                                    "tolerations": [
                                                        {
                                                            "key": "CriticalAddonsOnly",
                                                            "operator": "Exists"
                                                        },
                                                        {
                                                            "operator": "Exists"
                                                        }
                                                    ],
                                                    "volumes": [
                                                        {
                                                            "configMap": {
                                                                "defaultMode": 420,
                                                                "name": "kube-proxy"
                                                            },
                                                            "name": "kube-proxy"
                                                        },
                                                        {
                                                            "hostPath": {
                                                                "path": "/run/xtables.lock",
                                                                "type": "FileOrCreate"
                                                            },
                                                            "name": "xtables-lock"
                                                        },
                                                        {
                                                            "hostPath": {
                                                                "path": "/lib/modules",
                                                                "type": ""
                                                            },
                                                            "name": "lib-modules"
                                                        }
                                                    ]
                                                }
                                            },
                                            "updateStrategy": {
                                                "rollingUpdate": {
                                                    "maxUnavailable": 1
                                                },
                                                "type": "RollingUpdate"
                                            }
                                        },
                                        "status": {
                                            "currentNumberScheduled": 1,
                                            "desiredNumberScheduled": 1,
                                            "numberAvailable": 1,
                                            "numberMisscheduled": 0,
                                            "numberReady": 1,
                                            "observedGeneration": 1,
                                            "updatedNumberScheduled": 1
                                        }
                                    }
                                ]
                            },
                            "exception": {
                                "guid": "0e71632b-d377-4d54-a6d5-67643bcb6047",
                                "name": "exception_C-0017_kube-system_cc21d0ba28b70809f341e4809ecd51d2",
                                "attributes": {
                                    "namespaceOnly": "true"
                                },
                                "policyType": "postureExceptionPolicy",
                                "creationTime": "2021-10-14T12:13:00.213686",
                                "actions": [
                                    "alertOnly"
                                ],
                                "resources": [
                                    {
                                        "designatorType": "Attributes",
                                        "wlid": "",
                                        "wildwlid": "",
                                        "sid": "",
                                        "attributes": {
                                            "cluster": "minikube",
                                            "namespace": "kube-system"
                                        }
                                    }
                                ],
                                "posturePolicies": [
                                    {
                                        "frameworkName": "NSA",
                                        "controlName": "Immutable container filesystem",
                                        "ruleName": ""
                                    }
                                ]
                            }
                        }
                    ]
                }
            ],
            "remediation": "Set the filesystem of the container to read-only when possible (POD securityContext, readOnlyRootFilesystem: true). If containers application needs to write into the filesystem, it is recommended to mount secondary filesystems for specific directories where application require write access.",
            "description": "Mutable container filesystem can be abused to inject malicious code or data into containers. Use immutable (read-only) filesystem to limit potential attacks.",
            "score": 100
        },
        {
            "guid": "",
            "attributes": {
                "armoBuiltin": true
            },
            "id": "C-0013",
            "controlID": "C-0013",
            "name": "Non-root containers",
            "ruleReports": [
                {
                    "name": "non-root-containers",
                    "remediation": "",
                    "ruleStatus": {
                        "status": "success",
                        "message": ""
                    },
                    "ruleResponses": null
                }
            ],
            "remediation": "If your application does not need root privileges, make sure to define the runAsUser or runAsGroup under the PodSecurityContext and use user ID 1000 or higher. Do not turn on allowPrivlegeEscalation bit and make sure runAsNonRoot is true.",
            "description": "Potential attackers may gain access to a container and leverage its existing privileges to conduct an attack. Therefore, it is not recommended to deploy containers with root privileges unless it is absolutely necessary. This contol identifies all the Pods running as root or can escalate to root.",
            "score": 100
        },
        {
            "guid": "",
            "attributes": {
                "armoBuiltin": true,
                "microsoftMitreColumns": [
                    "Privilege escalation"
                ]
            },
            "id": "C-0057",
            "controlID": "C-0057",
            "name": "Privileged container",
            "ruleReports": [
                {
                    "name": "rule-privilege-escalation",
                    "remediation": "",
                    "ruleStatus": {
                        "status": "success",
                        "message": ""
                    },
                    "ruleResponses": [
                        {
                            "alertMessage": "DaemonSet: kube-proxy is defined as privileged:",
                            "ruleStatus": "warning",
                            "packagename": "armo_builtins",
                            "alertScore": 3,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "apps/v1",
                                        "kind": "DaemonSet",
                                        "metadata": {
                                            "annotations": {
                                                "deprecated.daemonset.template.generation": "1"
                                            },
                                            "creationTimestamp": "2021-06-20T09:07:08Z",
                                            "generation": 1,
                                            "labels": {
                                                "k8s-app": "kube-proxy"
                                            },
                                            "name": "kube-proxy",
                                            "namespace": "kube-system",
                                            "resourceVersion": "1450902",
                                            "selfLink": "/apis/apps/v1/namespaces/kube-system/daemonsets/kube-proxy",
                                            "uid": "dd1ba553-66da-47bc-8bc1-79c4b2f47dab"
                                        },
                                        "spec": {
                                            "revisionHistoryLimit": 10,
                                            "selector": {
                                                "matchLabels": {
                                                    "k8s-app": "kube-proxy"
                                                }
                                            },
                                            "template": {
                                                "metadata": {
                                                    "creationTimestamp": null,
                                                    "labels": {
                                                        "k8s-app": "kube-proxy"
                                                    }
                                                },
                                                "spec": {
                                                    "containers": [
                                                        {
                                                            "command": [
                                                                "/usr/local/bin/kube-proxy",
                                                                "--config=/var/lib/kube-proxy/config.conf",
                                                                "--hostname-override=$(NODE_NAME)"
                                                            ],
                                                            "env": [
                                                                {
                                                                    "name": "NODE_NAME",
                                                                    "valueFrom": {
                                                                        "fieldRef": {
                                                                            "apiVersion": "v1",
                                                                            "fieldPath": "spec.nodeName"
                                                                        }
                                                                    }
                                                                }
                                                            ],
                                                            "image": "k8s.gcr.io/kube-proxy:v1.16.0",
                                                            "imagePullPolicy": "IfNotPresent",
                                                            "name": "kube-proxy",
                                                            "resources": {},
                                                            "securityContext": {
                                                                "privileged": true
                                                            },
                                                            "terminationMessagePath": "/dev/termination-log",
                                                            "terminationMessagePolicy": "File",
                                                            "volumeMounts": [
                                                                {
                                                                    "mountPath": "/var/lib/kube-proxy",
                                                                    "name": "kube-proxy"
                                                                },
                                                                {
                                                                    "mountPath": "/run/xtables.lock",
                                                                    "name": "xtables-lock"
                                                                },
                                                                {
                                                                    "mountPath": "/lib/modules",
                                                                    "name": "lib-modules",
                                                                    "readOnly": true
                                                                }
                                                            ]
                                                        }
                                                    ],
                                                    "dnsPolicy": "ClusterFirst",
                                                    "hostNetwork": true,
                                                    "nodeSelector": {
                                                        "beta.kubernetes.io/os": "linux"
                                                    },
                                                    "priorityClassName": "system-node-critical",
                                                    "restartPolicy": "Always",
                                                    "schedulerName": "default-scheduler",
                                                    "securityContext": {},
                                                    "serviceAccount": "kube-proxy",
                                                    "serviceAccountName": "kube-proxy",
                                                    "terminationGracePeriodSeconds": 30,
                                                    "tolerations": [
                                                        {
                                                            "key": "CriticalAddonsOnly",
                                                            "operator": "Exists"
                                                        },
                                                        {
                                                            "operator": "Exists"
                                                        }
                                                    ],
                                                    "volumes": [
                                                        {
                                                            "configMap": {
                                                                "defaultMode": 420,
                                                                "name": "kube-proxy"
                                                            },
                                                            "name": "kube-proxy"
                                                        },
                                                        {
                                                            "hostPath": {
                                                                "path": "/run/xtables.lock",
                                                                "type": "FileOrCreate"
                                                            },
                                                            "name": "xtables-lock"
                                                        },
                                                        {
                                                            "hostPath": {
                                                                "path": "/lib/modules",
                                                                "type": ""
                                                            },
                                                            "name": "lib-modules"
                                                        }
                                                    ]
                                                }
                                            },
                                            "updateStrategy": {
                                                "rollingUpdate": {
                                                    "maxUnavailable": 1
                                                },
                                                "type": "RollingUpdate"
                                            }
                                        },
                                        "status": {
                                            "currentNumberScheduled": 1,
                                            "desiredNumberScheduled": 1,
                                            "numberAvailable": 1,
                                            "numberMisscheduled": 0,
                                            "numberReady": 1,
                                            "observedGeneration": 1,
                                            "updatedNumberScheduled": 1
                                        }
                                    }
                                ]
                            },
                            "exception": {
                                "guid": "3a2a581f-6d3d-46e7-9399-a5e1118ecf71",
                                "name": "exception_C-0057_kube-system_f0daa76858044ec8af027d4e9eb0174a",
                                "attributes": {
                                    "namespaceOnly": "true"
                                },
                                "policyType": "postureExceptionPolicy",
                                "creationTime": "2021-10-14T12:12:54.978764",
                                "actions": [
                                    "alertOnly"
                                ],
                                "resources": [
                                    {
                                        "designatorType": "Attributes",
                                        "wlid": "",
                                        "wildwlid": "",
                                        "sid": "",
                                        "attributes": {
                                            "cluster": "minikube",
                                            "namespace": "kube-system"
                                        }
                                    }
                                ],
                                "posturePolicies": [
                                    {
                                        "frameworkName": "NSA",
                                        "controlName": "Privileged container",
                                        "ruleName": ""
                                    }
                                ]
                            }
                        }
                    ]
                }
            ],
            "remediation": "Remove privileged capabilities by setting the securityContext.privileged to false. If you must deploy a Pod as privileged, add other restriction to it, such as network policy, Seccomp etc and still remove all unnecessary capabilities. Use the exception mechanism to remove unnecessary notifocations.",
            "description": "Potential attackers may gain access to privileged containers and inherit access to the host resources. Therefore, it is not recommended to deploy privileged containers unless it is absolutely necessary. This control identifies all the privileged Pods.",
            "score": 100
        },
        {
            "guid": "",
            "attributes": {
                "armoBuiltin": true
            },
            "id": "C-0006",
            "controlID": "C-0006",
            "name": "Allowed hostPath",
            "ruleReports": [
                {
                    "name": "alert-rw-hostpath",
                    "remediation": "",
                    "ruleStatus": {
                        "status": "success",
                        "message": ""
                    },
                    "ruleResponses": [
                        {
                            "alertMessage": "pod: etcd-david-virtualbox has: etcd-certs as hostPath volume",
                            "ruleStatus": "",
                            "packagename": "armo_builtins",
                            "alertScore": 7,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "Pod",
                                        "metadata": {
                                            "annotations": {
                                                "kubernetes.io/config.hash": "e0fcc6e4323055b5880f8aac4c950836",
                                                "kubernetes.io/config.mirror": "e0fcc6e4323055b5880f8aac4c950836",
                                                "kubernetes.io/config.seen": "2021-06-20T12:06:52.495386281+03:00",
                                                "kubernetes.io/config.source": "file"
                                            },
                                            "creationTimestamp": "2021-06-20T09:08:22Z",
                                            "labels": {
                                                "component": "etcd",
                                                "tier": "control-plane"
                                            },
                                            "name": "etcd-david-virtualbox",
                                            "namespace": "kube-system",
                                            "resourceVersion": "2110909",
                                            "selfLink": "/api/v1/namespaces/kube-system/pods/etcd-david-virtualbox",
                                            "uid": "154e7f87-907f-4edb-a73c-26e965d4fe02"
                                        },
                                        "spec": {
                                            "containers": [
                                                {
                                                    "command": [
                                                        "etcd",
                                                        "--advertise-client-urls=https://10.0.2.15:2379",
                                                        "--cert-file=/var/lib/minikube/certs/etcd/server.crt",
                                                        "--client-cert-auth=true",
                                                        "--data-dir=/var/lib/minikube/etcd",
                                                        "--initial-advertise-peer-urls=https://10.0.2.15:2380",
                                                        "--initial-cluster=david-virtualbox=https://10.0.2.15:2380",
                                                        "--key-file=/var/lib/minikube/certs/etcd/server.key",
                                                        "--listen-client-urls=https://127.0.0.1:2379,https://10.0.2.15:2379",
                                                        "--listen-metrics-urls=http://127.0.0.1:2381,http://10.0.2.15:2381",
                                                        "--listen-peer-urls=https://10.0.2.15:2380",
                                                        "--name=david-virtualbox",
                                                        "--peer-cert-file=/var/lib/minikube/certs/etcd/peer.crt",
                                                        "--peer-client-cert-auth=true",
                                                        "--peer-key-file=/var/lib/minikube/certs/etcd/peer.key",
                                                        "--peer-trusted-ca-file=/var/lib/minikube/certs/etcd/ca.crt",
                                                        "--snapshot-count=10000",
                                                        "--trusted-ca-file=/var/lib/minikube/certs/etcd/ca.crt"
                                                    ],
                                                    "image": "k8s.gcr.io/etcd:3.3.15-0",
                                                    "imagePullPolicy": "IfNotPresent",
                                                    "livenessProbe": {
                                                        "failureThreshold": 8,
                                                        "httpGet": {
                                                            "host": "127.0.0.1",
                                                            "path": "/health",
                                                            "port": 2381,
                                                            "scheme": "HTTP"
                                                        },
                                                        "initialDelaySeconds": 15,
                                                        "periodSeconds": 10,
                                                        "successThreshold": 1,
                                                        "timeoutSeconds": 15
                                                    },
                                                    "name": "etcd",
                                                    "resources": {},
                                                    "terminationMessagePath": "/dev/termination-log",
                                                    "terminationMessagePolicy": "File",
                                                    "volumeMounts": [
                                                        {
                                                            "mountPath": "/var/lib/minikube/etcd",
                                                            "name": "etcd-data"
                                                        },
                                                        {
                                                            "mountPath": "/var/lib/minikube/certs/etcd",
                                                            "name": "etcd-certs"
                                                        }
                                                    ]
                                                }
                                            ],
                                            "dnsPolicy": "ClusterFirst",
                                            "enableServiceLinks": true,
                                            "hostNetwork": true,
                                            "nodeName": "david-virtualbox",
                                            "priority": 2000000000,
                                            "priorityClassName": "system-cluster-critical",
                                            "restartPolicy": "Always",
                                            "schedulerName": "default-scheduler",
                                            "securityContext": {},
                                            "terminationGracePeriodSeconds": 30,
                                            "tolerations": [
                                                {
                                                    "effect": "NoExecute",
                                                    "operator": "Exists"
                                                }
                                            ],
                                            "volumes": [
                                                {
                                                    "hostPath": {
                                                        "path": "/var/lib/minikube/certs/etcd",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "etcd-certs"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/var/lib/minikube/etcd",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "etcd-data"
                                                }
                                            ]
                                        },
                                        "status": {
                                            "conditions": [
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-03T06:28:57Z",
                                                    "status": "True",
                                                    "type": "Initialized"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-06T05:23:22Z",
                                                    "status": "True",
                                                    "type": "Ready"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-06T05:23:22Z",
                                                    "status": "True",
                                                    "type": "ContainersReady"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-03T06:28:57Z",
                                                    "status": "True",
                                                    "type": "PodScheduled"
                                                }
                                            ],
                                            "containerStatuses": [
                                                {
                                                    "containerID": "docker://0cc62102444cc74e3dddb2b6cd7550036feaa0e61f4297c85fc72e3be2905ec9",
                                                    "image": "k8s.gcr.io/etcd:3.3.15-0",
                                                    "imageID": "docker-pullable://k8s.gcr.io/etcd@sha256:12c2c5e5731c3bcd56e6f1c05c0f9198b6f06793fa7fca2fb43aab9622dc4afa",
                                                    "lastState": {
                                                        "terminated": {
                                                            "containerID": "docker://61de0abaa35617dafed9274c5b03f99f269412bdea2f15d903b4293d620f4b9e",
                                                            "exitCode": 255,
                                                            "finishedAt": "2021-10-14T05:35:02Z",
                                                            "reason": "Error",
                                                            "startedAt": "2021-10-12T10:28:38Z"
                                                        }
                                                    },
                                                    "name": "etcd",
                                                    "ready": true,
                                                    "restartCount": 51,
                                                    "started": true,
                                                    "state": {
                                                        "running": {
                                                            "startedAt": "2021-10-14T05:35:12Z"
                                                        }
                                                    }
                                                }
                                            ],
                                            "hostIP": "10.0.2.15",
                                            "phase": "Running",
                                            "podIP": "10.0.2.15",
                                            "podIPs": [
                                                {
                                                    "ip": "10.0.2.15"
                                                }
                                            ],
                                            "qosClass": "BestEffort",
                                            "startTime": "2021-10-03T06:28:57Z"
                                        }
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "DaemonSet: kube-proxy has: xtables-lock as hostPath volume",
                            "ruleStatus": "",
                            "packagename": "armo_builtins",
                            "alertScore": 7,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "apps/v1",
                                        "kind": "DaemonSet",
                                        "metadata": {
                                            "annotations": {
                                                "deprecated.daemonset.template.generation": "1"
                                            },
                                            "creationTimestamp": "2021-06-20T09:07:08Z",
                                            "generation": 1,
                                            "labels": {
                                                "k8s-app": "kube-proxy"
                                            },
                                            "name": "kube-proxy",
                                            "namespace": "kube-system",
                                            "resourceVersion": "1450902",
                                            "selfLink": "/apis/apps/v1/namespaces/kube-system/daemonsets/kube-proxy",
                                            "uid": "dd1ba553-66da-47bc-8bc1-79c4b2f47dab"
                                        },
                                        "spec": {
                                            "revisionHistoryLimit": 10,
                                            "selector": {
                                                "matchLabels": {
                                                    "k8s-app": "kube-proxy"
                                                }
                                            },
                                            "template": {
                                                "metadata": {
                                                    "creationTimestamp": null,
                                                    "labels": {
                                                        "k8s-app": "kube-proxy"
                                                    }
                                                },
                                                "spec": {
                                                    "containers": [
                                                        {
                                                            "command": [
                                                                "/usr/local/bin/kube-proxy",
                                                                "--config=/var/lib/kube-proxy/config.conf",
                                                                "--hostname-override=$(NODE_NAME)"
                                                            ],
                                                            "env": [
                                                                {
                                                                    "name": "NODE_NAME",
                                                                    "valueFrom": {
                                                                        "fieldRef": {
                                                                            "apiVersion": "v1",
                                                                            "fieldPath": "spec.nodeName"
                                                                        }
                                                                    }
                                                                }
                                                            ],
                                                            "image": "k8s.gcr.io/kube-proxy:v1.16.0",
                                                            "imagePullPolicy": "IfNotPresent",
                                                            "name": "kube-proxy",
                                                            "resources": {},
                                                            "securityContext": {
                                                                "privileged": true
                                                            },
                                                            "terminationMessagePath": "/dev/termination-log",
                                                            "terminationMessagePolicy": "File",
                                                            "volumeMounts": [
                                                                {
                                                                    "mountPath": "/var/lib/kube-proxy",
                                                                    "name": "kube-proxy"
                                                                },
                                                                {
                                                                    "mountPath": "/run/xtables.lock",
                                                                    "name": "xtables-lock"
                                                                },
                                                                {
                                                                    "mountPath": "/lib/modules",
                                                                    "name": "lib-modules",
                                                                    "readOnly": true
                                                                }
                                                            ]
                                                        }
                                                    ],
                                                    "dnsPolicy": "ClusterFirst",
                                                    "hostNetwork": true,
                                                    "nodeSelector": {
                                                        "beta.kubernetes.io/os": "linux"
                                                    },
                                                    "priorityClassName": "system-node-critical",
                                                    "restartPolicy": "Always",
                                                    "schedulerName": "default-scheduler",
                                                    "securityContext": {},
                                                    "serviceAccount": "kube-proxy",
                                                    "serviceAccountName": "kube-proxy",
                                                    "terminationGracePeriodSeconds": 30,
                                                    "tolerations": [
                                                        {
                                                            "key": "CriticalAddonsOnly",
                                                            "operator": "Exists"
                                                        },
                                                        {
                                                            "operator": "Exists"
                                                        }
                                                    ],
                                                    "volumes": [
                                                        {
                                                            "configMap": {
                                                                "defaultMode": 420,
                                                                "name": "kube-proxy"
                                                            },
                                                            "name": "kube-proxy"
                                                        },
                                                        {
                                                            "hostPath": {
                                                                "path": "/run/xtables.lock",
                                                                "type": "FileOrCreate"
                                                            },
                                                            "name": "xtables-lock"
                                                        },
                                                        {
                                                            "hostPath": {
                                                                "path": "/lib/modules",
                                                                "type": ""
                                                            },
                                                            "name": "lib-modules"
                                                        }
                                                    ]
                                                }
                                            },
                                            "updateStrategy": {
                                                "rollingUpdate": {
                                                    "maxUnavailable": 1
                                                },
                                                "type": "RollingUpdate"
                                            }
                                        },
                                        "status": {
                                            "currentNumberScheduled": 1,
                                            "desiredNumberScheduled": 1,
                                            "numberAvailable": 1,
                                            "numberMisscheduled": 0,
                                            "numberReady": 1,
                                            "observedGeneration": 1,
                                            "updatedNumberScheduled": 1
                                        }
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "pod: kube-controller-manager-david-virtualbox has: flexvolume-dir as hostPath volume",
                            "ruleStatus": "",
                            "packagename": "armo_builtins",
                            "alertScore": 7,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "Pod",
                                        "metadata": {
                                            "annotations": {
                                                "kubernetes.io/config.hash": "a16b2d5766eae37796e4a8ed7f8ce12a",
                                                "kubernetes.io/config.mirror": "a16b2d5766eae37796e4a8ed7f8ce12a",
                                                "kubernetes.io/config.seen": "2021-06-20T12:06:52.495389283+03:00",
                                                "kubernetes.io/config.source": "file"
                                            },
                                            "creationTimestamp": "2021-06-20T09:08:00Z",
                                            "labels": {
                                                "component": "kube-controller-manager",
                                                "tier": "control-plane"
                                            },
                                            "name": "kube-controller-manager-david-virtualbox",
                                            "namespace": "kube-system",
                                            "resourceVersion": "2110899",
                                            "selfLink": "/api/v1/namespaces/kube-system/pods/kube-controller-manager-david-virtualbox",
                                            "uid": "6ca9d32c-21c3-4c0e-8087-5445c80a2bcc"
                                        },
                                        "spec": {
                                            "containers": [
                                                {
                                                    "command": [
                                                        "kube-controller-manager",
                                                        "--allocate-node-cidrs=true",
                                                        "--authentication-kubeconfig=/etc/kubernetes/controller-manager.conf",
                                                        "--authorization-kubeconfig=/etc/kubernetes/controller-manager.conf",
                                                        "--bind-address=127.0.0.1",
                                                        "--client-ca-file=/var/lib/minikube/certs/ca.crt",
                                                        "--cluster-cidr=10.244.0.0/16",
                                                        "--cluster-signing-cert-file=/var/lib/minikube/certs/ca.crt",
                                                        "--cluster-signing-key-file=/var/lib/minikube/certs/ca.key",
                                                        "--controllers=*,bootstrapsigner,tokencleaner",
                                                        "--kubeconfig=/etc/kubernetes/controller-manager.conf",
                                                        "--leader-elect=false",
                                                        "--node-cidr-mask-size=24",
                                                        "--requestheader-client-ca-file=/var/lib/minikube/certs/front-proxy-ca.crt",
                                                        "--root-ca-file=/var/lib/minikube/certs/ca.crt",
                                                        "--service-account-private-key-file=/var/lib/minikube/certs/sa.key",
                                                        "--service-cluster-ip-range=10.96.0.0/12",
                                                        "--use-service-account-credentials=true"
                                                    ],
                                                    "image": "k8s.gcr.io/kube-controller-manager:v1.16.0",
                                                    "imagePullPolicy": "IfNotPresent",
                                                    "livenessProbe": {
                                                        "failureThreshold": 8,
                                                        "httpGet": {
                                                            "host": "127.0.0.1",
                                                            "path": "/healthz",
                                                            "port": 10252,
                                                            "scheme": "HTTP"
                                                        },
                                                        "initialDelaySeconds": 15,
                                                        "periodSeconds": 10,
                                                        "successThreshold": 1,
                                                        "timeoutSeconds": 15
                                                    },
                                                    "name": "kube-controller-manager",
                                                    "resources": {
                                                        "requests": {
                                                            "cpu": "200m"
                                                        }
                                                    },
                                                    "terminationMessagePath": "/dev/termination-log",
                                                    "terminationMessagePolicy": "File",
                                                    "volumeMounts": [
                                                        {
                                                            "mountPath": "/etc/ssl/certs",
                                                            "name": "ca-certs",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/etc/ca-certificates",
                                                            "name": "etc-ca-certificates",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/etc/pki",
                                                            "name": "etc-pki",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/usr/libexec/kubernetes/kubelet-plugins/volume/exec",
                                                            "name": "flexvolume-dir"
                                                        },
                                                        {
                                                            "mountPath": "/var/lib/minikube/certs",
                                                            "name": "k8s-certs",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/etc/kubernetes/controller-manager.conf",
                                                            "name": "kubeconfig",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/usr/local/share/ca-certificates",
                                                            "name": "usr-local-share-ca-certificates",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/usr/share/ca-certificates",
                                                            "name": "usr-share-ca-certificates",
                                                            "readOnly": true
                                                        }
                                                    ]
                                                }
                                            ],
                                            "dnsPolicy": "ClusterFirst",
                                            "enableServiceLinks": true,
                                            "hostNetwork": true,
                                            "nodeName": "david-virtualbox",
                                            "priority": 2000000000,
                                            "priorityClassName": "system-cluster-critical",
                                            "restartPolicy": "Always",
                                            "schedulerName": "default-scheduler",
                                            "securityContext": {},
                                            "terminationGracePeriodSeconds": 30,
                                            "tolerations": [
                                                {
                                                    "effect": "NoExecute",
                                                    "operator": "Exists"
                                                }
                                            ],
                                            "volumes": [
                                                {
                                                    "hostPath": {
                                                        "path": "/etc/ssl/certs",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "ca-certs"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/etc/ca-certificates",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "etc-ca-certificates"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/etc/pki",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "etc-pki"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/usr/libexec/kubernetes/kubelet-plugins/volume/exec",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "flexvolume-dir"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/var/lib/minikube/certs",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "k8s-certs"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/etc/kubernetes/controller-manager.conf",
                                                        "type": "FileOrCreate"
                                                    },
                                                    "name": "kubeconfig"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/usr/local/share/ca-certificates",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "usr-local-share-ca-certificates"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/usr/share/ca-certificates",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "usr-share-ca-certificates"
                                                }
                                            ]
                                        },
                                        "status": {
                                            "conditions": [
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-14T05:35:11Z",
                                                    "status": "True",
                                                    "type": "Initialized"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-14T05:35:13Z",
                                                    "status": "True",
                                                    "type": "Ready"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-14T05:35:13Z",
                                                    "status": "True",
                                                    "type": "ContainersReady"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-14T05:35:11Z",
                                                    "status": "True",
                                                    "type": "PodScheduled"
                                                }
                                            ],
                                            "containerStatuses": [
                                                {
                                                    "containerID": "docker://eefb83fb1b81a497703013f609a54ffc03b57157395cb5f9a128dfeffea54b58",
                                                    "image": "k8s.gcr.io/kube-controller-manager:v1.16.0",
                                                    "imageID": "docker-pullable://k8s.gcr.io/kube-controller-manager@sha256:c156a05ee9d40e3ca2ebf9337f38a10558c1fc6c9124006f128a82e6c38cdf3e",
                                                    "lastState": {
                                                        "terminated": {
                                                            "containerID": "docker://81d8dca7f5b87e8b9bd8f515e9bcf86f4614306a9c729c30aad92560f520ed54",
                                                            "exitCode": 255,
                                                            "finishedAt": "2021-10-14T05:35:02Z",
                                                            "reason": "Error",
                                                            "startedAt": "2021-10-12T10:28:38Z"
                                                        }
                                                    },
                                                    "name": "kube-controller-manager",
                                                    "ready": true,
                                                    "restartCount": 55,
                                                    "started": true,
                                                    "state": {
                                                        "running": {
                                                            "startedAt": "2021-10-14T05:35:12Z"
                                                        }
                                                    }
                                                }
                                            ],
                                            "hostIP": "10.0.2.15",
                                            "phase": "Running",
                                            "podIP": "10.0.2.15",
                                            "podIPs": [
                                                {
                                                    "ip": "10.0.2.15"
                                                }
                                            ],
                                            "qosClass": "Burstable",
                                            "startTime": "2021-10-14T05:35:11Z"
                                        }
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "pod: storage-provisioner has: tmp as hostPath volume",
                            "ruleStatus": "",
                            "packagename": "armo_builtins",
                            "alertScore": 7,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "Pod",
                                        "metadata": {
                                            "annotations": {
                                                "kubectl.kubernetes.io/last-applied-configuration": "{\"apiVersion\":\"v1\",\"kind\":\"Pod\",\"metadata\":{\"annotations\":{},\"labels\":{\"addonmanager.kubernetes.io/mode\":\"Reconcile\",\"integration-test\":\"storage-provisioner\"},\"name\":\"storage-provisioner\",\"namespace\":\"kube-system\"},\"spec\":{\"containers\":[{\"command\":[\"/storage-provisioner\"],\"image\":\"gcr.io/k8s-minikube/storage-provisioner:v4\",\"imagePullPolicy\":\"IfNotPresent\",\"name\":\"storage-provisioner\",\"volumeMounts\":[{\"mountPath\":\"/tmp\",\"name\":\"tmp\"}]}],\"hostNetwork\":true,\"serviceAccountName\":\"storage-provisioner\",\"volumes\":[{\"hostPath\":{\"path\":\"/tmp\",\"type\":\"Directory\"},\"name\":\"tmp\"}]}}\n"
                                            },
                                            "creationTimestamp": "2021-06-20T09:07:09Z",
                                            "labels": {
                                                "addonmanager.kubernetes.io/mode": "Reconcile",
                                                "integration-test": "storage-provisioner"
                                            },
                                            "name": "storage-provisioner",
                                            "namespace": "kube-system",
                                            "resourceVersion": "2110982",
                                            "selfLink": "/api/v1/namespaces/kube-system/pods/storage-provisioner",
                                            "uid": "ea5dc2e2-4f7a-49f4-9e88-37e8e2d741a5"
                                        },
                                        "spec": {
                                            "containers": [
                                                {
                                                    "command": [
                                                        "/storage-provisioner"
                                                    ],
                                                    "image": "gcr.io/k8s-minikube/storage-provisioner:v4",
                                                    "imagePullPolicy": "IfNotPresent",
                                                    "name": "storage-provisioner",
                                                    "resources": {},
                                                    "terminationMessagePath": "/dev/termination-log",
                                                    "terminationMessagePolicy": "File",
                                                    "volumeMounts": [
                                                        {
                                                            "mountPath": "/tmp",
                                                            "name": "tmp"
                                                        },
                                                        {
                                                            "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount",
                                                            "name": "storage-provisioner-token-bbjlq",
                                                            "readOnly": true
                                                        }
                                                    ]
                                                }
                                            ],
                                            "dnsPolicy": "ClusterFirst",
                                            "enableServiceLinks": true,
                                            "hostNetwork": true,
                                            "nodeName": "david-virtualbox",
                                            "priority": 0,
                                            "restartPolicy": "Always",
                                            "schedulerName": "default-scheduler",
                                            "securityContext": {},
                                            "serviceAccount": "storage-provisioner",
                                            "serviceAccountName": "storage-provisioner",
                                            "terminationGracePeriodSeconds": 30,
                                            "tolerations": [
                                                {
                                                    "effect": "NoExecute",
                                                    "key": "node.kubernetes.io/not-ready",
                                                    "operator": "Exists",
                                                    "tolerationSeconds": 300
                                                },
                                                {
                                                    "effect": "NoExecute",
                                                    "key": "node.kubernetes.io/unreachable",
                                                    "operator": "Exists",
                                                    "tolerationSeconds": 300
                                                }
                                            ],
                                            "volumes": [
                                                {
                                                    "hostPath": {
                                                        "path": "/tmp",
                                                        "type": "Directory"
                                                    },
                                                    "name": "tmp"
                                                },
                                                {
                                                    "name": "storage-provisioner-token-bbjlq",
                                                    "secret": {
                                                        "defaultMode": 420,
                                                        "secretName": "storage-provisioner-token-bbjlq"
                                                    }
                                                }
                                            ]
                                        },
                                        "status": {
                                            "conditions": [
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-06-20T09:07:23Z",
                                                    "status": "True",
                                                    "type": "Initialized"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-14T05:36:17Z",
                                                    "status": "True",
                                                    "type": "Ready"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-14T05:36:17Z",
                                                    "status": "True",
                                                    "type": "ContainersReady"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-06-20T09:07:23Z",
                                                    "status": "True",
                                                    "type": "PodScheduled"
                                                }
                                            ],
                                            "containerStatuses": [
                                                {
                                                    "containerID": "docker://78d5935d6367da4a887877b19e472885edcfa17c8beb3425058f517002ccac4b",
                                                    "image": "gcr.io/k8s-minikube/storage-provisioner:v4",
                                                    "imageID": "docker-pullable://gcr.io/k8s-minikube/storage-provisioner@sha256:06f83c679a723d938b8776510d979c69549ad7df516279981e23554b3e68572f",
                                                    "lastState": {
                                                        "terminated": {
                                                            "containerID": "docker://4ac8e5fb478cd62a8edeb8a14ea0e63989df016d15f48a21279eee9a19e631a3",
                                                            "exitCode": 1,
                                                            "finishedAt": "2021-10-14T05:36:01Z",
                                                            "reason": "Error",
                                                            "startedAt": "2021-10-14T05:35:30Z"
                                                        }
                                                    },
                                                    "name": "storage-provisioner",
                                                    "ready": true,
                                                    "restartCount": 94,
                                                    "started": true,
                                                    "state": {
                                                        "running": {
                                                            "startedAt": "2021-10-14T05:36:16Z"
                                                        }
                                                    }
                                                }
                                            ],
                                            "hostIP": "10.0.2.15",
                                            "phase": "Running",
                                            "podIP": "10.0.2.15",
                                            "podIPs": [
                                                {
                                                    "ip": "10.0.2.15"
                                                }
                                            ],
                                            "qosClass": "BestEffort",
                                            "startTime": "2021-06-20T09:07:23Z"
                                        }
                                    }
                                ]
                            }
                        }
                    ]
                }
            ],
            "remediation": "Refrain from using host path mount.",
            "description": "Mounting host directory to the container can be abused to get access to sensitive data and gain persistence on the host machine.",
            "score": 42
        },
        {
            "guid": "",
            "attributes": {
                "armoBuiltin": true
            },
            "id": "C-0034",
            "controlID": "C-0034",
            "name": "Automatic mapping of service account",
            "ruleReports": [
                {
                    "name": "automount-service-account",
                    "remediation": "",
                    "ruleStatus": {
                        "status": "success",
                        "message": ""
                    },
                    "ruleResponses": [
                        {
                            "alertMessage": "the following service account: default in the following namespace: argocd mounts service account tokens in pods by default",
                            "ruleStatus": "failed",
                            "packagename": "armo_builtins",
                            "alertScore": 9,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "ServiceAccount",
                                        "metadata": {
                                            "creationTimestamp": "2021-08-31T13:59:37Z",
                                            "name": "default",
                                            "namespace": "argocd",
                                            "resourceVersion": "1635076",
                                            "selfLink": "/api/v1/namespaces/argocd/serviceaccounts/default",
                                            "uid": "a112204c-c9a6-4a1d-909d-32ffadbc27c7"
                                        },
                                        "secrets": [
                                            {
                                                "name": "default-token-8grjt"
                                            }
                                        ]
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "the following service account: default in the following namespace: default mounts service account tokens in pods by default",
                            "ruleStatus": "failed",
                            "packagename": "armo_builtins",
                            "alertScore": 9,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "ServiceAccount",
                                        "metadata": {
                                            "creationTimestamp": "2021-06-20T09:07:22Z",
                                            "name": "default",
                                            "namespace": "default",
                                            "resourceVersion": "316",
                                            "selfLink": "/api/v1/namespaces/default/serviceaccounts/default",
                                            "uid": "b73c6821-efe6-4dbd-987a-2c260f9791d6"
                                        },
                                        "secrets": [
                                            {
                                                "name": "default-token-x6lzl"
                                            }
                                        ]
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "the following service account: polaris in the following namespace: default mounts service account tokens in pods by default",
                            "ruleStatus": "failed",
                            "packagename": "armo_builtins",
                            "alertScore": 9,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "ServiceAccount",
                                        "metadata": {
                                            "annotations": {
                                                "meta.helm.sh/release-name": "polaris",
                                                "meta.helm.sh/release-namespace": "default"
                                            },
                                            "creationTimestamp": "2021-08-09T06:28:43Z",
                                            "labels": {
                                                "app": "polaris",
                                                "app.kubernetes.io/component": "dashboard",
                                                "app.kubernetes.io/instance": "polaris",
                                                "app.kubernetes.io/managed-by": "Helm",
                                                "app.kubernetes.io/name": "polaris",
                                                "app.kubernetes.io/part-of": "polaris",
                                                "app.kubernetes.io/version": "4.0.6",
                                                "helm.sh/chart": "polaris-4.0.6"
                                            },
                                            "name": "polaris",
                                            "namespace": "default",
                                            "resourceVersion": "1136720",
                                            "selfLink": "/api/v1/namespaces/default/serviceaccounts/polaris",
                                            "uid": "2155d340-5363-47f6-ae56-fb268df8eddc"
                                        },
                                        "secrets": [
                                            {
                                                "name": "polaris-token-2w67h"
                                            }
                                        ]
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "the following service account: default in the following namespace: kube-node-lease mounts service account tokens in pods by default",
                            "ruleStatus": "failed",
                            "packagename": "armo_builtins",
                            "alertScore": 9,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "ServiceAccount",
                                        "metadata": {
                                            "creationTimestamp": "2021-06-20T09:07:22Z",
                                            "name": "default",
                                            "namespace": "kube-node-lease",
                                            "resourceVersion": "313",
                                            "selfLink": "/api/v1/namespaces/kube-node-lease/serviceaccounts/default",
                                            "uid": "77a49d18-561d-4e29-8c8d-50b9adaa1330"
                                        },
                                        "secrets": [
                                            {
                                                "name": "default-token-mrhzm"
                                            }
                                        ]
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "the following service account: default in the following namespace: kube-public mounts service account tokens in pods by default",
                            "ruleStatus": "failed",
                            "packagename": "armo_builtins",
                            "alertScore": 9,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "ServiceAccount",
                                        "metadata": {
                                            "creationTimestamp": "2021-06-20T09:07:22Z",
                                            "name": "default",
                                            "namespace": "kube-public",
                                            "resourceVersion": "310",
                                            "selfLink": "/api/v1/namespaces/kube-public/serviceaccounts/default",
                                            "uid": "c68a25b5-6362-4fd8-8bda-85520398f67b"
                                        },
                                        "secrets": [
                                            {
                                                "name": "default-token-h4z28"
                                            }
                                        ]
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "the following service account: attachdetach-controller in the following namespace: kube-system mounts service account tokens in pods by default",
                            "ruleStatus": "warning",
                            "packagename": "armo_builtins",
                            "alertScore": 9,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "ServiceAccount",
                                        "metadata": {
                                            "creationTimestamp": "2021-06-20T09:07:10Z",
                                            "name": "attachdetach-controller",
                                            "namespace": "kube-system",
                                            "resourceVersion": "258",
                                            "selfLink": "/api/v1/namespaces/kube-system/serviceaccounts/attachdetach-controller",
                                            "uid": "267b387b-85e9-4f74-a6e5-976216db3894"
                                        },
                                        "secrets": [
                                            {
                                                "name": "attachdetach-controller-token-jhb2h"
                                            }
                                        ]
                                    }
                                ]
                            },
                            "exception": {
                                "guid": "f0c9dd3c-4c54-4485-b6f9-826b5f2a3dd3",
                                "name": "exception_C-0034_kube-system_d365c088322e7c9db2ca33edcbb28e9a",
                                "attributes": {
                                    "namespaceOnly": "true"
                                },
                                "policyType": "postureExceptionPolicy",
                                "creationTime": "2021-10-14T12:13:15.128513",
                                "actions": [
                                    "alertOnly"
                                ],
                                "resources": [
                                    {
                                        "designatorType": "Attributes",
                                        "wlid": "",
                                        "wildwlid": "",
                                        "sid": "",
                                        "attributes": {
                                            "cluster": "minikube",
                                            "namespace": "kube-system"
                                        }
                                    }
                                ],
                                "posturePolicies": [
                                    {
                                        "frameworkName": "NSA",
                                        "controlName": "Automatic mapping of service account",
                                        "ruleName": ""
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "the following service account: bootstrap-signer in the following namespace: kube-system mounts service account tokens in pods by default",
                            "ruleStatus": "warning",
                            "packagename": "armo_builtins",
                            "alertScore": 9,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "ServiceAccount",
                                        "metadata": {
                                            "creationTimestamp": "2021-06-20T09:07:10Z",
                                            "name": "bootstrap-signer",
                                            "namespace": "kube-system",
                                            "resourceVersion": "264",
                                            "selfLink": "/api/v1/namespaces/kube-system/serviceaccounts/bootstrap-signer",
                                            "uid": "f19d1763-94e9-4fb9-9790-2e84c5a5ef33"
                                        },
                                        "secrets": [
                                            {
                                                "name": "bootstrap-signer-token-mhn9d"
                                            }
                                        ]
                                    }
                                ]
                            },
                            "exception": {
                                "guid": "f0c9dd3c-4c54-4485-b6f9-826b5f2a3dd3",
                                "name": "exception_C-0034_kube-system_d365c088322e7c9db2ca33edcbb28e9a",
                                "attributes": {
                                    "namespaceOnly": "true"
                                },
                                "policyType": "postureExceptionPolicy",
                                "creationTime": "2021-10-14T12:13:15.128513",
                                "actions": [
                                    "alertOnly"
                                ],
                                "resources": [
                                    {
                                        "designatorType": "Attributes",
                                        "wlid": "",
                                        "wildwlid": "",
                                        "sid": "",
                                        "attributes": {
                                            "cluster": "minikube",
                                            "namespace": "kube-system"
                                        }
                                    }
                                ],
                                "posturePolicies": [
                                    {
                                        "frameworkName": "NSA",
                                        "controlName": "Automatic mapping of service account",
                                        "ruleName": ""
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "the following service account: certificate-controller in the following namespace: kube-system mounts service account tokens in pods by default",
                            "ruleStatus": "warning",
                            "packagename": "armo_builtins",
                            "alertScore": 9,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "ServiceAccount",
                                        "metadata": {
                                            "creationTimestamp": "2021-06-20T09:07:08Z",
                                            "name": "certificate-controller",
                                            "namespace": "kube-system",
                                            "resourceVersion": "226",
                                            "selfLink": "/api/v1/namespaces/kube-system/serviceaccounts/certificate-controller",
                                            "uid": "00035670-ee8a-4f8c-9e76-dbce5bfaf9a8"
                                        },
                                        "secrets": [
                                            {
                                                "name": "certificate-controller-token-l57qw"
                                            }
                                        ]
                                    }
                                ]
                            },
                            "exception": {
                                "guid": "f0c9dd3c-4c54-4485-b6f9-826b5f2a3dd3",
                                "name": "exception_C-0034_kube-system_d365c088322e7c9db2ca33edcbb28e9a",
                                "attributes": {
                                    "namespaceOnly": "true"
                                },
                                "policyType": "postureExceptionPolicy",
                                "creationTime": "2021-10-14T12:13:15.128513",
                                "actions": [
                                    "alertOnly"
                                ],
                                "resources": [
                                    {
                                        "designatorType": "Attributes",
                                        "wlid": "",
                                        "wildwlid": "",
                                        "sid": "",
                                        "attributes": {
                                            "cluster": "minikube",
                                            "namespace": "kube-system"
                                        }
                                    }
                                ],
                                "posturePolicies": [
                                    {
                                        "frameworkName": "NSA",
                                        "controlName": "Automatic mapping of service account",
                                        "ruleName": ""
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "the following service account: clusterrole-aggregation-controller in the following namespace: kube-system mounts service account tokens in pods by default",
                            "ruleStatus": "warning",
                            "packagename": "armo_builtins",
                            "alertScore": 9,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "ServiceAccount",
                                        "metadata": {
                                            "creationTimestamp": "2021-06-20T09:07:22Z",
                                            "name": "clusterrole-aggregation-controller",
                                            "namespace": "kube-system",
                                            "resourceVersion": "301",
                                            "selfLink": "/api/v1/namespaces/kube-system/serviceaccounts/clusterrole-aggregation-controller",
                                            "uid": "af591e31-4f92-4ec9-a73f-9daf38582561"
                                        },
                                        "secrets": [
                                            {
                                                "name": "clusterrole-aggregation-controller-token-zp4tc"
                                            }
                                        ]
                                    }
                                ]
                            },
                            "exception": {
                                "guid": "f0c9dd3c-4c54-4485-b6f9-826b5f2a3dd3",
                                "name": "exception_C-0034_kube-system_d365c088322e7c9db2ca33edcbb28e9a",
                                "attributes": {
                                    "namespaceOnly": "true"
                                },
                                "policyType": "postureExceptionPolicy",
                                "creationTime": "2021-10-14T12:13:15.128513",
                                "actions": [
                                    "alertOnly"
                                ],
                                "resources": [
                                    {
                                        "designatorType": "Attributes",
                                        "wlid": "",
                                        "wildwlid": "",
                                        "sid": "",
                                        "attributes": {
                                            "cluster": "minikube",
                                            "namespace": "kube-system"
                                        }
                                    }
                                ],
                                "posturePolicies": [
                                    {
                                        "frameworkName": "NSA",
                                        "controlName": "Automatic mapping of service account",
                                        "ruleName": ""
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "the following service account: coredns in the following namespace: kube-system mounts service account tokens in pods by default",
                            "ruleStatus": "warning",
                            "packagename": "armo_builtins",
                            "alertScore": 9,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "ServiceAccount",
                                        "metadata": {
                                            "creationTimestamp": "2021-06-20T09:07:08Z",
                                            "name": "coredns",
                                            "namespace": "kube-system",
                                            "resourceVersion": "196",
                                            "selfLink": "/api/v1/namespaces/kube-system/serviceaccounts/coredns",
                                            "uid": "44032b7d-5857-4cb5-b754-08942e123509"
                                        },
                                        "secrets": [
                                            {
                                                "name": "coredns-token-jfk4n"
                                            }
                                        ]
                                    }
                                ]
                            },
                            "exception": {
                                "guid": "f0c9dd3c-4c54-4485-b6f9-826b5f2a3dd3",
                                "name": "exception_C-0034_kube-system_d365c088322e7c9db2ca33edcbb28e9a",
                                "attributes": {
                                    "namespaceOnly": "true"
                                },
                                "policyType": "postureExceptionPolicy",
                                "creationTime": "2021-10-14T12:13:15.128513",
                                "actions": [
                                    "alertOnly"
                                ],
                                "resources": [
                                    {
                                        "designatorType": "Attributes",
                                        "wlid": "",
                                        "wildwlid": "",
                                        "sid": "",
                                        "attributes": {
                                            "cluster": "minikube",
                                            "namespace": "kube-system"
                                        }
                                    }
                                ],
                                "posturePolicies": [
                                    {
                                        "frameworkName": "NSA",
                                        "controlName": "Automatic mapping of service account",
                                        "ruleName": ""
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "the following service account: cronjob-controller in the following namespace: kube-system mounts service account tokens in pods by default",
                            "ruleStatus": "warning",
                            "packagename": "armo_builtins",
                            "alertScore": 9,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "ServiceAccount",
                                        "metadata": {
                                            "creationTimestamp": "2021-06-20T09:07:12Z",
                                            "name": "cronjob-controller",
                                            "namespace": "kube-system",
                                            "resourceVersion": "284",
                                            "selfLink": "/api/v1/namespaces/kube-system/serviceaccounts/cronjob-controller",
                                            "uid": "2306bc0a-3e40-4964-859c-36f4afe6a7ef"
                                        },
                                        "secrets": [
                                            {
                                                "name": "cronjob-controller-token-qdncz"
                                            }
                                        ]
                                    }
                                ]
                            },
                            "exception": {
                                "guid": "f0c9dd3c-4c54-4485-b6f9-826b5f2a3dd3",
                                "name": "exception_C-0034_kube-system_d365c088322e7c9db2ca33edcbb28e9a",
                                "attributes": {
                                    "namespaceOnly": "true"
                                },
                                "policyType": "postureExceptionPolicy",
                                "creationTime": "2021-10-14T12:13:15.128513",
                                "actions": [
                                    "alertOnly"
                                ],
                                "resources": [
                                    {
                                        "designatorType": "Attributes",
                                        "wlid": "",
                                        "wildwlid": "",
                                        "sid": "",
                                        "attributes": {
                                            "cluster": "minikube",
                                            "namespace": "kube-system"
                                        }
                                    }
                                ],
                                "posturePolicies": [
                                    {
                                        "frameworkName": "NSA",
                                        "controlName": "Automatic mapping of service account",
                                        "ruleName": ""
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "the following service account: daemon-set-controller in the following namespace: kube-system mounts service account tokens in pods by default",
                            "ruleStatus": "warning",
                            "packagename": "armo_builtins",
                            "alertScore": 9,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "ServiceAccount",
                                        "metadata": {
                                            "creationTimestamp": "2021-06-20T09:07:09Z",
                                            "name": "daemon-set-controller",
                                            "namespace": "kube-system",
                                            "resourceVersion": "250",
                                            "selfLink": "/api/v1/namespaces/kube-system/serviceaccounts/daemon-set-controller",
                                            "uid": "917484ff-7fed-494b-9f98-681889c427ac"
                                        },
                                        "secrets": [
                                            {
                                                "name": "daemon-set-controller-token-jgx68"
                                            }
                                        ]
                                    }
                                ]
                            },
                            "exception": {
                                "guid": "f0c9dd3c-4c54-4485-b6f9-826b5f2a3dd3",
                                "name": "exception_C-0034_kube-system_d365c088322e7c9db2ca33edcbb28e9a",
                                "attributes": {
                                    "namespaceOnly": "true"
                                },
                                "policyType": "postureExceptionPolicy",
                                "creationTime": "2021-10-14T12:13:15.128513",
                                "actions": [
                                    "alertOnly"
                                ],
                                "resources": [
                                    {
                                        "designatorType": "Attributes",
                                        "wlid": "",
                                        "wildwlid": "",
                                        "sid": "",
                                        "attributes": {
                                            "cluster": "minikube",
                                            "namespace": "kube-system"
                                        }
                                    }
                                ],
                                "posturePolicies": [
                                    {
                                        "frameworkName": "NSA",
                                        "controlName": "Automatic mapping of service account",
                                        "ruleName": ""
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "the following service account: default in the following namespace: kube-system mounts service account tokens in pods by default",
                            "ruleStatus": "warning",
                            "packagename": "armo_builtins",
                            "alertScore": 9,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "ServiceAccount",
                                        "metadata": {
                                            "creationTimestamp": "2021-06-20T09:07:22Z",
                                            "name": "default",
                                            "namespace": "kube-system",
                                            "resourceVersion": "317",
                                            "selfLink": "/api/v1/namespaces/kube-system/serviceaccounts/default",
                                            "uid": "d4582926-3998-4690-a70b-c95a23369341"
                                        },
                                        "secrets": [
                                            {
                                                "name": "default-token-q8bfg"
                                            }
                                        ]
                                    }
                                ]
                            },
                            "exception": {
                                "guid": "f0c9dd3c-4c54-4485-b6f9-826b5f2a3dd3",
                                "name": "exception_C-0034_kube-system_d365c088322e7c9db2ca33edcbb28e9a",
                                "attributes": {
                                    "namespaceOnly": "true"
                                },
                                "policyType": "postureExceptionPolicy",
                                "creationTime": "2021-10-14T12:13:15.128513",
                                "actions": [
                                    "alertOnly"
                                ],
                                "resources": [
                                    {
                                        "designatorType": "Attributes",
                                        "wlid": "",
                                        "wildwlid": "",
                                        "sid": "",
                                        "attributes": {
                                            "cluster": "minikube",
                                            "namespace": "kube-system"
                                        }
                                    }
                                ],
                                "posturePolicies": [
                                    {
                                        "frameworkName": "NSA",
                                        "controlName": "Automatic mapping of service account",
                                        "ruleName": ""
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "the following service account: deployment-controller in the following namespace: kube-system mounts service account tokens in pods by default",
                            "ruleStatus": "warning",
                            "packagename": "armo_builtins",
                            "alertScore": 9,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "ServiceAccount",
                                        "metadata": {
                                            "creationTimestamp": "2021-06-20T09:07:11Z",
                                            "name": "deployment-controller",
                                            "namespace": "kube-system",
                                            "resourceVersion": "278",
                                            "selfLink": "/api/v1/namespaces/kube-system/serviceaccounts/deployment-controller",
                                            "uid": "a8855f45-6f99-4597-9465-e97da41f394a"
                                        },
                                        "secrets": [
                                            {
                                                "name": "deployment-controller-token-g68t5"
                                            }
                                        ]
                                    }
                                ]
                            },
                            "exception": {
                                "guid": "f0c9dd3c-4c54-4485-b6f9-826b5f2a3dd3",
                                "name": "exception_C-0034_kube-system_d365c088322e7c9db2ca33edcbb28e9a",
                                "attributes": {
                                    "namespaceOnly": "true"
                                },
                                "policyType": "postureExceptionPolicy",
                                "creationTime": "2021-10-14T12:13:15.128513",
                                "actions": [
                                    "alertOnly"
                                ],
                                "resources": [
                                    {
                                        "designatorType": "Attributes",
                                        "wlid": "",
                                        "wildwlid": "",
                                        "sid": "",
                                        "attributes": {
                                            "cluster": "minikube",
                                            "namespace": "kube-system"
                                        }
                                    }
                                ],
                                "posturePolicies": [
                                    {
                                        "frameworkName": "NSA",
                                        "controlName": "Automatic mapping of service account",
                                        "ruleName": ""
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "the following service account: disruption-controller in the following namespace: kube-system mounts service account tokens in pods by default",
                            "ruleStatus": "warning",
                            "packagename": "armo_builtins",
                            "alertScore": 9,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "ServiceAccount",
                                        "metadata": {
                                            "creationTimestamp": "2021-06-20T09:07:08Z",
                                            "name": "disruption-controller",
                                            "namespace": "kube-system",
                                            "resourceVersion": "218",
                                            "selfLink": "/api/v1/namespaces/kube-system/serviceaccounts/disruption-controller",
                                            "uid": "53692379-90d3-430f-bb59-c40187663ff5"
                                        },
                                        "secrets": [
                                            {
                                                "name": "disruption-controller-token-pmmz7"
                                            }
                                        ]
                                    }
                                ]
                            },
                            "exception": {
                                "guid": "f0c9dd3c-4c54-4485-b6f9-826b5f2a3dd3",
                                "name": "exception_C-0034_kube-system_d365c088322e7c9db2ca33edcbb28e9a",
                                "attributes": {
                                    "namespaceOnly": "true"
                                },
                                "policyType": "postureExceptionPolicy",
                                "creationTime": "2021-10-14T12:13:15.128513",
                                "actions": [
                                    "alertOnly"
                                ],
                                "resources": [
                                    {
                                        "designatorType": "Attributes",
                                        "wlid": "",
                                        "wildwlid": "",
                                        "sid": "",
                                        "attributes": {
                                            "cluster": "minikube",
                                            "namespace": "kube-system"
                                        }
                                    }
                                ],
                                "posturePolicies": [
                                    {
                                        "frameworkName": "NSA",
                                        "controlName": "Automatic mapping of service account",
                                        "ruleName": ""
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "the following service account: endpoint-controller in the following namespace: kube-system mounts service account tokens in pods by default",
                            "ruleStatus": "warning",
                            "packagename": "armo_builtins",
                            "alertScore": 9,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "ServiceAccount",
                                        "metadata": {
                                            "creationTimestamp": "2021-06-20T09:07:09Z",
                                            "name": "endpoint-controller",
                                            "namespace": "kube-system",
                                            "resourceVersion": "240",
                                            "selfLink": "/api/v1/namespaces/kube-system/serviceaccounts/endpoint-controller",
                                            "uid": "067dfe84-b48b-4a08-9d58-f4f03944fb12"
                                        },
                                        "secrets": [
                                            {
                                                "name": "endpoint-controller-token-s2j5n"
                                            }
                                        ]
                                    }
                                ]
                            },
                            "exception": {
                                "guid": "f0c9dd3c-4c54-4485-b6f9-826b5f2a3dd3",
                                "name": "exception_C-0034_kube-system_d365c088322e7c9db2ca33edcbb28e9a",
                                "attributes": {
                                    "namespaceOnly": "true"
                                },
                                "policyType": "postureExceptionPolicy",
                                "creationTime": "2021-10-14T12:13:15.128513",
                                "actions": [
                                    "alertOnly"
                                ],
                                "resources": [
                                    {
                                        "designatorType": "Attributes",
                                        "wlid": "",
                                        "wildwlid": "",
                                        "sid": "",
                                        "attributes": {
                                            "cluster": "minikube",
                                            "namespace": "kube-system"
                                        }
                                    }
                                ],
                                "posturePolicies": [
                                    {
                                        "frameworkName": "NSA",
                                        "controlName": "Automatic mapping of service account",
                                        "ruleName": ""
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "the following service account: expand-controller in the following namespace: kube-system mounts service account tokens in pods by default",
                            "ruleStatus": "warning",
                            "packagename": "armo_builtins",
                            "alertScore": 9,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "ServiceAccount",
                                        "metadata": {
                                            "creationTimestamp": "2021-06-20T09:07:22Z",
                                            "name": "expand-controller",
                                            "namespace": "kube-system",
                                            "resourceVersion": "298",
                                            "selfLink": "/api/v1/namespaces/kube-system/serviceaccounts/expand-controller",
                                            "uid": "09c3fa61-b2af-41e2-8eee-a0f6423b70fc"
                                        },
                                        "secrets": [
                                            {
                                                "name": "expand-controller-token-4t7tj"
                                            }
                                        ]
                                    }
                                ]
                            },
                            "exception": {
                                "guid": "f0c9dd3c-4c54-4485-b6f9-826b5f2a3dd3",
                                "name": "exception_C-0034_kube-system_d365c088322e7c9db2ca33edcbb28e9a",
                                "attributes": {
                                    "namespaceOnly": "true"
                                },
                                "policyType": "postureExceptionPolicy",
                                "creationTime": "2021-10-14T12:13:15.128513",
                                "actions": [
                                    "alertOnly"
                                ],
                                "resources": [
                                    {
                                        "designatorType": "Attributes",
                                        "wlid": "",
                                        "wildwlid": "",
                                        "sid": "",
                                        "attributes": {
                                            "cluster": "minikube",
                                            "namespace": "kube-system"
                                        }
                                    }
                                ],
                                "posturePolicies": [
                                    {
                                        "frameworkName": "NSA",
                                        "controlName": "Automatic mapping of service account",
                                        "ruleName": ""
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "the following service account: generic-garbage-collector in the following namespace: kube-system mounts service account tokens in pods by default",
                            "ruleStatus": "warning",
                            "packagename": "armo_builtins",
                            "alertScore": 9,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "ServiceAccount",
                                        "metadata": {
                                            "creationTimestamp": "2021-06-20T09:07:11Z",
                                            "name": "generic-garbage-collector",
                                            "namespace": "kube-system",
                                            "resourceVersion": "272",
                                            "selfLink": "/api/v1/namespaces/kube-system/serviceaccounts/generic-garbage-collector",
                                            "uid": "5257b086-51a0-40d3-aae5-e80c39d36214"
                                        },
                                        "secrets": [
                                            {
                                                "name": "generic-garbage-collector-token-h4z87"
                                            }
                                        ]
                                    }
                                ]
                            },
                            "exception": {
                                "guid": "f0c9dd3c-4c54-4485-b6f9-826b5f2a3dd3",
                                "name": "exception_C-0034_kube-system_d365c088322e7c9db2ca33edcbb28e9a",
                                "attributes": {
                                    "namespaceOnly": "true"
                                },
                                "policyType": "postureExceptionPolicy",
                                "creationTime": "2021-10-14T12:13:15.128513",
                                "actions": [
                                    "alertOnly"
                                ],
                                "resources": [
                                    {
                                        "designatorType": "Attributes",
                                        "wlid": "",
                                        "wildwlid": "",
                                        "sid": "",
                                        "attributes": {
                                            "cluster": "minikube",
                                            "namespace": "kube-system"
                                        }
                                    }
                                ],
                                "posturePolicies": [
                                    {
                                        "frameworkName": "NSA",
                                        "controlName": "Automatic mapping of service account",
                                        "ruleName": ""
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "the following service account: horizontal-pod-autoscaler in the following namespace: kube-system mounts service account tokens in pods by default",
                            "ruleStatus": "warning",
                            "packagename": "armo_builtins",
                            "alertScore": 9,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "ServiceAccount",
                                        "metadata": {
                                            "creationTimestamp": "2021-06-20T09:07:07Z",
                                            "name": "horizontal-pod-autoscaler",
                                            "namespace": "kube-system",
                                            "resourceVersion": "170",
                                            "selfLink": "/api/v1/namespaces/kube-system/serviceaccounts/horizontal-pod-autoscaler",
                                            "uid": "d274d4f0-d256-4ec7-b63f-dad2eabc77df"
                                        },
                                        "secrets": [
                                            {
                                                "name": "horizontal-pod-autoscaler-token-5dj9b"
                                            }
                                        ]
                                    }
                                ]
                            },
                            "exception": {
                                "guid": "f0c9dd3c-4c54-4485-b6f9-826b5f2a3dd3",
                                "name": "exception_C-0034_kube-system_d365c088322e7c9db2ca33edcbb28e9a",
                                "attributes": {
                                    "namespaceOnly": "true"
                                },
                                "policyType": "postureExceptionPolicy",
                                "creationTime": "2021-10-14T12:13:15.128513",
                                "actions": [
                                    "alertOnly"
                                ],
                                "resources": [
                                    {
                                        "designatorType": "Attributes",
                                        "wlid": "",
                                        "wildwlid": "",
                                        "sid": "",
                                        "attributes": {
                                            "cluster": "minikube",
                                            "namespace": "kube-system"
                                        }
                                    }
                                ],
                                "posturePolicies": [
                                    {
                                        "frameworkName": "NSA",
                                        "controlName": "Automatic mapping of service account",
                                        "ruleName": ""
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "the following service account: job-controller in the following namespace: kube-system mounts service account tokens in pods by default",
                            "ruleStatus": "warning",
                            "packagename": "armo_builtins",
                            "alertScore": 9,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "ServiceAccount",
                                        "metadata": {
                                            "creationTimestamp": "2021-06-20T09:07:11Z",
                                            "name": "job-controller",
                                            "namespace": "kube-system",
                                            "resourceVersion": "275",
                                            "selfLink": "/api/v1/namespaces/kube-system/serviceaccounts/job-controller",
                                            "uid": "c2206aff-9fe6-4e58-a480-290f9adbbedd"
                                        },
                                        "secrets": [
                                            {
                                                "name": "job-controller-token-f94hz"
                                            }
                                        ]
                                    }
                                ]
                            },
                            "exception": {
                                "guid": "f0c9dd3c-4c54-4485-b6f9-826b5f2a3dd3",
                                "name": "exception_C-0034_kube-system_d365c088322e7c9db2ca33edcbb28e9a",
                                "attributes": {
                                    "namespaceOnly": "true"
                                },
                                "policyType": "postureExceptionPolicy",
                                "creationTime": "2021-10-14T12:13:15.128513",
                                "actions": [
                                    "alertOnly"
                                ],
                                "resources": [
                                    {
                                        "designatorType": "Attributes",
                                        "wlid": "",
                                        "wildwlid": "",
                                        "sid": "",
                                        "attributes": {
                                            "cluster": "minikube",
                                            "namespace": "kube-system"
                                        }
                                    }
                                ],
                                "posturePolicies": [
                                    {
                                        "frameworkName": "NSA",
                                        "controlName": "Automatic mapping of service account",
                                        "ruleName": ""
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "the following service account: kube-proxy in the following namespace: kube-system mounts service account tokens in pods by default",
                            "ruleStatus": "warning",
                            "packagename": "armo_builtins",
                            "alertScore": 9,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "ServiceAccount",
                                        "metadata": {
                                            "creationTimestamp": "2021-06-20T09:07:08Z",
                                            "name": "kube-proxy",
                                            "namespace": "kube-system",
                                            "resourceVersion": "205",
                                            "selfLink": "/api/v1/namespaces/kube-system/serviceaccounts/kube-proxy",
                                            "uid": "e389960b-cdb0-444d-908f-f97377c9d5f2"
                                        },
                                        "secrets": [
                                            {
                                                "name": "kube-proxy-token-k546c"
                                            }
                                        ]
                                    }
                                ]
                            },
                            "exception": {
                                "guid": "f0c9dd3c-4c54-4485-b6f9-826b5f2a3dd3",
                                "name": "exception_C-0034_kube-system_d365c088322e7c9db2ca33edcbb28e9a",
                                "attributes": {
                                    "namespaceOnly": "true"
                                },
                                "policyType": "postureExceptionPolicy",
                                "creationTime": "2021-10-14T12:13:15.128513",
                                "actions": [
                                    "alertOnly"
                                ],
                                "resources": [
                                    {
                                        "designatorType": "Attributes",
                                        "wlid": "",
                                        "wildwlid": "",
                                        "sid": "",
                                        "attributes": {
                                            "cluster": "minikube",
                                            "namespace": "kube-system"
                                        }
                                    }
                                ],
                                "posturePolicies": [
                                    {
                                        "frameworkName": "NSA",
                                        "controlName": "Automatic mapping of service account",
                                        "ruleName": ""
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "the following service account: namespace-controller in the following namespace: kube-system mounts service account tokens in pods by default",
                            "ruleStatus": "warning",
                            "packagename": "armo_builtins",
                            "alertScore": 9,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "ServiceAccount",
                                        "metadata": {
                                            "creationTimestamp": "2021-06-20T09:07:07Z",
                                            "name": "namespace-controller",
                                            "namespace": "kube-system",
                                            "resourceVersion": "164",
                                            "selfLink": "/api/v1/namespaces/kube-system/serviceaccounts/namespace-controller",
                                            "uid": "014a1dd6-8da8-4882-bc48-a576917028b1"
                                        },
                                        "secrets": [
                                            {
                                                "name": "namespace-controller-token-mpjf4"
                                            }
                                        ]
                                    }
                                ]
                            },
                            "exception": {
                                "guid": "f0c9dd3c-4c54-4485-b6f9-826b5f2a3dd3",
                                "name": "exception_C-0034_kube-system_d365c088322e7c9db2ca33edcbb28e9a",
                                "attributes": {
                                    "namespaceOnly": "true"
                                },
                                "policyType": "postureExceptionPolicy",
                                "creationTime": "2021-10-14T12:13:15.128513",
                                "actions": [
                                    "alertOnly"
                                ],
                                "resources": [
                                    {
                                        "designatorType": "Attributes",
                                        "wlid": "",
                                        "wildwlid": "",
                                        "sid": "",
                                        "attributes": {
                                            "cluster": "minikube",
                                            "namespace": "kube-system"
                                        }
                                    }
                                ],
                                "posturePolicies": [
                                    {
                                        "frameworkName": "NSA",
                                        "controlName": "Automatic mapping of service account",
                                        "ruleName": ""
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "the following service account: node-controller in the following namespace: kube-system mounts service account tokens in pods by default",
                            "ruleStatus": "warning",
                            "packagename": "armo_builtins",
                            "alertScore": 9,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "ServiceAccount",
                                        "metadata": {
                                            "creationTimestamp": "2021-06-20T09:07:07Z",
                                            "name": "node-controller",
                                            "namespace": "kube-system",
                                            "resourceVersion": "180",
                                            "selfLink": "/api/v1/namespaces/kube-system/serviceaccounts/node-controller",
                                            "uid": "afa8f1b4-4994-4904-bc23-96b973f60e08"
                                        },
                                        "secrets": [
                                            {
                                                "name": "node-controller-token-7v92h"
                                            }
                                        ]
                                    }
                                ]
                            },
                            "exception": {
                                "guid": "f0c9dd3c-4c54-4485-b6f9-826b5f2a3dd3",
                                "name": "exception_C-0034_kube-system_d365c088322e7c9db2ca33edcbb28e9a",
                                "attributes": {
                                    "namespaceOnly": "true"
                                },
                                "policyType": "postureExceptionPolicy",
                                "creationTime": "2021-10-14T12:13:15.128513",
                                "actions": [
                                    "alertOnly"
                                ],
                                "resources": [
                                    {
                                        "designatorType": "Attributes",
                                        "wlid": "",
                                        "wildwlid": "",
                                        "sid": "",
                                        "attributes": {
                                            "cluster": "minikube",
                                            "namespace": "kube-system"
                                        }
                                    }
                                ],
                                "posturePolicies": [
                                    {
                                        "frameworkName": "NSA",
                                        "controlName": "Automatic mapping of service account",
                                        "ruleName": ""
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "the following service account: persistent-volume-binder in the following namespace: kube-system mounts service account tokens in pods by default",
                            "ruleStatus": "warning",
                            "packagename": "armo_builtins",
                            "alertScore": 9,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "ServiceAccount",
                                        "metadata": {
                                            "creationTimestamp": "2021-06-20T09:07:07Z",
                                            "name": "persistent-volume-binder",
                                            "namespace": "kube-system",
                                            "resourceVersion": "176",
                                            "selfLink": "/api/v1/namespaces/kube-system/serviceaccounts/persistent-volume-binder",
                                            "uid": "636ddc2b-3399-4d80-aa07-d5a5ad1d0c8e"
                                        },
                                        "secrets": [
                                            {
                                                "name": "persistent-volume-binder-token-mcvcw"
                                            }
                                        ]
                                    }
                                ]
                            },
                            "exception": {
                                "guid": "f0c9dd3c-4c54-4485-b6f9-826b5f2a3dd3",
                                "name": "exception_C-0034_kube-system_d365c088322e7c9db2ca33edcbb28e9a",
                                "attributes": {
                                    "namespaceOnly": "true"
                                },
                                "policyType": "postureExceptionPolicy",
                                "creationTime": "2021-10-14T12:13:15.128513",
                                "actions": [
                                    "alertOnly"
                                ],
                                "resources": [
                                    {
                                        "designatorType": "Attributes",
                                        "wlid": "",
                                        "wildwlid": "",
                                        "sid": "",
                                        "attributes": {
                                            "cluster": "minikube",
                                            "namespace": "kube-system"
                                        }
                                    }
                                ],
                                "posturePolicies": [
                                    {
                                        "frameworkName": "NSA",
                                        "controlName": "Automatic mapping of service account",
                                        "ruleName": ""
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "the following service account: pod-garbage-collector in the following namespace: kube-system mounts service account tokens in pods by default",
                            "ruleStatus": "warning",
                            "packagename": "armo_builtins",
                            "alertScore": 9,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "ServiceAccount",
                                        "metadata": {
                                            "creationTimestamp": "2021-06-20T09:07:22Z",
                                            "name": "pod-garbage-collector",
                                            "namespace": "kube-system",
                                            "resourceVersion": "292",
                                            "selfLink": "/api/v1/namespaces/kube-system/serviceaccounts/pod-garbage-collector",
                                            "uid": "44d8cc24-e8cc-41de-bda9-3af8fbaddc34"
                                        },
                                        "secrets": [
                                            {
                                                "name": "pod-garbage-collector-token-49dkl"
                                            }
                                        ]
                                    }
                                ]
                            },
                            "exception": {
                                "guid": "f0c9dd3c-4c54-4485-b6f9-826b5f2a3dd3",
                                "name": "exception_C-0034_kube-system_d365c088322e7c9db2ca33edcbb28e9a",
                                "attributes": {
                                    "namespaceOnly": "true"
                                },
                                "policyType": "postureExceptionPolicy",
                                "creationTime": "2021-10-14T12:13:15.128513",
                                "actions": [
                                    "alertOnly"
                                ],
                                "resources": [
                                    {
                                        "designatorType": "Attributes",
                                        "wlid": "",
                                        "wildwlid": "",
                                        "sid": "",
                                        "attributes": {
                                            "cluster": "minikube",
                                            "namespace": "kube-system"
                                        }
                                    }
                                ],
                                "posturePolicies": [
                                    {
                                        "frameworkName": "NSA",
                                        "controlName": "Automatic mapping of service account",
                                        "ruleName": ""
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "the following service account: pv-protection-controller in the following namespace: kube-system mounts service account tokens in pods by default",
                            "ruleStatus": "warning",
                            "packagename": "armo_builtins",
                            "alertScore": 9,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "ServiceAccount",
                                        "metadata": {
                                            "creationTimestamp": "2021-06-20T09:07:11Z",
                                            "name": "pv-protection-controller",
                                            "namespace": "kube-system",
                                            "resourceVersion": "268",
                                            "selfLink": "/api/v1/namespaces/kube-system/serviceaccounts/pv-protection-controller",
                                            "uid": "c6d1f18e-b066-4666-a411-95b26e21e3e3"
                                        },
                                        "secrets": [
                                            {
                                                "name": "pv-protection-controller-token-h69ll"
                                            }
                                        ]
                                    }
                                ]
                            },
                            "exception": {
                                "guid": "f0c9dd3c-4c54-4485-b6f9-826b5f2a3dd3",
                                "name": "exception_C-0034_kube-system_d365c088322e7c9db2ca33edcbb28e9a",
                                "attributes": {
                                    "namespaceOnly": "true"
                                },
                                "policyType": "postureExceptionPolicy",
                                "creationTime": "2021-10-14T12:13:15.128513",
                                "actions": [
                                    "alertOnly"
                                ],
                                "resources": [
                                    {
                                        "designatorType": "Attributes",
                                        "wlid": "",
                                        "wildwlid": "",
                                        "sid": "",
                                        "attributes": {
                                            "cluster": "minikube",
                                            "namespace": "kube-system"
                                        }
                                    }
                                ],
                                "posturePolicies": [
                                    {
                                        "frameworkName": "NSA",
                                        "controlName": "Automatic mapping of service account",
                                        "ruleName": ""
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "the following service account: pvc-protection-controller in the following namespace: kube-system mounts service account tokens in pods by default",
                            "ruleStatus": "warning",
                            "packagename": "armo_builtins",
                            "alertScore": 9,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "ServiceAccount",
                                        "metadata": {
                                            "creationTimestamp": "2021-06-20T09:07:07Z",
                                            "name": "pvc-protection-controller",
                                            "namespace": "kube-system",
                                            "resourceVersion": "167",
                                            "selfLink": "/api/v1/namespaces/kube-system/serviceaccounts/pvc-protection-controller",
                                            "uid": "1c4e574f-f57e-4b1f-9a60-4fd5230103bb"
                                        },
                                        "secrets": [
                                            {
                                                "name": "pvc-protection-controller-token-ssj2k"
                                            }
                                        ]
                                    }
                                ]
                            },
                            "exception": {
                                "guid": "f0c9dd3c-4c54-4485-b6f9-826b5f2a3dd3",
                                "name": "exception_C-0034_kube-system_d365c088322e7c9db2ca33edcbb28e9a",
                                "attributes": {
                                    "namespaceOnly": "true"
                                },
                                "policyType": "postureExceptionPolicy",
                                "creationTime": "2021-10-14T12:13:15.128513",
                                "actions": [
                                    "alertOnly"
                                ],
                                "resources": [
                                    {
                                        "designatorType": "Attributes",
                                        "wlid": "",
                                        "wildwlid": "",
                                        "sid": "",
                                        "attributes": {
                                            "cluster": "minikube",
                                            "namespace": "kube-system"
                                        }
                                    }
                                ],
                                "posturePolicies": [
                                    {
                                        "frameworkName": "NSA",
                                        "controlName": "Automatic mapping of service account",
                                        "ruleName": ""
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "the following service account: replicaset-controller in the following namespace: kube-system mounts service account tokens in pods by default",
                            "ruleStatus": "warning",
                            "packagename": "armo_builtins",
                            "alertScore": 9,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "ServiceAccount",
                                        "metadata": {
                                            "creationTimestamp": "2021-06-20T09:07:08Z",
                                            "name": "replicaset-controller",
                                            "namespace": "kube-system",
                                            "resourceVersion": "210",
                                            "selfLink": "/api/v1/namespaces/kube-system/serviceaccounts/replicaset-controller",
                                            "uid": "4fb2ce39-50f2-4afc-a886-e01179131243"
                                        },
                                        "secrets": [
                                            {
                                                "name": "replicaset-controller-token-qkbwl"
                                            }
                                        ]
                                    }
                                ]
                            },
                            "exception": {
                                "guid": "f0c9dd3c-4c54-4485-b6f9-826b5f2a3dd3",
                                "name": "exception_C-0034_kube-system_d365c088322e7c9db2ca33edcbb28e9a",
                                "attributes": {
                                    "namespaceOnly": "true"
                                },
                                "policyType": "postureExceptionPolicy",
                                "creationTime": "2021-10-14T12:13:15.128513",
                                "actions": [
                                    "alertOnly"
                                ],
                                "resources": [
                                    {
                                        "designatorType": "Attributes",
                                        "wlid": "",
                                        "wildwlid": "",
                                        "sid": "",
                                        "attributes": {
                                            "cluster": "minikube",
                                            "namespace": "kube-system"
                                        }
                                    }
                                ],
                                "posturePolicies": [
                                    {
                                        "frameworkName": "NSA",
                                        "controlName": "Automatic mapping of service account",
                                        "ruleName": ""
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "the following service account: replication-controller in the following namespace: kube-system mounts service account tokens in pods by default",
                            "ruleStatus": "warning",
                            "packagename": "armo_builtins",
                            "alertScore": 9,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "ServiceAccount",
                                        "metadata": {
                                            "creationTimestamp": "2021-06-20T09:07:22Z",
                                            "name": "replication-controller",
                                            "namespace": "kube-system",
                                            "resourceVersion": "289",
                                            "selfLink": "/api/v1/namespaces/kube-system/serviceaccounts/replication-controller",
                                            "uid": "564cda5b-3847-4c61-a713-c9fbfed4c761"
                                        },
                                        "secrets": [
                                            {
                                                "name": "replication-controller-token-v5mk6"
                                            }
                                        ]
                                    }
                                ]
                            },
                            "exception": {
                                "guid": "f0c9dd3c-4c54-4485-b6f9-826b5f2a3dd3",
                                "name": "exception_C-0034_kube-system_d365c088322e7c9db2ca33edcbb28e9a",
                                "attributes": {
                                    "namespaceOnly": "true"
                                },
                                "policyType": "postureExceptionPolicy",
                                "creationTime": "2021-10-14T12:13:15.128513",
                                "actions": [
                                    "alertOnly"
                                ],
                                "resources": [
                                    {
                                        "designatorType": "Attributes",
                                        "wlid": "",
                                        "wildwlid": "",
                                        "sid": "",
                                        "attributes": {
                                            "cluster": "minikube",
                                            "namespace": "kube-system"
                                        }
                                    }
                                ],
                                "posturePolicies": [
                                    {
                                        "frameworkName": "NSA",
                                        "controlName": "Automatic mapping of service account",
                                        "ruleName": ""
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "the following service account: resourcequota-controller in the following namespace: kube-system mounts service account tokens in pods by default",
                            "ruleStatus": "warning",
                            "packagename": "armo_builtins",
                            "alertScore": 9,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "ServiceAccount",
                                        "metadata": {
                                            "creationTimestamp": "2021-06-20T09:07:08Z",
                                            "name": "resourcequota-controller",
                                            "namespace": "kube-system",
                                            "resourceVersion": "198",
                                            "selfLink": "/api/v1/namespaces/kube-system/serviceaccounts/resourcequota-controller",
                                            "uid": "5e88b92f-a1ce-4f67-ab53-48c65dd0eabc"
                                        },
                                        "secrets": [
                                            {
                                                "name": "resourcequota-controller-token-wfmpg"
                                            }
                                        ]
                                    }
                                ]
                            },
                            "exception": {
                                "guid": "f0c9dd3c-4c54-4485-b6f9-826b5f2a3dd3",
                                "name": "exception_C-0034_kube-system_d365c088322e7c9db2ca33edcbb28e9a",
                                "attributes": {
                                    "namespaceOnly": "true"
                                },
                                "policyType": "postureExceptionPolicy",
                                "creationTime": "2021-10-14T12:13:15.128513",
                                "actions": [
                                    "alertOnly"
                                ],
                                "resources": [
                                    {
                                        "designatorType": "Attributes",
                                        "wlid": "",
                                        "wildwlid": "",
                                        "sid": "",
                                        "attributes": {
                                            "cluster": "minikube",
                                            "namespace": "kube-system"
                                        }
                                    }
                                ],
                                "posturePolicies": [
                                    {
                                        "frameworkName": "NSA",
                                        "controlName": "Automatic mapping of service account",
                                        "ruleName": ""
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "the following service account: service-account-controller in the following namespace: kube-system mounts service account tokens in pods by default",
                            "ruleStatus": "warning",
                            "packagename": "armo_builtins",
                            "alertScore": 9,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "ServiceAccount",
                                        "metadata": {
                                            "creationTimestamp": "2021-06-20T09:07:22Z",
                                            "name": "service-account-controller",
                                            "namespace": "kube-system",
                                            "resourceVersion": "295",
                                            "selfLink": "/api/v1/namespaces/kube-system/serviceaccounts/service-account-controller",
                                            "uid": "4b034d22-4b95-46f1-a0e4-8360ab0d9b5e"
                                        },
                                        "secrets": [
                                            {
                                                "name": "service-account-controller-token-8htl5"
                                            }
                                        ]
                                    }
                                ]
                            },
                            "exception": {
                                "guid": "f0c9dd3c-4c54-4485-b6f9-826b5f2a3dd3",
                                "name": "exception_C-0034_kube-system_d365c088322e7c9db2ca33edcbb28e9a",
                                "attributes": {
                                    "namespaceOnly": "true"
                                },
                                "policyType": "postureExceptionPolicy",
                                "creationTime": "2021-10-14T12:13:15.128513",
                                "actions": [
                                    "alertOnly"
                                ],
                                "resources": [
                                    {
                                        "designatorType": "Attributes",
                                        "wlid": "",
                                        "wildwlid": "",
                                        "sid": "",
                                        "attributes": {
                                            "cluster": "minikube",
                                            "namespace": "kube-system"
                                        }
                                    }
                                ],
                                "posturePolicies": [
                                    {
                                        "frameworkName": "NSA",
                                        "controlName": "Automatic mapping of service account",
                                        "ruleName": ""
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "the following service account: service-controller in the following namespace: kube-system mounts service account tokens in pods by default",
                            "ruleStatus": "warning",
                            "packagename": "armo_builtins",
                            "alertScore": 9,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "ServiceAccount",
                                        "metadata": {
                                            "creationTimestamp": "2021-06-20T09:07:09Z",
                                            "name": "service-controller",
                                            "namespace": "kube-system",
                                            "resourceVersion": "230",
                                            "selfLink": "/api/v1/namespaces/kube-system/serviceaccounts/service-controller",
                                            "uid": "d122fb96-85f7-4d4e-9181-4e0dbec88905"
                                        },
                                        "secrets": [
                                            {
                                                "name": "service-controller-token-nq68f"
                                            }
                                        ]
                                    }
                                ]
                            },
                            "exception": {
                                "guid": "f0c9dd3c-4c54-4485-b6f9-826b5f2a3dd3",
                                "name": "exception_C-0034_kube-system_d365c088322e7c9db2ca33edcbb28e9a",
                                "attributes": {
                                    "namespaceOnly": "true"
                                },
                                "policyType": "postureExceptionPolicy",
                                "creationTime": "2021-10-14T12:13:15.128513",
                                "actions": [
                                    "alertOnly"
                                ],
                                "resources": [
                                    {
                                        "designatorType": "Attributes",
                                        "wlid": "",
                                        "wildwlid": "",
                                        "sid": "",
                                        "attributes": {
                                            "cluster": "minikube",
                                            "namespace": "kube-system"
                                        }
                                    }
                                ],
                                "posturePolicies": [
                                    {
                                        "frameworkName": "NSA",
                                        "controlName": "Automatic mapping of service account",
                                        "ruleName": ""
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "the following service account: statefulset-controller in the following namespace: kube-system mounts service account tokens in pods by default",
                            "ruleStatus": "warning",
                            "packagename": "armo_builtins",
                            "alertScore": 9,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "ServiceAccount",
                                        "metadata": {
                                            "creationTimestamp": "2021-06-20T09:07:12Z",
                                            "name": "statefulset-controller",
                                            "namespace": "kube-system",
                                            "resourceVersion": "281",
                                            "selfLink": "/api/v1/namespaces/kube-system/serviceaccounts/statefulset-controller",
                                            "uid": "024e1ca4-839f-48a9-a457-3db304566a5e"
                                        },
                                        "secrets": [
                                            {
                                                "name": "statefulset-controller-token-vjtz6"
                                            }
                                        ]
                                    }
                                ]
                            },
                            "exception": {
                                "guid": "f0c9dd3c-4c54-4485-b6f9-826b5f2a3dd3",
                                "name": "exception_C-0034_kube-system_d365c088322e7c9db2ca33edcbb28e9a",
                                "attributes": {
                                    "namespaceOnly": "true"
                                },
                                "policyType": "postureExceptionPolicy",
                                "creationTime": "2021-10-14T12:13:15.128513",
                                "actions": [
                                    "alertOnly"
                                ],
                                "resources": [
                                    {
                                        "designatorType": "Attributes",
                                        "wlid": "",
                                        "wildwlid": "",
                                        "sid": "",
                                        "attributes": {
                                            "cluster": "minikube",
                                            "namespace": "kube-system"
                                        }
                                    }
                                ],
                                "posturePolicies": [
                                    {
                                        "frameworkName": "NSA",
                                        "controlName": "Automatic mapping of service account",
                                        "ruleName": ""
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "the following service account: storage-provisioner in the following namespace: kube-system mounts service account tokens in pods by default",
                            "ruleStatus": "warning",
                            "packagename": "armo_builtins",
                            "alertScore": 9,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "ServiceAccount",
                                        "metadata": {
                                            "annotations": {
                                                "kubectl.kubernetes.io/last-applied-configuration": "{\"apiVersion\":\"v1\",\"kind\":\"ServiceAccount\",\"metadata\":{\"annotations\":{},\"labels\":{\"addonmanager.kubernetes.io/mode\":\"Reconcile\"},\"name\":\"storage-provisioner\",\"namespace\":\"kube-system\"}}\n"
                                            },
                                            "creationTimestamp": "2021-06-20T09:07:09Z",
                                            "labels": {
                                                "addonmanager.kubernetes.io/mode": "Reconcile"
                                            },
                                            "name": "storage-provisioner",
                                            "namespace": "kube-system",
                                            "resourceVersion": "235",
                                            "selfLink": "/api/v1/namespaces/kube-system/serviceaccounts/storage-provisioner",
                                            "uid": "32e0cb77-df25-4ff3-9eb2-512c3f012e26"
                                        },
                                        "secrets": [
                                            {
                                                "name": "storage-provisioner-token-bbjlq"
                                            }
                                        ]
                                    }
                                ]
                            },
                            "exception": {
                                "guid": "f0c9dd3c-4c54-4485-b6f9-826b5f2a3dd3",
                                "name": "exception_C-0034_kube-system_d365c088322e7c9db2ca33edcbb28e9a",
                                "attributes": {
                                    "namespaceOnly": "true"
                                },
                                "policyType": "postureExceptionPolicy",
                                "creationTime": "2021-10-14T12:13:15.128513",
                                "actions": [
                                    "alertOnly"
                                ],
                                "resources": [
                                    {
                                        "designatorType": "Attributes",
                                        "wlid": "",
                                        "wildwlid": "",
                                        "sid": "",
                                        "attributes": {
                                            "cluster": "minikube",
                                            "namespace": "kube-system"
                                        }
                                    }
                                ],
                                "posturePolicies": [
                                    {
                                        "frameworkName": "NSA",
                                        "controlName": "Automatic mapping of service account",
                                        "ruleName": ""
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "the following service account: token-cleaner in the following namespace: kube-system mounts service account tokens in pods by default",
                            "ruleStatus": "warning",
                            "packagename": "armo_builtins",
                            "alertScore": 9,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "ServiceAccount",
                                        "metadata": {
                                            "creationTimestamp": "2021-06-20T09:07:07Z",
                                            "name": "token-cleaner",
                                            "namespace": "kube-system",
                                            "resourceVersion": "173",
                                            "selfLink": "/api/v1/namespaces/kube-system/serviceaccounts/token-cleaner",
                                            "uid": "ee7167ef-5c2f-4446-a883-993ab465aada"
                                        },
                                        "secrets": [
                                            {
                                                "name": "token-cleaner-token-kpwv6"
                                            }
                                        ]
                                    }
                                ]
                            },
                            "exception": {
                                "guid": "f0c9dd3c-4c54-4485-b6f9-826b5f2a3dd3",
                                "name": "exception_C-0034_kube-system_d365c088322e7c9db2ca33edcbb28e9a",
                                "attributes": {
                                    "namespaceOnly": "true"
                                },
                                "policyType": "postureExceptionPolicy",
                                "creationTime": "2021-10-14T12:13:15.128513",
                                "actions": [
                                    "alertOnly"
                                ],
                                "resources": [
                                    {
                                        "designatorType": "Attributes",
                                        "wlid": "",
                                        "wildwlid": "",
                                        "sid": "",
                                        "attributes": {
                                            "cluster": "minikube",
                                            "namespace": "kube-system"
                                        }
                                    }
                                ],
                                "posturePolicies": [
                                    {
                                        "frameworkName": "NSA",
                                        "controlName": "Automatic mapping of service account",
                                        "ruleName": ""
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "the following service account: ttl-controller in the following namespace: kube-system mounts service account tokens in pods by default",
                            "ruleStatus": "warning",
                            "packagename": "armo_builtins",
                            "alertScore": 9,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "ServiceAccount",
                                        "metadata": {
                                            "creationTimestamp": "2021-06-20T09:07:09Z",
                                            "name": "ttl-controller",
                                            "namespace": "kube-system",
                                            "resourceVersion": "254",
                                            "selfLink": "/api/v1/namespaces/kube-system/serviceaccounts/ttl-controller",
                                            "uid": "6b691af3-fdbd-4bf0-9b3c-2069f0d99f4f"
                                        },
                                        "secrets": [
                                            {
                                                "name": "ttl-controller-token-mtgjj"
                                            }
                                        ]
                                    }
                                ]
                            },
                            "exception": {
                                "guid": "f0c9dd3c-4c54-4485-b6f9-826b5f2a3dd3",
                                "name": "exception_C-0034_kube-system_d365c088322e7c9db2ca33edcbb28e9a",
                                "attributes": {
                                    "namespaceOnly": "true"
                                },
                                "policyType": "postureExceptionPolicy",
                                "creationTime": "2021-10-14T12:13:15.128513",
                                "actions": [
                                    "alertOnly"
                                ],
                                "resources": [
                                    {
                                        "designatorType": "Attributes",
                                        "wlid": "",
                                        "wildwlid": "",
                                        "sid": "",
                                        "attributes": {
                                            "cluster": "minikube",
                                            "namespace": "kube-system"
                                        }
                                    }
                                ],
                                "posturePolicies": [
                                    {
                                        "frameworkName": "NSA",
                                        "controlName": "Automatic mapping of service account",
                                        "ruleName": ""
                                    }
                                ]
                            }
                        }
                    ]
                }
            ],
            "remediation": "Disable automatic mounting of service account tokens to PODs either at the service account level or at the individual POD level, by specifying the automountServiceAccountToken: false. Note that POD level takes precedence.",
            "description": "Potential attacker may gain access to a POD and steal its service account token. Therefore, it is recommended to disable automatic mapping of the service account tokens in service account configuration and enable it only for PODs that need to use them.",
            "score": 86
        },
        {
            "guid": "",
            "attributes": {
                "armoBuiltin": true
            },
            "id": "C-0041",
            "controlID": "C-0041",
            "name": "hostNetwork access",
            "ruleReports": [
                {
                    "name": "host-network-access",
                    "remediation": "",
                    "ruleStatus": {
                        "status": "success",
                        "message": ""
                    },
                    "ruleResponses": [
                        {
                            "alertMessage": "Pod: etcd-david-virtualbox is connected to the host network",
                            "ruleStatus": "",
                            "packagename": "armo_builtins",
                            "alertScore": 9,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "Pod",
                                        "metadata": {
                                            "annotations": {
                                                "kubernetes.io/config.hash": "e0fcc6e4323055b5880f8aac4c950836",
                                                "kubernetes.io/config.mirror": "e0fcc6e4323055b5880f8aac4c950836",
                                                "kubernetes.io/config.seen": "2021-06-20T12:06:52.495386281+03:00",
                                                "kubernetes.io/config.source": "file"
                                            },
                                            "creationTimestamp": "2021-06-20T09:08:22Z",
                                            "labels": {
                                                "component": "etcd",
                                                "tier": "control-plane"
                                            },
                                            "name": "etcd-david-virtualbox",
                                            "namespace": "kube-system",
                                            "resourceVersion": "2110909",
                                            "selfLink": "/api/v1/namespaces/kube-system/pods/etcd-david-virtualbox",
                                            "uid": "154e7f87-907f-4edb-a73c-26e965d4fe02"
                                        },
                                        "spec": {
                                            "containers": [
                                                {
                                                    "command": [
                                                        "etcd",
                                                        "--advertise-client-urls=https://10.0.2.15:2379",
                                                        "--cert-file=/var/lib/minikube/certs/etcd/server.crt",
                                                        "--client-cert-auth=true",
                                                        "--data-dir=/var/lib/minikube/etcd",
                                                        "--initial-advertise-peer-urls=https://10.0.2.15:2380",
                                                        "--initial-cluster=david-virtualbox=https://10.0.2.15:2380",
                                                        "--key-file=/var/lib/minikube/certs/etcd/server.key",
                                                        "--listen-client-urls=https://127.0.0.1:2379,https://10.0.2.15:2379",
                                                        "--listen-metrics-urls=http://127.0.0.1:2381,http://10.0.2.15:2381",
                                                        "--listen-peer-urls=https://10.0.2.15:2380",
                                                        "--name=david-virtualbox",
                                                        "--peer-cert-file=/var/lib/minikube/certs/etcd/peer.crt",
                                                        "--peer-client-cert-auth=true",
                                                        "--peer-key-file=/var/lib/minikube/certs/etcd/peer.key",
                                                        "--peer-trusted-ca-file=/var/lib/minikube/certs/etcd/ca.crt",
                                                        "--snapshot-count=10000",
                                                        "--trusted-ca-file=/var/lib/minikube/certs/etcd/ca.crt"
                                                    ],
                                                    "image": "k8s.gcr.io/etcd:3.3.15-0",
                                                    "imagePullPolicy": "IfNotPresent",
                                                    "livenessProbe": {
                                                        "failureThreshold": 8,
                                                        "httpGet": {
                                                            "host": "127.0.0.1",
                                                            "path": "/health",
                                                            "port": 2381,
                                                            "scheme": "HTTP"
                                                        },
                                                        "initialDelaySeconds": 15,
                                                        "periodSeconds": 10,
                                                        "successThreshold": 1,
                                                        "timeoutSeconds": 15
                                                    },
                                                    "name": "etcd",
                                                    "resources": {},
                                                    "terminationMessagePath": "/dev/termination-log",
                                                    "terminationMessagePolicy": "File",
                                                    "volumeMounts": [
                                                        {
                                                            "mountPath": "/var/lib/minikube/etcd",
                                                            "name": "etcd-data"
                                                        },
                                                        {
                                                            "mountPath": "/var/lib/minikube/certs/etcd",
                                                            "name": "etcd-certs"
                                                        }
                                                    ]
                                                }
                                            ],
                                            "dnsPolicy": "ClusterFirst",
                                            "enableServiceLinks": true,
                                            "hostNetwork": true,
                                            "nodeName": "david-virtualbox",
                                            "priority": 2000000000,
                                            "priorityClassName": "system-cluster-critical",
                                            "restartPolicy": "Always",
                                            "schedulerName": "default-scheduler",
                                            "securityContext": {},
                                            "terminationGracePeriodSeconds": 30,
                                            "tolerations": [
                                                {
                                                    "effect": "NoExecute",
                                                    "operator": "Exists"
                                                }
                                            ],
                                            "volumes": [
                                                {
                                                    "hostPath": {
                                                        "path": "/var/lib/minikube/certs/etcd",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "etcd-certs"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/var/lib/minikube/etcd",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "etcd-data"
                                                }
                                            ]
                                        },
                                        "status": {
                                            "conditions": [
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-03T06:28:57Z",
                                                    "status": "True",
                                                    "type": "Initialized"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-06T05:23:22Z",
                                                    "status": "True",
                                                    "type": "Ready"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-06T05:23:22Z",
                                                    "status": "True",
                                                    "type": "ContainersReady"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-03T06:28:57Z",
                                                    "status": "True",
                                                    "type": "PodScheduled"
                                                }
                                            ],
                                            "containerStatuses": [
                                                {
                                                    "containerID": "docker://0cc62102444cc74e3dddb2b6cd7550036feaa0e61f4297c85fc72e3be2905ec9",
                                                    "image": "k8s.gcr.io/etcd:3.3.15-0",
                                                    "imageID": "docker-pullable://k8s.gcr.io/etcd@sha256:12c2c5e5731c3bcd56e6f1c05c0f9198b6f06793fa7fca2fb43aab9622dc4afa",
                                                    "lastState": {
                                                        "terminated": {
                                                            "containerID": "docker://61de0abaa35617dafed9274c5b03f99f269412bdea2f15d903b4293d620f4b9e",
                                                            "exitCode": 255,
                                                            "finishedAt": "2021-10-14T05:35:02Z",
                                                            "reason": "Error",
                                                            "startedAt": "2021-10-12T10:28:38Z"
                                                        }
                                                    },
                                                    "name": "etcd",
                                                    "ready": true,
                                                    "restartCount": 51,
                                                    "started": true,
                                                    "state": {
                                                        "running": {
                                                            "startedAt": "2021-10-14T05:35:12Z"
                                                        }
                                                    }
                                                }
                                            ],
                                            "hostIP": "10.0.2.15",
                                            "phase": "Running",
                                            "podIP": "10.0.2.15",
                                            "podIPs": [
                                                {
                                                    "ip": "10.0.2.15"
                                                }
                                            ],
                                            "qosClass": "BestEffort",
                                            "startTime": "2021-10-03T06:28:57Z"
                                        }
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "Pod: kube-apiserver-david-virtualbox is connected to the host network",
                            "ruleStatus": "",
                            "packagename": "armo_builtins",
                            "alertScore": 9,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "Pod",
                                        "metadata": {
                                            "annotations": {
                                                "kubernetes.io/config.hash": "668e9396bc7b8c495d39f4a3bc479397",
                                                "kubernetes.io/config.mirror": "668e9396bc7b8c495d39f4a3bc479397",
                                                "kubernetes.io/config.seen": "2021-06-20T12:06:52.495387809+03:00",
                                                "kubernetes.io/config.source": "file"
                                            },
                                            "creationTimestamp": "2021-06-20T09:08:05Z",
                                            "labels": {
                                                "component": "kube-apiserver",
                                                "tier": "control-plane"
                                            },
                                            "name": "kube-apiserver-david-virtualbox",
                                            "namespace": "kube-system",
                                            "resourceVersion": "2110920",
                                            "selfLink": "/api/v1/namespaces/kube-system/pods/kube-apiserver-david-virtualbox",
                                            "uid": "327cbf13-97d6-42a3-8469-5acfdbe1be09"
                                        },
                                        "spec": {
                                            "containers": [
                                                {
                                                    "command": [
                                                        "kube-apiserver",
                                                        "--advertise-address=10.0.2.15",
                                                        "--allow-privileged=true",
                                                        "--authorization-mode=Node,RBAC",
                                                        "--client-ca-file=/var/lib/minikube/certs/ca.crt",
                                                        "--enable-admission-plugins=NamespaceLifecycle,LimitRanger,ServiceAccount,DefaultStorageClass,DefaultTolerationSeconds,NodeRestriction,MutatingAdmissionWebhook,ValidatingAdmissionWebhook,ResourceQuota",
                                                        "--enable-bootstrap-token-auth=true",
                                                        "--etcd-cafile=/var/lib/minikube/certs/etcd/ca.crt",
                                                        "--etcd-certfile=/var/lib/minikube/certs/apiserver-etcd-client.crt",
                                                        "--etcd-keyfile=/var/lib/minikube/certs/apiserver-etcd-client.key",
                                                        "--etcd-servers=https://127.0.0.1:2379",
                                                        "--insecure-port=0",
                                                        "--kubelet-client-certificate=/var/lib/minikube/certs/apiserver-kubelet-client.crt",
                                                        "--kubelet-client-key=/var/lib/minikube/certs/apiserver-kubelet-client.key",
                                                        "--kubelet-preferred-address-types=InternalIP,ExternalIP,Hostname",
                                                        "--proxy-client-cert-file=/var/lib/minikube/certs/front-proxy-client.crt",
                                                        "--proxy-client-key-file=/var/lib/minikube/certs/front-proxy-client.key",
                                                        "--requestheader-allowed-names=front-proxy-client",
                                                        "--requestheader-client-ca-file=/var/lib/minikube/certs/front-proxy-ca.crt",
                                                        "--requestheader-extra-headers-prefix=X-Remote-Extra-",
                                                        "--requestheader-group-headers=X-Remote-Group",
                                                        "--requestheader-username-headers=X-Remote-User",
                                                        "--secure-port=8443",
                                                        "--service-account-key-file=/var/lib/minikube/certs/sa.pub",
                                                        "--service-cluster-ip-range=10.96.0.0/12",
                                                        "--tls-cert-file=/var/lib/minikube/certs/apiserver.crt",
                                                        "--tls-private-key-file=/var/lib/minikube/certs/apiserver.key"
                                                    ],
                                                    "image": "k8s.gcr.io/kube-apiserver:v1.16.0",
                                                    "imagePullPolicy": "IfNotPresent",
                                                    "livenessProbe": {
                                                        "failureThreshold": 8,
                                                        "httpGet": {
                                                            "host": "10.0.2.15",
                                                            "path": "/healthz",
                                                            "port": 8443,
                                                            "scheme": "HTTPS"
                                                        },
                                                        "initialDelaySeconds": 15,
                                                        "periodSeconds": 10,
                                                        "successThreshold": 1,
                                                        "timeoutSeconds": 15
                                                    },
                                                    "name": "kube-apiserver",
                                                    "resources": {
                                                        "requests": {
                                                            "cpu": "250m"
                                                        }
                                                    },
                                                    "terminationMessagePath": "/dev/termination-log",
                                                    "terminationMessagePolicy": "File",
                                                    "volumeMounts": [
                                                        {
                                                            "mountPath": "/etc/ssl/certs",
                                                            "name": "ca-certs",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/etc/ca-certificates",
                                                            "name": "etc-ca-certificates",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/etc/pki",
                                                            "name": "etc-pki",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/var/lib/minikube/certs",
                                                            "name": "k8s-certs",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/usr/local/share/ca-certificates",
                                                            "name": "usr-local-share-ca-certificates",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/usr/share/ca-certificates",
                                                            "name": "usr-share-ca-certificates",
                                                            "readOnly": true
                                                        }
                                                    ]
                                                }
                                            ],
                                            "dnsPolicy": "ClusterFirst",
                                            "enableServiceLinks": true,
                                            "hostNetwork": true,
                                            "nodeName": "david-virtualbox",
                                            "priority": 2000000000,
                                            "priorityClassName": "system-cluster-critical",
                                            "restartPolicy": "Always",
                                            "schedulerName": "default-scheduler",
                                            "securityContext": {},
                                            "terminationGracePeriodSeconds": 30,
                                            "tolerations": [
                                                {
                                                    "effect": "NoExecute",
                                                    "operator": "Exists"
                                                }
                                            ],
                                            "volumes": [
                                                {
                                                    "hostPath": {
                                                        "path": "/etc/ssl/certs",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "ca-certs"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/etc/ca-certificates",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "etc-ca-certificates"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/etc/pki",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "etc-pki"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/var/lib/minikube/certs",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "k8s-certs"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/usr/local/share/ca-certificates",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "usr-local-share-ca-certificates"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/usr/share/ca-certificates",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "usr-share-ca-certificates"
                                                }
                                            ]
                                        },
                                        "status": {
                                            "conditions": [
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-03T06:28:57Z",
                                                    "status": "True",
                                                    "type": "Initialized"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-06T05:23:25Z",
                                                    "status": "True",
                                                    "type": "Ready"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-06T05:23:25Z",
                                                    "status": "True",
                                                    "type": "ContainersReady"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-03T06:28:57Z",
                                                    "status": "True",
                                                    "type": "PodScheduled"
                                                }
                                            ],
                                            "containerStatuses": [
                                                {
                                                    "containerID": "docker://c6d5933ac2752c18942e488d2d162fcc3141d1ae131be2abc3a50d2cb91e4e22",
                                                    "image": "k8s.gcr.io/kube-apiserver:v1.16.0",
                                                    "imageID": "docker-pullable://k8s.gcr.io/kube-apiserver@sha256:f4168527c91289da2708f62ae729fdde5fb484167dd05ffbb7ab666f60de96cd",
                                                    "lastState": {
                                                        "terminated": {
                                                            "containerID": "docker://3714c8295fe8b0ff97082305a44f08764e06d1a599ab3ee3ca6db15ab498a276",
                                                            "exitCode": 255,
                                                            "finishedAt": "2021-10-14T05:35:02Z",
                                                            "reason": "Error",
                                                            "startedAt": "2021-10-12T10:28:38Z"
                                                        }
                                                    },
                                                    "name": "kube-apiserver",
                                                    "ready": true,
                                                    "restartCount": 51,
                                                    "started": true,
                                                    "state": {
                                                        "running": {
                                                            "startedAt": "2021-10-14T05:35:12Z"
                                                        }
                                                    }
                                                }
                                            ],
                                            "hostIP": "10.0.2.15",
                                            "phase": "Running",
                                            "podIP": "10.0.2.15",
                                            "podIPs": [
                                                {
                                                    "ip": "10.0.2.15"
                                                }
                                            ],
                                            "qosClass": "Burstable",
                                            "startTime": "2021-10-03T06:28:57Z"
                                        }
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "Pod: kube-controller-manager-david-virtualbox is connected to the host network",
                            "ruleStatus": "",
                            "packagename": "armo_builtins",
                            "alertScore": 9,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "Pod",
                                        "metadata": {
                                            "annotations": {
                                                "kubernetes.io/config.hash": "a16b2d5766eae37796e4a8ed7f8ce12a",
                                                "kubernetes.io/config.mirror": "a16b2d5766eae37796e4a8ed7f8ce12a",
                                                "kubernetes.io/config.seen": "2021-06-20T12:06:52.495389283+03:00",
                                                "kubernetes.io/config.source": "file"
                                            },
                                            "creationTimestamp": "2021-06-20T09:08:00Z",
                                            "labels": {
                                                "component": "kube-controller-manager",
                                                "tier": "control-plane"
                                            },
                                            "name": "kube-controller-manager-david-virtualbox",
                                            "namespace": "kube-system",
                                            "resourceVersion": "2110899",
                                            "selfLink": "/api/v1/namespaces/kube-system/pods/kube-controller-manager-david-virtualbox",
                                            "uid": "6ca9d32c-21c3-4c0e-8087-5445c80a2bcc"
                                        },
                                        "spec": {
                                            "containers": [
                                                {
                                                    "command": [
                                                        "kube-controller-manager",
                                                        "--allocate-node-cidrs=true",
                                                        "--authentication-kubeconfig=/etc/kubernetes/controller-manager.conf",
                                                        "--authorization-kubeconfig=/etc/kubernetes/controller-manager.conf",
                                                        "--bind-address=127.0.0.1",
                                                        "--client-ca-file=/var/lib/minikube/certs/ca.crt",
                                                        "--cluster-cidr=10.244.0.0/16",
                                                        "--cluster-signing-cert-file=/var/lib/minikube/certs/ca.crt",
                                                        "--cluster-signing-key-file=/var/lib/minikube/certs/ca.key",
                                                        "--controllers=*,bootstrapsigner,tokencleaner",
                                                        "--kubeconfig=/etc/kubernetes/controller-manager.conf",
                                                        "--leader-elect=false",
                                                        "--node-cidr-mask-size=24",
                                                        "--requestheader-client-ca-file=/var/lib/minikube/certs/front-proxy-ca.crt",
                                                        "--root-ca-file=/var/lib/minikube/certs/ca.crt",
                                                        "--service-account-private-key-file=/var/lib/minikube/certs/sa.key",
                                                        "--service-cluster-ip-range=10.96.0.0/12",
                                                        "--use-service-account-credentials=true"
                                                    ],
                                                    "image": "k8s.gcr.io/kube-controller-manager:v1.16.0",
                                                    "imagePullPolicy": "IfNotPresent",
                                                    "livenessProbe": {
                                                        "failureThreshold": 8,
                                                        "httpGet": {
                                                            "host": "127.0.0.1",
                                                            "path": "/healthz",
                                                            "port": 10252,
                                                            "scheme": "HTTP"
                                                        },
                                                        "initialDelaySeconds": 15,
                                                        "periodSeconds": 10,
                                                        "successThreshold": 1,
                                                        "timeoutSeconds": 15
                                                    },
                                                    "name": "kube-controller-manager",
                                                    "resources": {
                                                        "requests": {
                                                            "cpu": "200m"
                                                        }
                                                    },
                                                    "terminationMessagePath": "/dev/termination-log",
                                                    "terminationMessagePolicy": "File",
                                                    "volumeMounts": [
                                                        {
                                                            "mountPath": "/etc/ssl/certs",
                                                            "name": "ca-certs",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/etc/ca-certificates",
                                                            "name": "etc-ca-certificates",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/etc/pki",
                                                            "name": "etc-pki",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/usr/libexec/kubernetes/kubelet-plugins/volume/exec",
                                                            "name": "flexvolume-dir"
                                                        },
                                                        {
                                                            "mountPath": "/var/lib/minikube/certs",
                                                            "name": "k8s-certs",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/etc/kubernetes/controller-manager.conf",
                                                            "name": "kubeconfig",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/usr/local/share/ca-certificates",
                                                            "name": "usr-local-share-ca-certificates",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/usr/share/ca-certificates",
                                                            "name": "usr-share-ca-certificates",
                                                            "readOnly": true
                                                        }
                                                    ]
                                                }
                                            ],
                                            "dnsPolicy": "ClusterFirst",
                                            "enableServiceLinks": true,
                                            "hostNetwork": true,
                                            "nodeName": "david-virtualbox",
                                            "priority": 2000000000,
                                            "priorityClassName": "system-cluster-critical",
                                            "restartPolicy": "Always",
                                            "schedulerName": "default-scheduler",
                                            "securityContext": {},
                                            "terminationGracePeriodSeconds": 30,
                                            "tolerations": [
                                                {
                                                    "effect": "NoExecute",
                                                    "operator": "Exists"
                                                }
                                            ],
                                            "volumes": [
                                                {
                                                    "hostPath": {
                                                        "path": "/etc/ssl/certs",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "ca-certs"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/etc/ca-certificates",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "etc-ca-certificates"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/etc/pki",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "etc-pki"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/usr/libexec/kubernetes/kubelet-plugins/volume/exec",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "flexvolume-dir"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/var/lib/minikube/certs",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "k8s-certs"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/etc/kubernetes/controller-manager.conf",
                                                        "type": "FileOrCreate"
                                                    },
                                                    "name": "kubeconfig"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/usr/local/share/ca-certificates",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "usr-local-share-ca-certificates"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/usr/share/ca-certificates",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "usr-share-ca-certificates"
                                                }
                                            ]
                                        },
                                        "status": {
                                            "conditions": [
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-14T05:35:11Z",
                                                    "status": "True",
                                                    "type": "Initialized"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-14T05:35:13Z",
                                                    "status": "True",
                                                    "type": "Ready"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-14T05:35:13Z",
                                                    "status": "True",
                                                    "type": "ContainersReady"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-14T05:35:11Z",
                                                    "status": "True",
                                                    "type": "PodScheduled"
                                                }
                                            ],
                                            "containerStatuses": [
                                                {
                                                    "containerID": "docker://eefb83fb1b81a497703013f609a54ffc03b57157395cb5f9a128dfeffea54b58",
                                                    "image": "k8s.gcr.io/kube-controller-manager:v1.16.0",
                                                    "imageID": "docker-pullable://k8s.gcr.io/kube-controller-manager@sha256:c156a05ee9d40e3ca2ebf9337f38a10558c1fc6c9124006f128a82e6c38cdf3e",
                                                    "lastState": {
                                                        "terminated": {
                                                            "containerID": "docker://81d8dca7f5b87e8b9bd8f515e9bcf86f4614306a9c729c30aad92560f520ed54",
                                                            "exitCode": 255,
                                                            "finishedAt": "2021-10-14T05:35:02Z",
                                                            "reason": "Error",
                                                            "startedAt": "2021-10-12T10:28:38Z"
                                                        }
                                                    },
                                                    "name": "kube-controller-manager",
                                                    "ready": true,
                                                    "restartCount": 55,
                                                    "started": true,
                                                    "state": {
                                                        "running": {
                                                            "startedAt": "2021-10-14T05:35:12Z"
                                                        }
                                                    }
                                                }
                                            ],
                                            "hostIP": "10.0.2.15",
                                            "phase": "Running",
                                            "podIP": "10.0.2.15",
                                            "podIPs": [
                                                {
                                                    "ip": "10.0.2.15"
                                                }
                                            ],
                                            "qosClass": "Burstable",
                                            "startTime": "2021-10-14T05:35:11Z"
                                        }
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "Pod: kube-scheduler-david-virtualbox is connected to the host network",
                            "ruleStatus": "",
                            "packagename": "armo_builtins",
                            "alertScore": 9,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "Pod",
                                        "metadata": {
                                            "annotations": {
                                                "kubernetes.io/config.hash": "b3d303074fe0ca1d42a8bd9ed248df09",
                                                "kubernetes.io/config.mirror": "b3d303074fe0ca1d42a8bd9ed248df09",
                                                "kubernetes.io/config.seen": "2021-06-20T12:06:52.495381685+03:00",
                                                "kubernetes.io/config.source": "file"
                                            },
                                            "creationTimestamp": "2021-06-20T09:08:11Z",
                                            "labels": {
                                                "component": "kube-scheduler",
                                                "tier": "control-plane"
                                            },
                                            "name": "kube-scheduler-david-virtualbox",
                                            "namespace": "kube-system",
                                            "resourceVersion": "2110888",
                                            "selfLink": "/api/v1/namespaces/kube-system/pods/kube-scheduler-david-virtualbox",
                                            "uid": "226e285e-b2b3-423c-8bd4-04cfb775cbc6"
                                        },
                                        "spec": {
                                            "containers": [
                                                {
                                                    "command": [
                                                        "kube-scheduler",
                                                        "--authentication-kubeconfig=/etc/kubernetes/scheduler.conf",
                                                        "--authorization-kubeconfig=/etc/kubernetes/scheduler.conf",
                                                        "--bind-address=127.0.0.1",
                                                        "--kubeconfig=/etc/kubernetes/scheduler.conf",
                                                        "--leader-elect=false"
                                                    ],
                                                    "image": "k8s.gcr.io/kube-scheduler:v1.16.0",
                                                    "imagePullPolicy": "IfNotPresent",
                                                    "livenessProbe": {
                                                        "failureThreshold": 8,
                                                        "httpGet": {
                                                            "host": "127.0.0.1",
                                                            "path": "/healthz",
                                                            "port": 10251,
                                                            "scheme": "HTTP"
                                                        },
                                                        "initialDelaySeconds": 15,
                                                        "periodSeconds": 10,
                                                        "successThreshold": 1,
                                                        "timeoutSeconds": 15
                                                    },
                                                    "name": "kube-scheduler",
                                                    "resources": {
                                                        "requests": {
                                                            "cpu": "100m"
                                                        }
                                                    },
                                                    "terminationMessagePath": "/dev/termination-log",
                                                    "terminationMessagePolicy": "File",
                                                    "volumeMounts": [
                                                        {
                                                            "mountPath": "/etc/kubernetes/scheduler.conf",
                                                            "name": "kubeconfig",
                                                            "readOnly": true
                                                        }
                                                    ]
                                                }
                                            ],
                                            "dnsPolicy": "ClusterFirst",
                                            "enableServiceLinks": true,
                                            "hostNetwork": true,
                                            "nodeName": "david-virtualbox",
                                            "priority": 2000000000,
                                            "priorityClassName": "system-cluster-critical",
                                            "restartPolicy": "Always",
                                            "schedulerName": "default-scheduler",
                                            "securityContext": {},
                                            "terminationGracePeriodSeconds": 30,
                                            "tolerations": [
                                                {
                                                    "effect": "NoExecute",
                                                    "operator": "Exists"
                                                }
                                            ],
                                            "volumes": [
                                                {
                                                    "hostPath": {
                                                        "path": "/etc/kubernetes/scheduler.conf",
                                                        "type": "FileOrCreate"
                                                    },
                                                    "name": "kubeconfig"
                                                }
                                            ]
                                        },
                                        "status": {
                                            "conditions": [
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-14T05:35:11Z",
                                                    "status": "True",
                                                    "type": "Initialized"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-14T05:35:13Z",
                                                    "status": "True",
                                                    "type": "Ready"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-14T05:35:13Z",
                                                    "status": "True",
                                                    "type": "ContainersReady"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-14T05:35:11Z",
                                                    "status": "True",
                                                    "type": "PodScheduled"
                                                }
                                            ],
                                            "containerStatuses": [
                                                {
                                                    "containerID": "docker://53c67aa4e9ac29bf19ebaa451fe9de11d2bad3442dc6ece2dc3e03a0dc266526",
                                                    "image": "k8s.gcr.io/kube-scheduler:v1.16.0",
                                                    "imageID": "docker-pullable://k8s.gcr.io/kube-scheduler@sha256:094023ab9cd02059eb0295d234ff9ea321e0e22e4813986d7f1a1ac4dc1990d0",
                                                    "lastState": {
                                                        "terminated": {
                                                            "containerID": "docker://fbe8092ea4bbc12f8ed65da47d80b24528557bc81a7fc54948da26e76d704b4a",
                                                            "exitCode": 255,
                                                            "finishedAt": "2021-10-14T05:35:02Z",
                                                            "reason": "Error",
                                                            "startedAt": "2021-10-12T10:28:38Z"
                                                        }
                                                    },
                                                    "name": "kube-scheduler",
                                                    "ready": true,
                                                    "restartCount": 52,
                                                    "started": true,
                                                    "state": {
                                                        "running": {
                                                            "startedAt": "2021-10-14T05:35:12Z"
                                                        }
                                                    }
                                                }
                                            ],
                                            "hostIP": "10.0.2.15",
                                            "phase": "Running",
                                            "podIP": "10.0.2.15",
                                            "podIPs": [
                                                {
                                                    "ip": "10.0.2.15"
                                                }
                                            ],
                                            "qosClass": "Burstable",
                                            "startTime": "2021-10-14T05:35:11Z"
                                        }
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "Pod: storage-provisioner is connected to the host network",
                            "ruleStatus": "",
                            "packagename": "armo_builtins",
                            "alertScore": 9,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "Pod",
                                        "metadata": {
                                            "annotations": {
                                                "kubectl.kubernetes.io/last-applied-configuration": "{\"apiVersion\":\"v1\",\"kind\":\"Pod\",\"metadata\":{\"annotations\":{},\"labels\":{\"addonmanager.kubernetes.io/mode\":\"Reconcile\",\"integration-test\":\"storage-provisioner\"},\"name\":\"storage-provisioner\",\"namespace\":\"kube-system\"},\"spec\":{\"containers\":[{\"command\":[\"/storage-provisioner\"],\"image\":\"gcr.io/k8s-minikube/storage-provisioner:v4\",\"imagePullPolicy\":\"IfNotPresent\",\"name\":\"storage-provisioner\",\"volumeMounts\":[{\"mountPath\":\"/tmp\",\"name\":\"tmp\"}]}],\"hostNetwork\":true,\"serviceAccountName\":\"storage-provisioner\",\"volumes\":[{\"hostPath\":{\"path\":\"/tmp\",\"type\":\"Directory\"},\"name\":\"tmp\"}]}}\n"
                                            },
                                            "creationTimestamp": "2021-06-20T09:07:09Z",
                                            "labels": {
                                                "addonmanager.kubernetes.io/mode": "Reconcile",
                                                "integration-test": "storage-provisioner"
                                            },
                                            "name": "storage-provisioner",
                                            "namespace": "kube-system",
                                            "resourceVersion": "2110982",
                                            "selfLink": "/api/v1/namespaces/kube-system/pods/storage-provisioner",
                                            "uid": "ea5dc2e2-4f7a-49f4-9e88-37e8e2d741a5"
                                        },
                                        "spec": {
                                            "containers": [
                                                {
                                                    "command": [
                                                        "/storage-provisioner"
                                                    ],
                                                    "image": "gcr.io/k8s-minikube/storage-provisioner:v4",
                                                    "imagePullPolicy": "IfNotPresent",
                                                    "name": "storage-provisioner",
                                                    "resources": {},
                                                    "terminationMessagePath": "/dev/termination-log",
                                                    "terminationMessagePolicy": "File",
                                                    "volumeMounts": [
                                                        {
                                                            "mountPath": "/tmp",
                                                            "name": "tmp"
                                                        },
                                                        {
                                                            "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount",
                                                            "name": "storage-provisioner-token-bbjlq",
                                                            "readOnly": true
                                                        }
                                                    ]
                                                }
                                            ],
                                            "dnsPolicy": "ClusterFirst",
                                            "enableServiceLinks": true,
                                            "hostNetwork": true,
                                            "nodeName": "david-virtualbox",
                                            "priority": 0,
                                            "restartPolicy": "Always",
                                            "schedulerName": "default-scheduler",
                                            "securityContext": {},
                                            "serviceAccount": "storage-provisioner",
                                            "serviceAccountName": "storage-provisioner",
                                            "terminationGracePeriodSeconds": 30,
                                            "tolerations": [
                                                {
                                                    "effect": "NoExecute",
                                                    "key": "node.kubernetes.io/not-ready",
                                                    "operator": "Exists",
                                                    "tolerationSeconds": 300
                                                },
                                                {
                                                    "effect": "NoExecute",
                                                    "key": "node.kubernetes.io/unreachable",
                                                    "operator": "Exists",
                                                    "tolerationSeconds": 300
                                                }
                                            ],
                                            "volumes": [
                                                {
                                                    "hostPath": {
                                                        "path": "/tmp",
                                                        "type": "Directory"
                                                    },
                                                    "name": "tmp"
                                                },
                                                {
                                                    "name": "storage-provisioner-token-bbjlq",
                                                    "secret": {
                                                        "defaultMode": 420,
                                                        "secretName": "storage-provisioner-token-bbjlq"
                                                    }
                                                }
                                            ]
                                        },
                                        "status": {
                                            "conditions": [
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-06-20T09:07:23Z",
                                                    "status": "True",
                                                    "type": "Initialized"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-14T05:36:17Z",
                                                    "status": "True",
                                                    "type": "Ready"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-14T05:36:17Z",
                                                    "status": "True",
                                                    "type": "ContainersReady"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-06-20T09:07:23Z",
                                                    "status": "True",
                                                    "type": "PodScheduled"
                                                }
                                            ],
                                            "containerStatuses": [
                                                {
                                                    "containerID": "docker://78d5935d6367da4a887877b19e472885edcfa17c8beb3425058f517002ccac4b",
                                                    "image": "gcr.io/k8s-minikube/storage-provisioner:v4",
                                                    "imageID": "docker-pullable://gcr.io/k8s-minikube/storage-provisioner@sha256:06f83c679a723d938b8776510d979c69549ad7df516279981e23554b3e68572f",
                                                    "lastState": {
                                                        "terminated": {
                                                            "containerID": "docker://4ac8e5fb478cd62a8edeb8a14ea0e63989df016d15f48a21279eee9a19e631a3",
                                                            "exitCode": 1,
                                                            "finishedAt": "2021-10-14T05:36:01Z",
                                                            "reason": "Error",
                                                            "startedAt": "2021-10-14T05:35:30Z"
                                                        }
                                                    },
                                                    "name": "storage-provisioner",
                                                    "ready": true,
                                                    "restartCount": 94,
                                                    "started": true,
                                                    "state": {
                                                        "running": {
                                                            "startedAt": "2021-10-14T05:36:16Z"
                                                        }
                                                    }
                                                }
                                            ],
                                            "hostIP": "10.0.2.15",
                                            "phase": "Running",
                                            "podIP": "10.0.2.15",
                                            "podIPs": [
                                                {
                                                    "ip": "10.0.2.15"
                                                }
                                            ],
                                            "qosClass": "BestEffort",
                                            "startTime": "2021-06-20T09:07:23Z"
                                        }
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "DaemonSet: kube-proxy has a pod connected to the host network",
                            "ruleStatus": "",
                            "packagename": "armo_builtins",
                            "alertScore": 9,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "apps/v1",
                                        "kind": "DaemonSet",
                                        "metadata": {
                                            "annotations": {
                                                "deprecated.daemonset.template.generation": "1"
                                            },
                                            "creationTimestamp": "2021-06-20T09:07:08Z",
                                            "generation": 1,
                                            "labels": {
                                                "k8s-app": "kube-proxy"
                                            },
                                            "name": "kube-proxy",
                                            "namespace": "kube-system",
                                            "resourceVersion": "1450902",
                                            "selfLink": "/apis/apps/v1/namespaces/kube-system/daemonsets/kube-proxy",
                                            "uid": "dd1ba553-66da-47bc-8bc1-79c4b2f47dab"
                                        },
                                        "spec": {
                                            "revisionHistoryLimit": 10,
                                            "selector": {
                                                "matchLabels": {
                                                    "k8s-app": "kube-proxy"
                                                }
                                            },
                                            "template": {
                                                "metadata": {
                                                    "creationTimestamp": null,
                                                    "labels": {
                                                        "k8s-app": "kube-proxy"
                                                    }
                                                },
                                                "spec": {
                                                    "containers": [
                                                        {
                                                            "command": [
                                                                "/usr/local/bin/kube-proxy",
                                                                "--config=/var/lib/kube-proxy/config.conf",
                                                                "--hostname-override=$(NODE_NAME)"
                                                            ],
                                                            "env": [
                                                                {
                                                                    "name": "NODE_NAME",
                                                                    "valueFrom": {
                                                                        "fieldRef": {
                                                                            "apiVersion": "v1",
                                                                            "fieldPath": "spec.nodeName"
                                                                        }
                                                                    }
                                                                }
                                                            ],
                                                            "image": "k8s.gcr.io/kube-proxy:v1.16.0",
                                                            "imagePullPolicy": "IfNotPresent",
                                                            "name": "kube-proxy",
                                                            "resources": {},
                                                            "securityContext": {
                                                                "privileged": true
                                                            },
                                                            "terminationMessagePath": "/dev/termination-log",
                                                            "terminationMessagePolicy": "File",
                                                            "volumeMounts": [
                                                                {
                                                                    "mountPath": "/var/lib/kube-proxy",
                                                                    "name": "kube-proxy"
                                                                },
                                                                {
                                                                    "mountPath": "/run/xtables.lock",
                                                                    "name": "xtables-lock"
                                                                },
                                                                {
                                                                    "mountPath": "/lib/modules",
                                                                    "name": "lib-modules",
                                                                    "readOnly": true
                                                                }
                                                            ]
                                                        }
                                                    ],
                                                    "dnsPolicy": "ClusterFirst",
                                                    "hostNetwork": true,
                                                    "nodeSelector": {
                                                        "beta.kubernetes.io/os": "linux"
                                                    },
                                                    "priorityClassName": "system-node-critical",
                                                    "restartPolicy": "Always",
                                                    "schedulerName": "default-scheduler",
                                                    "securityContext": {},
                                                    "serviceAccount": "kube-proxy",
                                                    "serviceAccountName": "kube-proxy",
                                                    "terminationGracePeriodSeconds": 30,
                                                    "tolerations": [
                                                        {
                                                            "key": "CriticalAddonsOnly",
                                                            "operator": "Exists"
                                                        },
                                                        {
                                                            "operator": "Exists"
                                                        }
                                                    ],
                                                    "volumes": [
                                                        {
                                                            "configMap": {
                                                                "defaultMode": 420,
                                                                "name": "kube-proxy"
                                                            },
                                                            "name": "kube-proxy"
                                                        },
                                                        {
                                                            "hostPath": {
                                                                "path": "/run/xtables.lock",
                                                                "type": "FileOrCreate"
                                                            },
                                                            "name": "xtables-lock"
                                                        },
                                                        {
                                                            "hostPath": {
                                                                "path": "/lib/modules",
                                                                "type": ""
                                                            },
                                                            "name": "lib-modules"
                                                        }
                                                    ]
                                                }
                                            },
                                            "updateStrategy": {
                                                "rollingUpdate": {
                                                    "maxUnavailable": 1
                                                },
                                                "type": "RollingUpdate"
                                            }
                                        },
                                        "status": {
                                            "currentNumberScheduled": 1,
                                            "desiredNumberScheduled": 1,
                                            "numberAvailable": 1,
                                            "numberMisscheduled": 0,
                                            "numberReady": 1,
                                            "observedGeneration": 1,
                                            "updatedNumberScheduled": 1
                                        }
                                    }
                                ]
                            }
                        }
                    ]
                }
            ],
            "remediation": "Only connect PODs to host network when it is necessary. If not, set the hostNetwork field of the pod spec to false, or completely remove it (false is the default). Whitelist only those PODs that must have access to host network by design.",
            "description": "Potential attackers may gain access to a POD and inherit access to the entire host network. For example, in AWS case, they will have access to the entire VPC. This control identifies all the PODs with host network access enabled.",
            "score": 14
        },
        {
            "guid": "",
            "attributes": {
                "armoBuiltin": true
            },
            "id": "C-0009",
            "controlID": "C-0009",
            "name": "Resource policies",
            "ruleReports": [
                {
                    "name": "resource-policies",
                    "remediation": "",
                    "ruleStatus": {
                        "status": "success",
                        "message": ""
                    },
                    "ruleResponses": [
                        {
                            "alertMessage": "there are no resource limits defined for container : etcd",
                            "ruleStatus": "warning",
                            "packagename": "armo_builtins",
                            "alertScore": 7,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "Pod",
                                        "metadata": {
                                            "annotations": {
                                                "kubernetes.io/config.hash": "e0fcc6e4323055b5880f8aac4c950836",
                                                "kubernetes.io/config.mirror": "e0fcc6e4323055b5880f8aac4c950836",
                                                "kubernetes.io/config.seen": "2021-06-20T12:06:52.495386281+03:00",
                                                "kubernetes.io/config.source": "file"
                                            },
                                            "creationTimestamp": "2021-06-20T09:08:22Z",
                                            "labels": {
                                                "component": "etcd",
                                                "tier": "control-plane"
                                            },
                                            "name": "etcd-david-virtualbox",
                                            "namespace": "kube-system",
                                            "resourceVersion": "2110909",
                                            "selfLink": "/api/v1/namespaces/kube-system/pods/etcd-david-virtualbox",
                                            "uid": "154e7f87-907f-4edb-a73c-26e965d4fe02"
                                        },
                                        "spec": {
                                            "containers": [
                                                {
                                                    "command": [
                                                        "etcd",
                                                        "--advertise-client-urls=https://10.0.2.15:2379",
                                                        "--cert-file=/var/lib/minikube/certs/etcd/server.crt",
                                                        "--client-cert-auth=true",
                                                        "--data-dir=/var/lib/minikube/etcd",
                                                        "--initial-advertise-peer-urls=https://10.0.2.15:2380",
                                                        "--initial-cluster=david-virtualbox=https://10.0.2.15:2380",
                                                        "--key-file=/var/lib/minikube/certs/etcd/server.key",
                                                        "--listen-client-urls=https://127.0.0.1:2379,https://10.0.2.15:2379",
                                                        "--listen-metrics-urls=http://127.0.0.1:2381,http://10.0.2.15:2381",
                                                        "--listen-peer-urls=https://10.0.2.15:2380",
                                                        "--name=david-virtualbox",
                                                        "--peer-cert-file=/var/lib/minikube/certs/etcd/peer.crt",
                                                        "--peer-client-cert-auth=true",
                                                        "--peer-key-file=/var/lib/minikube/certs/etcd/peer.key",
                                                        "--peer-trusted-ca-file=/var/lib/minikube/certs/etcd/ca.crt",
                                                        "--snapshot-count=10000",
                                                        "--trusted-ca-file=/var/lib/minikube/certs/etcd/ca.crt"
                                                    ],
                                                    "image": "k8s.gcr.io/etcd:3.3.15-0",
                                                    "imagePullPolicy": "IfNotPresent",
                                                    "livenessProbe": {
                                                        "failureThreshold": 8,
                                                        "httpGet": {
                                                            "host": "127.0.0.1",
                                                            "path": "/health",
                                                            "port": 2381,
                                                            "scheme": "HTTP"
                                                        },
                                                        "initialDelaySeconds": 15,
                                                        "periodSeconds": 10,
                                                        "successThreshold": 1,
                                                        "timeoutSeconds": 15
                                                    },
                                                    "name": "etcd",
                                                    "resources": {},
                                                    "terminationMessagePath": "/dev/termination-log",
                                                    "terminationMessagePolicy": "File",
                                                    "volumeMounts": [
                                                        {
                                                            "mountPath": "/var/lib/minikube/etcd",
                                                            "name": "etcd-data"
                                                        },
                                                        {
                                                            "mountPath": "/var/lib/minikube/certs/etcd",
                                                            "name": "etcd-certs"
                                                        }
                                                    ]
                                                }
                                            ],
                                            "dnsPolicy": "ClusterFirst",
                                            "enableServiceLinks": true,
                                            "hostNetwork": true,
                                            "nodeName": "david-virtualbox",
                                            "priority": 2000000000,
                                            "priorityClassName": "system-cluster-critical",
                                            "restartPolicy": "Always",
                                            "schedulerName": "default-scheduler",
                                            "securityContext": {},
                                            "terminationGracePeriodSeconds": 30,
                                            "tolerations": [
                                                {
                                                    "effect": "NoExecute",
                                                    "operator": "Exists"
                                                }
                                            ],
                                            "volumes": [
                                                {
                                                    "hostPath": {
                                                        "path": "/var/lib/minikube/certs/etcd",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "etcd-certs"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/var/lib/minikube/etcd",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "etcd-data"
                                                }
                                            ]
                                        },
                                        "status": {
                                            "conditions": [
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-03T06:28:57Z",
                                                    "status": "True",
                                                    "type": "Initialized"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-06T05:23:22Z",
                                                    "status": "True",
                                                    "type": "Ready"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-06T05:23:22Z",
                                                    "status": "True",
                                                    "type": "ContainersReady"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-03T06:28:57Z",
                                                    "status": "True",
                                                    "type": "PodScheduled"
                                                }
                                            ],
                                            "containerStatuses": [
                                                {
                                                    "containerID": "docker://0cc62102444cc74e3dddb2b6cd7550036feaa0e61f4297c85fc72e3be2905ec9",
                                                    "image": "k8s.gcr.io/etcd:3.3.15-0",
                                                    "imageID": "docker-pullable://k8s.gcr.io/etcd@sha256:12c2c5e5731c3bcd56e6f1c05c0f9198b6f06793fa7fca2fb43aab9622dc4afa",
                                                    "lastState": {
                                                        "terminated": {
                                                            "containerID": "docker://61de0abaa35617dafed9274c5b03f99f269412bdea2f15d903b4293d620f4b9e",
                                                            "exitCode": 255,
                                                            "finishedAt": "2021-10-14T05:35:02Z",
                                                            "reason": "Error",
                                                            "startedAt": "2021-10-12T10:28:38Z"
                                                        }
                                                    },
                                                    "name": "etcd",
                                                    "ready": true,
                                                    "restartCount": 51,
                                                    "started": true,
                                                    "state": {
                                                        "running": {
                                                            "startedAt": "2021-10-14T05:35:12Z"
                                                        }
                                                    }
                                                }
                                            ],
                                            "hostIP": "10.0.2.15",
                                            "phase": "Running",
                                            "podIP": "10.0.2.15",
                                            "podIPs": [
                                                {
                                                    "ip": "10.0.2.15"
                                                }
                                            ],
                                            "qosClass": "BestEffort",
                                            "startTime": "2021-10-03T06:28:57Z"
                                        }
                                    }
                                ]
                            },
                            "exception": {
                                "guid": "48418e49-2b0b-43dd-86b8-bb4935204634",
                                "name": "exception_C-0009_kube-system_a4a42f46c70fda8b3631275369505ed2",
                                "attributes": {
                                    "namespaceOnly": "true"
                                },
                                "policyType": "postureExceptionPolicy",
                                "creationTime": "2021-10-14T12:13:06.754963",
                                "actions": [
                                    "alertOnly"
                                ],
                                "resources": [
                                    {
                                        "designatorType": "Attributes",
                                        "wlid": "",
                                        "wildwlid": "",
                                        "sid": "",
                                        "attributes": {
                                            "cluster": "minikube",
                                            "namespace": "kube-system"
                                        }
                                    }
                                ],
                                "posturePolicies": [
                                    {
                                        "frameworkName": "NSA",
                                        "controlName": "Resource policies",
                                        "ruleName": ""
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "there are no resource limits defined for container : kube-apiserver",
                            "ruleStatus": "warning",
                            "packagename": "armo_builtins",
                            "alertScore": 7,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "Pod",
                                        "metadata": {
                                            "annotations": {
                                                "kubernetes.io/config.hash": "668e9396bc7b8c495d39f4a3bc479397",
                                                "kubernetes.io/config.mirror": "668e9396bc7b8c495d39f4a3bc479397",
                                                "kubernetes.io/config.seen": "2021-06-20T12:06:52.495387809+03:00",
                                                "kubernetes.io/config.source": "file"
                                            },
                                            "creationTimestamp": "2021-06-20T09:08:05Z",
                                            "labels": {
                                                "component": "kube-apiserver",
                                                "tier": "control-plane"
                                            },
                                            "name": "kube-apiserver-david-virtualbox",
                                            "namespace": "kube-system",
                                            "resourceVersion": "2110920",
                                            "selfLink": "/api/v1/namespaces/kube-system/pods/kube-apiserver-david-virtualbox",
                                            "uid": "327cbf13-97d6-42a3-8469-5acfdbe1be09"
                                        },
                                        "spec": {
                                            "containers": [
                                                {
                                                    "command": [
                                                        "kube-apiserver",
                                                        "--advertise-address=10.0.2.15",
                                                        "--allow-privileged=true",
                                                        "--authorization-mode=Node,RBAC",
                                                        "--client-ca-file=/var/lib/minikube/certs/ca.crt",
                                                        "--enable-admission-plugins=NamespaceLifecycle,LimitRanger,ServiceAccount,DefaultStorageClass,DefaultTolerationSeconds,NodeRestriction,MutatingAdmissionWebhook,ValidatingAdmissionWebhook,ResourceQuota",
                                                        "--enable-bootstrap-token-auth=true",
                                                        "--etcd-cafile=/var/lib/minikube/certs/etcd/ca.crt",
                                                        "--etcd-certfile=/var/lib/minikube/certs/apiserver-etcd-client.crt",
                                                        "--etcd-keyfile=/var/lib/minikube/certs/apiserver-etcd-client.key",
                                                        "--etcd-servers=https://127.0.0.1:2379",
                                                        "--insecure-port=0",
                                                        "--kubelet-client-certificate=/var/lib/minikube/certs/apiserver-kubelet-client.crt",
                                                        "--kubelet-client-key=/var/lib/minikube/certs/apiserver-kubelet-client.key",
                                                        "--kubelet-preferred-address-types=InternalIP,ExternalIP,Hostname",
                                                        "--proxy-client-cert-file=/var/lib/minikube/certs/front-proxy-client.crt",
                                                        "--proxy-client-key-file=/var/lib/minikube/certs/front-proxy-client.key",
                                                        "--requestheader-allowed-names=front-proxy-client",
                                                        "--requestheader-client-ca-file=/var/lib/minikube/certs/front-proxy-ca.crt",
                                                        "--requestheader-extra-headers-prefix=X-Remote-Extra-",
                                                        "--requestheader-group-headers=X-Remote-Group",
                                                        "--requestheader-username-headers=X-Remote-User",
                                                        "--secure-port=8443",
                                                        "--service-account-key-file=/var/lib/minikube/certs/sa.pub",
                                                        "--service-cluster-ip-range=10.96.0.0/12",
                                                        "--tls-cert-file=/var/lib/minikube/certs/apiserver.crt",
                                                        "--tls-private-key-file=/var/lib/minikube/certs/apiserver.key"
                                                    ],
                                                    "image": "k8s.gcr.io/kube-apiserver:v1.16.0",
                                                    "imagePullPolicy": "IfNotPresent",
                                                    "livenessProbe": {
                                                        "failureThreshold": 8,
                                                        "httpGet": {
                                                            "host": "10.0.2.15",
                                                            "path": "/healthz",
                                                            "port": 8443,
                                                            "scheme": "HTTPS"
                                                        },
                                                        "initialDelaySeconds": 15,
                                                        "periodSeconds": 10,
                                                        "successThreshold": 1,
                                                        "timeoutSeconds": 15
                                                    },
                                                    "name": "kube-apiserver",
                                                    "resources": {
                                                        "requests": {
                                                            "cpu": "250m"
                                                        }
                                                    },
                                                    "terminationMessagePath": "/dev/termination-log",
                                                    "terminationMessagePolicy": "File",
                                                    "volumeMounts": [
                                                        {
                                                            "mountPath": "/etc/ssl/certs",
                                                            "name": "ca-certs",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/etc/ca-certificates",
                                                            "name": "etc-ca-certificates",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/etc/pki",
                                                            "name": "etc-pki",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/var/lib/minikube/certs",
                                                            "name": "k8s-certs",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/usr/local/share/ca-certificates",
                                                            "name": "usr-local-share-ca-certificates",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/usr/share/ca-certificates",
                                                            "name": "usr-share-ca-certificates",
                                                            "readOnly": true
                                                        }
                                                    ]
                                                }
                                            ],
                                            "dnsPolicy": "ClusterFirst",
                                            "enableServiceLinks": true,
                                            "hostNetwork": true,
                                            "nodeName": "david-virtualbox",
                                            "priority": 2000000000,
                                            "priorityClassName": "system-cluster-critical",
                                            "restartPolicy": "Always",
                                            "schedulerName": "default-scheduler",
                                            "securityContext": {},
                                            "terminationGracePeriodSeconds": 30,
                                            "tolerations": [
                                                {
                                                    "effect": "NoExecute",
                                                    "operator": "Exists"
                                                }
                                            ],
                                            "volumes": [
                                                {
                                                    "hostPath": {
                                                        "path": "/etc/ssl/certs",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "ca-certs"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/etc/ca-certificates",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "etc-ca-certificates"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/etc/pki",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "etc-pki"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/var/lib/minikube/certs",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "k8s-certs"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/usr/local/share/ca-certificates",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "usr-local-share-ca-certificates"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/usr/share/ca-certificates",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "usr-share-ca-certificates"
                                                }
                                            ]
                                        },
                                        "status": {
                                            "conditions": [
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-03T06:28:57Z",
                                                    "status": "True",
                                                    "type": "Initialized"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-06T05:23:25Z",
                                                    "status": "True",
                                                    "type": "Ready"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-06T05:23:25Z",
                                                    "status": "True",
                                                    "type": "ContainersReady"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-03T06:28:57Z",
                                                    "status": "True",
                                                    "type": "PodScheduled"
                                                }
                                            ],
                                            "containerStatuses": [
                                                {
                                                    "containerID": "docker://c6d5933ac2752c18942e488d2d162fcc3141d1ae131be2abc3a50d2cb91e4e22",
                                                    "image": "k8s.gcr.io/kube-apiserver:v1.16.0",
                                                    "imageID": "docker-pullable://k8s.gcr.io/kube-apiserver@sha256:f4168527c91289da2708f62ae729fdde5fb484167dd05ffbb7ab666f60de96cd",
                                                    "lastState": {
                                                        "terminated": {
                                                            "containerID": "docker://3714c8295fe8b0ff97082305a44f08764e06d1a599ab3ee3ca6db15ab498a276",
                                                            "exitCode": 255,
                                                            "finishedAt": "2021-10-14T05:35:02Z",
                                                            "reason": "Error",
                                                            "startedAt": "2021-10-12T10:28:38Z"
                                                        }
                                                    },
                                                    "name": "kube-apiserver",
                                                    "ready": true,
                                                    "restartCount": 51,
                                                    "started": true,
                                                    "state": {
                                                        "running": {
                                                            "startedAt": "2021-10-14T05:35:12Z"
                                                        }
                                                    }
                                                }
                                            ],
                                            "hostIP": "10.0.2.15",
                                            "phase": "Running",
                                            "podIP": "10.0.2.15",
                                            "podIPs": [
                                                {
                                                    "ip": "10.0.2.15"
                                                }
                                            ],
                                            "qosClass": "Burstable",
                                            "startTime": "2021-10-03T06:28:57Z"
                                        }
                                    }
                                ]
                            },
                            "exception": {
                                "guid": "48418e49-2b0b-43dd-86b8-bb4935204634",
                                "name": "exception_C-0009_kube-system_a4a42f46c70fda8b3631275369505ed2",
                                "attributes": {
                                    "namespaceOnly": "true"
                                },
                                "policyType": "postureExceptionPolicy",
                                "creationTime": "2021-10-14T12:13:06.754963",
                                "actions": [
                                    "alertOnly"
                                ],
                                "resources": [
                                    {
                                        "designatorType": "Attributes",
                                        "wlid": "",
                                        "wildwlid": "",
                                        "sid": "",
                                        "attributes": {
                                            "cluster": "minikube",
                                            "namespace": "kube-system"
                                        }
                                    }
                                ],
                                "posturePolicies": [
                                    {
                                        "frameworkName": "NSA",
                                        "controlName": "Resource policies",
                                        "ruleName": ""
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "there are no resource limits defined for container : kube-controller-manager",
                            "ruleStatus": "warning",
                            "packagename": "armo_builtins",
                            "alertScore": 7,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "Pod",
                                        "metadata": {
                                            "annotations": {
                                                "kubernetes.io/config.hash": "a16b2d5766eae37796e4a8ed7f8ce12a",
                                                "kubernetes.io/config.mirror": "a16b2d5766eae37796e4a8ed7f8ce12a",
                                                "kubernetes.io/config.seen": "2021-06-20T12:06:52.495389283+03:00",
                                                "kubernetes.io/config.source": "file"
                                            },
                                            "creationTimestamp": "2021-06-20T09:08:00Z",
                                            "labels": {
                                                "component": "kube-controller-manager",
                                                "tier": "control-plane"
                                            },
                                            "name": "kube-controller-manager-david-virtualbox",
                                            "namespace": "kube-system",
                                            "resourceVersion": "2110899",
                                            "selfLink": "/api/v1/namespaces/kube-system/pods/kube-controller-manager-david-virtualbox",
                                            "uid": "6ca9d32c-21c3-4c0e-8087-5445c80a2bcc"
                                        },
                                        "spec": {
                                            "containers": [
                                                {
                                                    "command": [
                                                        "kube-controller-manager",
                                                        "--allocate-node-cidrs=true",
                                                        "--authentication-kubeconfig=/etc/kubernetes/controller-manager.conf",
                                                        "--authorization-kubeconfig=/etc/kubernetes/controller-manager.conf",
                                                        "--bind-address=127.0.0.1",
                                                        "--client-ca-file=/var/lib/minikube/certs/ca.crt",
                                                        "--cluster-cidr=10.244.0.0/16",
                                                        "--cluster-signing-cert-file=/var/lib/minikube/certs/ca.crt",
                                                        "--cluster-signing-key-file=/var/lib/minikube/certs/ca.key",
                                                        "--controllers=*,bootstrapsigner,tokencleaner",
                                                        "--kubeconfig=/etc/kubernetes/controller-manager.conf",
                                                        "--leader-elect=false",
                                                        "--node-cidr-mask-size=24",
                                                        "--requestheader-client-ca-file=/var/lib/minikube/certs/front-proxy-ca.crt",
                                                        "--root-ca-file=/var/lib/minikube/certs/ca.crt",
                                                        "--service-account-private-key-file=/var/lib/minikube/certs/sa.key",
                                                        "--service-cluster-ip-range=10.96.0.0/12",
                                                        "--use-service-account-credentials=true"
                                                    ],
                                                    "image": "k8s.gcr.io/kube-controller-manager:v1.16.0",
                                                    "imagePullPolicy": "IfNotPresent",
                                                    "livenessProbe": {
                                                        "failureThreshold": 8,
                                                        "httpGet": {
                                                            "host": "127.0.0.1",
                                                            "path": "/healthz",
                                                            "port": 10252,
                                                            "scheme": "HTTP"
                                                        },
                                                        "initialDelaySeconds": 15,
                                                        "periodSeconds": 10,
                                                        "successThreshold": 1,
                                                        "timeoutSeconds": 15
                                                    },
                                                    "name": "kube-controller-manager",
                                                    "resources": {
                                                        "requests": {
                                                            "cpu": "200m"
                                                        }
                                                    },
                                                    "terminationMessagePath": "/dev/termination-log",
                                                    "terminationMessagePolicy": "File",
                                                    "volumeMounts": [
                                                        {
                                                            "mountPath": "/etc/ssl/certs",
                                                            "name": "ca-certs",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/etc/ca-certificates",
                                                            "name": "etc-ca-certificates",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/etc/pki",
                                                            "name": "etc-pki",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/usr/libexec/kubernetes/kubelet-plugins/volume/exec",
                                                            "name": "flexvolume-dir"
                                                        },
                                                        {
                                                            "mountPath": "/var/lib/minikube/certs",
                                                            "name": "k8s-certs",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/etc/kubernetes/controller-manager.conf",
                                                            "name": "kubeconfig",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/usr/local/share/ca-certificates",
                                                            "name": "usr-local-share-ca-certificates",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/usr/share/ca-certificates",
                                                            "name": "usr-share-ca-certificates",
                                                            "readOnly": true
                                                        }
                                                    ]
                                                }
                                            ],
                                            "dnsPolicy": "ClusterFirst",
                                            "enableServiceLinks": true,
                                            "hostNetwork": true,
                                            "nodeName": "david-virtualbox",
                                            "priority": 2000000000,
                                            "priorityClassName": "system-cluster-critical",
                                            "restartPolicy": "Always",
                                            "schedulerName": "default-scheduler",
                                            "securityContext": {},
                                            "terminationGracePeriodSeconds": 30,
                                            "tolerations": [
                                                {
                                                    "effect": "NoExecute",
                                                    "operator": "Exists"
                                                }
                                            ],
                                            "volumes": [
                                                {
                                                    "hostPath": {
                                                        "path": "/etc/ssl/certs",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "ca-certs"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/etc/ca-certificates",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "etc-ca-certificates"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/etc/pki",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "etc-pki"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/usr/libexec/kubernetes/kubelet-plugins/volume/exec",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "flexvolume-dir"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/var/lib/minikube/certs",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "k8s-certs"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/etc/kubernetes/controller-manager.conf",
                                                        "type": "FileOrCreate"
                                                    },
                                                    "name": "kubeconfig"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/usr/local/share/ca-certificates",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "usr-local-share-ca-certificates"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/usr/share/ca-certificates",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "usr-share-ca-certificates"
                                                }
                                            ]
                                        },
                                        "status": {
                                            "conditions": [
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-14T05:35:11Z",
                                                    "status": "True",
                                                    "type": "Initialized"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-14T05:35:13Z",
                                                    "status": "True",
                                                    "type": "Ready"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-14T05:35:13Z",
                                                    "status": "True",
                                                    "type": "ContainersReady"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-14T05:35:11Z",
                                                    "status": "True",
                                                    "type": "PodScheduled"
                                                }
                                            ],
                                            "containerStatuses": [
                                                {
                                                    "containerID": "docker://eefb83fb1b81a497703013f609a54ffc03b57157395cb5f9a128dfeffea54b58",
                                                    "image": "k8s.gcr.io/kube-controller-manager:v1.16.0",
                                                    "imageID": "docker-pullable://k8s.gcr.io/kube-controller-manager@sha256:c156a05ee9d40e3ca2ebf9337f38a10558c1fc6c9124006f128a82e6c38cdf3e",
                                                    "lastState": {
                                                        "terminated": {
                                                            "containerID": "docker://81d8dca7f5b87e8b9bd8f515e9bcf86f4614306a9c729c30aad92560f520ed54",
                                                            "exitCode": 255,
                                                            "finishedAt": "2021-10-14T05:35:02Z",
                                                            "reason": "Error",
                                                            "startedAt": "2021-10-12T10:28:38Z"
                                                        }
                                                    },
                                                    "name": "kube-controller-manager",
                                                    "ready": true,
                                                    "restartCount": 55,
                                                    "started": true,
                                                    "state": {
                                                        "running": {
                                                            "startedAt": "2021-10-14T05:35:12Z"
                                                        }
                                                    }
                                                }
                                            ],
                                            "hostIP": "10.0.2.15",
                                            "phase": "Running",
                                            "podIP": "10.0.2.15",
                                            "podIPs": [
                                                {
                                                    "ip": "10.0.2.15"
                                                }
                                            ],
                                            "qosClass": "Burstable",
                                            "startTime": "2021-10-14T05:35:11Z"
                                        }
                                    }
                                ]
                            },
                            "exception": {
                                "guid": "48418e49-2b0b-43dd-86b8-bb4935204634",
                                "name": "exception_C-0009_kube-system_a4a42f46c70fda8b3631275369505ed2",
                                "attributes": {
                                    "namespaceOnly": "true"
                                },
                                "policyType": "postureExceptionPolicy",
                                "creationTime": "2021-10-14T12:13:06.754963",
                                "actions": [
                                    "alertOnly"
                                ],
                                "resources": [
                                    {
                                        "designatorType": "Attributes",
                                        "wlid": "",
                                        "wildwlid": "",
                                        "sid": "",
                                        "attributes": {
                                            "cluster": "minikube",
                                            "namespace": "kube-system"
                                        }
                                    }
                                ],
                                "posturePolicies": [
                                    {
                                        "frameworkName": "NSA",
                                        "controlName": "Resource policies",
                                        "ruleName": ""
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "there are no resource limits defined for container : kube-scheduler",
                            "ruleStatus": "warning",
                            "packagename": "armo_builtins",
                            "alertScore": 7,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "Pod",
                                        "metadata": {
                                            "annotations": {
                                                "kubernetes.io/config.hash": "b3d303074fe0ca1d42a8bd9ed248df09",
                                                "kubernetes.io/config.mirror": "b3d303074fe0ca1d42a8bd9ed248df09",
                                                "kubernetes.io/config.seen": "2021-06-20T12:06:52.495381685+03:00",
                                                "kubernetes.io/config.source": "file"
                                            },
                                            "creationTimestamp": "2021-06-20T09:08:11Z",
                                            "labels": {
                                                "component": "kube-scheduler",
                                                "tier": "control-plane"
                                            },
                                            "name": "kube-scheduler-david-virtualbox",
                                            "namespace": "kube-system",
                                            "resourceVersion": "2110888",
                                            "selfLink": "/api/v1/namespaces/kube-system/pods/kube-scheduler-david-virtualbox",
                                            "uid": "226e285e-b2b3-423c-8bd4-04cfb775cbc6"
                                        },
                                        "spec": {
                                            "containers": [
                                                {
                                                    "command": [
                                                        "kube-scheduler",
                                                        "--authentication-kubeconfig=/etc/kubernetes/scheduler.conf",
                                                        "--authorization-kubeconfig=/etc/kubernetes/scheduler.conf",
                                                        "--bind-address=127.0.0.1",
                                                        "--kubeconfig=/etc/kubernetes/scheduler.conf",
                                                        "--leader-elect=false"
                                                    ],
                                                    "image": "k8s.gcr.io/kube-scheduler:v1.16.0",
                                                    "imagePullPolicy": "IfNotPresent",
                                                    "livenessProbe": {
                                                        "failureThreshold": 8,
                                                        "httpGet": {
                                                            "host": "127.0.0.1",
                                                            "path": "/healthz",
                                                            "port": 10251,
                                                            "scheme": "HTTP"
                                                        },
                                                        "initialDelaySeconds": 15,
                                                        "periodSeconds": 10,
                                                        "successThreshold": 1,
                                                        "timeoutSeconds": 15
                                                    },
                                                    "name": "kube-scheduler",
                                                    "resources": {
                                                        "requests": {
                                                            "cpu": "100m"
                                                        }
                                                    },
                                                    "terminationMessagePath": "/dev/termination-log",
                                                    "terminationMessagePolicy": "File",
                                                    "volumeMounts": [
                                                        {
                                                            "mountPath": "/etc/kubernetes/scheduler.conf",
                                                            "name": "kubeconfig",
                                                            "readOnly": true
                                                        }
                                                    ]
                                                }
                                            ],
                                            "dnsPolicy": "ClusterFirst",
                                            "enableServiceLinks": true,
                                            "hostNetwork": true,
                                            "nodeName": "david-virtualbox",
                                            "priority": 2000000000,
                                            "priorityClassName": "system-cluster-critical",
                                            "restartPolicy": "Always",
                                            "schedulerName": "default-scheduler",
                                            "securityContext": {},
                                            "terminationGracePeriodSeconds": 30,
                                            "tolerations": [
                                                {
                                                    "effect": "NoExecute",
                                                    "operator": "Exists"
                                                }
                                            ],
                                            "volumes": [
                                                {
                                                    "hostPath": {
                                                        "path": "/etc/kubernetes/scheduler.conf",
                                                        "type": "FileOrCreate"
                                                    },
                                                    "name": "kubeconfig"
                                                }
                                            ]
                                        },
                                        "status": {
                                            "conditions": [
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-14T05:35:11Z",
                                                    "status": "True",
                                                    "type": "Initialized"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-14T05:35:13Z",
                                                    "status": "True",
                                                    "type": "Ready"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-14T05:35:13Z",
                                                    "status": "True",
                                                    "type": "ContainersReady"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-14T05:35:11Z",
                                                    "status": "True",
                                                    "type": "PodScheduled"
                                                }
                                            ],
                                            "containerStatuses": [
                                                {
                                                    "containerID": "docker://53c67aa4e9ac29bf19ebaa451fe9de11d2bad3442dc6ece2dc3e03a0dc266526",
                                                    "image": "k8s.gcr.io/kube-scheduler:v1.16.0",
                                                    "imageID": "docker-pullable://k8s.gcr.io/kube-scheduler@sha256:094023ab9cd02059eb0295d234ff9ea321e0e22e4813986d7f1a1ac4dc1990d0",
                                                    "lastState": {
                                                        "terminated": {
                                                            "containerID": "docker://fbe8092ea4bbc12f8ed65da47d80b24528557bc81a7fc54948da26e76d704b4a",
                                                            "exitCode": 255,
                                                            "finishedAt": "2021-10-14T05:35:02Z",
                                                            "reason": "Error",
                                                            "startedAt": "2021-10-12T10:28:38Z"
                                                        }
                                                    },
                                                    "name": "kube-scheduler",
                                                    "ready": true,
                                                    "restartCount": 52,
                                                    "started": true,
                                                    "state": {
                                                        "running": {
                                                            "startedAt": "2021-10-14T05:35:12Z"
                                                        }
                                                    }
                                                }
                                            ],
                                            "hostIP": "10.0.2.15",
                                            "phase": "Running",
                                            "podIP": "10.0.2.15",
                                            "podIPs": [
                                                {
                                                    "ip": "10.0.2.15"
                                                }
                                            ],
                                            "qosClass": "Burstable",
                                            "startTime": "2021-10-14T05:35:11Z"
                                        }
                                    }
                                ]
                            },
                            "exception": {
                                "guid": "48418e49-2b0b-43dd-86b8-bb4935204634",
                                "name": "exception_C-0009_kube-system_a4a42f46c70fda8b3631275369505ed2",
                                "attributes": {
                                    "namespaceOnly": "true"
                                },
                                "policyType": "postureExceptionPolicy",
                                "creationTime": "2021-10-14T12:13:06.754963",
                                "actions": [
                                    "alertOnly"
                                ],
                                "resources": [
                                    {
                                        "designatorType": "Attributes",
                                        "wlid": "",
                                        "wildwlid": "",
                                        "sid": "",
                                        "attributes": {
                                            "cluster": "minikube",
                                            "namespace": "kube-system"
                                        }
                                    }
                                ],
                                "posturePolicies": [
                                    {
                                        "frameworkName": "NSA",
                                        "controlName": "Resource policies",
                                        "ruleName": ""
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "there are no resource limits defined for container : storage-provisioner",
                            "ruleStatus": "warning",
                            "packagename": "armo_builtins",
                            "alertScore": 7,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "Pod",
                                        "metadata": {
                                            "annotations": {
                                                "kubectl.kubernetes.io/last-applied-configuration": "{\"apiVersion\":\"v1\",\"kind\":\"Pod\",\"metadata\":{\"annotations\":{},\"labels\":{\"addonmanager.kubernetes.io/mode\":\"Reconcile\",\"integration-test\":\"storage-provisioner\"},\"name\":\"storage-provisioner\",\"namespace\":\"kube-system\"},\"spec\":{\"containers\":[{\"command\":[\"/storage-provisioner\"],\"image\":\"gcr.io/k8s-minikube/storage-provisioner:v4\",\"imagePullPolicy\":\"IfNotPresent\",\"name\":\"storage-provisioner\",\"volumeMounts\":[{\"mountPath\":\"/tmp\",\"name\":\"tmp\"}]}],\"hostNetwork\":true,\"serviceAccountName\":\"storage-provisioner\",\"volumes\":[{\"hostPath\":{\"path\":\"/tmp\",\"type\":\"Directory\"},\"name\":\"tmp\"}]}}\n"
                                            },
                                            "creationTimestamp": "2021-06-20T09:07:09Z",
                                            "labels": {
                                                "addonmanager.kubernetes.io/mode": "Reconcile",
                                                "integration-test": "storage-provisioner"
                                            },
                                            "name": "storage-provisioner",
                                            "namespace": "kube-system",
                                            "resourceVersion": "2110982",
                                            "selfLink": "/api/v1/namespaces/kube-system/pods/storage-provisioner",
                                            "uid": "ea5dc2e2-4f7a-49f4-9e88-37e8e2d741a5"
                                        },
                                        "spec": {
                                            "containers": [
                                                {
                                                    "command": [
                                                        "/storage-provisioner"
                                                    ],
                                                    "image": "gcr.io/k8s-minikube/storage-provisioner:v4",
                                                    "imagePullPolicy": "IfNotPresent",
                                                    "name": "storage-provisioner",
                                                    "resources": {},
                                                    "terminationMessagePath": "/dev/termination-log",
                                                    "terminationMessagePolicy": "File",
                                                    "volumeMounts": [
                                                        {
                                                            "mountPath": "/tmp",
                                                            "name": "tmp"
                                                        },
                                                        {
                                                            "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount",
                                                            "name": "storage-provisioner-token-bbjlq",
                                                            "readOnly": true
                                                        }
                                                    ]
                                                }
                                            ],
                                            "dnsPolicy": "ClusterFirst",
                                            "enableServiceLinks": true,
                                            "hostNetwork": true,
                                            "nodeName": "david-virtualbox",
                                            "priority": 0,
                                            "restartPolicy": "Always",
                                            "schedulerName": "default-scheduler",
                                            "securityContext": {},
                                            "serviceAccount": "storage-provisioner",
                                            "serviceAccountName": "storage-provisioner",
                                            "terminationGracePeriodSeconds": 30,
                                            "tolerations": [
                                                {
                                                    "effect": "NoExecute",
                                                    "key": "node.kubernetes.io/not-ready",
                                                    "operator": "Exists",
                                                    "tolerationSeconds": 300
                                                },
                                                {
                                                    "effect": "NoExecute",
                                                    "key": "node.kubernetes.io/unreachable",
                                                    "operator": "Exists",
                                                    "tolerationSeconds": 300
                                                }
                                            ],
                                            "volumes": [
                                                {
                                                    "hostPath": {
                                                        "path": "/tmp",
                                                        "type": "Directory"
                                                    },
                                                    "name": "tmp"
                                                },
                                                {
                                                    "name": "storage-provisioner-token-bbjlq",
                                                    "secret": {
                                                        "defaultMode": 420,
                                                        "secretName": "storage-provisioner-token-bbjlq"
                                                    }
                                                }
                                            ]
                                        },
                                        "status": {
                                            "conditions": [
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-06-20T09:07:23Z",
                                                    "status": "True",
                                                    "type": "Initialized"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-14T05:36:17Z",
                                                    "status": "True",
                                                    "type": "Ready"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-14T05:36:17Z",
                                                    "status": "True",
                                                    "type": "ContainersReady"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-06-20T09:07:23Z",
                                                    "status": "True",
                                                    "type": "PodScheduled"
                                                }
                                            ],
                                            "containerStatuses": [
                                                {
                                                    "containerID": "docker://78d5935d6367da4a887877b19e472885edcfa17c8beb3425058f517002ccac4b",
                                                    "image": "gcr.io/k8s-minikube/storage-provisioner:v4",
                                                    "imageID": "docker-pullable://gcr.io/k8s-minikube/storage-provisioner@sha256:06f83c679a723d938b8776510d979c69549ad7df516279981e23554b3e68572f",
                                                    "lastState": {
                                                        "terminated": {
                                                            "containerID": "docker://4ac8e5fb478cd62a8edeb8a14ea0e63989df016d15f48a21279eee9a19e631a3",
                                                            "exitCode": 1,
                                                            "finishedAt": "2021-10-14T05:36:01Z",
                                                            "reason": "Error",
                                                            "startedAt": "2021-10-14T05:35:30Z"
                                                        }
                                                    },
                                                    "name": "storage-provisioner",
                                                    "ready": true,
                                                    "restartCount": 94,
                                                    "started": true,
                                                    "state": {
                                                        "running": {
                                                            "startedAt": "2021-10-14T05:36:16Z"
                                                        }
                                                    }
                                                }
                                            ],
                                            "hostIP": "10.0.2.15",
                                            "phase": "Running",
                                            "podIP": "10.0.2.15",
                                            "podIPs": [
                                                {
                                                    "ip": "10.0.2.15"
                                                }
                                            ],
                                            "qosClass": "BestEffort",
                                            "startTime": "2021-06-20T09:07:23Z"
                                        }
                                    }
                                ]
                            },
                            "exception": {
                                "guid": "48418e49-2b0b-43dd-86b8-bb4935204634",
                                "name": "exception_C-0009_kube-system_a4a42f46c70fda8b3631275369505ed2",
                                "attributes": {
                                    "namespaceOnly": "true"
                                },
                                "policyType": "postureExceptionPolicy",
                                "creationTime": "2021-10-14T12:13:06.754963",
                                "actions": [
                                    "alertOnly"
                                ],
                                "resources": [
                                    {
                                        "designatorType": "Attributes",
                                        "wlid": "",
                                        "wildwlid": "",
                                        "sid": "",
                                        "attributes": {
                                            "cluster": "minikube",
                                            "namespace": "kube-system"
                                        }
                                    }
                                ],
                                "posturePolicies": [
                                    {
                                        "frameworkName": "NSA",
                                        "controlName": "Resource policies",
                                        "ruleName": ""
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "there are no resource limits defined for container : kube-proxy",
                            "ruleStatus": "warning",
                            "packagename": "armo_builtins",
                            "alertScore": 7,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "apps/v1",
                                        "kind": "DaemonSet",
                                        "metadata": {
                                            "annotations": {
                                                "deprecated.daemonset.template.generation": "1"
                                            },
                                            "creationTimestamp": "2021-06-20T09:07:08Z",
                                            "generation": 1,
                                            "labels": {
                                                "k8s-app": "kube-proxy"
                                            },
                                            "name": "kube-proxy",
                                            "namespace": "kube-system",
                                            "resourceVersion": "1450902",
                                            "selfLink": "/apis/apps/v1/namespaces/kube-system/daemonsets/kube-proxy",
                                            "uid": "dd1ba553-66da-47bc-8bc1-79c4b2f47dab"
                                        },
                                        "spec": {
                                            "revisionHistoryLimit": 10,
                                            "selector": {
                                                "matchLabels": {
                                                    "k8s-app": "kube-proxy"
                                                }
                                            },
                                            "template": {
                                                "metadata": {
                                                    "creationTimestamp": null,
                                                    "labels": {
                                                        "k8s-app": "kube-proxy"
                                                    }
                                                },
                                                "spec": {
                                                    "containers": [
                                                        {
                                                            "command": [
                                                                "/usr/local/bin/kube-proxy",
                                                                "--config=/var/lib/kube-proxy/config.conf",
                                                                "--hostname-override=$(NODE_NAME)"
                                                            ],
                                                            "env": [
                                                                {
                                                                    "name": "NODE_NAME",
                                                                    "valueFrom": {
                                                                        "fieldRef": {
                                                                            "apiVersion": "v1",
                                                                            "fieldPath": "spec.nodeName"
                                                                        }
                                                                    }
                                                                }
                                                            ],
                                                            "image": "k8s.gcr.io/kube-proxy:v1.16.0",
                                                            "imagePullPolicy": "IfNotPresent",
                                                            "name": "kube-proxy",
                                                            "resources": {},
                                                            "securityContext": {
                                                                "privileged": true
                                                            },
                                                            "terminationMessagePath": "/dev/termination-log",
                                                            "terminationMessagePolicy": "File",
                                                            "volumeMounts": [
                                                                {
                                                                    "mountPath": "/var/lib/kube-proxy",
                                                                    "name": "kube-proxy"
                                                                },
                                                                {
                                                                    "mountPath": "/run/xtables.lock",
                                                                    "name": "xtables-lock"
                                                                },
                                                                {
                                                                    "mountPath": "/lib/modules",
                                                                    "name": "lib-modules",
                                                                    "readOnly": true
                                                                }
                                                            ]
                                                        }
                                                    ],
                                                    "dnsPolicy": "ClusterFirst",
                                                    "hostNetwork": true,
                                                    "nodeSelector": {
                                                        "beta.kubernetes.io/os": "linux"
                                                    },
                                                    "priorityClassName": "system-node-critical",
                                                    "restartPolicy": "Always",
                                                    "schedulerName": "default-scheduler",
                                                    "securityContext": {},
                                                    "serviceAccount": "kube-proxy",
                                                    "serviceAccountName": "kube-proxy",
                                                    "terminationGracePeriodSeconds": 30,
                                                    "tolerations": [
                                                        {
                                                            "key": "CriticalAddonsOnly",
                                                            "operator": "Exists"
                                                        },
                                                        {
                                                            "operator": "Exists"
                                                        }
                                                    ],
                                                    "volumes": [
                                                        {
                                                            "configMap": {
                                                                "defaultMode": 420,
                                                                "name": "kube-proxy"
                                                            },
                                                            "name": "kube-proxy"
                                                        },
                                                        {
                                                            "hostPath": {
                                                                "path": "/run/xtables.lock",
                                                                "type": "FileOrCreate"
                                                            },
                                                            "name": "xtables-lock"
                                                        },
                                                        {
                                                            "hostPath": {
                                                                "path": "/lib/modules",
                                                                "type": ""
                                                            },
                                                            "name": "lib-modules"
                                                        }
                                                    ]
                                                }
                                            },
                                            "updateStrategy": {
                                                "rollingUpdate": {
                                                    "maxUnavailable": 1
                                                },
                                                "type": "RollingUpdate"
                                            }
                                        },
                                        "status": {
                                            "currentNumberScheduled": 1,
                                            "desiredNumberScheduled": 1,
                                            "numberAvailable": 1,
                                            "numberMisscheduled": 0,
                                            "numberReady": 1,
                                            "observedGeneration": 1,
                                            "updatedNumberScheduled": 1
                                        }
                                    }
                                ]
                            },
                            "exception": {
                                "guid": "48418e49-2b0b-43dd-86b8-bb4935204634",
                                "name": "exception_C-0009_kube-system_a4a42f46c70fda8b3631275369505ed2",
                                "attributes": {
                                    "namespaceOnly": "true"
                                },
                                "policyType": "postureExceptionPolicy",
                                "creationTime": "2021-10-14T12:13:06.754963",
                                "actions": [
                                    "alertOnly"
                                ],
                                "resources": [
                                    {
                                        "designatorType": "Attributes",
                                        "wlid": "",
                                        "wildwlid": "",
                                        "sid": "",
                                        "attributes": {
                                            "cluster": "minikube",
                                            "namespace": "kube-system"
                                        }
                                    }
                                ],
                                "posturePolicies": [
                                    {
                                        "frameworkName": "NSA",
                                        "controlName": "Resource policies",
                                        "ruleName": ""
                                    }
                                ]
                            }
                        }
                    ]
                }
            ],
            "remediation": "Define LimitRange and ResourceQuota policies to limit resource usage for namespaces or nodes.",
            "description": "CPU and memory resources should have a limit set for every container to prevent resource exhaustion. This control identifies all the Pods without resource limit definition.",
            "score": 100
        },
        {
            "guid": "",
            "attributes": {
                "armoBuiltin": true,
                "microsoftMitreColumns": [
                    "Initial Access"
                ]
            },
            "id": "C-0047",
            "controlID": "C-0047",
            "name": "Exposed dashboard",
            "ruleReports": [
                {
                    "name": "rule-exposed-dashboard",
                    "remediation": "",
                    "ruleStatus": {
                        "status": "success",
                        "message": ""
                    },
                    "ruleResponses": null
                }
            ],
            "remediation": "Update dashboard version to v2.0.1 and above.",
            "description": "Kubernetes dashboard versions before v2.0.1 do not support user authentication. If exposed externally, it will allow unauthenticated remote management of the cluster. This control checks presence of the kubernetes-dashboard deployment and its version number.",
            "score": 100
        },
        {
            "guid": "",
            "attributes": {
                "armoBuiltin": true
            },
            "id": "C-0016",
            "controlID": "C-0016",
            "name": "Allow privilege escalation",
            "ruleReports": [
                {
                    "name": "rule-allow-privilege-escalation",
                    "remediation": "",
                    "ruleStatus": {
                        "status": "success",
                        "message": ""
                    },
                    "ruleResponses": null
                }
            ],
            "remediation": "If your application does not need it, make sure the allowPrivilegeEscalation field of the securityContext is set to false.",
            "description": "Attackers may gain access to a container and uplift its privilege to enable excessive capabilities.",
            "score": 100
        },
        {
            "guid": "",
            "attributes": {
                "armoBuiltin": true,
                "microsoftMitreColumns": [
                    "Credential access",
                    "Lateral Movement"
                ]
            },
            "id": "C-0012",
            "controlID": "C-0012",
            "name": "Applications credentials in configuration files",
            "ruleReports": [
                {
                    "name": "rule-credentials-in-env-var",
                    "remediation": "",
                    "ruleStatus": {
                        "status": "success",
                        "message": ""
                    },
                    "ruleResponses": null
                },
                {
                    "name": "rule-credentials-configmap",
                    "remediation": "",
                    "ruleStatus": {
                        "status": "success",
                        "message": ""
                    },
                    "ruleResponses": [
                        {
                            "alertMessage": "this configmap has sensitive information: extension-apiserver-authentication",
                            "ruleStatus": "",
                            "packagename": "armo_builtins",
                            "alertScore": 9,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "data": {
                                            "client-ca-file": "-----BEGIN CERTIFICATE-----\nMIIDBjCCAe6gAwIBAgIBATANBgkqhkiG9w0BAQsFADAVMRMwEQYDVQQDEwptaW5p\na3ViZUNBMB4XDTIxMDIyODA2MDAzN1oXDTMxMDIyNzA2MDAzN1owFTETMBEGA1UE\nAxMKbWluaWt1YmVDQTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAL6m\n1wdcQZnc+3wpRABksEgGaQnYFwENQxkSjct/DAr3bKcsZU3hg6fS7Qupw9A+xgfK\nVYo6KZIzjzOJVfpfZTIdXe5eRoGoiQQ2znl45+rnRoMf7q3vkbdvLQZBRWmeXNGl\ngcl64/lzc5H1udrAx2aaNDufB6hb4G9SOWG6UCaBvvMzQq/VUp8G0OV1A7N27Iyf\nURUgna43wgyuGozWiq7gwQ9C9j4xzZMjOFux5AFhWIwxXOBSY4u/UCo+C4YWFzRU\n5wIOIEeSyLDPPSP+Q6xlx2YOSM1iQSIqbS2Q3kd0jqTHMsKIdh5JajHhB8CrouUc\ncNZNX+D+qVOSDZY8Nk8CAwEAAaNhMF8wDgYDVR0PAQH/BAQDAgKkMB0GA1UdJQQW\nMBQGCCsGAQUFBwMCBggrBgEFBQcDATAPBgNVHRMBAf8EBTADAQH/MB0GA1UdDgQW\nBBQpkrXoKj6LBuW349y+wPB488LFrzANBgkqhkiG9w0BAQsFAAOCAQEARqTf4nNB\n6zV2f+GVH1GMVaEyzOu8CHN8pwz0kfwhP0hW8BDYGGwvIxWZ5IMlT1xBpiRYgQdk\n17VVzjQiZT+MNlIDUGi3f+iPSA+V9RJdP8kDG6e2OUIHyW2cSydxu7VqtwQKx/q8\nlJGCPmliFvMW/By4qP0/ffOvmRh218Vpv8igS+RXBcMZjfWCdt0NJsi5XHGvK5+C\nhKfZIE6Va33IkrykXuXtEiL8zYr7RaVy6KGWuoDONT3U+U1o4hegJfgv5p7wy+Zp\nBAH1XqHHpY8Gt8O4FC8QhWoYL1CFXS7KbXIJEqONzJw/6Cun23jkhavSZuzhhKj5\n8CrKoJvc2HbGYw==\n-----END CERTIFICATE-----\n",
                                            "requestheader-allowed-names": "[\"front-proxy-client\"]",
                                            "requestheader-client-ca-file": "-----BEGIN CERTIFICATE-----\nMIIC7zCCAdegAwIBAgIBADANBgkqhkiG9w0BAQsFADAZMRcwFQYDVQQDEw5mcm9u\ndC1wcm94eS1jYTAeFw0yMTA2MjAwODUzMTlaFw0zMTA2MTgwODUzMTlaMBkxFzAV\nBgNVBAMTDmZyb250LXByb3h5LWNhMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIB\nCgKCAQEAug8lXtjanZvnZ+0mBuDCQNgAqRyPMgkNCwu+je0v+hewZShqirGC/Lxi\nlDUW40Wmx+GH9G4MZ6gd+NJfT9hZ1suhXggifu2+UVv1nLuqxKwfYzzvgwHq7cAQ\nvDYUpzHCbdMYRB9FhCRJJWU7EmDxnlU5UAV1TrLlw7durE1XlAOS/POuRVMiZ4LK\nAGCGcaFpJStr4D2GPP7OhI6w4r/WOvyyjtNLJqWOOGNd/IaLV9Y+hlJq1Uc8zzuA\n/8kBB9S3DSuR1Ue7KTdT8wOBjPFX0knlo3QOspOHE6Plsz5iBkBsj5pMLgT++kwd\nvLIjT74v5sajB6D953T+5iXblhToLwIDAQABo0IwQDAOBgNVHQ8BAf8EBAMCAqQw\nDwYDVR0TAQH/BAUwAwEB/zAdBgNVHQ4EFgQUiTQU70lr6vpRs2m6dSj5JvfwCekw\nDQYJKoZIhvcNAQELBQADggEBAGuveOzNGo27wBnfvfO1Tzg35wzID86jF6S6dtcJ\nfzhNl7W7NuMIj2W/l01+2p6qLPio0N/HzPMpPwPf/MSJ+pr4l7i7wI55WAWJjD8I\nFhkIVxmgXk4Q6f0w8cn2ub2t6VX0QHeh6cYeK2QZ7XRNpWHm6sgTpyTV3lthOhGH\nAa0Snjyu+V0NGdGja6LQDRm6rmVjyn5GsnFuYzBk59h1BriC5Ss7b+jtm8qR1mQE\nCl0WeZLo9ZaqxeI8vJaDHTXPtY1tiR8slfdNktcXBJeLsXu09VxdeppJofz6UIFu\nMADKodcnm8A9KmXXX85gZ4ilJeLYCEfyKx3+zTkXGVa3gJg=\n-----END CERTIFICATE-----\n",
                                            "requestheader-extra-headers-prefix": "[\"X-Remote-Extra-\"]",
                                            "requestheader-group-headers": "[\"X-Remote-Group\"]",
                                            "requestheader-username-headers": "[\"X-Remote-User\"]"
                                        },
                                        "kind": "ConfigMap",
                                        "metadata": {
                                            "creationTimestamp": "2021-06-20T09:07:04Z",
                                            "name": "extension-apiserver-authentication",
                                            "namespace": "kube-system",
                                            "resourceVersion": "42",
                                            "selfLink": "/api/v1/namespaces/kube-system/configmaps/extension-apiserver-authentication",
                                            "uid": "b9bed01d-5bac-429b-b422-f88f809dfd57"
                                        }
                                    }
                                ]
                            }
                        }
                    ]
                }
            ],
            "remediation": "Use Kubernetes secrets or Key Management Systems to store credentials.",
            "description": "Attackers who have access to configuration files can steal the stored secrets and use them. This control checks if ConfigMaps or pod specifications have sensitive information in their configuration.",
            "score": 92
        },
        {
            "guid": "",
            "attributes": {
                "armoBuiltin": true,
                "microsoftMitreColumns": [
                    "Privilege escalation"
                ]
            },
            "id": "C-0035",
            "controlID": "C-0035",
            "name": "Cluster-admin binding",
            "ruleReports": [
                {
                    "name": "rule-list-all-cluster-admins",
                    "remediation": "",
                    "ruleStatus": {
                        "status": "success",
                        "message": ""
                    },
                    "ruleResponses": [
                        {
                            "alertMessage": "The following ServiceAccount: argocd-application-controller have high privileges, such as cluster-admin",
                            "ruleStatus": "failed",
                            "packagename": "armo_builtins",
                            "alertScore": 9,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "rbac.authorization.k8s.io/v1",
                                        "kind": "ClusterRole",
                                        "metadata": {
                                            "annotations": {
                                                "kubectl.kubernetes.io/last-applied-configuration": "{\"apiVersion\":\"rbac.authorization.k8s.io/v1\",\"kind\":\"ClusterRole\",\"metadata\":{\"annotations\":{},\"labels\":{\"app.kubernetes.io/component\":\"application-controller\",\"app.kubernetes.io/name\":\"argocd-application-controller\",\"app.kubernetes.io/part-of\":\"argocd\"},\"name\":\"argocd-application-controller\"},\"rules\":[{\"apiGroups\":[\"*\"],\"resources\":[\"*\"],\"verbs\":[\"*\"]},{\"nonResourceURLs\":[\"*\"],\"verbs\":[\"*\"]}]}\n"
                                            },
                                            "creationTimestamp": "2021-08-22T11:50:37Z",
                                            "labels": {
                                                "app.kubernetes.io/component": "application-controller",
                                                "app.kubernetes.io/name": "argocd-application-controller",
                                                "app.kubernetes.io/part-of": "argocd"
                                            },
                                            "name": "argocd-application-controller",
                                            "resourceVersion": "1427392",
                                            "selfLink": "/apis/rbac.authorization.k8s.io/v1/clusterroles/argocd-application-controller",
                                            "uid": "947b1d1e-8fae-4f64-8264-452275d0d2a3"
                                        },
                                        "rules": [
                                            {
                                                "apiGroups": [
                                                    "*"
                                                ],
                                                "resources": [
                                                    "*"
                                                ],
                                                "verbs": [
                                                    "*"
                                                ]
                                            },
                                            {
                                                "nonResourceURLs": [
                                                    "*"
                                                ],
                                                "verbs": [
                                                    "*"
                                                ]
                                            }
                                        ]
                                    },
                                    {
                                        "apiVersion": "rbac.authorization.k8s.io/v1",
                                        "kind": "ClusterRoleBinding",
                                        "metadata": {
                                            "annotations": {
                                                "kubectl.kubernetes.io/last-applied-configuration": "{\"apiVersion\":\"rbac.authorization.k8s.io/v1\",\"kind\":\"ClusterRoleBinding\",\"metadata\":{\"annotations\":{},\"labels\":{\"app.kubernetes.io/component\":\"application-controller\",\"app.kubernetes.io/name\":\"argocd-application-controller\",\"app.kubernetes.io/part-of\":\"argocd\"},\"name\":\"argocd-application-controller\"},\"roleRef\":{\"apiGroup\":\"rbac.authorization.k8s.io\",\"kind\":\"ClusterRole\",\"name\":\"argocd-application-controller\"},\"subjects\":[{\"kind\":\"ServiceAccount\",\"name\":\"argocd-application-controller\",\"namespace\":\"argocd\"}]}\n"
                                            },
                                            "creationTimestamp": "2021-08-22T11:50:37Z",
                                            "labels": {
                                                "app.kubernetes.io/component": "application-controller",
                                                "app.kubernetes.io/name": "argocd-application-controller",
                                                "app.kubernetes.io/part-of": "argocd"
                                            },
                                            "name": "argocd-application-controller",
                                            "resourceVersion": "1427398",
                                            "selfLink": "/apis/rbac.authorization.k8s.io/v1/clusterrolebindings/argocd-application-controller",
                                            "uid": "12055d3d-8f87-4051-bd24-334ba3c409b9"
                                        },
                                        "roleRef": {
                                            "apiGroup": "rbac.authorization.k8s.io",
                                            "kind": "ClusterRole",
                                            "name": "argocd-application-controller"
                                        },
                                        "subjects": [
                                            {
                                                "kind": "ServiceAccount",
                                                "name": "argocd-application-controller",
                                                "namespace": "argocd"
                                            }
                                        ]
                                    }
                                ],
                                "externalObjects": {
                                    "subject": [
                                        {
                                            "kind": "ServiceAccount",
                                            "name": "argocd-application-controller",
                                            "namespace": "argocd"
                                        }
                                    ]
                                }
                            }
                        },
                        {
                            "alertMessage": "The following Group: system:masters have high privileges, such as cluster-admin",
                            "ruleStatus": "failed",
                            "packagename": "armo_builtins",
                            "alertScore": 9,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "rbac.authorization.k8s.io/v1",
                                        "kind": "ClusterRole",
                                        "metadata": {
                                            "annotations": {
                                                "rbac.authorization.kubernetes.io/autoupdate": "true"
                                            },
                                            "creationTimestamp": "2021-06-20T09:07:04Z",
                                            "labels": {
                                                "kubernetes.io/bootstrapping": "rbac-defaults"
                                            },
                                            "name": "cluster-admin",
                                            "resourceVersion": "43",
                                            "selfLink": "/apis/rbac.authorization.k8s.io/v1/clusterroles/cluster-admin",
                                            "uid": "612c8a97-ef32-4995-8b1d-ebf60e74d514"
                                        },
                                        "rules": [
                                            {
                                                "apiGroups": [
                                                    "*"
                                                ],
                                                "resources": [
                                                    "*"
                                                ],
                                                "verbs": [
                                                    "*"
                                                ]
                                            },
                                            {
                                                "nonResourceURLs": [
                                                    "*"
                                                ],
                                                "verbs": [
                                                    "*"
                                                ]
                                            }
                                        ]
                                    },
                                    {
                                        "apiVersion": "rbac.authorization.k8s.io/v1",
                                        "kind": "ClusterRoleBinding",
                                        "metadata": {
                                            "annotations": {
                                                "rbac.authorization.kubernetes.io/autoupdate": "true"
                                            },
                                            "creationTimestamp": "2021-06-20T09:07:04Z",
                                            "labels": {
                                                "kubernetes.io/bootstrapping": "rbac-defaults"
                                            },
                                            "name": "cluster-admin",
                                            "resourceVersion": "96",
                                            "selfLink": "/apis/rbac.authorization.k8s.io/v1/clusterrolebindings/cluster-admin",
                                            "uid": "a594dc87-99a0-4154-9844-d7b5f6e1919c"
                                        },
                                        "roleRef": {
                                            "apiGroup": "rbac.authorization.k8s.io",
                                            "kind": "ClusterRole",
                                            "name": "cluster-admin"
                                        },
                                        "subjects": [
                                            {
                                                "apiGroup": "rbac.authorization.k8s.io",
                                                "kind": "Group",
                                                "name": "system:masters"
                                            }
                                        ]
                                    }
                                ],
                                "externalObjects": {
                                    "subject": [
                                        {
                                            "apiGroup": "rbac.authorization.k8s.io",
                                            "kind": "Group",
                                            "name": "system:masters"
                                        }
                                    ]
                                }
                            }
                        },
                        {
                            "alertMessage": "The following ServiceAccount: nginx-ingress have high privileges, such as cluster-admin",
                            "ruleStatus": "failed",
                            "packagename": "armo_builtins",
                            "alertScore": 9,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "rbac.authorization.k8s.io/v1",
                                        "kind": "ClusterRole",
                                        "metadata": {
                                            "annotations": {
                                                "kubectl.kubernetes.io/last-applied-configuration": "{\"apiVersion\":\"rbac.authorization.k8s.io/v1\",\"kind\":\"ClusterRole\",\"metadata\":{\"annotations\":{},\"labels\":{\"app\":\"nginx-ingress-roles\"},\"name\":\"nginx-ingress-roles\"},\"rules\":[{\"apiGroups\":[\"*\"],\"resources\":[\"*\"],\"verbs\":[\"*\"]}]}\n"
                                            },
                                            "creationTimestamp": "2021-08-18T14:58:32Z",
                                            "labels": {
                                                "app": "nginx-ingress-roles"
                                            },
                                            "name": "nginx-ingress-roles",
                                            "resourceVersion": "1355206",
                                            "selfLink": "/apis/rbac.authorization.k8s.io/v1/clusterroles/nginx-ingress-roles",
                                            "uid": "e262c786-9138-4db2-86be-1f6bd50f93bd"
                                        },
                                        "rules": [
                                            {
                                                "apiGroups": [
                                                    "*"
                                                ],
                                                "resources": [
                                                    "*"
                                                ],
                                                "verbs": [
                                                    "*"
                                                ]
                                            }
                                        ]
                                    },
                                    {
                                        "apiVersion": "rbac.authorization.k8s.io/v1",
                                        "kind": "ClusterRoleBinding",
                                        "metadata": {
                                            "annotations": {
                                                "kubectl.kubernetes.io/last-applied-configuration": "{\"apiVersion\":\"rbac.authorization.k8s.io/v1\",\"kind\":\"ClusterRoleBinding\",\"metadata\":{\"annotations\":{},\"labels\":{\"app\":\"nginx-ingress-roles-binding\"},\"name\":\"nginx-ingress-roles-binding\"},\"roleRef\":{\"apiGroup\":\"rbac.authorization.k8s.io\",\"kind\":\"ClusterRole\",\"name\":\"nginx-ingress-roles\"},\"subjects\":[{\"kind\":\"ServiceAccount\",\"name\":\"nginx-ingress\",\"namespace\":\"nginx-ingress\"}]}\n"
                                            },
                                            "creationTimestamp": "2021-08-18T14:58:57Z",
                                            "labels": {
                                                "app": "nginx-ingress-roles-binding"
                                            },
                                            "name": "nginx-ingress-roles-binding",
                                            "resourceVersion": "1355223",
                                            "selfLink": "/apis/rbac.authorization.k8s.io/v1/clusterrolebindings/nginx-ingress-roles-binding",
                                            "uid": "38821ce8-ea8f-4081-ba6c-5769a2382447"
                                        },
                                        "roleRef": {
                                            "apiGroup": "rbac.authorization.k8s.io",
                                            "kind": "ClusterRole",
                                            "name": "nginx-ingress-roles"
                                        },
                                        "subjects": [
                                            {
                                                "kind": "ServiceAccount",
                                                "name": "nginx-ingress",
                                                "namespace": "nginx-ingress"
                                            }
                                        ]
                                    }
                                ],
                                "externalObjects": {
                                    "subject": [
                                        {
                                            "kind": "ServiceAccount",
                                            "name": "nginx-ingress",
                                            "namespace": "nginx-ingress"
                                        }
                                    ]
                                }
                            }
                        }
                    ]
                }
            ],
            "remediation": "You should apply least privilege principle. Make sure cluster admin permissions are granted only when it is absolutely necessary. Don't use subjects with such high permissions for daily operations.",
            "description": "Attackers who have cluster admin permissions (can perform any action on any resource), can take advantage of their privileges for malicious activities. This control determines which subjects have cluster admin permissions.",
            "score": 95
        },
        {
            "guid": "",
            "attributes": {
                "armoBuiltin": true,
                "microsoftMitreColumns": [
                    "Execution"
                ]
            },
            "id": "C-0002",
            "controlID": "C-0002",
            "name": "Exec into container",
            "ruleReports": [
                {
                    "name": "exec-into-container",
                    "remediation": "",
                    "ruleStatus": {
                        "status": "success",
                        "message": ""
                    },
                    "ruleResponses": [
                        {
                            "alertMessage": "the following ServiceAccount: argocd-application-controller, can exec into  containers",
                            "ruleStatus": "",
                            "packagename": "armo_builtins",
                            "alertScore": 9,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "rbac.authorization.k8s.io/v1",
                                        "kind": "ClusterRole",
                                        "metadata": {
                                            "annotations": {
                                                "kubectl.kubernetes.io/last-applied-configuration": "{\"apiVersion\":\"rbac.authorization.k8s.io/v1\",\"kind\":\"ClusterRole\",\"metadata\":{\"annotations\":{},\"labels\":{\"app.kubernetes.io/component\":\"application-controller\",\"app.kubernetes.io/name\":\"argocd-application-controller\",\"app.kubernetes.io/part-of\":\"argocd\"},\"name\":\"argocd-application-controller\"},\"rules\":[{\"apiGroups\":[\"*\"],\"resources\":[\"*\"],\"verbs\":[\"*\"]},{\"nonResourceURLs\":[\"*\"],\"verbs\":[\"*\"]}]}\n"
                                            },
                                            "creationTimestamp": "2021-08-22T11:50:37Z",
                                            "labels": {
                                                "app.kubernetes.io/component": "application-controller",
                                                "app.kubernetes.io/name": "argocd-application-controller",
                                                "app.kubernetes.io/part-of": "argocd"
                                            },
                                            "name": "argocd-application-controller",
                                            "resourceVersion": "1427392",
                                            "selfLink": "/apis/rbac.authorization.k8s.io/v1/clusterroles/argocd-application-controller",
                                            "uid": "947b1d1e-8fae-4f64-8264-452275d0d2a3"
                                        },
                                        "rules": [
                                            {
                                                "apiGroups": [
                                                    "*"
                                                ],
                                                "resources": [
                                                    "*"
                                                ],
                                                "verbs": [
                                                    "*"
                                                ]
                                            },
                                            {
                                                "nonResourceURLs": [
                                                    "*"
                                                ],
                                                "verbs": [
                                                    "*"
                                                ]
                                            }
                                        ]
                                    },
                                    {
                                        "apiVersion": "rbac.authorization.k8s.io/v1",
                                        "kind": "ClusterRoleBinding",
                                        "metadata": {
                                            "annotations": {
                                                "kubectl.kubernetes.io/last-applied-configuration": "{\"apiVersion\":\"rbac.authorization.k8s.io/v1\",\"kind\":\"ClusterRoleBinding\",\"metadata\":{\"annotations\":{},\"labels\":{\"app.kubernetes.io/component\":\"application-controller\",\"app.kubernetes.io/name\":\"argocd-application-controller\",\"app.kubernetes.io/part-of\":\"argocd\"},\"name\":\"argocd-application-controller\"},\"roleRef\":{\"apiGroup\":\"rbac.authorization.k8s.io\",\"kind\":\"ClusterRole\",\"name\":\"argocd-application-controller\"},\"subjects\":[{\"kind\":\"ServiceAccount\",\"name\":\"argocd-application-controller\",\"namespace\":\"argocd\"}]}\n"
                                            },
                                            "creationTimestamp": "2021-08-22T11:50:37Z",
                                            "labels": {
                                                "app.kubernetes.io/component": "application-controller",
                                                "app.kubernetes.io/name": "argocd-application-controller",
                                                "app.kubernetes.io/part-of": "argocd"
                                            },
                                            "name": "argocd-application-controller",
                                            "resourceVersion": "1427398",
                                            "selfLink": "/apis/rbac.authorization.k8s.io/v1/clusterrolebindings/argocd-application-controller",
                                            "uid": "12055d3d-8f87-4051-bd24-334ba3c409b9"
                                        },
                                        "roleRef": {
                                            "apiGroup": "rbac.authorization.k8s.io",
                                            "kind": "ClusterRole",
                                            "name": "argocd-application-controller"
                                        },
                                        "subjects": [
                                            {
                                                "kind": "ServiceAccount",
                                                "name": "argocd-application-controller",
                                                "namespace": "argocd"
                                            }
                                        ]
                                    }
                                ],
                                "externalObjects": {
                                    "subject": [
                                        {
                                            "kind": "ServiceAccount",
                                            "name": "argocd-application-controller",
                                            "namespace": "argocd"
                                        }
                                    ]
                                }
                            }
                        },
                        {
                            "alertMessage": "the following Group: system:masters, can exec into  containers",
                            "ruleStatus": "",
                            "packagename": "armo_builtins",
                            "alertScore": 9,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "rbac.authorization.k8s.io/v1",
                                        "kind": "ClusterRole",
                                        "metadata": {
                                            "annotations": {
                                                "rbac.authorization.kubernetes.io/autoupdate": "true"
                                            },
                                            "creationTimestamp": "2021-06-20T09:07:04Z",
                                            "labels": {
                                                "kubernetes.io/bootstrapping": "rbac-defaults"
                                            },
                                            "name": "cluster-admin",
                                            "resourceVersion": "43",
                                            "selfLink": "/apis/rbac.authorization.k8s.io/v1/clusterroles/cluster-admin",
                                            "uid": "612c8a97-ef32-4995-8b1d-ebf60e74d514"
                                        },
                                        "rules": [
                                            {
                                                "apiGroups": [
                                                    "*"
                                                ],
                                                "resources": [
                                                    "*"
                                                ],
                                                "verbs": [
                                                    "*"
                                                ]
                                            },
                                            {
                                                "nonResourceURLs": [
                                                    "*"
                                                ],
                                                "verbs": [
                                                    "*"
                                                ]
                                            }
                                        ]
                                    },
                                    {
                                        "apiVersion": "rbac.authorization.k8s.io/v1",
                                        "kind": "ClusterRoleBinding",
                                        "metadata": {
                                            "annotations": {
                                                "rbac.authorization.kubernetes.io/autoupdate": "true"
                                            },
                                            "creationTimestamp": "2021-06-20T09:07:04Z",
                                            "labels": {
                                                "kubernetes.io/bootstrapping": "rbac-defaults"
                                            },
                                            "name": "cluster-admin",
                                            "resourceVersion": "96",
                                            "selfLink": "/apis/rbac.authorization.k8s.io/v1/clusterrolebindings/cluster-admin",
                                            "uid": "a594dc87-99a0-4154-9844-d7b5f6e1919c"
                                        },
                                        "roleRef": {
                                            "apiGroup": "rbac.authorization.k8s.io",
                                            "kind": "ClusterRole",
                                            "name": "cluster-admin"
                                        },
                                        "subjects": [
                                            {
                                                "apiGroup": "rbac.authorization.k8s.io",
                                                "kind": "Group",
                                                "name": "system:masters"
                                            }
                                        ]
                                    }
                                ],
                                "externalObjects": {
                                    "subject": [
                                        {
                                            "apiGroup": "rbac.authorization.k8s.io",
                                            "kind": "Group",
                                            "name": "system:masters"
                                        }
                                    ]
                                }
                            }
                        },
                        {
                            "alertMessage": "the following ServiceAccount: nginx-ingress, can exec into  containers",
                            "ruleStatus": "",
                            "packagename": "armo_builtins",
                            "alertScore": 9,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "rbac.authorization.k8s.io/v1",
                                        "kind": "ClusterRole",
                                        "metadata": {
                                            "annotations": {
                                                "kubectl.kubernetes.io/last-applied-configuration": "{\"apiVersion\":\"rbac.authorization.k8s.io/v1\",\"kind\":\"ClusterRole\",\"metadata\":{\"annotations\":{},\"labels\":{\"app\":\"nginx-ingress-roles\"},\"name\":\"nginx-ingress-roles\"},\"rules\":[{\"apiGroups\":[\"*\"],\"resources\":[\"*\"],\"verbs\":[\"*\"]}]}\n"
                                            },
                                            "creationTimestamp": "2021-08-18T14:58:32Z",
                                            "labels": {
                                                "app": "nginx-ingress-roles"
                                            },
                                            "name": "nginx-ingress-roles",
                                            "resourceVersion": "1355206",
                                            "selfLink": "/apis/rbac.authorization.k8s.io/v1/clusterroles/nginx-ingress-roles",
                                            "uid": "e262c786-9138-4db2-86be-1f6bd50f93bd"
                                        },
                                        "rules": [
                                            {
                                                "apiGroups": [
                                                    "*"
                                                ],
                                                "resources": [
                                                    "*"
                                                ],
                                                "verbs": [
                                                    "*"
                                                ]
                                            }
                                        ]
                                    },
                                    {
                                        "apiVersion": "rbac.authorization.k8s.io/v1",
                                        "kind": "ClusterRoleBinding",
                                        "metadata": {
                                            "annotations": {
                                                "kubectl.kubernetes.io/last-applied-configuration": "{\"apiVersion\":\"rbac.authorization.k8s.io/v1\",\"kind\":\"ClusterRoleBinding\",\"metadata\":{\"annotations\":{},\"labels\":{\"app\":\"nginx-ingress-roles-binding\"},\"name\":\"nginx-ingress-roles-binding\"},\"roleRef\":{\"apiGroup\":\"rbac.authorization.k8s.io\",\"kind\":\"ClusterRole\",\"name\":\"nginx-ingress-roles\"},\"subjects\":[{\"kind\":\"ServiceAccount\",\"name\":\"nginx-ingress\",\"namespace\":\"nginx-ingress\"}]}\n"
                                            },
                                            "creationTimestamp": "2021-08-18T14:58:57Z",
                                            "labels": {
                                                "app": "nginx-ingress-roles-binding"
                                            },
                                            "name": "nginx-ingress-roles-binding",
                                            "resourceVersion": "1355223",
                                            "selfLink": "/apis/rbac.authorization.k8s.io/v1/clusterrolebindings/nginx-ingress-roles-binding",
                                            "uid": "38821ce8-ea8f-4081-ba6c-5769a2382447"
                                        },
                                        "roleRef": {
                                            "apiGroup": "rbac.authorization.k8s.io",
                                            "kind": "ClusterRole",
                                            "name": "nginx-ingress-roles"
                                        },
                                        "subjects": [
                                            {
                                                "kind": "ServiceAccount",
                                                "name": "nginx-ingress",
                                                "namespace": "nginx-ingress"
                                            }
                                        ]
                                    }
                                ],
                                "externalObjects": {
                                    "subject": [
                                        {
                                            "kind": "ServiceAccount",
                                            "name": "nginx-ingress",
                                            "namespace": "nginx-ingress"
                                        }
                                    ]
                                }
                            }
                        }
                    ]
                }
            ],
            "remediation": "It is recommended to prohibit “kubectl exec” command in production environments. It is also recommended not to use subjects with this permission for daily cluster operations.",
            "description": "Attackers with relevant permissions can run malicious commands in the context of legitimate containers in the cluster using “kubectl exec” command. This control determines which subjects have permissions to use this command.",
            "score": 95
        },
        {
            "guid": "",
            "attributes": {
                "armoBuiltin": true
            },
            "id": "C-0028",
            "controlID": "C-0028",
            "name": "Dangerous capabilities",
            "ruleReports": [
                {
                    "name": "dangerous-capabilities",
                    "remediation": "",
                    "ruleStatus": {
                        "status": "success",
                        "message": ""
                    },
                    "ruleResponses": null
                }
            ],
            "remediation": "Check and remove all unnecessary capabilities from the POD security context of the containers and use the exception mechanism to remove warnings where these capabilities are necessary.",
            "description": "Giving dangerous and unnecessary LINUX capabilities to a container can increase the impact of the container compromise. This control identifies all the PODs with dangerous capabilities such as SYS_ADMIN and others.",
            "score": 100
        },
        {
            "guid": "",
            "attributes": {
                "armoBuiltin": true
            },
            "id": "C-0046",
            "controlID": "C-0046",
            "name": "Insecure capabilities",
            "ruleReports": [
                {
                    "name": "insecure-capabilities",
                    "remediation": "",
                    "ruleStatus": {
                        "status": "success",
                        "message": ""
                    },
                    "ruleResponses": [
                        {
                            "alertMessage": "container: coredns in workload: coredns  have dangerous capabilities",
                            "ruleStatus": "",
                            "packagename": "armo_builtins",
                            "alertScore": 7,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "apps/v1",
                                        "kind": "Deployment",
                                        "metadata": {
                                            "annotations": {
                                                "deployment.kubernetes.io/revision": "1"
                                            },
                                            "creationTimestamp": "2021-06-20T09:07:08Z",
                                            "generation": 2,
                                            "labels": {
                                                "k8s-app": "kube-dns"
                                            },
                                            "name": "coredns",
                                            "namespace": "kube-system",
                                            "resourceVersion": "2110974",
                                            "selfLink": "/apis/apps/v1/namespaces/kube-system/deployments/coredns",
                                            "uid": "fdab6339-f481-4fcb-aa9d-d6fa9260184b"
                                        },
                                        "spec": {
                                            "progressDeadlineSeconds": 600,
                                            "replicas": 1,
                                            "revisionHistoryLimit": 10,
                                            "selector": {
                                                "matchLabels": {
                                                    "k8s-app": "kube-dns"
                                                }
                                            },
                                            "strategy": {
                                                "rollingUpdate": {
                                                    "maxSurge": "25%",
                                                    "maxUnavailable": 1
                                                },
                                                "type": "RollingUpdate"
                                            },
                                            "template": {
                                                "metadata": {
                                                    "creationTimestamp": null,
                                                    "labels": {
                                                        "k8s-app": "kube-dns"
                                                    }
                                                },
                                                "spec": {
                                                    "containers": [
                                                        {
                                                            "args": [
                                                                "-conf",
                                                                "/etc/coredns/Corefile"
                                                            ],
                                                            "image": "k8s.gcr.io/coredns:1.6.2",
                                                            "imagePullPolicy": "IfNotPresent",
                                                            "livenessProbe": {
                                                                "failureThreshold": 5,
                                                                "httpGet": {
                                                                    "path": "/health",
                                                                    "port": 8080,
                                                                    "scheme": "HTTP"
                                                                },
                                                                "initialDelaySeconds": 60,
                                                                "periodSeconds": 10,
                                                                "successThreshold": 1,
                                                                "timeoutSeconds": 5
                                                            },
                                                            "name": "coredns",
                                                            "ports": [
                                                                {
                                                                    "containerPort": 53,
                                                                    "name": "dns",
                                                                    "protocol": "UDP"
                                                                },
                                                                {
                                                                    "containerPort": 53,
                                                                    "name": "dns-tcp",
                                                                    "protocol": "TCP"
                                                                },
                                                                {
                                                                    "containerPort": 9153,
                                                                    "name": "metrics",
                                                                    "protocol": "TCP"
                                                                }
                                                            ],
                                                            "readinessProbe": {
                                                                "failureThreshold": 3,
                                                                "httpGet": {
                                                                    "path": "/ready",
                                                                    "port": 8181,
                                                                    "scheme": "HTTP"
                                                                },
                                                                "periodSeconds": 10,
                                                                "successThreshold": 1,
                                                                "timeoutSeconds": 1
                                                            },
                                                            "resources": {
                                                                "limits": {
                                                                    "memory": "170Mi"
                                                                },
                                                                "requests": {
                                                                    "cpu": "100m",
                                                                    "memory": "70Mi"
                                                                }
                                                            },
                                                            "securityContext": {
                                                                "allowPrivilegeEscalation": false,
                                                                "capabilities": {
                                                                    "add": [
                                                                        "NET_BIND_SERVICE"
                                                                    ],
                                                                    "drop": [
                                                                        "all"
                                                                    ]
                                                                },
                                                                "readOnlyRootFilesystem": true
                                                            },
                                                            "terminationMessagePath": "/dev/termination-log",
                                                            "terminationMessagePolicy": "File",
                                                            "volumeMounts": [
                                                                {
                                                                    "mountPath": "/etc/coredns",
                                                                    "name": "config-volume",
                                                                    "readOnly": true
                                                                }
                                                            ]
                                                        }
                                                    ],
                                                    "dnsPolicy": "Default",
                                                    "nodeSelector": {
                                                        "beta.kubernetes.io/os": "linux"
                                                    },
                                                    "priorityClassName": "system-cluster-critical",
                                                    "restartPolicy": "Always",
                                                    "schedulerName": "default-scheduler",
                                                    "securityContext": {},
                                                    "serviceAccount": "coredns",
                                                    "serviceAccountName": "coredns",
                                                    "terminationGracePeriodSeconds": 30,
                                                    "tolerations": [
                                                        {
                                                            "key": "CriticalAddonsOnly",
                                                            "operator": "Exists"
                                                        },
                                                        {
                                                            "effect": "NoSchedule",
                                                            "key": "node-role.kubernetes.io/master"
                                                        }
                                                    ],
                                                    "volumes": [
                                                        {
                                                            "configMap": {
                                                                "defaultMode": 420,
                                                                "items": [
                                                                    {
                                                                        "key": "Corefile",
                                                                        "path": "Corefile"
                                                                    }
                                                                ],
                                                                "name": "coredns"
                                                            },
                                                            "name": "config-volume"
                                                        }
                                                    ]
                                                }
                                            }
                                        },
                                        "status": {
                                            "availableReplicas": 1,
                                            "conditions": [
                                                {
                                                    "lastTransitionTime": "2021-06-20T09:07:22Z",
                                                    "lastUpdateTime": "2021-06-20T09:07:22Z",
                                                    "message": "Deployment has minimum availability.",
                                                    "reason": "MinimumReplicasAvailable",
                                                    "status": "True",
                                                    "type": "Available"
                                                },
                                                {
                                                    "lastTransitionTime": "2021-06-20T09:07:22Z",
                                                    "lastUpdateTime": "2021-06-20T09:07:35Z",
                                                    "message": "ReplicaSet \"coredns-5644d7b6d9\" has successfully progressed.",
                                                    "reason": "NewReplicaSetAvailable",
                                                    "status": "True",
                                                    "type": "Progressing"
                                                }
                                            ],
                                            "observedGeneration": 2,
                                            "readyReplicas": 1,
                                            "replicas": 1,
                                            "updatedReplicas": 1
                                        }
                                    }
                                ]
                            }
                        }
                    ]
                }
            ],
            "remediation": "Remove all insecure capabilities which aren’t necessary for the container.",
            "description": "Giving insecure or excsessive capabilities to a container can increase the impact of the container compromise. This control identifies all the PODs with dangerous capabilities (see documentation pages for details).",
            "score": 85
        },
        {
            "guid": "",
            "attributes": {
                "armoBuiltin": true
            },
            "id": "C-0055",
            "controlID": "C-0055",
            "name": "Linux hardening",
            "ruleReports": [
                {
                    "name": "linux-hardening",
                    "remediation": "",
                    "ruleStatus": {
                        "status": "success",
                        "message": ""
                    },
                    "ruleResponses": [
                        {
                            "alertMessage": "Pod: etcd-david-virtualbox does not define any linux security hardening",
                            "ruleStatus": "",
                            "packagename": "armo_builtins",
                            "alertScore": 7,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "Pod",
                                        "metadata": {
                                            "annotations": {
                                                "kubernetes.io/config.hash": "e0fcc6e4323055b5880f8aac4c950836",
                                                "kubernetes.io/config.mirror": "e0fcc6e4323055b5880f8aac4c950836",
                                                "kubernetes.io/config.seen": "2021-06-20T12:06:52.495386281+03:00",
                                                "kubernetes.io/config.source": "file"
                                            },
                                            "creationTimestamp": "2021-06-20T09:08:22Z",
                                            "labels": {
                                                "component": "etcd",
                                                "tier": "control-plane"
                                            },
                                            "name": "etcd-david-virtualbox",
                                            "namespace": "kube-system",
                                            "resourceVersion": "2110909",
                                            "selfLink": "/api/v1/namespaces/kube-system/pods/etcd-david-virtualbox",
                                            "uid": "154e7f87-907f-4edb-a73c-26e965d4fe02"
                                        },
                                        "spec": {
                                            "containers": [
                                                {
                                                    "command": [
                                                        "etcd",
                                                        "--advertise-client-urls=https://10.0.2.15:2379",
                                                        "--cert-file=/var/lib/minikube/certs/etcd/server.crt",
                                                        "--client-cert-auth=true",
                                                        "--data-dir=/var/lib/minikube/etcd",
                                                        "--initial-advertise-peer-urls=https://10.0.2.15:2380",
                                                        "--initial-cluster=david-virtualbox=https://10.0.2.15:2380",
                                                        "--key-file=/var/lib/minikube/certs/etcd/server.key",
                                                        "--listen-client-urls=https://127.0.0.1:2379,https://10.0.2.15:2379",
                                                        "--listen-metrics-urls=http://127.0.0.1:2381,http://10.0.2.15:2381",
                                                        "--listen-peer-urls=https://10.0.2.15:2380",
                                                        "--name=david-virtualbox",
                                                        "--peer-cert-file=/var/lib/minikube/certs/etcd/peer.crt",
                                                        "--peer-client-cert-auth=true",
                                                        "--peer-key-file=/var/lib/minikube/certs/etcd/peer.key",
                                                        "--peer-trusted-ca-file=/var/lib/minikube/certs/etcd/ca.crt",
                                                        "--snapshot-count=10000",
                                                        "--trusted-ca-file=/var/lib/minikube/certs/etcd/ca.crt"
                                                    ],
                                                    "image": "k8s.gcr.io/etcd:3.3.15-0",
                                                    "imagePullPolicy": "IfNotPresent",
                                                    "livenessProbe": {
                                                        "failureThreshold": 8,
                                                        "httpGet": {
                                                            "host": "127.0.0.1",
                                                            "path": "/health",
                                                            "port": 2381,
                                                            "scheme": "HTTP"
                                                        },
                                                        "initialDelaySeconds": 15,
                                                        "periodSeconds": 10,
                                                        "successThreshold": 1,
                                                        "timeoutSeconds": 15
                                                    },
                                                    "name": "etcd",
                                                    "resources": {},
                                                    "terminationMessagePath": "/dev/termination-log",
                                                    "terminationMessagePolicy": "File",
                                                    "volumeMounts": [
                                                        {
                                                            "mountPath": "/var/lib/minikube/etcd",
                                                            "name": "etcd-data"
                                                        },
                                                        {
                                                            "mountPath": "/var/lib/minikube/certs/etcd",
                                                            "name": "etcd-certs"
                                                        }
                                                    ]
                                                }
                                            ],
                                            "dnsPolicy": "ClusterFirst",
                                            "enableServiceLinks": true,
                                            "hostNetwork": true,
                                            "nodeName": "david-virtualbox",
                                            "priority": 2000000000,
                                            "priorityClassName": "system-cluster-critical",
                                            "restartPolicy": "Always",
                                            "schedulerName": "default-scheduler",
                                            "securityContext": {},
                                            "terminationGracePeriodSeconds": 30,
                                            "tolerations": [
                                                {
                                                    "effect": "NoExecute",
                                                    "operator": "Exists"
                                                }
                                            ],
                                            "volumes": [
                                                {
                                                    "hostPath": {
                                                        "path": "/var/lib/minikube/certs/etcd",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "etcd-certs"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/var/lib/minikube/etcd",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "etcd-data"
                                                }
                                            ]
                                        },
                                        "status": {
                                            "conditions": [
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-03T06:28:57Z",
                                                    "status": "True",
                                                    "type": "Initialized"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-06T05:23:22Z",
                                                    "status": "True",
                                                    "type": "Ready"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-06T05:23:22Z",
                                                    "status": "True",
                                                    "type": "ContainersReady"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-03T06:28:57Z",
                                                    "status": "True",
                                                    "type": "PodScheduled"
                                                }
                                            ],
                                            "containerStatuses": [
                                                {
                                                    "containerID": "docker://0cc62102444cc74e3dddb2b6cd7550036feaa0e61f4297c85fc72e3be2905ec9",
                                                    "image": "k8s.gcr.io/etcd:3.3.15-0",
                                                    "imageID": "docker-pullable://k8s.gcr.io/etcd@sha256:12c2c5e5731c3bcd56e6f1c05c0f9198b6f06793fa7fca2fb43aab9622dc4afa",
                                                    "lastState": {
                                                        "terminated": {
                                                            "containerID": "docker://61de0abaa35617dafed9274c5b03f99f269412bdea2f15d903b4293d620f4b9e",
                                                            "exitCode": 255,
                                                            "finishedAt": "2021-10-14T05:35:02Z",
                                                            "reason": "Error",
                                                            "startedAt": "2021-10-12T10:28:38Z"
                                                        }
                                                    },
                                                    "name": "etcd",
                                                    "ready": true,
                                                    "restartCount": 51,
                                                    "started": true,
                                                    "state": {
                                                        "running": {
                                                            "startedAt": "2021-10-14T05:35:12Z"
                                                        }
                                                    }
                                                }
                                            ],
                                            "hostIP": "10.0.2.15",
                                            "phase": "Running",
                                            "podIP": "10.0.2.15",
                                            "podIPs": [
                                                {
                                                    "ip": "10.0.2.15"
                                                }
                                            ],
                                            "qosClass": "BestEffort",
                                            "startTime": "2021-10-03T06:28:57Z"
                                        }
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "Pod: kube-apiserver-david-virtualbox does not define any linux security hardening",
                            "ruleStatus": "",
                            "packagename": "armo_builtins",
                            "alertScore": 7,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "Pod",
                                        "metadata": {
                                            "annotations": {
                                                "kubernetes.io/config.hash": "668e9396bc7b8c495d39f4a3bc479397",
                                                "kubernetes.io/config.mirror": "668e9396bc7b8c495d39f4a3bc479397",
                                                "kubernetes.io/config.seen": "2021-06-20T12:06:52.495387809+03:00",
                                                "kubernetes.io/config.source": "file"
                                            },
                                            "creationTimestamp": "2021-06-20T09:08:05Z",
                                            "labels": {
                                                "component": "kube-apiserver",
                                                "tier": "control-plane"
                                            },
                                            "name": "kube-apiserver-david-virtualbox",
                                            "namespace": "kube-system",
                                            "resourceVersion": "2110920",
                                            "selfLink": "/api/v1/namespaces/kube-system/pods/kube-apiserver-david-virtualbox",
                                            "uid": "327cbf13-97d6-42a3-8469-5acfdbe1be09"
                                        },
                                        "spec": {
                                            "containers": [
                                                {
                                                    "command": [
                                                        "kube-apiserver",
                                                        "--advertise-address=10.0.2.15",
                                                        "--allow-privileged=true",
                                                        "--authorization-mode=Node,RBAC",
                                                        "--client-ca-file=/var/lib/minikube/certs/ca.crt",
                                                        "--enable-admission-plugins=NamespaceLifecycle,LimitRanger,ServiceAccount,DefaultStorageClass,DefaultTolerationSeconds,NodeRestriction,MutatingAdmissionWebhook,ValidatingAdmissionWebhook,ResourceQuota",
                                                        "--enable-bootstrap-token-auth=true",
                                                        "--etcd-cafile=/var/lib/minikube/certs/etcd/ca.crt",
                                                        "--etcd-certfile=/var/lib/minikube/certs/apiserver-etcd-client.crt",
                                                        "--etcd-keyfile=/var/lib/minikube/certs/apiserver-etcd-client.key",
                                                        "--etcd-servers=https://127.0.0.1:2379",
                                                        "--insecure-port=0",
                                                        "--kubelet-client-certificate=/var/lib/minikube/certs/apiserver-kubelet-client.crt",
                                                        "--kubelet-client-key=/var/lib/minikube/certs/apiserver-kubelet-client.key",
                                                        "--kubelet-preferred-address-types=InternalIP,ExternalIP,Hostname",
                                                        "--proxy-client-cert-file=/var/lib/minikube/certs/front-proxy-client.crt",
                                                        "--proxy-client-key-file=/var/lib/minikube/certs/front-proxy-client.key",
                                                        "--requestheader-allowed-names=front-proxy-client",
                                                        "--requestheader-client-ca-file=/var/lib/minikube/certs/front-proxy-ca.crt",
                                                        "--requestheader-extra-headers-prefix=X-Remote-Extra-",
                                                        "--requestheader-group-headers=X-Remote-Group",
                                                        "--requestheader-username-headers=X-Remote-User",
                                                        "--secure-port=8443",
                                                        "--service-account-key-file=/var/lib/minikube/certs/sa.pub",
                                                        "--service-cluster-ip-range=10.96.0.0/12",
                                                        "--tls-cert-file=/var/lib/minikube/certs/apiserver.crt",
                                                        "--tls-private-key-file=/var/lib/minikube/certs/apiserver.key"
                                                    ],
                                                    "image": "k8s.gcr.io/kube-apiserver:v1.16.0",
                                                    "imagePullPolicy": "IfNotPresent",
                                                    "livenessProbe": {
                                                        "failureThreshold": 8,
                                                        "httpGet": {
                                                            "host": "10.0.2.15",
                                                            "path": "/healthz",
                                                            "port": 8443,
                                                            "scheme": "HTTPS"
                                                        },
                                                        "initialDelaySeconds": 15,
                                                        "periodSeconds": 10,
                                                        "successThreshold": 1,
                                                        "timeoutSeconds": 15
                                                    },
                                                    "name": "kube-apiserver",
                                                    "resources": {
                                                        "requests": {
                                                            "cpu": "250m"
                                                        }
                                                    },
                                                    "terminationMessagePath": "/dev/termination-log",
                                                    "terminationMessagePolicy": "File",
                                                    "volumeMounts": [
                                                        {
                                                            "mountPath": "/etc/ssl/certs",
                                                            "name": "ca-certs",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/etc/ca-certificates",
                                                            "name": "etc-ca-certificates",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/etc/pki",
                                                            "name": "etc-pki",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/var/lib/minikube/certs",
                                                            "name": "k8s-certs",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/usr/local/share/ca-certificates",
                                                            "name": "usr-local-share-ca-certificates",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/usr/share/ca-certificates",
                                                            "name": "usr-share-ca-certificates",
                                                            "readOnly": true
                                                        }
                                                    ]
                                                }
                                            ],
                                            "dnsPolicy": "ClusterFirst",
                                            "enableServiceLinks": true,
                                            "hostNetwork": true,
                                            "nodeName": "david-virtualbox",
                                            "priority": 2000000000,
                                            "priorityClassName": "system-cluster-critical",
                                            "restartPolicy": "Always",
                                            "schedulerName": "default-scheduler",
                                            "securityContext": {},
                                            "terminationGracePeriodSeconds": 30,
                                            "tolerations": [
                                                {
                                                    "effect": "NoExecute",
                                                    "operator": "Exists"
                                                }
                                            ],
                                            "volumes": [
                                                {
                                                    "hostPath": {
                                                        "path": "/etc/ssl/certs",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "ca-certs"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/etc/ca-certificates",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "etc-ca-certificates"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/etc/pki",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "etc-pki"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/var/lib/minikube/certs",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "k8s-certs"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/usr/local/share/ca-certificates",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "usr-local-share-ca-certificates"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/usr/share/ca-certificates",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "usr-share-ca-certificates"
                                                }
                                            ]
                                        },
                                        "status": {
                                            "conditions": [
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-03T06:28:57Z",
                                                    "status": "True",
                                                    "type": "Initialized"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-06T05:23:25Z",
                                                    "status": "True",
                                                    "type": "Ready"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-06T05:23:25Z",
                                                    "status": "True",
                                                    "type": "ContainersReady"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-03T06:28:57Z",
                                                    "status": "True",
                                                    "type": "PodScheduled"
                                                }
                                            ],
                                            "containerStatuses": [
                                                {
                                                    "containerID": "docker://c6d5933ac2752c18942e488d2d162fcc3141d1ae131be2abc3a50d2cb91e4e22",
                                                    "image": "k8s.gcr.io/kube-apiserver:v1.16.0",
                                                    "imageID": "docker-pullable://k8s.gcr.io/kube-apiserver@sha256:f4168527c91289da2708f62ae729fdde5fb484167dd05ffbb7ab666f60de96cd",
                                                    "lastState": {
                                                        "terminated": {
                                                            "containerID": "docker://3714c8295fe8b0ff97082305a44f08764e06d1a599ab3ee3ca6db15ab498a276",
                                                            "exitCode": 255,
                                                            "finishedAt": "2021-10-14T05:35:02Z",
                                                            "reason": "Error",
                                                            "startedAt": "2021-10-12T10:28:38Z"
                                                        }
                                                    },
                                                    "name": "kube-apiserver",
                                                    "ready": true,
                                                    "restartCount": 51,
                                                    "started": true,
                                                    "state": {
                                                        "running": {
                                                            "startedAt": "2021-10-14T05:35:12Z"
                                                        }
                                                    }
                                                }
                                            ],
                                            "hostIP": "10.0.2.15",
                                            "phase": "Running",
                                            "podIP": "10.0.2.15",
                                            "podIPs": [
                                                {
                                                    "ip": "10.0.2.15"
                                                }
                                            ],
                                            "qosClass": "Burstable",
                                            "startTime": "2021-10-03T06:28:57Z"
                                        }
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "Pod: kube-controller-manager-david-virtualbox does not define any linux security hardening",
                            "ruleStatus": "",
                            "packagename": "armo_builtins",
                            "alertScore": 7,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "Pod",
                                        "metadata": {
                                            "annotations": {
                                                "kubernetes.io/config.hash": "a16b2d5766eae37796e4a8ed7f8ce12a",
                                                "kubernetes.io/config.mirror": "a16b2d5766eae37796e4a8ed7f8ce12a",
                                                "kubernetes.io/config.seen": "2021-06-20T12:06:52.495389283+03:00",
                                                "kubernetes.io/config.source": "file"
                                            },
                                            "creationTimestamp": "2021-06-20T09:08:00Z",
                                            "labels": {
                                                "component": "kube-controller-manager",
                                                "tier": "control-plane"
                                            },
                                            "name": "kube-controller-manager-david-virtualbox",
                                            "namespace": "kube-system",
                                            "resourceVersion": "2110899",
                                            "selfLink": "/api/v1/namespaces/kube-system/pods/kube-controller-manager-david-virtualbox",
                                            "uid": "6ca9d32c-21c3-4c0e-8087-5445c80a2bcc"
                                        },
                                        "spec": {
                                            "containers": [
                                                {
                                                    "command": [
                                                        "kube-controller-manager",
                                                        "--allocate-node-cidrs=true",
                                                        "--authentication-kubeconfig=/etc/kubernetes/controller-manager.conf",
                                                        "--authorization-kubeconfig=/etc/kubernetes/controller-manager.conf",
                                                        "--bind-address=127.0.0.1",
                                                        "--client-ca-file=/var/lib/minikube/certs/ca.crt",
                                                        "--cluster-cidr=10.244.0.0/16",
                                                        "--cluster-signing-cert-file=/var/lib/minikube/certs/ca.crt",
                                                        "--cluster-signing-key-file=/var/lib/minikube/certs/ca.key",
                                                        "--controllers=*,bootstrapsigner,tokencleaner",
                                                        "--kubeconfig=/etc/kubernetes/controller-manager.conf",
                                                        "--leader-elect=false",
                                                        "--node-cidr-mask-size=24",
                                                        "--requestheader-client-ca-file=/var/lib/minikube/certs/front-proxy-ca.crt",
                                                        "--root-ca-file=/var/lib/minikube/certs/ca.crt",
                                                        "--service-account-private-key-file=/var/lib/minikube/certs/sa.key",
                                                        "--service-cluster-ip-range=10.96.0.0/12",
                                                        "--use-service-account-credentials=true"
                                                    ],
                                                    "image": "k8s.gcr.io/kube-controller-manager:v1.16.0",
                                                    "imagePullPolicy": "IfNotPresent",
                                                    "livenessProbe": {
                                                        "failureThreshold": 8,
                                                        "httpGet": {
                                                            "host": "127.0.0.1",
                                                            "path": "/healthz",
                                                            "port": 10252,
                                                            "scheme": "HTTP"
                                                        },
                                                        "initialDelaySeconds": 15,
                                                        "periodSeconds": 10,
                                                        "successThreshold": 1,
                                                        "timeoutSeconds": 15
                                                    },
                                                    "name": "kube-controller-manager",
                                                    "resources": {
                                                        "requests": {
                                                            "cpu": "200m"
                                                        }
                                                    },
                                                    "terminationMessagePath": "/dev/termination-log",
                                                    "terminationMessagePolicy": "File",
                                                    "volumeMounts": [
                                                        {
                                                            "mountPath": "/etc/ssl/certs",
                                                            "name": "ca-certs",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/etc/ca-certificates",
                                                            "name": "etc-ca-certificates",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/etc/pki",
                                                            "name": "etc-pki",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/usr/libexec/kubernetes/kubelet-plugins/volume/exec",
                                                            "name": "flexvolume-dir"
                                                        },
                                                        {
                                                            "mountPath": "/var/lib/minikube/certs",
                                                            "name": "k8s-certs",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/etc/kubernetes/controller-manager.conf",
                                                            "name": "kubeconfig",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/usr/local/share/ca-certificates",
                                                            "name": "usr-local-share-ca-certificates",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/usr/share/ca-certificates",
                                                            "name": "usr-share-ca-certificates",
                                                            "readOnly": true
                                                        }
                                                    ]
                                                }
                                            ],
                                            "dnsPolicy": "ClusterFirst",
                                            "enableServiceLinks": true,
                                            "hostNetwork": true,
                                            "nodeName": "david-virtualbox",
                                            "priority": 2000000000,
                                            "priorityClassName": "system-cluster-critical",
                                            "restartPolicy": "Always",
                                            "schedulerName": "default-scheduler",
                                            "securityContext": {},
                                            "terminationGracePeriodSeconds": 30,
                                            "tolerations": [
                                                {
                                                    "effect": "NoExecute",
                                                    "operator": "Exists"
                                                }
                                            ],
                                            "volumes": [
                                                {
                                                    "hostPath": {
                                                        "path": "/etc/ssl/certs",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "ca-certs"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/etc/ca-certificates",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "etc-ca-certificates"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/etc/pki",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "etc-pki"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/usr/libexec/kubernetes/kubelet-plugins/volume/exec",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "flexvolume-dir"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/var/lib/minikube/certs",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "k8s-certs"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/etc/kubernetes/controller-manager.conf",
                                                        "type": "FileOrCreate"
                                                    },
                                                    "name": "kubeconfig"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/usr/local/share/ca-certificates",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "usr-local-share-ca-certificates"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/usr/share/ca-certificates",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "usr-share-ca-certificates"
                                                }
                                            ]
                                        },
                                        "status": {
                                            "conditions": [
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-14T05:35:11Z",
                                                    "status": "True",
                                                    "type": "Initialized"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-14T05:35:13Z",
                                                    "status": "True",
                                                    "type": "Ready"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-14T05:35:13Z",
                                                    "status": "True",
                                                    "type": "ContainersReady"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-14T05:35:11Z",
                                                    "status": "True",
                                                    "type": "PodScheduled"
                                                }
                                            ],
                                            "containerStatuses": [
                                                {
                                                    "containerID": "docker://eefb83fb1b81a497703013f609a54ffc03b57157395cb5f9a128dfeffea54b58",
                                                    "image": "k8s.gcr.io/kube-controller-manager:v1.16.0",
                                                    "imageID": "docker-pullable://k8s.gcr.io/kube-controller-manager@sha256:c156a05ee9d40e3ca2ebf9337f38a10558c1fc6c9124006f128a82e6c38cdf3e",
                                                    "lastState": {
                                                        "terminated": {
                                                            "containerID": "docker://81d8dca7f5b87e8b9bd8f515e9bcf86f4614306a9c729c30aad92560f520ed54",
                                                            "exitCode": 255,
                                                            "finishedAt": "2021-10-14T05:35:02Z",
                                                            "reason": "Error",
                                                            "startedAt": "2021-10-12T10:28:38Z"
                                                        }
                                                    },
                                                    "name": "kube-controller-manager",
                                                    "ready": true,
                                                    "restartCount": 55,
                                                    "started": true,
                                                    "state": {
                                                        "running": {
                                                            "startedAt": "2021-10-14T05:35:12Z"
                                                        }
                                                    }
                                                }
                                            ],
                                            "hostIP": "10.0.2.15",
                                            "phase": "Running",
                                            "podIP": "10.0.2.15",
                                            "podIPs": [
                                                {
                                                    "ip": "10.0.2.15"
                                                }
                                            ],
                                            "qosClass": "Burstable",
                                            "startTime": "2021-10-14T05:35:11Z"
                                        }
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "Pod: kube-scheduler-david-virtualbox does not define any linux security hardening",
                            "ruleStatus": "",
                            "packagename": "armo_builtins",
                            "alertScore": 7,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "Pod",
                                        "metadata": {
                                            "annotations": {
                                                "kubernetes.io/config.hash": "b3d303074fe0ca1d42a8bd9ed248df09",
                                                "kubernetes.io/config.mirror": "b3d303074fe0ca1d42a8bd9ed248df09",
                                                "kubernetes.io/config.seen": "2021-06-20T12:06:52.495381685+03:00",
                                                "kubernetes.io/config.source": "file"
                                            },
                                            "creationTimestamp": "2021-06-20T09:08:11Z",
                                            "labels": {
                                                "component": "kube-scheduler",
                                                "tier": "control-plane"
                                            },
                                            "name": "kube-scheduler-david-virtualbox",
                                            "namespace": "kube-system",
                                            "resourceVersion": "2110888",
                                            "selfLink": "/api/v1/namespaces/kube-system/pods/kube-scheduler-david-virtualbox",
                                            "uid": "226e285e-b2b3-423c-8bd4-04cfb775cbc6"
                                        },
                                        "spec": {
                                            "containers": [
                                                {
                                                    "command": [
                                                        "kube-scheduler",
                                                        "--authentication-kubeconfig=/etc/kubernetes/scheduler.conf",
                                                        "--authorization-kubeconfig=/etc/kubernetes/scheduler.conf",
                                                        "--bind-address=127.0.0.1",
                                                        "--kubeconfig=/etc/kubernetes/scheduler.conf",
                                                        "--leader-elect=false"
                                                    ],
                                                    "image": "k8s.gcr.io/kube-scheduler:v1.16.0",
                                                    "imagePullPolicy": "IfNotPresent",
                                                    "livenessProbe": {
                                                        "failureThreshold": 8,
                                                        "httpGet": {
                                                            "host": "127.0.0.1",
                                                            "path": "/healthz",
                                                            "port": 10251,
                                                            "scheme": "HTTP"
                                                        },
                                                        "initialDelaySeconds": 15,
                                                        "periodSeconds": 10,
                                                        "successThreshold": 1,
                                                        "timeoutSeconds": 15
                                                    },
                                                    "name": "kube-scheduler",
                                                    "resources": {
                                                        "requests": {
                                                            "cpu": "100m"
                                                        }
                                                    },
                                                    "terminationMessagePath": "/dev/termination-log",
                                                    "terminationMessagePolicy": "File",
                                                    "volumeMounts": [
                                                        {
                                                            "mountPath": "/etc/kubernetes/scheduler.conf",
                                                            "name": "kubeconfig",
                                                            "readOnly": true
                                                        }
                                                    ]
                                                }
                                            ],
                                            "dnsPolicy": "ClusterFirst",
                                            "enableServiceLinks": true,
                                            "hostNetwork": true,
                                            "nodeName": "david-virtualbox",
                                            "priority": 2000000000,
                                            "priorityClassName": "system-cluster-critical",
                                            "restartPolicy": "Always",
                                            "schedulerName": "default-scheduler",
                                            "securityContext": {},
                                            "terminationGracePeriodSeconds": 30,
                                            "tolerations": [
                                                {
                                                    "effect": "NoExecute",
                                                    "operator": "Exists"
                                                }
                                            ],
                                            "volumes": [
                                                {
                                                    "hostPath": {
                                                        "path": "/etc/kubernetes/scheduler.conf",
                                                        "type": "FileOrCreate"
                                                    },
                                                    "name": "kubeconfig"
                                                }
                                            ]
                                        },
                                        "status": {
                                            "conditions": [
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-14T05:35:11Z",
                                                    "status": "True",
                                                    "type": "Initialized"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-14T05:35:13Z",
                                                    "status": "True",
                                                    "type": "Ready"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-14T05:35:13Z",
                                                    "status": "True",
                                                    "type": "ContainersReady"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-14T05:35:11Z",
                                                    "status": "True",
                                                    "type": "PodScheduled"
                                                }
                                            ],
                                            "containerStatuses": [
                                                {
                                                    "containerID": "docker://53c67aa4e9ac29bf19ebaa451fe9de11d2bad3442dc6ece2dc3e03a0dc266526",
                                                    "image": "k8s.gcr.io/kube-scheduler:v1.16.0",
                                                    "imageID": "docker-pullable://k8s.gcr.io/kube-scheduler@sha256:094023ab9cd02059eb0295d234ff9ea321e0e22e4813986d7f1a1ac4dc1990d0",
                                                    "lastState": {
                                                        "terminated": {
                                                            "containerID": "docker://fbe8092ea4bbc12f8ed65da47d80b24528557bc81a7fc54948da26e76d704b4a",
                                                            "exitCode": 255,
                                                            "finishedAt": "2021-10-14T05:35:02Z",
                                                            "reason": "Error",
                                                            "startedAt": "2021-10-12T10:28:38Z"
                                                        }
                                                    },
                                                    "name": "kube-scheduler",
                                                    "ready": true,
                                                    "restartCount": 52,
                                                    "started": true,
                                                    "state": {
                                                        "running": {
                                                            "startedAt": "2021-10-14T05:35:12Z"
                                                        }
                                                    }
                                                }
                                            ],
                                            "hostIP": "10.0.2.15",
                                            "phase": "Running",
                                            "podIP": "10.0.2.15",
                                            "podIPs": [
                                                {
                                                    "ip": "10.0.2.15"
                                                }
                                            ],
                                            "qosClass": "Burstable",
                                            "startTime": "2021-10-14T05:35:11Z"
                                        }
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "Pod: storage-provisioner does not define any linux security hardening",
                            "ruleStatus": "",
                            "packagename": "armo_builtins",
                            "alertScore": 7,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "Pod",
                                        "metadata": {
                                            "annotations": {
                                                "kubectl.kubernetes.io/last-applied-configuration": "{\"apiVersion\":\"v1\",\"kind\":\"Pod\",\"metadata\":{\"annotations\":{},\"labels\":{\"addonmanager.kubernetes.io/mode\":\"Reconcile\",\"integration-test\":\"storage-provisioner\"},\"name\":\"storage-provisioner\",\"namespace\":\"kube-system\"},\"spec\":{\"containers\":[{\"command\":[\"/storage-provisioner\"],\"image\":\"gcr.io/k8s-minikube/storage-provisioner:v4\",\"imagePullPolicy\":\"IfNotPresent\",\"name\":\"storage-provisioner\",\"volumeMounts\":[{\"mountPath\":\"/tmp\",\"name\":\"tmp\"}]}],\"hostNetwork\":true,\"serviceAccountName\":\"storage-provisioner\",\"volumes\":[{\"hostPath\":{\"path\":\"/tmp\",\"type\":\"Directory\"},\"name\":\"tmp\"}]}}\n"
                                            },
                                            "creationTimestamp": "2021-06-20T09:07:09Z",
                                            "labels": {
                                                "addonmanager.kubernetes.io/mode": "Reconcile",
                                                "integration-test": "storage-provisioner"
                                            },
                                            "name": "storage-provisioner",
                                            "namespace": "kube-system",
                                            "resourceVersion": "2110982",
                                            "selfLink": "/api/v1/namespaces/kube-system/pods/storage-provisioner",
                                            "uid": "ea5dc2e2-4f7a-49f4-9e88-37e8e2d741a5"
                                        },
                                        "spec": {
                                            "containers": [
                                                {
                                                    "command": [
                                                        "/storage-provisioner"
                                                    ],
                                                    "image": "gcr.io/k8s-minikube/storage-provisioner:v4",
                                                    "imagePullPolicy": "IfNotPresent",
                                                    "name": "storage-provisioner",
                                                    "resources": {},
                                                    "terminationMessagePath": "/dev/termination-log",
                                                    "terminationMessagePolicy": "File",
                                                    "volumeMounts": [
                                                        {
                                                            "mountPath": "/tmp",
                                                            "name": "tmp"
                                                        },
                                                        {
                                                            "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount",
                                                            "name": "storage-provisioner-token-bbjlq",
                                                            "readOnly": true
                                                        }
                                                    ]
                                                }
                                            ],
                                            "dnsPolicy": "ClusterFirst",
                                            "enableServiceLinks": true,
                                            "hostNetwork": true,
                                            "nodeName": "david-virtualbox",
                                            "priority": 0,
                                            "restartPolicy": "Always",
                                            "schedulerName": "default-scheduler",
                                            "securityContext": {},
                                            "serviceAccount": "storage-provisioner",
                                            "serviceAccountName": "storage-provisioner",
                                            "terminationGracePeriodSeconds": 30,
                                            "tolerations": [
                                                {
                                                    "effect": "NoExecute",
                                                    "key": "node.kubernetes.io/not-ready",
                                                    "operator": "Exists",
                                                    "tolerationSeconds": 300
                                                },
                                                {
                                                    "effect": "NoExecute",
                                                    "key": "node.kubernetes.io/unreachable",
                                                    "operator": "Exists",
                                                    "tolerationSeconds": 300
                                                }
                                            ],
                                            "volumes": [
                                                {
                                                    "hostPath": {
                                                        "path": "/tmp",
                                                        "type": "Directory"
                                                    },
                                                    "name": "tmp"
                                                },
                                                {
                                                    "name": "storage-provisioner-token-bbjlq",
                                                    "secret": {
                                                        "defaultMode": 420,
                                                        "secretName": "storage-provisioner-token-bbjlq"
                                                    }
                                                }
                                            ]
                                        },
                                        "status": {
                                            "conditions": [
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-06-20T09:07:23Z",
                                                    "status": "True",
                                                    "type": "Initialized"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-14T05:36:17Z",
                                                    "status": "True",
                                                    "type": "Ready"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-14T05:36:17Z",
                                                    "status": "True",
                                                    "type": "ContainersReady"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-06-20T09:07:23Z",
                                                    "status": "True",
                                                    "type": "PodScheduled"
                                                }
                                            ],
                                            "containerStatuses": [
                                                {
                                                    "containerID": "docker://78d5935d6367da4a887877b19e472885edcfa17c8beb3425058f517002ccac4b",
                                                    "image": "gcr.io/k8s-minikube/storage-provisioner:v4",
                                                    "imageID": "docker-pullable://gcr.io/k8s-minikube/storage-provisioner@sha256:06f83c679a723d938b8776510d979c69549ad7df516279981e23554b3e68572f",
                                                    "lastState": {
                                                        "terminated": {
                                                            "containerID": "docker://4ac8e5fb478cd62a8edeb8a14ea0e63989df016d15f48a21279eee9a19e631a3",
                                                            "exitCode": 1,
                                                            "finishedAt": "2021-10-14T05:36:01Z",
                                                            "reason": "Error",
                                                            "startedAt": "2021-10-14T05:35:30Z"
                                                        }
                                                    },
                                                    "name": "storage-provisioner",
                                                    "ready": true,
                                                    "restartCount": 94,
                                                    "started": true,
                                                    "state": {
                                                        "running": {
                                                            "startedAt": "2021-10-14T05:36:16Z"
                                                        }
                                                    }
                                                }
                                            ],
                                            "hostIP": "10.0.2.15",
                                            "phase": "Running",
                                            "podIP": "10.0.2.15",
                                            "podIPs": [
                                                {
                                                    "ip": "10.0.2.15"
                                                }
                                            ],
                                            "qosClass": "BestEffort",
                                            "startTime": "2021-06-20T09:07:23Z"
                                        }
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "Workload: kube-proxy does not define any linux security hardening",
                            "ruleStatus": "",
                            "packagename": "armo_builtins",
                            "alertScore": 7,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "apps/v1",
                                        "kind": "DaemonSet",
                                        "metadata": {
                                            "annotations": {
                                                "deprecated.daemonset.template.generation": "1"
                                            },
                                            "creationTimestamp": "2021-06-20T09:07:08Z",
                                            "generation": 1,
                                            "labels": {
                                                "k8s-app": "kube-proxy"
                                            },
                                            "name": "kube-proxy",
                                            "namespace": "kube-system",
                                            "resourceVersion": "1450902",
                                            "selfLink": "/apis/apps/v1/namespaces/kube-system/daemonsets/kube-proxy",
                                            "uid": "dd1ba553-66da-47bc-8bc1-79c4b2f47dab"
                                        },
                                        "spec": {
                                            "revisionHistoryLimit": 10,
                                            "selector": {
                                                "matchLabels": {
                                                    "k8s-app": "kube-proxy"
                                                }
                                            },
                                            "template": {
                                                "metadata": {
                                                    "creationTimestamp": null,
                                                    "labels": {
                                                        "k8s-app": "kube-proxy"
                                                    }
                                                },
                                                "spec": {
                                                    "containers": [
                                                        {
                                                            "command": [
                                                                "/usr/local/bin/kube-proxy",
                                                                "--config=/var/lib/kube-proxy/config.conf",
                                                                "--hostname-override=$(NODE_NAME)"
                                                            ],
                                                            "env": [
                                                                {
                                                                    "name": "NODE_NAME",
                                                                    "valueFrom": {
                                                                        "fieldRef": {
                                                                            "apiVersion": "v1",
                                                                            "fieldPath": "spec.nodeName"
                                                                        }
                                                                    }
                                                                }
                                                            ],
                                                            "image": "k8s.gcr.io/kube-proxy:v1.16.0",
                                                            "imagePullPolicy": "IfNotPresent",
                                                            "name": "kube-proxy",
                                                            "resources": {},
                                                            "securityContext": {
                                                                "privileged": true
                                                            },
                                                            "terminationMessagePath": "/dev/termination-log",
                                                            "terminationMessagePolicy": "File",
                                                            "volumeMounts": [
                                                                {
                                                                    "mountPath": "/var/lib/kube-proxy",
                                                                    "name": "kube-proxy"
                                                                },
                                                                {
                                                                    "mountPath": "/run/xtables.lock",
                                                                    "name": "xtables-lock"
                                                                },
                                                                {
                                                                    "mountPath": "/lib/modules",
                                                                    "name": "lib-modules",
                                                                    "readOnly": true
                                                                }
                                                            ]
                                                        }
                                                    ],
                                                    "dnsPolicy": "ClusterFirst",
                                                    "hostNetwork": true,
                                                    "nodeSelector": {
                                                        "beta.kubernetes.io/os": "linux"
                                                    },
                                                    "priorityClassName": "system-node-critical",
                                                    "restartPolicy": "Always",
                                                    "schedulerName": "default-scheduler",
                                                    "securityContext": {},
                                                    "serviceAccount": "kube-proxy",
                                                    "serviceAccountName": "kube-proxy",
                                                    "terminationGracePeriodSeconds": 30,
                                                    "tolerations": [
                                                        {
                                                            "key": "CriticalAddonsOnly",
                                                            "operator": "Exists"
                                                        },
                                                        {
                                                            "operator": "Exists"
                                                        }
                                                    ],
                                                    "volumes": [
                                                        {
                                                            "configMap": {
                                                                "defaultMode": 420,
                                                                "name": "kube-proxy"
                                                            },
                                                            "name": "kube-proxy"
                                                        },
                                                        {
                                                            "hostPath": {
                                                                "path": "/run/xtables.lock",
                                                                "type": "FileOrCreate"
                                                            },
                                                            "name": "xtables-lock"
                                                        },
                                                        {
                                                            "hostPath": {
                                                                "path": "/lib/modules",
                                                                "type": ""
                                                            },
                                                            "name": "lib-modules"
                                                        }
                                                    ]
                                                }
                                            },
                                            "updateStrategy": {
                                                "rollingUpdate": {
                                                    "maxUnavailable": 1
                                                },
                                                "type": "RollingUpdate"
                                            }
                                        },
                                        "status": {
                                            "currentNumberScheduled": 1,
                                            "desiredNumberScheduled": 1,
                                            "numberAvailable": 1,
                                            "numberMisscheduled": 0,
                                            "numberReady": 1,
                                            "observedGeneration": 1,
                                            "updatedNumberScheduled": 1
                                        }
                                    }
                                ]
                            }
                        }
                    ]
                }
            ],
            "remediation": "You can use AppArmor, Seccomp, SELinux and Linux Capabilities mechanisms to restrict containers abilities to utilize unwanted privileges.",
            "description": "Containers may be given more privileges than they actually need. This can increase the potential impact of a container compromise.",
            "score": 14
        },
        {
            "guid": "",
            "attributes": {
                "armoBuiltin": true
            },
            "id": "C-0030",
            "controlID": "C-0030",
            "name": "Ingress and Egress blocked",
            "ruleReports": [
                {
                    "name": "ingress-and-egress-blocked",
                    "remediation": "",
                    "ruleStatus": {
                        "status": "success",
                        "message": ""
                    },
                    "ruleResponses": [
                        {
                            "alertMessage": "Pod: etcd-david-virtualbox does not have ingress/egress defined",
                            "ruleStatus": "failed",
                            "packagename": "armo_builtins",
                            "alertScore": 7,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "Pod",
                                        "metadata": {
                                            "annotations": {
                                                "kubernetes.io/config.hash": "e0fcc6e4323055b5880f8aac4c950836",
                                                "kubernetes.io/config.mirror": "e0fcc6e4323055b5880f8aac4c950836",
                                                "kubernetes.io/config.seen": "2021-06-20T12:06:52.495386281+03:00",
                                                "kubernetes.io/config.source": "file"
                                            },
                                            "creationTimestamp": "2021-06-20T09:08:22Z",
                                            "labels": {
                                                "component": "etcd",
                                                "tier": "control-plane"
                                            },
                                            "name": "etcd-david-virtualbox",
                                            "namespace": "kube-system",
                                            "resourceVersion": "2110909",
                                            "selfLink": "/api/v1/namespaces/kube-system/pods/etcd-david-virtualbox",
                                            "uid": "154e7f87-907f-4edb-a73c-26e965d4fe02"
                                        },
                                        "spec": {
                                            "containers": [
                                                {
                                                    "command": [
                                                        "etcd",
                                                        "--advertise-client-urls=https://10.0.2.15:2379",
                                                        "--cert-file=/var/lib/minikube/certs/etcd/server.crt",
                                                        "--client-cert-auth=true",
                                                        "--data-dir=/var/lib/minikube/etcd",
                                                        "--initial-advertise-peer-urls=https://10.0.2.15:2380",
                                                        "--initial-cluster=david-virtualbox=https://10.0.2.15:2380",
                                                        "--key-file=/var/lib/minikube/certs/etcd/server.key",
                                                        "--listen-client-urls=https://127.0.0.1:2379,https://10.0.2.15:2379",
                                                        "--listen-metrics-urls=http://127.0.0.1:2381,http://10.0.2.15:2381",
                                                        "--listen-peer-urls=https://10.0.2.15:2380",
                                                        "--name=david-virtualbox",
                                                        "--peer-cert-file=/var/lib/minikube/certs/etcd/peer.crt",
                                                        "--peer-client-cert-auth=true",
                                                        "--peer-key-file=/var/lib/minikube/certs/etcd/peer.key",
                                                        "--peer-trusted-ca-file=/var/lib/minikube/certs/etcd/ca.crt",
                                                        "--snapshot-count=10000",
                                                        "--trusted-ca-file=/var/lib/minikube/certs/etcd/ca.crt"
                                                    ],
                                                    "image": "k8s.gcr.io/etcd:3.3.15-0",
                                                    "imagePullPolicy": "IfNotPresent",
                                                    "livenessProbe": {
                                                        "failureThreshold": 8,
                                                        "httpGet": {
                                                            "host": "127.0.0.1",
                                                            "path": "/health",
                                                            "port": 2381,
                                                            "scheme": "HTTP"
                                                        },
                                                        "initialDelaySeconds": 15,
                                                        "periodSeconds": 10,
                                                        "successThreshold": 1,
                                                        "timeoutSeconds": 15
                                                    },
                                                    "name": "etcd",
                                                    "resources": {},
                                                    "terminationMessagePath": "/dev/termination-log",
                                                    "terminationMessagePolicy": "File",
                                                    "volumeMounts": [
                                                        {
                                                            "mountPath": "/var/lib/minikube/etcd",
                                                            "name": "etcd-data"
                                                        },
                                                        {
                                                            "mountPath": "/var/lib/minikube/certs/etcd",
                                                            "name": "etcd-certs"
                                                        }
                                                    ]
                                                }
                                            ],
                                            "dnsPolicy": "ClusterFirst",
                                            "enableServiceLinks": true,
                                            "hostNetwork": true,
                                            "nodeName": "david-virtualbox",
                                            "priority": 2000000000,
                                            "priorityClassName": "system-cluster-critical",
                                            "restartPolicy": "Always",
                                            "schedulerName": "default-scheduler",
                                            "securityContext": {},
                                            "terminationGracePeriodSeconds": 30,
                                            "tolerations": [
                                                {
                                                    "effect": "NoExecute",
                                                    "operator": "Exists"
                                                }
                                            ],
                                            "volumes": [
                                                {
                                                    "hostPath": {
                                                        "path": "/var/lib/minikube/certs/etcd",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "etcd-certs"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/var/lib/minikube/etcd",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "etcd-data"
                                                }
                                            ]
                                        },
                                        "status": {
                                            "conditions": [
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-03T06:28:57Z",
                                                    "status": "True",
                                                    "type": "Initialized"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-06T05:23:22Z",
                                                    "status": "True",
                                                    "type": "Ready"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-06T05:23:22Z",
                                                    "status": "True",
                                                    "type": "ContainersReady"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-03T06:28:57Z",
                                                    "status": "True",
                                                    "type": "PodScheduled"
                                                }
                                            ],
                                            "containerStatuses": [
                                                {
                                                    "containerID": "docker://0cc62102444cc74e3dddb2b6cd7550036feaa0e61f4297c85fc72e3be2905ec9",
                                                    "image": "k8s.gcr.io/etcd:3.3.15-0",
                                                    "imageID": "docker-pullable://k8s.gcr.io/etcd@sha256:12c2c5e5731c3bcd56e6f1c05c0f9198b6f06793fa7fca2fb43aab9622dc4afa",
                                                    "lastState": {
                                                        "terminated": {
                                                            "containerID": "docker://61de0abaa35617dafed9274c5b03f99f269412bdea2f15d903b4293d620f4b9e",
                                                            "exitCode": 255,
                                                            "finishedAt": "2021-10-14T05:35:02Z",
                                                            "reason": "Error",
                                                            "startedAt": "2021-10-12T10:28:38Z"
                                                        }
                                                    },
                                                    "name": "etcd",
                                                    "ready": true,
                                                    "restartCount": 51,
                                                    "started": true,
                                                    "state": {
                                                        "running": {
                                                            "startedAt": "2021-10-14T05:35:12Z"
                                                        }
                                                    }
                                                }
                                            ],
                                            "hostIP": "10.0.2.15",
                                            "phase": "Running",
                                            "podIP": "10.0.2.15",
                                            "podIPs": [
                                                {
                                                    "ip": "10.0.2.15"
                                                }
                                            ],
                                            "qosClass": "BestEffort",
                                            "startTime": "2021-10-03T06:28:57Z"
                                        }
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "Pod: kube-apiserver-david-virtualbox does not have ingress/egress defined",
                            "ruleStatus": "failed",
                            "packagename": "armo_builtins",
                            "alertScore": 7,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "Pod",
                                        "metadata": {
                                            "annotations": {
                                                "kubernetes.io/config.hash": "668e9396bc7b8c495d39f4a3bc479397",
                                                "kubernetes.io/config.mirror": "668e9396bc7b8c495d39f4a3bc479397",
                                                "kubernetes.io/config.seen": "2021-06-20T12:06:52.495387809+03:00",
                                                "kubernetes.io/config.source": "file"
                                            },
                                            "creationTimestamp": "2021-06-20T09:08:05Z",
                                            "labels": {
                                                "component": "kube-apiserver",
                                                "tier": "control-plane"
                                            },
                                            "name": "kube-apiserver-david-virtualbox",
                                            "namespace": "kube-system",
                                            "resourceVersion": "2110920",
                                            "selfLink": "/api/v1/namespaces/kube-system/pods/kube-apiserver-david-virtualbox",
                                            "uid": "327cbf13-97d6-42a3-8469-5acfdbe1be09"
                                        },
                                        "spec": {
                                            "containers": [
                                                {
                                                    "command": [
                                                        "kube-apiserver",
                                                        "--advertise-address=10.0.2.15",
                                                        "--allow-privileged=true",
                                                        "--authorization-mode=Node,RBAC",
                                                        "--client-ca-file=/var/lib/minikube/certs/ca.crt",
                                                        "--enable-admission-plugins=NamespaceLifecycle,LimitRanger,ServiceAccount,DefaultStorageClass,DefaultTolerationSeconds,NodeRestriction,MutatingAdmissionWebhook,ValidatingAdmissionWebhook,ResourceQuota",
                                                        "--enable-bootstrap-token-auth=true",
                                                        "--etcd-cafile=/var/lib/minikube/certs/etcd/ca.crt",
                                                        "--etcd-certfile=/var/lib/minikube/certs/apiserver-etcd-client.crt",
                                                        "--etcd-keyfile=/var/lib/minikube/certs/apiserver-etcd-client.key",
                                                        "--etcd-servers=https://127.0.0.1:2379",
                                                        "--insecure-port=0",
                                                        "--kubelet-client-certificate=/var/lib/minikube/certs/apiserver-kubelet-client.crt",
                                                        "--kubelet-client-key=/var/lib/minikube/certs/apiserver-kubelet-client.key",
                                                        "--kubelet-preferred-address-types=InternalIP,ExternalIP,Hostname",
                                                        "--proxy-client-cert-file=/var/lib/minikube/certs/front-proxy-client.crt",
                                                        "--proxy-client-key-file=/var/lib/minikube/certs/front-proxy-client.key",
                                                        "--requestheader-allowed-names=front-proxy-client",
                                                        "--requestheader-client-ca-file=/var/lib/minikube/certs/front-proxy-ca.crt",
                                                        "--requestheader-extra-headers-prefix=X-Remote-Extra-",
                                                        "--requestheader-group-headers=X-Remote-Group",
                                                        "--requestheader-username-headers=X-Remote-User",
                                                        "--secure-port=8443",
                                                        "--service-account-key-file=/var/lib/minikube/certs/sa.pub",
                                                        "--service-cluster-ip-range=10.96.0.0/12",
                                                        "--tls-cert-file=/var/lib/minikube/certs/apiserver.crt",
                                                        "--tls-private-key-file=/var/lib/minikube/certs/apiserver.key"
                                                    ],
                                                    "image": "k8s.gcr.io/kube-apiserver:v1.16.0",
                                                    "imagePullPolicy": "IfNotPresent",
                                                    "livenessProbe": {
                                                        "failureThreshold": 8,
                                                        "httpGet": {
                                                            "host": "10.0.2.15",
                                                            "path": "/healthz",
                                                            "port": 8443,
                                                            "scheme": "HTTPS"
                                                        },
                                                        "initialDelaySeconds": 15,
                                                        "periodSeconds": 10,
                                                        "successThreshold": 1,
                                                        "timeoutSeconds": 15
                                                    },
                                                    "name": "kube-apiserver",
                                                    "resources": {
                                                        "requests": {
                                                            "cpu": "250m"
                                                        }
                                                    },
                                                    "terminationMessagePath": "/dev/termination-log",
                                                    "terminationMessagePolicy": "File",
                                                    "volumeMounts": [
                                                        {
                                                            "mountPath": "/etc/ssl/certs",
                                                            "name": "ca-certs",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/etc/ca-certificates",
                                                            "name": "etc-ca-certificates",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/etc/pki",
                                                            "name": "etc-pki",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/var/lib/minikube/certs",
                                                            "name": "k8s-certs",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/usr/local/share/ca-certificates",
                                                            "name": "usr-local-share-ca-certificates",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/usr/share/ca-certificates",
                                                            "name": "usr-share-ca-certificates",
                                                            "readOnly": true
                                                        }
                                                    ]
                                                }
                                            ],
                                            "dnsPolicy": "ClusterFirst",
                                            "enableServiceLinks": true,
                                            "hostNetwork": true,
                                            "nodeName": "david-virtualbox",
                                            "priority": 2000000000,
                                            "priorityClassName": "system-cluster-critical",
                                            "restartPolicy": "Always",
                                            "schedulerName": "default-scheduler",
                                            "securityContext": {},
                                            "terminationGracePeriodSeconds": 30,
                                            "tolerations": [
                                                {
                                                    "effect": "NoExecute",
                                                    "operator": "Exists"
                                                }
                                            ],
                                            "volumes": [
                                                {
                                                    "hostPath": {
                                                        "path": "/etc/ssl/certs",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "ca-certs"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/etc/ca-certificates",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "etc-ca-certificates"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/etc/pki",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "etc-pki"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/var/lib/minikube/certs",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "k8s-certs"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/usr/local/share/ca-certificates",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "usr-local-share-ca-certificates"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/usr/share/ca-certificates",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "usr-share-ca-certificates"
                                                }
                                            ]
                                        },
                                        "status": {
                                            "conditions": [
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-03T06:28:57Z",
                                                    "status": "True",
                                                    "type": "Initialized"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-06T05:23:25Z",
                                                    "status": "True",
                                                    "type": "Ready"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-06T05:23:25Z",
                                                    "status": "True",
                                                    "type": "ContainersReady"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-03T06:28:57Z",
                                                    "status": "True",
                                                    "type": "PodScheduled"
                                                }
                                            ],
                                            "containerStatuses": [
                                                {
                                                    "containerID": "docker://c6d5933ac2752c18942e488d2d162fcc3141d1ae131be2abc3a50d2cb91e4e22",
                                                    "image": "k8s.gcr.io/kube-apiserver:v1.16.0",
                                                    "imageID": "docker-pullable://k8s.gcr.io/kube-apiserver@sha256:f4168527c91289da2708f62ae729fdde5fb484167dd05ffbb7ab666f60de96cd",
                                                    "lastState": {
                                                        "terminated": {
                                                            "containerID": "docker://3714c8295fe8b0ff97082305a44f08764e06d1a599ab3ee3ca6db15ab498a276",
                                                            "exitCode": 255,
                                                            "finishedAt": "2021-10-14T05:35:02Z",
                                                            "reason": "Error",
                                                            "startedAt": "2021-10-12T10:28:38Z"
                                                        }
                                                    },
                                                    "name": "kube-apiserver",
                                                    "ready": true,
                                                    "restartCount": 51,
                                                    "started": true,
                                                    "state": {
                                                        "running": {
                                                            "startedAt": "2021-10-14T05:35:12Z"
                                                        }
                                                    }
                                                }
                                            ],
                                            "hostIP": "10.0.2.15",
                                            "phase": "Running",
                                            "podIP": "10.0.2.15",
                                            "podIPs": [
                                                {
                                                    "ip": "10.0.2.15"
                                                }
                                            ],
                                            "qosClass": "Burstable",
                                            "startTime": "2021-10-03T06:28:57Z"
                                        }
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "Pod: kube-controller-manager-david-virtualbox does not have ingress/egress defined",
                            "ruleStatus": "failed",
                            "packagename": "armo_builtins",
                            "alertScore": 7,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "Pod",
                                        "metadata": {
                                            "annotations": {
                                                "kubernetes.io/config.hash": "a16b2d5766eae37796e4a8ed7f8ce12a",
                                                "kubernetes.io/config.mirror": "a16b2d5766eae37796e4a8ed7f8ce12a",
                                                "kubernetes.io/config.seen": "2021-06-20T12:06:52.495389283+03:00",
                                                "kubernetes.io/config.source": "file"
                                            },
                                            "creationTimestamp": "2021-06-20T09:08:00Z",
                                            "labels": {
                                                "component": "kube-controller-manager",
                                                "tier": "control-plane"
                                            },
                                            "name": "kube-controller-manager-david-virtualbox",
                                            "namespace": "kube-system",
                                            "resourceVersion": "2110899",
                                            "selfLink": "/api/v1/namespaces/kube-system/pods/kube-controller-manager-david-virtualbox",
                                            "uid": "6ca9d32c-21c3-4c0e-8087-5445c80a2bcc"
                                        },
                                        "spec": {
                                            "containers": [
                                                {
                                                    "command": [
                                                        "kube-controller-manager",
                                                        "--allocate-node-cidrs=true",
                                                        "--authentication-kubeconfig=/etc/kubernetes/controller-manager.conf",
                                                        "--authorization-kubeconfig=/etc/kubernetes/controller-manager.conf",
                                                        "--bind-address=127.0.0.1",
                                                        "--client-ca-file=/var/lib/minikube/certs/ca.crt",
                                                        "--cluster-cidr=10.244.0.0/16",
                                                        "--cluster-signing-cert-file=/var/lib/minikube/certs/ca.crt",
                                                        "--cluster-signing-key-file=/var/lib/minikube/certs/ca.key",
                                                        "--controllers=*,bootstrapsigner,tokencleaner",
                                                        "--kubeconfig=/etc/kubernetes/controller-manager.conf",
                                                        "--leader-elect=false",
                                                        "--node-cidr-mask-size=24",
                                                        "--requestheader-client-ca-file=/var/lib/minikube/certs/front-proxy-ca.crt",
                                                        "--root-ca-file=/var/lib/minikube/certs/ca.crt",
                                                        "--service-account-private-key-file=/var/lib/minikube/certs/sa.key",
                                                        "--service-cluster-ip-range=10.96.0.0/12",
                                                        "--use-service-account-credentials=true"
                                                    ],
                                                    "image": "k8s.gcr.io/kube-controller-manager:v1.16.0",
                                                    "imagePullPolicy": "IfNotPresent",
                                                    "livenessProbe": {
                                                        "failureThreshold": 8,
                                                        "httpGet": {
                                                            "host": "127.0.0.1",
                                                            "path": "/healthz",
                                                            "port": 10252,
                                                            "scheme": "HTTP"
                                                        },
                                                        "initialDelaySeconds": 15,
                                                        "periodSeconds": 10,
                                                        "successThreshold": 1,
                                                        "timeoutSeconds": 15
                                                    },
                                                    "name": "kube-controller-manager",
                                                    "resources": {
                                                        "requests": {
                                                            "cpu": "200m"
                                                        }
                                                    },
                                                    "terminationMessagePath": "/dev/termination-log",
                                                    "terminationMessagePolicy": "File",
                                                    "volumeMounts": [
                                                        {
                                                            "mountPath": "/etc/ssl/certs",
                                                            "name": "ca-certs",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/etc/ca-certificates",
                                                            "name": "etc-ca-certificates",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/etc/pki",
                                                            "name": "etc-pki",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/usr/libexec/kubernetes/kubelet-plugins/volume/exec",
                                                            "name": "flexvolume-dir"
                                                        },
                                                        {
                                                            "mountPath": "/var/lib/minikube/certs",
                                                            "name": "k8s-certs",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/etc/kubernetes/controller-manager.conf",
                                                            "name": "kubeconfig",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/usr/local/share/ca-certificates",
                                                            "name": "usr-local-share-ca-certificates",
                                                            "readOnly": true
                                                        },
                                                        {
                                                            "mountPath": "/usr/share/ca-certificates",
                                                            "name": "usr-share-ca-certificates",
                                                            "readOnly": true
                                                        }
                                                    ]
                                                }
                                            ],
                                            "dnsPolicy": "ClusterFirst",
                                            "enableServiceLinks": true,
                                            "hostNetwork": true,
                                            "nodeName": "david-virtualbox",
                                            "priority": 2000000000,
                                            "priorityClassName": "system-cluster-critical",
                                            "restartPolicy": "Always",
                                            "schedulerName": "default-scheduler",
                                            "securityContext": {},
                                            "terminationGracePeriodSeconds": 30,
                                            "tolerations": [
                                                {
                                                    "effect": "NoExecute",
                                                    "operator": "Exists"
                                                }
                                            ],
                                            "volumes": [
                                                {
                                                    "hostPath": {
                                                        "path": "/etc/ssl/certs",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "ca-certs"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/etc/ca-certificates",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "etc-ca-certificates"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/etc/pki",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "etc-pki"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/usr/libexec/kubernetes/kubelet-plugins/volume/exec",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "flexvolume-dir"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/var/lib/minikube/certs",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "k8s-certs"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/etc/kubernetes/controller-manager.conf",
                                                        "type": "FileOrCreate"
                                                    },
                                                    "name": "kubeconfig"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/usr/local/share/ca-certificates",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "usr-local-share-ca-certificates"
                                                },
                                                {
                                                    "hostPath": {
                                                        "path": "/usr/share/ca-certificates",
                                                        "type": "DirectoryOrCreate"
                                                    },
                                                    "name": "usr-share-ca-certificates"
                                                }
                                            ]
                                        },
                                        "status": {
                                            "conditions": [
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-14T05:35:11Z",
                                                    "status": "True",
                                                    "type": "Initialized"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-14T05:35:13Z",
                                                    "status": "True",
                                                    "type": "Ready"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-14T05:35:13Z",
                                                    "status": "True",
                                                    "type": "ContainersReady"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-14T05:35:11Z",
                                                    "status": "True",
                                                    "type": "PodScheduled"
                                                }
                                            ],
                                            "containerStatuses": [
                                                {
                                                    "containerID": "docker://eefb83fb1b81a497703013f609a54ffc03b57157395cb5f9a128dfeffea54b58",
                                                    "image": "k8s.gcr.io/kube-controller-manager:v1.16.0",
                                                    "imageID": "docker-pullable://k8s.gcr.io/kube-controller-manager@sha256:c156a05ee9d40e3ca2ebf9337f38a10558c1fc6c9124006f128a82e6c38cdf3e",
                                                    "lastState": {
                                                        "terminated": {
                                                            "containerID": "docker://81d8dca7f5b87e8b9bd8f515e9bcf86f4614306a9c729c30aad92560f520ed54",
                                                            "exitCode": 255,
                                                            "finishedAt": "2021-10-14T05:35:02Z",
                                                            "reason": "Error",
                                                            "startedAt": "2021-10-12T10:28:38Z"
                                                        }
                                                    },
                                                    "name": "kube-controller-manager",
                                                    "ready": true,
                                                    "restartCount": 55,
                                                    "started": true,
                                                    "state": {
                                                        "running": {
                                                            "startedAt": "2021-10-14T05:35:12Z"
                                                        }
                                                    }
                                                }
                                            ],
                                            "hostIP": "10.0.2.15",
                                            "phase": "Running",
                                            "podIP": "10.0.2.15",
                                            "podIPs": [
                                                {
                                                    "ip": "10.0.2.15"
                                                }
                                            ],
                                            "qosClass": "Burstable",
                                            "startTime": "2021-10-14T05:35:11Z"
                                        }
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "Pod: kube-scheduler-david-virtualbox does not have ingress/egress defined",
                            "ruleStatus": "failed",
                            "packagename": "armo_builtins",
                            "alertScore": 7,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "Pod",
                                        "metadata": {
                                            "annotations": {
                                                "kubernetes.io/config.hash": "b3d303074fe0ca1d42a8bd9ed248df09",
                                                "kubernetes.io/config.mirror": "b3d303074fe0ca1d42a8bd9ed248df09",
                                                "kubernetes.io/config.seen": "2021-06-20T12:06:52.495381685+03:00",
                                                "kubernetes.io/config.source": "file"
                                            },
                                            "creationTimestamp": "2021-06-20T09:08:11Z",
                                            "labels": {
                                                "component": "kube-scheduler",
                                                "tier": "control-plane"
                                            },
                                            "name": "kube-scheduler-david-virtualbox",
                                            "namespace": "kube-system",
                                            "resourceVersion": "2110888",
                                            "selfLink": "/api/v1/namespaces/kube-system/pods/kube-scheduler-david-virtualbox",
                                            "uid": "226e285e-b2b3-423c-8bd4-04cfb775cbc6"
                                        },
                                        "spec": {
                                            "containers": [
                                                {
                                                    "command": [
                                                        "kube-scheduler",
                                                        "--authentication-kubeconfig=/etc/kubernetes/scheduler.conf",
                                                        "--authorization-kubeconfig=/etc/kubernetes/scheduler.conf",
                                                        "--bind-address=127.0.0.1",
                                                        "--kubeconfig=/etc/kubernetes/scheduler.conf",
                                                        "--leader-elect=false"
                                                    ],
                                                    "image": "k8s.gcr.io/kube-scheduler:v1.16.0",
                                                    "imagePullPolicy": "IfNotPresent",
                                                    "livenessProbe": {
                                                        "failureThreshold": 8,
                                                        "httpGet": {
                                                            "host": "127.0.0.1",
                                                            "path": "/healthz",
                                                            "port": 10251,
                                                            "scheme": "HTTP"
                                                        },
                                                        "initialDelaySeconds": 15,
                                                        "periodSeconds": 10,
                                                        "successThreshold": 1,
                                                        "timeoutSeconds": 15
                                                    },
                                                    "name": "kube-scheduler",
                                                    "resources": {
                                                        "requests": {
                                                            "cpu": "100m"
                                                        }
                                                    },
                                                    "terminationMessagePath": "/dev/termination-log",
                                                    "terminationMessagePolicy": "File",
                                                    "volumeMounts": [
                                                        {
                                                            "mountPath": "/etc/kubernetes/scheduler.conf",
                                                            "name": "kubeconfig",
                                                            "readOnly": true
                                                        }
                                                    ]
                                                }
                                            ],
                                            "dnsPolicy": "ClusterFirst",
                                            "enableServiceLinks": true,
                                            "hostNetwork": true,
                                            "nodeName": "david-virtualbox",
                                            "priority": 2000000000,
                                            "priorityClassName": "system-cluster-critical",
                                            "restartPolicy": "Always",
                                            "schedulerName": "default-scheduler",
                                            "securityContext": {},
                                            "terminationGracePeriodSeconds": 30,
                                            "tolerations": [
                                                {
                                                    "effect": "NoExecute",
                                                    "operator": "Exists"
                                                }
                                            ],
                                            "volumes": [
                                                {
                                                    "hostPath": {
                                                        "path": "/etc/kubernetes/scheduler.conf",
                                                        "type": "FileOrCreate"
                                                    },
                                                    "name": "kubeconfig"
                                                }
                                            ]
                                        },
                                        "status": {
                                            "conditions": [
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-14T05:35:11Z",
                                                    "status": "True",
                                                    "type": "Initialized"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-14T05:35:13Z",
                                                    "status": "True",
                                                    "type": "Ready"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-14T05:35:13Z",
                                                    "status": "True",
                                                    "type": "ContainersReady"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-14T05:35:11Z",
                                                    "status": "True",
                                                    "type": "PodScheduled"
                                                }
                                            ],
                                            "containerStatuses": [
                                                {
                                                    "containerID": "docker://53c67aa4e9ac29bf19ebaa451fe9de11d2bad3442dc6ece2dc3e03a0dc266526",
                                                    "image": "k8s.gcr.io/kube-scheduler:v1.16.0",
                                                    "imageID": "docker-pullable://k8s.gcr.io/kube-scheduler@sha256:094023ab9cd02059eb0295d234ff9ea321e0e22e4813986d7f1a1ac4dc1990d0",
                                                    "lastState": {
                                                        "terminated": {
                                                            "containerID": "docker://fbe8092ea4bbc12f8ed65da47d80b24528557bc81a7fc54948da26e76d704b4a",
                                                            "exitCode": 255,
                                                            "finishedAt": "2021-10-14T05:35:02Z",
                                                            "reason": "Error",
                                                            "startedAt": "2021-10-12T10:28:38Z"
                                                        }
                                                    },
                                                    "name": "kube-scheduler",
                                                    "ready": true,
                                                    "restartCount": 52,
                                                    "started": true,
                                                    "state": {
                                                        "running": {
                                                            "startedAt": "2021-10-14T05:35:12Z"
                                                        }
                                                    }
                                                }
                                            ],
                                            "hostIP": "10.0.2.15",
                                            "phase": "Running",
                                            "podIP": "10.0.2.15",
                                            "podIPs": [
                                                {
                                                    "ip": "10.0.2.15"
                                                }
                                            ],
                                            "qosClass": "Burstable",
                                            "startTime": "2021-10-14T05:35:11Z"
                                        }
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "Pod: storage-provisioner does not have ingress/egress defined",
                            "ruleStatus": "failed",
                            "packagename": "armo_builtins",
                            "alertScore": 7,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "Pod",
                                        "metadata": {
                                            "annotations": {
                                                "kubectl.kubernetes.io/last-applied-configuration": "{\"apiVersion\":\"v1\",\"kind\":\"Pod\",\"metadata\":{\"annotations\":{},\"labels\":{\"addonmanager.kubernetes.io/mode\":\"Reconcile\",\"integration-test\":\"storage-provisioner\"},\"name\":\"storage-provisioner\",\"namespace\":\"kube-system\"},\"spec\":{\"containers\":[{\"command\":[\"/storage-provisioner\"],\"image\":\"gcr.io/k8s-minikube/storage-provisioner:v4\",\"imagePullPolicy\":\"IfNotPresent\",\"name\":\"storage-provisioner\",\"volumeMounts\":[{\"mountPath\":\"/tmp\",\"name\":\"tmp\"}]}],\"hostNetwork\":true,\"serviceAccountName\":\"storage-provisioner\",\"volumes\":[{\"hostPath\":{\"path\":\"/tmp\",\"type\":\"Directory\"},\"name\":\"tmp\"}]}}\n"
                                            },
                                            "creationTimestamp": "2021-06-20T09:07:09Z",
                                            "labels": {
                                                "addonmanager.kubernetes.io/mode": "Reconcile",
                                                "integration-test": "storage-provisioner"
                                            },
                                            "name": "storage-provisioner",
                                            "namespace": "kube-system",
                                            "resourceVersion": "2110982",
                                            "selfLink": "/api/v1/namespaces/kube-system/pods/storage-provisioner",
                                            "uid": "ea5dc2e2-4f7a-49f4-9e88-37e8e2d741a5"
                                        },
                                        "spec": {
                                            "containers": [
                                                {
                                                    "command": [
                                                        "/storage-provisioner"
                                                    ],
                                                    "image": "gcr.io/k8s-minikube/storage-provisioner:v4",
                                                    "imagePullPolicy": "IfNotPresent",
                                                    "name": "storage-provisioner",
                                                    "resources": {},
                                                    "terminationMessagePath": "/dev/termination-log",
                                                    "terminationMessagePolicy": "File",
                                                    "volumeMounts": [
                                                        {
                                                            "mountPath": "/tmp",
                                                            "name": "tmp"
                                                        },
                                                        {
                                                            "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount",
                                                            "name": "storage-provisioner-token-bbjlq",
                                                            "readOnly": true
                                                        }
                                                    ]
                                                }
                                            ],
                                            "dnsPolicy": "ClusterFirst",
                                            "enableServiceLinks": true,
                                            "hostNetwork": true,
                                            "nodeName": "david-virtualbox",
                                            "priority": 0,
                                            "restartPolicy": "Always",
                                            "schedulerName": "default-scheduler",
                                            "securityContext": {},
                                            "serviceAccount": "storage-provisioner",
                                            "serviceAccountName": "storage-provisioner",
                                            "terminationGracePeriodSeconds": 30,
                                            "tolerations": [
                                                {
                                                    "effect": "NoExecute",
                                                    "key": "node.kubernetes.io/not-ready",
                                                    "operator": "Exists",
                                                    "tolerationSeconds": 300
                                                },
                                                {
                                                    "effect": "NoExecute",
                                                    "key": "node.kubernetes.io/unreachable",
                                                    "operator": "Exists",
                                                    "tolerationSeconds": 300
                                                }
                                            ],
                                            "volumes": [
                                                {
                                                    "hostPath": {
                                                        "path": "/tmp",
                                                        "type": "Directory"
                                                    },
                                                    "name": "tmp"
                                                },
                                                {
                                                    "name": "storage-provisioner-token-bbjlq",
                                                    "secret": {
                                                        "defaultMode": 420,
                                                        "secretName": "storage-provisioner-token-bbjlq"
                                                    }
                                                }
                                            ]
                                        },
                                        "status": {
                                            "conditions": [
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-06-20T09:07:23Z",
                                                    "status": "True",
                                                    "type": "Initialized"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-14T05:36:17Z",
                                                    "status": "True",
                                                    "type": "Ready"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-10-14T05:36:17Z",
                                                    "status": "True",
                                                    "type": "ContainersReady"
                                                },
                                                {
                                                    "lastProbeTime": null,
                                                    "lastTransitionTime": "2021-06-20T09:07:23Z",
                                                    "status": "True",
                                                    "type": "PodScheduled"
                                                }
                                            ],
                                            "containerStatuses": [
                                                {
                                                    "containerID": "docker://78d5935d6367da4a887877b19e472885edcfa17c8beb3425058f517002ccac4b",
                                                    "image": "gcr.io/k8s-minikube/storage-provisioner:v4",
                                                    "imageID": "docker-pullable://gcr.io/k8s-minikube/storage-provisioner@sha256:06f83c679a723d938b8776510d979c69549ad7df516279981e23554b3e68572f",
                                                    "lastState": {
                                                        "terminated": {
                                                            "containerID": "docker://4ac8e5fb478cd62a8edeb8a14ea0e63989df016d15f48a21279eee9a19e631a3",
                                                            "exitCode": 1,
                                                            "finishedAt": "2021-10-14T05:36:01Z",
                                                            "reason": "Error",
                                                            "startedAt": "2021-10-14T05:35:30Z"
                                                        }
                                                    },
                                                    "name": "storage-provisioner",
                                                    "ready": true,
                                                    "restartCount": 94,
                                                    "started": true,
                                                    "state": {
                                                        "running": {
                                                            "startedAt": "2021-10-14T05:36:16Z"
                                                        }
                                                    }
                                                }
                                            ],
                                            "hostIP": "10.0.2.15",
                                            "phase": "Running",
                                            "podIP": "10.0.2.15",
                                            "podIPs": [
                                                {
                                                    "ip": "10.0.2.15"
                                                }
                                            ],
                                            "qosClass": "BestEffort",
                                            "startTime": "2021-06-20T09:07:23Z"
                                        }
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "Deployment: coredns has Pods which don't have ingress/egress defined",
                            "ruleStatus": "failed",
                            "packagename": "armo_builtins",
                            "alertScore": 7,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "apps/v1",
                                        "kind": "Deployment",
                                        "metadata": {
                                            "annotations": {
                                                "deployment.kubernetes.io/revision": "1"
                                            },
                                            "creationTimestamp": "2021-06-20T09:07:08Z",
                                            "generation": 2,
                                            "labels": {
                                                "k8s-app": "kube-dns"
                                            },
                                            "name": "coredns",
                                            "namespace": "kube-system",
                                            "resourceVersion": "2110974",
                                            "selfLink": "/apis/apps/v1/namespaces/kube-system/deployments/coredns",
                                            "uid": "fdab6339-f481-4fcb-aa9d-d6fa9260184b"
                                        },
                                        "spec": {
                                            "progressDeadlineSeconds": 600,
                                            "replicas": 1,
                                            "revisionHistoryLimit": 10,
                                            "selector": {
                                                "matchLabels": {
                                                    "k8s-app": "kube-dns"
                                                }
                                            },
                                            "strategy": {
                                                "rollingUpdate": {
                                                    "maxSurge": "25%",
                                                    "maxUnavailable": 1
                                                },
                                                "type": "RollingUpdate"
                                            },
                                            "template": {
                                                "metadata": {
                                                    "creationTimestamp": null,
                                                    "labels": {
                                                        "k8s-app": "kube-dns"
                                                    }
                                                },
                                                "spec": {
                                                    "containers": [
                                                        {
                                                            "args": [
                                                                "-conf",
                                                                "/etc/coredns/Corefile"
                                                            ],
                                                            "image": "k8s.gcr.io/coredns:1.6.2",
                                                            "imagePullPolicy": "IfNotPresent",
                                                            "livenessProbe": {
                                                                "failureThreshold": 5,
                                                                "httpGet": {
                                                                    "path": "/health",
                                                                    "port": 8080,
                                                                    "scheme": "HTTP"
                                                                },
                                                                "initialDelaySeconds": 60,
                                                                "periodSeconds": 10,
                                                                "successThreshold": 1,
                                                                "timeoutSeconds": 5
                                                            },
                                                            "name": "coredns",
                                                            "ports": [
                                                                {
                                                                    "containerPort": 53,
                                                                    "name": "dns",
                                                                    "protocol": "UDP"
                                                                },
                                                                {
                                                                    "containerPort": 53,
                                                                    "name": "dns-tcp",
                                                                    "protocol": "TCP"
                                                                },
                                                                {
                                                                    "containerPort": 9153,
                                                                    "name": "metrics",
                                                                    "protocol": "TCP"
                                                                }
                                                            ],
                                                            "readinessProbe": {
                                                                "failureThreshold": 3,
                                                                "httpGet": {
                                                                    "path": "/ready",
                                                                    "port": 8181,
                                                                    "scheme": "HTTP"
                                                                },
                                                                "periodSeconds": 10,
                                                                "successThreshold": 1,
                                                                "timeoutSeconds": 1
                                                            },
                                                            "resources": {
                                                                "limits": {
                                                                    "memory": "170Mi"
                                                                },
                                                                "requests": {
                                                                    "cpu": "100m",
                                                                    "memory": "70Mi"
                                                                }
                                                            },
                                                            "securityContext": {
                                                                "allowPrivilegeEscalation": false,
                                                                "capabilities": {
                                                                    "add": [
                                                                        "NET_BIND_SERVICE"
                                                                    ],
                                                                    "drop": [
                                                                        "all"
                                                                    ]
                                                                },
                                                                "readOnlyRootFilesystem": true
                                                            },
                                                            "terminationMessagePath": "/dev/termination-log",
                                                            "terminationMessagePolicy": "File",
                                                            "volumeMounts": [
                                                                {
                                                                    "mountPath": "/etc/coredns",
                                                                    "name": "config-volume",
                                                                    "readOnly": true
                                                                }
                                                            ]
                                                        }
                                                    ],
                                                    "dnsPolicy": "Default",
                                                    "nodeSelector": {
                                                        "beta.kubernetes.io/os": "linux"
                                                    },
                                                    "priorityClassName": "system-cluster-critical",
                                                    "restartPolicy": "Always",
                                                    "schedulerName": "default-scheduler",
                                                    "securityContext": {},
                                                    "serviceAccount": "coredns",
                                                    "serviceAccountName": "coredns",
                                                    "terminationGracePeriodSeconds": 30,
                                                    "tolerations": [
                                                        {
                                                            "key": "CriticalAddonsOnly",
                                                            "operator": "Exists"
                                                        },
                                                        {
                                                            "effect": "NoSchedule",
                                                            "key": "node-role.kubernetes.io/master"
                                                        }
                                                    ],
                                                    "volumes": [
                                                        {
                                                            "configMap": {
                                                                "defaultMode": 420,
                                                                "items": [
                                                                    {
                                                                        "key": "Corefile",
                                                                        "path": "Corefile"
                                                                    }
                                                                ],
                                                                "name": "coredns"
                                                            },
                                                            "name": "config-volume"
                                                        }
                                                    ]
                                                }
                                            }
                                        },
                                        "status": {
                                            "availableReplicas": 1,
                                            "conditions": [
                                                {
                                                    "lastTransitionTime": "2021-06-20T09:07:22Z",
                                                    "lastUpdateTime": "2021-06-20T09:07:22Z",
                                                    "message": "Deployment has minimum availability.",
                                                    "reason": "MinimumReplicasAvailable",
                                                    "status": "True",
                                                    "type": "Available"
                                                },
                                                {
                                                    "lastTransitionTime": "2021-06-20T09:07:22Z",
                                                    "lastUpdateTime": "2021-06-20T09:07:35Z",
                                                    "message": "ReplicaSet \"coredns-5644d7b6d9\" has successfully progressed.",
                                                    "reason": "NewReplicaSetAvailable",
                                                    "status": "True",
                                                    "type": "Progressing"
                                                }
                                            ],
                                            "observedGeneration": 2,
                                            "readyReplicas": 1,
                                            "replicas": 1,
                                            "updatedReplicas": 1
                                        }
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "DaemonSet: kube-proxy has Pods which don't have ingress/egress defined",
                            "ruleStatus": "failed",
                            "packagename": "armo_builtins",
                            "alertScore": 7,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "apps/v1",
                                        "kind": "DaemonSet",
                                        "metadata": {
                                            "annotations": {
                                                "deprecated.daemonset.template.generation": "1"
                                            },
                                            "creationTimestamp": "2021-06-20T09:07:08Z",
                                            "generation": 1,
                                            "labels": {
                                                "k8s-app": "kube-proxy"
                                            },
                                            "name": "kube-proxy",
                                            "namespace": "kube-system",
                                            "resourceVersion": "1450902",
                                            "selfLink": "/apis/apps/v1/namespaces/kube-system/daemonsets/kube-proxy",
                                            "uid": "dd1ba553-66da-47bc-8bc1-79c4b2f47dab"
                                        },
                                        "spec": {
                                            "revisionHistoryLimit": 10,
                                            "selector": {
                                                "matchLabels": {
                                                    "k8s-app": "kube-proxy"
                                                }
                                            },
                                            "template": {
                                                "metadata": {
                                                    "creationTimestamp": null,
                                                    "labels": {
                                                        "k8s-app": "kube-proxy"
                                                    }
                                                },
                                                "spec": {
                                                    "containers": [
                                                        {
                                                            "command": [
                                                                "/usr/local/bin/kube-proxy",
                                                                "--config=/var/lib/kube-proxy/config.conf",
                                                                "--hostname-override=$(NODE_NAME)"
                                                            ],
                                                            "env": [
                                                                {
                                                                    "name": "NODE_NAME",
                                                                    "valueFrom": {
                                                                        "fieldRef": {
                                                                            "apiVersion": "v1",
                                                                            "fieldPath": "spec.nodeName"
                                                                        }
                                                                    }
                                                                }
                                                            ],
                                                            "image": "k8s.gcr.io/kube-proxy:v1.16.0",
                                                            "imagePullPolicy": "IfNotPresent",
                                                            "name": "kube-proxy",
                                                            "resources": {},
                                                            "securityContext": {
                                                                "privileged": true
                                                            },
                                                            "terminationMessagePath": "/dev/termination-log",
                                                            "terminationMessagePolicy": "File",
                                                            "volumeMounts": [
                                                                {
                                                                    "mountPath": "/var/lib/kube-proxy",
                                                                    "name": "kube-proxy"
                                                                },
                                                                {
                                                                    "mountPath": "/run/xtables.lock",
                                                                    "name": "xtables-lock"
                                                                },
                                                                {
                                                                    "mountPath": "/lib/modules",
                                                                    "name": "lib-modules",
                                                                    "readOnly": true
                                                                }
                                                            ]
                                                        }
                                                    ],
                                                    "dnsPolicy": "ClusterFirst",
                                                    "hostNetwork": true,
                                                    "nodeSelector": {
                                                        "beta.kubernetes.io/os": "linux"
                                                    },
                                                    "priorityClassName": "system-node-critical",
                                                    "restartPolicy": "Always",
                                                    "schedulerName": "default-scheduler",
                                                    "securityContext": {},
                                                    "serviceAccount": "kube-proxy",
                                                    "serviceAccountName": "kube-proxy",
                                                    "terminationGracePeriodSeconds": 30,
                                                    "tolerations": [
                                                        {
                                                            "key": "CriticalAddonsOnly",
                                                            "operator": "Exists"
                                                        },
                                                        {
                                                            "operator": "Exists"
                                                        }
                                                    ],
                                                    "volumes": [
                                                        {
                                                            "configMap": {
                                                                "defaultMode": 420,
                                                                "name": "kube-proxy"
                                                            },
                                                            "name": "kube-proxy"
                                                        },
                                                        {
                                                            "hostPath": {
                                                                "path": "/run/xtables.lock",
                                                                "type": "FileOrCreate"
                                                            },
                                                            "name": "xtables-lock"
                                                        },
                                                        {
                                                            "hostPath": {
                                                                "path": "/lib/modules",
                                                                "type": ""
                                                            },
                                                            "name": "lib-modules"
                                                        }
                                                    ]
                                                }
                                            },
                                            "updateStrategy": {
                                                "rollingUpdate": {
                                                    "maxUnavailable": 1
                                                },
                                                "type": "RollingUpdate"
                                            }
                                        },
                                        "status": {
                                            "currentNumberScheduled": 1,
                                            "desiredNumberScheduled": 1,
                                            "numberAvailable": 1,
                                            "numberMisscheduled": 0,
                                            "numberReady": 1,
                                            "observedGeneration": 1,
                                            "updatedNumberScheduled": 1
                                        }
                                    }
                                ]
                            }
                        }
                    ]
                }
            ],
            "remediation": "Define a network policy that restricts ingress and egress connections.",
            "description": "Disable Ingress and Egress traffic on all pods wherever possible. It is recommended to define restrictive network policy on all new PODs, and then enable sources/destinations that this POD must communicate with.",
            "score": 0
        },
        {
            "guid": "",
            "attributes": {
                "armoBuiltin": true
            },
            "id": "C-0044",
            "controlID": "C-0044",
            "name": "Container hostPort",
            "ruleReports": [
                {
                    "name": "container-hostPort",
                    "remediation": "",
                    "ruleStatus": {
                        "status": "success",
                        "message": ""
                    },
                    "ruleResponses": null
                }
            ],
            "remediation": "Avoid usage of hostPort unless it is absolutely necessary. Use NodePort / ClusterIP instead.",
            "description": "Configuring hostPort limits you to a particular port, and if any two workloads that specify the same HostPort they cannot be deployed to the same node. Therefore, if the number of replica of such workload is higher than the number of nodes, the deployment will fail.",
            "score": 100
        },
        {
            "guid": "",
            "attributes": {
                "armoBuiltin": true
            },
            "id": "C-0011",
            "controlID": "C-0011",
            "name": "Network policies",
            "ruleReports": [
                {
                    "name": "internal-networking",
                    "remediation": "",
                    "ruleStatus": {
                        "status": "success",
                        "message": ""
                    },
                    "ruleResponses": [
                        {
                            "alertMessage": "no policy is defined for namespace argocd",
                            "ruleStatus": "",
                            "packagename": "armo_builtins",
                            "alertScore": 9,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "Namespace",
                                        "metadata": {
                                            "creationTimestamp": "2021-08-31T13:59:37Z",
                                            "name": "argocd",
                                            "resourceVersion": "1635073",
                                            "selfLink": "/api/v1/namespaces/argocd",
                                            "uid": "5b39f04a-a01b-4306-8e73-b79354ed6baf"
                                        },
                                        "spec": {
                                            "finalizers": [
                                                "kubernetes"
                                            ]
                                        },
                                        "status": {
                                            "phase": "Active"
                                        }
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "no policy is defined for namespace default",
                            "ruleStatus": "",
                            "packagename": "armo_builtins",
                            "alertScore": 9,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "Namespace",
                                        "metadata": {
                                            "creationTimestamp": "2021-06-20T09:07:06Z",
                                            "name": "default",
                                            "resourceVersion": "227958",
                                            "selfLink": "/api/v1/namespaces/default",
                                            "uid": "0e39819c-c076-466a-a392-7d7326cc3c07"
                                        },
                                        "spec": {
                                            "finalizers": [
                                                "kubernetes"
                                            ]
                                        },
                                        "status": {
                                            "phase": "Active"
                                        }
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "no policy is defined for namespace kube-node-lease",
                            "ruleStatus": "",
                            "packagename": "armo_builtins",
                            "alertScore": 9,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "Namespace",
                                        "metadata": {
                                            "creationTimestamp": "2021-06-20T09:07:03Z",
                                            "name": "kube-node-lease",
                                            "resourceVersion": "39",
                                            "selfLink": "/api/v1/namespaces/kube-node-lease",
                                            "uid": "4438f409-7c9e-4134-95f8-e25485b84f55"
                                        },
                                        "spec": {
                                            "finalizers": [
                                                "kubernetes"
                                            ]
                                        },
                                        "status": {
                                            "phase": "Active"
                                        }
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "no policy is defined for namespace kube-public",
                            "ruleStatus": "",
                            "packagename": "armo_builtins",
                            "alertScore": 9,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "Namespace",
                                        "metadata": {
                                            "creationTimestamp": "2021-06-20T09:07:03Z",
                                            "name": "kube-public",
                                            "resourceVersion": "31",
                                            "selfLink": "/api/v1/namespaces/kube-public",
                                            "uid": "4e17425d-d1f8-463a-b8a1-cfa0b568c84b"
                                        },
                                        "spec": {
                                            "finalizers": [
                                                "kubernetes"
                                            ]
                                        },
                                        "status": {
                                            "phase": "Active"
                                        }
                                    }
                                ]
                            }
                        },
                        {
                            "alertMessage": "no policy is defined for namespace kube-system",
                            "ruleStatus": "",
                            "packagename": "armo_builtins",
                            "alertScore": 9,
                            "alertObject": {
                                "k8sApiObjects": [
                                    {
                                        "apiVersion": "v1",
                                        "kind": "Namespace",
                                        "metadata": {
                                            "creationTimestamp": "2021-06-20T09:07:03Z",
                                            "name": "kube-system",
                                            "resourceVersion": "24",
                                            "selfLink": "/api/v1/namespaces/kube-system",
                                            "uid": "07dd0388-3e6a-4af2-83c1-52b040706985"
                                        },
                                        "spec": {
                                            "finalizers": [
                                                "kubernetes"
                                            ]
                                        },
                                        "status": {
                                            "phase": "Active"
                                        }
                                    }
                                ]
                            }
                        }
                    ]
                }
            ],
            "remediation": "Define network policies or use similar network protection mechanisms.",
            "description": "If no network policy is defined, attackers who gain access to a single container may use it to probe the network. This control lists all namespaces in which no network policies are defined.",
            "score": 0
        },
        {
            "guid": "",
            "attributes": {
                "armoBuiltin": true
            },
            "id": "C-0058",
            "controlID": "C-0058",
            "name": "CVE-2021-25741 - Using symlink for arbitrary host file system access.",
            "ruleReports": [
                {
                    "name": "Symlink-Exchange-Can-Allow-Host-Filesystem-Access",
                    "remediation": "",
                    "ruleStatus": {
                        "status": "success",
                        "message": ""
                    },
                    "ruleResponses": null
                }
            ],
            "remediation": "To mitigate this vulnerability without upgrading kubelet, you can disable the VolumeSubpath feature gate on kubelet and kube-apiserver, or remove any existing Pods using subPath or subPathExpr feature.",
            "description": "A user may be able to create a container with subPath or subPathExpr volume mounts to access files \u0026 directories anywhere on the host filesystem. Following Kubernetes versions are affected: v1.22.0 - v1.22.1, v1.21.0 - v1.21.4, v1.20.0 - v1.20.10, version v1.19.14 and lower. This control checks the vulnerable versions and the actual usage of the subPath feature in all Pods in the cluster.",
            "score": 100
        }
    ],
    "score": 10
}
`
