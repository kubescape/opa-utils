package resourcesresults

import (
	"encoding/json"
	"testing"

	_ "embed"

	"github.com/kubescape/opa-utils/reporthandling"
	"github.com/kubescape/opa-utils/reporthandling/apis"
	"github.com/stretchr/testify/assert"
)

func TestSetGetControlID(t *testing.T) {
	r := ResourceAssociatedControl{}
	id := "C-0078"
	r.SetID(id)
	assert.Equal(t, id, r.GetID())
}

func TestSetStatus(t *testing.T) {
	r1 := mockResourceAssociatedControlPassed()
	r1.SetStatus(reporthandling.Control{})
	assert.Equal(t, apis.StatusPassed, r1.GetStatus(nil).Status())
	assert.Equal(t, apis.SubStatusUnknown, r1.GetSubStatus())
	assert.True(t, r1.GetStatus(nil).IsPassed())
	assert.False(t, r1.GetStatus(nil).IsFailed())
	assert.False(t, r1.GetStatus(nil).IsSkipped())

	r2 := mockResourceAssociatedControlFailed()
	r2.SetStatus(reporthandling.Control{})
	assert.Equal(t, apis.StatusFailed, r2.GetStatus(nil).Status())
	assert.Equal(t, apis.SubStatusUnknown, r2.GetSubStatus())
	assert.False(t, r2.GetStatus(nil).IsPassed())
	assert.True(t, r2.GetStatus(nil).IsFailed())
	assert.False(t, r2.GetStatus(nil).IsSkipped())

	r3 := mockResourceAssociatedControlException()
	r3.SetStatus(reporthandling.Control{})
	assert.Equal(t, apis.StatusPassed, r3.GetStatus(nil).Status())
	assert.Equal(t, apis.SubStatusException, r3.GetSubStatus())
	assert.True(t, r3.GetStatus(nil).IsPassed())
	assert.False(t, r3.GetStatus(nil).IsFailed())
	assert.False(t, r3.GetStatus(nil).IsSkipped())

	r4 := mockResourceAssociatedControlConfiguration()
	r4.SetStatus(*mockControlWithActionRequiredConfiguration())
	assert.Equal(t, apis.StatusSkipped, r4.GetStatus(nil).Status())
	assert.Equal(t, apis.SubStatusConfiguration, r4.GetSubStatus())
	assert.False(t, r4.GetStatus(nil).IsPassed())
	assert.False(t, r4.GetStatus(nil).IsFailed())
	assert.True(t, r4.GetStatus(nil).IsSkipped())

	r5 := mockResourceAssociatedControlFailed()
	r5.SetStatus(*mockControlWithActionRequiredManualReview())
	assert.Equal(t, apis.StatusSkipped, r5.GetStatus(nil).Status())
	assert.Equal(t, apis.SubStatusManualReview, r5.GetSubStatus())
	assert.False(t, r5.GetStatus(nil).IsPassed())
	assert.False(t, r5.GetStatus(nil).IsFailed())
	assert.True(t, r5.GetStatus(nil).IsSkipped())

	r6 := mockResourceAssociatedControlFailed()
	r6.SetStatus(*mockControlWithActionRequiredRequiresReview())
	assert.Equal(t, apis.StatusSkipped, r6.GetStatus(nil).Status())
	assert.Equal(t, apis.SubStatusRequiresReview, r6.GetSubStatus())
	assert.False(t, r6.GetStatus(nil).IsPassed())
	assert.False(t, r6.GetStatus(nil).IsFailed())
	assert.True(t, r6.GetStatus(nil).IsSkipped())

}

func TestResourceAssociatedControl_SetName(t *testing.T) {
	type fields struct {
		ControlID               string
		Name                    string
		ResourceAssociatedRules []ResourceAssociatedRule
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "TestResourceAssociatedControl_SetName",
			fields: struct {
				ControlID               string
				Name                    string
				ResourceAssociatedRules []ResourceAssociatedRule
			}{
				ControlID:               "C-0078",
				Name:                    "TestResourceAssociatedControl_SetName",
				ResourceAssociatedRules: []ResourceAssociatedRule{},
			},
			args: struct{ name string }{name: "TestResourceAssociatedControl_SetName"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			control := &ResourceAssociatedControl{
				ControlID:               tt.fields.ControlID,
				Name:                    tt.fields.Name,
				ResourceAssociatedRules: tt.fields.ResourceAssociatedRules,
			}
			control.SetName(tt.args.name)
		})
	}
}

