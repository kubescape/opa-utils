package helpers

import (
	"fmt"
	"strings"

	"github.com/kubescape/opa-utils/reporthandling/apis"
)

// Support the different cloud providers naming conventions.
// For example, in GKE, the cluster name is in the format of "gke_project_zone_cluster-name", and in AKS, the cluster name is in the format of "arn:aws:eks:eu-west-1:id:cluster/my-cluster"
// Read more about the different cloud providers naming conventions here: https://stackoverflow.com/questions/74516648/eks-aks-cluster-name-convention/74534378#74534378
//
// The following functions are used to parse the cluster name and return the prefix and suffix of the name.
//
// The developer using this package should implement a function that looks like this:
//
// func GetCloudMetadata() apis.ICloudParser {
// 	var name string
// 	var provider apis.CloudProviderName
// 	switch provider {
// 	case apis.GCP, apis.GKE:
// 		return NewGKEMetadata(name)
// 	case apis.AKS:
// 		return NewAKSMetadata() // TODO: pass the relevant parameters
// 	case apis.EKS:
// 		return NewEKSMetadata(name)
// 	}
// 	return nil
// }
//

// =============================== GCP ===============================

type GKEMetadata struct {
	name string
}

func NewGKEMetadata(name string) *GKEMetadata {
	return &GKEMetadata{name: name}
}

// GetName returns the full name of the cluster
func (gke *GKEMetadata) GetName() string {
	return gke.name
}

func (gke GKEMetadata) Provider() apis.CloudProviderName {
	return apis.GKE
}

func (gke *GKEMetadata) Parse() (string, string, error) {
	sliced := gke.split(gke.name)
	if len(sliced) < 4 {
		return "", "", fmt.Errorf("cluster name '%s' is not a valid GCP cluster name", gke.name)
	}
	return gke.join(sliced[:3]), gke.join(sliced[3:]), nil
}

func (gke GKEMetadata) split(s string) []string {
	return strings.Split(s, "_")
}

func (gke GKEMetadata) join(s []string) string {
	return strings.Join(s, "_")
}

// =============================== EKS ===============================

type EKSMetadata struct {
	name string
}

func NewEKSMetadata(name string) *EKSMetadata {
	return &EKSMetadata{name: name}
}

// GetName returns the full name of the cluster
func (eks *EKSMetadata) GetName() string {
	return eks.name
}

func (eks EKSMetadata) Provider() apis.CloudProviderName {
	return apis.EKS
}

func (eks *EKSMetadata) Parse() (string, string, error) {
	prefix, suffix, err := eks.parseKubernetesContextName()
	if err == nil {
		return prefix, suffix, nil
	}
	return eks.parseNormalizedClusterName()
}

// parseKubernetesContextName parses a standard ARN format EKS cluster name from Kubernetes context name format
// Format: "arn:aws:eks:{region}:{account}:cluster/{cluster-name}"
func (eks *EKSMetadata) parseKubernetesContextName() (string, string, error) {
	sliced := eks.split(eks.name)
	if len(sliced) < 6 {
		return "", "", fmt.Errorf("cluster name '%s' is not a valid EKS cluster name", eks.name)
	}

	return eks.join(sliced[:5]), strings.Replace(eks.join(sliced[5:]), "cluster/", "", 1), nil
}

func (eks EKSMetadata) split(s string) []string {
	return strings.Split(s, ":")
}

func (eks EKSMetadata) join(s []string) string {
	return strings.Join(s, ":")
}

// parseNormalizedClusterName parses a normalized EKS cluster name
// Format: "arn-aws-eks-{region}-{account}-cluster-{cluster-name}"
func (eks *EKSMetadata) parseNormalizedClusterName() (string, string, error) {
	// Remove "arn-aws-eks-" prefix
	withoutPrefix := strings.TrimPrefix(eks.name, "arn-aws-eks-")

	// Find the "-cluster-" separator
	clusterIndex := strings.Index(withoutPrefix, "-cluster-")
	if clusterIndex == -1 {
		return "", "", fmt.Errorf("cluster name '%s' is not a valid normalized EKS cluster name", eks.name)
	}

	// Extract region and account (everything before "-cluster-")
	regionAndAccount := withoutPrefix[:clusterIndex]

	// Extract cluster name (everything after "-cluster-")
	clusterName := withoutPrefix[clusterIndex+len("-cluster-"):]

	// Split region and account by finding the last hyphen (account is typically numeric)
	// We need to find where the region ends and account begins
	// Region format can be like "eu-central-1" or "us-east-1", account is numeric
	parts := strings.Split(regionAndAccount, "-")
	if len(parts) < 2 {
		return "", "", fmt.Errorf("cluster name '%s' is not a valid normalized EKS cluster name", eks.name)
	}

	// Account is the last part, region is everything before it
	account := parts[len(parts)-1]
	region := strings.Join(parts[:len(parts)-1], "-")

	// Build prefix: "arn:aws:eks:{region}:{account}"
	prefix := fmt.Sprintf("arn:aws:eks:%s:%s", region, account)

	return prefix, clusterName, nil
}

// =============================== AKS ===============================
// TODO: implement AKS parser support

type AKSMetadata struct {
	name string
}

func NewAKSMetadata(name string) *AKSMetadata {
	return &AKSMetadata{
		name: name,
	}
}

// GetName returns the full name of the cluster
func (aks *AKSMetadata) GetName() string {
	return aks.name
}

func (aks AKSMetadata) Provider() apis.CloudProviderName {
	return apis.AKS
}

func (aks *AKSMetadata) Parse() (string, string, error) {
	return "", aks.name, nil
}
