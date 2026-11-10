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

func TestListRuleExceptionsPreservesMatchingPolicyTuples(t *testing.T) {
	p := NewProcessor()
	exception := armotypes.PostureExceptionPolicy{
		PosturePolicies: []armotypes.PosturePolicy{
			{FrameworkName: "NSA", ControlID: "C-0034", RuleName: "R1"},
			{FrameworkName: "MITRE", ControlID: "C-0034", RuleName: "R2"},
		},
	}

	matched := p.ListRuleExceptions([]armotypes.PostureExceptionPolicy{exception}, "", "", "R1")

	require.Len(t, matched, 1)
	require.Len(t, matched[0].PosturePolicies, 1)
	assert.Equal(t, armotypes.PosturePolicy{
		FrameworkName: "NSA",
		ControlID:     "C-0034",
		RuleName:      "R1",
	}, matched[0].PosturePolicies[0])
	assert.Len(t, exception.PosturePolicies, 2, "filtering must not mutate the caller-owned exception")
}

func TestListRuleExceptionsPreservesGlobalPolicyRepresentations(t *testing.T) {
	p := NewProcessor()
	tests := []struct {
		name     string
		policies []armotypes.PosturePolicy
	}{
		{name: "nil posture policies", policies: nil},
		{name: "empty posture policies", policies: []armotypes.PosturePolicy{}},
		{name: "one empty posture policy", policies: []armotypes.PosturePolicy{{}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exception := armotypes.PostureExceptionPolicy{PosturePolicies: tt.policies}
			matched := p.ListRuleExceptions([]armotypes.PostureExceptionPolicy{exception}, "MITRE", "C-0034", "R1")
			require.Len(t, matched, 1)
			assert.Equal(t, tt.policies, matched[0].PosturePolicies)
		})
	}
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

	// alternation is anchored as a whole
	assert.True(t, c.compareCluster("prod|staging", "staging"))
	assert.False(t, c.compareCluster("prod|staging", "prod-eu"))
	assert.False(t, c.compareCluster("prod|staging", "pre-staging"))
	assert.True(t, c.regexCompareI("NSA|MITRE", "mitre"))
	assert.False(t, c.regexCompareI("NSA|MITRE", "nsa-custom"))

	// a bare "*" keeps working as a match-all wildcard
	assert.True(t, c.compareCluster("*", "bez-minikube-25-10"))
	assert.True(t, c.regexCompareI("*", "C-0001"))
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
			assert.Equal(t, tt.expected, processor.hasException(tt.clusterName, tt.designator, tt.workload, nil, nil))
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
			assert.Equal(t, tt.expected, p.iterateRegoResponseVector(tt.workload, tt.designator.DigestPortalDesignator(), nil, nil))
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
			assert.Equal(t, tt.expected, p.metadataHasException(pod, attrs, nil, nil))
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
	result := p.hasException("", designator, vector, nil, nil)
	_ = attrs
	assert.False(t, result, "containerName on a RegoResponseVector must not fall back to base-object name match")
}

