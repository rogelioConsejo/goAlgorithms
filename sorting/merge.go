package sorting

type MergeSorter struct {

}

func (s *MergeSorter) Sort(input []int) (sorted []int)  {
	if len(input) <= 1 {
		return input
	} else {
		firstHalf := input[0:(len(input)/2)]
		secondHalf := input[(len(input)/2):]
		s := new(MergeSorter)
		firstHalf = s.Sort(firstHalf)
		secondHalf = s.Sort(secondHalf)
		sorted = mergeSort(firstHalf, secondHalf)
	}
	return
}

func mergeSort(firstHalf []int, secondHalf []int) []int {
	p1 := 0
	p2 := 0
	var result []int
	for p1 < len(firstHalf) || p2 < len(secondHalf) {
		if p1 >= len(firstHalf) {
			result, p2 = addToMergeFrom(result, secondHalf, p2)
		} else if p2 >= len(secondHalf) {
			result, p1 = addToMergeFrom(result, firstHalf, p1)
		} else {
			if firstHalf[p1] < secondHalf[p2] {
				result, p1 = addToMergeFrom(result, firstHalf, p1)
			} else {
				result, p2 = addToMergeFrom(result, secondHalf, p2)
			}
		}
	}
	return result
}

func addToMergeFrom(merge []int, arr []int, pointer int) ([]int, int) {
	merge = append(merge, arr[pointer])
	pointer++
	return merge, pointer
}

