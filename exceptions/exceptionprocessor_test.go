package exceptions

import (
	"testing"

	"github.com/armosec/armoapi-go/identifiers"

	"github.com/kubescape/k8s-interface/workloadinterface"
	"github.com/kubescape/opa-utils/objectsenvelopes"
	"github.com/kubescape/opa-utils/reporthandling"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/armosec/armoapi-go/armotypes"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

/* unused for now
func postureExceptionPolicyDisableMock() *armotypes.PostureExceptionPolicy {
	return &armotypes.PostureExceptionPolicy{}
}
*/

func postureExceptionPolicyAlertOnlyMock() *armotypes.PostureExceptionPolicy {
	return &armotypes.PostureExceptionPolicy{
		PortalBase: armotypes.PortalBase{
			Name: "postureExceptionPolicyAlertOnlyMock",
		},
		PolicyType: "postureExceptionPolicy",
		Actions:    []armotypes.PostureExceptionPolicyActions{armotypes.AlertOnly},
		Resources: []identifiers.PortalDesignator{
			{
				DesignatorType: identifiers.DesignatorAttributes,
				Attributes: map[string]string{
					identifiers.AttributeNamespace: "default",
					identifiers.AttributeCluster:   "unittest",
				},
			},
		},
		PosturePolicies: []armotypes.PosturePolicy{
			{
				FrameworkName: "MIT.*",
			},
		},
	}
}

func postureLabelsRegexExceptionPolicyAlertOnlyMock() *armotypes.PostureExceptionPolicy {
	return &armotypes.PostureExceptionPolicy{
		PortalBase: armotypes.PortalBase{
			Name: "postureLabelsRegexExceptionPolicyAlertOnlyMock",
		},
		PolicyType: "postureExceptionPolicy",
		Actions:    []armotypes.PostureExceptionPolicyActions{armotypes.AlertOnly},
		Resources: []identifiers.PortalDesignator{
			{
				DesignatorType: identifiers.DesignatorAttributes,
				Attributes: map[string]string{
					"myLabelOrAnnotation": "static_test",
				},
			},
		},
		PosturePolicies: []armotypes.PosturePolicy{
			{
				FrameworkName: "MIT.*",
			},
		},
	}
}

func postureResourceIDExceptionPolicyMock(resourceID string) *armotypes.PostureExceptionPolicy {
	return &armotypes.PostureExceptionPolicy{
		PortalBase: armotypes.PortalBase{
			Name: "postureResourceIDExceptionPolicyMock",
		},
		PolicyType: "postureExceptionPolicy",
		Actions:    []armotypes.PostureExceptionPolicyActions{armotypes.AlertOnly},
		Resources: []identifiers.PortalDesignator{
			{
				DesignatorType: identifiers.DesignatorAttributes,
				Attributes: map[string]string{
					identifiers.AttributeCluster:    "test",
					identifiers.AttributeResourceID: resourceID,
				},
			},
		},
		PosturePolicies: []armotypes.PosturePolicy{
			{
				FrameworkName: "MIT.*",
			},
		},
	}
}

func emptyPostureExceptionPolicyAlertOnlyMock() *armotypes.PostureExceptionPolicy {
	return &armotypes.PostureExceptionPolicy{
		PortalBase: armotypes.PortalBase{
			Name: "postureExceptionPolicyAlertOnlyMock",
		},
		PolicyType: "postureExceptionPolicy",
		Actions:    []armotypes.PostureExceptionPolicyActions{armotypes.AlertOnly},
		Resources: []identifiers.PortalDesignator{
			{
				DesignatorType: identifiers.DesignatorAttributes,
				Attributes:     map[string]string{},
			},
		},
		PosturePolicies: []armotypes.PosturePolicy{},
	}
}
func TestListRuleExceptions(t *testing.T) {
	p := NewProcessor()
	exceptionPolicies := []armotypes.PostureExceptionPolicy{*postureExceptionPolicyAlertOnlyMock()}
	res1 := p.ListRuleExceptions(exceptionPolicies, "MITRE", "", "")
	assert.Equal(t, 1, len(res1))

	res2 := p.ListRuleExceptions(exceptionPolicies, "", "", "")
	assert.Equal(t, len(res2), 1)

	res3 := p.ListRuleExceptions(exceptionPolicies, "NSA", "", "")
	assert.Equal(t, len(res3), 0)

}

func TestListRuleExceptionsRegex(t *testing.T) {
	p := NewProcessor()
	exceptionPolicy := emptyPostureExceptionPolicyAlertOnlyMock()
	res1 := p.ListRuleExceptions([]armotypes.PostureExceptionPolicy{*exceptionPolicy}, "MITRE", "", "")
	assert.Equal(t, 1, len(res1))

	exceptionPolicy.PosturePolicies = append(exceptionPolicy.PosturePolicies, armotypes.PosturePolicy{
		FrameworkName: "MIT.*",
	})

	res2 := p.ListRuleExceptions([]armotypes.PostureExceptionPolicy{*exceptionPolicy}, "2MITRE", "", "")
	assert.Equal(t, 0, len(res2))

	exceptionPolicy.PosturePolicies[0] = armotypes.PosturePolicy{
		FrameworkName: "mit.*",
	}
	res2 = p.ListRuleExceptions([]armotypes.PostureExceptionPolicy{*exceptionPolicy}, "MITRE", "", "")
	assert.Equal(t, 1, len(res2))

	exceptionPolicy.PosturePolicies[0] = armotypes.PosturePolicy{
		FrameworkName: "mitre",
	}
	res2 = p.ListRuleExceptions([]armotypes.PostureExceptionPolicy{*exceptionPolicy}, "MITRE", "", "")
	assert.Equal(t, 1, len(res2))

	exceptionPolicy.PosturePolicies[0] = armotypes.PosturePolicy{
		FrameworkName: "MITRE",
		ControlName:   "my.*", // deprecated
		RuleName:      "rule.*vk",
	}

	res3 := p.ListRuleExceptions([]armotypes.PostureExceptionPolicy{*exceptionPolicy}, "MITRE", "", "")
	assert.Equal(t, 1, len(res3))

	res3 = p.ListRuleExceptions([]armotypes.PostureExceptionPolicy{*exceptionPolicy}, "MITRE", "", "")
	assert.Equal(t, 1, len(res3))

	res3 = p.ListRuleExceptions([]armotypes.PostureExceptionPolicy{*exceptionPolicy}, "MITRE", "", "rulebla -bla vk")
	assert.Equal(t, 1, len(res3))

	res3 = p.ListRuleExceptions([]armotypes.PostureExceptionPolicy{*exceptionPolicy}, "MITRE", "", "rulebla -bla")
	assert.Equal(t, 0, len(res3))
}

func TestGetResourceExceptions(t *testing.T) {
	emptyObj, err := workloadinterface.NewBaseObjBytes([]byte(`{"apiVersion": "v1", "kind":"Deployment", "metadata": {"name": "test"}}`))
	require.NoError(t, err)

	withLabelObj, err := workloadinterface.NewBaseObjBytes([]byte(`{"apiVersion": "v1", "kind":"Deployment", "metadata": {"name": "test", "labels": {"myLabelOrAnnotation" : "static_test", "mySecondLabelOrAnnotation" : "second_static_test"}}}`))
	require.NoError(t, err)

	withAnnotationObj, err := workloadinterface.NewBaseObjBytes([]byte(`{"apiVersion": "v1", "kind":"Deployment", "metadata": {"name": "test", "annotations": {"myLabelOrAnnotation" : "static_test", "mySecondLabelOrAnnotation" : "second_static_test"}}}`))
	require.NoError(t, err)

	idObj, err := workloadinterface.NewBaseObjBytes([]byte(`{"apiVersion": "v1/core", "kind":"Deployment", "metadata": {"name": "test", "namespace": "default"}}`))
	require.NoError(t, err)

	exceptionPolicyResourceID := postureResourceIDExceptionPolicyMock(idObj.GetID())
	exceptionPolicyResourceIDRegex := postureResourceIDExceptionPolicyMock("*")
	exceptionPolicyResourceOtherID := postureResourceIDExceptionPolicyMock("v1/core/default/ConfigMap/test")

	exceptionPolicy := postureLabelsRegexExceptionPolicyAlertOnlyMock()
	exceptionPolicyRegex := postureLabelsRegexExceptionPolicyAlertOnlyMock()
	exceptionPolicyRegex.Resources[0].Attributes["myLabelOrAnnotation"] = "static_.*"

	p := NewProcessor()

	testCases := []struct {
		workloadObj             workloadinterface.IMetadata
		exceptionPolicy         *armotypes.PostureExceptionPolicy
		desc                    string
		expectedExceptionsCount int
	}{
		{
			desc:                    "no label nor annotation",
			exceptionPolicy:         exceptionPolicy,
			workloadObj:             emptyObj,
			expectedExceptionsCount: 0,
		},
		{
			desc:                    "no label nor annotation (regexp)",
			exceptionPolicy:         exceptionPolicyRegex,
			workloadObj:             emptyObj,
			expectedExceptionsCount: 0,
		},
		{
			desc:                    "static label",
			exceptionPolicy:         exceptionPolicy,
			workloadObj:             withLabelObj,
			expectedExceptionsCount: 1,
		},
		{
			desc:                    "static annotation",
			exceptionPolicy:         exceptionPolicy,
			workloadObj:             withAnnotationObj,
			expectedExceptionsCount: 1,
		},
		{
			desc:                    "regex label",
			exceptionPolicy:         exceptionPolicyRegex,
			workloadObj:             withLabelObj,
			expectedExceptionsCount: 1,
		},
		{
			desc:                    "regex annotation",
			exceptionPolicy:         exceptionPolicyRegex,
			workloadObj:             withAnnotationObj,
			expectedExceptionsCount: 1,
		},
		{
			desc:                    "exception by ID",
			exceptionPolicy:         exceptionPolicyResourceID,
			workloadObj:             idObj,
			expectedExceptionsCount: 1,
		},
		{
			desc:                    "exception by ID regex",
			exceptionPolicy:         exceptionPolicyResourceIDRegex,
			workloadObj:             idObj,
			expectedExceptionsCount: 1,
		},
		{
			desc:                    "exception with not matching ID",
			exceptionPolicy:         exceptionPolicyResourceOtherID,
			workloadObj:             idObj,
			expectedExceptionsCount: 0,
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			res := p.GetResourceExceptions([]armotypes.PostureExceptionPolicy{*test.exceptionPolicy}, test.workloadObj, "test")
			assert.Equal(t, test.expectedExceptionsCount, len(res))
		})
	}
}

