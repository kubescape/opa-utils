package reportsummary

import (
	"testing"

	"github.com/armosec/opa-utils/reporthandling/apis"
	"github.com/stretchr/testify/assert"
)

func TestSetStatus(t *testing.T) {
	status := []apis.ScanningStatus{}

	f := mockFrameworkSummaryFailExclude()
	status = append(status, f.GetStatus().Status())

	f.Status = apis.StatusUnknown
	for k, v := range f.Controls {
		status = append(status, v.GetStatus().Status())
		v.Status = apis.StatusUnknown
		f.Controls[k] = v
	}

	i := 0
	assert.Equal(t, status[i], f.GetStatus().Status())
	for _, v := range f.Controls {
		i++
		assert.Equal(t, status[i], v.GetStatus().Status())
	}

}
