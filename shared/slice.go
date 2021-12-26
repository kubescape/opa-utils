package shared

// StringInSlice return true if string found in slice of strings
func StringInSlice(strSlice []string, str string) bool {
	for i := range strSlice {
		if strSlice[i] == str {
			return true
		}
	}
	return false
}

// MapStringToSlice returns map's keys
func MapStringToSlice(strMap map[string]interface{}) []string {
	strSlice := []string{}
	for k := range strMap {
		strSlice = append(strSlice, k)
	}
	return strSlice
}

// SliceStringToUnique returns unique values of slice
func SliceStringToUnique(strSlice []string) []string {
	strMap := map[string]interface{}{}
	for i := range strSlice {
		strMap[strSlice[i]] = nil
	}
	return MapStringToSlice(strMap)
}
