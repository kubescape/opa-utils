package exceptions

import (
	"testing"

	"github.com/kubescape/k8s-interface/workloadinterface"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/armosec/armoapi-go/armotypes"
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
		Resources: []armotypes.PortalDesignator{
			{
				DesignatorType: armotypes.DesignatorAttributes,
				Attributes: map[string]string{
					armotypes.AttributeNamespace: "default",
					armotypes.AttributeCluster:   "unittest",
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
		Resources: []armotypes.PortalDesignator{
			{
				DesignatorType: armotypes.DesignatorAttributes,
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

func emptyPostureExceptionPolicyAlertOnlyMock() *armotypes.PostureExceptionPolicy {
	return &armotypes.PostureExceptionPolicy{
		PortalBase: armotypes.PortalBase{
			Name: "postureExceptionPolicyAlertOnlyMock",
		},
		PolicyType: "postureExceptionPolicy",
		Actions:    []armotypes.PostureExceptionPolicyActions{armotypes.AlertOnly},
		Resources: []armotypes.PortalDesignator{
			{
				DesignatorType: armotypes.DesignatorAttributes,
				Attributes:     map[string]string{},
			},
		},
		PosturePolicies: []armotypes.PosturePolicy{},
	}
}
func TestListRuleExceptions(t *testing.T) {
	p := NewProcessor()
	exceptionPolicies := []armotypes.PostureExceptionPolicy{*postureExceptionPolicyAlertOnlyMock()}
	res1 := p.ListRuleExceptions(exceptionPolicies, "MITRE", "", "", "")
	assert.Equal(t, 1, len(res1))

	res2 := p.ListRuleExceptions(exceptionPolicies, "", "hostPath mount", "", "")
	assert.Equal(t, len(res2), 1)

	res3 := p.ListRuleExceptions(exceptionPolicies, "NSA", "", "", "")
	assert.Equal(t, len(res3), 0)

}

func TestListRuleExceptionsRegex(t *testing.T) {
	p := NewProcessor()
	exceptionPolicy := emptyPostureExceptionPolicyAlertOnlyMock()
	res1 := p.ListRuleExceptions([]armotypes.PostureExceptionPolicy{*exceptionPolicy}, "MITRE", "", "", "")
	assert.Equal(t, 1, len(res1))

	exceptionPolicy.PosturePolicies = append(exceptionPolicy.PosturePolicies, armotypes.PosturePolicy{
		FrameworkName: "MIT.*",
	})

	res2 := p.ListRuleExceptions([]armotypes.PostureExceptionPolicy{*exceptionPolicy}, "MITRE", "", "", "")
	assert.Equal(t, 1, len(res2))

	res2 = p.ListRuleExceptions([]armotypes.PostureExceptionPolicy{*exceptionPolicy}, "2MITRE", "", "", "")
	assert.Equal(t, 0, len(res2))

	exceptionPolicy.PosturePolicies[0] = armotypes.PosturePolicy{
		FrameworkName: "mit.*",
	}
	res2 = p.ListRuleExceptions([]armotypes.PostureExceptionPolicy{*exceptionPolicy}, "MITRE", "", "", "")
	assert.Equal(t, 1, len(res2))

	exceptionPolicy.PosturePolicies[0] = armotypes.PosturePolicy{
		FrameworkName: "mitre",
	}
	res2 = p.ListRuleExceptions([]armotypes.PostureExceptionPolicy{*exceptionPolicy}, "MITRE", "", "", "")
	assert.Equal(t, 1, len(res2))

	exceptionPolicy.PosturePolicies[0] = armotypes.PosturePolicy{
		FrameworkName: "MITRE",
		ControlName:   "my.*",
		RuleName:      "rule.*vk",
	}

	res3 := p.ListRuleExceptions([]armotypes.PostureExceptionPolicy{*exceptionPolicy}, "MITRE", "", "", "")
	assert.Equal(t, 1, len(res3))

	res3 = p.ListRuleExceptions([]armotypes.PostureExceptionPolicy{*exceptionPolicy}, "MITRE", "my-control", "", "")
	assert.Equal(t, 1, len(res3))

	res3 = p.ListRuleExceptions([]armotypes.PostureExceptionPolicy{*exceptionPolicy}, "MITRE", "control-my", "", "")
	assert.Equal(t, 0, len(res3))

	res3 = p.ListRuleExceptions([]armotypes.PostureExceptionPolicy{*exceptionPolicy}, "MITRE", "my-control", "", "rulebla -bla vk")
	assert.Equal(t, 1, len(res3))

	res3 = p.ListRuleExceptions([]armotypes.PostureExceptionPolicy{*exceptionPolicy}, "MITRE", "control-my", "", "rulebla -bla")
	assert.Equal(t, 0, len(res3))
}

func TestGetResourceExceptions(t *testing.T) {
	emptyObj, err := workloadinterface.NewBaseObjBytes([]byte(`{"apiVersion": "v1", "kind":"Deployment", "metadata": {"name": "test"}}`))
	require.NoError(t, err)

	withLabelObj, err := workloadinterface.NewBaseObjBytes([]byte(`{"apiVersion": "v1", "kind":"Deployment", "metadata": {"name": "test", "labels": {"myLabelOrAnnotation" : "static_test"}}}`))
	require.NoError(t, err)

	withAnnotationObj, err := workloadinterface.NewBaseObjBytes([]byte(`{"apiVersion": "v1", "kind":"Deployment", "metadata": {"name": "test", "annotations": {"myLabelOrAnnotation" : "static_test"}}}`))
	require.NoError(t, err)

	exceptionPolicy := postureLabelsRegexExceptionPolicyAlertOnlyMock()
	exceptionPolicyRegex := postureLabelsRegexExceptionPolicyAlertOnlyMock()
	exceptionPolicyRegex.Resources[0].Attributes["myLabelOrAnnotation"] = "static_.*"

	p := NewProcessor()

	testCases := []struct {
		desc                    string
		exceptionPolicy         *armotypes.PostureExceptionPolicy
		workloadObj             workloadinterface.IMetadata
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
	}

	for _, test := range testCases {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			res := p.GetResourceExceptions([]armotypes.PostureExceptionPolicy{*exceptionPolicy}, test.workloadObj, "")
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

// func TestGetException(t *testing.T) {
// 	exceptionPolicies := []armotypes.PostureExceptionPolicy{*PostureExceptionPolicyAlertOnlyMock()}
// 	res1 := ListRuleExceptions(exceptionPolicies, "MITRE", "", "")
// 	if len(res1) != 1 {
// 		t.Errorf("expecting 1 exception")
// 	}
// 	res2 := ListRuleExceptions(exceptionPolicies, "", "hostPath mount", "")
// 	if len(res2) != 0 {
// 		t.Errorf("expecting 0 exception")
// 	}
// }
