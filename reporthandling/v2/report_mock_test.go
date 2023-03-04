package v2

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/kubescape/opa-utils/reporthandling/mock"
	"github.com/stretchr/testify/require"
)

func MockPostureReport() *PostureReport {
	v := mock.MockData[PostureReport]()

	return &v
}

func TestPostureReportJSON(t *testing.T) {
	v := MockPostureReport()

	spew.Dump(v)

	buf, err := json.Marshal(v)
	require.NoError(t, err)

	require.NoError(t,
		os.WriteFile("fake_posture_report.json", buf, 0666),
	)
}
