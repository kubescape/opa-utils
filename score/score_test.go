package score

import (
	"fmt"
	"testing"

	"github.com/armosec/armoapi-go/armotypes"
	k8sinterface "github.com/kubescape/k8s-interface/k8sinterface"
	"github.com/kubescape/k8s-interface/workloadinterface"
	armoupautils "github.com/kubescape/opa-utils/objectsenvelopes"
	"github.com/kubescape/opa-utils/reporthandling"
	"github.com/kubescape/opa-utils/reporthandling/apis"
	helpers "github.com/kubescape/opa-utils/reporthandling/helpers/v1"
	"github.com/kubescape/opa-utils/reporthandling/results/v1/reportsummary"
	"github.com/kubescape/opa-utils/reporthandling/results/v1/resourcesresults"
	v2 "github.com/kubescape/opa-utils/reporthandling/v2"
	"github.com/kubescape/opa-utils/score/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetScore(t *testing.T) {
	t.Parallel()

	t.Run("with workloads with replicas", func(t *testing.T) {
		for _, toPin := range []string{
			"deployment",
			"replicaSet",
		} {
			resourceType := toPin

			t.Run(fmt.Sprintf("with %s workload, 3 replicas", resourceType), func(t *testing.T) {
				t.Parallel()

				deployment := mocks.GetResourceByType(t, "deployment")
				wl := workloadinterface.NewWorkloadObj(deployment)
				require.NotNil(t, wl,
					"invalid deployment workload in mock",
				)

				require.Equalf(t, 3, wl.GetReplicas(),
					"invalid wl was put into the test, should have 3 replicas %v", deployment,
				)

				var s ScoreUtil
				score := s.GetScore(deployment)

				const expected = float32(3.30)
				require.InDeltaf(t, expected, score, 1e-6,
					"invalid score: should be %v~(numerical errrors considered), but got: %v", expected, score,
				)
			})

			t.Run(fmt.Sprintf("with %s workload, 1 replica", resourceType), func(t *testing.T) {
				t.Parallel()

				deployment := mocks.GetResourceByType(t, "deployment", mocks.WithReplicas(1))
				wl := workloadinterface.NewWorkloadObj(deployment)
				require.NotNil(t, wl,
					"invalid deployment workload in mock",
				)
				require.Equalf(t, 1, wl.GetReplicas(),
					"invalid wl was put into the test, should have 1 replica %v", deployment,
				)

				var s ScoreUtil
				score := s.GetScore(deployment)

				const expected = float32(1.00)
				require.InDeltaf(t, expected, score, 1e-6,
					"invalid score: should be %v~(numerical errrors considered), but got: %v", expected, score,
				)
			})
		}
	})

	t.Run("with daemonSet workload", func(t *testing.T) {
		t.Run("with active daemonSet", func(t *testing.T) {
			t.Parallel()

			ds := mocks.GetResourceByType(t, "daemonset")
			var s ScoreUtil
			score := s.GetScore(ds)

			const expected = float32(13) // value defined by the mock
			require.InDeltaf(t, expected, score, 1e-6,
				"invalid score: should be %v~(numerical errrors considered), but got: %v", expected, score,
			)
		})

		t.Run("a status with zero desiredNumberScheduled should yield the default score", func(t *testing.T) {
			t.Parallel()

			ds := mocks.GetResourceByType(t, "daemonset", mocks.WithDesiredNumberScheduled(0))

			var s ScoreUtil
			score := s.GetScore(ds)

			const expected = float32(1)
			require.InDeltaf(t, expected, score, 1e-6,
				"invalid score: should be %v~(numerical errrors considered), but got: %v", expected, score,
			)
		})

		t.Run("a non-zero desiredNumberScheduled should multiply the score", func(t *testing.T) {
			t.Parallel()

			ds := mocks.GetResourceByType(t, "daemonset", mocks.WithDesiredNumberScheduled(10))

			var s ScoreUtil
			score := s.GetScore(ds)

			const expected = float32(10)
			require.InDeltaf(t, expected, score, 1e-6,
				"invalid score: should be %v~(numerical errrors considered), but got: %v", expected, score,
			)
		})
	})

	t.Run("with workloads with default score", func(t *testing.T) {
		for _, toPin := range []string{
			"pod",
		} {
			resourceType := toPin
			const expected = float32(1)

			t.Run(fmt.Sprintf("with %s workload", resourceType), func(t *testing.T) {
				t.Parallel()

				ds := mocks.GetResourceByType(t, resourceType)

				var s ScoreUtil
				score := s.GetScore(ds)

				require.InDeltaf(t, expected, score, 1e-6,
					"invalid score: should be %v~(numerical errrors considered), but got: %v", expected, score,
				)
			})
		}
	})

	t.Run("with non-workload resources, expect default score", func(t *testing.T) {
		const expected = float32(1)

		for _, toPin := range []string{
			"role",
			"secret",
		} {
			resourceType := toPin

			t.Run(fmt.Sprintf("with %s resource", resourceType), func(t *testing.T) {
				t.Parallel()

				ds := mocks.GetResourceByType(t, resourceType)
				require.Truef(t, k8sinterface.IsTypeWorkload(ds), // IsTypeWorkload also report non-workload objects as "legit" k8s resources
					"k8sinterface.IsTypeWorkload has changed behavior",
				)

				var s ScoreUtil
				score := s.GetScore(ds)

				require.InDeltaf(t, expected, score, 1e-6,
					"invalid score: should be %v~(numerical errrors considered), but got: %v", expected, score,
				)
			})
		}
	})

	t.Run("with invalid resources, expect default score", func(t *testing.T) {
		const expected = float32(1)

		for _, toPin := range mocks.GetInvalidResources(t) {
			resource := toPin
			kind := resource["kind"]
			if kind == "" {
				kind = "unknown"
			}

			t.Run(fmt.Sprintf("with kind=%q", kind), func(t *testing.T) {
				t.Parallel()

				require.Falsef(t, k8sinterface.IsTypeWorkload(resource), // IsTypeWorkload checks on non-empty kind and version
					"k8sinterface.IsTypeWorkload has changed behavior",
				)

				var s ScoreUtil
				score := s.GetScore(resource)

				require.InDeltaf(t, expected, score, 1e-6,
					"invalid score: should be %v~(numerical errrors considered), but got: %v", expected, score,
				)

			})
		}
	})

	t.Run("with resource with related objects", func(t *testing.T) {
		const expected = float32(1)

		for _, toPin := range []string{
			"serviceaccount",
		} {
			resourceType := toPin

			t.Run(fmt.Sprintf("with %s resource", resourceType), func(t *testing.T) {
				t.Parallel()

				ds := mocks.GetResourceByType(t, resourceType)
				require.Falsef(t, k8sinterface.IsTypeWorkload(ds),
					"k8sinterface.IsTypeWorkload has changed behavior",
				)
				require.Truef(t, armoupautils.IsTypeRegoResponseVector(ds),
					"objectsenvelopes.IsTypeRegoResponseVector has changed behavior",
				)

				var s ScoreUtil
				score := s.GetScore(ds)

				require.InDeltaf(t, expected, score, 1e-6,
					"invalid score: should be %v~(numerical errrors considered), but got: %v", expected, score,
				)
			})
		}
	})

	t.Run("invalid DaemonSet workload should leave input score unchanged [edge]", func(t *testing.T) {
		t.Parallel()

		ds := mocks.GetResourceByType(t, "daemonset", mocks.WithDesiredNumberScheduled(10))
		wl := workloadinterface.NewWorkloadObj(ds)
		require.NotNil(t, wl)

		// inject invalid map that fails to marshal (edge case)
		v := map[string]interface{}{
			"cannot-marshal": struct {
				F func() `json:"fail"`
			}{
				F: func() {},
			},
		}

		var s ScoreUtil
		const input = float32(3.141558)
		score := s.processWorkload(wl, input, v)

		require.Equal(t, input, score)
	})

	t.Run("non-workload related objects in rego response vector should not contribute to the score", func(t *testing.T) {
		t.Parallel()
		const expected = float32(1)

		ds := mocks.GetResourceByType(t, "serviceaccount", mocks.WithRelatedObjects(map[string]interface{}{
			"apiVersion": "hostdata.kubescape.cloud/v1", // object recognized as host sensor
			"name":       "not-a-workload",
		}))
		require.Falsef(t, k8sinterface.IsTypeWorkload(ds),
			"k8sinterface.IsTypeWorkload has changed behavior",
		)
		require.Truef(t, armoupautils.IsTypeRegoResponseVector(ds),
			"objectsenvelopes.IsTypeRegoResponseVector has changed behavior",
		)

		vec := armoupautils.NewRegoResponseVectorObject(ds)
		related := vec.GetRelatedObjects()
		require.Lenf(t, related, 3, "test conditions have changed: we want to simulate a non-workload related obect")

		var s ScoreUtil
		score := s.GetScore(ds)

		require.InDeltaf(t, expected, score, 1e-6,
			"invalid score: should be %v~(numerical errrors considered), but got: %v", expected, score,
		)
	})
}

