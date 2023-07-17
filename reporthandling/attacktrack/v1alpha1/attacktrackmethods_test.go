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
		v        IAttackTrack
		name     string
		expected bool
	}{
		{
			name: "Cyclic Definition",
			v: GetAttackTrackMock(
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
			v: GetAttackTrackMock(
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
			v: GetAttackTrackMock(
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
			v: GetAttackTrackMock(
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
	v := GetAttackTrackMock(
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
		attackTrack IAttackTrack
		controlsMap AttackTrackControlsLookup
		want        [][]string
	}{
		{
			name: "Valid Attack Graph",
			attackTrack: GetAttackTrackMock(
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
	attackTrack := GetAttackTrackMock(
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

func TestCalculatePathsRootToLeaf(t *testing.T) {
	tests := []struct {
		name        string
		attackTrack IAttackTrack
		controlsMap AttackTrackControlsLookup
		want        [][]string
	}{
		{
			name: "Found attack chain",
			attackTrack: GetAttackTrackMock(
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
					"C": {
						control_5,
					},
					"B": {
						control_6,
					},
					"G": {
						control_7, control_8,
					},
				},
			},
			want: [][]string{
				{"A", "C", "B"},
				{"A", "F"},
			},
		},
		{
			name: "No attack chain",
			attackTrack: GetAttackTrackMock(
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
					"B": {
						control_6,
					},
				},
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := NewAttackTrackAllPathsHandler(tt.attackTrack, &tt.controlsMap)
			paths := handler.CalculatePathsRootToLeaf()

			if tt.want != nil || paths != nil {
				assert.Equalf(t, len(tt.want), len(paths), "CalculatePathsRootToLeaf should return the correct number of paths. expected: %v, actual: %v", len(tt.want), len(paths))
				for i, path := range paths {
					for j, step := range path {
						assert.Equalf(t, tt.want[i][j], step.GetName(), "CalculatePathsRootToLeaf should return the correct paths. expected: %v, actual: %v", tt.want, paths)
					}
				}

			}

		})
	}
}

func TestGenerateAttackTrackFromPaths(t *testing.T) {
	tests := []struct {
		name        string
		attackTrack IAttackTrack
		controlsMap AttackTrackControlsLookup
		want        IAttackTrack
	}{
		{
			name: "Found attack chain",
			attackTrack: GetAttackTrackMock(
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
					"C": {
						control_5,
					},
					"B": {
						control_6,
					},
					"G": {
						control_7, control_8,
					},
				},
			},
			want: GetAttackTrackMock(
				AttackTrackStep{
					Name: "A",
					SubSteps: []AttackTrackStep{
						{
							Name: "C",
							SubSteps: []AttackTrackStep{
								{
									Name: "B",
								},
							},
						},
						{
							Name: "F",
						},
					},
				}),
		},
		{
			name: "No attack chain",
			attackTrack: GetAttackTrackMock(
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
					"B": {
						control_6,
					},
				},
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := NewAttackTrackAllPathsHandler(tt.attackTrack, &tt.controlsMap)
			paths := handler.CalculatePathsRootToLeaf()
			result := handler.GenerateAttackTrackFromPaths(paths)

			if !(tt.want == nil && result == nil) {
				if tt.want == nil {
					assert.Fail(t, "Expected is nil while actual is not nil")
				} else if result == nil {
					assert.Fail(t, "Actual is nil while expected is not nil")
				} else {
					assert.True(t, result.GetData().(*AttackTrackStep).Equal(tt.want.GetData().(*AttackTrackStep), false))
				}
			}

		})
	}
}

func TestFilterNodesWithControls(t *testing.T) {

	// Create some sample paths
	paths := [][]AttackTrackStepMock{
		{
			AttackTrackStepMock{Name: "A", Controls: []IAttackTrackControl{}},
			AttackTrackStepMock{Name: "B", Controls: []IAttackTrackControl{}},
		},
		{
			AttackTrackStepMock{Name: "A", Controls: []IAttackTrackControl{}},
			AttackTrackStepMock{Name: "C", Controls: []IAttackTrackControl{}},
			AttackTrackStepMock{Name: "D", Controls: []IAttackTrackControl{}},
		},
		{
			AttackTrackStepMock{Name: "B", Controls: []IAttackTrackControl{}},
			AttackTrackStepMock{Name: "E", Controls: []IAttackTrackControl{}},
		},
	}

	// Create the test cases
	testCases := []struct {
		name           string
		step           *AttackTrackStepMock
		expectedResult *AttackTrackStep
	}{
		{
			name: "No controls, no substeps",
			step: &AttackTrackStepMock{
				Name:        "A",
				Description: "Step A",
				SubSteps:    []AttackTrackStepMock{},
				Controls:    []IAttackTrackControl{},
			},
			expectedResult: nil,
		},
		{
			name: "Controls, no substeps",
			step: &AttackTrackStepMock{
				Name:        "B",
				Description: "Step B",
				SubSteps:    []AttackTrackStepMock{},
				Controls:    []IAttackTrackControl{control_1},
			},
			expectedResult: &AttackTrackStep{
				Name:        "B",
				Description: "Step B",
				SubSteps:    []AttackTrackStep{},
				Controls:    []IAttackTrackControl{control_1},
			},
		},
		{
			name: "No controls, substeps present in paths",
			step: &AttackTrackStepMock{
				Name:        "C",
				Description: "Step C",
				SubSteps: []AttackTrackStepMock{
					{Name: "D", Controls: []IAttackTrackControl{}},
				},
				Controls: []IAttackTrackControl{},
			},
			expectedResult: nil,
		},
		{
			name: "Controls, substeps not present in paths",
			step: &AttackTrackStepMock{
				Name:        "E",
				Description: "Step E",
				SubSteps: []AttackTrackStepMock{
					{Name: "F", Controls: []IAttackTrackControl{}},
					{Name: "G", Controls: []IAttackTrackControl{}},
				},
				Controls: []IAttackTrackControl{control_2},
			},
			expectedResult: &AttackTrackStep{
				Name:        "E",
				Description: "Step E",
				SubSteps:    []AttackTrackStep{},
				Controls:    []IAttackTrackControl{control_2},
			},
		},
		{
			name: "Controls, substeps present in paths",
			step: &AttackTrackStepMock{
				Name:        "A",
				Description: "Step A",
				SubSteps: []AttackTrackStepMock{
					{Name: "B", Controls: []IAttackTrackControl{}},
					{Name: "C", Controls: []IAttackTrackControl{}},
				},
				Controls: []IAttackTrackControl{control_1},
			},
			expectedResult: &AttackTrackStep{
				Name:        "A",
				Description: "Step A",
				SubSteps:    []AttackTrackStep{},
				Controls:    []IAttackTrackControl{control_1},
			},
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			handler := AttackTrackAllPathsHandler{
				// attackTrack: AttackTrackMock(tc.step),
				attackTrack: AttackTrackMock{Spec: MockAttackTrackSpecification{Data: tc.step}},
			}

			pathsCopy := make([][]IAttackTrackStep, len(paths))
			for i := range paths {
				pathsCopy[i] = make([]IAttackTrackStep, len(paths[i]))
				for j, step := range paths[i] {
					pathsCopy[i][j] = step
				}
			}

			result := handler.filterNodesWithControls(handler.attackTrack.GetData(), pathsCopy)

			if !(result == nil && tc.expectedResult == nil) && result.Equal(tc.expectedResult, true) == false {
				t.Errorf("Unexpected result.\nExpected: %+v\nGot: %+v", tc.expectedResult, result)
			}
		})
	}
}

func TestGetSubstepsWithVulnerabilities(t *testing.T) {
	// Create an AttackTrack object with substeps having different values for ChecksVulnerabilities
	attackTrack := AttackTrack{
		ApiVersion: "v1",
		Kind:       "AttackTrack",
		Metadata:   map[string]interface{}{},
		Spec: AttackTrackSpecification{
			Version:     "1.0",
			Description: "Example attack track",
			Data: AttackTrackStep{
				Name:                  "Step 1",
				Description:           "First step",
				ChecksVulnerabilities: true,
				SubSteps: []AttackTrackStep{
					{
						Name:                  "Substep 1.1",
						Description:           "Substep 1.1 description",
						ChecksVulnerabilities: true,
					},
					{
						Name:                  "Substep 1.2",
						Description:           "Substep 1.2 description",
						ChecksVulnerabilities: false,
					},
				},
			},
		},
	}

	// Call the method being tested
	substepNames := attackTrack.GetSubstepsWithVulnerabilities()

	// Define the expected substep names with ChecksVulnerabilities set to true
	expectedSubstepNames := []string{"Step 1", "Substep 1.1"}

	// Check if the returned substep names match the expected substep names
	if len(substepNames) != len(expectedSubstepNames) {
		t.Errorf("Unexpected number of substep names. Expected: %d, Got: %d", len(expectedSubstepNames), len(substepNames))
	}

	for i, name := range substepNames {
		if name != expectedSubstepNames[i] {
			t.Errorf("Mismatched substep name. Expected: %s, Got: %s", expectedSubstepNames[i], name)
		}
	}
}
