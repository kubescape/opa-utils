package gitregostore

import (
	"fmt"
	"strings"

	// "github.com/armosec/capacketsgo/opapolicy"
	"github.com/armosec/armoapi-go/armotypes"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	opapolicy "github.com/kubescape/opa-utils/reporthandling"
	"github.com/kubescape/opa-utils/reporthandling/attacktrack/v1alpha1"
)

// GetOPAPolicies returns all the policies of given customer
func (gs *GitRegoStore) GetOPAPolicies() ([]opapolicy.PolicyRule, error) {
	if gs.Rules == nil {
		return nil, fmt.Errorf("no rules found in GitRegoStore")
	}
	return gs.Rules, nil
}

func (gs *GitRegoStore) GetOPAPoliciesNamesList() ([]string, error) {
	gs.rulesLock.RLock()
	defer gs.rulesLock.RUnlock()
	var policiesNameList []string
	for _, rule := range gs.Rules {
		policiesNameList = append(policiesNameList, rule.Name)
	}
	return policiesNameList, nil
}

// GetOPAPolicyByName returns specific policy by the name
func (gs *GitRegoStore) GetOPAPolicyByName(ruleName string) (*opapolicy.PolicyRule, error) {
	gs.rulesLock.RLock()
	defer gs.rulesLock.RUnlock()
	for _, rule := range gs.Rules {
		if strings.EqualFold(rule.Name, ruleName) {
			return &rule, nil
		}
	}
	return nil, fmt.Errorf("rule '%s' not found", ruleName)
}

func (gs *GitRegoStore) fillRulesAndRulesIDsInControl(control *opapolicy.Control) error {
	fil := gs.ControlRuleRelations.Filter(
		dataframe.F{Colname: "ControlID", Comparator: series.Eq, Comparando: control.ControlID},
	)
	var rulesList []opapolicy.PolicyRule
	var rulesIDList []string

	for row := 0; row < fil.Nrow(); row++ {
		ruleName := fil.Elem(row, 1)
		rule, err := gs.GetOPAPolicyByName(ruleName.String())
		if err != nil {
			return err
		}
		// add rule to control.rules
		rulesList = append(rulesList, *rule)
		// add ruleId ro control.ruleIds
		rulesIDList = append(rulesIDList, rule.GUID)
	}
	control.Rules = rulesList
	control.RulesIDs = &rulesIDList
	return nil
}

func (gs *GitRegoStore) GetAttackTracks() ([]v1alpha1.AttackTrack, error) {
	gs.attackTracksLock.RLock()
	defer gs.attackTracksLock.RUnlock()

	if gs.AttackTracks == nil {
		return nil, fmt.Errorf("no attack tracks found in GitRegoStore")
	}
	return gs.AttackTracks, nil
}

// GetOPAControlByName returns specific control by the name
func (gs *GitRegoStore) GetOPAControlByName(controlName string) (*opapolicy.Control, error) {
	gs.controlsLock.RLock()
	defer gs.controlsLock.RUnlock()
	for _, control := range gs.Controls {
		if strings.EqualFold(control.Name, controlName) || strings.EqualFold(control.Name, getNewControlName(control.ControlID, controlName)) {
			err := gs.fillRulesAndRulesIDsInControl(&control)
			if err != nil {
				return nil, err
			}
			return &control, nil
		}
	}
	return nil, fmt.Errorf("control '%s' not found", controlName)
}

// GetOPAControlByID returns specific control by the ID
func (gs *GitRegoStore) GetOPAControlByID(controlID string) (*opapolicy.Control, error) {
	gs.controlsLock.RLock()
	defer gs.controlsLock.RUnlock()
	for _, control := range gs.Controls {
		if strings.EqualFold(control.ControlID, controlID) || strings.EqualFold(control.ControlID, getNewControlID(controlID)) {
			err := gs.fillRulesAndRulesIDsInControl(&control)
			if err != nil {
				return nil, err
			}
			return &control, nil
		}
	}
	return nil, fmt.Errorf("control '%s' not found", controlID)
}

// GetOPAControl returns specific control by the name or ID
func (gs *GitRegoStore) GetOPAControl(c string) (*opapolicy.Control, error) {
	if isControlID(c) {
		return gs.GetOPAControlByID(c)
	} else {
		return gs.GetOPAControlByName(c)
	}
}

// GetOPAControls returns all the controls of given customer
func (gs *GitRegoStore) GetOPAControls() ([]opapolicy.Control, error) {
	gs.controlsLock.RLock()
	defer gs.controlsLock.RUnlock()
	var controlsList []opapolicy.Control
	if gs.Controls == nil {
		return nil, fmt.Errorf("no controls found in GitRegoStore")
	}
	for _, control := range gs.Controls {
		err := gs.fillRulesAndRulesIDsInControl(&control)
		if err != nil {
			return nil, err
		}
		controlsList = append(controlsList, control)
	}
	return controlsList, nil
}

