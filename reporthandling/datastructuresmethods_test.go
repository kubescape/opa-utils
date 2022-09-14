package reporthandling

import (
	"encoding/json"
	"testing"

	"github.com/kubescape/opa-utils/reporthandling/mock"
	"github.com/stretchr/testify/assert"
)

func FrameworkResultsMock(report string) (*FrameworkReport, error) {
	frameworkReport := &FrameworkReport{}
	if err := json.Unmarshal([]byte(report), frameworkReport); err != nil {
		return nil, err
	}
	return frameworkReport, nil
}

func ControlsResultsMock(report string) ([]ControlReport, error) {
	f, err := FrameworkResultsMock(report)
	if err != nil {
		return nil, err
	}
	return f.ControlReports, nil
}

func TestControlsResults(t *testing.T) {

	framework, err := FrameworkResultsMock(mock.NSAScanV10119)
	assert.NoError(t, err, err)
	assert.Equal(t, len(framework.ControlReports), 21)

	SetUniqueResourcesCounter(framework)

	assert.Equal(t, 28, framework.GetNumberOfFailedResources(), "framework.GetNumberOfFailedResources")
	assert.Equal(t, 31, framework.GetNumberOfWarningResources(), "framework.GetNumberOfWarningResources")

	for _, control := range framework.ControlReports {
		switch control.ControlID {
		case "C-0005":
			assert.Equal(t, 0, control.GetNumberOfFailedResources(), "C-0005: control.GetNumberOfFailedResources")
			assert.Equal(t, 0, control.GetNumberOfWarningResources(), "C-0005: GetNumberOfWarningResources")
			assert.True(t, control.Passed(), "C-0005: Passed")
			assert.False(t, control.Warning(), "C-0005: Warning")
			assert.False(t, control.Failed(), "C-0005: Failed")
		case "C-0038":
			assert.Equal(t, 0, control.GetNumberOfFailedResources(), "C-0038: GetNumberOfFailedResources")
			assert.Equal(t, 0, control.GetNumberOfWarningResources(), "C-0038: GetNumberOfWarningResources")
			assert.True(t, control.Passed(), "C-0038: Passed")
			assert.False(t, control.Warning(), "C-0038: Warning")
			assert.False(t, control.Failed(), "C-0038: Failed")
		case "C-0017":
			assert.Equal(t, 0, control.GetNumberOfFailedResources(), "C-0017: GetNumberOfFailedResources")
			assert.Equal(t, 6, control.GetNumberOfWarningResources(), "C-0017: GetNumberOfWarningResources")
			assert.False(t, control.Passed(), "C-0017: Passed")
			assert.True(t, control.Warning(), "C-0017: Warning")
			assert.False(t, control.Failed(), "C-0017: Failed")
			for _, rule := range control.RuleReports {
				assert.Equal(t, 0, rule.GetNumberOfFailedResources(), "C-0017: rule.GetNumberOfFailedResources")
				assert.Equal(t, 6, rule.GetNumberOfWarningResources(), "C-0017: rule.GetNumberOfWarningResources")
			}
		case "C-0009":
			assert.Equal(t, 0, control.GetNumberOfFailedResources(), "C-0009: GetNumberOfFailedResources")
			assert.Equal(t, 6, control.GetNumberOfWarningResources(), "C-0009: GetNumberOfWarningResources")
			assert.False(t, control.Passed(), "C-0009: Passed")
			assert.True(t, control.Warning(), "C-0009: Warning")
			assert.False(t, control.Failed(), "C-0009: Failed")
		case "C-0030":
			assert.Equal(t, 7, control.GetNumberOfFailedResources(), "C-0030: GetNumberOfFailedResources")
			assert.Equal(t, 0, control.GetNumberOfWarningResources(), "C-0030: GetNumberOfWarningResources")
			assert.False(t, control.Passed(), "C-0030: Passed")
			assert.False(t, control.Warning(), "C-0030: Warning")
			assert.True(t, control.Failed(), "C-0030: Failed")
		case "C-0013":
			assert.Equal(t, 0, control.GetNumberOfFailedResources(), "C-0013: GetNumberOfFailedResources")
			assert.Equal(t, 0, control.GetNumberOfWarningResources(), "C-0013: GetNumberOfWarningResources")
			assert.True(t, control.Passed(), "C-0013: Passed")
			assert.False(t, control.Warning(), "C-0013: Warning")
			assert.False(t, control.Failed(), "C-0013: Failed")
		case "C-0034":
			assert.Equal(t, 5, control.GetNumberOfFailedResources(), "C-0034: GetNumberOfFailedResources")
			assert.Equal(t, 31, control.GetNumberOfWarningResources(), "C-0034: GetNumberOfWarningResources")
			assert.False(t, control.Passed(), "C-0034: Passed")
			assert.False(t, control.Warning(), "C-0034: Warning")
			assert.True(t, control.Failed(), "C-0034: Failed")
		case "C-0035":
			assert.Equal(t, 4, control.GetNumberOfFailedResources(), "C-0035: GetNumberOfFailedResources")
			assert.Equal(t, 0, control.GetNumberOfWarningResources(), "C-0035: GetNumberOfWarningResources")
			assert.False(t, control.Passed(), "C-0035: Passed")
			assert.False(t, control.Warning(), "C-0035: Warning")
			assert.True(t, control.Failed(), "C-0035: Failed")
		case "C-0016":
			assert.Equal(t, 0, control.GetNumberOfFailedResources(), "C-0016: GetNumberOfFailedResources")
			assert.Equal(t, 0, control.GetNumberOfWarningResources(), "C-0016: GetNumberOfWarningResources")
			assert.True(t, control.Passed(), "C-0016: Passed")
			assert.False(t, control.Warning(), "C-0016: Warning")
			assert.False(t, control.Failed(), "C-0016: Failed")
		}
	}
}

