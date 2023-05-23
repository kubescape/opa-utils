package resourcesresults

import (
	"testing"

	"github.com/armosec/armoapi-go/armotypes"
	"github.com/kubescape/k8s-interface/workloadinterface"
	"github.com/kubescape/opa-utils/exceptions"
	"github.com/kubescape/opa-utils/reporthandling"
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

func mockControlsList() map[string]reporthandling.Control {
	return map[string]reporthandling.Control{
		"C-0087": {},
		"C-0088": {},
		"C-0089": {},
	}
}

func TestSetExceptions(t *testing.T) {
	w := workloadinterface.NewWorkloadMock(nil)
	processor := exceptions.NewProcessor()

	exceptions := []armotypes.PostureExceptionPolicy{}
	exceptions = append(exceptions, *mockExceptionDeploymentC0087())
	exceptions = append(exceptions, *mockExceptionUnitestDeploymentC0087())
	exceptions = append(exceptions, *mockExceptionUnitestC0088())
	exceptions = append(exceptions, *mockExceptionDeploymentC0089())
	c := mockControlsList()
	// simple test
	result1 := mockResultFailed()
	result1.SetExceptions(w, exceptions, "", c, WithExceptionsProcessor(processor))
	assert.Equal(t, 2, result1.ListControlsIDs(nil).Passed())
	assert.Equal(t, 1, result1.ListControlsIDs(nil).Failed())

	// without option to reuse the processor
	result1.SetExceptions(w, exceptions, "", c)
	assert.Equal(t, 2, result1.ListControlsIDs(nil).Passed())
	assert.Equal(t, 1, result1.ListControlsIDs(nil).Failed())

	// test cluster name
	result2 := mockResultFailed()
	result2.SetExceptions(w, exceptions, "unitest", c, WithExceptionsProcessor(processor))
	assert.Equal(t, 3, result2.ListControlsIDs(nil).Passed())
	assert.Equal(t, 0, result2.ListControlsIDs(nil).Failed())

	// test wrong cluster name
	result3 := mockResultFailed()
	result3.SetExceptions(w, exceptions, "unitest2", c, WithExceptionsProcessor(processor))
	assert.Equal(t, 2, result3.ListControlsIDs(nil).Passed())
	assert.Equal(t, 1, result3.ListControlsIDs(nil).Failed())
}
