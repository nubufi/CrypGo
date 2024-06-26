package caesar

import (
	"fmt"

	"github.com/nubufi/CrypGo/utils"
)

// Encrypt applies the Caesar cipher to the given plainText using the provided shift value.
// It returns the encrypted text and an error if the character is not found in the alphabet.
//
// Parameters:
//
// - alphabet: A string containing the characters of the alphabet.
//
// - plainText: The text to be encrypted.
//
// - shift: The number of positions to shift each character in the plainText.
//
// - ignoreUnknown: A boolean value to determine whether to ignore unknown characters in the plainText.
// if set to true, unknown characters will be left as is in the encrypted text else an error will be returned.
//
// Returns:
//
// - string: The encrypted text.
//
// - error: An error if the character is not found in the alphabet.
//
// Example:
//
// - encryptedText, err := caesar.Encrypt("abcdefghijklmnopqrstuvwxyz", "hello", 3,true)
func Encrypt(alphabet, plainText string, shift int, ignoreUnknown bool) (string, error) {
	encryptedText := ""
	alphRune := []rune(alphabet)

	for _, char := range plainText {
		currentIndex := utils.IndexOfRune(alphRune, char)
		if currentIndex == -1 {
			if ignoreUnknown {
				encryptedText += string(char)
				continue
			}
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
//
// Parameters:
//
// - alphabet: A string containing the characters of the alphabet.
//
// - encryptedText: The text to be decrypted.
//
// - shift: The number of positions to shift each character in the encrypted text.
//
// - ignoreUnknown: A boolean value to determine whether to ignore unknown characters in the plainText.
// if set to true, unknown characters will be left as is in the encrypted text else an error will be returned.
//
// Returns:
//
// - string: The decrypted text.
//
// - error: An error if the character is not found in the alphabet.
//
// Example:
//
// - decryptedText, err := caesar.Decrypt("abcdefghijklmnopqrstuvwxyz", "khoor", 3,true)
func Decrypt(alphabet, encryptedText string, shift int, ignoreUnknown bool) (string, error) {
	decryptedText := ""
	alphRune := []rune(alphabet)

	for _, char := range encryptedText {
		if char == ' ' {
			decryptedText += " "
			continue
		}
		currentIndex := utils.IndexOfRune(alphRune, char)
		if currentIndex == -1 {
			if ignoreUnknown {
				decryptedText += string(char)
				continue
			}
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
