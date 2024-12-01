package hex

import (
	"testing"

	. "github.com/jnsoft/jngo/testhelper"
)

func TestHex(t *testing.T) {
	str := "Hello, World!"

	t.Run("XOR", func(t *testing.T) {
		arr1 := []byte{1, 2, 3, 4, 5}
		arr2 := []byte{5, 4, 3, 2, 1}
		arr3 := []byte{4, 6, 0, 6, 4}
		result := XOR(arr1, arr2)
		CollectionAssertEqual(t, result, arr3)
	})

	t.Run("To/From Hex String", func(t *testing.T) {
		bytes := []byte(str)
		hex_str_lower := ToHexString(bytes, false)
		hex_str_upper := ToHexString(bytes, true)
		returned_bytes1, err := FromHexString(hex_str_lower)
		AssertTrue(t, err == nil)
		returned_bytes2, err := FromHexString(hex_str_upper)
		AssertTrue(t, err == nil)

		CollectionAssertEqual(t, returned_bytes1, bytes)
		CollectionAssertEqual(t, returned_bytes2, bytes)
	})
}
