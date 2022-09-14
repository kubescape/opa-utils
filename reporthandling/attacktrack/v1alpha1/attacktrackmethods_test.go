package v1alpha1

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	control_1 = &AttackTrackControlMock{ControlId: "1", Tags: []string{"security"}, Categories: []string{"A"}}
	control_2 = &AttackTrackControlMock{ControlId: "2", Tags: []string{"security"}, Categories: []string{"A", "Z"}}
	control_3 = &AttackTrackControlMock{ControlId: "3", Tags: []string{"security"}, Categories: []string{"Z"}}
	control_4 = &AttackTrackControlMock{ControlId: "4", Tags: []string{"security"}, Categories: []string{"B"}}
	control_5 = &AttackTrackControlMock{ControlId: "5", Tags: []string{"security"}, Categories: []string{"C"}}
	control_6 = &AttackTrackControlMock{ControlId: "6", Tags: []string{"security-impact"}, Categories: []string{"D"}}
	control_7 = &AttackTrackControlMock{ControlId: "7", Tags: []string{"security"}, Categories: []string{"E"}}
	control_8 = &AttackTrackControlMock{ControlId: "8", Tags: []string{"security-impact"}, Categories: []string{"F"}}
)

func TestAttackTrack_IsValid(t *testing.T) {
	tests := []struct {
		name     string
		v        *AttackTrack
		expected bool
	}{
		{
			name: "Cyclic Definition",
			v: AttackTrackMock(
				AttackTrackStep{
					Name: "A",
					SubSteps: []AttackTrackStep{
						{
							Name: "B",
							SubSteps: []AttackTrackStep{
								{
									Name: "A",
								},
							},
						},
					},
				},
			),
			expected: false,
		},
		{
			name: "Duplicate step",
			v: AttackTrackMock(
				AttackTrackStep{
					Name: "A",
					SubSteps: []AttackTrackStep{
						{
							Name: "B",
						},
						{
							Name: "B",
						},
					},
				},
			),
			expected: false,
		},
		{
			name: "Valid Attack Graph",
			v: AttackTrackMock(
				AttackTrackStep{
					Name: "A",
					SubSteps: []AttackTrackStep{
						{
							Name: "B",
						},
						{
							Name: "C",
						},
					},
				},
			),
			expected: true,
		},
		{
			name: "Invalid tree with duplicate step",
			v: AttackTrackMock(
				AttackTrackStep{
					Name: "A",
					SubSteps: []AttackTrackStep{
						{
							Name: "C",
							SubSteps: []AttackTrackStep{
								{
									Name: "B",
								},
								{
									Name: "D",
								},
								{
									Name: "E",
									SubSteps: []AttackTrackStep{
										{
											Name: "D",
										},
									},
								},
							},
						},
					},
				},
			),
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.v.IsValid())
		})
	}
}

func TestIterator(t *testing.T) {
	v := AttackTrackMock(
		AttackTrackStep{
			Name: "A",
			SubSteps: []AttackTrackStep{
				{
					Name: "C",
					SubSteps: []AttackTrackStep{
						{
							Name: "B",
						},
						{
							Name: "D",
						},
						{
							Name: "E",
						},
					},
				},
				{
					Name: "F",
				},
			},
		},
	)

	expected := []string{"A", "F", "C", "E", "D", "B"}
	it := v.Iterator()
	assert.Truef(t, it.HasNext(), "iterator should have next")

	actual := []string{}
	for it.HasNext() {
		actual = append(actual, it.Next().GetName())
	}
	assert.Equalf(t, expected, actual, "iterator should return the correct order. expected: %v, actual: %v", expected, actual)
}

func TestAttackTrack_CalculateAllPaths(t *testing.T) {
	tests := []struct {
		name        string
		attackTrack *AttackTrack
		controlsMap AttackTrackControlsLookup
		want        [][]string
	}{
		{
			name: "Valid Attack Graph",
			attackTrack: AttackTrackMock(
				AttackTrackStep{
					Name: "A",
					SubSteps: []AttackTrackStep{
						{
							Name: "C",
							SubSteps: []AttackTrackStep{
								{
									Name: "B",
								},
								{
									Name: "D",
								},
								{
									Name: "E",
									SubSteps: []AttackTrackStep{
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
			controlsMap: AttackTrackControlsLookup{
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
			want: [][]string{
				{"A", "F"},
				{"E", "G"},
				{"D"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := NewAttackTrackAllPathsHandler(tt.attackTrack, &tt.controlsMap)
			allPaths := handler.CalculateAllPaths()

			assert.Equalf(t, len(tt.want), len(allPaths), "CalculateAllPaths should return the correct number of paths. expected: %v, actual: %v", len(tt.want), len(allPaths))
			for i, path := range allPaths {
				for j, step := range path {
					assert.Equalf(t, tt.want[i][j], step.GetName(), "CalculateAllPaths should return the correct paths. expected: %v, actual: %v", tt.want, allPaths)
				}
			}
		})
	}
}

func TestNewAttackTrackControlsLookup(t *testing.T) {
	attackTrack := AttackTrackMock(
		AttackTrackStep{
			Name: "A",
			SubSteps: []AttackTrackStep{
				{
					Name: "C",
					SubSteps: []AttackTrackStep{
						{
							Name: "B",
						},
						{
							Name: "D",
						},
						{
							Name: "E",
							SubSteps: []AttackTrackStep{
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
	)

	attackTracks := []IAttackTrack{attackTrack}
	failedControls := []string{"1", "2", "3", "4", "5", "6", "7", "8"}
	allControls := map[string]IAttackTrackControl{
		"1": control_1,
		"2": control_2,
		"3": control_3,
		"4": control_4,
		"5": control_5,
		"6": control_6,
		"7": control_7,
		"8": control_8,
	}

	result := NewAttackTrackControlsLookup(attackTracks, failedControls, allControls)
	expected := AttackTrackControlsLookup{
		"TestAttackTrack": {
			"A": []IAttackTrackControl{control_1, control_2},
			"Z": []IAttackTrackControl{control_2, control_3},
			"B": []IAttackTrackControl{control_4},
			"C": []IAttackTrackControl{control_5},
			"D": []IAttackTrackControl{control_6},
			"E": []IAttackTrackControl{control_7},
			"F": []IAttackTrackControl{control_8},
		},
	}

	resultJson, _ := json.Marshal(result)
	expectedJson, _ := json.Marshal(expected)

	assert.JSONEq(t, string(expectedJson), string(resultJson))
}
