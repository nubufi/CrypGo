package utils

func IndexOfRune(r []rune, char rune) int {
	for i, v := range r {
		if v == char {
			return i
		}
	}
	return -1
}