//go:embed testdata/old_resource_associated_controls.json
var oldResourceAssociatedControlsTestData []byte

func Test_GetStatusAndSubStatus_OldResourceAssociatedControls(t *testing.T) {
	controls := []ResourceAssociatedControl{}
	err := json.Unmarshal([]byte(oldResourceAssociatedControlsTestData), &controls)
	assert.NoError(t, err)

	controlIdToExpectedStatus := map[string]string{
		"C-0054": "passed", // passed w/ exception
		"C-0067": "failed", // failed
		"C-0002": "passed", // passed
	}

	controlIdToExpectedSubStatus := map[string]string{
		"C-0054": "w/exceptions", // passed w/ exception
		"C-0067": "",             // failed
		"C-0002": "",             // passed
	}

	for _, control := range controls {
		assert.True(t, control.isOldControl())
		assert.Equal(t, controlIdToExpectedStatus[control.ControlID], string(control.GetStatus(nil).Status()))
		assert.Equal(t, controlIdToExpectedSubStatus[control.ControlID], string(control.GetSubStatus()))
	}
}

//go:embed testdata/new_resource_associated_controls.json
var newResourceAssociatedControlsTestData []byte

func Test_GetStatusAndSubStatus_NewResourceAssociatedControls(t *testing.T) {
	controls := []ResourceAssociatedControl{}
	err := json.Unmarshal([]byte(newResourceAssociatedControlsTestData), &controls)
	assert.NoError(t, err)

	controlIdToExpectedStatus := map[string]string{
		"C-0053": "passed", // passed w/ exception
		"C-0014": "passed", // passed
		"C-0212": "failed", // failed
	}

	controlIdToExpectedSubStatus := map[string]string{
		"C-0053": "w/exceptions", // passed w/ exception
		"C-0014": "",             // passed
		"C-0212": "",             // failed
	}

	for _, control := range controls {
		assert.False(t, control.isOldControl())
		assert.Equal(t, controlIdToExpectedStatus[control.ControlID], string(control.GetStatus(nil).Status()))
		assert.Equal(t, controlIdToExpectedSubStatus[control.ControlID], string(control.GetSubStatus()))
	}
}

func TestControlMissingAllConfigurations(t *testing.T) {
	tests := []struct {
		name    string
		control *ResourceAssociatedControl
		want    bool
	}{
		{
			name: "TestControlNoConfigurations",
			control: &ResourceAssociatedControl{
				ResourceAssociatedRules: []ResourceAssociatedRule{
					{
						ControlConfigurations: map[string][]string{},
					},
				},
			},
			want: true,
		}, {
			name: "TestControlOneEmptyConfiguration",
			control: &ResourceAssociatedControl{
				ResourceAssociatedRules: []ResourceAssociatedRule{
					{
						ControlConfigurations: map[string][]string{
							"EmptyConfiguration": {},
						},
					},
				},
			},
			want: true,
		},
		{
			name: "TestControlOneNonEmptyConfiguration",
			control: &ResourceAssociatedControl{
				ResourceAssociatedRules: []ResourceAssociatedRule{
					{
						ControlConfigurations: map[string][]string{
							"NonEmptyConfiguration": {
								"key", "value",
							},
						},
					},
				},
			},
			want: false,
		},
		{
			name: "TestControlMultipleConfigurations",
			control: &ResourceAssociatedControl{
				ResourceAssociatedRules: []ResourceAssociatedRule{
					{
						ControlConfigurations: map[string][]string{
							"EmptyConfiguration": {},
							"NonEmptyConfiguration": {
								"key", "value",
							},
						},
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := controlMissingAllConfigurations(tt.control); got != tt.want {
				t.Errorf("Control.missingAllConfigurations() = %v, want %v", got, tt.want)
			}
		})
	}
}