func TestCalculatePostureReportV2(t *testing.T) {
	t.Parallel()

	t.Run("with empty report", func(t *testing.T) {
		t.Parallel()

		s := NewScore(map[string]workloadinterface.IMetadata{})
		report := &v2.PostureReport{
			SummaryDetails: reportsummary.SummaryDetails{Frameworks: []reportsummary.FrameworkSummary{{Name: "empty", Controls: reportsummary.ControlSummaries{}}}},
			Results:        []resourcesresults.Result{},
			Resources:      []reporthandling.Resource{},
		}

		require.Errorf(t, s.CalculatePostureReportV2(report),
			"empty framework should return an error",
		)

		require.Equal(t, float32(0), report.SummaryDetails.Frameworks[0].Score,
			"empty framework should return an error and have a score equals to 0",
		)
	})

	t.Run("with skipped report", func(t *testing.T) {
		t.Parallel()

		s := NewScore(map[string]workloadinterface.IMetadata{})
		report := &v2.PostureReport{
			SummaryDetails: reportsummary.SummaryDetails{Frameworks: []reportsummary.FrameworkSummary{{Name: "skipped", Controls: reportsummary.ControlSummaries{
				"skipped1": reportsummary.ControlSummary{
					Name:        "skipped1",
					ControlID:   "Skippie1",
					Description: "skipper",
				},
				"skipped2": reportsummary.ControlSummary{
					Name:        "skipped2",
					ControlID:   "Skippie2",
					Description: "skipper",
				},
			}}}},
			Results:   []resourcesresults.Result{},
			Resources: []reporthandling.Resource{},
		}

		require.Errorf(t, s.CalculatePostureReportV2(report),
			"empty framework should return an error",
		)

		require.Equal(t, float32(0), report.SummaryDetails.Frameworks[0].Score,
			"empty framework should return an error and have a score equals to 0",
		)
	})

	t.Run("with mock report", func(t *testing.T) {
		t.Parallel()

		resources, report := mockPostureReportV2(t)
		s := ScoreUtil{
			isDebugMode: true,
			resources:   resources,
		}

		require.NoErrorf(t, s.CalculatePostureReportV2(report),
			"mock framework should not return an error",
		)

		const (
			expectedForFramework1 = float32(62.577965)
			expectedForFramework2 = float32(46.42857)
			expectedSummary       = float32(51.280453)
		)

		t.Run("assert control scores", func(t *testing.T) {
			require.Len(t, report.SummaryDetails.Controls, 4)
			for _, control := range report.SummaryDetails.Controls {
				var expectedForControl float64

				switch control.ControlID {
				case "control-1":
					expectedForControl = 81.13208
				case "control-2":
					expectedForControl = 0 // passed
				case "control-3":
					expectedForControl = 66.666664
				case "control-4":
					expectedForControl = 0 // passed
				}

				assert.InDeltaf(t, expectedForControl, control.Score, 1e-6,
					"unexpected summarized score for control %q", control.ControlID,
				)
			}
		})

		t.Run("assert framework scores", func(t *testing.T) {
			assert.InDeltaf(t, expectedForFramework1, report.SummaryDetails.Frameworks[0].Score, 1e-6,
				"unexpected summarized score for framework[0]",
			)
			assert.InDeltaf(t, expectedForFramework2, report.SummaryDetails.Frameworks[1].Score, 1e-6,
				"unexpected summarized score for framework[1]",
			)
		})

		t.Run("assert final score", func(t *testing.T) {
			assert.InDeltaf(t, expectedSummary, report.SummaryDetails.Score, 1e-6,
				"unexpected summarized final score",
			)
		})
	})
}

