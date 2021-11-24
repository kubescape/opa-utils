package reporthandling

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	container "cloud.google.com/go/container/apiv1"
	k8sinterface "github.com/armosec/k8s-interface/k8sinterface"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/eks"
	containerpb "google.golang.org/genproto/googleapis/container/v1"
)

func GetClusterInfoForEKS() (map[string]interface{}, error) {
	s, err := session.NewSession()
	if err != nil {
		return nil, err
	}
	currContext := k8sinterface.GetCurrentContext()
	if currContext == nil {
		return nil, nil
	}
	region := strings.Split(k8sinterface.GetCurrentContext().Cluster, ".")[1]
	svc := eks.New(s, &aws.Config{Region: aws.String(region)})
	input := &eks.DescribeClusterInput{
		Name: aws.String(k8sinterface.GetCurrentContext().Cluster),
	}

	result, err := svc.DescribeCluster(input)
	if err != nil {
		return nil, err
	}
	resultInJson, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}
	var clusterInfo map[string]interface{}
	err = json.Unmarshal(resultInJson, &clusterInfo)
	if err != nil {
		return nil, err
	}
	return clusterInfo, nil
}

func GetClusterInfoForGKE() (map[string]interface{}, error) {
	ctx := context.Background()
	c, err := container.NewClusterManagerClient(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	parsedName := strings.Split(k8sinterface.GetClusterName(), "_")
	clusterName := fmt.Sprintf("projects/%s/locations/%s/clusters/%s", parsedName[1], parsedName[2], parsedName[3])
	req := &containerpb.GetClusterRequest{
		Name: clusterName,
	}
	result, err := c.GetCluster(ctx, req)
	if err != nil {
		return nil, err
	}
	resultInJson, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}
	var clusterInfo map[string]interface{}
	err = json.Unmarshal(resultInJson, &clusterInfo)
	if err != nil {
		return nil, err
	}
	return clusterInfo, nil

}