// postureObjectSelectorExceptionMock builds an exception scoped by the given
// designator attributes (so the designator axis is controllable) and carries the
// given ObjectSelector as the additional label axis.
func postureObjectSelectorExceptionMock(designatorAttrs map[string]string, sel *armotypes.LabelSelector) *armotypes.PostureExceptionPolicy {
	return &armotypes.PostureExceptionPolicy{
		PortalBase: armotypes.PortalBase{Name: "postureObjectSelectorExceptionMock"},
		PolicyType: "postureExceptionPolicy",
		Actions:    []armotypes.PostureExceptionPolicyActions{armotypes.AlertOnly},
		Resources: []identifiers.PortalDesignator{
			{
				DesignatorType: identifiers.DesignatorAttributes,
				Attributes:     designatorAttrs,
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
		{
			// A malformed selector (In with no values is invalid) must match nothing,
			// not degrade into match-all even though the designator (namespace) matches.
			desc: "malformed selector matches nothing",
			selector: &armotypes.LabelSelector{MatchExpressions: []armotypes.LabelSelectorRequirement{
				{Key: "env", Operator: metav1.LabelSelectorOpIn, Values: nil},
			}},
			workloadObj:             labeledObj,
			expectedExceptionsCount: 0,
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			ex := postureObjectSelectorExceptionMock(map[string]string{identifiers.AttributeNamespace: "default"}, test.selector)
			res := p.GetResourceExceptions([]armotypes.PostureExceptionPolicy{*ex}, test.workloadObj, "test")
			assert.Equal(t, test.expectedExceptionsCount, len(res))
		})
	}
}

// TestGetResourceExceptions_ApiGroup exercises apiGroup through the public entry point,
// covering the distinction between an absent key and one explicitly set to the empty string.
// An absent apiGroup imposes no constraint; an explicit "" names the core group and has to
// exclude resources from named groups, rather than being read as "no constraint".
func TestGetResourceExceptions_ApiGroup(t *testing.T) {
	p := NewProcessor()

	pod := workloadinterface.NewWorkloadObj(podObject([]string{"app"}, nil))
	deployment := workloadinterface.NewWorkloadObj(deploymentObject("apps/v1", nil))

	testCases := []struct {
		desc                    string
		attributes              map[string]string
		workloadObj             workloadinterface.IMetadata
		expectedExceptionsCount int
	}{
		{
			desc:                    "explicit empty apiGroup matches a core group resource",
			attributes:              map[string]string{identifiers.AttributeApiGroup: ""},
			workloadObj:             pod,
			expectedExceptionsCount: 1,
		},
		{
			desc:                    "explicit empty apiGroup does not match an apps/v1 resource",
			attributes:              map[string]string{identifiers.AttributeApiGroup: ""},
			workloadObj:             deployment,
			expectedExceptionsCount: 0,
		},
		{
			desc:                    "named apiGroup matches its own group",
			attributes:              map[string]string{identifiers.AttributeApiGroup: "apps"},
			workloadObj:             deployment,
			expectedExceptionsCount: 1,
		},
		{
			desc:                    "named apiGroup does not match a core group resource",
			attributes:              map[string]string{identifiers.AttributeApiGroup: "apps"},
			workloadObj:             pod,
			expectedExceptionsCount: 0,
		},
		{
			desc:                    "an absent apiGroup imposes no constraint",
			attributes:              map[string]string{identifiers.AttributeNamespace: "default"},
			workloadObj:             deployment,
			expectedExceptionsCount: 1,
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			ex := postureObjectSelectorExceptionMock(test.attributes, nil)
			res := p.GetResourceExceptions([]armotypes.PostureExceptionPolicy{*ex}, test.workloadObj, "test")
			assert.Equal(t, test.expectedExceptionsCount, len(res))
		})
	}
}

// TestGetResourceExceptions_RegexAlternation pins that a designator pattern with
// alternation has to match the whole value. Anchoring "a|b" as "^a|b$" would accept
// any value that starts with a or ends with b, so an exception written for two
// namespaces or names would also silence findings on unrelated resources.
func TestGetResourceExceptions_RegexAlternation(t *testing.T) {
	p := NewProcessor()

	workloadIn := func(name, namespace string) workloadinterface.IMetadata {
		return workloadinterface.NewWorkloadObj(map[string]interface{}{
			"apiVersion": "apps/v1",
			"kind":       "Deployment",
			"metadata":   map[string]interface{}{"name": name, "namespace": namespace},
		})
	}

	testCases := []struct {
		desc                    string
		attributes              map[string]string
		workloadObj             workloadinterface.IMetadata
		expectedExceptionsCount int
	}{
		{
			desc:                    "namespace alternation matches its first alternative",
			attributes:              map[string]string{identifiers.AttributeNamespace: "kube-system|kube-public"},
			workloadObj:             workloadIn("coredns", "kube-system"),
			expectedExceptionsCount: 1,
		},
		{
			desc:                    "namespace alternation matches its last alternative",
			attributes:              map[string]string{identifiers.AttributeNamespace: "kube-system|kube-public"},
			workloadObj:             workloadIn("coredns", "kube-public"),
			expectedExceptionsCount: 1,
		},
		{
			desc:                    "namespace alternation does not match a namespace that only starts with the first alternative",
			attributes:              map[string]string{identifiers.AttributeNamespace: "kube-system|kube-public"},
			workloadObj:             workloadIn("coredns", "kube-system-tenant"),
			expectedExceptionsCount: 0,
		},
		{
			desc:                    "namespace alternation does not match a namespace that only ends with the last alternative",
			attributes:              map[string]string{identifiers.AttributeNamespace: "kube-system|kube-public"},
			workloadObj:             workloadIn("coredns", "not-kube-public"),
			expectedExceptionsCount: 0,
		},
		{
			desc:                    "name alternation does not match a longer name",
			attributes:              map[string]string{identifiers.AttributeName: "nginx|redis"},
			workloadObj:             workloadIn("nginx-backdoor", "default"),
			expectedExceptionsCount: 0,
		},
		{
			desc:                    "name alternation matches an exact name",
			attributes:              map[string]string{identifiers.AttributeName: "nginx|redis"},
			workloadObj:             workloadIn("redis", "default"),
			expectedExceptionsCount: 1,
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			ex := postureObjectSelectorExceptionMock(test.attributes, nil)
			res := p.GetResourceExceptions([]armotypes.PostureExceptionPolicy{*ex}, test.workloadObj, "test")
			assert.Equal(t, test.expectedExceptionsCount, len(res))
		})
	}
}

