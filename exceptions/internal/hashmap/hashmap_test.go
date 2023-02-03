package hashmap

import (
	"crypto/rand"
	"log"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHashCollisions(t *testing.T) {
	// Test hashes on a large number of random map[string]string.
	// NOTE: this test is not very powerful and can only assert gross blunders
	// in the hash implementation. In particular, it is not large enough to
	// compare MSet-XOR-Hash and MSet-Mu-Hash.
	const (
		// n      = 100000000 // This is on a local computer with a lot of memory...
		n      = 1000 // This is for smoke-testing on CI
		size   = 8
		length = 4 // string lengths
	)

	hashes := make(map[uint64]map[string]string, n)
	generator := &mapsGenerator{n: n, size: size, length: length}

	for {
		attrs := generator.Next()
		if attrs == nil {
			break
		}

		hash := HashMap(attrs)
		collided, found := hashes[hash]
		if found {
			// verify that the maps are actually the same
			require.EqualValues(t, collided, attrs,
				"expected hash to be unique for different maps: found a collision: %v ~ %v", attrs, collided,
			)
		}
		hashes[hash] = attrs
	}

	t.Logf("successfully compared %d hashes of random maps of size %d", n, size)
}

// mapsGenerator generate maps of "size" elements with keys and values random strings of length "length"
type mapsGenerator struct {
	n         int
	size      int
	length    int
	generated int
}

func (g *mapsGenerator) Next() map[string]string {
	if g.generated >= g.size {
		return nil
	}

	result := make(map[string]string, g.size)
	for i := 0; i < g.size; i++ {
		key := make([]byte, g.length)
		_, err := rand.Read(key)
		if err != nil {
			log.Printf("warning: random generator error: %v", err)
			return nil
		}
		value := make([]byte, g.length)
		_, err = rand.Read(value)
		if err != nil {
			log.Printf("warning: random generator error: %v", err)
			return nil
		}

		result[string(key)] = string(value)
	}

	g.generated++

	return result
}

func multmod(a, b uint64) uint64 {
	// a reference implementation of a * b mod q, without any optimization
	x := big.NewInt(0)
	x.SetUint64(a)

	y := big.NewInt(0)
	y.SetUint64(b)

	prime := big.NewInt(0)
	prime.SetUint64(q)

	z := big.NewInt(0).Mul(x, y)
	m := big.NewInt(0).Mod(z, prime)
	if !m.IsUint64() {
		panic("must be representable as a uint64")
	}

	return m.Uint64()
}

func TestFieldMult(t *testing.T) {
	operands := []uint64{ // taken from sample outputs from FNV hashes
		4645756927087552409,
		16162368515513178030,
		18164748899638307352,
		10422122549694193689,
		17734110399310544832,
		1972183573942468894,
		13759071485432025751,
		11592539056557247347,
	}

	for i, a := range operands {
		var b uint64
		if i == len(operands)-1 {
			b = operands[0]
		} else {
			b = operands[i+1]
		}

		p1 := multmod(a, b)
		p2 := fieldMul64(a, b)
		require.Equalf(t, p1, p2,
			"expected %d x %d mod %d = %d but got %d", a, b, q, p1, p2,
		)
	}
}
