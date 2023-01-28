package slices

import (
	"math/rand"
	"testing"

	"github.com/armosec/utils-go/str"
)

var (
	fixedInput8, fixedInput16, fixedInput32, manyInput32, fixedTrimList           []string
	currentInput8, currentInput16, currentInput32, currentMany32, currentTrimList []string
)

func init() {
	// We make sure that we don't bias in benchmark alloc reporting.
	// Allocation is performed only once, and each (serialized) benchmark
	// copies the original values to the allocated space.
	//
	// Since the benchmarked methods rely on in-place modification, we want to
	// compare every iteration with the same original input, while not introducing
	// extraneous allocation due to the benchmarking machinery.
	fixedInput8 = []string{
		"A", "C", "B", "B", "A", "D", "A", "C",
	}
	fixedInput16 = []string{
		"A", "C", "B", "B", "A", "D", "A", "C",
		"A", "C", "B", "B", "A", "D", "A", "C",
	}
	fixedInput32 = []string{
		"A", "C", "B", "B", "A", "D", "A", "C",
		"A", "C", "B", "B", "A", "D", "A", "C",
		"A", "C", "B", "B", "A", "D", "A", "C",
		"A", "C", "B", "B", "A", "D", "A", "C",
	}
	manyInput32 = []string{
		"A", "C", "B", "E", "A", "D", "F", "G",
		"H", "H", "I", "J", "K", "L", "A", "C",
		"M", "N", "B", "O", "P", "Q", "E", "R",
		"S", "T", "U", "B", "V", "W", "X", "Y",
	}
	fixedTrimList = []string{
		"D", "A", "E",
	}

	currentInput8 = make([]string, len(fixedInput8))
	currentInput16 = make([]string, len(fixedInput16))
	currentInput32 = make([]string, len(fixedInput32))
	currentMany32 = make([]string, len(manyInput32))
	currentTrimList = make([]string, len(fixedTrimList))
}

func input8() []string {
	// copy original content, randomly shuffled
	copy(currentInput8, fixedInput8)

	return shuffle(currentInput8)
}

func input16() []string {
	// copy original content, randomly shuffled
	copy(currentInput16, fixedInput16)

	return shuffle(currentInput16)
}

func input32() []string {
	// copy original content, randomly shuffled
	copy(currentInput16, fixedInput16)

	return shuffle(currentInput16)
}

func many32() []string {
	// copy original content, randomly shuffled
	copy(currentInput16, fixedInput16)

	return shuffle(currentInput16)
}

func trimList() []string {
	copy(currentTrimList, fixedTrimList)

	return shuffle(currentTrimList)
}

// shuffle a slice
func shuffle(input []string) []string {
	for i := 0; i < len(input); i++ {
		j := rand.Intn(len(input)) //nolint:gosec
		input[i], input[j] = input[j], input[i]
	}

	return input
}

func BenchmarkUnique(b *testing.B) {
	// do not run in parallel

	b.Run("UniqueStrings x 8", runUniqueBenchmark(UniqueStrings, input8))
	b.Run("UniqueStrings x 16", runUniqueBenchmark(UniqueStrings, input16))
	b.Run("UniqueStrings x 32", runUniqueBenchmark(UniqueStrings, input32))
	b.Run("UniqueStrings x 32 (many)", runUniqueBenchmark(UniqueStrings, many32))

	b.Run("SliceStringToUnique x 8", runUniqueBenchmark(str.SliceStringToUnique, input16))
	b.Run("SliceStringToUnique x 16", runUniqueBenchmark(str.SliceStringToUnique, input16))
	b.Run("SliceStringToUnique x 32", runUniqueBenchmark(str.SliceStringToUnique, input32))
	b.Run("SliceStringToUnique x 32 (many)", runUniqueBenchmark(str.SliceStringToUnique, many32))
}

func BenchmarkTrim(b *testing.B) {
	// do not run in parallel

	b.Run("Trim x 8", runTrimBenchmark(Trim, input8, trimList))
	b.Run("Trim x 16", runTrimBenchmark(Trim, input16, trimList))
	b.Run("Trim x 32", runTrimBenchmark(Trim, input32, trimList))
	b.Run("Trim x 32 (many)", runTrimBenchmark(Trim, many32, trimList))

	b.Run("TrimStable x 8", runTrimBenchmark(TrimStable, input8, trimList))
	b.Run("TrimStable x 16", runTrimBenchmark(TrimStable, input16, trimList))
	b.Run("TrimStable x 32", runTrimBenchmark(TrimStable, input32, trimList))
	b.Run("TrimStable x 32", runTrimBenchmark(TrimStable, many32, trimList))

	b.Run("TrimSwap x 8 (original version)", runTrimBenchmark(originalTrimmer, input8, trimList))
	b.Run("TrimSwap x 16 (original version)", runTrimBenchmark(originalTrimmer, input16, trimList))
	b.Run("TrimSwap x 32 (original version)", runTrimBenchmark(originalTrimmer, input32, trimList))
	b.Run("TrimSwap x 32 (many, original version)", runTrimBenchmark(originalTrimmer, many32, trimList))

	b.Run("TrimUnique x 8", runTrimBenchmark(TrimUnique, input8, trimList))
	b.Run("TrimUnique x 16", runTrimBenchmark(TrimUnique, input16, trimList))
	b.Run("TrimUnique x 32", runTrimBenchmark(TrimUnique, input32, trimList))
	b.Run("TrimUnique x 32 (many)", runTrimBenchmark(TrimUnique, many32, trimList))

	b.Run("TrimStableUnique x 8", runTrimBenchmark(TrimStableUnique, input8, trimList))
	b.Run("TrimStableUnique x 16", runTrimBenchmark(TrimStableUnique, input16, trimList))
	b.Run("TrimStableUnique x 32", runTrimBenchmark(TrimStableUnique, input32, trimList))
	b.Run("TrimStableUnique x 32 (many)", runTrimBenchmark(TrimStableUnique, many32, trimList))
}

func runUniqueBenchmark(fn func([]string) []string, input func() []string) func(*testing.B) {
	return func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()

		for n := 0; n < b.N; n++ {
			input := input()
			_ = fn(input)
		}
	}
}

func runTrimBenchmark(fn func([]string, []string) []string, input, trimList func() []string) func(*testing.B) {
	return func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()

		for n := 0; n < b.N; n++ {
			input, trimList := input(), trimList()
			_ = fn(input, trimList)
		}
	}
}

// originalTrimmer reminds us of the original implementation of the trimming, for the purpose
// of comparing benchmarks.
func originalTrimmer(origin, trimFrom []string) []string {
	if len(origin) == 0 || len(trimFrom) == 0 { // if there is nothing to trim
		return origin
	}
	toRemove := make(map[string]bool, len(trimFrom))

	for i := range trimFrom {
		toRemove[trimFrom[i]] = true
	}
	originLen := len(origin)
	for i := 0; i < originLen; {
		if _, ok := toRemove[origin[i]]; ok {
			str.RemoveIndexFromStringSlice(&origin, i) // swap i-th element with the last one
			originLen--
		} else {
			i++
		}
	}
	return origin
}
