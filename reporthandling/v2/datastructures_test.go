package v2

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"

	"github.com/armosec/armoapi-go/armotypes"
	"github.com/francoispqt/gojay"
	"github.com/kubescape/opa-utils/objectsenvelopes"
	"github.com/kubescape/opa-utils/reporthandling"
	"github.com/kubescape/opa-utils/reporthandling/results/v1/reportsummary"
	"github.com/kubescape/opa-utils/reporthandling/results/v1/resourcesresults"
	"github.com/stretchr/testify/assert"
)

func GetPostureReportMock() *PostureReport {

	resource := []reporthandling.Resource{}
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
		Metadata: Metadata{
			ClusterMetadata: ClusterMetadata{
				NumberOfWorkerNodes: 8,
				ContextName:         "minikube",
			},
			ScanMetadata: ScanMetadata{
				Format:             "json,pdf,pretty-printer",
				Formats:            []string{"json", "pdf", "pretty-printer"},
				ExcludedNamespaces: []string{"exclude-namespace"},
				IncludeNamespaces:  []string{"include-namespace"},
				FailThreshold:      23.54,
				Submit:             true,
				HostScanner:        true,
				Logger:             "fantastic-logger",
				TargetType:         "framework",
				TargetNames:        []string{"framework"},
				UseExceptions:      "/path/to/exceptions",
				ControlsInputs:     "/path/to/ctrls",
				VerboseMode:        true,
			},
		},
	}
}

func TestPostureReportMock(t *testing.T) {
	p := GetPostureReportMock()
	assert.Equal(t, 22, len(p.Resources))
	// t.Error(p.ToString())
}

// TestPostureReportGojayUnmarshal validates the unmarshaller that is used to validate the posture object in e.r and other places
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

	// Metadata
	assert.Equal(t, original.Metadata.ClusterMetadata.NumberOfWorkerNodes, postureReport.Metadata.ClusterMetadata.NumberOfWorkerNodes)
	assert.Equal(t, original.Metadata.ClusterMetadata.ContextName, postureReport.Metadata.ClusterMetadata.ContextName)
	assert.Equal(t, original.Metadata.ScanMetadata.Format, postureReport.Metadata.ScanMetadata.Format)
	assert.Equal(t, original.Metadata.ScanMetadata.Formats, postureReport.Metadata.ScanMetadata.Formats)
	assert.Equal(t, original.Metadata.ScanMetadata.ExcludedNamespaces, postureReport.Metadata.ScanMetadata.ExcludedNamespaces)
	assert.Equal(t, original.Metadata.ScanMetadata.IncludeNamespaces, postureReport.Metadata.ScanMetadata.IncludeNamespaces)
	assert.Equal(t, original.Metadata.ScanMetadata.FailThreshold, postureReport.Metadata.ScanMetadata.FailThreshold)
	assert.Equal(t, original.Metadata.ScanMetadata.Submit, postureReport.Metadata.ScanMetadata.Submit)
	assert.Equal(t, original.Metadata.ScanMetadata.HostScanner, postureReport.Metadata.ScanMetadata.HostScanner)
	assert.Equal(t, original.Metadata.ScanMetadata.Logger, postureReport.Metadata.ScanMetadata.Logger)
	assert.Equal(t, original.Metadata.ScanMetadata.TargetType, postureReport.Metadata.ScanMetadata.TargetType)
	assert.Equal(t, original.Metadata.ScanMetadata.TargetNames, postureReport.Metadata.ScanMetadata.TargetNames)
	assert.Equal(t, original.Metadata.ScanMetadata.UseExceptions, postureReport.Metadata.ScanMetadata.UseExceptions)
	assert.Equal(t, original.Metadata.ScanMetadata.ControlsInputs, postureReport.Metadata.ScanMetadata.ControlsInputs)
	assert.Equal(t, original.Metadata.ScanMetadata.VerboseMode, postureReport.Metadata.ScanMetadata.VerboseMode)
}
