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

const (
	supportBackwardCompatibility = true
)

// GetOPAPolicies returns all the policies of given customer
func (gs *GitRegoStore) GetOPAPolicies() ([]opapolicy.PolicyRule, error) {
	gs.rulesLock.RLock()
	defer gs.rulesLock.RUnlock()

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

	return gs.getOPAPolicyByName(ruleName)
}

func (gs *GitRegoStore) getOPAPolicyByName(ruleName string) (*opapolicy.PolicyRule, error) {
	for _, rule := range gs.Rules {
		if strings.EqualFold(rule.Name, ruleName) {
			return &rule, nil
		}
	}

	return nil, fmt.Errorf("rule '%s' not found", ruleName)
}

func (gs *GitRegoStore) GetAttackTracks() ([]v1alpha1.AttackTrack, error) {
	gs.attackTracksLock.RLock()
	defer gs.attackTracksLock.RUnlock()

	if gs.AttackTracks == nil {
		return nil, fmt.Errorf("no attack tracks found in GitRegoStore")
	}
	return gs.AttackTracks, nil
}

// DEPECATED
// GetOPAControlByName returns specific BaseControl by the name
func (gs *GitRegoStore) GetOPAControlByName(controlName string) (*opapolicy.Control, error) {
	gs.controlsLock.RLock()
	defer gs.controlsLock.RUnlock()

	for _, controlToPin := range gs.Controls {
		control := controlToPin
		if strings.EqualFold(control.Name, controlName) ||
			// If backward compatibility is supported, extract from patched control name the new name.
			(supportBackwardCompatibility && strings.EqualFold(control.Name, baseControlName(control.ControlID, controlName))) {
			if len(control.Rules) == 0 {
				err := gs.fillRulesAndRulesIDsInControl(&control)
				if err != nil {
					return nil, err
				}
			}
			return &control, nil
		}
	}
	return nil, fmt.Errorf("control '%s' not found", controlName)
}

// GetOPAControlByID returns specific BaseControl by the ID
func (gs *GitRegoStore) GetOPAControlByID(controlID string) (*opapolicy.Control, error) {
	gs.controlsLock.RLock()
	defer gs.controlsLock.RUnlock()

	return gs.getOPAControlByID(controlID)
}

func (gs *GitRegoStore) getOPAControlByID(controlID string) (*opapolicy.Control, error) {
	for _, controlToPin := range gs.Controls {
		control := controlToPin
		if strings.EqualFold(control.ControlID, controlID) ||
			// If backward compatibility is supported,try to find if the controlID sent has a new controlID
			(supportBackwardCompatibility && strings.EqualFold(control.ControlID, newControlID(controlID))) {

			if len(control.Rules) == 0 {
				err := gs.fillRulesAndRulesIDsInControl(&control)
				if err != nil {
					return nil, err
				}
			}

			return &control, nil
		}
	}
	return nil, fmt.Errorf("control '%s' not found", controlID)
}

