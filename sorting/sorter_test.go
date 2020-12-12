package sorting

import (
	"testing"
)

func TestIntSorter(t *testing.T) {
	testArray := []int{5,3,7,18,1,9,8,6,6,0}
	sorter := intSorter()

	if isInitialized(sorter) {
		sorted := sorter.Sort(testArray)
		if areDifferentSize(testArray, sorted) {
			t.Error("sorted array has invalid size")
		}
		for i := 0; i < len(sorted)-1; i++ {
			if isNotOrdered(sorted[i], sorted[i+1]) {
				t.Error("values not ordered")
			}
		}
	} else {
		t.Error("sorter not initialized")
	}

}

func areDifferentSize(array1 []int, array2 []int) bool {
	return len(array1) != len(array2)
}

func isInitialized(sorter IntSorter) bool {
	return !(sorter == nil)
}

func isNotOrdered(i int, i2 int) bool {
	return !(i<=i2)
}

func intSorter() IntSorter{
	return new(BubbleSorter)
}