// TestListRuleExceptionsRegexAlternation pins whole-value matching for the
// case-insensitive policy patterns (framework, control ID, rule name).
func TestListRuleExceptionsRegexAlternation(t *testing.T) {
	p := NewProcessor()
	exception := armotypes.PostureExceptionPolicy{
		PosturePolicies: []armotypes.PosturePolicy{{ControlID: "C-0001|C-0002"}},
	}

	assert.Len(t, p.ListRuleExceptions([]armotypes.PostureExceptionPolicy{exception}, "", "c-0001", ""), 1)
	assert.Len(t, p.ListRuleExceptions([]armotypes.PostureExceptionPolicy{exception}, "", "C-0002", ""), 1)
	assert.Empty(t, p.ListRuleExceptions([]armotypes.PostureExceptionPolicy{exception}, "", "C-00012", ""))
	assert.Empty(t, p.ListRuleExceptions([]armotypes.PostureExceptionPolicy{exception}, "", "C-10002", ""))

	assert.True(t, p.RegexCompareControlID("C-0001|C-0002", "C-0002"))
	assert.False(t, p.RegexCompareControlID("C-0001|C-0002", "C-00015"))
}

// TestGetResourceExceptions_ObjectSelector_RegoResponseVector pins that the
// objectSelector is evaluated against the *related* workload of a RegoResponseVector
// (the real object), not the label-less vector envelope, and that the selector and
// designator must be satisfied by the *same* related object (no cross-object match).
func TestGetResourceExceptions_ObjectSelector_RegoResponseVector(t *testing.T) {
	p := NewProcessor()

	// Vector whose single related object is a labeled Deployment in "default".
	vectorMatching := objectsenvelopes.NewRegoResponseVectorObject(map[string]interface{}{
		"kind": "RegoResponseVector",
		"name": "web",
		"relatedObjects": []interface{}{
			map[string]interface{}{
				"apiVersion": "apps/v1",
				"kind":       "Deployment",
				"metadata": map[string]interface{}{
					"name":      "web",
					"namespace": "default",
					"labels":    map[string]interface{}{"app": "web", "env": "prod"},
				},
			},
		},
	})

	// Same, but the related workload carries a non-matching label.
	vectorNonMatching := objectsenvelopes.NewRegoResponseVectorObject(map[string]interface{}{
		"kind": "RegoResponseVector",
		"name": "web",
		"relatedObjects": []interface{}{
			map[string]interface{}{
				"apiVersion": "apps/v1",
				"kind":       "Deployment",
				"metadata": map[string]interface{}{
					"name":      "web",
					"namespace": "default",
					"labels":    map[string]interface{}{"app": "api"},
				},
			},
		},
	})

	// Two related objects: the Deployment matches the (kind-scoped) designator but not
	// the selector; the Service matches the selector but not the designator.
	vectorCrossObject := objectsenvelopes.NewRegoResponseVectorObject(map[string]interface{}{
		"kind": "RegoResponseVector",
		"name": "web",
		"relatedObjects": []interface{}{
			map[string]interface{}{
				"apiVersion": "apps/v1",
				"kind":       "Deployment",
				"metadata": map[string]interface{}{
					"name":      "web",
					"namespace": "default",
					"labels":    map[string]interface{}{"app": "api"},
				},
			},
			map[string]interface{}{
				"apiVersion": "v1",
				"kind":       "Service",
				"metadata": map[string]interface{}{
					"name":      "web",
					"namespace": "default",
					"labels":    map[string]interface{}{"app": "web"},
				},
			},
		},
	})

	testCases := []struct {
		workloadObj             workloadinterface.IMetadata
		designatorAttrs         map[string]string
		desc                    string
		expectedExceptionsCount int
	}{
		{
			desc:                    "selector matches the related workload of a vector",
			designatorAttrs:         map[string]string{identifiers.AttributeNamespace: "default"},
			workloadObj:             vectorMatching,
			expectedExceptionsCount: 1,
		},
		{
			desc:                    "selector does not match the related workload of a vector",
			designatorAttrs:         map[string]string{identifiers.AttributeNamespace: "default"},
			workloadObj:             vectorNonMatching,
			expectedExceptionsCount: 0,
		},
		{
			desc:                    "selector and designator satisfied by different related objects does not match",
			designatorAttrs:         map[string]string{identifiers.AttributeNamespace: "default", identifiers.AttributeKind: "Deployment"},
			workloadObj:             vectorCrossObject,
			expectedExceptionsCount: 0,
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			ex := postureObjectSelectorExceptionMock(test.designatorAttrs, &armotypes.LabelSelector{MatchLabels: map[string]string{"app": "web"}})
			res := p.GetResourceExceptions([]armotypes.PostureExceptionPolicy{*ex}, test.workloadObj, "test")
			assert.Equal(t, test.expectedExceptionsCount, len(res))
		})
	}
}