func TestCalculatePostureReportV1(t *testing.T) {
	t.Parallel()

	t.Run("with empty report", func(t *testing.T) {
		t.Parallel()

		s := NewScore(map[string]workloadinterface.IMetadata{})
		report := reporthandling.PostureReport{FrameworkReports: []reporthandling.FrameworkReport{{
			Name:           "empty",
			ControlReports: []reporthandling.ControlReport{},
		}}}

		require.NoErrorf(t, s.Calculate(report.FrameworkReports),
			"empty framework should not return an error in V1",
		)
		require.Equalf(t, float32(0.0), report.FrameworkReports[0].Score,
			"empty framework should have a score equals to 0",
		)
	})

	t.Run("with mock report", func(t *testing.T) {
		t.Parallel()

		resources, reports := mockPostureReportV1(t)
		s := ScoreUtil{
			isDebugMode: true,
			resources:   resources,
		}

		require.NoErrorf(t, s.Calculate(reports),
			"mock framework should not return an error",
		)

		const (
			expectedForFramework1 = float32(62.577965)
			expectedForFramework2 = float32(54.736843) // NOTE(fredbi): V2 yields float32(46.42857)
			expectedSummary       = float32(51.280453)
		)

		t.Run("assert control scores", func(t *testing.T) {
			require.Len(t, reports[0].ControlReports, 2)
			require.Len(t, reports[1].ControlReports, 2)
			allControlReports := make([]reporthandling.ControlReport, 0, len(reports[0].ControlReports)+len(reports[1].ControlReports))
			allControlReports = append(allControlReports, reports[0].ControlReports...)
			allControlReports = append(allControlReports, reports[1].ControlReports...)

			for _, control := range allControlReports {
				var expectedForControl float64

				switch control.ControlID {
				case "control-1":
					expectedForControl = 81.13208
				case "control-2":
					expectedForControl = 0
				case "control-3":
					expectedForControl = 66.666664
				case "control-4":
					expectedForControl = 0
				}

				assert.InDeltaf(t, expectedForControl, control.Score, 1e-6,
					"unexpected summarized score for control %q", control.ControlID,
				)
			}
		})

		t.Run("assert framework scores", func(t *testing.T) {
			assert.InDeltaf(t, expectedForFramework1, reports[0].Score, 1e-6,
				"unexpected summarized score for framework[0]",
			)
			assert.InDeltaf(t, expectedForFramework2, reports[1].Score, 1e-6,
				"unexpected summarized score for framework[1]",
			)
		})
	})
}

