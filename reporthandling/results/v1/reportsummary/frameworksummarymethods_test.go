package reportsummary

import (
	"reflect"
	"testing"

	"github.com/kubescape/opa-utils/reporthandling/apis"
	"github.com/stretchr/testify/assert"
)

func TestSetStatus(t *testing.T) {
	status := []apis.ScanningStatus{}

	f := mockFrameworkSummaryFailException()
	status = append(status, f.GetStatus().Status())

	f.Status = apis.StatusUnknown
	for k, v := range f.Controls {
		status = append(status, v.GetStatus().Status())
		f.Controls[k] = v
	}

	i := 0
	assert.Equal(t, status[i], f.GetStatus().Status())
	for _, v := range f.Controls {
		i++
		assert.Equal(t, status[i], v.GetStatus().Status())
	}

}

func TestStatusInfoNotPresent(t *testing.T) {

	f := mockSummaryDetailsNoInnerStatus() // Status: skipped , InnerStatus empty
	for _, v := range f.Controls {
		status := v.GetStatus()
		assert.Equal(t, reflect.TypeOf(status), reflect.TypeOf(&apis.StatusInfo{}))
		assert.Equal(t, status.Status(), apis.StatusSkipped)
		assert.Equal(t, status.Info(), "")
	}

}

func TestStatusEmpty(t *testing.T) {

	f := mockSummaryDetailsStatusEmpty()
	for _, v := range f.Controls {
		v.Status = apis.StatusPassed
		v.StatusInfo.SubStatus = apis.SubStatusIrrelevant
		status := v.GetStatus()
		subStatus := v.GetSubStatus()
		assert.Equal(t, reflect.TypeOf(status), reflect.TypeOf(&apis.StatusInfo{}))
		assert.Equal(t, reflect.TypeOf(subStatus), reflect.TypeOf(apis.SubStatusIrrelevant))
		assert.Equal(t, status.Status(), apis.StatusPassed)
		assert.Equal(t, subStatus, apis.SubStatusIrrelevant)
		assert.Equal(t, status.Info(), "")
	}

}

func TestStatusInfoSkipped(t *testing.T) {
	var status apis.ScanningStatus
	var info string

	f := mockSummaryDetailsStatusSkipped() // control -> status: "skipped", info: "no host sensor flag"

	for _, v := range f.Controls {
		status = v.GetStatus().Status()
		info = v.GetStatus().Info()
		assert.Equal(t, status, apis.StatusSkipped)
		assert.Equal(t, info, "no host sensor flag")
	}

}

func TestStatusIrrelevant(t *testing.T) {
	var status apis.ScanningStatus
	var subStatus apis.ScanningSubStatus

	f := mockSummaryDetailsStatusIrrelevant() // control -> status: "passed", subStatus: "irrelevant"

	for _, v := range f.Controls {
		status = v.GetStatus().Status()
		subStatus = v.GetSubStatus()
		assert.Equal(t, status, apis.StatusPassed)
		assert.Equal(t, subStatus, apis.SubStatusIrrelevant)
	}

}

func TestFrameworkControlsSummariesCounters(t *testing.T) {
	f := mockFrameworkSummaryFailPass()
	assert.Equal(t, len(f.Controls), f.GetControls().NumberOfControls().All(), "invalid total control count")
	// assert.Equal(t, len(f.GetControls().NumberOfControls().), f.GetControls().NumberOfControls().Failed(), "invalid total failed control count")
	assert.Equal(t, f.GetControls().ListControlsIDs(nil).Failed(), f.GetControls().NumberOfControls().Failed(), "invalid total failed control count")
	assert.Equal(t, f.GetControls().ListControlsIDs(nil).Passed(), f.GetControls().NumberOfControls().Passed(), "invalid total passed control count")
	assert.Equal(t, f.GetControls().ListControlsIDs(nil).Skipped(), f.GetControls().NumberOfControls().Skipped(), "invalid total skipped control count")
}

func TestFrameworkGettingSpecificControl(t *testing.T) {
	f := mockFrameworkSummaryFailPass()
	a := f.GetControls().GetControl(EControlCriteriaID, "1234")
	assert.Nil(t, a, "control id '1234' shouldn't exist")
	assert.Equal(t, "control-fail-pass", f.GetControls().GetControl(EControlCriteriaID, "C-0001").GetName(), "wrong control retrieved")
}
