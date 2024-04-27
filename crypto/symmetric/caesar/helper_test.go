package caesar

import (
	"fmt"
	"testing"

	"github.com/nubufi/CrypGo/data"
)

func TestGetShiftNumber(t *testing.T) {
	text := "akdaisdsklasis iassdkiasdkai"

	engFreq := []rune(data.EnglishAlphabetFrequency)
	engAlph := []rune(data.EnglishAlphabet)

	shift := GetShiftNumber(text, engFreq, engAlph)

	fmt.Println(shift)
}