func TestRegexCompare(t *testing.T) {
	c := newComparator()

	assert.True(t, c.compareCluster(".*minikube.*", "bez-minikube-25-10"))
	assert.True(t, c.compareCluster("bez-minikube-25-10", "bez-minikube-25-10"))
	assert.False(t, c.compareCluster("minikube", "bez-minikube-25-10"))
	assert.False(t, c.compareCluster("bla", "bez-minikube-25-10"))
}

func TestHasException(t *testing.T) {
	processor := NewProcessor()

	tests := []struct {
		workload    workloadinterface.IMetadata
		designator  *identifiers.PortalDesignator
		name        string
		clusterName string
		expected    bool
	}{
		{
			name:        "Test case: Missing attributes",
			clusterName: "cluster1",
			designator: &identifiers.PortalDesignator{
				DesignatorType: identifiers.DesignatorAttributes,
				Attributes:     map[string]string{},
			},
			workload: workloadinterface.NewWorkloadObj(
				map[string]interface{}{},
			),
			expected: false,
		},
		{
			name:        "Test case: Matching cluster name",
			clusterName: "cluster1",
			designator: &identifiers.PortalDesignator{
				DesignatorType: identifiers.DesignatorAttributes,
				Attributes: map[string]string{
					"cluster": "cluster1",
				},
			},
			workload: workloadinterface.NewWorkloadObj(
				map[string]interface{}{},
			),
			expected: true,
		},
		{
			name:        "Test case: Non-matching cluster name",
			clusterName: "cluster1",
			designator: &identifiers.PortalDesignator{
				DesignatorType: identifiers.DesignatorAttributes,
				Attributes: map[string]string{
					"cluster": "cluster2",
				},
			},
			workload: workloadinterface.NewWorkloadObj(
				map[string]interface{}{},
			),
			expected: false,
		},
		{
			name:        "Test case: Matching cluster name with regex",
			clusterName: "cluster1",
			designator: &identifiers.PortalDesignator{
				DesignatorType: identifiers.DesignatorAttributes,
				Attributes: map[string]string{
					"cluster": "cluster.*",
				},
			},
			workload: workloadinterface.NewWorkloadObj(
				map[string]interface{}{},
			),
			expected: true,
		},
		{
			name: "Test case: Kind matches",
			designator: &identifiers.PortalDesignator{
				DesignatorType: identifiers.DesignatorAttributes,
				Attributes: map[string]string{
					"kind": "Deployment",
				},
			},
			workload: workloadinterface.NewWorkloadObj(
				map[string]interface{}{
					"kind": "Deployment",
				},
			),
			expected: true,
		},
		{
			name: "Test case: Name matches",
			designator: &identifiers.PortalDesignator{
				DesignatorType: identifiers.DesignatorAttributes,
				Attributes: map[string]string{
					"name": "test-workload",
				},
			},
			workload: workloadinterface.NewWorkloadObj(
				map[string]interface{}{
					"metadata": map[string]interface{}{
						"name": "test-workload",
					},
				},
			),
			expected: true,
		},
		{
			name: "Test case: Namespace matches",
			designator: &identifiers.PortalDesignator{
				DesignatorType: identifiers.DesignatorAttributes,
				Attributes: map[string]string{
					"namespace": "test-namespace",
				},
			},
			workload: workloadinterface.NewWorkloadObj(
				map[string]interface{}{
					"metadata": map[string]interface{}{
						"namespace": "test-namespace",
					},
				},
			),
			expected: true,
		},
		{
			name:        "Test case: Kind matches with regex",
			clusterName: "cluster1",
			designator: &identifiers.PortalDesignator{
				DesignatorType: identifiers.DesignatorAttributes,
				Attributes: map[string]string{
					"kind": "Deploy.*",
				},
			},
			workload: workloadinterface.NewWorkloadObj(
				map[string]interface{}{
					"kind": "Deployment",
				},
			),
			expected: true,
		},
		{
			name:        "Test case 3: Name matches with regex",
			clusterName: "cluster1",
			designator: &identifiers.PortalDesignator{
				DesignatorType: identifiers.DesignatorAttributes,
				Attributes: map[string]string{
					"name": "test-.*",
				},
			},
			workload: workloadinterface.NewWorkloadObj(
				map[string]interface{}{
					"metadata": map[string]interface{}{
						"name": "test-workload",
					},
				},
			),
			expected: true,
		},
		{
			name:        "Test case 4: Namespace matches with regex",
			clusterName: "cluster1",
			designator: &identifiers.PortalDesignator{
				DesignatorType: identifiers.DesignatorAttributes,
				Attributes: map[string]string{
					"namespace": "test-.*",
				},
			},
			workload: workloadinterface.NewWorkloadObj(
				map[string]interface{}{
					"metadata": map[string]interface{}{
						"namespace": "test-namespace",
					},
				},
			),
			expected: true,
		},
		{
			name:        "Test case 5: Kind does not match",
			clusterName: "cluster1",
			designator: &identifiers.PortalDesignator{
				DesignatorType: identifiers.DesignatorAttributes,
				Attributes: map[string]string{
					"kind": "Service",
				},
			},
			workload: workloadinterface.NewWorkloadObj(
				map[string]interface{}{
					"kind": "Deployment",
				},
			),
			expected: false,
		},
		{
			name:        "Test case 6: Name does not match",
			clusterName: "cluster1",
			designator: &identifiers.PortalDesignator{
				DesignatorType: identifiers.DesignatorAttributes,
				Attributes: map[string]string{
					"name": "different-workload",
				},
			},
			workload: workloadinterface.NewWorkloadObj(
				map[string]interface{}{
					"metadata": map[string]interface{}{
						"name": "test-workload",
					},
				},
			),
			expected: false,
		},
		{
			name:        "Test case: Namespace does not match",
			clusterName: "cluster1",
			designator: &identifiers.PortalDesignator{
				DesignatorType: identifiers.DesignatorAttributes,
				Attributes: map[string]string{
					"namespace": "different-namespace",
				},
			},
			workload: workloadinterface.NewWorkloadObj(
				map[string]interface{}{
					"metadata": map[string]interface{}{
						"namespace": "test-namespace",
					},
				},
			),
			expected: false,
		},
		{
			name: "Test case: Path matches",
			designator: &identifiers.PortalDesignator{
				DesignatorType: identifiers.DesignatorAttributes,
				Attributes: map[string]string{
					"path": "/path/to/source",
				},
			},
			workload: workloadinterface.NewWorkloadObj(
				map[string]interface{}{
					"sourcePath": "/path/to/source",
				},
			),
			expected: true,
		},
		{
			name: "Test case: Path matches with regex",
			designator: &identifiers.PortalDesignator{
				DesignatorType: identifiers.DesignatorAttributes,
				Attributes: map[string]string{
					"path": "/path/.*/source",
				},
			},
			workload: workloadinterface.NewWorkloadObj(
				map[string]interface{}{
					"sourcePath": "/path/to/source",
				},
			),
			expected: true,
		},
		{
			name: "Test case: Path does not match",
			designator: &identifiers.PortalDesignator{
				DesignatorType: identifiers.DesignatorAttributes,
				Attributes: map[string]string{
					"path": "/path/to/source",
				},
			},
			workload: workloadinterface.NewWorkloadObj(
				map[string]interface{}{
					"sourcePath": "/path/to/dest",
				},
			),
			expected: false,
		},
		{
			name: "Test case: Labels match",
			designator: &identifiers.PortalDesignator{
				DesignatorType: identifiers.DesignatorAttributes,
				Attributes: map[string]string{
					"key1": "val1",
					"key2": "val2",
				},
			},
			workload: workloadinterface.NewWorkloadObj(
				map[string]interface{}{
					"apiVersion": "v1",
					"kind":       "Namespace",
					"metadata": map[string]interface{}{
						"name": "test",
						"labels": map[string]interface{}{
							"key1": "val1",
							"key2": "val2",
						},
					},
				},
			),
			expected: true,
		},
		{
			name: "Test case: Labels do not match",
			designator: &identifiers.PortalDesignator{
				DesignatorType: identifiers.DesignatorAttributes,
				Attributes: map[string]string{
					"key1": "val1",
					"key2": "val2",
				},
			},
			workload: workloadinterface.NewWorkloadObj(
				map[string]interface{}{
					"apiVersion": "v1",
					"kind":       "Namespace",
					"metadata": map[string]interface{}{
						"name": "test",
						"labels": map[string]interface{}{
							"key1": "val1",
							"key2": "val3",
						},
					},
				},
			),
			expected: false,
		},
		{
			name: "Test case: Labels missing",
			designator: &identifiers.PortalDesignator{
				DesignatorType: identifiers.DesignatorAttributes,
				Attributes: map[string]string{
					"key1": "val1",
					"key2": "val2",
				},
			},
			workload: workloadinterface.NewWorkloadObj(
				map[string]interface{}{
					"apiVersion": "v1",
					"kind":       "Namespace",
					"metadata": map[string]interface{}{
						"name": "test",
						"labels": map[string]interface{}{
							"key1": "val1",
						},
					},
				},
			),
			expected: false,
		},
		{
			name: "Test case: Labels match regex",
			designator: &identifiers.PortalDesignator{
				DesignatorType: identifiers.DesignatorAttributes,
				Attributes: map[string]string{
					"key1": ".*",
					"key2": ".*",
				},
			},
			workload: workloadinterface.NewWorkloadObj(
				map[string]interface{}{
					"apiVersion": "v1",
					"kind":       "Namespace",
					"metadata": map[string]interface{}{
						"name": "test",
						"labels": map[string]interface{}{
							"key1": "val1",
							"key2": "val2",
						},
					},
				},
			),
			expected: true,
		},
		{
			name: "Test case: Labels dont match regex",
			designator: &identifiers.PortalDesignator{
				DesignatorType: identifiers.DesignatorAttributes,
				Attributes: map[string]string{
					"key1": "val.*",
					"key2": "val.*",
				},
			},
			workload: workloadinterface.NewWorkloadObj(
				map[string]interface{}{
					"apiVersion": "v1",
					"kind":       "Namespace",
					"metadata": map[string]interface{}{
						"name": "test",
						"labels": map[string]interface{}{
							"key1": "val1",
							"key2": "bla2",
						},
					},
				},
			),
			expected: false,
		},
		{
			name: "Test case: Annotations match",
			designator: &identifiers.PortalDesignator{
				DesignatorType: identifiers.DesignatorAttributes,
				Attributes: map[string]string{
					"key1": "val1",
					"key2": "val2",
				},
			},
			workload: workloadinterface.NewWorkloadObj(
				map[string]interface{}{
					"apiVersion": "v1",
					"kind":       "Namespace",
					"metadata": map[string]interface{}{
						"name": "test",
						"annotations": map[string]interface{}{
							"key1": "val1",
							"key2": "val2",
						},
					},
				},
			),
			expected: true,
		},
		{
			name: "Test case: Annotations do not match",
			designator: &identifiers.PortalDesignator{
				DesignatorType: identifiers.DesignatorAttributes,
				Attributes: map[string]string{
					"key1": "val1",
					"key2": "val2",
				},
			},
			workload: workloadinterface.NewWorkloadObj(
				map[string]interface{}{
					"apiVersion": "v1",
					"kind":       "Namespace",
					"metadata": map[string]interface{}{
						"name": "test",
						"annotations": map[string]interface{}{
							"key1": "val1",
							"key2": "val3",
						},
					},
				},
			),
			expected: false,
		},
		{
			name: "Test case: annotations missing",
			designator: &identifiers.PortalDesignator{
				DesignatorType: identifiers.DesignatorAttributes,
				Attributes: map[string]string{
					"key1": "val1",
					"key2": "val2",
				},
			},
			workload: workloadinterface.NewWorkloadObj(
				map[string]interface{}{
					"apiVersion": "v1",
					"kind":       "Namespace",
					"metadata": map[string]interface{}{
						"name": "test",
						"annotations": map[string]interface{}{
							"key1": "val1",
						},
					},
				},
			),
			expected: false,
		},
		{
			name: "Test case: annotations match regex",
			designator: &identifiers.PortalDesignator{
				DesignatorType: identifiers.DesignatorAttributes,
				Attributes: map[string]string{
					"key1": ".*",
					"key2": ".*",
				},
			},
			workload: workloadinterface.NewWorkloadObj(
				map[string]interface{}{
					"apiVersion": "v1",
					"kind":       "Namespace",
					"metadata": map[string]interface{}{
						"name": "test",
						"annotations": map[string]interface{}{
							"key1": "val1",
							"key2": "val2",
						},
					},
				},
			),
			expected: true,
		},
		{
			name: "Test case: annotations dont match regex",
			designator: &identifiers.PortalDesignator{
				DesignatorType: identifiers.DesignatorAttributes,
				Attributes: map[string]string{
					"key1": "val.*",
					"key2": "val.*",
				},
			},
			workload: workloadinterface.NewWorkloadObj(
				map[string]interface{}{
					"apiVersion": "v1",
					"kind":       "Namespace",
					"metadata": map[string]interface{}{
						"name": "test",
						"annotations": map[string]interface{}{
							"key1": "val1",
							"key2": "bla2",
						},
					},
				},
			),
			expected: false,
		},
		{
			name: "Labels and annotations match in related object",
			designator: &identifiers.PortalDesignator{
				DesignatorType: identifiers.DesignatorAttributes,
				Attributes: map[string]string{
					"key1": "val.*",
					"key2": "val.*",
				},
			},
			workload: objectsenvelopes.NewRegoResponseVectorObject(
				map[string]interface{}{
					"kind": "RegoResponseVector",
					"name": "test",
					"relatedObjects": []interface{}{
						map[string]interface{}{
							"apiVersion": "apps/v1",
							"kind":       "Deployment",
							"metadata": map[string]interface{}{
								"name":      "test",
								"namespace": "test-namespace",
								"labels": map[string]interface{}{
									"key1": "val1",
									"key2": "val2",
								},
							},
						},
						map[string]interface{}{
							"apiVersion": "apps/v1",
							"kind":       "Deployment",
							"metadata": map[string]interface{}{
								"name":      "test-2",
								"namespace": "test-namespace",
								"annotations": map[string]interface{}{
									"key1": "val1",
									"key2": "val2",
								},
							},
						},
					},
				},
			),
			expected: true,
		},
		{
			name: "Test case: Name matches in base object",
			workload: objectsenvelopes.NewRegoResponseVectorObject(
				map[string]interface{}{
					"kind": "RegoResponseVector",
					"name": "base",
					"relatedObjects": []interface{}{
						map[string]interface{}{
							"apiVersion": "apps/v1",
							"kind":       "Deployment",
							"metadata": map[string]interface{}{
								"name":      "test",
								"namespace": "test-namespace",
							},
						},
						map[string]interface{}{
							"apiVersion": "apps/v1",
							"kind":       "Deployment",
							"metadata": map[string]interface{}{
								"name":      "test-2",
								"namespace": "test-namespace",
							},
						},
					},
				},
			),
			designator: &identifiers.PortalDesignator{
				DesignatorType: identifiers.DesignatorAttributes,
				Attributes: map[string]string{
					"name": "base",
				},
			},
			expected: true,
		},
		{
			name: "Test case: Kind matches in base object",
			workload: objectsenvelopes.NewRegoResponseVectorObject(
				map[string]interface{}{
					"kind": "RegoResponseVector",
					"name": "base",
					"relatedObjects": []interface{}{
						map[string]interface{}{
							"apiVersion": "apps/v1",
							"kind":       "Deployment",
							"metadata": map[string]interface{}{
								"name":      "test",
								"namespace": "test-namespace",
							},
						},
						map[string]interface{}{
							"apiVersion": "apps/v1",
							"kind":       "Deployment",
							"metadata": map[string]interface{}{
								"name":      "test-2",
								"namespace": "test-namespace",
							},
						},
					},
				},
			),
			designator: &identifiers.PortalDesignator{
				DesignatorType: identifiers.DesignatorAttributes,
				Attributes: map[string]string{
					"kind": "RegoResponseVector",
				},
			},
			expected: true,
		},
		{
			name: "Test case: Name mismatches in base object",
			workload: objectsenvelopes.NewRegoResponseVectorObject(
				map[string]interface{}{
					"kind": "ServiceAccount",
					"name": "base",
					"relatedObjects": []interface{}{
						map[string]interface{}{
							"apiVersion": "apps/v1",
							"kind":       "Deployment",
							"metadata": map[string]interface{}{
								"name":      "test",
								"namespace": "test-namespace",
							},
						},
						map[string]interface{}{
							"apiVersion": "apps/v1",
							"kind":       "Deployment",
							"metadata": map[string]interface{}{
								"name":      "test-2",
								"namespace": "test-namespace",
							},
						},
					},
				},
			),
			designator: &identifiers.PortalDesignator{
				DesignatorType: identifiers.DesignatorAttributes,
				Attributes: map[string]string{
					"name": "base-2",
				},
			},
			expected: false,
		},
		{
			name: "Test case: Kind mismatches in base object",
			workload: objectsenvelopes.NewRegoResponseVectorObject(
				map[string]interface{}{
					"kind": "ServiceAccount",
					"name": "base",
					"relatedObjects": []interface{}{
						map[string]interface{}{
							"apiVersion": "apps/v1",
							"kind":       "Deployment",
							"metadata": map[string]interface{}{
								"name":      "test",
								"namespace": "test-namespace",
							},
						},
						map[string]interface{}{
							"apiVersion": "apps/v1",
							"kind":       "Deployment",
							"metadata": map[string]interface{}{
								"name":      "test-2",
								"namespace": "test-namespace",
							},
						},
					},
				},
			),
			designator: &identifiers.PortalDesignator{
				DesignatorType: identifiers.DesignatorAttributes,
				Attributes: map[string]string{
					"kind": "RegoResponseVector",
				},
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, processor.hasException(tt.clusterName, tt.designator, tt.workload, nil))
		})
	}
}