// GetOPAControlByFrameworkNameAndControlName - get framework name and control name and return the relevant control object
func (gs *GitRegoStore) GetOPAControlByFrameworkNameAndControlName(frameworkName string, controlName string) (*opapolicy.Control, error) {
	gs.frameworksLock.RLock()
	defer gs.frameworksLock.RUnlock()

	gs.controlsLock.RLock()
	defer gs.controlsLock.RUnlock()

	fw, err := gs.getOPAFrameworkByName(frameworkName) // locks framework
	if err != nil {
		return nil, err
	}

	for _, controlToPin := range fw.Controls {
		control := controlToPin
		if strings.EqualFold(control.Name, controlName) ||
			// If backward compatibility is supported, extract from patched control name the new name.
			(supportBackwardCompatibility && strings.EqualFold(control.Name, baseControlName(control.ControlID, controlName))) {
			if len(control.Rules) == 0 {
				err := gs.fillRulesAndRulesIDsInControl(&control)
				if err != nil {
					return nil, err
				}
			}
			return &control, nil
		}
	}

	return nil, fmt.Errorf("control  name '%s' not found in framework '%s'", controlName, fw.Name)

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
	for _, controlToPin := range gs.Controls {
		control := controlToPin
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

// GetOpaFrameworkListByControlName return a list of fw names this control is in
func (gs *GitRegoStore) GetOpaFrameworkListByControlName(controlName string) []string {
	gs.frameworkRelationsLock.RLock()
	defer gs.frameworkRelationsLock.RUnlock()

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
	gs.frameworkRelationsLock.RLock()
	defer gs.frameworkRelationsLock.RUnlock()

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
	for _, frameworkToPin := range gs.Frameworks {
		fw := frameworkToPin
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

	return gs.getOPAFrameworkByName(frameworkName)
}

func (gs *GitRegoStore) getOPAFrameworkByName(frameworkName string) (*opapolicy.Framework, error) {
	for _, frameworkToPin := range gs.Frameworks {
		fw := frameworkToPin
		if strings.EqualFold(fw.Name, frameworkName) ||
			// If backward compatibility is supported,try to compare the new CIS name.
			(true && strings.EqualFold(fw.Name, newFrameworkName(frameworkName))) {
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

func (gs *GitRegoStore) fillRulesAndRulesIDsInControl(control *opapolicy.Control) error {
	gs.rulesLock.RLock()
	defer gs.rulesLock.RUnlock()

	gs.controlRelationsLock.RLock()
	defer gs.controlRelationsLock.RUnlock()

	fil := gs.ControlRuleRelations.Filter(
		dataframe.F{Colname: "ControlID", Comparator: series.Eq, Comparando: control.ControlID},
	)
	var rulesList []opapolicy.PolicyRule
	var rulesIDList []string

	for row := 0; row < fil.Nrow(); row++ {
		ruleName := fil.Elem(row, 1)
		rule, err := gs.getOPAPolicyByName(ruleName.String()) // requires R-Lock on Rules
		if err != nil {
			return err
		}
		// add rule to control.rules
		rulesList = append(rulesList, *rule)
		// add ruleId ro control.ruleIds
		rulesIDList = append(rulesIDList, rule.GUID)
	}

	gs.controlEscalatedLock.Lock() // this locks all concurrent attempts to fill any control
	control.Rules = rulesList
	control.RulesIDs = &rulesIDList
	gs.controlEscalatedLock.Unlock()

	return nil
}

func (gs *GitRegoStore) fillControlsAndControlIDsInFramework(fw *opapolicy.Framework) error {
	gs.rulesLock.RLock()
	defer gs.rulesLock.RUnlock()

	gs.frameworkRelationsLock.RLock()
	defer gs.frameworkRelationsLock.RUnlock()

	fil := gs.FrameworkControlRelations.Filter(
		dataframe.F{Colname: "frameworkName", Comparator: series.Eq, Comparando: fw.Name},
	)
	var controlsList []opapolicy.Control
	var controlsIDList []string

	// if there are no controls in framework, need to populate them all from base controls.
	if len(fw.Controls) == 0 {
		for row := 0; row < fil.Nrow(); row++ {
			controlID := fil.Elem(row, 1)
			control, err := gs.getOPAControlByID(controlID.String()) // requires R-Lock on Rules
			if err != nil {
				return err
			}
			// add control to controlsList
			controlsList = append(controlsList, *control)
			// add controlID to controlsIDList
			controlsIDList = append(controlsIDList, control.ControlID)

		}

		gs.frameworkEscalatedLock.Lock()
		fw.Controls = controlsList
		fw.ControlsIDs = &controlsIDList
		gs.frameworkEscalatedLock.Unlock()
	} else {
		// if there are controls, need to populate only the rules.
		for i := range fw.Controls {
			if len(fw.Controls[i].Rules) == 0 {

				// getting the control object using GetOPAControlByID as it handles backward compatibility
				tmpControl, err := gs.getOPAControlByID(fw.Controls[i].ControlID)
				if err != nil {
					return err
				}

				gs.frameworkEscalatedLock.Lock()
				fw.Controls[i].Rules = tmpControl.Rules
				fw.Controls[i].RulesIDs = tmpControl.RulesIDs
				gs.frameworkEscalatedLock.Unlock()
			}

		}

	}

	return nil
}
