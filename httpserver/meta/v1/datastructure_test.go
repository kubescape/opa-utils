package v1

import (
	"encoding/json"
	"testing"
)

func TestPostScanRequestPreservesControlExclusions(t *testing.T) {
	var request PostScanRequest
	input := []byte(`{"targetType":"framework","targetNames":["nsa"],"excludeControls":["C-0069","C-0070"]}`)
	if err := json.Unmarshal(input, &request); err != nil {
		t.Fatal(err)
	}
	output, err := json.Marshal(request)
	if err != nil {
		t.Fatal(err)
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(output, &fields); err != nil {
		t.Fatal(err)
	}
	if string(fields["excludeControls"]) != `["C-0069","C-0070"]` {
		t.Fatalf("control exclusions lost in request round-trip: %s", output)
	}
}

func TestPostScanRequestOmitsEmptyControlExclusions(t *testing.T) {
	for _, input := range []string{`{}`, `{"excludeControls":[]}`, `{"excludeControls":null}`} {
		var request PostScanRequest
		if err := json.Unmarshal([]byte(input), &request); err != nil {
			t.Fatal(err)
		}
		output, err := json.Marshal(request)
		if err != nil {
			t.Fatal(err)
		}
		var fields map[string]json.RawMessage
		if err := json.Unmarshal(output, &fields); err != nil {
			t.Fatal(err)
		}
		if _, exists := fields["excludeControls"]; exists {
			t.Fatalf("empty exclusions must be omitted: %s", output)
		}
	}
}