func TestProcessor_iterateRegoResponseVector(t *testing.T) {
	p := NewProcessor()

	tests := []struct {
		workload   workloadinterface.IMetadata
		designator *identifiers.PortalDesignator
		name       string
		expected   bool
	}{
		{
			name: "Labels match in one related object",
			workload: objectsenvelopes.NewRegoResponseVectorObject(
				map[string]interface{}{
					"kind": "RegoResponseVector",
					"name": "test",
					"relatedObjects": []interface{}{
						map[string]interface{}{
							"apiVersion": "apps/v1",
							"kind":       "Deployment",
							"metadata": map[string]interface{}{
								"name":      "test",
								"namespace": "test-namespace",
								"labels": map[string]interface{}{
									"app": "test-app",
								},
							},
						},
						map[string]interface{}{
							"apiVersion": "apps/v1",
							"kind":       "Deployment",
							"metadata": map[string]interface{}{
								"name":      "test-2",
								"namespace": "test-namespace",
							},
						},
					},
				},
			),
			designator: &identifiers.PortalDesignator{
				DesignatorType: identifiers.DesignatorAttributes,
				Attributes: map[string]string{
					"app": "test-app",
				},
			},
			expected: true,
		},
		{
			name: "Labels match in one related object and mismatch in another related object",
			workload: objectsenvelopes.NewRegoResponseVectorObject(
				map[string]interface{}{
					"kind": "RegoResponseVector",
					"name": "test",
					"relatedObjects": []interface{}{
						map[string]interface{}{
							"apiVersion": "apps/v1",
							"kind":       "Deployment",
							"metadata": map[string]interface{}{
								"name":      "test",
								"namespace": "test-namespace",
								"labels": map[string]interface{}{
									"app": "different-app",
								},
							},
						},
						map[string]interface{}{
							"apiVersion": "apps/v1",
							"kind":       "Deployment",
							"metadata": map[string]interface{}{
								"name":      "test-2",
								"namespace": "test-namespace",
								"labels": map[string]interface{}{
									"app": "test-app",
								},
							},
						},
					},
				},
			),
			designator: &identifiers.PortalDesignator{
				DesignatorType: identifiers.DesignatorAttributes,
				Attributes: map[string]string{
					"app": "test-app",
				},
			},
			expected: true,
		},
		{
			name: "Annotations match in one related object",
			workload: objectsenvelopes.NewRegoResponseVectorObject(
				map[string]interface{}{
					"kind": "RegoResponseVector",
					"name": "test",
					"relatedObjects": []interface{}{
						map[string]interface{}{
							"apiVersion": "apps/v1",
							"kind":       "Deployment",
							"metadata": map[string]interface{}{
								"name":      "test",
								"namespace": "test-namespace",
								"annotations": map[string]interface{}{
									"app": "test-app",
								},
							},
						},
						map[string]interface{}{
							"apiVersion": "apps/v1",
							"kind":       "Deployment",
							"metadata": map[string]interface{}{
								"name":      "test-2",
								"namespace": "test-namespace",
							},
						},
					},
				},
			),
			designator: &identifiers.PortalDesignator{
				DesignatorType: identifiers.DesignatorAttributes,
				Attributes: map[string]string{
					"app": "test-app",
				},
			},
			expected: true,
		},
		{
			name: "Annotations match in one related object and mismatch in another related object",
			workload: objectsenvelopes.NewRegoResponseVectorObject(
				map[string]interface{}{
					"kind": "RegoResponseVector",
					"name": "test",
					"relatedObjects": []interface{}{
						map[string]interface{}{
							"apiVersion": "apps/v1",
							"kind":       "Deployment",
							"metadata": map[string]interface{}{
								"name":      "test",
								"namespace": "test-namespace",
								"annotations": map[string]interface{}{
									"app": "different-app",
								},
							},
						},
						map[string]interface{}{
							"apiVersion": "apps/v1",
							"kind":       "Deployment",
							"metadata": map[string]interface{}{
								"name":      "test-2",
								"namespace": "test-namespace",
								"annotations": map[string]interface{}{
									"app": "test-app",
								},
							},
						},
					},
				},
			),
			designator: &identifiers.PortalDesignator{
				DesignatorType: identifiers.DesignatorAttributes,
				Attributes: map[string]string{
					"app": "test-app",
				},
			},
			expected: true,
		},
		{
			name: "Labels and Annotations do not match",
			workload: objectsenvelopes.NewRegoResponseVectorObject(
				map[string]interface{}{
					"kind": "RegoResponseVector",
					"name": "test",
					"relatedObjects": []interface{}{
						map[string]interface{}{
							"apiVersion": "apps/v1",
							"kind":       "Deployment",
							"metadata": map[string]interface{}{
								"name":      "test",
								"namespace": "test-namespace",
								"labels": map[string]interface{}{
									"app": "test-app",
								},
								"annotations": map[string]interface{}{
									"app": "test-app",
								},
							},
						},
						map[string]interface{}{
							"apiVersion": "apps/v1",
							"kind":       "Deployment",
							"metadata": map[string]interface{}{
								"name":      "test-2",
								"namespace": "test-namespace",
								"labels": map[string]interface{}{
									"app": "test-app",
								},
								"annotations": map[string]interface{}{
									"app": "test-app",
								},
							},
						},
					},
				},
			),
			designator: &identifiers.PortalDesignator{
				DesignatorType: identifiers.DesignatorAttributes,
				Attributes: map[string]string{
					"app": "different-app",
				},
			},
			expected: false,
		},
		{
			name: "Labels and Annotations are missing in related objects",
			workload: objectsenvelopes.NewRegoResponseVectorObject(
				map[string]interface{}{
					"kind": "RegoResponseVector",
					"name": "test",
					"relatedObjects": []interface{}{
						map[string]interface{}{
							"apiVersion": "apps/v1",
							"kind":       "Deployment",
							"metadata": map[string]interface{}{
								"name":      "test",
								"namespace": "test-namespace",
							},
						},
						map[string]interface{}{
							"apiVersion": "apps/v1",
							"kind":       "Deployment",
							"metadata": map[string]interface{}{
								"name":      "test-2",
								"namespace": "test-namespace",
							},
						},
					},
				},
			),
			designator: &identifiers.PortalDesignator{
				DesignatorType: identifiers.DesignatorAttributes,
				Attributes: map[string]string{
					"app": "test-app",
				},
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, p.iterateRegoResponseVector(tt.workload, tt.designator.DigestPortalDesignator(), nil))
		})
	}
}

