package localworkload

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func NewLocalWorkloadMck() *LocalWorkload {
	object := map[string]interface{}{
		"kind":  "b",
		PathKey: "/path/file",
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
	assert.Equal(t, "path=1336429864/api=///b/", m.GetID())
}

func TestSetID(t *testing.T) {
	m := NewLocalWorkloadMck()
	m.SetPath("")
	assert.Equal(t, "path=2166136261/api=///b/", m.GetID())
}

func TestDeletePathEntry(t *testing.T) {
	m := NewLocalWorkloadMck()
	assert.Equal(t, "/path/file", m.GetObject()[PathKey].(string))

	m.DeletePathEntry()
	assert.Equal(t, nil, m.GetObject()[PathKey])
}
