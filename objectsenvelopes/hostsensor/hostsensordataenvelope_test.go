package hostsensor

import (
	"testing"
)

func TestIsTypeTypeHostSensor(t *testing.T) {
	tests := []struct {
		name   string
		object map[string]interface{}
		want   bool
	}{
		{
			name: "valid apiVersion",
			object: map[string]interface{}{
				"apiVersion": "hostdata.kubescape.cloud/v1",
			},
			want: true,
		},
		{
			name: "apiVersion does not match GroupHostSensor",
			object: map[string]interface{}{
				"apiVersion": "someOtherGroup/v1",
			},
			want: false,
		},
		{
			name: "apiVersion is an integer",
			object: map[string]interface{}{
				"apiVersion": 12345,
			},
			want: false,
		},
		{
			name: "missing apiVersion",
			object: map[string]interface{}{
				"someOtherKey": "someValue",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsTypeTypeHostSensor(tt.object); got != tt.want {
				t.Errorf("%q: IsTypeTypeHostSensor() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

