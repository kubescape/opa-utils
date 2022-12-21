package reporthandling

type ICloudMetadata interface {
	// Get the suffix of the name. For example, for a cluster named "gke_project_zone_my-cluster" in GKE, the suffix is "my-cluster"
	GetName() string

	// Get full name as it appears in the config file
	GetFullName() string

	// Get provider name
	GetProvider() string

	// Get the prefix of the name. For example, for a cluster named "gke_project_zone_my-cluster" in GKE, the prefix is "gke_project_zone"
	GetPrefix() string
}
