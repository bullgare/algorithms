package sortinsertion

// Sort sorts given slice of integers with a insertion sort algorithm
func Sort(in []int) {
	for index, num := range in {
		if index == 0 {
			continue
		}

		for i, before := range in[:index - 1] {
			if before > num {
				in[index] = before
				in[i] = num
				break
			}
		}
	}
}