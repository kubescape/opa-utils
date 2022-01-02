package resourcesresults

import (
	"github.com/armosec/armoapi-go/armotypes"
	"github.com/armosec/k8s-interface/workloadinterface"
)

func MockResults() []Result {
	return []Result{
		*mockResultPassed(),
		*mockResultFailed(),
		// *mockResultSkipped(),
	}
}

func mockResultPassed() *Result {
	w := workloadinterface.NewWorkloadMock(nil)
	return &Result{
		ResourceID: w.GetID(),
		AssociatedControls: []ResourceAssociatedControl{
			*mockResourceAssociatedControl0089Passed(),
		},
	}
}

// func mockResultSkipped() *Result {
// 	return &Result{
// 		ResourceID:         "resource/passed",
// 		AssociatedControls: []ResourceAssociatedControl{},
// 	}
// }
func mockResultFailed() *Result {
	w := workloadinterface.NewWorkloadMock(nil)
	return &Result{
		ResourceID: w.GetID(),
		AssociatedControls: []ResourceAssociatedControl{
			*mockResourceAssociatedControl0087Failed(),
			*mockResourceAssociatedControl0088Failed(),
			*mockResourceAssociatedControl0089Passed(),
		},
	}
}
func mockResourceAssociatedControl0087Failed() *ResourceAssociatedControl {
	return &ResourceAssociatedControl{
		ControlID: "C-0087",
		ResourceAssociatedRules: []ResourceAssociatedRule{
			*mockResourceAssociatedRuleA(),
			*mockResourceAssociatedRuleB(),
		},
	}
}

func mockResourceAssociatedControl0088Failed() *ResourceAssociatedControl {
	return &ResourceAssociatedControl{
		ControlID: "C-0088",
		ResourceAssociatedRules: []ResourceAssociatedRule{
			*mockResourceAssociatedRuleB(),
		},
	}
}

func mockResourceAssociatedControl0089Passed() *ResourceAssociatedControl {
	return &ResourceAssociatedControl{
		ControlID: "C-0089",
	}
}
func mockResourceAssociatedRuleA() *ResourceAssociatedRule {
	return &ResourceAssociatedRule{
		Name:                  "ruleA",
		Paths:                 []Path{{FailedPath: "path/to/fail/A"}},
		Exception:             []armotypes.PostureExceptionPolicy{},
		ControlConfigurations: nil,
	}
}

func mockResourceAssociatedRuleB() *ResourceAssociatedRule {
	return &ResourceAssociatedRule{
		Name:                  "ruleB",
		Paths:                 []Path{{FailedPath: "path/to/fail/B"}},
		Exception:             []armotypes.PostureExceptionPolicy{},
		ControlConfigurations: nil,
	}
}

// func mockResourceAssociatedRuleWithFWException() *ResourceAssociatedRule {
// 	return &ResourceAssociatedRule{
// 		Name:        "ruleB",
// 		FailedPaths: []string{"path/to/fail/B"},
// 		Exception:   []armotypes.PostureExceptionPolicy{},
// 	}
// }
