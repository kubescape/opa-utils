package exceptions

import (
	// "hash/maphash"
	"testing"

	"github.com/armosec/armoapi-go/armotypes"
	// "github.com/pkg/profile"
	"github.com/stretchr/testify/require"
)

func TestDesignatorsCache(t *testing.T) {
	cache := &designatorCache{
		innerMap: make(map[portalDesignatorKey]armotypes.AttributesDesignators, 1000),
		// seed:     maphash.MakeSeed(), // for maphash version
	}

	t.Run("should retrieve cached designator", func(t *testing.T) {
		designator := &armotypes.PortalDesignator{
			DesignatorType: armotypes.DesignatorAttributes,
			WLID:           "x",
			WildWLID:       "y",
			SID:            "z",
			Attributes: map[string]string{
				"1": "2",
				"3": "4",
			},
		}

		attrs := designator.DigestPortalDesignator()
		_, found := cache.Get(designator)
		require.False(t, found)

		cache.Set(designator, attrs)

		retrieved, found := cache.Get(designator)
		require.True(t, found)

		require.EqualValues(t, attrs, retrieved)
	})

	t.Run("should not collide with previously cached designator (WLID differs)", func(t *testing.T) {
		designator := &armotypes.PortalDesignator{
			DesignatorType: armotypes.DesignatorAttributes,
			WLID:           "x1",
			WildWLID:       "y",
			SID:            "z",
			Attributes: map[string]string{
				"1": "2",
				"3": "4",
			},
		}

		attrs := designator.DigestPortalDesignator()
		_, found := cache.Get(designator)
		require.False(t, found)

		cache.Set(designator, attrs)

		retrieved, found := cache.Get(designator)
		require.True(t, found)

		require.EqualValues(t, attrs, retrieved)
	})

	t.Run("should not collide with previously cached designator (attributes differ)", func(t *testing.T) {
		designator := &armotypes.PortalDesignator{
			DesignatorType: armotypes.DesignatorAttributes,
			WLID:           "x",
			WildWLID:       "y",
			SID:            "z",
			Attributes: map[string]string{
				"1": "2",
				"3": "4",
				"5": "6",
			},
		}

		attrs := designator.DigestPortalDesignator()
		_, found := cache.Get(designator)
		require.False(t, found)

		cache.Set(designator, attrs)

		retrieved, found := cache.Get(designator)
		require.True(t, found)

		require.EqualValues(t, attrs, retrieved)
	})

	t.Run("should support empty attributes", func(t *testing.T) {
		designator := &armotypes.PortalDesignator{
			DesignatorType: armotypes.DesignatorAttributes,
			WLID:           "x",
			WildWLID:       "y",
			SID:            "z",
		}

		attrs := designator.DigestPortalDesignator()
		_, found := cache.Get(designator)
		require.False(t, found)

		cache.Set(designator, attrs)

		retrieved, found := cache.Get(designator)
		require.True(t, found)

		require.EqualValues(t, attrs, retrieved)
	})
}

func BenchmarkCache(b *testing.B) {
	cache := &designatorCache{
		innerMap: make(map[portalDesignatorKey]armotypes.AttributesDesignators, 1000),
		// seed:     maphash.MakeSeed(), // for maphash version
	}

	designator := &armotypes.PortalDesignator{
		DesignatorType: armotypes.DesignatorAttributes,
		WLID:           "x",
		WildWLID:       "y",
		SID:            "z",
		Attributes: map[string]string{
			"1": "2",
			"3": "4",
			"5": "6",
		},
	}
	attrs := designator.DigestPortalDesignator()

	b.ResetTimer()
	b.ReportAllocs()
	b.SetBytes(0)

	for n := 0; n < b.N; n++ {
		_, _ = cache.Get(designator)
		cache.Set(designator, attrs)
		_, _ = cache.Get(designator)
	}
}

/* TODO: profiling code - remove from the final (merged) version
func TestProfileCache(t *testing.T) {
	cache := &designatorCache{
		innerMap: make(map[portalDesignatorKey]armotypes.AttributesDesignators, 1000),
		// seed:     maphash.MakeSeed(), // for maphash version
	}

	designator := &armotypes.PortalDesignator{
		DesignatorType: armotypes.DesignatorAttributes,
		WLID:           "x",
		WildWLID:       "y",
		SID:            "z",
		Attributes: map[string]string{
			"1": "2",
			"3": "4",
			"5": "6",
		},
	}
	attrs := designator.DigestPortalDesignator()
	defer profile.Start(
		profile.CPUProfile,
		profile.ProfilePath("cpuprof"),
		profile.NoShutdownHook,
	).Stop()

	for n := 0; n < 1000000; n++ {
		_, _ = cache.Get(designator)
		cache.Set(designator, attrs)
		_, _ = cache.Get(designator)
	}
}
*/
