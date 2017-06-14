package insertion

// Sort sorts given slice of integers with a insertion sort algorithm
func Sort(in []int) {
	for index, num := range in {
		if index == 0 {
			continue
		}

		for i := index - 1; i >= 0; i -- {
			if in[i] > num {
				in[i + 1] = in[i]
				in[i] = num
			} else {
				break
			}
		}
	}
}