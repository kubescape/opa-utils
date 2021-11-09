package gitregostore

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	// "github.com/armosec/capacketsgo/opapolicy"
	opapolicy "github.com/armosec/opa-utils/reporthandling"
	"github.com/go-gota/gota/dataframe"
	"go.uber.org/zap"
)

type storeSetter func(*GitRegoStore, string) error

const (
	frameworksJsonFileName            = "frameworks"
	controlsJsonFileName              = "controls"
	rulesJsonFileName                 = "rules"
	frameworkControlRelationsFileName = "FWName_CID_CName"
	ControlRuleRelationsFileName      = "ControlID_RuleName"
)

var storeSetterMapping = map[string]storeSetter{
	frameworksJsonFileName:            (*GitRegoStore).setFrameworks,
	controlsJsonFileName:              (*GitRegoStore).setControls,
	rulesJsonFileName:                 (*GitRegoStore).setRules,
	frameworkControlRelationsFileName: (*GitRegoStore).setFrameworkControlRelationsFileName,
	ControlRuleRelationsFileName:      (*GitRegoStore).setControlRuleRelationsFileName,
}

type InnerTree []struct {
	PATH string `json:"path"`
}
type Tree struct {
	TREE InnerTree `json:"tree"`
}

// func setURL()
func (gs *GitRegoStore) setURL() {
	var url string

	if isUrlRelease(gs.Path) {
		url = gs.BaseUrl + "/" + gs.Owner + "/" + gs.Repository + "/" + gs.Path + "/" + gs.Tag
	} else {
		url = gs.BaseUrl + "/" + gs.Owner + "/" + gs.Repository + "/" + gs.Path + "/" + gs.Branch + "?recursive=1"
	}
	gs.URL = url
}

func (gs *GitRegoStore) setObjects() {
	if isUrlRelease(gs.URL) {
		gs.setObjectsFromReleaseLoop()
	} else {
		gs.setObjectsFromRepoLoop()
	}
}

func isUrlRelease(u string) bool {
	return strings.Contains(u, "releases")
}

// ========================== set Objects From Repo =====================================

