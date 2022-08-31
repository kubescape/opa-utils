package prioritization

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewControlPriorityVector(t *testing.T) {
	tests := []struct {
		name string
		want *PriorityVector
	}{
		{
			name: "control priority vector initialization",
			want: &PriorityVector{
				Score:  0,
				Type:   ControlPriorityVectorType,
				Vector: []string{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewControlPriorityVector(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewControlPriorityVector() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetScore(t *testing.T) {
	pv := &PriorityVector{
		Type:   ControlPriorityVectorType,
		Vector: []string{},
	}
	assert.Equalf(t, float64(0), pv.GetScore(), "PriorityVector.GetScore() = %v, expected %v", pv.GetScore(), float64(0))

	pv1 := &PriorityVector{
		Type:   ControlPriorityVectorType,
		Vector: []string{},
		Score:  5,
	}
	assert.Equalf(t, float64(5), pv1.GetScore(), "PriorityVector.GetScore() = %v, expected %v", pv1.GetScore(), float64(5))

}

func TestSetScore(t *testing.T) {
	pv := &PriorityVector{
		Type:   ControlPriorityVectorType,
		Vector: []string{"C1", "C2"},
		Score:  3,
	}
	var expected float64 = 3
	assert.Equalf(t, expected, pv.Score, "PriorityVector.Score = %v, expected %v", pv.Score, expected)
	pv.SetScore(5)

	expected = 5
	assert.Equalf(t, expected, pv.Score, "PriorityVector.Score = %v, expected %v", pv.Score, expected)
}

func TestAdd(t *testing.T) {
	pv := &PriorityVector{
		Type:   ControlPriorityVectorType,
		Vector: []string{"C1", "C2"},
		Score:  3,
	}
	expected := []string{"C1", "C2"}
	assert.ElementsMatchf(t, expected, pv.Vector, "PriorityVector.Vector = %v, expected %v", pv.Vector, expected)

	pv.Add("C3")
	expected = []string{"C1", "C2", "C3"}

	assert.ElementsMatchf(t, expected, pv.Vector, "PriorityVector.Vector = %v, expected %v", pv.Vector, expected)
}

func TestList(t *testing.T) {
	pv := &PriorityVector{
		Type:   ControlPriorityVectorType,
		Vector: []string{"C1", "C2"},
		Score:  3,
	}

	result := pv.List()
	expected := []string{"C1", "C2"}

	assert.ElementsMatchf(t, expected, result, "PriorityVector.List() = %v, expected %v", result, expected)
}

func TestGetType(t *testing.T) {
	pv := &PriorityVector{
		Type:   ControlPriorityVectorType,
		Vector: []string{"C1", "C2"},
		Score:  3,
	}
	result := pv.GetType()
	expected := ControlPriorityVectorType

	assert.Equalf(t, expected, result, "PriorityVector.GetType() = %v, expected %v", result, expected)
}