func TestMetadataHasException_ContainerName(t *testing.T) {
	p := NewProcessor()

	pod := workloadinterface.NewWorkloadObj(podObject([]string{"app", "sidecar"}, []string{"init-setup"}))

	tests := []struct {
		name       string
		designator *identifiers.PortalDesignator
		expected   bool
	}{
		{
			name: "exception targets matching container",
			designator: &identifiers.PortalDesignator{
				DesignatorType: identifiers.DesignatorAttributes,
				Attributes:     map[string]string{identifiers.AttributeContainerName: "sidecar"},
			},
			expected: true,
		},
		{
			name: "exception targets non-existing container",
			designator: &identifiers.PortalDesignator{
				DesignatorType: identifiers.DesignatorAttributes,
				Attributes:     map[string]string{identifiers.AttributeContainerName: "other"},
			},
			expected: false,
		},
		{
			name: "exception targets init container",
			designator: &identifiers.PortalDesignator{
				DesignatorType: identifiers.DesignatorAttributes,
				Attributes:     map[string]string{identifiers.AttributeContainerName: "init-setup"},
			},
			expected: true,
		},
		{
			name: "containerName with regex wildcard",
			designator: &identifiers.PortalDesignator{
				DesignatorType: identifiers.DesignatorAttributes,
				Attributes:     map[string]string{identifiers.AttributeContainerName: "side.*"},
			},
			expected: true,
		},
		{
			name: "containerName combined with matching namespace",
			designator: &identifiers.PortalDesignator{
				DesignatorType: identifiers.DesignatorAttributes,
				Attributes: map[string]string{
					identifiers.AttributeNamespace:     "default",
					identifiers.AttributeContainerName: "app",
				},
			},
			expected: true,
		},
		{
			name: "containerName combined with non-matching namespace",
			designator: &identifiers.PortalDesignator{
				DesignatorType: identifiers.DesignatorAttributes,
				Attributes: map[string]string{
					identifiers.AttributeNamespace:     "other-ns",
					identifiers.AttributeContainerName: "app",
				},
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			attrs := tt.designator.DigestPortalDesignator()
			assert.Equal(t, tt.expected, p.metadataHasException(pod, attrs, nil))
		})
	}
}

