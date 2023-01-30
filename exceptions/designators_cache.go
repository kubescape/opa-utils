package exceptions

import (
	"sync"

	"github.com/armosec/armoapi-go/armotypes"
	"github.com/kubescape/opa-utils/exceptions/internal/hashmap"
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
		LenAttributes  int
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
		LenAttributes:  len(designator.Attributes),
		AttributesHash: hashmap.HashMap(designator.Attributes),
	}
}
