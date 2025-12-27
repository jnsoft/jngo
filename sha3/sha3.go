package sha3

// Keccak/SHA3 parameters (Keccak-f[1600])
const (
	stateSize = 25 // 5×5 lanes
	numRounds = 24

	sha3_256Rate    = 136 // 1088-bit rate
	sha3_256HashLen = 32  // 256 bits

	sha3_512Rate    = 72 // 576-bit rate
	sha3_512HashLen = 64 // 512 bits

	shake256Rate = 136 // 1088-bit rate
)

// round constants for Keccak-f[1600]
var roundConstants = [numRounds]uint64{
	0x0000000000000001,
	0x0000000000008082,
	0x800000000000808A,
	0x8000000080008000,
	0x000000000000808B,
	0x0000000080000001,
	0x8000000080008081,
	0x8000000000008009,
	0x000000000000008A,
	0x0000000000000088,
	0x0000000080008009,
	0x000000008000000A,
	0x000000008000808B,
	0x800000000000008B,
	0x8000000000008089,
	0x8000000000008003,
	0x8000000000008002,
	0x8000000000000080,
	0x000000000000800A,
	0x800000008000000A,
	0x8000000080008081,
	0x8000000000008080,
	0x0000000080000001,
	0x8000000080008008,
}

// rotation offsets r[x][y]
var rotationOffsets = [5][5]uint{
	{0, 36, 3, 41, 18},
	{1, 44, 10, 45, 2},
	{62, 6, 43, 15, 61},
	{28, 55, 25, 21, 56},
	{27, 20, 39, 8, 14},
}

// Hash256 computes the SHA3-256 hash of the given data and returns
// the 32-byte digest.
func Hash256(data []byte) []byte {
	return keccakHash(data, sha3_256Rate, 0x06, sha3_256HashLen)
}

// Hash512 computes the SHA3-512 hash of the given data and returns
// the 64-byte digest.
func Hash512(data []byte) []byte {
	return keccakHash(data, sha3_512Rate, 0x06, sha3_512HashLen)
}

// Shake256 computes the SHAKE256 extendable-output function (XOF) of the
// given data and returns outLen bytes of output.
func Shake256(data []byte, outLen int) []byte {
	if outLen <= 0 {
		return []byte{}
	}
	return keccakHash(data, shake256Rate, 0x1F, outLen)
}

// keccakHash is a generic sponge-based Keccak/SHA3/SHAKE helper.
func keccakHash(data []byte, rate int, ds byte, outLen int) []byte {
	var state [stateSize]uint64
	idx := 0

	// Absorb phase
	for _, b := range data {
		lane := idx / 8
		shift := uint((idx % 8) * 8)
		state[lane] ^= uint64(b) << shift
		idx++
		if idx == rate {
			keccakF1600(&state)
			idx = 0
		}
	}

	// Padding: domain separation byte `ds` and final 0x80 bit
	{
		lane := idx / 8
		shift := uint((idx % 8) * 8)
		state[lane] ^= uint64(ds) << shift

		lastIdx := rate - 1
		lane = lastIdx / 8
		shift = uint((lastIdx % 8) * 8)
		state[lane] ^= uint64(0x80) << shift
	}

	keccakF1600(&state)

	// Squeeze phase (may span multiple blocks for XOF / large hashes)
	out := make([]byte, outLen)
	outIdx := 0

	for outIdx < outLen {
		for i := 0; i < rate && outIdx < outLen; i++ {
			lane := i / 8
			shift := uint((i % 8) * 8)
			out[outIdx] = byte((state[lane] >> shift) & 0xFF)
			outIdx++
		}
		if outIdx < outLen {
			keccakF1600(&state)
		}
	}

	return out
}

// keccakF1600 applies the Keccak-f[1600] permutation to the given state in place.
// The state a is a 5×5 array of 64-bit lanes, flattened into a 25-element slice,
// as specified in the Keccak/SHA-3 standard (see FIPS 202 and the Keccak reference).
func keccakF1600(a *[25]uint64) {
	var C, D [5]uint64
	var B [25]uint64

	for round := 0; round < numRounds; round++ {
		// θ step
		for x := 0; x < 5; x++ {
			C[x] = (*a)[x] ^ (*a)[x+5] ^ (*a)[x+10] ^ (*a)[x+15] ^ (*a)[x+20]
		}
		for x := 0; x < 5; x++ {
			D[x] = C[(x+4)%5] ^ rotl64(C[(x+1)%5], 1)
		}
		for x := 0; x < 5; x++ {
			for y := 0; y < 5; y++ {
				(*a)[x+5*y] ^= D[x]
			}
		}

		// ρ and π steps
		for x := 0; x < 5; x++ {
			for y := 0; y < 5; y++ {
				newX := y
				newY := (2*x + 3*y) % 5
				from := x + 5*y
				to := newX + 5*newY
				B[to] = rotl64((*a)[from], rotationOffsets[x][y])
			}
		}

		// χ step
		for y := 0; y < 5; y++ {
			for x := 0; x < 5; x++ {
				idx := x + 5*y
				(*a)[idx] = B[idx] ^ ((^B[(x+1)%5+5*y]) & B[(x+2)%5+5*y])
			}
		}

		// ι step
		(*a)[0] ^= roundConstants[round]
	}
}

// rotl64 performs a left rotation of the 64-bit value x by n bits.
func rotl64(x uint64, n uint) uint64 {
	if n == 0 {
		return x
	}
	return (x << n) | (x >> (64 - n))
}
