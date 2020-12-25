package sorting

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)


//I extracted performTest() to be able to test all the sorting implementations in one test
func TestIntSorter(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	//Test with a big pseudo-random array
	println("Bubble")
	performTest(t,bubbleSorter(), rand.Perm(100000))			// O(n^2)
	println("Insertion")
	performTest(t,insertionSorter(), rand.Perm(100000))			// O(n^2)
	println("Merge")
	performTest(t,mergeSorter(), rand.Perm(100000))			// O(n log n)
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
			if areNotOrdered(sorted[i], sorted[i+1]) {
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

func areNotOrdered(i int, i2 int) bool {
	return !(i<=i2)
}

func bubbleSorter() IntSorter{
	return new(BubbleSorter)
}

func insertionSorter() IntSorter{
	return new(InsertionSorter)
}

func mergeSorter() IntSorter{
	return new(MergeSorter)
}