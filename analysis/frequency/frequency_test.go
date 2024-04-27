package frequency

import (
	"testing"
)

func TestCountLetters(t *testing.T) {
	text := "hello"

	counts := CountLetters(text)

	if counts['l'] != 2 {
		t.Errorf("Expected count of 'l' to be 2, but got %d", counts['h'])
	}
}

func TestSortByFrequency(t *testing.T) {
	text := "hello"

	counts := CountLetters(text)

	sorted := SortByFrequency(counts)

	if sorted[0] != 'l' {
		t.Errorf("Expected first character to be 'l', but got %s", string(sorted[0]))
	}
}