func TestSetRuleResponsExceptions_ContainerNamePrecision(t *testing.T) {
	p := NewProcessor()

	// Pod has two containers: app (index 0) and sidecar (index 1).
	podObj := podObject([]string{"app", "sidecar"}, nil)
	exception := armotypes.PostureExceptionPolicy{
		Resources: []identifiers.PortalDesignator{
			{
				DesignatorType: identifiers.DesignatorAttributes,
				Attributes:     map[string]string{identifiers.AttributeContainerName: "sidecar"},
			},
		},
	}

	tests := []struct {
		name        string
		failedPaths []string
		wantExcept  bool
	}{
		{
			name:        "finding on sidecar (containers[1]) — exception matches",
			failedPaths: []string{"spec.containers[1].securityContext.privileged"},
			wantExcept:  true,
		},
		{
			name:        "finding on app (containers[0]) — exception must NOT apply",
			failedPaths: []string{"spec.containers[0].securityContext.privileged"},
			wantExcept:  false,
		},
		{
			name:        "no failed paths — falls back to workload scan, sidecar found",
			failedPaths: nil,
			wantExcept:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := reporthandling.RuleResponse{
				AlertObject: reporthandling.AlertObject{
					K8SApiObjects: []map[string]interface{}{podObj},
				},
				AssistedRemediation: reporthandling.AssistedRemediation{
					FailedPaths: tt.failedPaths,
				},
			}
			results := []reporthandling.RuleResponse{result}
			p.SetRuleResponsExceptions(results, []armotypes.PostureExceptionPolicy{exception}, "")
			if tt.wantExcept {
				assert.NotNil(t, results[0].Exception, "expected exception to be set")
			} else {
				assert.Nil(t, results[0].Exception, "expected exception NOT to be set")
			}
		})
	}
}

