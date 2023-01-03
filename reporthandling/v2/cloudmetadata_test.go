package v2

import (
	"reflect"
	"testing"

	"github.com/kubescape/opa-utils/reporthandling/apis"
	v1 "github.com/kubescape/opa-utils/reporthandling/helpers/v1"
)

func TestNewCloudMetadata(t *testing.T) {
	type args struct {
		cloudParser apis.ICloudParser
	}
	tests := []struct {
		args args
		want *CloudMetadata
		name string
	}{
		{
			name: "TestNewCloudMetadata - EKS",
			args: args{
				cloudParser: v1.NewEKSMetadata("arn:aws:eks:eu-west-1:id:cluster/my-cluster"),
			},
			want: &CloudMetadata{
				CloudProvider: apis.EKS,
				FullName:      "arn:aws:eks:eu-west-1:id:cluster/my-cluster",
				ShortName:     "my-cluster",
				PrefixName:    "arn:aws:eks:eu-west-1:id",
			},
		},
		{
			name: "TestNewCloudMetadata - GKE",
			args: args{
				cloudParser: v1.NewGKEMetadata("gke_project_zone_my-cluster"),
			},
			want: &CloudMetadata{
				CloudProvider: apis.GKE,
				FullName:      "gke_project_zone_my-cluster",
				ShortName:     "my-cluster",
				PrefixName:    "gke_project_zone",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCloudMetadata(tt.args.cloudParser); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCloudMetadata() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCloudMetadata_GetName(t *testing.T) {
	type fields struct {
		ShortName string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "TestCloudMetadata_GetName",
			fields: fields{
				ShortName: "name",
			},
			want: "name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cloudMetadata := &CloudMetadata{
				ShortName: tt.fields.ShortName,
			}
			if got := cloudMetadata.GetName(); got != tt.want {
				t.Errorf("CloudMetadata.GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCloudMetadata_GetFullName(t *testing.T) {
	type fields struct {
		FullName string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "TestCloudMetadata_GetFullName",
			fields: fields{
				FullName: "FullName",
			},
			want: "FullName",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cloudMetadata := &CloudMetadata{
				FullName: tt.fields.FullName,
			}
			if got := cloudMetadata.GetFullName(); got != tt.want {
				t.Errorf("CloudMetadata.GetFullName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCloudMetadata_GetProvider(t *testing.T) {
	type fields struct {
		CloudProvider apis.CloudProviderName
	}
	tests := []struct {
		name   string
		fields fields
		want   apis.CloudProviderName
	}{
		{
			name: "TestCloudMetadata_GetProvider",
			fields: fields{
				CloudProvider: apis.EKS,
			},
			want: apis.EKS,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cloudMetadata := &CloudMetadata{
				CloudProvider: tt.fields.CloudProvider,
			}
			if got := cloudMetadata.GetProvider(); got != tt.want {
				t.Errorf("CloudMetadata.GetProvider() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCloudMetadata_GetPrefix(t *testing.T) {
	type fields struct {
		PrefixName string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "TestCloudMetadata_GetPrefix",
			fields: fields{
				PrefixName: "PrefixName",
			},
			want: "PrefixName",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cloudMetadata := &CloudMetadata{
				PrefixName: tt.fields.PrefixName,
			}
			if got := cloudMetadata.GetPrefix(); got != tt.want {
				t.Errorf("CloudMetadata.GetPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
