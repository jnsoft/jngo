package sorting

import "github.com/jnsoft/jngo/misc"

// MergeSort and InsertionSort are stable -> preserves the relative order of equal keys in the array
// QuickSort is unstable -> does not preserve the relative order of equal keys in the array
// InsertionSort is slowest (between N and N^2) but uses the least extra space (1)
// QuickSort sorts in NlogN and uses logN extra space, but 3-way quicksort is the fastest and sorts in between N and NlogN depending on input distribution
// MergeSort sorts in NlogN and uses N extra space

const MERGE_CUTOFF = 7

func InsertionSort[T misc.Ordered](arr []T) {
	insertionSort(arr, 0, len(arr)-1)
}

func InsertionSortImmutable[T misc.Ordered](arr []T) []T {
	sortedArr := append([]T(nil), arr...)
	insertionSort(sortedArr, 0, len(sortedArr)-1)
	return sortedArr
}

func insertionSort[T misc.Ordered](arr []T, lo, hi int) {
	for i := lo; i <= hi; i++ {
		for j := i; j > lo && arr[j] < arr[j-1]; j-- {
			swap := arr[j]
			arr[j] = arr[j-1]
			arr[j-1] = swap
		}
	}
}

// sorting an array using an optimized version of mergesort (top-down)
func MergeSort[T misc.Ordered](arr []T) {
	aux := append([]T(nil), arr...)
	mergeSort(aux, arr, 0, len(arr)-1)
}

// ret[i] is the index of the ith smallest entry in arr
// or: if we sort the array, ret[i] gives the original position for sorted[i]
func MergeSortIndices[T misc.Ordered](arr []T) []int {
	N := len(arr)
	index := misc.Sequence(0, N-1, 1)
	aux := make([]int, N)
	mergeSort_aux(arr, index, aux, 0, N-1)
	return index
}

func mergeSort[T misc.Ordered](src, dst []T, lo, hi int) {
	// if hi <= lo {return}

	// Switching to insertion sort for small subarrays will improve the running time
	// of a typical mergesort implementation by 10 to 15 percent
	if hi <= lo+MERGE_CUTOFF {
		insertionSort(dst, lo, hi)
		return
	}

	mid := lo + (hi-lo)/2
	mergeSort(dst, src, lo, mid)
	mergeSort(dst, src, mid+1, hi)

	// test whether array is already in order, thus reducing running time by skipping call to merge
	if src[mid+1] >= src[mid] {
		copy(dst[lo:hi+1], src[lo:hi+1])
		return
	}
	merge(src, dst, lo, mid, hi)
}

// mergesort a[lo..hi] using auxiliary array aux[lo..hi]
func mergeSort_aux[T misc.Ordered](a []T, index, aux []int, lo, hi int) {
	if hi <= lo {
		return
	}
	mid := lo + (hi-lo)/2
	mergeSort_aux(a, index, aux, lo, mid)
	mergeSort_aux(a, index, aux, mid+1, hi)
	merge_aux(a, index, aux, lo, mid, hi)
}

// precondition: src[lo .. mid] and src[mid+1 .. hi] are sorted subarrays
func merge[T misc.Ordered](src, dst []T, lo, mid, hi int) {
	i := lo
	j := mid + 1
	for k := lo; k <= hi; k++ {
		if i > mid {
			dst[k] = src[j]
			j++
		} else if j > hi {
			dst[k] = src[i]
			i++
		} else if src[j] < src[i] {
			dst[k] = src[j] // to ensure stability
			j++
		} else {
			dst[k] = src[i]
			i++
		}
	}
}

func merge_aux[T misc.Ordered](a []T, index, aux []int, lo, mid, hi int) {
	// copy to aux[]
	for k := lo; k <= hi; k++ {
		aux[k] = index[k]
	}

	// merge back to a[]
	i := lo
	j := mid + 1
	for k := lo; k <= hi; k++ {
		if i > mid {
			index[k] = aux[j]
			j++
		} else if j > hi {
			index[k] = aux[i]
			i++
		} else if a[aux[j]] < a[aux[i]] {
			index[k] = aux[j]
			j++
		} else {
			index[k] = aux[i]
			i++
		}
	}
}

// optimized version of quicksort (using Bentley-McIlroy 3-way partitioning, Tukey's ninther, and cutoff to insertion sort)
func QuickSort[T misc.Ordered](arr []T) {
	qsort(arr, 0, len(arr)-1)
}

func qsort[T misc.Ordered](a []T, lo, hi int) {
	N := hi - lo + 1
	// cutoff to insertion sort
	if N <= MERGE_CUTOFF {
		insertionSort(a, lo, hi)
		return
	} else if N <= 40 { // use median-of-3 as partitioning element
		m := median3(a, lo, lo+N/2, hi)
		exch(a, m, lo)
	} else { // use Tukey ninther as partitioning element
		eps := N / 8
		mid := lo + N/2
		m1 := median3(a, lo, lo+eps, lo+eps+eps)
		m2 := median3(a, mid-eps, mid, mid+eps)
		m3 := median3(a, hi-eps-eps, hi-eps, hi)
		ninther := median3(a, m1, m2, m3)
		exch(a, ninther, lo)
	}

	// Bentley-McIlroy 3-way partitioning
	i := lo
	j := hi + 1
	p := lo
	q := hi + 1
	v := a[lo]
	for {
		i++
		for a[i] < v {
			if i == hi {
				break
			}
			i++
		}
		j--
		for v < a[j] {
			if j == lo {
				break
			}
			j--
		}

		// pointers cross
		if i == j && a[i] == v {
			p++
			exch(a, p, i)
		}
		if i >= j {
			break
		}

		exch(a, i, j)
		if a[i] == v {
			p++
			exch(a, p, i)
		}
		if a[j] == v {
			q--
			exch(a, q, j)
		}
	}

	i = j + 1
	for k := lo; k <= p; k++ {
		exch(a, k, j)
		j--
	}
	for k := hi; k >= q; k-- {
		exch(a, k, i)
		i++
	}

	qsort(a, lo, j)
	qsort(a, i, hi)
}

func exch[T misc.Ordered](a []T, i, j int) {
	swap := a[i]
	a[i] = a[j]
	a[j] = swap
}

func median3[T misc.Ordered](a []T, i, j, k int) int {
	if a[i] < a[j] {
		if a[j] < a[k] {
			return j
		} else if a[i] < a[k] {
			return k
		} else {
			return i
		}
	} else {
		if a[k] < a[j] {
			return j
		} else if a[k] < a[i] {
			return k
		} else {
			return i
		}
	}
}

func IsSorted[T misc.Ordered](arr []T) bool {
	return isSorted(arr, 0, len(arr)-1)
}

func isSorted[T misc.Ordered](a []T, lo, hi int) bool {
	for i := lo + 1; i <= hi; i++ {
		if a[i] < a[i-1] {
			return false
		}
	}
	return true
}
