package sha2

import "encoding/binary"

const SHA256_BLOCK_SIZE = 512

// constants [§4.2.2]

// initial hash value [§5.3.3]
// H(i) is the i:th hash value. H(0) is the initial hash value; H(N) is the final hash value and is used to determine the message digest
// For SHA-256, the initial hash value, H(0), shall consist of the following eight 32-bit words
var H0 = []uint{0x6a09e667, 0xbb67ae85, 0x3c6ef372, 0xa54ff53a, 0x510e527f, 0x9b05688c, 0x1f83d9ab, 0x5be0cd19}

// K: Constant value to be used for the iteration t of the hash computation
// These words represent the first thirty-two bits of the fractional parts of the cube roots of the first sixty-four prime numbers
var K = []uint{
	0x428a2f98, 0x71374491, 0xb5c0fbcf, 0xe9b5dba5, 0x3956c25b, 0x59f111f1, 0x923f82a4, 0xab1c5ed5,
	0xd807aa98, 0x12835b01, 0x243185be, 0x550c7dc3, 0x72be5d74, 0x80deb1fe, 0x9bdc06a7, 0xc19bf174,
	0xe49b69c1, 0xefbe4786, 0x0fc19dc6, 0x240ca1cc, 0x2de92c6f, 0x4a7484aa, 0x5cb0a9dc, 0x76f988da,
	0x983e5152, 0xa831c66d, 0xb00327c8, 0xbf597fc7, 0xc6e00bf3, 0xd5a79147, 0x06ca6351, 0x14292967,
	0x27b70a85, 0x2e1b2138, 0x4d2c6dfc, 0x53380d13, 0x650a7354, 0x766a0abb, 0x81c2c92e, 0x92722c85,
	0xa2bfe8a1, 0xa81a664b, 0xc24b8b70, 0xc76c51a3, 0xd192e819, 0xd6990624, 0xf40e3585, 0x106aa070,
	0x19a4c116, 0x1e376c08, 0x2748774c, 0x34b0bcb5, 0x391c0cb3, 0x4ed8aa4a, 0x5b9cca4f, 0x682e6ff3,
	0x748f82ee, 0x78a5636f, 0x84c87814, 0x8cc70208, 0x90befffa, 0xa4506ceb, 0xbef9a3f7, 0xc67178f2,
}

func Hash256(data []byte) []byte {
	l := len(data)
	H := make([]uint, 8)
	copy(H, H0[:8])

	ix := 0
	var block [64]byte

	// main loop
	for i := 0; i < l; i++ {
		block[ix] = data[i]
		ix++

		if ix == 64 { // 512 bit blocks
			sha256Transform(&H, block)
			ix = 0
		}
	}

	// padding

	if ix < 56 { // size will fit in this block
		block[ix] = 0x80 // hex 80 = binary 1000 0000 (append 1 followed by zeroes to data)
		ix++
		for ix < 56 {
			block[ix] = 0x00
			ix++
		}
	} else {
		block[ix] = 0x80 // hex 80 = binary 1000 0000 (append 1 followed by zeroes to data)
		ix++
		for ix < 63 {
			block[ix] = 0x00
			ix++
		}

		sha256Transform(&H, block)

		ix = 0
		block = [64]byte{}
	}

	// Get the length in bits and convert to bytes
	lengthInBits := uint64(l * 8)
	lenBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(lenBytes, lengthInBits)

	// Append the length bytes to the Block array
	for i := 0; i < len(lenBytes); i++ {
		block[63-i] = lenBytes[i]
	}

	sha256Transform(&H, block)

	return getHashBytes(H)
}

func sha256Transform(H *[]uint, data [64]byte) {
	var a, b, c, d, e, f, g, h, i, j, t1, t2 uint
	var m [64]uint

	for i, j = 0, 0; i < 16; i, j = i+1, j+4 {
		m[i] = uint(data[j])<<24 | uint(data[j+1])<<16 | uint(data[j+2])<<8 | uint(data[j+3])
	}

	for i := 16; i < 64; i++ {
		m[i] = σ1(m[i-2]) + m[i-7] + σ0(m[i-15]) + m[i-16]
	}

	a = (*H)[0]
	b = (*H)[1]
	c = (*H)[2]
	d = (*H)[3]
	e = (*H)[4]
	f = (*H)[5]
	g = (*H)[6]
	h = (*H)[7]

	for i = 0; i < 64; i++ {

		t1 = h + Σ1(e) + ch(e, f, g) + K[i] + m[i]
		t2 = Σ0(a) + maj(a, b, c)
		h = g
		g = f
		f = e
		e = d + t1
		d = c
		c = b
		b = a
		a = t1 + t2
	}

	(*H)[0] += a
	(*H)[1] += b
	(*H)[2] += c
	(*H)[3] += d
	(*H)[4] += e
	(*H)[5] += f
	(*H)[6] += g
	(*H)[7] += h
}

// Rotates right (circular right shift) value x by n positions [§3.2.4].
func rotr(x uint, n uint8) uint {
	return (x >> n) | (x << (32 - n))
}

// Logical functions [§4.1.2].
func Σ0(x uint) uint { // lsigma0
	return rotr(x, 2) ^ rotr(x, 13) ^ rotr(x, 22)
}
func Σ1(x uint) uint { //lsigma1
	return rotr(x, 6) ^ rotr(x, 11) ^ rotr(x, 25)
}
func σ0(x uint) uint { // ssigma0
	return rotr(x, 7) ^ rotr(x, 18) ^ (x >> 3)
}
func σ1(x uint) uint { // ssigma1
	return rotr(x, 17) ^ rotr(x, 19) ^ (x >> 10)
}
func ch(x, y, z uint) uint { // 'choice'
	return (x & y) ^ (^x & z)
}
func maj(x, y, z uint) uint { // 'majority'
	return (x & y) ^ (x & z) ^ (y & z)
}

// GetHashBytes converts an array of uint32 to a byte array
func getHashBytes(H []uint) []byte {
	hash := make([]byte, 32)
	for i := 0; i < 4; i++ {
		hash[i] = byte((H[0] >> (24 - i*8)) & 0x000000ff)
		hash[i+4] = byte((H[1] >> (24 - i*8)) & 0x000000ff)
		hash[i+8] = byte((H[2] >> (24 - i*8)) & 0x000000ff)
		hash[i+12] = byte((H[3] >> (24 - i*8)) & 0x000000ff)
		hash[i+16] = byte((H[4] >> (24 - i*8)) & 0x000000ff)
		hash[i+20] = byte((H[5] >> (24 - i*8)) & 0x000000ff)
		hash[i+24] = byte((H[6] >> (24 - i*8)) & 0x000000ff)
		hash[i+28] = byte((H[7] >> (24 - i*8)) & 0x000000ff)
	}
	return hash
}