// TestGetResourceExceptions_ObjectSelectorAndContainerName pins that objectSelector and a
// containerName designator are ANDed: the exception applies only when the selector matches
// the workload's labels AND the named container is one of the workload's containers. The
// selector is evaluated (in metadataHasException) before the containerName check, so a
// selector miss short-circuits regardless of the container match.
func TestGetResourceExceptions_ObjectSelectorAndContainerName(t *testing.T) {
	labeledPod, err := workloadinterface.NewBaseObjBytes([]byte(`{"apiVersion":"v1","kind":"Pod","metadata":{"name":"web","namespace":"default","labels":{"app":"web"}},"spec":{"containers":[{"name":"app"},{"name":"sidecar"}]}}`))
	require.NoError(t, err)

	p := NewProcessor()

	testCases := []struct {
		selector                *armotypes.LabelSelector
		desc                    string
		containerName           string
		expectedExceptionsCount int
	}{
		{
			desc:                    "selector matches and container matches",
			selector:                &armotypes.LabelSelector{MatchLabels: map[string]string{"app": "web"}},
			containerName:           "sidecar",
			expectedExceptionsCount: 1,
		},
		{
			desc:                    "selector matches but container does not",
			selector:                &armotypes.LabelSelector{MatchLabels: map[string]string{"app": "web"}},
			containerName:           "missing",
			expectedExceptionsCount: 0,
		},
		{
			desc:                    "container matches but selector does not",
			selector:                &armotypes.LabelSelector{MatchLabels: map[string]string{"app": "api"}},
			containerName:           "sidecar",
			expectedExceptionsCount: 0,
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			ex := postureObjectSelectorExceptionMock(map[string]string{identifiers.AttributeContainerName: test.containerName}, test.selector)
			res := p.GetResourceExceptions([]armotypes.PostureExceptionPolicy{*ex}, labeledPod, "test")
			assert.Equal(t, test.expectedExceptionsCount, len(res))
		})
	}
}

