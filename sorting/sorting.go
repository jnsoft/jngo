package sorting

import "github.com/jnsoft/jngo/misc"

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
