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
				found := false
				for i := len(sorted)-1 ; i >= 0; i-- {
					if sorted[i] <= value {
						var cache []int
						cache = append([]int{}, sorted...)
						sorted = append(sorted[0:i+1],value)
						sorted = append(sorted, cache[i+1:]...)
						found = true
						break
					}
				}
				if !found{
					sorted = append([]int{value}, sorted...)
				}

			} else {
				sorted = append(sorted, value)
			}
		}
	}

	return sorted
}
