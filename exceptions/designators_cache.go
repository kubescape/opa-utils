package exceptions

import (
	"hash/maphash"
	"sync"

	"github.com/armosec/armoapi-go/armotypes"
)

var (
	globalDesignatorCache *designatorCache
	setGlobalCacheOnce    sync.Once
)

type (
	// designable knows how to identify the attributes from a Designator.
	//
	// TODO(fred): perf - this interface circumvents the unexported nature
	// of armotypes.attributesDesignator.
	// We could avoid this extra level of indirection by exporting that type.
	designable interface {
		GetCluster() string
		GetNamespace() string
		GetKind() string
		GetName() string
		GetPath() string
		GetLabels() map[string]string
	}

	// designatorCache knows how to cache designators.
	designatorCache struct {
		innerMap sync.Map
		seed     maphash.Seed
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
			seed: maphash.MakeSeed(),
		}
	})

	return globalDesignatorCache
}

func (c *designatorCache) Get(designator *armotypes.PortalDesignator) (designable, bool) {
	val, ok := c.innerMap.Load(c.toDesignatorKey(designator))
	if ok {
		return val.(designable), true
	}

	return nil, false
}

func (c *designatorCache) Set(designator *armotypes.PortalDesignator, value designable) {
	c.innerMap.Store(c.toDesignatorKey(designator), value)
}

func (c *designatorCache) toDesignatorKey(designator *armotypes.PortalDesignator) portalDesignatorKey {
	return portalDesignatorKey{
		DesignatorType: designator.DesignatorType,
		WLID:           designator.WLID,
		WildWLID:       designator.WildWLID,
		SID:            designator.SID,
		AttributesHash: c.hashMap(designator.Attributes), // a unique hash on attributes in order to make the key indexable
	}
}

func (c *designatorCache) hashMap(input map[string]string) uint64 {
	var (
		sum uint64
		h   maphash.Hash
	)

	for k, v := range input {
		// NOTE: the hash is case sensitive. This shouldn't have any significant impact on performances
		h.WriteString(k)
		h.WriteByte(0)
		h.WriteString(v)

		// final hash has to be insensitive to the ordering of the map
		sum += h.Sum64()
	}

	return sum
}
