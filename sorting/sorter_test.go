package sorting

import (
	"testing"
)


//I extracted performTest() to be able to test all the sorting implementations in one test
func TestIntSorter(t *testing.T) {
	performTest(t,bubbleSorter(), []int{5,3,7,18,1,9,8,6,6,0})
	performTest(t,bubbleSorter(), []int{5,3,7,4,1,9,8,6,6,18})
	performTest(t,bubbleSorter(), []int{})
	performTest(t,bubbleSorter(), []int{1})
	performTest(t,bubbleSorter(), []int{1,1})
	performTest(t,bubbleSorter(), []int{1,2})
	performTest(t,bubbleSorter(), []int{2,1})
}

func performTest(t *testing.T, sorter IntSorter, testArray []int){

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

func bubbleSorter() IntSorter{
	return new(BubbleSorter)
}