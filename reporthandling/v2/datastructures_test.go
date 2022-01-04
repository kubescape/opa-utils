package v2

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"

	"github.com/armosec/armoapi-go/armotypes"
	"github.com/armosec/opa-utils/objectsenvelopes"
	"github.com/armosec/opa-utils/reporthandling/results/v1/reportsummary"
	"github.com/armosec/opa-utils/reporthandling/results/v1/resourcesresults"
	"github.com/francoispqt/gojay"
	"github.com/stretchr/testify/assert"
)

func GetPostureReportMock() *PostureReport {

	resource := []Resource{}
	err := json.Unmarshal([]byte(ResourcesListMock), &resource)
	if err != nil {
		panic(err)
	}
	i := 0
	results := []resourcesresults.Result{}
	for i = 0; i < 5; i++ {
		results = append(results, resourcesresults.Result{
			ResourceID: objectsenvelopes.NewObject(resource[i].Object.(map[string]interface{})).GetID(),
			AssociatedControls: []resourcesresults.ResourceAssociatedControl{
				{
					ControlID: "C-0045",
					ResourceAssociatedRules: []resourcesresults.ResourceAssociatedRule{
						{
							Name:      "bla-bla",
							Exception: []armotypes.PostureExceptionPolicy{},
						},
					},
				},
			},
		},
		)
	}
	for j := i; j < len(resource); j++ {
		results = append(results, resourcesresults.Result{
			ResourceID: objectsenvelopes.NewObject(resource[j].Object.(map[string]interface{})).GetID(),
			AssociatedControls: []resourcesresults.ResourceAssociatedControl{
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
		SummaryDetails: reportsummary.SummaryDetails{
			Frameworks: []reportsummary.FrameworkSummary{
				{
					Name:  "NSA",
					Score: 68,
					Controls: map[string]reportsummary.ControlSummary{
						"C-0045": {
							Score: 68,
							ResourceCounters: reportsummary.ResourceCounters{
								PassedResources:   17,
								FailedResources:   5,
								ExcludedResources: 0,
							},
						},
					},
				},
			},
			Controls: map[string]reportsummary.ControlSummary{
				"C-0045": {
					Score: 68,
					ResourceCounters: reportsummary.ResourceCounters{
						PassedResources:   17,
						FailedResources:   5,
						ExcludedResources: 0,
					},
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

//this test validates the unmarshaller that used to validate the posture object in e.r and other places
func TestPostureReportGojayUnmarshal(t *testing.T) {
	postureReport := &PostureReport{}
	original := GetPostureReportMock()
	asBytes, err := json.Marshal(original)
	assert.NoError(t, err, "failed to marshal postureReport")

	err = gojay.NewDecoder(bytes.NewReader(asBytes)).Decode(postureReport)
	assert.NoError(t, err, "failed to unmarshal using gojay postureReport")

	assert.Equal(t, original.ReportID, postureReport.ReportID)
	assert.Equal(t, original.CustomerGUID, postureReport.CustomerGUID)
	assert.Equal(t, original.ClusterName, postureReport.ClusterName)
	assert.Equal(t, original.ReportGenerationTime.UTC(), postureReport.ReportGenerationTime.UTC())

}