// TestHasException_RegoResponseVector_BaseFallthroughWithSelector locks the interaction
// between a non-empty objectSelector and the RegoResponseVector base-object fall-through.
// When a designator matches only the base envelope (not any related object), a non-empty
// selector must still fail-closed: the label-less envelope cannot satisfy it, and a selector
// match on a related object must NOT cross-combine with a designator match on the base.
func TestHasException_RegoResponseVector_BaseFallthroughWithSelector(t *testing.T) {
	p := NewProcessor()

	// Vector base name "vec"; the single related object is a Deployment named "web"
	// carrying app=web, so it never matches the name=vec designator.
	vector := objectsenvelopes.NewRegoResponseVectorObject(map[string]interface{}{
		"kind": "RegoResponseVector",
		"name": "vec",
		"relatedObjects": []interface{}{
			map[string]interface{}{
				"apiVersion": "apps/v1",
				"kind":       "Deployment",
				"metadata": map[string]interface{}{
					"name":      "web",
					"namespace": "default",
					"labels":    map[string]interface{}{"app": "web"},
				},
			},
		},
	})

	designator := &identifiers.PortalDesignator{
		DesignatorType: identifiers.DesignatorAttributes,
		Attributes:     map[string]string{identifiers.AttributeName: "vec"}, // matches only the base envelope
	}

	// No selector: the designator matches the base envelope, so the exception applies.
	noSelector, ok := parseObjectSelector(&armotypes.PostureExceptionPolicy{})
	require.True(t, ok)
	assert.True(t, p.hasException("", designator, vector, nil, noSelector),
		"name matching the base envelope with no selector must apply")

	// Non-empty selector matching the related object (app=web): the base envelope has no
	// labels, so the selector cannot match there, and the related object does not match the
	// name=vec designator — the selector and designator are never satisfied by one object.
	selector, ok := parseObjectSelector(&armotypes.PostureExceptionPolicy{
		ObjectSelector: &armotypes.LabelSelector{MatchLabels: map[string]string{"app": "web"}},
	})
	require.True(t, ok)
	assert.False(t, p.hasException("", designator, vector, nil, selector),
		"objectSelector must not cross-combine a related-object label match with a base-envelope designator match")
}

// TestHasException_RegoResponseVector_BaseFallthroughNegativeSelector pins the
// fail-closed behavior for negative label-selector operators (DoesNotExist / NotIn).
// Unlike a positive selector, a negative requirement matches an *empty* label set, so
// it would spuriously pass against the label-less base envelope on fall-through. The
// only real workload (a related object carrying env=prod) does not satisfy the
// selector, so the finding must NOT be suppressed — the exception must not apply.
func TestHasException_RegoResponseVector_BaseFallthroughNegativeSelector(t *testing.T) {
	p := NewProcessor()

	// Vector base name "web"; the single related object is a Deployment named "web"
	// carrying env=prod. A name=web designator matches both the base envelope and the
	// related object, so only the selector should keep the exception from applying.
	vector := objectsenvelopes.NewRegoResponseVectorObject(map[string]interface{}{
		"kind": "RegoResponseVector",
		"name": "web",
		"relatedObjects": []interface{}{
			map[string]interface{}{
				"apiVersion": "apps/v1",
				"kind":       "Deployment",
				"metadata": map[string]interface{}{
					"name":      "web",
					"namespace": "default",
					"labels":    map[string]interface{}{"env": "prod"},
				},
			},
		},
	})

	designator := &identifiers.PortalDesignator{
		DesignatorType: identifiers.DesignatorAttributes,
		Attributes:     map[string]string{identifiers.AttributeName: "web"},
	}

	testCases := []struct {
		selector *armotypes.LabelSelector
		desc     string
	}{
		{
			desc: "DoesNotExist over a related workload that has the key",
			selector: &armotypes.LabelSelector{MatchExpressions: []armotypes.LabelSelectorRequirement{
				{Key: "env", Operator: metav1.LabelSelectorOpDoesNotExist},
			}},
		},
		{
			desc: "NotIn over a related workload whose value is excluded",
			selector: &armotypes.LabelSelector{MatchExpressions: []armotypes.LabelSelectorRequirement{
				{Key: "env", Operator: metav1.LabelSelectorOpNotIn, Values: []string{"prod"}},
			}},
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			selector, ok := parseObjectSelector(&armotypes.PostureExceptionPolicy{ObjectSelector: test.selector})
			require.True(t, ok)
			assert.False(t, p.hasException("", designator, vector, nil, selector),
				"a negative objectSelector must not match the label-less base envelope on fall-through")
		})
	}
}

