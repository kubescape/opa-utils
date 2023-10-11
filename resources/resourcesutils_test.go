package resources

import (
	"strings"
	"testing"

	"github.com/kubescape/opa-utils/reporthandling"
	"github.com/stretchr/testify/assert"
)

func TestGetFilteredPostureControlInputs(t *testing.T) {
	regoInputData := RegoDependenciesData{}
	regoInputData.PostureControlInputs = map[string][]string{"sensitiveKeyNames": {"keyA", "keyB"}}
	s := []string{"settings.postureControlInputs.sensitiveKeyNames", "settings.postureControlInputs.blabla"}
	postureControlInputs := regoInputData.GetFilteredPostureControlInputs(s)
	splitted0 := strings.Split(s[0], ".")
	_, ok := postureControlInputs[splitted0[2]]
	assert.True(t, ok)

	splitted1 := strings.Split(s[1], ".")
	_, ok = postureControlInputs[splitted1[2]]
	assert.False(t, ok)
}

func TestGetFilteredPostureControlConfigInputs(t *testing.T) {
	regoInputData := RegoDependenciesData{}
	regoInputData.PostureControlInputs = map[string][]string{"sensitiveKeyNames": {"keyA", "keyB"}}

	inputs := []reporthandling.ControlConfigInputs{
		{
			Path: "settings.postureControlInputs.sensitiveKeyNames",
			Name: "Sensitive Key Names",
		},
		{
			Path: "settings.postureControlInputs.blabla",
			Name: "Blabla",
		},
	}

	postureControlInputs := regoInputData.GetFilteredPostureControlConfigInputs(inputs)

	splitted0 := strings.Split(inputs[0].Path, ".")
	_, ok := postureControlInputs[splitted0[2]]
	assert.True(t, ok)

	splitted1 := strings.Split(inputs[1].Path, ".")
	_, ok = postureControlInputs[splitted1[2]]
	assert.False(t, ok)
}
