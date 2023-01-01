package resourcesresults

import (
	"testing"

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
