package reportsummary

import "testing"

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
				ResourcesWithCriticalSeverityCounter: tt.fields.ResourcesWithCriticalSeverityCounter,
				ResourcesWithHighSeverityCounter:     tt.fields.ResourcesWithHighSeverityCounter,
				ResourcesWithMediumSeverityCounter:   tt.fields.ResourcesWithMediumSeverityCounter,
				ResourcesWithLowSeverityCounter:      tt.fields.ResourcesWithLowSeverityCounter,
			}
			if got := sc.NumberOfResourcesWithCriticalSeverity(); got != tt.want.ResourcesWithCriticalSeverityCounter {
				t.Errorf("SeverityCounters.NumberOfResourcesWithCriticalSeverity() = %v, want %v", got, tt.want)
			}
			if got := sc.NumberOfResourcesWithHighSeverity(); got != tt.want.ResourcesWithHighSeverityCounter {
				t.Errorf("SeverityCounters.NumberOfResourcesWithCriticalSeverity() = %v, want %v", got, tt.want)
			}
			if got := sc.NumberOfResourcesWithMediumSeverity(); got != tt.want.ResourcesWithMediumSeverityCounter {
				t.Errorf("SeverityCounters.NumberOfResourcesWithCriticalSeverity() = %v, want %v", got, tt.want)
			}
			if got := sc.NumberOfResourcesWithLowSeverity(); got != tt.want.ResourcesWithLowSeverityCounter {
				t.Errorf("SeverityCounters.NumberOfResourcesWithCriticalSeverity() = %v, want %v", got, tt.want)
			}
		})
	}
}
