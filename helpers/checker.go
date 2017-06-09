package helpers

import "fmt"

func CheckSliceIsSorted(ints []int) error {
	max := -1

	for index, i := range ints {
		if i < max {
			return fmt.Errorf("Element %d (index %d) is not sorted", i, index)
		}
		max = i
	}

	return nil
}
