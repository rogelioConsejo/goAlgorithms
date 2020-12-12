package sorting

import "fmt"

type BubbleSorter struct {
}

func (b *BubbleSorter) Sort(input []int) (sorted []int) {
	for i:=0; i<len(input)-1 ; i++{
		for index, value := range input {
			if index < len(input)-1 && value > input[index+1] {
				input[index] = input[index+1]
				input[index+1] = value
			}
		}
	}

	sorted = input
	fmt.Printf("%+v", input)
	return
}
