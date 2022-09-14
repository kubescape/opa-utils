package prioritization

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetResourceID(t *testing.T) {
	pr := &PrioritizedResource{
		ResourceID:     "aa/bb/cc",
		PriorityVector: []*ControlsVector{},
	}
	assert.Equal(t, "aa/bb/cc", pr.ResourceID)

	pr.SetResourceID("xx")
	assert.Equal(t, "xx", pr.ResourceID)
}

func TestGetResourceID(t *testing.T) {
	pr := &PrioritizedResource{
		ResourceID:     "aa/bb/cc",
		PriorityVector: []*ControlsVector{},
	}
	assert.Equal(t, "aa/bb/cc", pr.GetResourceID())
}

func TestPrioritizedResourceCalculateScore(t *testing.T) {
	type fields struct {
		ResourceID     string
		PriorityVector []*ControlsVector
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "empty priority vector",
			fields: fields{
				ResourceID:     "A",
				PriorityVector: []*ControlsVector{},
			},
			want: 0,
		},
		{
			name: "non-empty priority vector",
			fields: fields{
				ResourceID: "A",
				PriorityVector: []*ControlsVector{
					{
						Type:   ControlPriorityVectorType,
						Score:  4.5,
						Vector: []*PriorityVectorControl{{ControlID: "C1"}, {ControlID: "C2"}},
					},
					{
						Type:   ControlPriorityVectorType,
						Score:  2.5,
						Vector: []*PriorityVectorControl{{ControlID: "C3"}},
					},
				},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pr := &PrioritizedResource{
				ResourceID:     tt.fields.ResourceID,
				PriorityVector: tt.fields.PriorityVector,
			}
			if got := pr.CalculateScore(); got != tt.want {
				t.Errorf("PrioritizedResource.CalculateScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListControlsIDs(t *testing.T) {
	type fields struct {
		ResourceID     string
		PriorityVector []*ControlsVector
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "empty priority vector",
			fields: fields{
				ResourceID:     "A",
				PriorityVector: []*ControlsVector{},
			},
			want: []string{},
		},
		{
			name: "non-empty priority vector",
			fields: fields{
				ResourceID: "A",
				PriorityVector: []*ControlsVector{
					{
						Type:   ControlPriorityVectorType,
						Score:  4.5,
						Vector: []*PriorityVectorControl{{ControlID: "C1"}, {ControlID: "C2"}},
					},
					{
						Type:   ControlPriorityVectorType,
						Score:  2.5,
						Vector: []*PriorityVectorControl{{ControlID: "C3"}},
					},
				},
			},
			want: []string{"C1", "C2", "C3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pr := &PrioritizedResource{
				ResourceID:     tt.fields.ResourceID,
				PriorityVector: tt.fields.PriorityVector,
			}
			if got := pr.ListControlsIDs(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PrioritizedResource.ListControlsIDs() = %v, want %v", got, tt.want)
			}
		})
	}
}
