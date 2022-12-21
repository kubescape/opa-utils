package v2

import (
	"github.com/kubescape/opa-utils/reporthandling/apis"
)

// =============================== CloudMetadata ===============================

// NewCloudMetadata creates a new CloudMetadata object
func NewCloudMetadata(cloudParser apis.ICloudParser) *CloudMetadata {
	prefix, suffix, _ := cloudParser.Parse()
	return &CloudMetadata{
		CloudProvider: cloudParser.Provider(),
		FullName:      cloudParser.GetName(),
		ShortName:     suffix,
		PrefixName:    prefix,
	}
}

// Get the suffix of the name. For example, for a cluster named "gke_project_zone_my-cluster" in GKE, the suffix is "my-cluster"
func (cloudMetadata *CloudMetadata) GetName() string {
	return cloudMetadata.ShortName
}

// Get full name as it appears in the config file
func (cloudMetadata *CloudMetadata) GetFullName() string {
	return cloudMetadata.FullName
}

// Get provider name
func (cloudMetadata *CloudMetadata) GetProvider() apis.CloudProviderName {
	return cloudMetadata.CloudProvider
}

// Get the prefix of the name. For example, for a cluster named "gke_project_zone_my-cluster" in GKE, the prefix is "gke_project_zone"
func (cloudMetadata *CloudMetadata) GetPrefix() string {
	return cloudMetadata.PrefixName
}
