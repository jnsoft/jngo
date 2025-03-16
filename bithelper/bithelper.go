package bithelper

import (
	"reflect"
	"strconv"
	"strings"
)

type UnsignedInteger interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type IntegerNumber interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | UnsignedInteger
}

func IsEven[T IntegerNumber](n T) bool {
	return n&1 == 0
}

// number of leading zeros
func Nlz[T IntegerNumber](n T) T {
	var y, x T
	x = 0
	y = n
	for {
		if n < 0 {
			return x
		}

		if y == 0 {
			return 32 - x
		}
		x = x + 1
		n = n << 1
		y = y >> 1
	}
}

// number of trailing zeros
func Ntz[T IntegerNumber](n T) T {
	var x T
	n = ^n & (n - 1)
	x = 0 // x = 32
	for n != 0 {
		x = x + 1
		n = n >> 1 // n = n + n;
	}
	return x // return x
}

// set(0000,1) -> 0010
func SetBit[T IntegerNumber](n T, pos int) T {
	return (1 << pos) | n
}

// clear(0010,1) -> 0000
func ClearBit[T IntegerNumber](n T, pos int) T {
	return ^(1 << pos) & n
}

// toggle(0010,2) -> 0110
func ToggleBit[T IntegerNumber](n T, pos int) T {
	return (1 << pos) ^ n
}

// conv(1100) -> 1111
func ConvertTralingZeros[T UnsignedInteger](n T) T {
	return (n - 1) | n
}

// Extract (isolate) least significant bit
func LSB[T UnsignedInteger](n T) T {
	return n & -n // -n = (~n)+1 in 2-complement rep.
}

// copy bits from b into a where mask = 1
func MaskedCopy[T IntegerNumber](a T, b T, m T) T {
	return (b & m) | (a & ^m)
}

func SwapBits[T IntegerNumber](n T, a T, b T) T {
	p := (n >> a) ^ (n>>b)&1
	r := n ^ (p << a)
	return r ^ (p << b)
}

// Kerninghan's population count, count number of bits set to 1
func PopulationCount[T IntegerNumber](n T) T {
	var c T
	for c = 0; n != 0; c++ {
		n &= (n - 1) // n & (n - 1) zeros the lowest one bit
	}
	return c
}

// 11001010 -> 3
func BitIslandsCount[T IntegerNumber](n T) T {
	return (n & 1) + PopulationCount(n^(n>>1))/2
}

// find index of the least significant bit (x86 has instructions BSF for this)
func BitScanForwards(n uint32) int {
	if n == 0 {
		return -1 // No bits set
	}

	pos := 0
	for n&1 == 0 {
		pos++
		n >>= 1
	}
	return pos
}

func BitScanForwards2(n uint32) int {
	if n == 0 {
		return -1
	}
	n = n & -n
	var c int = 0
	if (n & 0xffff0000) != 0 {
		c += 16
	}
	if (n & 0xff00ff00) != 0 {
		c += 8
	}
	if (n & 0xf0f0f0f0) != 0 {
		c += 4
	}
	if (n & 0xcccccccc) != 0 {
		c += 2
	}
	if (n & 0xaaaaaaaa) != 0 {
		c += 1
	}
	return c
}

// find next larger int with same number of bits set, 01001100 -> 01010001 -> 01010010 -> 01010100 ...
func NextLexicographicPermutaion(n uint32) uint32 {
	// Step 1: Identify the rightmost 0 that is followed by 1
	c := n & -n                     // Isolate the rightmost 1-bit
	r := n + c                      // Add c to the number to get the next larger number
	return (((n ^ r) / c) >> 1) | r // Rearrange bits to the right of the flipped position
}

func NextLexicographicPermutaion2(n uint32) uint32 {
	t := ConvertTralingZeros(n)
	return (t + 1) | (LSB(^t)-1)>>(BitScanForwards(n)+1)
}

func bitCount(n uint32) int {
	count := 0
	for n > 0 {
		count++
		n >>= 1
	}
	return count
}

func ToBinaryString[T IntegerNumber](i T) string {
	return strconv.FormatInt(int64(i), 2)
}

func FromBinaryString[T IntegerNumber](b string) (T, error) {
	var zero T
	bitSize := int(reflect.TypeOf(zero).Size() * 8)

	i, err := strconv.ParseInt(b, 2, bitSize)
	return T(i), err
}

func ToBinary[T IntegerNumber](i T) string {
	if i == 0 {
		return "0"
	}
	if i == 1 {
		return "1"
	}

	divisor := i / 2
	remainder := i % 2

	return ToBinary(divisor) + strconv.Itoa(int(remainder))
}

func Pad(s string, length int, pad rune) string {
	if len(s) >= length {
		return s
	}
	return strings.Repeat(string(pad), length-len(s)) + s
}