// TestGetResourceExceptions_EmptyResourcesMatchesEverywhere pins issue-3368
// (kubescape/kubescape#3368): an exception with no Resources at all must apply
// with no scope constraint - the same "no resources = no scope constraint"
// convention kubescape already documents and relies on for manual controls
// (hasExplicitControlException) - instead of matching nothing, which is what
// getResourceExceptions did before this fix (its designator loop has zero
// iterations when Resources is empty).
func TestGetResourceExceptions_EmptyResourcesMatchesEverywhere(t *testing.T) {
	p := NewProcessor()

	anyWorkload, err := workloadinterface.NewBaseObjBytes([]byte(`{"apiVersion":"v1","kind":"Deployment","metadata":{"name":"anything","namespace":"anywhere"}}`))
	require.NoError(t, err)

	scopeLessException := armotypes.PostureExceptionPolicy{
		PortalBase:      armotypes.PortalBase{Name: "scope-less"},
		PolicyType:      "postureExceptionPolicy",
		PosturePolicies: []armotypes.PosturePolicy{{FrameworkName: "MITRE"}},
		// Resources deliberately nil.
	}

	res := p.GetResourceExceptions([]armotypes.PostureExceptionPolicy{scopeLessException}, anyWorkload, "any-cluster")
	require.Len(t, res, 1, "a scope-less exception (no Resources) must match an arbitrary workload on an arbitrary cluster")
}

// TestGetResourceExceptions_EmptyResourcesWithObjectSelectorStillConstrains
// confirms the fix is additive, not a blanket bypass: when a scope-less
// exception also carries an ObjectSelector, that selector still constrains
// which workloads it applies to - only the designator axis (cluster/
// namespace/name/kind/labels/path) is left unconstrained by an empty
// Resources list.
func TestGetResourceExceptions_EmptyResourcesWithObjectSelectorStillConstrains(t *testing.T) {
	p := NewProcessor()

	labeledWorkload, err := workloadinterface.NewBaseObjBytes([]byte(`{"apiVersion":"v1","kind":"Deployment","metadata":{"name":"web","namespace":"default","labels":{"app":"web"}}}`))
	require.NoError(t, err)
	unlabeledWorkload, err := workloadinterface.NewBaseObjBytes([]byte(`{"apiVersion":"v1","kind":"Deployment","metadata":{"name":"web","namespace":"default"}}`))
	require.NoError(t, err)

	scopeLessWithSelector := armotypes.PostureExceptionPolicy{
		PortalBase:      armotypes.PortalBase{Name: "scope-less-with-selector"},
		PolicyType:      "postureExceptionPolicy",
		PosturePolicies: []armotypes.PosturePolicy{{FrameworkName: "MITRE"}},
		ObjectSelector:  &armotypes.LabelSelector{MatchLabels: map[string]string{"app": "web"}},
		// Resources deliberately nil.
	}

	matched := p.GetResourceExceptions([]armotypes.PostureExceptionPolicy{scopeLessWithSelector}, labeledWorkload, "any-cluster")
	assert.Len(t, matched, 1, "the workload carries the selected label, so the scope-less exception should still apply")

	unmatched := p.GetResourceExceptions([]armotypes.PostureExceptionPolicy{scopeLessWithSelector}, unlabeledWorkload, "any-cluster")
	assert.Empty(t, unmatched, "the workload lacks the selected label, so a scope-less exception must not blanket-match it")
}

