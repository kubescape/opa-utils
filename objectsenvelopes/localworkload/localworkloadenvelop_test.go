package localworkload

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func NewLocalWorkloadMck() *LocalWorkload {
	object := map[string]interface{}{
		"kind": "b",
		"path": "/path/file",
	}
	return NewLocalWorkload(object)
}

func TestGetPath(t *testing.T) {
	m := NewLocalWorkloadMck()
	assert.Equal(t, "/path/file", m.GetPath())
}

func TestSetPath(t *testing.T) {
	m := NewLocalWorkloadMck()
	m.SetPath("/bla")
	assert.Equal(t, "/bla", m.GetPath())
}

func TestGetKind(t *testing.T) {
	m := NewLocalWorkloadMck()
	assert.Equal(t, "b", m.GetKind())
}

func TestGetID(t *testing.T) {
	m := NewLocalWorkloadMck()
	assert.Equal(t, "path=L3BhdGgvZmlsZQ==////b/", m.GetID())
}
