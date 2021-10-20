package gitregostore

import (
	"fmt"
	"strings"

	// "github.com/armosec/capacketsgo/opapolicy"
	opapolicy "github.com/armosec/opa-utils/reporthandling"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

// GetOPAPolicies returns all the policies of given customer
func (gs *GitRegoStore) GetOPAPolicies() ([]opapolicy.PolicyRule, error) {
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

// GetOPAPolicy returns specific policy
func (gs *GitRegoStore) GetOPAPolicy(policyGUID string) (*opapolicy.PolicyRule, error) {
	gs.rulesLock.RLock()
	defer gs.rulesLock.RUnlock()
	for _, rule := range gs.Rules {
		if rule.GUID == policyGUID {
			return &rule, nil
		}
	}
	return nil, fmt.Errorf("rule '%s' not found", policyGUID)
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

// GetOPAControlByName returns specific control by the name
func (gs *GitRegoStore) GetOPAControlByName(controlName string) (*opapolicy.Control, error) {
	gs.controlsLock.RLock()
	defer gs.controlsLock.RUnlock()
	for _, control := range gs.Controls {
		if strings.EqualFold(control.Name, controlName) {
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
		if strings.EqualFold(control.ControlID, controlID) {
			err := gs.fillRulesAndRulesIDsInControl(&control)
			if err != nil {
				return nil, err
			}
			return &control, nil
		}
	}
	return nil, fmt.Errorf("control '%s' not found", controlID)
}

// GetOPAControls returns all the controls of given customer
func (gs *GitRegoStore) GetOPAControls() ([]opapolicy.Control, error) {
	gs.controlsLock.RLock()
	defer gs.controlsLock.RUnlock()
	var controlsList []opapolicy.Control
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
	fmt.Printf("in GetOPAControlsNamesList")
	gs.controlsLock.RLock()
	defer gs.controlsLock.RUnlock()
	var controlsNameList []string
	for _, control := range gs.Controls {
		controlsNameList = append(controlsNameList, control.Name)
	}
	return controlsNameList, nil
}

func (gs *GitRegoStore) GetOPAControlsIDsList() ([]string, error) {
	fmt.Printf("in GetOPAControlsNamesList")
	gs.controlsLock.RLock()
	defer gs.controlsLock.RUnlock()
	var controlsIDList []string
	for _, control := range gs.Controls {
		controlsIDList = append(controlsIDList, control.ControlID)
	}
	return controlsIDList, nil
}

// GetOPAControl returns specific  policy
func (gs *GitRegoStore) GetOPAControl(policyGUID string) (*opapolicy.Control, error) {
	gs.controlsLock.RLock()
	defer gs.controlsLock.RUnlock()
	for _, control := range gs.Controls {
		if control.Name == policyGUID {
			err := gs.fillRulesAndRulesIDsInControl(&control)
			if err != nil {
				return nil, err
			}
			return &control, nil
		}
	}
	return nil, fmt.Errorf("control '%s' not found", policyGUID)
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
		controlsIDList = append(controlsIDList, control.GUID)
	}
	fw.Controls = controlsList
	fw.ControlsIDs = &controlsIDList
	return nil
}

// GetOPAFrameworks returns all the frameworks of given customer
func (gs *GitRegoStore) GetOPAFrameworks() ([]opapolicy.Framework, error) {
	gs.frameworksLock.RLock()
	defer gs.frameworksLock.RUnlock()
	var frameworksList []opapolicy.Framework
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

// GetOPAFramework returns specific framework
func (gs *GitRegoStore) GetOPAFramework(frameworkGUID string) (*opapolicy.Framework, error) {
	gs.frameworksLock.RLock()
	defer gs.frameworksLock.RUnlock()
	for _, fw := range gs.Frameworks {
		if fw.GUID == frameworkGUID {
			err := gs.fillControlsAndControlIDsInFramework(&fw)
			if err != nil {
				return nil, err
			}
			return &fw, nil
		}
	}
	return nil, fmt.Errorf("control '%s' not found", frameworkGUID)
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
	return nil, fmt.Errorf("control '%s' not found", frameworkName)
}