// TestGetResourceExceptions_EmptyResourcesDifferentControlDoesNotMatch confirms
// the scope-less match is still gated by ListRuleExceptions/PosturePolicies
// upstream in the normal pipeline - GetResourceExceptions itself does not
// filter by control, but this documents that a scope-less exception is not
// some universal always-true policy independent of what selected it.
func TestGetResourceExceptions_EmptyResourcesDifferentControlDoesNotMatch(t *testing.T) {
	p := NewProcessor()

	workload, err := workloadinterface.NewBaseObjBytes([]byte(`{"apiVersion":"v1","kind":"Deployment","metadata":{"name":"web","namespace":"default"}}`))
	require.NoError(t, err)

	scopeLessException := armotypes.PostureExceptionPolicy{
		PortalBase:      armotypes.PortalBase{Name: "scope-less-for-C-0034"},
		PolicyType:      "postureExceptionPolicy",
		PosturePolicies: []armotypes.PosturePolicy{{ControlID: "C-0034"}},
	}

	// ListRuleExceptions is the actual PosturePolicies gate; GetResourceExceptions
	// receives only what already survived it, matching the real pipeline
	// (SetRuleExceptions -> ListRuleExceptions -> SetRuleResponsExceptions -> getResourceExceptions).
	forThisControl := p.ListRuleExceptions([]armotypes.PostureExceptionPolicy{scopeLessException}, "", "C-0034", "")
	require.Len(t, forThisControl, 1)
	matched := p.GetResourceExceptions(forThisControl, workload, "any-cluster")
	assert.Len(t, matched, 1)

	forOtherControl := p.ListRuleExceptions([]armotypes.PostureExceptionPolicy{scopeLessException}, "", "C-0035", "")
	assert.Empty(t, forOtherControl, "a scope-less exception still only targets the control(s) its PosturePolicies name")
}

// TestSetRuleResponsExceptions_EmptyResourcesMatchesEverywhere exercises the
// fix through SetRuleResponsExceptions (the entry point SetRuleExceptions ->
// SetControlExceptions -> SetFrameworkExceptions actually use), not just the
// lower-level GetResourceExceptions, since issue-3368 named both entry points.
func TestSetRuleResponsExceptions_EmptyResourcesMatchesEverywhere(t *testing.T) {
	p := NewProcessor()

	workloadObj, err := workloadinterface.NewBaseObjBytes([]byte(`{"apiVersion":"v1","kind":"Deployment","metadata":{"name":"web","namespace":"default"}}`))
	require.NoError(t, err)

	scopeLessException := armotypes.PostureExceptionPolicy{
		PortalBase: armotypes.PortalBase{Name: "scope-less"},
		PolicyType: "postureExceptionPolicy",
	}

	results := []reporthandling.RuleResponse{
		{
			AlertObject: reporthandling.AlertObject{
				K8SApiObjects: []map[string]interface{}{workloadObj.GetObject()},
			},
		},
	}

	p.SetRuleResponsExceptions(results, []armotypes.PostureExceptionPolicy{scopeLessException}, "any-cluster")

	require.NotNil(t, results[0].Exception, "a scope-less exception must still be attached via SetRuleResponsExceptions")
	assert.Equal(t, "scope-less", results[0].Exception.GetName())
}