func TestControlScoreError(t *testing.T) {
	// test behavior on edge case (invalid report inputs)
	t.Parallel()

	resources := mockResources(t)
	s := ScoreUtil{isDebugMode: true, resources: resources}
	controlReport := &reporthandling.ControlReport{
		Name:      "mock-control-1",
		ControlID: "control-1",
		RuleReports: []reporthandling.RuleReport{
			{
				ListInputKinds: []string{},
				RuleResponses: []reporthandling.RuleResponse{
					{ // failed
						AlertObject: reporthandling.AlertObject{
							K8SApiObjects: []map[string]interface{}{
								resources["resource-1"].GetObject(),
								resources["resource-2"].GetObject(),
							},
						},
					},
					{ // passed
						Exception: &armotypes.PostureExceptionPolicy{},
						AlertObject: reporthandling.AlertObject{
							K8SApiObjects: []map[string]interface{}{
								resources["resource-3"].GetObject(),
							},
						},
					},
				},
			},
		},
		BaseScore: 0.00,
	}
	wcs, score := s.ControlScore(controlReport, "")

	require.Equal(t, float32(0), wcs)
	require.Equal(t, float32(0), score)
}

func mockPostureReportV2(t testing.TB) (map[string]workloadinterface.IMetadata, *v2.PostureReport) {
	resources := mockResources(t)

	var resourceWithFailed, resourceWithPassed helpers.AllLists
	resourceWithFailed.Append(apis.StatusFailed, "resource-1", "resource-2")
	resourceWithFailed.Append(apis.StatusPassed, "resource-3")
	resourceWithPassed.Append(apis.StatusPassed, "resource-4")

	var resourceWithFailed2, resourceWithPassed2 helpers.AllLists
	resourceWithFailed2.Append(apis.StatusFailed, "resource-5", "resource-6")
	resourceWithFailed2.Append(apis.StatusPassed, "resource-7", "resource-8")
	resourceWithPassed2.Append(apis.StatusPassed, "resource-9", "resource-10")

	report := &v2.PostureReport{
		SummaryDetails: reportsummary.SummaryDetails{
			Frameworks: []reportsummary.FrameworkSummary{
				// total unnormalized score: 82.1
				// total wcs: 160.1
				// expected final score: 51.28%
				{
					// total unnormalized score: 30.1
					// total wcs: 48.1
					// expected score: 62.57%
					Name: "mock-fw-summary-1",
					Controls: reportsummary.ControlSummaries{
						// 2 failed resources:
						// un-normalized : 7 + 7*3*1.1 (deployment with replicas) = 30.1,
						// wcs: 30.1 + 7 = 37.1
						// expected control score: 81.13%
						"summary-1": reportsummary.ControlSummary{
							Name:        "mock-control-1",
							ControlID:   "control-1",
							ResourceIDs: resourceWithFailed,
							ScoreFactor: 7.00,
						},
						// 0 failed resources:
						// un-normalized : 0
						// wcs: 11
						// expected control score: 0%
						"summary-2": reportsummary.ControlSummary{
							Name:        "mock-control-2",
							ControlID:   "control-2",
							ResourceIDs: resourceWithPassed,
							ScoreFactor: 11.00, // 0 failed resources
						},
					},
				},
				{
					// total unnormalized score: 52
					// total wcs: 112
					// expected score: 46.42%
					Name: "mock-fw-summary-2",
					Controls: reportsummary.ControlSummaries{
						// 2 failed resources:
						// un-normalized :  13 + 13*3 (desired number scheduled) = 52
						// wcs: 52+2*13 = 78
						// expected control score: 66.66%
						"summary-1": reportsummary.ControlSummary{
							Name:        "mock-control-3",
							ControlID:   "control-3",
							ResourceIDs: resourceWithFailed2,
							ScoreFactor: 13.00,
						},
						// 0 failed resources:
						// un-normalized : 0
						// wcs: 17 * 2 = 34
						// expected control score: 0%
						"summary-2": reportsummary.ControlSummary{
							Name:        "mock-control-4",
							ControlID:   "control-4",
							ResourceIDs: resourceWithPassed2,
							ScoreFactor: 17.00,
						},
					},
				},
			},
		},
		Results:   []resourcesresults.Result{},
		Resources: []reporthandling.Resource{},
	}

	//  initialize redundant score factors and resources
	controls := make(reportsummary.ControlSummaries, 4)
	for _, fw := range report.SummaryDetails.Frameworks {
		for _, control := range fw.Controls {
			controls[control.GetID()] = control
		}
	}
	report.SummaryDetails.Controls = controls

	return resources, report
}

