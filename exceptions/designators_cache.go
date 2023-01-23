package exceptions

import (
	"hash/fnv"
	// "hash/maphash"
	"sync"
	"unsafe"

	"github.com/armosec/armoapi-go/armotypes"
)

var (
	globalDesignatorCache *designatorCache
	setGlobalCacheOnce    sync.Once
)

type (
	// designatorCache knows how to cache designators.
	//
	// We use a plain map with mutex instead of sync.Map, so we may preallocate
	// a few slots for designators.
	designatorCache struct {
		mx       sync.RWMutex
		innerMap map[portalDesignatorKey]armotypes.AttributesDesignators
		// seed     maphash.Seed
	}

	portalDesignatorKey struct {
		DesignatorType armotypes.DesignatorType
		WLID           string
		WildWLID       string
		SID            string
		AttributesHash uint64 // summarizes the map[string]string with a single hash value
	}
)

// newDesignatorCache builds a chache for designators.
//
// This builds a single global instance only on the first call.
func newDesignatorCache() *designatorCache {
	setGlobalCacheOnce.Do(func() {
		globalDesignatorCache = &designatorCache{
			innerMap: make(map[portalDesignatorKey]armotypes.AttributesDesignators, 1000),
			// seed:     maphash.MakeSeed(), // for map/maphash
		}
	})

	return globalDesignatorCache
}

func (c *designatorCache) Get(designator *armotypes.PortalDesignator) (armotypes.AttributesDesignators, bool) {
	key := c.toDesignatorKey(designator)

	c.mx.RLock()
	defer c.mx.RUnlock()

	val, ok := c.innerMap[key]

	return val, ok
}

func (c *designatorCache) Set(designator *armotypes.PortalDesignator, value armotypes.AttributesDesignators) {
	key := c.toDesignatorKey(designator)

	c.mx.Lock()
	defer c.mx.Unlock()

	c.innerMap[key] = value
}

func (c *designatorCache) toDesignatorKey(designator *armotypes.PortalDesignator) portalDesignatorKey {
	return portalDesignatorKey{
		DesignatorType: designator.DesignatorType,
		WLID:           designator.WLID,
		WildWLID:       designator.WildWLID,
		SID:            designator.SID,
		// This feeds a unique hash of the attributes in order to make the key indexable.
		AttributesHash: c.hashMap(designator.Attributes),
	}
}

/* hash/maphash version
   	Notice the use of the package level hasher, rather than the hash struct.
	Hack:  the expression k+"|"+v does not generate extra allocation.

	goos: linux
	goarch: amd64
	pkg: github.com/kubescape/opa-utils/exceptions
	cpu: Intel(R) Core(TM) i5-6200U CPU @ 2.30GHz
	BenchmarkCache
	BenchmarkCache-4   	 1240087	       960.4 ns/op	       0 B/op	       0 allocs/op

func (c *designatorCache) hashMap(input map[string]string) uint64 {
	var sum uint64

	for k, v := range input {
		// NOTE: the hash is case sensitive. This shouldn't have any significant impact on performances
		// final hash has to be insensitive to the ordering of the map
		sum += maphash.String(c.seed, k+"|"+v) // NOTE: the string expression is normally recognized by the compiler and does not generate extra allocation
	}

	return sum
}
*/

/*
	 fnv version:
	 * benefits: ~30% faster than hash/maphash, no seeding
	 * shortcomings: no optimized API to write strings - requires a hack to reuse strings as []byte without extra alloc

		goos: linux
		goarch: amd64
		pkg: github.com/kubescape/opa-utils/exceptions
		cpu: Intel(R) Core(TM) i5-6200U CPU @ 2.30GHz
		BenchmarkCache
		BenchmarkCache-4   1715834	      705.2 ns/op	       0 B/op	       0 allocs/op
*/
func (c *designatorCache) hashMap(input map[string]string) uint64 {
	if len(input) == 0 {
		return 0
	}

	var sum uint64
	h := fnv.New32a() // NOTE: we limit the inner hash to 32 bits to reduce the likelihood of overfow

	for k, v := range input {
		bk, bv := hackZeroAlloc(k, v)

		// NOTE: the hash is case sensitive. This shouldn't have any significant impact on performances
		_, _ = h.Write(bk)
		_, _ = h.Write([]byte{'0'})
		_, _ = h.Write(bv)
		// The final hash has to be insensitive to the order in which the map is iterated.
		// Summing partial hashes increase the likelhood of collisions on the final result.
		sum += uint64(h.Sum32())
		h.Reset()
	}

	return sum
}

// internalString representation by golang runtime
type internalString struct {
	Data unsafe.Pointer
	Len  int
}

// hackZeroAlloc reuses a common hack found in the standard library
// to avoid allocating the underlying bytes of a string when converting.
//
// This assumes that the caller does not use the returned []byte slices after
// having relinquished the input strings to the garbage collector.
func hackZeroAlloc(k, v string) ([]byte, []byte) {
	addrK := (*internalString)(unsafe.Pointer(&k)).Data
	bk := unsafe.Slice((*byte)(addrK), len(k))

	addrV := (*internalString)(unsafe.Pointer(&v)).Data
	bv := unsafe.Slice((*byte)(addrV), len(v))

	return bv, bk
}
