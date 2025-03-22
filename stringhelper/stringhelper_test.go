package stringhelper

import (
	"testing"

	. "github.com/jnsoft/jngo/testhelper"
)

func TestStringHelper(t *testing.T) {
	t.Run("Test reverse", func(t *testing.T) {
		str := "I am a string!"
		AssertEqual(t, Reverse(Reverse(str)), str)
	})

	t.Run("Test NextPermutation", func(t *testing.T) {
		tests := []struct {
			input    string
			expected string
			success  bool
		}{
			{"123", "132", true},
			{"321", "321", false},
			{"a", "a", false},
			{"ab", "ba", true},
			{"aabb", "abab", true},
			{"edcba", "edcba", false},
			{"abcde", "abced", true},
			{"abedc", "acbde", true},
			{"", "", false},
		}
		for _, test := range tests {
			success, result := NextPermutation(&test.input)
			AssertEqual(t, result, test.expected)
			AssertEqual(t, success, test.success)
		}
	})

}
