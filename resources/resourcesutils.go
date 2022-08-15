package resources

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/open-policy-agent/opa/storage"
	"github.com/open-policy-agent/opa/storage/inmem"

	k8sinterface "github.com/kubescape/k8s-interface/k8sinterface"
	"k8s.io/client-go/rest"
)

type RegoDependenciesData struct {
	ClusterName          string              `json:"clusterName"`
	PostureControlInputs map[string][]string `json:"postureControlInputs"`
	K8sConfig            RegoK8sConfig       `json:"k8sconfig"`
}

type RegoK8sConfig struct {
	Token         string `json:"token"`
	IP            string `json:"ip"`
	Host          string `json:"host"`
	Port          string `json:"port"`
	CrtFile       string `json:"crtfile"`
	ClientCrtFile string `json:"clientcrtfile"`
	ClientKeyFile string `json:"clientkeyfile"`
	// ClientKeyFile string `json:"crtfile"`
}

func NewRegoDependenciesDataMock() *RegoDependenciesData {
	return NewRegoDependenciesData(k8sinterface.GetK8sConfig(), "")
}

func NewRegoDependenciesData(k8sConfig *rest.Config, clusterName string) *RegoDependenciesData {

	regoDependenciesData := RegoDependenciesData{
		ClusterName: clusterName,
	}
	if k8sConfig != nil {
		regoDependenciesData.K8sConfig = *NewRegoK8sConfig(k8sConfig)
	}
	return &regoDependenciesData
}
func NewRegoK8sConfig(k8sConfig *rest.Config) *RegoK8sConfig {

	host := k8sConfig.Host
	if host == "" {
		ip := os.Getenv("KUBERNETES_SERVICE_HOST")
		port := os.Getenv("KUBERNETES_SERVICE_PORT")
		host = fmt.Sprintf("https://%s:%s", ip, port)
	}

	token := ""
	if k8sConfig.BearerToken != "" {
		token = fmt.Sprintf("Bearer %s", k8sConfig.BearerToken)
	}

	regoK8sConfig := RegoK8sConfig{
		Token:         token,
		Host:          host,
		CrtFile:       k8sConfig.CAFile,
		ClientCrtFile: k8sConfig.CertFile,
		ClientKeyFile: k8sConfig.KeyFile,
	}
	return &regoK8sConfig
}

func (data *RegoDependenciesData) GetFilteredPostureControlInputs(settings []string) map[string][]string {
	jsonObj := map[string][]string{}
	for i := range settings {
		splitted := strings.Split(settings[i], ".")
		if len(splitted) != 3 {
			continue
		}
		if v, k := data.PostureControlInputs[splitted[2]]; k && v != nil {
			jsonObj[splitted[2]] = v
		}
	}
	return jsonObj
}
func TOStorage(postureControlInputs map[string][]string) (storage.Store, error) {
	return inmem.NewFromObject(map[string]interface{}{"postureControlInputs": postureControlInputs}), nil
}

// LoadRegoDependenciesFromDir loads the policies list from *.rego file in given directory
func LoadRegoFiles(dir string) map[string]string {

	modules := make(map[string]string)

	// Compile the module. The keys are used as identifiers in error messages.
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err == nil && strings.HasSuffix(path, ".rego") && !info.IsDir() {
			content, err := os.ReadFile(path)
			if err != nil {
				fmt.Printf("LoadRegoFiles, Failed to load: %s: %v", path, err)
			} else {
				modules[strings.Trim(filepath.Base(path), ".rego")] = string(content)
			}
		}
		return nil
	})

	return modules
}

// LoadRegoModules loads the policies from variables
func LoadRegoModules() map[string]string {

	modules := make(map[string]string)
	modules["cautils"] = RegoCAUtils
	modules["designators"] = RegoDesignators
	modules["kubernetes.api.client"] = RegoKubernetesApiClient

	return modules
}
