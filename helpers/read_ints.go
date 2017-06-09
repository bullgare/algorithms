package helpers

import (
	"io/ioutil"
	"fmt"
	"regexp"
	"strconv"
)

var reInt = regexp.MustCompile(`(?m)^\d+$`)

func ReadIntsFromFile(filepath string) ([]int, error) {
	contents, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("Error reading file %s: %s", filepath, err)
	}

	matches := reInt.FindAll(contents, -1)
	ints := make([]int, 0, len(matches))
	for k, m := range matches {
		i, err := strconv.Atoi(string(m))
		if err != nil {
			return nil, fmt.Errorf("Error converting %s to int (line %d of %s)", m, k + 1, filepath)
		}
		ints = append(ints, i)
	}
	return ints, nil
}