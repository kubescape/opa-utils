package helpers

import (
	"testing"

	"github.com/kubescape/opa-utils/reporthandling/apis"
)

func BenchmarkToUniqueResources(b *testing.B) {
	listA := mockAllListsA()
	listA.Append(apis.StatusExcluded, "b")
	listA.Append(apis.StatusExcluded, "b")
	listA.Append(apis.StatusExcluded, "b")
	b.ResetTimer()
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		listA.ToUniqueResources()
	}
}
