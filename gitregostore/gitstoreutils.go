package gitregostore

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	// "github.com/armosec/capacketsgo/opapolicy"
	opapolicy "github.com/armosec/opa-utils/reporthandling"
	"github.com/go-gota/gota/dataframe"
	"go.uber.org/zap"
)

var (
	frameworksJsonFileName            = "frameworks"
	controlsJsonFileName              = "controls"
	rulesJsonFileName                 = "rules"
	frameworkControlRelationsFileName = "FWName_CID_CName"
	ControlRuleRelationsFileName      = "ControlID_RuleName"
)

type InnerTree []struct {
	PATH string `json:"path"`
}
type Tree struct {
	TREE InnerTree `json:"tree"`
}

func (gs *GitRegoStore) setURL() {
	var url string
	if gs.Path == "releases" {
		url = gs.BaseUrl + "/" + gs.Owner + "/" + gs.Repository + "/" + gs.Path + "/" + gs.Tag
		// "https://api.github.com/repos/armosec/regolibrary/releases/latest"
	} else {
		url = gs.BaseUrl + "/" + gs.Owner + "/" + gs.Repository + "/" + gs.Path + "/" + gs.Branch + "?recursive=1"
	}
	gs.URL = url
}

func (gs *GitRegoStore) setObjects() {
	if strings.Contains(gs.URL, "releases") {
		gs.setObjectsFromReleaseOnce()
		go gs.setObjectsFromReleaseLoop()
	} else {
		gs.setObjectsFromRepoOnce()
		go gs.setObjectsFromRepoLoop()
	}
}

// ========================== set Objects From Repo =====================================

func (gs *GitRegoStore) setObjectsFromRepoLoop() {
	for {
		gs.setObjectsFromRepoOnce()
		time.Sleep(time.Duration(gs.FrequencyPullFromGitMinutes) * time.Minute)
	}
}

func (gs *GitRegoStore) setObjectsFromRepoOnce() {
	body, err := HttpGetter(gs.httpClient, gs.URL)
	if err != nil {
		log.Print(err.Error())
	}
	var trees Tree
	err = json.Unmarshal([]byte(body), &trees)
	if err != nil {
		log.Print(fmt.Errorf("failed to unmarshal response body from '%s', reason: %s", gs.URL, err.Error()))
	}

	// use only json files from relevant dirs
	for _, path := range trees.TREE {
		rawDataPath := "https://raw.githubusercontent.com/" + gs.Owner + "/" + gs.Repository + "/" + gs.Branch + "/" + path.PATH
		if strings.HasPrefix(path.PATH, "rules") && strings.HasSuffix(path.PATH, ".json") {
			respStr, err := HttpGetter(gs.httpClient, rawDataPath)
			if err != nil {
				log.Print(err.Error())
			}
			gs.setRulesWithRawRego(respStr, rawDataPath)
		} else if strings.HasPrefix(path.PATH, "controls") && strings.HasSuffix(path.PATH, ".json") {
			respStr, err := HttpGetter(gs.httpClient, rawDataPath)
			if err != nil {
				log.Print(err.Error())
			}
			gs.setControl(respStr)
		} else if strings.HasPrefix(path.PATH, "frameworks") && strings.HasSuffix(path.PATH, ".json") {
			respStr, err := HttpGetter(gs.httpClient, rawDataPath)
			if err != nil {
				log.Print(err.Error())
			}
			gs.setFramework(respStr)
		} else if strings.HasSuffix(path.PATH, "ControlID_RuleName.csv") {
			respStr, err := HttpGetter(gs.httpClient, rawDataPath)
			if err != nil {
				log.Print(err.Error())
			}
			gs.setControlRuleRelationsFileName(respStr)
		} else if strings.HasSuffix(path.PATH, "FWName_CID_CName.csv") {
			respStr, err := HttpGetter(gs.httpClient, rawDataPath)
			if err != nil {
				log.Print(err.Error())
			}
			gs.setFrameworkControlRelationsFileName(respStr)
		}
		// gs.setRelationObjectsFromJsons()
	}
}

func (gs *GitRegoStore) setFramework(respStr string) {
	framework := &opapolicy.Framework{}
	if err := JSONDecoder(respStr).Decode(framework); err != nil {
		log.Print(err.Error())
	}
	gs.frameworksLock.Lock()
	gs.Frameworks = append(gs.Frameworks, *framework)
	gs.frameworksLock.Unlock()
}

func (gs *GitRegoStore) setControl(respStr string) {
	control := &opapolicy.Control{}
	if err := JSONDecoder(respStr).Decode(control); err != nil {
		log.Print(err.Error())
	}
	gs.controlsLock.Lock()
	gs.Controls = append(gs.Controls, *control)
	gs.controlsLock.Unlock()
}

func (gs *GitRegoStore) setRulesWithRawRego(respStr string, path string) {
	rule := &opapolicy.PolicyRule{}
	if err := JSONDecoder(respStr).Decode(rule); err != nil {
		log.Print(err.Error())
	}
	rawRegoPath := path[:strings.LastIndex(path, "/")] + "/raw.rego"
	respString, err := HttpGetter(gs.httpClient, rawRegoPath)
	if err != nil {
		log.Print(err.Error())
	}
	rule.Rule = respString
	gs.rulesLock.Lock()
	gs.Rules = append(gs.Rules, *rule)
	gs.rulesLock.Unlock()
}

