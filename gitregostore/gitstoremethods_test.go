package gitregostore

import (
	"fmt"
	"testing"
)

func TestInitGitRegoStoreFromRelease(t *testing.T) {
	baseUrl := "https://api.github.com/repos"
	owner := "armosec"
	repository := "regolibrary"
	path := "releases"
	tag := "51267909"
	branch := "master"
	frequency := 15
	gs := InitGitRegoStore(baseUrl, owner, repository, path, tag, branch, frequency)
	if gs.Rules == nil {
		t.Errorf("failed to decode")
	}
	fmt.Println(gs.URL)
}

func TestInitGitRegoStoreFromRepo(t *testing.T) {
	baseUrl := "https://api.github.com/repos"
	owner := "armosec"
	repository := "regolibrary"
	path := "git/trees"
	tag := ""
	branch := "dev"
	frequency := 15
	gs := InitGitRegoStore(baseUrl, owner, repository, path, tag, branch, frequency)
	if gs.Controls == nil {
		t.Errorf("failed to decode controls")
	}
	if gs.Frameworks == nil {
		t.Errorf("failed to decode frameworks")
	}
	if gs.Rules == nil {
		t.Errorf("failed to decode rules")
	}
	fmt.Println(gs.URL)
}

func TestGetPoliciesMethods(t *testing.T) {
	baseUrl := "https://api.github.com/repos"
	owner := "armosec"
	repository := "regolibrary"
	path := "releases"
	tag := "51267909"
	branch := "master"
	frequency := 15
	gs := InitGitRegoStore(baseUrl, owner, repository, path, tag, branch, frequency)
	// Rules
	policies, err := gs.GetOPAPolicies()
	if err != nil || policies == nil {
		t.Errorf("failed to get all policies %s", err.Error())
	}
	policiesNames, err := gs.GetOPAPoliciesNamesList()
	if err != nil || len(policiesNames) == 0 {
		t.Errorf("failed to get policies names list %s", err.Error())
	}
	policy, err := gs.GetOPAPolicy(policies[0].GUID)
	if err != nil || policy == nil {
		t.Errorf("failed to get policy by guid: '%s', %s", policies[0].GUID, err.Error())
	}
	policy, err = gs.GetOPAPolicyByName(policiesNames[0])
	if err != nil || policy == nil {
		t.Errorf("failed to get policy by name: '%s', %s", policiesNames[0], err.Error())
	}
	// Controls
	controls, err := gs.GetOPAControls()
	if err != nil || controls == nil {
		t.Errorf("failed to get all controls %s", err.Error())
	}
	controlsNames, err := gs.GetOPAControlsNamesList()
	if err != nil || len(controlsNames) == 0 {
		t.Errorf("failed to get controls names list %s", err.Error())
	}
	// control, err := gs.GetOPAControl(controls[0].GUID)
	// if err != nil || control == nil {
	// 	t.Errorf("failed to get control by guid: '%s', %s", controls[0].GUID, err.Error())
	// }
	control, err := gs.GetOPAControlByName(controlsNames[0])
	if err != nil || control == nil {
		t.Errorf("failed to get control by name: '%s', %s", controlsNames[0], err.Error())
	}
	// Frameworks
	frameworks, err := gs.GetOPAFrameworks()
	if err != nil || frameworks == nil {
		t.Errorf("failed to get all frameworks %s", err.Error())
	}
	frameworksNames, err := gs.GetOPAFrameworksNamesList()
	if err != nil || len(frameworksNames) == 0 {
		t.Errorf("failed to get frameworks names list %s", err.Error())
	}
	// framework, err := gs.GetOPAFramework(frameworks[0].GUID)
	// if err != nil || framework == nil {
	// 	t.Errorf("failed to get framework by guid: '%s', %s", frameworks[0].GUID, err.Error())
	// }
	framework, err := gs.GetOPAFrameworkByName(frameworksNames[0])
	if err != nil || framework == nil {
		t.Errorf("failed to get framework by name: '%s', %s", frameworksNames[0], err.Error())
	}
}
