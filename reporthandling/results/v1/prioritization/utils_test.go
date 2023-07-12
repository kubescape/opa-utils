package prioritization

import (
	"reflect"
	"testing"

	"github.com/kubescape/opa-utils/reporthandling/attacktrack/v1alpha1"
	"github.com/stretchr/testify/assert"
)

var (
	control_1 = &v1alpha1.AttackTrackControlMock{ControlId: "1", Tags: []string{"security"}, Categories: []string{"A"}}
	control_2 = &v1alpha1.AttackTrackControlMock{ControlId: "2", Tags: []string{"security"}, Categories: []string{"A", "Z"}}
	control_3 = &v1alpha1.AttackTrackControlMock{ControlId: "3", Tags: []string{"security"}, Categories: []string{"Z"}}
	control_4 = &v1alpha1.AttackTrackControlMock{ControlId: "4", Tags: []string{"security"}, Categories: []string{"B"}}
	control_5 = &v1alpha1.AttackTrackControlMock{ControlId: "5", Tags: []string{"security"}, Categories: []string{"C"}}
	control_6 = &v1alpha1.AttackTrackControlMock{ControlId: "6", Tags: []string{"security-impact"}, Categories: []string{"D"}}
	control_7 = &v1alpha1.AttackTrackControlMock{ControlId: "7", Tags: []string{"security"}, Categories: []string{"E"}}
	control_8 = &v1alpha1.AttackTrackControlMock{ControlId: "8", Tags: []string{"security-impact"}, Categories: []string{"F"}}
)

func TestControlsVectorFromAttackTrackPaths(t *testing.T) {
	tests := []struct {
		name        string
		attackTrack v1alpha1.IAttackTrack
		controlsMap v1alpha1.AttackTrackControlsLookup
		expected    []ControlsVector
	}{
		{
			name: "Valid Attack Graph",
			attackTrack: v1alpha1.GetAttackTrackMock(
				v1alpha1.AttackTrackStep{
					Name: "A",
					SubSteps: []v1alpha1.AttackTrackStep{
						{
							Name: "C",
							SubSteps: []v1alpha1.AttackTrackStep{
								{
									Name: "B",
								},
								{
									Name: "D",
								},
								{
									Name: "E",
									SubSteps: []v1alpha1.AttackTrackStep{
										{
											Name: "G",
										},
									},
								},
							},
						},
						{
							Name: "F",
						},
					},
				},
			),
			controlsMap: v1alpha1.AttackTrackControlsLookup{
				"TestAttackTrack": {
					"A": {
						control_1, control_2,
					},
					"F": {
						control_3, control_4,
					},
					"D": {
						control_5,
					},
					"E": {
						control_6,
					},
					"G": {
						control_7, control_8,
					},
				},
			},
			expected: []ControlsVector{
				{
					AttackTrackName: "TestAttackTrack",
					Type:            "control",
					Score:           0,
					Severity:        0,
					Vector: []PriorityVectorControl{
						{ControlID: "1", Category: "A", Tags: []string{"security"}},
						{ControlID: "3", Category: "F", Tags: []string{"security"}},
					},
				},
				{
					AttackTrackName: "TestAttackTrack",
					Type:            "control",
					Score:           0,
					Severity:        0,
					Vector: []PriorityVectorControl{
						{ControlID: "1", Category: "A", Tags: []string{"security"}},
						{ControlID: "4", Category: "F", Tags: []string{"security"}},
					},
				},
				{
					AttackTrackName: "TestAttackTrack",
					Type:            "control",
					Score:           0,
					Severity:        0,
					Vector: []PriorityVectorControl{
						{ControlID: "2", Category: "A", Tags: []string{"security"}},
						{ControlID: "3", Category: "F", Tags: []string{"security"}},
					},
				},
				{
					AttackTrackName: "TestAttackTrack",
					Type:            "control",
					Score:           0,
					Severity:        0,
					Vector: []PriorityVectorControl{
						{ControlID: "2", Category: "A", Tags: []string{"security"}},
						{ControlID: "4", Category: "F", Tags: []string{"security"}},
					},
				},
				{
					AttackTrackName: "TestAttackTrack",
					Type:            "control",
					Score:           0,
					Severity:        0,
					Vector: []PriorityVectorControl{
						{ControlID: "6", Category: "E", Tags: []string{"security-impact"}},
						{ControlID: "7", Category: "G", Tags: []string{"security"}},
					},
				},
				{
					AttackTrackName: "TestAttackTrack",
					Type:            "control",
					Score:           0,
					Severity:        0,
					Vector: []PriorityVectorControl{
						{ControlID: "5", Category: "D", Tags: []string{"security"}},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := v1alpha1.NewAttackTrackAllPathsHandler(tt.attackTrack, &tt.controlsMap)
			allPaths := handler.CalculateAllPaths()

			result := ControlsVectorFromAttackTrackPaths(tt.attackTrack, allPaths)

			if assert.Equalf(t, len(tt.expected), len(result), "ControlsVectorFromAttackTrackPaths should return the correct number of vectors. expected: %v, actual: %v", len(tt.expected), len(result)) {
				for i := range tt.expected {
					if !reflect.DeepEqual(result[i], tt.expected[i]) {
						t.Errorf("result = %v, expected %v", result[i], tt.expected[i])
					}
				}
			}
		})
	}
}
