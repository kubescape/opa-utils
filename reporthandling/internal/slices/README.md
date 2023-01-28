# slices

Internal utility functions that manipulate slices of strings.

All these methods alter their input slice and do not allocate extra memory.

`UniqueStrings` is provided as a faster equivalent to `github.com/armosec/utils-go/str.SliceStringToUnique`.

`Trim` is the same as the previously private function `trimUnique()` with similar performances. Code is just more straightforward.

`TrimStable` has the same intent as `Trim` but does not alter the order of the input.

`TrimUnique` and `TrimStableUnique` combine `UniqueStrings` and `Trim` (resp. `TrimStable`) in a single iteration of the input slice.

## Benchmarks

```
go test -v -run Bench -bench . -benchtime 1s
goos: linux
goarch: amd64
pkg: github.com/kubescape/opa-utils/reporthandling/internal/slices
cpu: AMD Ryzen 7 5800X 8-Core Processor             
BenchmarkUnique
BenchmarkUnique/UniqueStrings_x_8
BenchmarkUnique/UniqueStrings_x_8-16         	 6007258	       192.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnique/UniqueStrings_x_16
BenchmarkUnique/UniqueStrings_x_16-16        	 3084657	       386.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnique/UniqueStrings_x_32
BenchmarkUnique/UniqueStrings_x_32-16        	 3148867	       387.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnique/SliceStringToUnique_x_8
BenchmarkUnique/SliceStringToUnique_x_8-16   	 2188087	       543.4 ns/op	      64 B/op	       1 allocs/op
BenchmarkUnique/SliceStringToUnique_x_16
BenchmarkUnique/SliceStringToUnique_x_16-16  	 2219115	       544.2 ns/op	      64 B/op	       1 allocs/op
BenchmarkUnique/SliceStringToUnique_x_32
BenchmarkUnique/SliceStringToUnique_x_32-16  	 2175615	       548.6 ns/op	      64 B/op	       1 allocs/op
BenchmarkTrim
BenchmarkTrim/Trim_x_8
BenchmarkTrim/Trim_x_8-16                    	 5328590	       222.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkTrim/Trim_x_16
BenchmarkTrim/Trim_x_16-16                   	 2748080	       422.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkTrim/Trim_x_32
BenchmarkTrim/Trim_x_32-16                   	 2732326	       425.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkTrim/TrimStable_x_8
BenchmarkTrim/TrimStable_x_8-16              	 5337442	       216.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkTrim/TrimStable_x_16
BenchmarkTrim/TrimStable_x_16-16             	 2752138	       435.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkTrim/TrimStable_x_32
BenchmarkTrim/TrimStable_x_32-16             	 2770256	       432.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkTrim/TrimSwap_x_8_(original_version)
BenchmarkTrim/TrimSwap_x_8_(original_version)-16         	 5299237	       225.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkTrim/TrimSwap_x_16_(original_version)
BenchmarkTrim/TrimSwap_x_16_(original_version)-16        	 2744328	       437.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkTrim/TrimSwap_x_32_(original_version)
BenchmarkTrim/TrimSwap_x_32_(original_version)-16        	 2800281	       440.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkTrim/TrimUnique_x_8
BenchmarkTrim/TrimUnique_x_8-16                          	 4220919	       284.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkTrim/TrimUnique_x_16
BenchmarkTrim/TrimUnique_x_16-16                         	 2469658	       482.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkTrim/TrimUnique_x_32
BenchmarkTrim/TrimUnique_x_32-16                         	 2408644	       494.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkTrim/TrimStableUnique_x_8
BenchmarkTrim/TrimStableUnique_x_8-16                    	 4250454	       288.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkTrim/TrimStableUnique_x_16
BenchmarkTrim/TrimStableUnique_x_16-16                   	 2390988	       486.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkTrim/TrimStableUnique_x_32
BenchmarkTrim/TrimStableUnique_x_32-16                   	 2416376	       495.0 ns/op	       0 B/op	       0 allocs/op
PASS
ok 
```
