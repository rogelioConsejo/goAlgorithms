package sorting

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)


//I extracted performTest() to be able to test all the sorting implementations in one test
func TestIntSorter(t *testing.T) {
	//Test with a big pseudo-random array
	performTest(t,bubbleSorter(), rand.Perm(100000))			// O(n^2)
}


func performTest(t *testing.T, sorter IntSorter, testArray []int){


	if isInitialized(sorter) {
		start := time.Now()
		sorted := sorter.Sort(testArray)
		fmt.Printf("Time: %s\n", time.Since(start))
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