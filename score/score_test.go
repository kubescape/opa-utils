package score

import (
	"testing"
)

func TestFrameworkMock(t *testing.T) {
}

func TestDaemonsetRule(t *testing.T) {
	// desiredType := "daemonset"
	// r := getResouceByType(desiredType)
	// if r == nil {
	// 	t.Errorf("no %v was found in the mock, should be 1", desiredType)
	// }
	// su := NewScore(nil, "")

	// resources := []map[string]interface{}{r}
	// weights := su.resourceRules(resources)
	// expecting := 13 * su.ResourceTypeScores[desiredType]
	// if weights != expecting {
	// 	t.Errorf("no %v unexpected weights were calculated expecting: %v got %v", desiredType, expecting, weights)
	// }
}

func TestMultipleReplicasRule(t *testing.T) {
	// desiredType := "deployment"
	// r := getResouceByType(desiredType)
	// if r == nil {
	// 	t.Errorf("no %v was found in the mock, should be 1", desiredType)
	// }
	// su := NewScore(nil, "")

	// resources := []map[string]interface{}{r}
	// weights := su.resourceRules(resources)
	// expecting := 3 * su.ResourceTypeScores[desiredType] * su.ResourceTypeScores["replicaset"]
	// if weights != expecting {
	// 	t.Errorf("no %v unexpected weights were calculated expecting: %v got %v", desiredType, expecting, weights)
	// }
}
