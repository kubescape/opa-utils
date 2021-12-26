package v2

import (
	"github.com/armosec/opa-utils/reporthandling/apis"
	helpersv1 "github.com/armosec/opa-utils/reporthandling/helpers/v1"
)

// type Report struct{
// 	postureReport *PostureReport
// 	filter *Filter
// }
func NewPostureReport() *PostureReport {
	return &PostureReport{}
}
func (postureReport *PostureReport) IsPassed(f *helpersv1.FilterPassed) bool {
	return false
}

func (postureReport *PostureReport) IsFailed(f *helpersv1.FilterFailed) bool {
	return false
}

func (postureReport *PostureReport) IsExcluded(f *helpersv1.FilterExcluded) bool {
	return false
}

func (postureReport *PostureReport) Status(f *helpersv1.Filters) apis.ScanningStatus {
	return apis.StatusSkipped
}

func (postureReport *PostureReport) Excluded(f *helpersv1.FilterExcluded) []string {
	return []string{}
}

func (postureReport *PostureReport) Passed(f *helpersv1.FilterPassed) []string {
	return []string{}
}

func (postureReport *PostureReport) Skipped(f *helpersv1.FilterSkipped) []string {
	return []string{}
}

func (postureReport *PostureReport) All(f *helpersv1.Filters) []string {
	return []string{}
}

func (postureReport *PostureReport) Failed(f *helpersv1.FilterFailed) []string {
	return []string{}
}
