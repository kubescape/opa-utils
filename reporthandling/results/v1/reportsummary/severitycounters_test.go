package reportsummary

import (
	"testing"

	"github.com/kubescape/opa-utils/reporthandling/apis"
)

func TestSeverityCounters_NumberOfResourcesWithCriticalSeverity(t *testing.T) {
	type fields struct {
		ResourcesWithCriticalSeverityCounter int
		ResourcesWithHighSeverityCounter     int
		ResourcesWithMediumSeverityCounter   int
		ResourcesWithLowSeverityCounter      int
	}
	tests := []struct {
		name   string
		fields fields
		want   fields
	}{
		{
			name: "test return methods",
			fields: fields{
				ResourcesWithCriticalSeverityCounter: 1,
				ResourcesWithHighSeverityCounter:     2,
				ResourcesWithMediumSeverityCounter:   3,
				ResourcesWithLowSeverityCounter:      4,
			},
			want: fields{
				ResourcesWithCriticalSeverityCounter: 1,
				ResourcesWithHighSeverityCounter:     2,
				ResourcesWithMediumSeverityCounter:   3,
				ResourcesWithLowSeverityCounter:      4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sc := &SeverityCounters{
				CriticalSeverityCounter: tt.fields.ResourcesWithCriticalSeverityCounter,
				HighSeverityCounter:     tt.fields.ResourcesWithHighSeverityCounter,
				MediumSeverityCounter:   tt.fields.ResourcesWithMediumSeverityCounter,
				LowSeverityCounter:      tt.fields.ResourcesWithLowSeverityCounter,
			}
			if got := sc.NumberOfCriticalSeverity(); got != tt.want.ResourcesWithCriticalSeverityCounter {
				t.Errorf("SeverityCounters.NumberOfResourcesWithCriticalSeverity() = %v, want %v", got, tt.want)
			}
			if got := sc.NumberOfHighSeverity(); got != tt.want.ResourcesWithHighSeverityCounter {
				t.Errorf("SeverityCounters.NumberOfResourcesWithCriticalSeverity() = %v, want %v", got, tt.want)
			}
			if got := sc.NumberOfMediumSeverity(); got != tt.want.ResourcesWithMediumSeverityCounter {
				t.Errorf("SeverityCounters.NumberOfResourcesWithCriticalSeverity() = %v, want %v", got, tt.want)
			}
			if got := sc.NumberOfLowSeverity(); got != tt.want.ResourcesWithLowSeverityCounter {
				t.Errorf("SeverityCounters.NumberOfResourcesWithCriticalSeverity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSeverityCounters_Increase(t *testing.T) {
	type fields struct {
		ResourcesWithCriticalSeverityCounter int
		ResourcesWithHighSeverityCounter     int
		ResourcesWithMediumSeverityCounter   int
		ResourcesWithLowSeverityCounter      int
	}
	type args struct {
		severity string
		amount   int
	}
	tests := []struct {
		name     string
		args     args
		expected int
		fields   fields
	}{
		{
			name: "increase critical severity",
			args: args{
				severity: apis.SeverityCriticalString,
				amount:   1,
			},
			fields: fields{
				ResourcesWithCriticalSeverityCounter: 3,
			},
			expected: 4,
		},
		{
			name: "increase high severity",
			args: args{
				severity: apis.SeverityHighString,
				amount:   2,
			},
			fields: fields{
				ResourcesWithHighSeverityCounter: 2,
			},
			expected: 4,
		},
		{
			name: "increase medium severity",
			args: args{
				severity: apis.SeverityMediumString,
				amount:   3,
			},
			fields: fields{
				ResourcesWithMediumSeverityCounter: 1,
			},
			expected: 4,
		},
		{
			name: "increase low severity",
			args: args{
				severity: apis.SeverityLowString,
				amount:   0,
			},
			fields: fields{
				ResourcesWithLowSeverityCounter: 4,
			},
			expected: 4,
		},
		{
			name: "increase undefine severity",
			args: args{
				severity: "",
				amount:   3,
			},
			fields:   fields{},
			expected: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sc := &SeverityCounters{
				CriticalSeverityCounter: tt.fields.ResourcesWithCriticalSeverityCounter,
				HighSeverityCounter:     tt.fields.ResourcesWithHighSeverityCounter,
				MediumSeverityCounter:   tt.fields.ResourcesWithMediumSeverityCounter,
				LowSeverityCounter:      tt.fields.ResourcesWithLowSeverityCounter,
			}
			sc.Increase(tt.args.severity, tt.args.amount)

			severityCounter := 0
			switch tt.args.severity {
			case apis.SeverityCriticalString:
				severityCounter = sc.CriticalSeverityCounter
			case apis.SeverityHighString:
				severityCounter = sc.HighSeverityCounter
			case apis.SeverityMediumString:
				severityCounter = sc.MediumSeverityCounter
			case apis.SeverityLowString:
				severityCounter = sc.LowSeverityCounter
			}

			if tt.expected != severityCounter {
				t.Errorf("severity: %s, counter = %d, want = %d", tt.args.severity, severityCounter, tt.expected)
			}
		})
	}
}
