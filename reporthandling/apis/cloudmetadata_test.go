package apis

import (
	"testing"
)

func TestCloudProviderName_Compare(t *testing.T) {
	type args struct {
		other string
	}
	tests := []struct {
		name string
		c    CloudProviderName
		args args
		want bool
	}{
		{
			name: "GKE - uppercase",
			c:    GKE,
			args: args{other: "GKE"},
			want: true,
		},
		{
			name: "gke - lowercase",
			c:    GKE,
			args: args{other: "gke"},
			want: true,
		},
		{
			name: "gke/eks - mixed",
			c:    GKE,
			args: args{other: "eks"},
			want: false,
		},
		{
			name: "eks",
			c:    EKS,
			args: args{other: "EKS"},
			want: true,
		},
		{
			name: "eks",
			c:    EKS,
			args: args{other: "EKS"},
			want: true,
		},
		{
			name: "aks",
			c:    AKS,
			args: args{other: "aks"},
			want: true,
		},
		{
			name: "gcp",
			c:    GCP,
			args: args{other: "gcp"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Compare(tt.args.other); got != tt.want {
				t.Errorf("CloudProviderName.Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCloudProviderName_ToString(t *testing.T) {
	tests := []struct {
		name string
		c    CloudProviderName
		want string
	}{
		{
			name: "GKE",
			c:    GKE,
			want: "GKE",
		},
		{
			name: "EKS",
			c:    EKS,
			want: "EKS",
		},
		{
			name: "AKS",
			c:    AKS,
			want: "AKS",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.ToString(); got != tt.want {
				t.Errorf("CloudProviderName.ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
