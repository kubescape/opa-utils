package reporthandling

import (
	"bytes"
	"encoding/json"
)

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
