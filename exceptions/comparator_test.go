package exceptions

import (
	"testing"

	"github.com/kubescape/k8s-interface/workloadinterface"
	"github.com/kubescape/opa-utils/objectsenvelopes"
	"github.com/kubescape/opa-utils/objectsenvelopes/localworkload"

	"github.com/stretchr/testify/assert"
)

func TestComparator_compareNamespace(t *testing.T) {
	c := &comparator{} // Assuming you have initialized the comparator instance

	tests := []struct {
		name      string
		workload  workloadinterface.IMetadata
		namespace string
		expected  bool
	}{
		{
			name: "Workload kind is Namespace, regex match",
			workload: workloadinterface.NewWorkloadObj(
				map[string]interface{}{
					"apiVersion": "v1",
					"kind":       "Namespace",
					"metadata": map[string]interface{}{
						"name": "test-namespace",
					},
				},
			),
			namespace: "test-namespace",
			expected:  true,
		},
		{
			name: "Workload kind is Namespace, regex does not match",
			workload: workloadinterface.NewWorkloadObj(
				map[string]interface{}{
					"apiVersion": "v1",
					"kind":       "Namespace",
					"metadata": map[string]interface{}{
						"name": "test-namespace",
					},
				},
			),
			namespace: "different-namespace",
			expected:  false,
		},
		{
			name: "Workload kind is not Namespace, regex match",
			workload: workloadinterface.NewWorkloadObj(
				map[string]interface{}{
					"apiVersion": "v1",
					"kind":       "Deployment",
					"metadata": map[string]interface{}{
						"name":      "test",
						"namespace": "test-namespace",
					},
				},
			),
			namespace: "test-namespace",
			expected:  true,
		},
		{
			name: "Workload kind is not Namespace, regex does not match",
			workload: workloadinterface.NewWorkloadObj(
				map[string]interface{}{
					"apiVersion": "v1",
					"kind":       "Deployment",
					"metadata": map[string]interface{}{
						"name":      "test",
						"namespace": "different-namespace",
					},
				},
			),
			namespace: "test-namespace",
			expected:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, c.compareNamespace(tt.workload, tt.namespace))
		})
	}
}
func TestComparator_compareKind(t *testing.T) {
	c := &comparator{} // Assuming you have initialized the comparator instance

	tests := []struct {
		name     string
		workload workloadinterface.IMetadata
		kind     string
		expected bool
	}{
		{
			name: "Kind matches",
			workload: workloadinterface.NewWorkloadObj(
				map[string]interface{}{
					"apiVersion": "v1",
					"kind":       "Deployment",
					"metadata": map[string]interface{}{
						"name":      "test",
						"namespace": "test-namespace",
					},
				},
			),
			kind:     "Deployment",
			expected: true,
		},
		{
			name: "Kind does not match",
			workload: workloadinterface.NewWorkloadObj(
				map[string]interface{}{
					"apiVersion": "v1",
					"kind":       "Deployment",
					"metadata": map[string]interface{}{
						"name":      "test",
						"namespace": "test-namespace",
					},
				},
			),
			kind:     "Pod",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, c.compareKind(tt.workload, tt.kind))
		})
	}
}

