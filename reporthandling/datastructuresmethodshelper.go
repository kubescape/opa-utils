package reporthandling

import (
	"bytes"
	"encoding/json"
)

type ResourcesIDs struct {
	allResources     []string
	failedResources  []string
	warningResources []string
}

func (r *ResourcesIDs) append(a *ResourcesIDs) {
	r.setAllResources(append(r.allResources, a.allResources...))
	r.setFailedResources(append(r.failedResources, a.failedResources...))
	r.setWarningResources(append(r.warningResources, a.warningResources...))
}
func (r *ResourcesIDs) GetAllResources() []string {
	return r.allResources
}
func (r *ResourcesIDs) GetFailedResources() []string {
	return r.failedResources
}
func (r *ResourcesIDs) GetWarningResources() []string {
	return r.warningResources
}

func (r *ResourcesIDs) setAllResources(a []string) {
	r.allResources = GetUniqueResourcesIDs(a)
}
func (r *ResourcesIDs) setFailedResources(a []string) {
	r.failedResources = GetUniqueResourcesIDs(a)
}
func (r *ResourcesIDs) setWarningResources(a []string) {
	r.warningResources = TrimUniqueIDs(GetUniqueResourcesIDs(a), r.failedResources)
}

func (pn *PolicyNotification) ToJSONBytesBuffer() (*bytes.Buffer, error) {
	res, err := json.Marshal(pn)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(res), err
}

func deleteFromMap(m map[string]interface{}, keepFields []string) {
	for k := range m {
		if StringInSlice(keepFields, k) {
			continue
		}
		delete(m, k)
	}
}

func StringInSlice(strSlice []string, str string) bool {
	for i := range strSlice {
		if strSlice[i] == str {
			return true
		}
	}
	return false
}

func RemoveResponse(slice []RuleResponse, index int) []RuleResponse {
	return append(slice[:index], slice[index+1:]...)
}

func percentage(big, small int) int {
	if big == 0 {
		if small == 0 {
			return 100
		}
		return 0
	}
	return int(float64(float64(big-small)/float64(big)) * 100)
}
