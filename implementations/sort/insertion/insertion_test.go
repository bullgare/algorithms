package insertion_test

import (
	"testing"

	"github.com/bullgare/algorithms/implementations/sort/insertion"
	"github.com/bullgare/algorithms/helpers"
)

func TestEmpty(t *testing.T) {
	var ints []int = nil
	insertion.Sort(ints)

	if len(ints) > 0 {
		t.Error("Expected empty slice on empty input")
	}
}

func Test1Element(t *testing.T) {
	ints := []int{1}
	insertion.Sort(ints)

	if len(ints) != 1 {
		t.Error("Expected slice with one element")
	}

	if ints[0] != 1 {
		t.Error("Expected res[0]=1, got ", ints[0])
	}
}

func Test2Elements(t *testing.T) {
	ints := []int{2, 1}
	insertion.Sort(ints)

	if len(ints) != 2 {
		t.Error("Expected slice with 2 elements")
	}

	if ints[0] != 1 {
		t.Error("Expected res[0]=1, got ", ints[0])
	}
	if ints[1] != 2 {
		t.Error("Expected res[1]=2, got ", ints[1])
	}
}

func Test30Elements(t *testing.T) {
	ints, err := helpers.ReadIntsFromFile("../../../data/30.txt")
	if err != nil {
		t.Error("Error reading file")
	}

	insertion.Sort(ints)

	if len(ints) != 30 {
		t.Errorf("Expected slice with 30 elements, got %d", len(ints))
	}
	if err = helpers.CheckSliceIsSorted(ints); err != nil {
		t.Error(err)
	}
}
