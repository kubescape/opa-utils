package resourcesresults

import (
	"testing"

	"github.com/kubescape/opa-utils/reporthandling/apis"
	"github.com/stretchr/testify/assert"
)

func TestSetGetControlID(t *testing.T) {
	r := ResourceAssociatedControl{}
	id := "C-0078"
	r.SetID(id)
	assert.Equal(t, id, r.GetID())
}

func TestControlStatus(t *testing.T) {
	r := mockResourceAssociatedControl0087Failed()
	assert.Equal(t, apis.StatusFailed, r.GetStatus(nil).Status())
	assert.True(t, r.GetStatus(nil).IsFailed())
	assert.False(t, r.GetStatus(nil).IsPassed())
	assert.False(t, r.GetStatus(nil).IsExcluded())
	assert.False(t, r.GetStatus(nil).IsSkipped())

	r2 := mockResourceAssociatedControl0089Passed()
	assert.Equal(t, apis.StatusPassed, r2.GetStatus(nil).Status())
	assert.True(t, r2.GetStatus(nil).IsPassed())
	assert.False(t, r2.GetStatus(nil).IsFailed())
	assert.False(t, r2.GetStatus(nil).IsExcluded())
	assert.False(t, r2.GetStatus(nil).IsSkipped())

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
