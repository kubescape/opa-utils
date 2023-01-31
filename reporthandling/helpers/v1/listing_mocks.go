package v1

func MockAllListsForIntegration() *AllLists {
	return &AllLists{
		passed:  []string{"a", "b"},
		failed:  []string{"a", "e"},
		skipped: []string{"f"},
		other:   []string{"i", "g", "h", "i"},
	}
}
