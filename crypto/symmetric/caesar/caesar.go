package caesar

import (
	"fmt"
	"strings"

	"github.com/nubufi/CrypGo/utils"
)

// Encrypt applies the Caesar cipher to the given plainText using the provided shift value.
// It returns the encrypted text and an error if the character is not found in the alphabet.
func Encrypt(alphabet, plainText string, shift int) (string, error) {
	encryptedText := ""
	plainText = strings.ToLower(plainText)
	alphRune := []rune(alphabet)

	for _, char := range plainText {
		if char == ' ' {
			encryptedText += " "
			continue
		}
		currentIndex := utils.IndexOfRune(alphRune, char)
		if currentIndex == -1 {
			return "", fmt.Errorf("character %s not found in alphabet", string(char))
		}
		encryptedIndex := (currentIndex + shift) % len(alphRune)

		encryptedText += string(alphRune[encryptedIndex])

	}

	return encryptedText, nil
}

// Decrypt takes an alphabet, encrypted text, and a shift value as input and returns the decrypted text.
// It uses the Caesar cipher algorithm to decrypt the text by shifting each character in the encrypted text
// backwards by the specified shift value. The decrypted text is returned along with an error, if any.
func Decrypt(alphabet, encryptedText string, shift int) (string, error) {
	decryptedText := ""
	encryptedText = strings.ToLower(encryptedText)
	alphRune := []rune(alphabet)

	for _, char := range encryptedText {
		if char == ' ' {
			decryptedText += " "
			continue
		}
		currentIndex := utils.IndexOfRune(alphRune, char)
		if currentIndex == -1 {
			return "", fmt.Errorf("character %s not found in alphabet", string(char))
		}
		decryptedIndex := (currentIndex - shift) % len(alphRune)

		if decryptedIndex < 0 {
			decryptedIndex += len(alphRune)
		}
		decryptedText += string(alphRune[decryptedIndex])
	}

	return decryptedText, nil
}
