package score

import (
	"testing"

	"github.com/kubescape/opa-utils/reporthandling/apis"
	"github.com/kubescape/opa-utils/reporthandling/results/v1/reportsummary"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// A framework whose controls were all skipped (e.g. not evaluated or missing
// configuration) has no resources to weigh, so its score cannot be computed.
// That must not stop the other frameworks, the summary and the compliance
// scores from being calculated.
func TestSetPostureReportComplianceScores_UnscorableFrameworkDoesNotBlockOthers(t *testing.T) {
	t.Parallel()

	resources, report := mockPostureReportV2(t)

	skipped := reportsummary.ControlSummary{Name: "skipped", ControlID: "skipped-1"}
	skipped.SetStatus(&apis.StatusInfo{InnerStatus: apis.StatusSkipped, SubStatus: apis.SubStatusConfiguration})
	unscorable := reportsummary.FrameworkSummary{
		Name:     "unscorable",
		Controls: reportsummary.ControlSummaries{"skipped-1": skipped},
	}
	report.SummaryDetails.Frameworks = append([]reportsummary.FrameworkSummary{unscorable}, report.SummaryDetails.Frameworks...)

	s := ScoreUtil{resources: resources}

	err := s.SetPostureReportComplianceScores(report)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unscorable")

	fws := report.SummaryDetails.Frameworks
	assert.Equal(t, float32(0), fws[0].Score)
	assert.Equal(t, float32(0), fws[0].ComplianceScore)

	// same values as the "with mock report" case of TestSetPostureReportComplianceScores
	assert.InDelta(t, float32(62.577965), fws[1].Score, 1e-6)
	assert.InDelta(t, float32(46.42857), fws[2].Score, 1e-6)
	assert.InDelta(t, float32(100), fws[3].Score, 1e-6)
	assert.InDelta(t, float32(66.66667), fws[1].ComplianceScore, 1e-6)
	assert.InDelta(t, float32(75), fws[2].ComplianceScore, 1e-6)
	assert.InDelta(t, float32(100), fws[3].ComplianceScore, 1e-6)

	assert.InDelta(t, float32(51.280453), report.SummaryDetails.Score, 1e-6)
	assert.InDelta(t, float32(76.66667), report.SummaryDetails.ComplianceScore, 1e-6)
	for id, ctrl := range report.SummaryDetails.Controls {
		require.NotNilf(t, ctrl.ComplianceScore, "control %q has no compliance score", id)
	}
}

// Every unscorable framework is reported, not just the first one.
func TestCalculatePostureReportV2_ReportsEveryUnscorableFramework(t *testing.T) {
	t.Parallel()

	resources, report := mockPostureReportV2(t)
	report.SummaryDetails.Frameworks = append(report.SummaryDetails.Frameworks,
		reportsummary.FrameworkSummary{Name: "empty-a", Controls: reportsummary.ControlSummaries{}},
		reportsummary.FrameworkSummary{Name: "empty-b", Controls: reportsummary.ControlSummaries{}},
	)

	s := ScoreUtil{resources: resources}

	err := s.CalculatePostureReportV2(report)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "empty-a")
	assert.Contains(t, err.Error(), "empty-b")
	assert.InDelta(t, float32(51.280453), report.SummaryDetails.Score, 1e-6)
}