func TestListResourcesIDs(t *testing.T) {
	framework, err := FrameworkResultsMock(mock.NSAScanV10119)
	assert.NoError(t, err, err)
	assert.Equal(t, len(framework.ControlReports), 21)

	for _, controlReport := range framework.ControlReports {
		for _, ruleReport := range controlReport.RuleReports {
			if ruleReport.Name == "immutable-container-filesystem" {
				assert.Equal(t, 6, len(ruleReport.ListResourcesIDs().GetWarningResources()))
				return
			}
		}
	}
}

func TestControl_GetAttackTrackCategories(t *testing.T) {
	validControlJson := `{"name":"TEST","attributes":{"armoBuiltin":true,"controlTypeTags":["security","compliance"],"attackTracks":[{"attackTrack": "container","categories": ["Execution","Initial access"]}]},"description":"","remediation":"","rulesNames":["CVE-2022-0185"],"id":"C-0079","long_description":"","test":"","controlID":"C-0079","baseScore":4,"example":""}`
	var validControl Control
	err := json.Unmarshal([]byte(validControlJson), &validControl)
	assert.NoError(t, err, err)
	assert.Equal(t, []string{"Execution", "Initial access"}, validControl.GetAttackTrackCategories("container"))
	assert.Equal(t, []string{}, validControl.GetAttackTrackCategories("test"))

	invalidControlJson1 := `{"name":"TEST","attributes":{"armoBuiltin":true,"controlTypeTags":["security","compliance"],"attackTracks":{"container": "x"}},"description":"","remediation":"","rulesNames":["CVE-2022-0185"],"id":"C-0079","long_description":"","test":"","controlID":"C-0079","baseScore":4,"example":""}`
	var invalidControl1 Control
	err = json.Unmarshal([]byte(invalidControlJson1), &invalidControl1)
	assert.NoError(t, err, err)
	assert.Equal(t, []string{}, invalidControl1.GetAttackTrackCategories("container"))

	invalidControlJson2 := `{"name":"TEST","attributes":{"armoBuiltin":true,"controlTypeTags":["security","compliance"],"attack":{"container": "x"}},"description":"","remediation":"","rulesNames":["CVE-2022-0185"],"id":"C-0079","long_description":"","test":"","controlID":"C-0079","baseScore":4,"example":""}`
	var invalidControl2 Control
	err = json.Unmarshal([]byte(invalidControlJson2), &invalidControl2)
	assert.NoError(t, err, err)
	assert.Equal(t, []string{}, invalidControl2.GetAttackTrackCategories("container"))
}

func TestControl_GetControlTypeTags(t *testing.T) {
	validControlJson := `{"name":"TEST","attributes":{"armoBuiltin":true,"controlTypeTags":["security","compliance"],"attackTracks":{"container":["Privilege escalation"]}},"description":"","remediation":"","rulesNames":["CVE-2022-0185"],"id":"C-0079","long_description":"","test":"","controlID":"C-0079","baseScore":4,"example":""}`
	var validControl Control
	err := json.Unmarshal([]byte(validControlJson), &validControl)
	assert.NoError(t, err, err)
	assert.Equal(t, []string{"security", "compliance"}, validControl.GetControlTypeTags())

	missingAttributeControlJson := `{"name":"TEST","attributes":{"armoBuiltin":true,"attackTracks":{"container": "x"}},"description":"","remediation":"","rulesNames":["CVE-2022-0185"],"id":"C-0079","long_description":"","test":"","controlID":"C-0079","baseScore":4,"example":""}`
	var missingAttributeControl Control
	err = json.Unmarshal([]byte(missingAttributeControlJson), &missingAttributeControl)
	assert.NoError(t, err, err)
	assert.Equal(t, []string{}, missingAttributeControl.GetControlTypeTags())
}
