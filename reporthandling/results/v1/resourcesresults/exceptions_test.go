package resourcesresults

import (
	"testing"

	"github.com/armosec/armoapi-go/armotypes"
	"github.com/armosec/k8s-interface/workloadinterface"
	v1 "github.com/armosec/opa-utils/reporthandling/helpers/v1"
	"github.com/stretchr/testify/assert"
)

func mockExceptionDeploymentC0087() *armotypes.PostureExceptionPolicy {
	return &armotypes.PostureExceptionPolicy{
		PortalBase: armotypes.PortalBase{
			Name: "DeploymentC0087",
		},
		Actions: []armotypes.PostureExceptionPolicyActions{armotypes.AlertOnly},
		Resources: []armotypes.PortalDesignator{
			{
				DesignatorType: armotypes.DesignatorAttributes,
				Attributes: map[string]string{
					armotypes.AttributeKind: "Deployment",
				},
			},
		},
		PosturePolicies: []armotypes.PosturePolicy{
			{
				ControlID: "C-0087",
			},
		},
	}
}

func mockExceptionUnitestDeploymentC0087() *armotypes.PostureExceptionPolicy {
	return &armotypes.PostureExceptionPolicy{
		PortalBase: armotypes.PortalBase{
			Name: "unitestDeploymentC0087",
		},
		Actions: []armotypes.PostureExceptionPolicyActions{armotypes.AlertOnly},
		Resources: []armotypes.PortalDesignator{
			{
				DesignatorType: armotypes.DesignatorAttributes,
				Attributes: map[string]string{
					armotypes.AttributeCluster: "unitest",
					armotypes.AttributeKind:    "Deployment",
				},
			},
		},
		PosturePolicies: []armotypes.PosturePolicy{
			{
				ControlID: "C-0087",
			},
		},
	}
}

func mockExceptionUnitestC0088() *armotypes.PostureExceptionPolicy {
	return &armotypes.PostureExceptionPolicy{
		PortalBase: armotypes.PortalBase{
			Name: "unitestC0088",
		},
		Actions: []armotypes.PostureExceptionPolicyActions{armotypes.AlertOnly},
		Resources: []armotypes.PortalDesignator{
			{
				DesignatorType: armotypes.DesignatorAttributes,
				Attributes: map[string]string{
					armotypes.AttributeCluster: "unitest",
				},
			},
		},
		PosturePolicies: []armotypes.PosturePolicy{
			{
				ControlID: "C-0088",
			},
		},
	}
}

func mockExceptionNSAUnitest() *armotypes.PostureExceptionPolicy {
	return &armotypes.PostureExceptionPolicy{
		PortalBase: armotypes.PortalBase{
			Name: "NSAUnitest",
		},
		Actions: []armotypes.PostureExceptionPolicyActions{armotypes.AlertOnly},
		Resources: []armotypes.PortalDesignator{
			{
				DesignatorType: armotypes.DesignatorAttributes,
				Attributes: map[string]string{
					armotypes.AttributeCluster: "unitest",
				},
			},
		},
		PosturePolicies: []armotypes.PosturePolicy{
			{
				FrameworkName: "NSA",
			},
		},
	}
}

func mockExceptionDeploymentC0089() *armotypes.PostureExceptionPolicy {
	return &armotypes.PostureExceptionPolicy{
		PortalBase: armotypes.PortalBase{
			Name: "Deployment0089",
		},
		Actions: []armotypes.PostureExceptionPolicyActions{armotypes.AlertOnly},
		Resources: []armotypes.PortalDesignator{
			{
				DesignatorType: armotypes.DesignatorAttributes,
				Attributes: map[string]string{
					armotypes.AttributeKind: "Deployment",
				},
			},
		},
		PosturePolicies: []armotypes.PosturePolicy{
			{
				ControlID: "C-0089",
			},
		},
	}
}

func TestSetExceptions(t *testing.T) {
	w := workloadinterface.NewWorkloadMock(nil)

	exceptions := []armotypes.PostureExceptionPolicy{}
	exceptions = append(exceptions, *mockExceptionDeploymentC0087())
	exceptions = append(exceptions, *mockExceptionUnitestDeploymentC0087())
	exceptions = append(exceptions, *mockExceptionUnitestC0088())
	exceptions = append(exceptions, *mockExceptionDeploymentC0089())

	// simple test
	result1 := mockResultFailed()
	result1.SetExceptions(w, exceptions, "")
	assert.Equal(t, 1, len(result1.ListControls(nil).Excluded()))
	assert.Equal(t, 1, len(result1.ListControls(nil).Passed()))
	assert.Equal(t, 1, len(result1.ListControls(nil).Failed()))

	// test cluster name
	result2 := mockResultFailed()
	result2.SetExceptions(w, exceptions, "unitest")
	assert.Equal(t, 2, len(result2.ListControls(nil).Excluded()))
	assert.Equal(t, 1, len(result2.ListControls(nil).Passed()))
	assert.Equal(t, 0, len(result2.ListControls(nil).Failed()))

	// test wrong cluster name
	result3 := mockResultFailed()
	result3.SetExceptions(w, exceptions, "unitest2")
	assert.Equal(t, 1, len(result3.ListControls(nil).Excluded()))
	assert.Equal(t, 1, len(result3.ListControls(nil).Passed()))
	assert.Equal(t, 1, len(result3.ListControls(nil).Failed()))

	// test filters on frameworks
	exceptions = []armotypes.PostureExceptionPolicy{}
	exceptions = append(exceptions, *mockExceptionUnitestC0088())
	exceptions = append(exceptions, *mockExceptionNSAUnitest())

	result4 := mockResultFailed()
	result4.SetExceptions(w, exceptions, "unitest")
	assert.Equal(t, 2, len(result4.ListControls(&v1.Filters{FrameworkNames: []string{"nsa"}}).Excluded()))
	assert.Equal(t, 2, len(result4.ListControls(&v1.Filters{FrameworkNames: []string{"nsa"}}).Excluded()))
	assert.Equal(t, 2, len(result4.ListControls(&v1.Filters{FrameworkNames: []string{""}}).Excluded()))
	assert.Equal(t, 2, len(result4.ListControls(nil).Excluded()))
	assert.Equal(t, 2, len(result4.ListControls(&v1.Filters{FrameworkNames: []string{"mitre", "nsa"}}).Excluded()))
	assert.Equal(t, 1, len(result4.ListControls(&v1.Filters{FrameworkNames: []string{"mitre"}}).Excluded()))
}
