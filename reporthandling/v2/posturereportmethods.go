package v2

import (
	"fmt"
	"time"

	"github.com/armosec/utils-go/str"
	"github.com/kubescape/opa-utils/reporthandling/apis"
	helpersv1 "github.com/kubescape/opa-utils/reporthandling/helpers/v1"
	"github.com/kubescape/opa-utils/reporthandling/results/v1/reportsummary"
	"github.com/kubescape/opa-utils/reporthandling/results/v1/resourcesresults"
)

// Status get the overall scanning status
func (postureReport *PostureReport) GetStatus() *helpersv1.Status {
	return postureReport.SummaryDetails.GetStatus()
}

// =========================================== List resources ====================================

func (postureReport *PostureReport) ListResourcesIDs(f *helpersv1.Filters) *helpersv1.AllLists {
	return postureReport.SummaryDetails.ListResourcesIDs(nil)
}

// =========================================== List Frameworks ====================================

// ListFrameworksNames list all framework policies summary
func (postureReport *PostureReport) ListFrameworks() []reportsummary.IFrameworkSummary {
	return postureReport.SummaryDetails.ListFrameworks()
}

// ListFrameworksNames list all frameworks names
func (postureReport *PostureReport) ListFrameworksNames() *helpersv1.AllLists {
	return postureReport.SummaryDetails.ListFrameworksNames()
}

// =========================================== List Controls ====================================
// ListControls list all controls policies summary
func (postureReport *PostureReport) ListControls() []reportsummary.IControlSummary {
	return postureReport.SummaryDetails.ListControls()
}

// ListControlsNames list all controls names
func (postureReport *PostureReport) ListControlsNames() *helpersv1.AllLists {
	return postureReport.SummaryDetails.ListControlsNames()
}

// ListControlsIDs list all controls names
func (postureReport *PostureReport) ListControlsIDs() *helpersv1.AllLists {
	return postureReport.SummaryDetails.ListControlsIDs()
}

// ==================================== Resource =============================================

// ResourceStatus get single resource status. If resource not found will return an empty string
func (postureReport *PostureReport) ResourceStatus(resourceID string, f *helpersv1.Filters) apis.IStatus {
	for i := range postureReport.Results {
		if postureReport.Results[i].ResourceID == resourceID {
			return postureReport.Results[i].GetStatus(f)
		}
	}
	return helpersv1.NewStatus(apis.StatusUnknown)
}

// ResourceResult get the result of a single resource. If resource not found will return nil
func (postureReport *PostureReport) ResourceResult(resourceID string) *resourcesresults.Result {
	for i := range postureReport.Results {
		if postureReport.Results[i].ResourceID == resourceID {
			return &postureReport.Results[i]
		}
	}
	return nil
}

// UpdateSummary get the result of a single resource. If resource not found will return nil
func (postureReport *PostureReport) InitializeSummary() {

	for i := range postureReport.Results {
		postureReport.AppendResourceResultToSummary(&postureReport.Results[i])
	}
	postureReport.SummaryDetails.InitResourcesSummary(nil)
}

// AppendResourceResultToSummary get the result of a single resource. If resource not found will return nil
func (postureReport *PostureReport) AppendResourceResultToSummary(resourceResult *resourcesresults.Result) {
	postureReport.SummaryDetails.AppendResourceResult(resourceResult)
}

// ========IPostureReport implementations =======
func (postureReport *PostureReport) GetCustomerGUID() string {
	return postureReport.CustomerGUID
}

func (postureReport *PostureReport) GetClusterName() string {
	return postureReport.ClusterName
}

func (postureReport *PostureReport) GetReportGUID() string {
	return postureReport.ReportID
}

func (postureReport *PostureReport) GetJobID() string {
	return postureReport.JobID
}

func (postureReport *PostureReport) GetTimestamp() *time.Time {
	return &postureReport.ReportGenerationTime
}

func (postureReport *PostureReport) GetScannigTarget() ScanningTarget {
	return postureReport.Metadata.ScanMetadata.ScanningTarget
}

func (postureReport *PostureReport) GetContextMetadata() *ContextMetadata {
	return &postureReport.Metadata.ContextMetadata
}

func (postureReport *PostureReport) GetRepositoryHash() string {
	contextMetadata := postureReport.GetContextMetadata()
	if contextMetadata == nil || contextMetadata.RepoContextMetadata == nil {
		return ""
	}

	repoIdentifier := fmt.Sprintf("%s/%s/%s/%s",
		contextMetadata.RepoContextMetadata.Provider,
		contextMetadata.RepoContextMetadata.Owner,
		contextMetadata.RepoContextMetadata.Repo,
		contextMetadata.RepoContextMetadata.Branch)

	return str.AsFNVHash(repoIdentifier)
}