// TestMetadataHasException_ApiGroup covers a SecurityException scoped by apiGroup.
//
// armoapi-go's DigestAttributesDesignator has no case for apiGroup, so it arrives in the
// labels map. Comparing it as a Kubernetes label can only fail, since apiGroup is a schema
// field and no resource carries a label by that name, which left every designator using it
// unmatchable. Matching it explicitly has to both restore the match and keep apiGroup
// constraining, rather than dropping it and widening the exception to every group.
func TestMetadataHasException_ApiGroup(t *testing.T) {
	p := NewProcessor()

	deployment := workloadinterface.NewWorkloadObj(deploymentObject("apps/v1", map[string]string{"app": "test-app"}))
	pod := workloadinterface.NewWorkloadObj(podObject([]string{"app"}, nil))

	tests := []struct {
		name       string
		workload   workloadinterface.IMetadata
		attributes map[string]string
		expected   bool
	}{
		{
			name:       "apiGroup alone matches",
			workload:   deployment,
			attributes: map[string]string{identifiers.AttributeApiGroup: "apps"},
			expected:   true,
		},
		{
			name:     "apiGroup with kind matches",
			workload: deployment,
			attributes: map[string]string{
				identifiers.AttributeApiGroup: "apps",
				identifiers.AttributeKind:     "Deployment",
			},
			expected: true,
		},
		{
			name:     "a different apiGroup still constrains a matching kind",
			workload: deployment,
			attributes: map[string]string{
				identifiers.AttributeApiGroup: "batch",
				identifiers.AttributeKind:     "Deployment",
			},
			expected: false,
		},
		{
			name:       "named group does not match a core group resource",
			workload:   pod,
			attributes: map[string]string{identifiers.AttributeApiGroup: "apps"},
			expected:   false,
		},
		{
			name:       "an explicit empty apiGroup names the core group and matches a core resource",
			workload:   pod,
			attributes: map[string]string{identifiers.AttributeApiGroup: ""},
			expected:   true,
		},
		{
			name:       "an explicit empty apiGroup does not match a named group resource",
			workload:   deployment,
			attributes: map[string]string{identifiers.AttributeApiGroup: ""},
			expected:   false,
		},
		{
			name:     "an explicit empty apiGroup constrains an otherwise matching kind",
			workload: deployment,
			attributes: map[string]string{
				identifiers.AttributeApiGroup: "",
				identifiers.AttributeKind:     "Deployment",
			},
			expected: false,
		},
		{
			name:     "apiGroup alongside a real label that matches",
			workload: deployment,
			attributes: map[string]string{
				identifiers.AttributeApiGroup: "apps",
				"app":                         "test-app",
			},
			expected: true,
		},
		{
			name:     "apiGroup alongside a real label that does not match",
			workload: deployment,
			attributes: map[string]string{
				identifiers.AttributeApiGroup: "apps",
				"app":                         "other-app",
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			designator := &identifiers.PortalDesignator{
				DesignatorType: identifiers.DesignatorAttributes,
				Attributes:     tt.attributes,
			}
			assert.Equal(t, tt.expected, p.metadataHasException(tt.workload, designator.DigestPortalDesignator(), nil, nil))
		})
	}
}

// subjectGroupVector is a RegoResponseVector built from an RBAC subject. It carries a bare
// apiGroup and no apiVersion, which is what distinguishes it from an ordinary workload.
const subjectGroupVector = `{"apiGroup":"rbac.authorization.k8s.io","kind":"Group","name":"system:masters","relatedObjects":[{"apiVersion":"rbac.authorization.k8s.io/v1","kind":"ClusterRoleBinding","metadata":{"name":"cluster-admin"},"roleRef":{"apiGroup":"rbac.authorization.k8s.io","kind":"ClusterRole","name":"cluster-admin"},"subjects":[{"apiGroup":"rbac.authorization.k8s.io","kind":"Group","name":"system:masters"}]}]}`

// TestHasException_ApiGroup_RegoResponseVector covers the RBAC subject case, where
// GetApiVersion returns a bare API group rather than a group/version pair. Splitting that as
// an apiVersion yields no group, which silently stopped a correctly scoped subject exception
// from matching.
func TestHasException_ApiGroup_RegoResponseVector(t *testing.T) {
	p := NewProcessor()

	vector, err := objectsenvelopes.NewRegoResponseVectorObjectFromBytes([]byte(subjectGroupVector))
	require.NoError(t, err)

	tests := []struct {
		name       string
		attributes map[string]string
		expected   bool
	}{
		{
			name: "matching named group attaches",
			attributes: map[string]string{
				identifiers.AttributeApiGroup: "rbac.authorization.k8s.io",
				identifiers.AttributeKind:     "Group",
				identifiers.AttributeName:     "system:masters",
			},
			expected: true,
		},
		{
			name: "mismatched named group does not attach",
			attributes: map[string]string{
				identifiers.AttributeApiGroup: "apps",
				identifiers.AttributeKind:     "Group",
				identifiers.AttributeName:     "system:masters",
			},
			expected: false,
		},
		{
			name: "an explicit core group does not match a named group subject",
			attributes: map[string]string{
				identifiers.AttributeApiGroup: "",
				identifiers.AttributeKind:     "Group",
				identifiers.AttributeName:     "system:masters",
			},
			expected: false,
		},
		{
			name: "apiGroup alone attaches",
			attributes: map[string]string{
				identifiers.AttributeApiGroup: "rbac.authorization.k8s.io",
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			designator := &identifiers.PortalDesignator{
				DesignatorType: identifiers.DesignatorAttributes,
				Attributes:     tt.attributes,
			}
			assert.Equal(t, tt.expected, p.hasException("test", designator, vector, nil, nil))
		})
	}
}
