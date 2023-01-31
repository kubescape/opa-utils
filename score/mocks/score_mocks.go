package mocks

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/kubescape/k8s-interface/workloadinterface"
	"github.com/stretchr/testify/require"
)

type Option func(map[string]interface{})

// WithName overrides the name of a mock resource
func WithName(name string) Option {
	return func(workload map[string]interface{}) {
		tmp, ok := workload["metadata"].(map[string]interface{})
		if !ok {
			return
		}
		tmp["name"] = name
		workload["metadata"] = tmp
	}
}

// WithID overrides the ID of a mock resource
func WithID(name string) Option {
	return func(workload map[string]interface{}) {
		tmp, ok := workload["metadata"].(map[string]interface{})
		if !ok {
			return
		}
		tmp["uid"] = name
		workload["metadata"] = tmp
	}
}

// WithReplicas overrides the replicas in the mock "spec" section, if any.
func WithReplicas(replicas int) Option {
	return func(workload map[string]interface{}) {
		tmp, ok := workload["spec"].(map[string]interface{})
		if !ok {
			return
		}
		tmp["replicas"] = 1
		workload["spec"] = tmp
	}
}

// WithDesiredNumberScheduled overrides the desired number scheduled in the status section, if any.
func WithDesiredNumberScheduled(desired int32) Option {
	return func(workload map[string]interface{}) {
		tmp, ok := workload["status"].(map[string]interface{})
		if !ok {
			return
		}

		tmp["desiredNumberScheduled"] = desired
		workload["status"] = tmp
	}
}

// WithRelatedObjects adds related objects to the mocked object.
func WithRelatedObjects(related ...map[string]interface{}) Option {
	return func(workload map[string]interface{}) {
		tmp, ok := workload["relatedObjects"].([]interface{})
		if !ok {
			tmp = make([]interface{}, len(related))
		}

		for _, add := range related {
			tmp = append(tmp, add)
		}

		workload["relatedObjects"] = tmp
	}
}

// GetResourceByType fetches a resource of the desired type from the resources mock fixture.
//
// For the moment, only 1 mock is available for each resource type.
func GetResourceByType(t testing.TB, desiredType string, opts ...Option) map[string]interface{} {
	for _, v := range loadResourcesMock(t) {
		wl := workloadinterface.NewWorkloadObj(v)
		if wl != nil { // workload fixture
			if strings.EqualFold(wl.GetKind(), desiredType) {
				for _, apply := range opts {
					apply(v)
				}

				return v
			}

			continue
		}

		// other resource
		for k := range v {
			if strings.EqualFold(k, desiredType) {
				for _, apply := range opts {
					apply(v)
				}

				return v
			}
		}
	}

	t.Fatalf("resource type %q not found in mock", desiredType)

	return nil
}

// GetInvalidResources fetches examples of invalid resource definitions.
func GetInvalidResources(t testing.TB) []map[string]interface{} {
	return loadInvalidResourcesMock(t)
}

func loadResourcesMock(t testing.TB) []map[string]interface{} {
	resources := make([]map[string]interface{}, 0)
	require.NoError(t, json.Unmarshal(loadFixture(t, "resourcemocks"), &resources))

	return resources
}

func loadInvalidResourcesMock(t testing.TB) []map[string]interface{} {
	resources := make([]map[string]interface{}, 0)
	require.NoError(t, json.Unmarshal(loadFixture(t, "invalidresources"), &resources))

	return resources
}

func loadFixture(t testing.TB, name string) []byte {
	dat, err := os.ReadFile(filepath.Join(currentDir(t), fmt.Sprintf("%s.json", name)))
	require.NoError(t, err)

	return dat
}

func currentDir(t testing.TB) string {
	_, filename, _, ok := runtime.Caller(1)
	require.Truef(t, ok, "could not resolve current folder")

	return filepath.Dir(filename)
}

/* unused for now
func getMITREFrameworkResultMock() []reporthandling.FrameworkReport {
 	l := make([]reporthandling.FrameworkReport, 0)
 	report := loadFrameworkMock()
 	resources := loadResourcesMock()
 	if report != nil && resources != nil {

 		report.ControlReports[0].RuleReports[0].GetNumberOfResources() = resources
 		l = append(l, *report)
 	}

 	return l
 }

func loadFrameworkMock(t testing.TB) *reporthandling.FrameworkReport {
	report := &reporthandling.FrameworkReport{}

	require.NoError(t, json.Unmarshal(loadFixture(t, "frameworkmock"), &report))

	return report
}
*/
