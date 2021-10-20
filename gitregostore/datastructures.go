package gitregostore

import (
	"net/http"
	"sync"

	// "github.com/armosec/capacketsgo/opapolicy"
	opapolicy "github.com/armosec/opa-utils/reporthandling"

	"github.com/go-gota/gota/dataframe"
)

type GitRegoStore struct {
	Frameworks                  []opapolicy.Framework
	Controls                    []opapolicy.Control
	Rules                       []opapolicy.PolicyRule
	FrameworkControlRelations   dataframe.DataFrame
	ControlRuleRelations        dataframe.DataFrame
	frameworksLock              sync.RWMutex
	controlsLock                sync.RWMutex
	rulesLock                   sync.RWMutex
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
}

func newGitRegoStore(baseUrl string, owner string, repository string, path string, tag string, branch string, frequency int) *GitRegoStore {
	return &GitRegoStore{httpClient: &http.Client{},
		BaseUrl:                     baseUrl,
		Owner:                       owner,
		Repository:                  repository,
		Path:                        path,
		Tag:                         tag,
		Branch:                      branch,
		FrequencyPullFromGitMinutes: frequency}
}

func InitGitRegoStore(baseUrl string, owner string, repository string, path string, tag string, branch string, frequency int) *GitRegoStore {
	gs := newGitRegoStore(baseUrl, owner, repository, path, tag, branch, frequency)
	gs.setURL()
	gs.setObjects()
	return gs
}

func InitDefaultGitRegoStore() *GitRegoStore {
	return InitGitRegoStore("https://api.github.com/repos", "armosec", "regolibrary", "releases", "latest", "master", 1)
}