func TestSetRuleResponsExceptions_RegoResponseVector_ContainerNamePrecision(t *testing.T) {
	p := NewProcessor()

	// A RegoResponseVector wrapping a pod that has containers [app (index 0), sidecar (index 1)].
	// The exception targets "sidecar". The FailedPaths point to containers[0] = "app".
	// The exception must NOT be applied because the failing container is "app", not "sidecar".
	//
	// Note: IsTypeRegoResponseVector requires top-level "kind", "name", and "relatedObjects"
	// keys — not the nested metadata.name used by regular Kubernetes objects.
	podObj := podObject([]string{"app", "sidecar"}, nil)
	vector := objectsenvelopes.NewRegoResponseVectorObject(map[string]interface{}{
		"kind":           "RegoResponseVector",
		"name":           "vec",
		"relatedObjects": []interface{}{podObj},
	})
	vectorRaw := vector.GetObject()

	exception := armotypes.PostureExceptionPolicy{
		Resources: []identifiers.PortalDesignator{
			{
				DesignatorType: identifiers.DesignatorAttributes,
				Attributes:     map[string]string{identifiers.AttributeContainerName: "sidecar"},
			},
		},
	}

	tests := []struct {
		name        string
		failedPaths []string
		wantExcept  bool
	}{
		{
			name:        "vector finding on app (containers[0]) — sidecar exception must NOT apply",
			failedPaths: []string{"spec.containers[0].securityContext.privileged"},
			wantExcept:  false,
		},
		{
			name:        "vector finding on sidecar (containers[1]) — sidecar exception must apply",
			failedPaths: []string{"spec.containers[1].securityContext.privileged"},
			wantExcept:  true,
		},
		{
			name:        "no failed paths — falls back to full scan, sidecar found",
			failedPaths: nil,
			wantExcept:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := reporthandling.RuleResponse{
				AlertObject: reporthandling.AlertObject{
					K8SApiObjects: []map[string]interface{}{vectorRaw},
				},
				AssistedRemediation: reporthandling.AssistedRemediation{
					FailedPaths: tt.failedPaths,
				},
			}
			results := []reporthandling.RuleResponse{result}
			p.SetRuleResponsExceptions(results, []armotypes.PostureExceptionPolicy{exception}, "")
			if tt.wantExcept {
				assert.NotNil(t, results[0].Exception, "expected exception to be set")
			} else {
				assert.Nil(t, results[0].Exception, "expected exception NOT to be set")
			}
		})
	}
}

