package v1alpha1

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAttackTrackUnmarshal(t *testing.T) {
	var obj AttackTrack
	file, _ := os.ReadFile(filepath.Join("testdata", "attacktrack.json"))
	err := json.Unmarshal([]byte(file), &obj)
	assert.NoError(t, err)
}
