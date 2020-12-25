package sorting

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"
	"time"
)

const (
	bubble = "Bubble"
	insertion = "Insertion"
	merge = "Merge"
)

var sorters = map[string]IntSorter{
	bubble: new(BubbleSorter),
	insertion: new(InsertionSorter),
	merge: new(MergeSorter),
}


//I extracted performTest() to be able to test all the sorting implementations in one test
func TestIntSorter(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	//Test with a big array
	testArray := rand.Perm(100000)
	println(bubble)
	performTest(t,sorters[bubble], append([]int{},testArray...))			// O(n^2)
	println(insertion)
	performTest(t,sorters[insertion], append([]int{},testArray...))			// O(n^2)
	println(merge)
	performTest(t,sorters[merge], append([]int{},testArray...))			// O(n log n)
}


func performTest(t *testing.T, sorter IntSorter, testArray []int){
	if isInitialized(sorter) {
		start := time.Now()
		sorted := sorter.Sort(testArray)
		m := &runtime.MemStats{}
		runtime.ReadMemStats(m)
		fmt.Printf("Mem (MB): %f\n", float64(m.Alloc)/1024/1024)
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

