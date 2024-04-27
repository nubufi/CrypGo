package caesar

import (
	"fmt"

	"github.com/nubufi/CrypGo/analysis/frequency"
	"github.com/nubufi/CrypGo/utils"
)

func GetShiftNumber(text string, alphabetFrequency, alphabet []rune) int {
	counts := frequency.CountLetters(text)
	sorted := frequency.SortByFrequency(counts)

	match := matchByFrequency(sorted, alphabetFrequency)

	shiftList := getShiftList(match, alphabet)

	fmt.Println(shiftList)
	shift := utils.GetMostRepeated(shiftList)

	return shift
}

func matchByFrequency(textFreq, alphabetFreq []rune) map[rune]rune {
	match := make(map[rune]rune)
	for i, char := range textFreq {
		match[char] = alphabetFreq[i]
	}
	return match
}

func getShiftList(match map[rune]rune, alphabet []rune) []int {
	shiftList := make([]int, 0)

	// i is the index of the rune in the text, j is the index of the rune in the alphabet
	for i, j := range match {
		indexOfI := utils.IndexOfRune(alphabet, i)
		indexOfJ := utils.IndexOfRune(alphabet, j)

		shift := indexOfJ - indexOfI
		if shift < 0 {
			shift += len(alphabet)
		}
		shiftList = append(shiftList, shift)
	}
	return shiftList
}
