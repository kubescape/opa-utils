package apis

import "strings"

type CloudProviderName string

const (
	GCP CloudProviderName = "GCP" // TODO: Why do we have this name ??
	GKE CloudProviderName = "GKE"
	EKS CloudProviderName = "EKS"
	AKS CloudProviderName = "AKS"
)

type ICloudParser interface {

	// Get the full and original name of the cluster. For example, for a cluster named "gke_project_zone_my-cluster" in GKE, the name is "gke_project_zone_my-cluster"
	GetName() string

	// Get provider name
	Provider() CloudProviderName

	// Parse the name of the cluster and return the prefix, suffix. For example, for a cluster named "gke_project_zone_my-cluster" in GKE, the prefix is "gke_project_zone" and the suffix is "my-cluster"
	Parse() (string, string, error)

	// // Parse the name of the cluster and return the prefix. For example, for a cluster named "gke_project_zone_my-cluster" in GKE, the prefix is "gke_project_zone"
	// Prefix() (string, error)

	// // Parse the name of the cluster and return the suffix. For example, for a cluster named "gke_project_zone_my-cluster" in GKE, the suffix is "my-cluster"
	// Suffix() (string, error)

	// // Split the name of the cluster. For example, for a cluster named "gke_project_zone_my-cluster" in GKE, the function should return ["gke","project","zone","my-cluster"]
	// Split(string) ([]string, error)

	// // Join the different elements of the cluster name. For example, for a cluster named ["gke","project","zone","my-cluster"] in GKE, the function will return "gke_project_zone_my-cluster"
	// Join([]string) (string, error)
}

// Compare returns true if the given string is equal to the cloud provider name
func (c CloudProviderName) Compare(other string) bool {
	return strings.ToUpper(other) == string(c)
}

// Convert the cloud provider name to string
func (c CloudProviderName) ToString() string {
	return string(c)
}
