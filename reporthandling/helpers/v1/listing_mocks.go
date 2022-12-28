package v1

func MockAllListsForIntegration() *AllLists {
	return &AllLists{
		passed:                []string{"a", "b"},
		passedExceptions:      []string{"c"},
		passedIrrelevant:      []string{"d"},
		skippedConfiguration:  []string{"e", "a"},
		skippedIntegration:    []string{"g"},
		skippedRequiresReview: []string{"h", "i"},
		skippedManualReview:   []string{"j", "k"},
		failed:                []string{"e", "g"},
		other:                 []string{"i", "l", "m", "n"},
	}
}
