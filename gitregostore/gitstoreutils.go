package gitregostore

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
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
		gs.setObjectsFromReleaseLoop()
	} else {
		gs.setObjectsFromRepoLoop()
	}
}

// ========================== set Objects From Repo =====================================

func (gs *GitRegoStore) setObjectsFromRepoLoop() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			if err := gs.setObjectsFromRepoOnce(); err != nil {
				fmt.Println(err)
			}
			if !gs.Watch {
				return
			}
			time.Sleep(time.Duration(gs.FrequencyPullFromGitMinutes) * time.Minute)
		}
	}()
	wg.Wait()
}

func (gs *GitRegoStore) setObjectsFromRepoOnce() error {
	body, err := HttpGetter(gs.httpClient, gs.URL)
	if err != nil {
		return err
	}
	var trees Tree
	err = json.Unmarshal([]byte(body), &trees)
	if err != nil {
		return fmt.Errorf("failed to unmarshal response body from '%s', reason: %s", gs.URL, err.Error())
	}

	// use only json files from relevant dirs
	for _, path := range trees.TREE {
		rawDataPath := "https://raw.githubusercontent.com/" + gs.Owner + "/" + gs.Repository + "/" + gs.Branch + "/" + path.PATH
		if strings.HasPrefix(path.PATH, "rules") && strings.HasSuffix(path.PATH, ".json") {
			respStr, err := HttpGetter(gs.httpClient, rawDataPath)
			if err != nil {
				return err
			}
			if err := gs.setRulesWithRawRego(respStr, rawDataPath); err != nil {
				return err
			}
		} else if strings.HasPrefix(path.PATH, "controls") && strings.HasSuffix(path.PATH, ".json") {
			respStr, err := HttpGetter(gs.httpClient, rawDataPath)
			if err != nil {
				return err
			}
			gs.setControl(respStr)
		} else if strings.HasPrefix(path.PATH, "frameworks") && strings.HasSuffix(path.PATH, ".json") {
			respStr, err := HttpGetter(gs.httpClient, rawDataPath)
			if err != nil {
				return err
			}
			gs.setFramework(respStr)
		} else if strings.HasSuffix(path.PATH, "ControlID_RuleName.csv") {
			respStr, err := HttpGetter(gs.httpClient, rawDataPath)
			if err != nil {
				return err
			}
			gs.setControlRuleRelationsFileName(respStr)
		} else if strings.HasSuffix(path.PATH, "FWName_CID_CName.csv") {
			respStr, err := HttpGetter(gs.httpClient, rawDataPath)
			if err != nil {
				return err
			}
			gs.setFrameworkControlRelationsFileName(respStr)
		}
	}
	return nil
}

func (gs *GitRegoStore) setFramework(respStr string) error {
	framework := &opapolicy.Framework{}
	if err := JSONDecoder(respStr).Decode(framework); err != nil {
		return err
	}
	gs.frameworksLock.Lock()
	defer gs.frameworksLock.Unlock()
	gs.Frameworks = append(gs.Frameworks, *framework)
	return nil
}

func (gs *GitRegoStore) setControl(respStr string) error {
	control := &opapolicy.Control{}
	if err := JSONDecoder(respStr).Decode(control); err != nil {
		return err
	}
	gs.controlsLock.Lock()
	defer gs.controlsLock.Unlock()
	gs.Controls = append(gs.Controls, *control)

	return nil
}

func (gs *GitRegoStore) setRulesWithRawRego(respStr string, path string) error {
	rule := &opapolicy.PolicyRule{}
	if err := JSONDecoder(respStr).Decode(rule); err != nil {
		return err
	}
	rawRegoPath := path[:strings.LastIndex(path, "/")] + "/raw.rego"
	respString, err := HttpGetter(gs.httpClient, rawRegoPath)
	if err != nil {
		return err
	}
	rule.Rule = respString

	gs.rulesLock.Lock()
	defer gs.rulesLock.Unlock()
	gs.Rules = append(gs.Rules, *rule)

	return nil
}

// ======================== set Objects From Release =============================================

func (gs *GitRegoStore) setObjectsFromReleaseLoop() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if err := gs.setObjectsFromReleaseOnce(); err != nil {
				fmt.Println(err)
			}
			if !gs.Watch {
				return
			}
			time.Sleep(time.Duration(gs.FrequencyPullFromGitMinutes) * time.Minute)
		}
	}()
	wg.Wait()
}

func (gs *GitRegoStore) setObjectsFromReleaseOnce() error {

	// TODO - support mock respons
	resp, err := http.Get(gs.URL)
	if err != nil {
		return fmt.Errorf("failed to get latest releases from '%s', reason: %s", gs.URL, err.Error())
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || 301 < resp.StatusCode {
		return fmt.Errorf("failed to download file, status code: %s", resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body from '%s', reason: %s", gs.URL, err.Error())
	}
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return fmt.Errorf("failed to unmarshal response body from '%s', reason: %s", gs.URL, err.Error())
	}

	// TODO - move to parse response dedicated functions
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
									return err
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
	return nil
}

func (gs *GitRegoStore) setFrameworks(respStr string) {
	frameworks := &[]opapolicy.Framework{}
	if err := JSONDecoder(respStr).Decode(frameworks); err != nil {
		log.Print(err.Error())
	}
	gs.frameworksLock.Lock()
	defer gs.frameworksLock.Unlock()
	gs.Frameworks = *frameworks
}

func (gs *GitRegoStore) setControls(respStr string) {
	controls := &[]opapolicy.Control{}
	if err := JSONDecoder(respStr).Decode(controls); err != nil {
		log.Print(err.Error())
	}
	gs.controlsLock.Lock()
	defer gs.controlsLock.Unlock()
	gs.Controls = *controls
}

func (gs *GitRegoStore) setRules(respStr string) {
	rules := &[]opapolicy.PolicyRule{}
	if err := JSONDecoder(respStr).Decode(rules); err != nil {
		log.Print(err.Error())
	}
	gs.rulesLock.Lock()
	defer gs.rulesLock.Unlock()
	gs.Rules = *rules
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
