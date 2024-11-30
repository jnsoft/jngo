package hex

import (
	"testing"

	. "github.com/jnsoft/jngo/testhelper"
)

func TestHex(t *testing.T) {
	str := "Hello, World!"
	t.Run("To/From Hex String", func(t *testing.T) {
		bytes := []byte(str)
		hex_str_lower := ToHexString(bytes, true)
		hex_str_upper := ToHexString(bytes, false)
		returned_bytes1, err := FromHexString(hex_str_lower)
		AssertTrue(t, err == nil)
		returned_bytes2, err := FromHexString(hex_str_upper)
		AssertTrue(t, err == nil)

		CollectionAssertEqual(t, returned_bytes1, bytes)
		CollectionAssertEqual(t, returned_bytes2, bytes)

	})
}
