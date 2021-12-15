package v2

import "encoding/json"

func (pr *PostureReport) ToBytes() ([]byte, error) {
	return json.Marshal(*pr)
}

func (pr *PostureReport) ToString() string {
	if bpr, err := pr.ToBytes(); err == nil {
		return string(bpr)
	}
	return "{}"
}
