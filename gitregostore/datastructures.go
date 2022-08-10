package gitregostore

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"sync"

	// "github.com/armosec/capacketsgo/opapolicy"
	"github.com/armosec/armoapi-go/armotypes"
	opapolicy "github.com/kubescape/opa-utils/reporthandling"

	"github.com/go-gota/gota/dataframe"
)

type GitRegoStore struct {
	Frameworks                  []opapolicy.Framework
	Controls                    []opapolicy.Control
	Rules                       []opapolicy.PolicyRule
	FrameworkControlRelations   dataframe.DataFrame
	ControlRuleRelations        dataframe.DataFrame
	DefaultConfigInputs         armotypes.CustomerConfig
	frameworksLock              sync.RWMutex
	controlsLock                sync.RWMutex
	rulesLock                   sync.RWMutex
	DefaultConfigInputsLock     sync.RWMutex
	URL                         string
	httpClient                  *http.Client
	BaseUrl                     string
	Owner                       string
	Repository                  string
	Path                        string
	Tag                         string
	Branch                      string
	FrequencyPullFromGitMinutes int
	CurGitVersion               string
	Watch                       bool
}

func newGitRegoStore(baseUrl string, owner string, repository string, path string, tag string, branch string, frequency int) *GitRegoStore {
	watch := false
	if frequency > 0 {
		watch = true
	}
	return &GitRegoStore{httpClient: &http.Client{},
		BaseUrl:                     baseUrl,
		Owner:                       owner,
		Repository:                  repository,
		Path:                        path,
		Tag:                         tag,
		Branch:                      branch,
		FrequencyPullFromGitMinutes: frequency,
		Watch:                       watch,
	}
}

// NewGitRegoStore return gitregostore obj with basic fields, before pulling from git
func NewGitRegoStore(baseUrl string, owner string, repository string, path string, tag string, branch string, frequency int) *GitRegoStore {
	gs := newGitRegoStore(baseUrl, owner, repository, path, tag, branch, frequency)
	gs.setURL()
	return gs
}

// SetRegoObjects pulls opa obj from git and stores in gitregostore
func (gs *GitRegoStore) SetRegoObjects() error {
	err := gs.setObjects()
	return err
}

func NewDefaultGitRegoStore(frequency int) *GitRegoStore {
	return NewGitRegoStore("https://github.com", "armosec", "regolibrary", "releases", "latest/download", "", frequency)
}

// Deprecated
// if frequency < 0 will pull only once
func InitGitRegoStore(baseUrl string, owner string, repository string, path string, tag string, branch string, frequency int) *GitRegoStore {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("InitGitRegoStore failed: stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	gs := newGitRegoStore(baseUrl, owner, repository, path, tag, branch, frequency)
	gs.setURL()
	gs.setObjects()
	return gs
}

// Deprecated
func InitDefaultGitRegoStore(frequency int) *GitRegoStore {
	return InitGitRegoStore("https://github.com", "armosec", "regolibrary", "releases", "latest/download", "", frequency)
}
