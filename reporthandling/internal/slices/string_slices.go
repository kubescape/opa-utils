package slices

// UniqueStrings returns the unique (unsorted) values of a slice.
//
// NOTE: this function does not allocate extra memory: the input slice is altered in-place.
//
// The returned slice is truncated in capacity to the number of unique elements.
func UniqueStrings(strSlice []string) []string {
	if len(strSlice) < 2 {
		return strSlice
	}

	strMap := make(map[string]struct{}, len(strSlice))
	uniqueIndex := 1
	strMap[strSlice[0]] = struct{}{}

	for _, val := range strSlice[1:] {
		if _, isDupe := strMap[val]; isDupe {
			continue
		}

		strSlice[uniqueIndex] = val
		strMap[val] = struct{}{}
		uniqueIndex++
	}

	return strSlice[:uniqueIndex:uniqueIndex]
}

// Trim returns the origin slice with all the elements from "trimFrom" removed.
//
// The original order of the original elements may be altered.
//
// NOTE: this function does not allocate extra memory: the input slice is altered in-place.
//
// The returned slice is truncated in capacity to the number of unique elements.
func Trim(origin, trimFrom []string) []string {
	length := len(origin)
	if length == 0 || len(trimFrom) == 0 { // nothing to trim
		return origin
	}

	// For a map that does not escape the func, the go runtime will grow the buckets in the stack if we keep the initial hint to 0.
	// See our comment above for function UniqueStrings().
	inTrimList := make(map[string]struct{}, len(trimFrom))

	for _, val := range trimFrom {
		inTrimList[val] = struct{}{}
	}

	// NOTE: this version is very similar to the original "trimUnique" implementation, but does not
	// need an extra function that manipulates pointers. Its performance is also about the same.
	for i := 0; i < length; {
		if _, found := inTrimList[origin[i]]; found {
			origin[i] = origin[length-1] // drop the i-th element and replace it by the last element
			length--

			continue
		}

		i++
	}

	return origin[:length:length]
}

// TrimStable returns the origin slice with all the elements from "trimFrom" removed.
//
// The original order of the original elements is maintained. Memory and CPU efficiency is about the same as Trim().
//
// NOTE: this function does not allocate extra memory: the input slice is altered in-place.
//
// The returned slice is truncated in capacity to the number of unique elements.
func TrimStable(origin, trimFrom []string) []string {
	length := len(origin)
	if length == 0 || len(trimFrom) == 0 { // nothing to trim
		return origin
	}

	// See our comment above for function UniqueStrings().
	inTrimList := make(map[string]struct{}, len(trimFrom))

	for _, val := range trimFrom {
		inTrimList[val] = struct{}{}
	}

	trimmedIndex := 0
	for i, val := range origin {
		if _, found := inTrimList[val]; !found {
			if trimmedIndex < i { // copy only if shifting has begun
				origin[trimmedIndex] = val // shift the current value left
			}
			trimmedIndex++
		}
	}

	return origin[:trimmedIndex:trimmedIndex]
}

// TrimStableUnique combines UniqueStrings and Trim in a single iteration.
//
// It returns the unique elements of the origin slice with all the elements from "trimFrom" removed.
//
// NOTE: this function does not allocate extra memory: the input slice is altered in-place.
//
// The returned slice is truncated in capacity to the number of unique elements.
//
// Calling TrimUnique in combination is slighly more efficient than calling these functions separately
// (the longer the slices, the larger the savings).
func TrimStableUnique(origin, trimFrom []string) []string {
	if len(origin) == 0 || len(trimFrom) == 0 { // nothing to trim
		return origin
	}

	// See our comment above for function UniqueStrings().
	inTrimList := make(map[string]struct{}, len(trimFrom))
	strMap := make(map[string]struct{}, len(origin))

	for _, val := range trimFrom {
		inTrimList[val] = struct{}{}
	}

	trimmedIndex := 0
	for _, val := range origin {
		if _, isDupe := strMap[val]; !isDupe {
			strMap[val] = struct{}{}

			if _, found := inTrimList[val]; !found {
				origin[trimmedIndex] = val
				trimmedIndex++
			}
		}
	}

	return origin[:trimmedIndex:trimmedIndex]
}

// TrimUnique combines UniqueStrings and Trim in a single iteration.
//
// It returns the unique elements of the origin slice with all the elements from "trimFrom" removed.
//
// NOTE: this function does not allocate extra memory: the input slice is altered in-place.
//
// The returned slice is truncated in capacity to the number of unique elements.
//
// Calling TrimUnique in combination is slighly more efficient than calling these functions separately
// (the longer the slices, the larger the savings).
func TrimUnique(origin, trimFrom []string) []string {
	length := len(origin)
	if len(origin) == 0 || len(trimFrom) == 0 { // nothing to trim
		return origin
	}

	inTrimList := make(map[string]struct{}, len(trimFrom))
	strMap := make(map[string]struct{}, len(origin))

	for _, val := range trimFrom {
		inTrimList[val] = struct{}{}
	}

	for i := 0; i < length; {
		val := origin[i]
		if _, isDupe := strMap[val]; !isDupe {
			strMap[val] = struct{}{}

			if _, found := inTrimList[val]; found {
				origin[i] = origin[length-1] // drop the i-th element and replace if by the last element
				length--

				continue
			}
		}

		i++
	}

	return origin[:length:length]
}

// StringInSlice determines if a needle is in a haystack of strings.
func StringInSlice(haystack []string, needle string) bool {
	for _, val := range haystack {
		if val == needle {
			return true
		}
	}

	return false
}
