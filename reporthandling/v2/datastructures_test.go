package v2

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/armosec/armoapi-go/armotypes"
	"github.com/armosec/opa-utils/objectsenvelopes"
	"github.com/armosec/opa-utils/reporthandling"
	"github.com/stretchr/testify/assert"
)

func GetPostureReportMock() *PostureReport {
	resource := []Resource{}
	err := json.Unmarshal([]byte(ResourcesListMock), &resource)
	if err != nil {
		panic(err)
	}
	i := 0
	results := []Result{}
	for i = 0; i < 5; i++ {
		results = append(results, Result{
			ResourceID: objectsenvelopes.NewObject(resource[i].Object.(map[string]interface{})).GetID(),
			AssociatedControls: []ResourceAssociatedControl{
				{
					ControlID: "C-0045",
					ResourceAssociatedRules: []ResourceAssociatedRule{
						{
							RuleName:    "bla-bla",
							FailedPaths: []string{},
							Exception:   &armotypes.PostureExceptionPolicy{},
							Status:      reporthandling.StatusPassed,
						},
					},
				},
			},
		},
		)
	}
	for j := i; j < len(resource); j++ {
		results = append(results, Result{
			ResourceID: objectsenvelopes.NewObject(resource[j].Object.(map[string]interface{})).GetID(),
			AssociatedControls: []ResourceAssociatedControl{
				{
					ControlID: "C-0045",
				},
			},
		},
		)
	}

	return &PostureReport{
		CustomerGUID:         "0343c0ee-22ab-4d90-8fbf-2a145a311b90",
		ClusterName:          "minikube",
		ReportID:             "9001c1da-3840-4f9e-a7d3-65eda7faf2e3",
		ReportGenerationTime: time.Now().UTC(),
		SummaryDetails: SummaryDetails{
			Frameworks: []FrameworkSummary{
				{
					Name:  "NSA",
					Score: 68,
					Controls: map[string]ControlSummary{
						"C-0045": {
							Score:            68,
							PassedResources:  17,
							FailedResources:  5,
							WarningResources: 0,
							Status:           reporthandling.StatusFailed,
						},
					},
				},
			},
			Controls: map[string]ControlSummary{
				"C-0045": {
					Score:            68,
					PassedResources:  17,
					FailedResources:  5,
					WarningResources: 0,
					Status:           reporthandling.StatusFailed,
				},
			},
		},
		Results:   results,
		Resources: GetResourcesListMock(),
	}
}

func TestPostureReportMock(t *testing.T) {
	p := GetPostureReportMock()
	assert.Equal(t, 22, len(p.Resources))
	// t.Error(p.ToString())
}
