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

func QuickSort[T misc.Ordered](arr []T) {
	if len(arr) < 2 {
		return
	}
	// Choose a pivot element (typically the last element)
	pivotIndex := len(arr) - 1
	pivot := arr[pivotIndex]

	// Partition the array into two halves
	left := 0
	right := pivotIndex - 1

	for left <= right {
		// Move left pointer to the right until we find an element >= pivot
		for arr[left] < pivot {
			left++
		}

		// Move right pointer to the left until we find an element <= pivot
		for arr[right] > pivot {
			right--
		}

		// Swap elements if they are in the wrong order
		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}
	QuickSort(arr[:right+1])
	QuickSort(arr[left:])
}
