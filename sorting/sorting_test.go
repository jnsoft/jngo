package sorting

import (
	"testing"

	. "github.com/jnsoft/jngo/testhelper"
)

func TestSorting(t *testing.T) {

	var intArray = []int{5, -3, 8, 1, 2, -7, 6, 4, 0}
	var floatArray = []float64{3.14, -1.5, 2.71, 0.0, -3.14, 1.618, 42.0}
	var stringArray = []string{"banana", "apple", "grape", "orange", "cherry", "blueberry"}

	var intArray_sorted = []int{-7, -3, 0, 1, 2, 4, 5, 6, 8}
	var floatArray_sorted = []float64{-3.14, -1.5, 0, 1.618, 2.71, 3.14, 42}
	var stringArray_sorted = []string{"apple", "banana", "blueberry", "cherry", "grape", "orange"}

	unsorted := []int{
		93, 27, 84, 12, 75, 3, 61, 48, 89, 20,
		7, 95, 31, 69, 41, 16, 55, 74, 68, 24,
		87, 1, 36, 57, 92, 23, 80, 47, 10, 62,
		4, 53, 33, 76, 50, 45, 6, 26, 30, 97,
		66, 43, 90, 19, 2, 58, 70, 14, 79, 39,
		17, 60, 9, 82, 40, 29, 35, 96, 56, 8,
		49, 98, 34, 21, 72, 11, 42, 59, 32, 63,
		22, 54, 13, 64, 88, 77, 91, 5, 85, 67,
		51, 86, 18, 28, 46, 37, 78, 71, 38, 15,
		44, 65, 25, 52, 81, 100, 83, 73, 94, 99,
	}

	sorted := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
		21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
		31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
		41, 42, 43, 44, 45, 46, 47, 48, 49, 50,
		51, 52, 53, 54, 55, 56, 57, 58, 59, 60,
		61, 62, 63, 64, 65, 66, 67, 68, 69, 70,
		71, 72, 73, 74, 75, 76, 77, 78, 79, 80,
		81, 82, 83, 84, 85, 86, 87, 88, 89, 90,
		91, 92, 93, 94, 95, 96, 97, 98, 99, 100,
	}

	t.Run("Insertion sort", func(t *testing.T) {

		t1 := InsertionSortImmutable(intArray)
		t2 := InsertionSortImmutable(floatArray)
		t3 := InsertionSortImmutable(stringArray)
		t4 := InsertionSortImmutable(unsorted)

		CollectionAssertEqual(t, t1, intArray_sorted)
		CollectionAssertEqual(t, t2, floatArray_sorted)
		CollectionAssertEqual(t, t3, stringArray_sorted)
		CollectionAssertEqual(t, t4, sorted)
	})

	t.Run("Merge sort", func(t *testing.T) {

		ints := append([]int(nil), unsorted...)
		MergeSort(ints)
		CollectionAssertEqual(t, ints, sorted)
	})

	t.Run("Quick sort", func(t *testing.T) {

		ints := append([]int(nil), unsorted...)
		QuickSort(ints)
		CollectionAssertEqual(t, ints, sorted)
	})

}
