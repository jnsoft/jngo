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
		hash_str := hex.ToHexString(hash, false)

		hash2 := sha256.New()
		hash2.Write([]byte(str))
		hashBytes := hash2.Sum(nil)
		hash_str2 := hex.ToHexString(hashBytes, false)

		AssertEqual(t, hash_str, hash_hexstring)
		AssertEqual(t, hash_str2, hash_hexstring)
	})

	t.Run("SHA256_nist_test_vector_1", func(t *testing.T) {
		abc := "abc"
		abcHash := "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad"

		abcBytes := []byte(abc)
		abcHashBytes := Hash256(abcBytes)
		abcHashBytesHexString := hex.ToHexString(abcHashBytes, false)

		AssertEqual(t, abcHashBytesHexString, abcHash)
	})

	t.Run("SHA256_nist_test_vector_2", func(t *testing.T) {
		s := "abcdbcdecdefdefgefghfghighijhijkijkljklmklmnlmnomnopnopq"
		hash := "248D6A61D20638B8E5C026930C3E6039A33CE45964FF2167F6ECEDD419DB06C1"

		data := []byte(s)
		hashbytes := Hash256(data)
		hex := hex.ToHexString(hashbytes, true)

		AssertEqual(t, hex, hash)
	})

}