func TestComparator_compareName(t *testing.T) {
	c := &comparator{} // Assuming you have initialized the comparator instance

	tests := []struct {
		name         string
		workload     workloadinterface.IMetadata
		workloadName string
		expected     bool
	}{
		{
			name: "Name matches",
			workload: workloadinterface.NewWorkloadObj(
				map[string]interface{}{
					"apiVersion": "v1",
					"kind":       "Deployment",
					"metadata": map[string]interface{}{
						"name":      "test",
						"namespace": "test-namespace",
					},
				},
			),
			workloadName: "test",
			expected:     true,
		},
		{
			name: "Name does not match",
			workload: workloadinterface.NewWorkloadObj(
				map[string]interface{}{
					"apiVersion": "v1",
					"kind":       "Deployment",
					"metadata": map[string]interface{}{
						"name":      "test",
						"namespace": "test-namespace",
					},
				},
			),
			workloadName: "different",
			expected:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, c.compareName(tt.workload, tt.workloadName))
		})
	}
}
func TestComparator_comparePath(t *testing.T) {
	c := &comparator{} // Assuming you have initialized the comparator instance

	tests := []struct {
		name     string
		workload workloadinterface.IMetadata
		path     string
		expected bool
	}{
		{
			name: "Workload has sourcePath, regex match",
			workload: workloadinterface.NewWorkloadObj(
				map[string]interface{}{
					"apiVersion": "v1",
					"kind":       "Deployment",
					"metadata": map[string]interface{}{
						"name":      "test",
						"namespace": "test-namespace",
					},
					"sourcePath": "/test/path",
				},
			),
			path:     "/test/path",
			expected: true,
		},
		{
			name: "Workload has sourcePath, regex does not match",
			workload: workloadinterface.NewWorkloadObj(
				map[string]interface{}{
					"apiVersion": "v1",
					"kind":       "Deployment",
					"metadata": map[string]interface{}{
						"name":      "test",
						"namespace": "test-namespace",
					},
					"sourcePath": "/test/path",
				},
			),
			path:     "/different/path",
			expected: false,
		},
		{
			name: "Workload does not have sourcePath",
			workload: workloadinterface.NewWorkloadObj(
				map[string]interface{}{
					"apiVersion": "v1",
					"kind":       "Deployment",
					"metadata": map[string]interface{}{
						"name":      "test",
						"namespace": "test-namespace",
					},
				},
			),
			path:     "/test/path",
			expected: false,
		},
		{
			name: "Not a workload",
			workload: workloadinterface.NewWorkloadObj(
				map[string]interface{}{
					"kind": "Something",
					"name": "test",
				},
			),
			path:     "/test/path",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, c.comparePath(tt.workload, tt.path))
		})
	}
}

func TestIsTypeWorkload(t *testing.T) {
	tests := []struct {
		workload workloadinterface.IMetadata
		name     string
		expected bool
	}{
		{
			name: "Workload type is WorkloadObject",
			workload: workloadinterface.NewWorkloadObj(
				map[string]interface{}{
					"apiVersion": "apps/v1",
					"kind":       "Deployment",
					"metadata": map[string]interface{}{
						"name":      "test",
						"namespace": "test-namespace",
					},
				},
			),
			expected: true,
		},
		{
			name: "Workload type is ListWorkloads",
			workload: workloadinterface.NewListWorkloadsObj(
				map[string]interface{}{
					"apiVersion": "v1",
					"kind":       "List",
					"metadata": map[string]interface{}{
						"namespace": "test-namespace",
					},
					"items": []interface{}{},
				},
			),
			expected: true,
		},
		{
			name: "Workload type is BaseObject",
			workload: workloadinterface.NewBaseObject(
				map[string]interface{}{
					"apiVersion": "v1",
					"kind":       "Namespace",
					"metadata": map[string]interface{}{
						"name": "test",
					},
				},
			),
			expected: true,
		},
		{
			name: "Workload type is LocalWorkload",
			workload: localworkload.NewLocalWorkload(
				map[string]interface{}{
					"apiVersion": "v1",
					"kind":       "Deployment",
					"metadata": map[string]interface{}{
						"name":      "test",
						"namespace": "test-namespace",
					},
					"sourcePath": "/test/path",
				},
			),
			expected: true,
		},
		{
			name: "Workload type is not a recognized type",
			workload: workloadinterface.NewWorkloadObj(
				map[string]interface{}{
					"objectType": "UnrecognizedType",
				},
			),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, isTypeWorkload(tt.workload))
		})
	}
}

func TestIsTypeRegoResponseVector(t *testing.T) {
	tests := []struct {
		workload workloadinterface.IMetadata
		name     string
		expected bool
	}{
		{
			name: "Workload type is RegoResponseVectorObject",
			workload: objectsenvelopes.NewRegoResponseVectorObject(
				map[string]interface{}{
					"kind": "RegoResponseVector",
					"name": "test",
					"relatedObjects": []interface{}{
						map[string]interface{}{
							"apiVersion": "apps/v1",
							"kind":       "Deployment",
							"metadata": map[string]interface{}{
								"name":      "test",
								"namespace": "test-namespace",
							},
						},
						map[string]interface{}{
							"apiVersion": "apps/v1",
							"kind":       "Deployment",
							"metadata": map[string]interface{}{
								"name":      "test-2",
								"namespace": "test-namespace",
							},
						},
					},
				},
			),
			expected: true,
		},
		{
			name: "Workload type is not RegoResponseVectorObject",
			workload: workloadinterface.NewBaseObject(
				map[string]interface{}{
					"apiVersion": "v1",
					"kind":       "Namespace",
					"metadata": map[string]interface{}{
						"name": "test",
					},
				},
			),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, isTypeRegoResponseVector(tt.workload))
		})
	}
}