func TestSetRuleResponsExceptions_ExternalObjects_RegoResponseVector(t *testing.T) {
	p := NewProcessor()

	// Same scenario as the K8SApiObjects test but delivered via AlertObject.ExternalObjects
	// (a single map[string]interface{}, not a list).
	// The vector wraps a pod with containers [app (0), sidecar (1)].
	// Exception targets "sidecar"; FailedPaths point at containers[0] = "app".
	podObj := podObject([]string{"app", "sidecar"}, nil)
	vectorRaw := map[string]interface{}{
		"kind":           "RegoResponseVector",
		"name":           "vec",
		"relatedObjects": []interface{}{podObj},
	}

	exception := armotypes.PostureExceptionPolicy{
		Resources: []identifiers.PortalDesignator{
			{
				DesignatorType: identifiers.DesignatorAttributes,
				Attributes:     map[string]string{identifiers.AttributeContainerName: "sidecar"},
			},
		},
	}

	tests := []struct {
		name        string
		failedPaths []string
		wantExcept  bool
	}{
		{
			name:        "ExternalObjects vector: finding on app — sidecar exception must NOT apply",
			failedPaths: []string{"spec.containers[0].securityContext.privileged"},
			wantExcept:  false,
		},
		{
			name:        "ExternalObjects vector: finding on sidecar — exception must apply",
			failedPaths: []string{"spec.containers[1].securityContext.privileged"},
			wantExcept:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := reporthandling.RuleResponse{
				AlertObject: reporthandling.AlertObject{
					ExternalObjects: vectorRaw,
				},
				AssistedRemediation: reporthandling.AssistedRemediation{
					FailedPaths: tt.failedPaths,
				},
			}
			results := []reporthandling.RuleResponse{result}
			p.SetRuleResponsExceptions(results, []armotypes.PostureExceptionPolicy{exception}, "")
			if tt.wantExcept {
				assert.NotNil(t, results[0].Exception, "expected exception to be set")
			} else {
				assert.Nil(t, results[0].Exception, "expected exception NOT to be set")
			}
		})
	}
}

