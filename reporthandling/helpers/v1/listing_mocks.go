package helpers

import "github.com/kubescape/opa-utils/reporthandling/apis"

func MockAllListsForIntegration() *AllLists {
	mock := &AllLists{}
	mock.Append(apis.StatusPassed, "a", "b")
	mock.Append(apis.StatusFailed, "a", "e")
	mock.Append(apis.StatusSkipped, "f")
	mock.Append(apis.StatusUnknown, "i", "g", "h", "i")

	return mock
}
