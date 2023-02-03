package hashmap

import (
	"hash/fnv"
	"math/bits"
	"unsafe"
)

const (
	q       uint64 = 18446744069414584321 // some frequently used high prime number to produce finite field arithmetics modulo q
	qInvNeg uint64 = 18446744069414584319 // q + r'.r = 1, i.e., qInvNeg = - q⁻¹ mod r
	rSquare uint64 = 18446744065119617025 // the Montgommery constant
)

// HashMap computes a set-hash as a unique signature for the content of the input map.
//
// FNV64a is the unitary hash used for every (key,value) pair.
//
// The hash combination is based on the MSet-Mu-Hash method described in [1](https://people.csail.mit.edu/devadas/pubs/mhashes.pdf) (section 5).
//
// The faster MSet-XOR-Hash method has been discarded because of some extra requirements (unique nonce, equality with modulo)
// that make it unsuitable for cache indexing.
func HashMap(input map[string]string) uint64 {
	if len(input) == 0 {
		return 0
	}

	sum := uint64(1)
	h := fnv.New64a()

	for k, v := range input {
		bk, bv := hackZeroAlloc(k, v)

		// The hash is case sensitive. This shouldn't have any significant impact on performances
		_, _ = h.Write(bk)
		_, _ = h.Write([]byte{'0'})
		_, _ = h.Write(bv)

		// MSet-Mu-Hash: sum = sum * h.Sum64() mod q (product in finite field)
		// NOTE: for the record with MSet-Xor-Hash, we have sum ^= h.Sum64() (with initial value 0)
		sum = fieldMul64(sum, h.Sum64())
		h.Reset()
	}

	return sum
}

// fieldMul64 yields x*y mod q on uint64 operands.
//
// From https://github.com/consensys/gnark-crypto/blob/v0.9.0/field/goldilocks/element_ops_purego.go#L59
//
// Also see:
// * https://en.wikipedia.org/wiki/Montgomery_modular_multiplication
func fieldMul64(x, y uint64) uint64 {
	// to Montgomery form
	x, y = tomont(x), tomont(y)

	var r uint64

	hi, lo := bits.Mul64(x, y)
	if lo != 0 {
		hi++ // x * y ≤ 2¹²⁸ - 2⁶⁵ + 1, meaning hi ≤ 2⁶⁴ - 2 so no need to worry about overflow
	}

	m := lo * qInvNeg
	hi2, _ := bits.Mul64(m, q)
	r, carry := bits.Add64(hi2, hi, 0)
	if carry != 0 || r >= q {
		r -= q // we need to reduce
	}

	// from Montgomery form
	return frommont(r)
}

func tomont(x uint64) uint64 {
	var r uint64

	hi, lo := bits.Mul64(x, rSquare)
	if lo != 0 {
		hi++
	}

	m := lo * qInvNeg
	hi2, _ := bits.Mul64(m, q)
	r, carry := bits.Add64(hi2, hi, 0)
	if carry != 0 || r >= q {
		r -= q
	}

	return r
}

func frommont(x uint64) uint64 {
	m := x * qInvNeg
	c := madd0(m, q, x)
	x = c

	if x >= q {
		x -= q
	}

	return x
}

// madd0 hi = a*b + c (discards lo bits)
func madd0(a, b, c uint64) (hi uint64) {
	var carry, lo uint64
	hi, lo = bits.Mul64(a, b)
	_, carry = bits.Add64(lo, c, 0)
	hi, _ = bits.Add64(hi, 0, carry)
	return
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
// having relinquished the input string to the garbage collector.
func hackZeroAlloc(k, v string) ([]byte, []byte) {
	addrK := (*internalString)(unsafe.Pointer(&k)).Data
	bk := unsafe.Slice((*byte)(addrK), len(k))

	addrV := (*internalString)(unsafe.Pointer(&v)).Data
	bv := unsafe.Slice((*byte)(addrV), len(v))

	return bv, bk
}
