package helpers

import (
	"testing"

	"github.com/armosec/armoapi-go/armotypes"
	"github.com/stretchr/testify/assert"
)

func mockNSAFW() *armotypes.PostureExceptionPolicy {
	return &armotypes.PostureExceptionPolicy{
		PosturePolicies: []armotypes.PosturePolicy{
			{
				FrameworkName: "NSA",
			},
		},
	}

}

func mockMITREFW() *armotypes.PostureExceptionPolicy {
	return &armotypes.PostureExceptionPolicy{
		PosturePolicies: []armotypes.PosturePolicy{
			{
				FrameworkName: "MITRE",
			},
		},
	}

}

func mockEmptyFW() *armotypes.PostureExceptionPolicy {
	return &armotypes.PostureExceptionPolicy{
		PosturePolicies: []armotypes.PosturePolicy{
			{
				FrameworkName: "",
			},
		},
	}

}
func TestFilterExceptions(t *testing.T) {
	f := Filters{}
	exceptions := []armotypes.PostureExceptionPolicy{
		*mockEmptyFW(),
		*mockMITREFW(),
	}
	exceptions2 := f.FilterExceptions(exceptions)
	assert.Equal(t, len(exceptions), len(exceptions2))

	exceptions = append(exceptions, *mockNSAFW())
	f.FrameworkNames = []string{"NSA"}
	exceptions3 := f.FilterExceptions(exceptions)
	assert.Equal(t, len(exceptions)-1, len(exceptions3))

	f.FrameworkNames = []string{"NSA", "MITRE"}
	exceptions4 := f.FilterExceptions(exceptions)
	assert.Equal(t, len(exceptions), len(exceptions4))

	f.FrameworkNames = []string{}
	exceptions5 := f.FilterExceptions(exceptions)
	assert.Equal(t, len(exceptions), len(exceptions5))

	exceptions6 := []armotypes.PostureExceptionPolicy{
		*mockMITREFW(),
	}
	f.FrameworkNames = []string{"NSA"}
	exceptions7 := f.FilterExceptions(exceptions6)
	assert.Equal(t, 0, len(exceptions7))
}
