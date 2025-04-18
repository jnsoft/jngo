package stringhelper

import (
	"math"
	"sort"
	"strings"
)

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func Sort(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

func IsPalindrome(s string) bool {
	normalized := strings.ToLower(strings.ReplaceAll(s, " ", ""))
	for i := 0; i < len(normalized)/2; i++ {
		if normalized[i] != normalized[len(normalized)-1-i] {
			return false
		}
	}
	return true
}

// The Hamming Distance measures the minimum number of substitutions required to change one string into the other
func HammingDistance(s1, s2 string) int {
	if len(s1) != len(s2) {
		return -1
	}

	distance := 0
	for i := range s1 {
		if s1[i] != s2[i] {
			distance++
		}
	}
	return distance
}

func EditDistance(s1, s2 string) int {
	n := len(s1)
	m := len(s2)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			if i == 0 {
				dp[i][j] = j // First row: All insertions
			} else if j == 0 {
				dp[i][j] = i // First column: All deletions
			} else if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] // No operation required
			} else {
				deletion := dp[i-1][j]
				insertion := dp[i][j-1]
				substitution := dp[i-1][j-1]
				dp[i][j] = 1 + int(math.Min(math.Min(float64(deletion), float64(insertion)), float64(substitution)))
			}
		}
	}

	// The bottom-right cell contains the result
	return dp[n][m]
}

func SplitStrings(input []string, delimiter string) [][]string {
	var result [][]string
	for _, str := range input {
		splitStr := strings.Split(str, delimiter)
		result = append(result, splitStr)
	}
	return result
}

func ToLines(str string) []string {
	return strings.Split(strings.ReplaceAll(str, "\r\n", "\n"), "\n")
}

// next lexicographically greater permutation of a word
func NextPermutation(s *string) (bool, string) {
	runes := []rune(*s)
	n := len(runes)
	// Find the largest index i such that runes[i] < runes[i+1]
	i := n - 2
	for i >= 0 && runes[i] >= runes[i+1] {
		i--
	}
	if i < 0 { // No such index
		return false, string(runes)
	} else {
		// Find the largest index j such that runes[i] < runes[j]
		ix := bsearch(runes, i+1, n-1, runes[i])

		// Swap runes[i] and runes[ix]
		swap(&runes[i], &runes[ix])

		// Reverse the sublist runes[start:end+1]
		rev(runes, i+1, n-1)

		return true, string(runes)
	}
}

func bsearch(s []rune, l, r int, key rune) int {
	ix := -1
	for l <= r {
		mid := l + (r-l)/2
		if s[mid] <= key {
			r = mid - 1
		} else {
			l = mid + 1
			if ix == -1 || s[ix] >= s[mid] {
				ix = mid
			}
		}
	}
	return ix
}

func rev(s []rune, l, r int) {
	for l < r {
		swap(&s[l], &s[r])
		l++
		r--
	}
}

func swap(a, b *rune) {
	if *a == *b {
		return
	}
	*a ^= *b
	*b ^= *a
	*a ^= *b
}
