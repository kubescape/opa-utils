package resourcesresults

import (
	"github.com/armosec/armoapi-go/armotypes"
)

func MockResults() []Result {
	return []Result{
		*mockResultPassed(),
		*mockResultFailed(),
		// *mockResultSkipped(),
	}
}

func mockResultPassed() *Result {
	return &Result{
		ResourceID: "resource/passed",
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
	return &Result{
		ResourceID: "resource/failed",
		AssociatedControls: []ResourceAssociatedControl{
			*mockResourceAssociatedControl0087Failed(),
			*mockResourceAssociatedControl0088Failed(),
			*mockResourceAssociatedControl0089Passed(),
		},
	}
}
func mockResourceAssociatedControl0087Failed() *ResourceAssociatedControl {
	return &ResourceAssociatedControl{
		ControlID:             "C-0087",
		ControlConfigurations: nil,
		ResourceAssociatedRules: []ResourceAssociatedRule{
			*mockResourceAssociatedRuleA(),
			*mockResourceAssociatedRuleB(),
		},
	}
}

func mockResourceAssociatedControl0088Failed() *ResourceAssociatedControl {
	return &ResourceAssociatedControl{
		ControlID:             "C-0088",
		ControlConfigurations: nil,
		ResourceAssociatedRules: []ResourceAssociatedRule{
			*mockResourceAssociatedRuleB(),
		},
	}
}

func mockResourceAssociatedControl0089Passed() *ResourceAssociatedControl {
	return &ResourceAssociatedControl{
		ControlID:             "C-0089",
		ControlConfigurations: nil,
	}
}
func mockResourceAssociatedRuleA() *ResourceAssociatedRule {
	return &ResourceAssociatedRule{
		Name:        "ruleA",
		FailedPaths: []string{"path/to/fail/A"},
		Exception:   []armotypes.PostureExceptionPolicy{},
	}
}

func mockResourceAssociatedRuleB() *ResourceAssociatedRule {
	return &ResourceAssociatedRule{
		Name:        "ruleB",
		FailedPaths: []string{"path/to/fail/B"},
		Exception:   []armotypes.PostureExceptionPolicy{},
	}
}

// func mockResourceAssociatedRuleWithFWException() *ResourceAssociatedRule {
// 	return &ResourceAssociatedRule{
// 		Name:        "ruleB",
// 		FailedPaths: []string{"path/to/fail/B"},
// 		Exception:   []armotypes.PostureExceptionPolicy{},
// 	}
// }