func (gs *GitRegoStore) GetOPAControlsNamesList() ([]string, error) {
	gs.controlsLock.RLock()
	defer gs.controlsLock.RUnlock()
	var controlsNameList []string
	for _, control := range gs.Controls {
		controlsNameList = append(controlsNameList, control.Name)
	}
	return controlsNameList, nil
}

func (gs *GitRegoStore) GetOPAControlsIDsList() ([]string, error) {
	gs.controlsLock.RLock()
	defer gs.controlsLock.RUnlock()
	var controlsIDList []string
	for _, control := range gs.Controls {
		controlsIDList = append(controlsIDList, control.ControlID)
	}
	return controlsIDList, nil
}

func (gs *GitRegoStore) fillControlsAndControlIDsInFramework(fw *opapolicy.Framework) error {
	fil := gs.FrameworkControlRelations.Filter(
		dataframe.F{Colname: "frameworkName", Comparator: series.Eq, Comparando: fw.Name},
	)
	var controlsList []opapolicy.Control
	var controlsIDList []string

	for row := 0; row < fil.Nrow(); row++ {
		controlName := fil.Elem(row, 2)
		control, err := gs.GetOPAControlByName(controlName.String())
		if err != nil {
			return err
		}
		// add rule to control.rules
		controlsList = append(controlsList, *control)
		// add ruleId ro control.ruleIds
		controlsIDList = append(controlsIDList, control.ControlID)
	}
	fw.Controls = controlsList
	fw.ControlsIDs = &controlsIDList
	return nil
}

// GetOpaFrameworkListByControlName return a list of fw names this control is in
func (gs *GitRegoStore) GetOpaFrameworkListByControlName(controlName string) []string {
	var frameworksNameList []string
	fil := gs.FrameworkControlRelations.Filter(
		dataframe.F{Colname: "ControlName", Comparator: series.Eq, Comparando: controlName},
	)
	for row := 0; row < fil.Nrow(); row++ {
		fwName := fil.Elem(row, 0)
		frameworksNameList = append(frameworksNameList, fwName.String())
	}
	return frameworksNameList
}

// GetOpaFrameworkListByControlID return a list of fw names this control is in
func (gs *GitRegoStore) GetOpaFrameworkListByControlID(controlID string) []string {
	var frameworksNameList []string
	fil := gs.FrameworkControlRelations.Filter(
		dataframe.F{Colname: "ControlID", Comparator: series.Eq, Comparando: controlID},
	)
	for row := 0; row < fil.Nrow(); row++ {
		fwName := fil.Elem(row, 0)
		frameworksNameList = append(frameworksNameList, fwName.String())
	}
	return frameworksNameList
}

// GetOPAFrameworks returns all the frameworks of given customer
func (gs *GitRegoStore) GetOPAFrameworks() ([]opapolicy.Framework, error) {
	gs.frameworksLock.RLock()
	defer gs.frameworksLock.RUnlock()
	var frameworksList []opapolicy.Framework
	if gs.Frameworks == nil {
		return nil, fmt.Errorf("no frameworks found in GitRegoStore")
	}
	for _, fw := range gs.Frameworks {
		err := gs.fillControlsAndControlIDsInFramework(&fw)
		if err != nil {
			return nil, err
		}
		frameworksList = append(frameworksList, fw)
	}
	return frameworksList, nil
}

func (gs *GitRegoStore) GetOPAFrameworksNamesList() ([]string, error) {
	gs.frameworksLock.RLock()
	defer gs.frameworksLock.RUnlock()
	var frameworksNameList []string
	for _, framework := range gs.Frameworks {
		frameworksNameList = append(frameworksNameList, framework.Name)
	}
	return frameworksNameList, nil
}

// GetOPAFrameworkByName returns specific framework by the name
func (gs *GitRegoStore) GetOPAFrameworkByName(frameworkName string) (*opapolicy.Framework, error) {
	gs.frameworksLock.RLock()
	defer gs.frameworksLock.RUnlock()
	for _, fw := range gs.Frameworks {
		if strings.EqualFold(fw.Name, frameworkName) {
			err := gs.fillControlsAndControlIDsInFramework(&fw)
			if err != nil {
				return nil, err
			}
			return &fw, nil
		}
	}
	return nil, fmt.Errorf("framework '%s' not found", frameworkName)
}

func (gs *GitRegoStore) GetDefaultConfigInputs() (armotypes.CustomerConfig, error) {
	gs.DefaultConfigInputsLock.RLock()
	defer gs.DefaultConfigInputsLock.RUnlock()
	return gs.DefaultConfigInputs, nil
}

func (gs *GitRegoStore) GetSystemPostureExceptionPolicies() ([]armotypes.PostureExceptionPolicy, error) {
	gs.systemPostureExceptionPoliciesLock.RLock()
	defer gs.systemPostureExceptionPoliciesLock.RUnlock()
	return gs.SystemPostureExceptionPolicies, nil
}
