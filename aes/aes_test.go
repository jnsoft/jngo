package aes

import (
	"testing"

	"github.com/jnsoft/jngo/hex"
	. "github.com/jnsoft/jngo/testhelper"
)

var NIST_PLAINTEXT = []byte{0x6b, 0xc1, 0xbe, 0xe2, 0x2e, 0x40, 0x9f, 0x96, 0xe9, 0x3d, 0x7e, 0x11, 0x73, 0x93, 0x17, 0x2a,
	0xae, 0x2d, 0x8a, 0x57, 0x1e, 0x03, 0xac, 0x9c, 0x9e, 0xb7, 0x6f, 0xac, 0x45, 0xaf, 0x8e, 0x51,
	0x30, 0xc8, 0x1c, 0x46, 0xa3, 0x5c, 0xe4, 0x11, 0xe5, 0xfb, 0xc1, 0x19, 0x1a, 0x0a, 0x52, 0xef,
	0xf6, 0x9f, 0x24, 0x45, 0xdf, 0x4f, 0x9b, 0x17, 0xad, 0x2b, 0x41, 0x7b, 0xe6, 0x6c, 0x37, 0x10}

var KEY = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
	0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f}

func TestAES(t *testing.T) {
	plaintext := []byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff}
	key := []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x10, 0x11,
		0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f}
	ciphertext := []byte{0x8e, 0xa2, 0xb7, 0xca, 0x51, 0x67, 0x45, 0xbf, 0xea, 0xfc, 0x49, 0x90, 0x4b, 0x49, 0x60, 0x89}

	t.Run("AES Cipher", func(t *testing.T) {
		ciphertext_out := make([]byte, 16)
		encrypted := "8ea2b7ca516745bfeafc49904b496089"

		key_schedule := KeyExpansion(key)
		Cipher(plaintext, ciphertext_out, key_schedule)
		hex_str := hex.ToHexString(ciphertext_out, false)
		AssertEqual(t, hex_str, encrypted)
	})

	t.Run("AES Decipher", func(t *testing.T) {
		plaintext_out := make([]byte, 16)
		decrypted := "00112233445566778899aabbccddeeff"

		key_schedule := KeyExpansion(key)
		Decipher(ciphertext, plaintext_out, key_schedule)
		hex_str := hex.ToHexString(plaintext_out, false)
		AssertEqual(t, hex_str, decrypted)
	})
}

func TestECR(t *testing.T) {
	t.Run("ECR Encode / Decode", func(t *testing.T) {
		ciphertext := ECB_Encrypt(NIST_PLAINTEXT, KEY)
		decrypted := ECB_Decrypt(ciphertext, KEY)
		CollectionAssertEqual(t, decrypted, NIST_PLAINTEXT)
	})
}

func TestCBC(t *testing.T) {
	t.Run("CBC Encode / Decode", func(t *testing.T) {
		IV := make([]byte, 16)
		ciphertext := CBC_Encrypt(NIST_PLAINTEXT, KEY, IV)
		decrypted := CBC_Decrypt(ciphertext, KEY, IV)
		CollectionAssertEqual(t, decrypted, NIST_PLAINTEXT)
	})
}
