package exceptions

import (
	"sync"

	"github.com/armosec/armoapi-go/identifiers"
	"github.com/kubescape/opa-utils/exceptions/internal/hashmap"
)

type (
	// designatorCache knows how to cache designators.
	//
	// We use a plain map with mutex instead of sync.Map, so we may preallocate
	// a few slots for designators.
	designatorCache struct {
		mx       sync.RWMutex
		innerMap map[portalDesignatorKey]identifiers.AttributesDesignators
	}

	portalDesignatorKey struct {
		DesignatorType identifiers.DesignatorType
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
		innerMap: make(map[portalDesignatorKey]identifiers.AttributesDesignators, heuristicAllocDesignators),
	}
}

func (c *designatorCache) Get(designator *identifiers.PortalDesignator) (identifiers.AttributesDesignators, bool) {
	key := c.toDesignatorKey(designator)

	c.mx.RLock()
	defer c.mx.RUnlock()

	val, ok := c.innerMap[key]

	return val, ok
}

func (c *designatorCache) Set(designator *identifiers.PortalDesignator, value identifiers.AttributesDesignators) {
	key := c.toDesignatorKey(designator)

	c.mx.Lock()
	defer c.mx.Unlock()

	c.innerMap[key] = value
}

func (c *designatorCache) toDesignatorKey(designator *identifiers.PortalDesignator) portalDesignatorKey {
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
