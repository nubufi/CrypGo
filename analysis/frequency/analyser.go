package frequency

func MatchFrequency(text string, frequency string) int {
	matches := 0
	for i, char := range text {
		if i >= len(frequency) {
			break
		}
		if char == rune(frequency[i]) {
			matches++
		}
	}
	return matches
}
