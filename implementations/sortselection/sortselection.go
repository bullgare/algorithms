package sortselection

// Sort sorts given slice of integers with a selection sort algorythm
func Sort(in []int) {
	l := len(in)
	for i, num := range in {
		min := num
		jMin := i
		for j := i; j < l; j ++ {
			if in[j] < min {
				min = in[j]
				jMin = j
			}
		}

		if i != jMin {
			in[jMin] = in[i]
			in[i] = min
		}
	}
}
