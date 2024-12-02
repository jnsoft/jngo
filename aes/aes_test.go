package aes

import (
	"testing"

	"github.com/jnsoft/jngo/hex"
	. "github.com/jnsoft/jngo/testhelper"
)

func TestAES(t *testing.T) {
	plaintext := []byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff}
	key := []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x10, 0x11,
		0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f}
	ciphertext := []byte{0x8e, 0xa2, 0xb7, 0xca, 0x51, 0x67, 0x45, 0xbf, 0xea, 0xfc, 0x49, 0x90, 0x4b, 0x49, 0x60, 0x89}

	t.Run("AES Encrypt", func(t *testing.T) {
		ciphertext_out := make([]byte, 16)
		key_schedule := make([]uint32, 60)
		encrypted := "8ea2b7ca516745bfeafc49904b496089"

		KeyExpansion(key, key_schedule, 256)
		Encrypt(plaintext, ciphertext_out, key_schedule, 256)
		hex_str := hex.ToHexString(ciphertext_out, false)
		AssertEqual(t, hex_str, encrypted)
	})

	t.Run("AES Decrypt", func(t *testing.T) {
		plaintext_out := make([]byte, 16)
		key_schedule := make([]uint32, 60)
		decrypted := "00112233445566778899aabbccddeeff"

		KeyExpansion(key, key_schedule, 256)
		Decrypt(ciphertext, plaintext_out, key_schedule, 256)
		hex_str := hex.ToHexString(plaintext_out, false)
		AssertEqual(t, hex_str, decrypted)
	})
}
