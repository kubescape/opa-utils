package v1

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
func (gke GKEMetadata) GetName() string {
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
func (eks EKSMetadata) GetName() string {
	return eks.name
}

func (eks EKSMetadata) Provider() apis.CloudProviderName {
	return apis.EKS
}

func (eks *EKSMetadata) Parse() (string, string, error) {

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

// =============================== AKS ===============================
// TODO: implement AKS parser support

type AKSMetadata struct {
	// name string
}

func NewAKSMetadata() *AKSMetadata {
	return &AKSMetadata{}
}

// GetName returns the full name of the cluster
func (aks AKSMetadata) GetName() string {
	return ""
}

func (aks AKSMetadata) Provider() apis.CloudProviderName {
	return apis.AKS
}

func (aks *AKSMetadata) Parse() (string, string, error) {
	return "", "", nil
}