func mockPostureReportV1(t testing.TB) (map[string]workloadinterface.IMetadata, []reporthandling.FrameworkReport) {
	// mock providing the same scores as with the V2 model, but hydrates a V1 model.
	// NOTE: the weighted result for V1 is different than for V2, because of a difference in how we compute wcs for passed items
	// (in the testcase below, control-4 contributes a different weight in V1 and V2).

	resources := mockResources(t)
	reports := []reporthandling.FrameworkReport{
		// total unnormalized score: 82.1
		// total wcs: 160.1
		// expected final score: 51.28%
		{
			// total unnormalized score: 30.1
			// total wcs: 48.1
			// expected score: 62.57%
			Name: "mock-fw-summary-1",
			ControlReports: []reporthandling.ControlReport{
				// 2 failed resources:
				// un-normalized : 7 + 7*3*1.1 (deployment with replicas) = 30.1,
				// wcs: 30.1 + 7 = 37.1
				// expected control score: 81.13%
				{
					Name:      "mock-control-1",
					ControlID: "control-1",
					RuleReports: []reporthandling.RuleReport{
						{
							ListInputKinds: []string{
								resources["resource-1"].GetID(),
								resources["resource-2"].GetID(),
								resources["resource-3"].GetID(),
							},
							RuleResponses: []reporthandling.RuleResponse{
								{ // failed
									AlertObject: reporthandling.AlertObject{
										K8SApiObjects: []map[string]interface{}{
											resources["resource-1"].GetObject(),
											resources["resource-2"].GetObject(),
										},
									},
								},
								{ // passed
									Exception: &armotypes.PostureExceptionPolicy{},
									AlertObject: reporthandling.AlertObject{
										K8SApiObjects: []map[string]interface{}{
											resources["resource-3"].GetObject(),
										},
									},
								},
							},
						},
					},
					BaseScore: 7.00,
				},
				// 0 failed resources:
				// un-normalized : 0
				// wcs: 11
				// expected control score: 0%
				{
					Name:      "mock-control-2",
					ControlID: "control-2",
					RuleReports: []reporthandling.RuleReport{
						{
							ListInputKinds: []string{
								resources["resource-4"].GetID(),
							},
							RuleResponses: []reporthandling.RuleResponse{
								{ // passed
									Exception: &armotypes.PostureExceptionPolicy{},
									AlertObject: reporthandling.AlertObject{
										K8SApiObjects: []map[string]interface{}{
											resources["resource-4"].GetObject(),
										},
									},
								},
							},
						},
					},
					BaseScore: 11.00, // 0 failed resources
				},
			},
		},
		{
			// total unnormalized score: 52
			// total wcs: 112
			// expected score: 46.42%
			Name: "mock-fw-summary-2",
			ControlReports: []reporthandling.ControlReport{
				{
					// 2 failed resources:
					// un-normalized :  13 + 13*3 (desired number scheduled) = 52
					// wcs: 52+2*13 = 78
					// expected control score: 66.66%
					Name:      "mock-control-3",
					ControlID: "control-3",
					RuleReports: []reporthandling.RuleReport{
						{
							ListInputKinds: []string{
								resources["resource-5"].GetID(),
								resources["resource-6"].GetID(),
								resources["resource-7"].GetID(),
								resources["resource-8"].GetID(),
							},
							RuleResponses: []reporthandling.RuleResponse{
								{ // failed
									AlertObject: reporthandling.AlertObject{
										K8SApiObjects: []map[string]interface{}{
											resources["resource-5"].GetObject(),
											resources["resource-6"].GetObject(),
										},
									},
								},
								{ // passed
									Exception: &armotypes.PostureExceptionPolicy{},
									AlertObject: reporthandling.AlertObject{
										K8SApiObjects: []map[string]interface{}{
											resources["resource-7"].GetObject(),
											resources["resource-8"].GetObject(),
										},
									},
								},
							},
						},
					},
					BaseScore: 13.00,
				},
				// 0 failed resources:
				// un-normalized : 0
				// wcs: 17 * 1 = 17 - Whereas V2 yields 17 * 2 = 34
				// expected control score: 0%
				{
					Name:      "mock-control-4",
					ControlID: "control-4",
					RuleReports: []reporthandling.RuleReport{
						{
							ListInputKinds: []string{
								resources["resource-9"].GetID(),
								resources["resource-10"].GetID(),
							},
							RuleResponses: []reporthandling.RuleResponse{
								{ // passed
									Exception: &armotypes.PostureExceptionPolicy{},
									AlertObject: reporthandling.AlertObject{
										K8SApiObjects: []map[string]interface{}{
											resources["resource-9"].GetObject(),
										},
									},
								},
								{ // passed
									Exception: &armotypes.PostureExceptionPolicy{},
									AlertObject: reporthandling.AlertObject{
										K8SApiObjects: []map[string]interface{}{
											resources["resource-10"].GetObject(),
										},
									},
								},
							},
						},
					},
					BaseScore: 17.00,
				},
			},
		},
	}

	resourcesWithID := make(map[string]workloadinterface.IMetadata, len(resources))
	for _, resource := range resources {
		resourcesWithID[resource.GetID()] = resource
	}

	return resourcesWithID, reports
}

