package v1

import "k8s.io/apimachinery/pkg/util/sets"

func MockAllListsForIntegration() *AllLists {
	return &AllLists{
		passed:   sets.New("a", "b"),
		failed:   sets.New("a", "e"),
		skipped:  sets.New("f"),
		other:    sets.New("i", "g", "h", "i"),
	}
}
