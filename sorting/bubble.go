package sorting

type BubbleSorter struct {
}

//sorted is only used to make the method signature clearer
func (b *BubbleSorter) Sort(input []int) (sorted []int) {
	for i:=0; i<len(input)-1 ; i++{
		isSorted := true
		for index, value := range input {
			if index < len(input)-1 && value > input[index+1] {
				input[index] = input[index+1]
				input[index+1] = value
				isSorted = false
			}
		}
		if isSorted{
			break
		}
	}

	return input
}