func mockResources(t testing.TB) map[string]workloadinterface.IMetadata {
	// Declare the resources referred to by the mock reports.
	//
	// Mock k8s objects are retrieved from the fixture in the ./mocks folder.

	return map[string]workloadinterface.IMetadata{
		"resource-1":  reporthandling.NewResource(mocks.GetResourceByType(t, "Role", mocks.WithName("resource-1"))),
		"resource-2":  reporthandling.NewResource(mocks.GetResourceByType(t, "Deployment", mocks.WithName("resource-2"))),
		"resource-3":  reporthandling.NewResource(mocks.GetResourceByType(t, "Pod", mocks.WithName("resource-3"))),
		"resource-4":  reporthandling.NewResource(mocks.GetResourceByType(t, "Pod", mocks.WithName("resource-4"))),
		"resource-5":  reporthandling.NewResource(mocks.GetResourceByType(t, "Secret", mocks.WithName("resource-5"))),
		"resource-6":  reporthandling.NewResource(mocks.GetResourceByType(t, "DaemonSet", mocks.WithName("resource-6"), mocks.WithDesiredNumberScheduled(3))),
		"resource-7":  reporthandling.NewResource(mocks.GetResourceByType(t, "Pod", mocks.WithName("resource-7"))),
		"resource-8":  reporthandling.NewResource(mocks.GetResourceByType(t, "Pod", mocks.WithName("resource-8"))),
		"resource-9":  reporthandling.NewResource(mocks.GetResourceByType(t, "Pod", mocks.WithName("resource-9"))),
		"resource-10": reporthandling.NewResource(mocks.GetResourceByType(t, "Pod", mocks.WithName("resource-10"))),
	}
}

