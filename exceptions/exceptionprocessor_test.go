package exceptions

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/armosec/armoapi-go/armotypes"
)

func postureExceptionPolicyDisableMock() *armotypes.PostureExceptionPolicy {
	return &armotypes.PostureExceptionPolicy{}
}

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
				FrameworkName: "MITRE",
			},
		},
	}
}

func TestListRuleExceptions(t *testing.T) {
	exceptionPolicies := []armotypes.PostureExceptionPolicy{*postureExceptionPolicyAlertOnlyMock()}
	res1 := ListRuleExceptions(exceptionPolicies, "MITRE", "", "", "")
	assert.Equal(t, 1, len(res1))

	res2 := ListRuleExceptions(exceptionPolicies, "", "hostPath mount", "", "")
	assert.Equal(t, len(res2), 1)

	res3 := ListRuleExceptions(exceptionPolicies, "NSA", "", "", "")
	assert.Equal(t, len(res3), 0)
}

func TestRegexCompare(t *testing.T) {
	assert.True(t, compareCluster(".*minikube.*", "bez-minikube-25-10"))
	assert.True(t, compareCluster("bez-minikube-25-10", "bez-minikube-25-10"))
	assert.False(t, compareCluster("minikube", "bez-minikube-25-10"))
	assert.False(t, compareCluster("bla", "bez-minikube-25-10"))
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
