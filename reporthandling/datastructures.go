package reporthandling

import (
	"github.com/armosec/armoapi-go/armotypes"
)

const (
	ControlAttributeKeyTypeTag      = "controlTypeTags"
	ControlAttributeKeyAttackTracks = "attackTracks"
)

type RuleLanguages string

const (
	RegoLanguage  RuleLanguages = "Rego"
	RegoLanguage2 RuleLanguages = "rego"
)

// RuleMatchObjects defines which objects this rule applied on
type RuleMatchObjects struct {
	APIGroups   []string `json:"apiGroups" bson:"apiGroups"`     // apps
	APIVersions []string `json:"apiVersions" bson:"apiVersions"` // v1/ v1beta1 / *
	Resources   []string `json:"resources" bson:"resources"`     // dep.., pods,
}

type RuleDependency struct {
	PackageName string `json:"packageName" bson:"packageName"`
}

type ControlConfigInputs struct {
	Path        string `json:"path" bson:"path"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
}

// PolicyRule represents single rule, the fundamental executable block of policy
type PolicyRule struct {
	armotypes.PortalBase   `json:",inline" bson:"inline"`
	CreationTime           string                `json:"creationTime" bson:"creationTime"`
	Rule                   string                `json:"rule" bson:"rule"`                                     // multiline string of raw.rego
	ResourceEnumerator     string                `json:"resourceEnumerator" bson:"resourceEnumerator"`         // multiline string of filter.rego, if exists
	RuleLanguage           RuleLanguages         `json:"ruleLanguage" bson:"ruleLanguage"`                     // default "rego"
	Match                  []RuleMatchObjects    `json:"match" bson:"match"`                                   // k8s resources this rule needs as inputs
	DynamicMatch           []RuleMatchObjects    `json:"dynamicMatch,omitempty" bson:"dynamicMatch,omitempty"` // NON-k8s resources this rule needs as inputs, acquired by host-scanner
	RuleDependencies       []RuleDependency      `json:"ruleDependencies" bson:"ruleDependencies"`             // packages this rule uses
	ConfigInputs           []string              `json:"configInputs" bson:"configInputs"`                     // DEPRECATED
	ControlConfigInputs    []ControlConfigInputs `json:"controlConfigInputs" bson:"controlConfigInputs" `      // list of inputs from postureControlInputs in customerConfig for this rule
	Description            string                `json:"description" bson:"description"`
	Remediation            string                `json:"remediation" bson:"remediation"`
	RuleQuery              string                `json:"ruleQuery" bson:"ruleQuery" `                          // default "armo_builtins" - DEPRECATED
	RelevantCloudProviders []string              `json:"relevantCloudProviders" bson:"relevantCloudProviders"` // rule is relevant only to clusters in these cloud providers
}

// Control represents a collection of rules which are combined together to single purpose
type Control struct {
	armotypes.PortalBase `json:",inline" bson:"inline"`
	Control_ID           string              `json:"id,omitempty" bson:"id,omitempty"  ` // to be Deprecated
	ControlID            string              `json:"controlID" bson:"controlID"`
	CreationTime         string              `json:"creationTime" bson:"creationTime"`
	Description          string              `json:"description" bson:"description"`
	Remediation          string              `json:"remediation" bson:"remediation"`
	Rules                []PolicyRule        `json:"rules" bson:"rules,omitempty"`
	FrameworkNames       []string            `json:"frameworkNames,omitempty" bson:"frameworkNames,omitempty"` // frameworks this control is part of
	FixedInput           map[string][]string `json:"fixedInput,omitempty"`                                     // DEPRECATED
	// for new list of  rules in POST/UPADTE requests
	RulesIDs              *[]string `json:"rulesIDs,omitempty" bson:"rulesIDs,omitempty"`
	BaseScore             float32   `json:"baseScore,omitempty" bson:"baseScore,omitempty"`
	ARMOImprovementFactor float32   `json:"ARMOImprovementFactor,omitempty" bson:"ARMOImprovementFactor,omitempty"`
}

type UpdatedControl struct {
	Control `json:",inline"`
	Rules   []interface{} `json:"rules"`
}

// Framework represents a collection of controls which are combined together to expose comprehensive behavior
type Framework struct {
	armotypes.PortalBase `json:",inline" bson:"inline"`
	CreationTime         string    `json:"creationTime" bson:"creationTime"`
	Description          string    `json:"description" bson:"description"`
	Controls             []Control `json:"controls" bson:"-"`
	// for new list of  controls in POST/UPADTE requests
	ControlsIDs *[]string                       `json:"controlsIDs,omitempty" bson:"controlsIDs,omitempty"`
	SubSections map[string]*FrameworkSubSection `json:"subSections,omitempty" bson:"subSections,omitempty"`
}

type UpdatedFramework struct {
	Framework `json:",inline"`
	Controls  []interface{} `json:"controls"`
}

type FrameworkSubSection struct {
	armotypes.PortalBase `json:",inline" bson:"inline"`
	ID                   string                          `json:"id" bson:"id"`                                       // unique id inside the framework.
	SubSections          map[string]*FrameworkSubSection `json:"subSections,omitempty" bson:"subSections,omitempty"` // inner subsection
	ControlIDs           []string                        `json:"controlsIDs,omitempty" bson:"controlsIDs,omitempty"` // control ids.
	Controls             []*Control                      `json:"-" bson:"-"`                                         // controls list for fast access
}

type AttackTrackCategories struct {
	AttackTrack string
	Categories  []string
}