// ================================ compliance score tests ================================

func TestGetControlComplianceScore(t *testing.T) {
	var resourceWithFailed, resourceWithPassed helpers.AllLists
	resourceWithFailed.Append(apis.StatusFailed, "resource-1", "resource-2")
	resourceWithFailed.Append(apis.StatusPassed, "resource-3")
	resourceWithPassed.Append(apis.StatusPassed, "resource-4")

	var resourceWithFailed2, resourceWithPassed2 helpers.AllLists
	resourceWithFailed2.Append(apis.StatusFailed, "resource-5", "resource-6")
	resourceWithFailed2.Append(apis.StatusPassed, "resource-7", "resource-8")
	resourceWithPassed2.Append(apis.StatusPassed, "resource-9", "resource-10")
	t.Parallel()

	t.Run("with empty control report", func(t *testing.T) {
		t.Parallel()

		resources := mockResources(t)
		s := ScoreUtil{isDebugMode: true, resources: resources}
		controlReport := reportsummary.ControlSummary{
			Name:        "empty",
			ControlID:   "empty",
			ResourceIDs: helpers.AllLists{},
		}

		require.Equal(t, float32(0), s.GetControlComplianceScore(&controlReport, ""),
			"empty control report should return a score equals to 0",
		)
	})

	t.Run("with control report", func(t *testing.T) {
		t.Parallel()

		resources := mockResources(t)
		s := ScoreUtil{isDebugMode: true, resources: resources}
		controlReport := reportsummary.ControlSummary{
			Name:        "mock-control-1",
			ControlID:   "mock-control-1",
			ResourceIDs: resourceWithFailed2,
		}

		require.Equal(t, float32(50), s.GetControlComplianceScore(&controlReport, ""),
			"control report should return a score equals to 50",
		)
	})
}

