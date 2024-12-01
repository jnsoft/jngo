package sha2

import (
	"crypto/sha256"
	"testing"

	"github.com/jnsoft/jngo/hex"
	. "github.com/jnsoft/jngo/testhelper"
)

func TestSha2(t *testing.T) {
	str := "Hello, World!"
	hash_hexstring := "dffd6021bb2bd5b0af676290809ec3a53191dd81c7f70a4b28688a362182986f"
	t.Run("SHA256", func(t *testing.T) {
		bytes := []byte(str)
		hash := Hash256(bytes)
		hash_str := hex.ToHexString(hash, true)

		hash2 := sha256.New()
		hash2.Write([]byte(str))
		hashBytes := hash2.Sum(nil)
		hash_str2 := hex.ToHexString(hashBytes, true)

		AssertEqual(t, hash_str, hash_hexstring)
		AssertEqual(t, hash_str2, hash_hexstring)
	})
}
