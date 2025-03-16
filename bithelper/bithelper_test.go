package bithelper

import (
	"testing"

	. "github.com/jnsoft/jngo/testhelper"
)

func TestBitHelper(t *testing.T) {

	t.Run("IsEven int16", func(t *testing.T) {
		var n_even int16 = 16
		var n_odd int16 = 15
		AssertTrue(t, IsEven(n_even))
		AssertFalse(t, IsEven(n_odd))
	})

	t.Run("IsEven uint64", func(t *testing.T) {
		var n_even uint64 = 1626348234
		var n_odd uint64 = 154356457
		AssertTrue(t, IsEven(n_even))
		AssertFalse(t, IsEven(n_odd))
	})

	t.Run("Leading zeros", func(t *testing.T) {
		testCases := []struct {
			input    uint32
			expected uint32
		}{
			{0, 32},         // All bits are zero, so 32 leading zeros
			{1, 31},         // Binary: 00000000000000000000000000000001
			{2, 30},         // Binary: 00000000000000000000000000000010
			{4, 29},         // Binary: 00000000000000000000000000000100
			{8, 28},         // Binary: 00000000000000000000000000001000
			{16, 27},        // Binary: 00000000000000000000000000010000
			{255, 24},       // Binary: 00000000000000000000000011111111
			{1024, 21},      // Binary: 00000000000000000000010000000000
			{32768, 16},     // Binary: 00000000000000001000000000000000
			{2147483648, 0}, // Binary: 10000000000000000000000000000000 (no leading zeros)
		}
		for _, tc := range testCases {
			result := Nlz(tc.input)
			AssertEqual(t, result, tc.expected)
		}
	})

	t.Run("Trailing zeros", func(t *testing.T) {
		testCases := []struct {
			input    uint32
			expected uint32
		}{
			{0, 32},          // All bits are zero, so 32 trailing zeros
			{1, 0},           // Binary: 00000000000000000000000000000001
			{2, 1},           // Binary: 00000000000000000000000000000010
			{4, 2},           // Binary: 00000000000000000000000000000100
			{8, 3},           // Binary: 00000000000000000000000000001000
			{16, 4},          // Binary: 00000000000000000000000000010000
			{1024, 10},       // Binary: 00000000000000000000010000000000
			{32768, 15},      // Binary: 00000000000000001000000000000000
			{2147483648, 31}, // Binary: 10000000000000000000000000000000 (no leading zeros)
		}
		for _, tc := range testCases {
			result := Ntz(tc.input)
			AssertEqual(t, result, tc.expected)
		}
	})

	t.Run("Test NextLexicographicPermutaion", func(t *testing.T) {
		testCases := []struct {
			input    uint32
			expected uint32
		}{
			//{0b00000001, 0b00000010}, // Simple case: next permutation
			//{0b00000010, 0b00000100}, // Simple case: next permutation
			//{0b00010000, 0b00100000}, // Next higher binary permutation
			//{0b111, 0b1011},          // 111 -> 1011
			//{0b1010, 0b1100}, // 1010 -> 1100
			{0b1000, 0b1001},         // 1000 -> 1001 (next lexicographic)
			{0b11111111, 0b11111111}, // Edge case where all bits are set (no higher permutation)
			//{0b10101010, 0b10101100}, // A pattern case
		}

		for _, tc := range testCases {
			result := NextLexicographicPermutaion(tc.input)
			println(ToBinaryString(result))
			AssertEqual(t, result, tc.expected)
		}
	})

	t.Run("String conversions", func(t *testing.T) {
		//in1 := ""
		var n_odd uint64 = 154356457
		//AssertTrue(t, IsEven(n_even))
		AssertFalse(t, IsEven(n_odd))
	})
}
