package stringhelper

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
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
