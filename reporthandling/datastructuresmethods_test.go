package reporthandling

import (
	"encoding/json"
	"testing"

	"github.com/armosec/opa-utils/reporthandling/mock"
	"github.com/stretchr/testify/assert"
)

func FrameworkResultsMock(report string) (*FrameworkReport, error) {
	frameworkReport := &FrameworkReport{}
	if err := json.Unmarshal([]byte(report), frameworkReport); err != nil {
		return nil, err
	}
	return frameworkReport, nil
}

func ControlsResultsMock(report string) ([]ControlReport, error) {
	f, err := FrameworkResultsMock(report)
	if err != nil {
		return nil, err
	}
	return f.ControlReports, nil
}

func TestControlsResults(t *testing.T) {
	controls, err := ControlsResultsMock(mock.NSAScanWithExceptions)
	assert.NoError(t, err, err)
	assert.Equal(t, len(controls), 21)

	for _, control := range controls {
		switch control.ControlID {
		// case "C-0005":
		// 	assert.Equal(t, 39, control.GetNumberOfResources(), "C-0005: GetNumberOfResources", control.GetNumberOfResources())
		// 	assert.Equal(t, 0, control.GetNumberOfFailedResources(), "C-0005: GetNumberOfFailedResources")
		// 	assert.Equal(t, 0, control.GetNumberOfWarningResources(), "C-0005: GetNumberOfWarningResources")
		// 	assert.Equal(t, 0, len(control.ListControlsInputKinds()), "C-0005: ListControlsInputKinds")
		// 	assert.True(t, control.Passed(), "C-0005: Passed")
		// 	assert.False(t, control.Warning(), "C-0005: Warning")
		// 	assert.False(t, control.Failed(), "C-0005: Failed")
		// case "C-0038":
		// 	assert.Equal(t, 39, control.GetNumberOfResources(), "C-0038: GetNumberOfResources", control.GetNumberOfResources())
		// 	assert.Equal(t, 0, control.GetNumberOfFailedResources(), "C-0038: GetNumberOfFailedResources")
		// 	assert.Equal(t, 0, control.GetNumberOfWarningResources(), "C-0038: GetNumberOfWarningResources")
		// 	assert.Equal(t, 0, len(control.ListControlsInputKinds()), "C-0038: ListControlsInputKinds")
		// 	assert.True(t, control.Passed(), "C-0038: Passed")
		// 	assert.False(t, control.Warning(), "C-0038: Warning")
		// 	assert.False(t, control.Failed(), "C-0038: Failed")
		case "C-0017": // TODO - test
			assert.Equal(t, 39, control.GetNumberOfResources(), "C-0017: GetNumberOfResources")
			assert.Equal(t, 30, control.GetNumberOfFailedResources(), "C-0017: GetNumberOfFailedResources")
			assert.Equal(t, 10, control.GetNumberOfWarningResources(), "C-0017: GetNumberOfWarningResources")
			assert.Equal(t, 0, len(control.ListControlsInputKinds()), "C-0017: ListControlsInputKinds")
			assert.False(t, control.Passed(), "C-0017: Passed")
			assert.False(t, control.Warning(), "C-0017: Warning")
			assert.True(t, control.Failed(), "C-0017: Failed")
		case "C-0009": // TODO - test
			assert.Equal(t, 46, control.GetNumberOfResources(), "C-0009: GetNumberOfResources")
			assert.Equal(t, 0, control.GetNumberOfFailedResources(), "C-0009: GetNumberOfFailedResources")
			assert.Equal(t, 14, control.GetNumberOfWarningResources(), "C-0009: GetNumberOfWarningResources")
			assert.Equal(t, 0, len(control.ListControlsInputKinds()), "C-0009: ListControlsInputKinds")
			assert.False(t, control.Passed(), "C-0009: Passed")
			assert.True(t, control.Warning(), "C-0009: Warning")
			assert.False(t, control.Failed(), "C-0009: Failed")
		case "C-0030": // TODO - test
			assert.Equal(t, 39, control.GetNumberOfResources(), "C-0030: GetNumberOfResources")
			assert.Equal(t, 29, control.GetNumberOfFailedResources(), "C-0030: GetNumberOfFailedResources")
			assert.Equal(t, 10, control.GetNumberOfWarningResources(), "C-0030: GetNumberOfWarningResources")
			assert.Equal(t, 0, len(control.ListControlsInputKinds()), "C-0030: ListControlsInputKinds")
			assert.False(t, control.Passed(), "C-0030: Passed")
			assert.False(t, control.Warning(), "C-0030: Warning")
			assert.True(t, control.Failed(), "C-0030: Failed")
		}
	}
}
