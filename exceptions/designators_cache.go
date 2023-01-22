package exceptions

import (
	"hash/fnv"
	"sync"
	"unsafe"

	"github.com/armosec/armoapi-go/armotypes"
)

type (
	// designatorCache knows how to cache designators.
	//
	// We use a plain map with mutex instead of sync.Map, so we may preallocate
	// a few slots for designators.
	designatorCache struct {
		mx       sync.RWMutex
		innerMap map[portalDesignatorKey]armotypes.AttributesDesignators
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
//
// NOTE(fredbi): the inner cache uses the FNV hashing function.
// Experiments with hash/maphash proved to be about 30% slower.
func newDesignatorCache() *designatorCache {
	const heuristicAllocDesignators = 1000 // this is a hint on the number of AttributeDesignators to hold, in order to minimize dynamic reallocations for this map

	return &designatorCache{
		innerMap: make(map[portalDesignatorKey]armotypes.AttributesDesignators, heuristicAllocDesignators),
	}
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
		// this feeds a unique hash of the attributes in order to make the key indexable.
		AttributesHash: c.hashMap(designator.Attributes),
	}
}

// hashMap computes a FNV hash as a unique signature for the content of the input map.
//
// NOTE(fredbi):
// * benefits: ~30% faster than hash/maphash, no seeding
// * shortcomings: no optimized API to write strings: requires a hack to reuse strings as []byte without extra alloc
//
//	goos: linux
//	goarch: amd64
//	pkg: github.com/kubescape/opa-utils/exceptions
//	cpu: Intel(R) Core(TM) i5-6200U CPU @ 2.30GHz
//	BenchmarkCache
//	BenchmarkCache-4   1715834	      705.2 ns/op	       0 B/op	       0 allocs/op
func (c *designatorCache) hashMap(input map[string]string) uint64 {
	if len(input) == 0 {
		return 0
	}

	var sum uint64
	h := fnv.New32a() // we limit the inner hash to 32 bits to reduce the likelihood of collisions when summing up

	for k, v := range input {
		bk, bv := hackZeroAlloc(k, v)

		// The hash is case sensitive. This shouldn't have any significant impact on performances
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

// internalString representation of a string by the golang runtime
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