// ======================== set Objects From Release =============================================

func (gs *GitRegoStore) setObjectsFromReleaseLoop() {
	for {
		gs.setObjectsFromReleaseOnce()
		time.Sleep(time.Duration(gs.FrequencyPullFromGitMinutes) * time.Minute)
	}
}

func (gs *GitRegoStore) setObjectsFromReleaseOnce() {
	resp, err := http.Get(gs.URL)
	if err != nil {
		log.Printf("failed to get latest releases from '%s', reason: %s", gs.URL, err.Error())
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || 301 < resp.StatusCode {
		log.Printf("failed to download file, status code: %s", resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("failed to read response body from '%s', reason: %s", gs.URL, err.Error())
	}
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Printf("failed to unmarshal response body from '%s', reason: %s", gs.URL, err.Error())
	}
	if tagName, ok := data["tag_name"]; ok {
		if gs.CurGitVersion != tagName.(string) {
			gs.CurGitVersion = tagName.(string)
			if assets, ok := data["assets"].([]interface{}); ok {
				for i := range assets {
					if asset, ok := assets[i].(map[string]interface{}); ok {
						if name, ok := asset["name"].(string); ok {
							if url, ok := asset["browser_download_url"].(string); ok {
								respStr, err := HttpGetter(gs.httpClient, url)
								if err != nil {
									log.Print(err.Error())
								}
								switch name {
								case frameworksJsonFileName:
									gs.setFrameworks(respStr)
								case controlsJsonFileName:
									gs.setControls(respStr)
								case rulesJsonFileName:
									gs.setRules(respStr)
								case frameworkControlRelationsFileName:
									gs.setFrameworkControlRelationsFileName(respStr)
								case ControlRuleRelationsFileName:
									gs.setControlRuleRelationsFileName(respStr)
								}
							}
						}
					}
				}
			}
		}
	}
}

func (gs *GitRegoStore) setFrameworks(respStr string) {
	frameworks := &[]opapolicy.Framework{}
	if err := JSONDecoder(respStr).Decode(frameworks); err != nil {
		log.Print(err.Error())
	}
	gs.frameworksLock.Lock()
	gs.Frameworks = *frameworks
	gs.frameworksLock.Unlock()
}

func (gs *GitRegoStore) setControls(respStr string) {
	controls := &[]opapolicy.Control{}
	if err := JSONDecoder(respStr).Decode(controls); err != nil {
		log.Print(err.Error())
	}
	gs.controlsLock.Lock()
	gs.Controls = *controls
	gs.controlsLock.Unlock()
}

func (gs *GitRegoStore) setRules(respStr string) {
	rules := &[]opapolicy.PolicyRule{}
	if err := JSONDecoder(respStr).Decode(rules); err != nil {
		log.Print(err.Error())
	}
	gs.rulesLock.Lock()
	gs.Rules = *rules
	gs.rulesLock.Unlock()
}

func (gs *GitRegoStore) setFrameworkControlRelationsFileName(respStr string) {
	df := dataframe.ReadCSV(strings.NewReader(respStr))
	gs.FrameworkControlRelations = df
}

func (gs *GitRegoStore) setControlRuleRelationsFileName(respStr string) {
	df := dataframe.ReadCSV(strings.NewReader(respStr))
	gs.ControlRuleRelations = df
}

// JSONDecoder returns JSON decoder for given string
func JSONDecoder(origin string) *json.Decoder {
	dec := json.NewDecoder(strings.NewReader(origin))
	dec.UseNumber()
	return dec
}

func HttpGetter(httpClient *http.Client, fullURL string) (string, error) {
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return "", err
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	respStr, err := HTTPRespToString(resp)
	if err != nil {
		return "", err
	}
	return respStr, nil
}

// HTTPRespToString parses the body as string and checks the HTTP status code, it closes the body reader at the end
func HTTPRespToString(resp *http.Response) (string, error) {
	if resp == nil || resp.Body == nil {
		return "", nil
	}
	strBuilder := strings.Builder{}
	defer resp.Body.Close()
	if resp.ContentLength > 0 {
		strBuilder.Grow(int(resp.ContentLength))
	}
	bytesNum, err := io.Copy(&strBuilder, resp.Body)
	respStr := strBuilder.String()
	if err != nil {
		respStrNewLen := len(respStr)
		if respStrNewLen > 1024 {
			respStrNewLen = 1024
		}
		return "", fmt.Errorf("HTTP request failed. URL: '%s', Read-ERROR: '%s', HTTP-CODE: '%s', BODY(top): '%s', HTTP-HEADERS: %v, HTTP-BODY-BUFFER-LENGTH: %v", resp.Request.URL.RequestURI(), err, resp.Status, respStr[:respStrNewLen], resp.Header, bytesNum)
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		respStrNewLen := len(respStr)
		if respStrNewLen > 1024 {
			respStrNewLen = 1024
		}
		err = fmt.Errorf("HTTP request failed. URL: '%s', HTTP-ERROR: '%s', BODY: '%s', HTTP-HEADERS: %v, HTTP-BODY-BUFFER-LENGTH: %v", resp.Request.URL.RequestURI(), resp.Status, respStr[:respStrNewLen], resp.Header, bytesNum)
	}
	zap.L().Debug("In HTTPRespToString - request end succesfully",
		zap.String("URL", resp.Request.URL.String()), zap.Int("contentLength", int(resp.ContentLength)))

	return respStr, err
}
