package utils

import "testing"

func TestIndexOfRune(t *testing.T) {
	r := []rune{'a', 'b', 'c', 'd', 'e'}

	index := IndexOfRune(r, 'c')

	if index != 2 {
		t.Errorf("Expected index of 'c' to be 2, but got %d", index)
	}
}

func TestGetMostRepeated(t *testing.T) {
	array := []int{1, 2, 3, 4, 5, 5, 5, 6, 7, 8, 9, 9, 9, 9}

	max := GetMostRepeated(array)

	if max != 9 {
		t.Errorf("Expected most repeated element to be 9, but got %d", max)
	}
}
