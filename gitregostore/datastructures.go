package gitregostore

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"sync"

	// "github.com/armosec/capacketsgo/opapolicy"
	"github.com/armosec/armoapi-go/armotypes"
	opapolicy "github.com/kubescape/opa-utils/reporthandling"
	"github.com/kubescape/opa-utils/reporthandling/attacktrack/v1alpha1"

	"github.com/go-gota/gota/dataframe"
)

type GitRegoStore struct {
	frameworksLock              sync.RWMutex
	DefaultConfigInputsLock     sync.RWMutex
	rulesLock                   sync.RWMutex
	controlsLock                sync.RWMutex
	attackTracksLock            sync.RWMutex
	ControlRuleRelations        dataframe.DataFrame
	FrameworkControlRelations   dataframe.DataFrame
	httpClient                  *http.Client
	Tag                         string
	Owner                       string
	CurGitVersion               string
	Branch                      string
	URL                         string
	Path                        string
	BaseUrl                     string
	Repository                  string
	DefaultConfigInputs         armotypes.CustomerConfig
	AttackTracks                []v1alpha1.AttackTrack
	Frameworks                  []opapolicy.Framework
	Controls                    []opapolicy.Control
	Rules                       []opapolicy.PolicyRule
	FrequencyPullFromGitMinutes int
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
	return NewGitRegoStore("https://github.com", "kubescape", "regolibrary", "releases", "latest/download", "", frequency)
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
	return InitGitRegoStore("https://github.com", "kubescape", "regolibrary", "releases", "latest/download", "", frequency)
}
