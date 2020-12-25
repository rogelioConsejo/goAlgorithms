package sorting

type InsertionSorter struct {
}

func (s *InsertionSorter) Sort(input []int) (sorted []int) {
	sorted = []int{}
	for index, value := range input {
		if index == 0 {
			sorted = append(sorted, value)
		} else {
			if value < sorted[len(sorted)-1] {
				sorted = placeInOrder(sorted, value)
			} else {
				sorted = append(sorted, value)
			}
		}
	}

	return sorted
}

func placeInOrder(sorted []int, value int) []int {
	found := false
	for i := len(sorted) - 1; i >= 0; i-- {
		if sorted[i] <= value {
			cache := sorted[i+1:]
			sorted = append(sorted[0:i+1], value)
			sorted = append(sorted, cache...)
			found = true
			break
		}
	}
	if !found {
		sorted = append([]int{value}, sorted...)
	}
	return sorted
}
