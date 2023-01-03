package v1

import (
	"reflect"
	"testing"

	"github.com/kubescape/opa-utils/reporthandling/apis"
)

const (
	gkeName = "gke_project_zone_my-cluster"
	eksName = "arn:aws:eks:eu-west-1:id:cluster/my-cluster"
)

func TestNewGKEMetadata(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		want *GKEMetadata
		name string
		args args
	}{
		{
			name: "TestNewGCPMetadata",
			args: args{name: gkeName},
			want: &GKEMetadata{name: gkeName},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGKEMetadata(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGCPMetadata() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGKEMetadata_GetName(t *testing.T) {
	type fields struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "TestGKEMetadata_GetName",
			fields: fields{name: gkeName},
			want:   gkeName,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gke := GKEMetadata{
				name: tt.fields.name,
			}
			if got := gke.GetName(); got != tt.want {
				t.Errorf("GKEMetadata.GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGKEMetadata_Provider(t *testing.T) {
	type fields struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		want   apis.CloudProviderName
	}{
		{
			name:   "TestGKEMetadata_Provider",
			fields: fields{name: gkeName},
			want:   apis.GKE,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gke := GKEMetadata{
				name: tt.fields.name,
			}
			if got := gke.Provider(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GKEMetadata.Provider() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGKEMetadata_Parse(t *testing.T) {
	type fields struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		want1   string
		wantErr bool
	}{
		{
			name:    "TestGKEMetadata_Parse",
			fields:  fields{name: gkeName},
			want:    "gke_project_zone",
			want1:   "my-cluster",
			wantErr: false,
		},
		{
			name:    "TestGKEMetadata_Parse - wrong format",
			fields:  fields{name: "gke_project_zone"},
			want:    "",
			want1:   "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gke := &GKEMetadata{
				name: tt.fields.name,
			}
			got, got1, err := gke.Parse()
			if (err != nil) != tt.wantErr {
				t.Errorf("GKEMetadata.Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GKEMetadata.Parse() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GKEMetadata.Parse() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNewEKSMetadata(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		want *EKSMetadata
		name string
		args args
	}{
		{
			name: "TestNewEKSMetadata",
			args: args{name: eksName},
			want: &EKSMetadata{name: eksName},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEKSMetadata(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEKSMetadata() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEKSMetadata_GetName(t *testing.T) {
	type fields struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "TestEKSMetadata_GetName",
			fields: fields{name: eksName},
			want:   eksName,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			eks := EKSMetadata{
				name: tt.fields.name,
			}
			if got := eks.GetName(); got != tt.want {
				t.Errorf("EKSMetadata.GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEKSMetadata_Provider(t *testing.T) {
	type fields struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		want   apis.CloudProviderName
	}{
		{
			name:   "TestEKSMetadata_Provider",
			fields: fields{name: eksName},
			want:   apis.EKS,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			eks := EKSMetadata{
				name: tt.fields.name,
			}
			if got := eks.Provider(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EKSMetadata.Provider() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEKSMetadata_Parse(t *testing.T) {
	type fields struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		want1   string
		wantErr bool
	}{
		{
			name:    "TestEKSMetadata_Parse",
			fields:  fields{name: eksName},
			want:    "arn:aws:eks:eu-west-1:id",
			want1:   "my-cluster",
			wantErr: false,
		},
		{
			name:    "TestEKSMetadata_Parse - wrong format",
			fields:  fields{name: "arn:aws:eks:eu-west-1:cluster/my-cluster"},
			want:    "",
			want1:   "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			eks := &EKSMetadata{
				name: tt.fields.name,
			}
			got, got1, err := eks.Parse()
			if (err != nil) != tt.wantErr {
				t.Errorf("EKSMetadata.Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("EKSMetadata.Parse() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("EKSMetadata.Parse() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNewAKSMetadata(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		want *AKSMetadata
		name string
		args args
	}{
		{
			name: "TestNewAKSMetadata",
			args: args{name: "aksName"},
			want: &AKSMetadata{name: "aksName"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAKSMetadata(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAKSMetadata() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAKSMetadata_GetName(t *testing.T) {
	type fields struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "TestAKSMetadata_GetName",
			fields: fields{name: "aksName"},
			want:   "aksName",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			aks := &AKSMetadata{
				name: tt.fields.name,
			}
			if got := aks.GetName(); got != tt.want {
				t.Errorf("AKSMetadata.GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAKSMetadata_Provider(t *testing.T) {
	type fields struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		want   apis.CloudProviderName
	}{
		{
			name:   "TestAKSMetadata_Provider",
			fields: fields{name: ""},
			want:   apis.AKS,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			aks := AKSMetadata{
				name: tt.fields.name,
			}
			if got := aks.Provider(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AKSMetadata.Provider() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAKSMetadata_Parse(t *testing.T) {
	type fields struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		want1   string
		wantErr bool
	}{
		{
			name:   "TestAKSMetadata_Parse",
			fields: fields{name: "aksName"},
			want:   "",
			want1:  "aksName",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			aks := &AKSMetadata{
				name: tt.fields.name,
			}
			got, got1, err := aks.Parse()
			if (err != nil) != tt.wantErr {
				t.Errorf("AKSMetadata.Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AKSMetadata.Parse() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("AKSMetadata.Parse() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
