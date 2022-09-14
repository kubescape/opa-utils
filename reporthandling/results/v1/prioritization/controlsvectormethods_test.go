package prioritization

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewControlsVector(t *testing.T) {
	tests := []struct {
		name string
		want *ControlsVector
	}{
		{
			name: "controls vector initialization",
			want: &ControlsVector{
				AttackTrackName: "test",
				Score:           0,
				Type:            ControlPriorityVectorType,
				Vector:          []*PriorityVectorControl{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewControlsVector("test"); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewControlsVector() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetScore(t *testing.T) {
	cv := &ControlsVector{
		Type:   ControlPriorityVectorType,
		Vector: []*PriorityVectorControl{},
	}
	assert.Equalf(t, float64(0), cv.GetScore(), "ControlsVector.GetScore() = %v, expected %v", cv.GetScore(), float64(0))

	cv1 := &ControlsVector{
		Type:   ControlPriorityVectorType,
		Vector: []*PriorityVectorControl{},
		Score:  5,
	}
	assert.Equalf(t, float64(5), cv1.GetScore(), "ControlsVector.GetScore() = %v, expected %v", cv1.GetScore(), float64(5))

}

func TestSetScore(t *testing.T) {
	cv := &ControlsVector{
		Type: ControlPriorityVectorType,
		Vector: []*PriorityVectorControl{
			{ControlID: "C1", Category: "X"},
			{ControlID: "C2", Category: "Y"},
		},
		Score:    3,
		Severity: 1,
	}
	var expected float64 = 3
	assert.Equalf(t, expected, cv.Score, "ControlsVector.Score = %v, expected %v", cv.Score, expected)
	cv.SetScore(5)

	expected = 5
	assert.Equalf(t, expected, cv.Score, "ControlsVector.Score = %v, expected %v", cv.Score, expected)
}

func TestAdd(t *testing.T) {
	cv := &ControlsVector{
		Type: ControlPriorityVectorType,
		Vector: []*PriorityVectorControl{
			{ControlID: "C1", Category: "X"},
			{ControlID: "C2", Category: "Y"},
		},
		Score:    3,
		Severity: 4,
	}
	expected := []*PriorityVectorControl{
		{ControlID: "C1", Category: "X"},
		{ControlID: "C2", Category: "Y"},
	}
	assert.ElementsMatchf(t, expected, cv.Vector, "ControlsVector.Vector = %v, expected %v", cv.Vector, expected)

	err := cv.Add(PriorityVectorControl{ControlID: "C3"})
	assert.NoError(t, err)
	expected = []*PriorityVectorControl{
		{ControlID: "C1", Category: "X"},
		{ControlID: "C2", Category: "Y"},
		{ControlID: "C3"},
	}

	assert.ElementsMatchf(t, expected, cv.Vector, "ControlsVector.Vector = %v, expected %v", cv.Vector, expected)

	err = cv.Add("C3")
	assert.Errorf(t, err, "expected error when trying to add unsupported type")
}

func TestList(t *testing.T) {
	cv := &ControlsVector{
		Type: ControlPriorityVectorType,
		Vector: []*PriorityVectorControl{
			{ControlID: "C1", Category: "X"},
			{ControlID: "C2", Category: "Y"},
		},
		Score: 3,
	}

	result := cv.List()
	expected := []*PriorityVectorControl{
		{ControlID: "C1", Category: "X"},
		{ControlID: "C2", Category: "Y"},
	}

	assert.ElementsMatchf(t, expected, result, "ControlsVector.List() = %v, expected %v", result, expected)
}

func TestGetType(t *testing.T) {
	cv := &ControlsVector{
		Type: ControlPriorityVectorType,
		Vector: []*PriorityVectorControl{
			{ControlID: "C1", Category: "X"},
			{ControlID: "C2", Category: "Y"},
		},
		Score: 3,
	}
	result := cv.GetType()
	expected := ControlPriorityVectorType

	assert.Equalf(t, expected, result, "ControlsVector.GetType() = %v, expected %v", result, expected)
}

func TestGetSeverity(t *testing.T) {
	cv := &ControlsVector{
		Type: ControlPriorityVectorType,
		Vector: []*PriorityVectorControl{
			{ControlID: "C1", Category: "X"},
			{ControlID: "C2", Category: "Y"},
		},
		Score:    3.2,
		Severity: 4,
	}
	result := cv.GetSeverity()
	expected := 4
	assert.Equalf(t, expected, result, "ControlsVector.GetSeverity() = %v, expected %v", result, expected)
}

func TestSetSeverity(t *testing.T) {
	cv := &ControlsVector{
		Type: ControlPriorityVectorType,
		Vector: []*PriorityVectorControl{
			{ControlID: "C1", Category: "X"},
			{ControlID: "C2", Category: "Y"},
		},
		Score:    3.2,
		Severity: 4,
	}
	cv.SetSeverity(5)

	result := cv.GetSeverity()
	expected := 5
	assert.Equalf(t, expected, result, "ControlsVector.GetSeverity() = %v, expected %v", result, expected)
}