func (gs *GitRegoStore) setObjectsFromRepoLoop() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		f := true
		for {
			if err := gs.setObjectsFromRepoOnce(); err != nil {
				fmt.Println(err)
			}
			if f {
				wg.Done() // first update to done
				f = false
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
	gs.frameworksLock.Lock()
	gs.controlsLock.Lock()
	gs.rulesLock.Lock()
	defer gs.frameworksLock.Unlock()
	defer gs.controlsLock.Unlock()
	defer gs.rulesLock.Unlock()
	gs.Frameworks = []opapolicy.Framework{}
	gs.Controls = []opapolicy.Control{}
	gs.Rules = []opapolicy.PolicyRule{}

	// use only json files from relevant dirs
	for _, path := range trees.TREE {
		rawDataPath := "https://raw.githubusercontent.com/" + gs.Owner + "/" + gs.Repository + "/" + gs.Branch + "/" + path.PATH

		// for kind, storeSetterMappingFunc := range storeSetterMapping {
		// 	respStr, err := HttpGetter(gs.httpClient, fmt.Sprintf("%s/%s", gs.URL, kind))
		// 	if err != nil {
		// 		return err
		// 	}
		// 	if err = storeSetterMappingFunc(gs, respStr); err != nil {
		// 		return err
		// 	}
		// }
		// // return nil

		if strings.HasPrefix(path.PATH, rulesJsonFileName) && strings.HasSuffix(path.PATH, ".json") {
			respStr, err := HttpGetter(gs.httpClient, rawDataPath)
			if err != nil {
				return err
			}
			if err := gs.setRulesWithRawRego(respStr, rawDataPath); err != nil {
				return err
			}
		} else if strings.HasPrefix(path.PATH, controlsJsonFileName) && strings.HasSuffix(path.PATH, ".json") {
			respStr, err := HttpGetter(gs.httpClient, rawDataPath)
			if err != nil {
				return err
			}
			if err := gs.setControl(respStr); err != nil {
				return err
			}
		} else if strings.HasPrefix(path.PATH, frameworksJsonFileName) && strings.HasSuffix(path.PATH, ".json") {
			respStr, err := HttpGetter(gs.httpClient, rawDataPath)
			if err != nil {
				return err
			}
			if err := gs.setFramework(respStr); err != nil {
				return err
			}
		} else if strings.HasSuffix(path.PATH, ControlRuleRelationsFileName+".csv") {
			respStr, err := HttpGetter(gs.httpClient, rawDataPath)
			if err != nil {
				return err
			}
			gs.setControlRuleRelationsFileName(respStr)
		} else if strings.HasSuffix(path.PATH, frameworkControlRelationsFileName+".csv") {
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
	gs.Frameworks = append(gs.Frameworks, *framework)
	return nil
}

func (gs *GitRegoStore) setControl(respStr string) error {
	control := &opapolicy.Control{}
	if err := JSONDecoder(respStr).Decode(control); err != nil {
		return err
	}
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
	gs.Rules = append(gs.Rules, *rule)
	return nil
}

// ======================== set Objects From Release =============================================

func (gs *GitRegoStore) setObjectsFromReleaseLoop() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		f := true
		for {
			if err := gs.setObjectsFromReleaseOnce(); err != nil {
				fmt.Println(err)
			}
			if f {
				wg.Done() // first update to done
				f = false
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

	for kind, storeSetterMappingFunc := range storeSetterMapping {
		respStr, err := HttpGetter(gs.httpClient, fmt.Sprintf("%s/%s", gs.URL, kind))
		if err != nil {
			return err
		}
		if err = storeSetterMappingFunc(gs, respStr); err != nil {
			return err
		}
	}
	return nil
}

// DEPRECATED
//
// func (gs *GitRegoStore) setObjectsFromReleaseOnce() error {

// 	// TODO - support mock respons
// 	resp, err := http.Get(gs.URL)
// 	if err != nil {
// 		return fmt.Errorf("failed to get latest releases from '%s', reason: %s", gs.URL, err.Error())
// 	}
// 	defer resp.Body.Close()
// 	if resp.StatusCode < 200 || 301 < resp.StatusCode {
// 		return fmt.Errorf("failed to download file, status code: %s", resp.Status)
// 	}
// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return fmt.Errorf("failed to read response body from '%s', reason: %s", gs.URL, err.Error())
// 	}
// 	var data map[string]interface{}
// 	err = json.Unmarshal(body, &data)
// 	if err != nil {
// 		return fmt.Errorf("failed to unmarshal response body from '%s', reason: %s", gs.URL, err.Error())
// 	}

// 	// TODO - move to parse response dedicated functions
// 	if tagName, ok := data["tag_name"]; ok {
// 		if gs.CurGitVersion != tagName.(string) {
// 			gs.CurGitVersion = tagName.(string)
// 			if assets, ok := data["assets"].([]interface{}); ok {
// 				for i := range assets {
// 					if asset, ok := assets[i].(map[string]interface{}); ok {
// 						if name, ok := asset["name"].(string); ok {
// 							if url, ok := asset["browser_download_url"].(string); ok {
// 								respStr, err := HttpGetter(gs.httpClient, url)
// 								if err != nil {
// 									return err
// 								}
// 								switch name {
// 								case frameworksJsonFileName:
// 									err = gs.setFrameworks(respStr)
// 								case controlsJsonFileName:
// 									err = gs.setControls(respStr)
// 								case rulesJsonFileName:
// 									err = gs.setRules(respStr)
// 								case frameworkControlRelationsFileName:
// 									gs.setFrameworkControlRelationsFileName(respStr)
// 								case ControlRuleRelationsFileName:
// 									gs.setControlRuleRelationsFileName(respStr)
// 								}
// 								if err != nil {
// 									return err
// 								}
// 							}
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return nil
// }

func (gs *GitRegoStore) setFrameworks(respStr string) error {
	frameworks := []opapolicy.Framework{}
	if err := JSONDecoder(respStr).Decode(&frameworks); err != nil {
		return err
	}
	gs.frameworksLock.Lock()
	defer gs.frameworksLock.Unlock()
	gs.Frameworks = frameworks
	return nil
}

func (gs *GitRegoStore) setControls(respStr string) error {
	controls := []opapolicy.Control{}
	if err := JSONDecoder(respStr).Decode(&controls); err != nil {
		return err
	}
	gs.controlsLock.Lock()
	defer gs.controlsLock.Unlock()
	gs.Controls = controls
	return nil
}

func (gs *GitRegoStore) setRules(respStr string) error {
	rules := &[]opapolicy.PolicyRule{}
	if err := JSONDecoder(respStr).Decode(rules); err != nil {
		return err
	}
	gs.rulesLock.Lock()
	defer gs.rulesLock.Unlock()
	gs.Rules = *rules
	return nil
}

func (gs *GitRegoStore) setFrameworkControlRelationsFileName(respStr string) error {
	df := dataframe.ReadCSV(strings.NewReader(respStr))
	gs.FrameworkControlRelations = df
	return nil
}

func (gs *GitRegoStore) setControlRuleRelationsFileName(respStr string) error {
	df := dataframe.ReadCSV(strings.NewReader(respStr))
	gs.ControlRuleRelations = df
	return nil
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
