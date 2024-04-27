package frequency

import "sort"

// CountLetters returns a map with the count of each character in the text.
func CountLetters(text string) map[rune]int {
	counts := make(map[rune]int)
	for _, char := range text {
		if char == ' ' {
			continue
		}
		counts[char]++
	}

	return counts
}

// SortByFrequency returns a slice of characters sorted by frequency in descending order.
func SortByFrequency(counts map[rune]int) []rune {
	var sorted []rune
	for char := range counts {
		sorted = append(sorted, char)
	}

	sort.Slice(sorted, func(i, j int) bool {
		return counts[sorted[i]] > counts[sorted[j]]
	})

	return sorted
}
