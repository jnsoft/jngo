package hmac

import (
	"github.com/jnsoft/jngo/hex"
	"github.com/jnsoft/jngo/sha2"
)

// RFC 2104

const (
	OUTER_PAD_CHAR = 0x5c
	INNER_PAD_CHAR = 0x36
)

// HMCAC(K,m) = H( (K' x opad) || H((K' x ipad) || m) )
// K' = len(K) > blocksize ? H(K) : K
// || = concat, x = bitwise xor
// opad = block size length repeated values of 0x5c
// ipad = block size length repeated values of 0x36

func Compute(key, msg []byte) []byte {
	key = setKey(key)
	outerKey := keyXorPad(key, OUTER_PAD_CHAR, sha2.SHA256_BLOCK_SIZE/8)
	innerKey := keyXorPad(key, INNER_PAD_CHAR, sha2.SHA256_BLOCK_SIZE/8)
	return sha2.Hash256(concat(outerKey, sha2.Hash256(concat(innerKey, msg))))
}

func Verify(key, msg, mac []byte) bool {
	digest := Compute(key, msg)
	for i := 0; i < len(digest); i++ {
		if digest[i] != mac[i] {
			return false
		}
	}
	return true
}

func concat(left, right []byte) []byte {
	arr := make([]byte, len(left)+len(right))
	copy(arr[0:], left[0:len(left)])
	copy(arr[len(left):], right[0:len(right)])
	return arr
}

func setKey(key []byte) []byte {
	if len(key) > sha2.SHA256_BLOCK_SIZE/8 {
		return sha2.Hash256(key)
	} else if len(key) < sha2.SHA256_BLOCK_SIZE/8 {
		return zeroPad(key, sha2.SHA256_BLOCK_SIZE/8)
	}
	return key
}

func keyXorPad(key []byte, padChar byte, blockSize int) []byte {
	return hex.XOR(key, getPad(blockSize, padChar))
}

func getPad(blockSize int, padChar byte) []byte {
	pad := make([]byte, blockSize)
	for i := 0; i < len(pad); i++ {
		pad[i] = padChar
	}
	return pad
}

func zeroPad(data []byte, blockSize int) []byte {
	pad := blockSize - len(data)%blockSize
	newLen := len(data) + pad
	block := make([]byte, newLen)
	copy(block, data)
	return block
}