func TestHasException_RegoResponseVector_ContainerNameFallbackBlocked(t *testing.T) {
	p := NewProcessor()

	// A RegoResponseVector whose base name is "base-subject".
	// The related object has containers [app], none named "missing".
	vector := objectsenvelopes.NewRegoResponseVectorObject(map[string]interface{}{
		"kind": "RegoResponseVector",
		"name": "base-subject",
		"relatedObjects": []interface{}{
			map[string]interface{}{
				"apiVersion": "v1",
				"kind":       "Pod",
				"metadata":   map[string]interface{}{"name": "base-subject"},
				"spec": map[string]interface{}{
					"containers": []interface{}{
						map[string]interface{}{"name": "app"},
					},
				},
			},
		},
	})

	designator := &identifiers.PortalDesignator{
		DesignatorType: identifiers.DesignatorAttributes,
		Attributes: map[string]string{
			identifiers.AttributeName:          "base-subject",
			identifiers.AttributeContainerName: "missing",
		},
	}

	attrs := designator.DigestPortalDesignator()
	// The base vector's name matches, but "missing" is not a container —
	// hasException must return false, not bypass the container check.
	result := p.hasException("", designator, vector, nil)
	_ = attrs
	assert.False(t, result, "containerName on a RegoResponseVector must not fall back to base-object name match")
}

// postureObjectSelectorExceptionMock builds an exception scoped to namespace
// "default" via a designator (so the designator axis always matches the test
// workloads) and carries the given ObjectSelector as the additional label axis.
func postureObjectSelectorExceptionMock(sel *armotypes.LabelSelector) *armotypes.PostureExceptionPolicy {
	return &armotypes.PostureExceptionPolicy{
		PortalBase: armotypes.PortalBase{Name: "postureObjectSelectorExceptionMock"},
		PolicyType: "postureExceptionPolicy",
		Actions:    []armotypes.PostureExceptionPolicyActions{armotypes.AlertOnly},
		Resources: []identifiers.PortalDesignator{
			{
				DesignatorType: identifiers.DesignatorAttributes,
				Attributes: map[string]string{
					identifiers.AttributeNamespace: "default",
				},
			},
		},
		ObjectSelector:  sel,
		PosturePolicies: []armotypes.PosturePolicy{{FrameworkName: "MIT.*"}},
	}
}

// TestGetResourceExceptions_ObjectSelector pins the ObjectSelector axis added for
// SecurityException spec.match.objectSelector. The selector is ANDed with the
// designator (namespace) match, and supports the full label-selector grammar —
// including the set-based operators that the regex-designator path cannot express.
func TestGetResourceExceptions_ObjectSelector(t *testing.T) {
	labeledObj, err := workloadinterface.NewBaseObjBytes([]byte(`{"apiVersion":"v1","kind":"Deployment","metadata":{"name":"web","namespace":"default","labels":{"app":"web","env":"prod"}}}`))
	require.NoError(t, err)

	noLabelsObj, err := workloadinterface.NewBaseObjBytes([]byte(`{"apiVersion":"v1","kind":"Deployment","metadata":{"name":"web","namespace":"default"}}`))
	require.NoError(t, err)

	p := NewProcessor()

	testCases := []struct {
		selector                *armotypes.LabelSelector
		workloadObj             workloadinterface.IMetadata
		desc                    string
		expectedExceptionsCount int
	}{
		{
			desc:                    "nil selector imposes no label constraint (designator decides)",
			selector:                nil,
			workloadObj:             labeledObj,
			expectedExceptionsCount: 1,
		},
		{
			// The empty-selector trap: empty must mean "no constraint", never
			// promoted to labels.Everything() (which would also match noLabelsObj
			// here) — but it must also not collapse to labels.Nothing().
			desc:                    "empty selector imposes no label constraint (designator decides)",
			selector:                &armotypes.LabelSelector{},
			workloadObj:             labeledObj,
			expectedExceptionsCount: 1,
		},
		{
			desc:                    "matchLabels matches",
			selector:                &armotypes.LabelSelector{MatchLabels: map[string]string{"app": "web"}},
			workloadObj:             labeledObj,
			expectedExceptionsCount: 1,
		},
		{
			desc:                    "matchLabels does not match",
			selector:                &armotypes.LabelSelector{MatchLabels: map[string]string{"app": "api"}},
			workloadObj:             labeledObj,
			expectedExceptionsCount: 0,
		},
		{
			desc: "matchExpressions In matches",
			selector: &armotypes.LabelSelector{MatchExpressions: []armotypes.LabelSelectorRequirement{
				{Key: "env", Operator: metav1.LabelSelectorOpIn, Values: []string{"prod", "staging"}},
			}},
			workloadObj:             labeledObj,
			expectedExceptionsCount: 1,
		},
		{
			desc: "matchExpressions NotIn excludes",
			selector: &armotypes.LabelSelector{MatchExpressions: []armotypes.LabelSelectorRequirement{
				{Key: "env", Operator: metav1.LabelSelectorOpNotIn, Values: []string{"prod"}},
			}},
			workloadObj:             labeledObj,
			expectedExceptionsCount: 0,
		},
		{
			desc: "matchExpressions Exists matches",
			selector: &armotypes.LabelSelector{MatchExpressions: []armotypes.LabelSelectorRequirement{
				{Key: "app", Operator: metav1.LabelSelectorOpExists},
			}},
			workloadObj:             labeledObj,
			expectedExceptionsCount: 1,
		},
		{
			desc: "matchExpressions DoesNotExist excludes when key present",
			selector: &armotypes.LabelSelector{MatchExpressions: []armotypes.LabelSelectorRequirement{
				{Key: "app", Operator: metav1.LabelSelectorOpDoesNotExist},
			}},
			workloadObj:             labeledObj,
			expectedExceptionsCount: 0,
		},
		{
			desc:                    "non-empty selector against a workload with no labels matches nothing",
			selector:                &armotypes.LabelSelector{MatchLabels: map[string]string{"app": "web"}},
			workloadObj:             noLabelsObj,
			expectedExceptionsCount: 0,
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			ex := postureObjectSelectorExceptionMock(test.selector)
			res := p.GetResourceExceptions([]armotypes.PostureExceptionPolicy{*ex}, test.workloadObj, "test")
			assert.Equal(t, test.expectedExceptionsCount, len(res))
		})
	}
}