func TestSetPostureReportComplianceScores(t *testing.T) {
	t.Parallel()

	t.Run("with empty report", func(t *testing.T) {
		t.Parallel()

		s := NewScore(map[string]workloadinterface.IMetadata{})
		report := &v2.PostureReport{
			SummaryDetails: reportsummary.SummaryDetails{Frameworks: []reportsummary.FrameworkSummary{{Name: "empty", Controls: reportsummary.ControlSummaries{}}}},
			Results:        []resourcesresults.Result{},
			Resources:      []reporthandling.Resource{},
		}

		require.Errorf(t, s.SetPostureReportComplianceScores(report),
			"empty framework should return an error",
		)

		require.Equal(t, float32(0), report.SummaryDetails.Frameworks[0].Score,
			"empty framework should return an error and have a score equals to 0",
		)
	})

	t.Run("with skipped report", func(t *testing.T) {
		t.Parallel()

		s := NewScore(map[string]workloadinterface.IMetadata{})
		report := &v2.PostureReport{
			SummaryDetails: reportsummary.SummaryDetails{Frameworks: []reportsummary.FrameworkSummary{{Name: "skipped", Controls: reportsummary.ControlSummaries{
				"skipped1": reportsummary.ControlSummary{
					Name:        "skipped1",
					ControlID:   "Skippie1",
					Description: "skipper",
				},
				"skipped2": reportsummary.ControlSummary{
					Name:        "skipped2",
					ControlID:   "Skippie2",
					Description: "skipper",
				},
			}}}},
			Results:   []resourcesresults.Result{},
			Resources: []reporthandling.Resource{},
		}

		require.Errorf(t, s.SetPostureReportComplianceScores(report),
			"empty framework should return an error",
		)

		require.Equal(t, float32(0), report.SummaryDetails.Frameworks[0].Score,
			"empty framework should return an error and have a score equals to 0",
		)
	})

	t.Run("with mock report", func(t *testing.T) {
		t.Parallel()

		resources, report := mockPostureReportV2(t)
		s := ScoreUtil{
			isDebugMode: true,
			resources:   resources,
		}

		require.NoErrorf(t, s.SetPostureReportComplianceScores(report),
			"mock framework should not return an error",
		)

		const (
			expectedScoreFramework1           = float32(62.577965)
			expectedScoreFramework2           = float32(46.42857)
			expectedComplianceScoreFramework1 = float32(66.66667)
			expectedComplianceScoreFramework2 = float32(75)
			expectedSummary                   = float32(70.833336)
		)

		t.Run("assert control scores", func(t *testing.T) {
			require.Len(t, report.SummaryDetails.Controls, 4)
			for _, control := range report.SummaryDetails.Controls {
				var expectedForControl float64

				switch control.ControlID {
				case "control-1":
					expectedForControl = 33.333336
				case "control-2":
					expectedForControl = 100 // passed
				case "control-3":
					expectedForControl = 50
				case "control-4":
					expectedForControl = 100 // passed
				}

				assert.InDeltaf(t, expectedForControl, control.Score, 1e-6,
					"unexpected summarized score for control %q", control.ControlID,
				)
			}
		})

		t.Run("assert framework scores", func(t *testing.T) {
			assert.InDeltaf(t, expectedScoreFramework1, report.SummaryDetails.Frameworks[0].Score, 1e-6,
				"unexpected summarized score for framework[0]",
			)
			assert.InDeltaf(t, expectedScoreFramework2, report.SummaryDetails.Frameworks[1].Score, 1e-6,
				"unexpected summarized score for framework[1]",
			)
		})

		t.Run("assert framework compliance scores", func(t *testing.T) {
			assert.InDeltaf(t, expectedComplianceScoreFramework1, report.SummaryDetails.Frameworks[0].ComplianceScore, 1e-6,
				"unexpected summarized compliance score for framework[0]",
			)
			assert.InDeltaf(t, expectedComplianceScoreFramework2, report.SummaryDetails.Frameworks[1].ComplianceScore, 1e-6,
				"unexpected summarized compliance score for framework[1]",
			)
		})

		t.Run("assert final score", func(t *testing.T) {
			assert.InDeltaf(t, expectedSummary, report.SummaryDetails.ComplianceScore, 1e-6,
				"unexpected summarized final score",
			)
		})
	})
}
