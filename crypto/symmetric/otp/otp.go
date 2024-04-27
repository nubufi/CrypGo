package otp

import (
	"fmt"

	"github.com/nubufi/CrypGo/utils"
)

// Encrypt applies the One Time Path algorithm to the given plainText using the provided key.
// It returns the encrypted text and an error if any character in the plainText or the key is not found in the alphabet.
//
// Parameters:
//
// - alphabet: A string containing the characters of the alphabet.
//
// - plainText: The text to be encrypted.
//
// - key: The random sequence of number of positions to shift each character in the plainText.
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
// - randomSequence := otp.RandomSequence("hello", 26)
// - encryptedText, err := otp.Encrypt("abcdefghijklmnopqrstuvwxyz", "hello", randomSequence, true)
func Encrypt(alphabet, plainText string, key []int, ignoreUnknown bool) (string, error) {
	if len(plainText) != len(key) {
		return "", fmt.Errorf("text and key must have the same length")
	}
	alphRune := []rune(alphabet)

	encryptedText := ""
	for i, char := range plainText {
		currentIndex := utils.IndexOfRune(alphRune, char)
		if currentIndex == -1 {
			if ignoreUnknown {
				encryptedText += string(char)
				continue
			}
			return "", fmt.Errorf("character %s not found in alphabet", string(char))
		}

		encryptedIndex := (currentIndex + key[i]) % len(alphRune)
		encryptedText += string(alphRune[encryptedIndex])
	}
	return encryptedText, nil
}

// Decrypt takes an alphabet, encrypted text, and a key as input and returns the decrypted text.
// It uses the One Time Path algorithm to decrypt the text by shifting each character in the encrypted text
// backwards by the specified shift value. The decrypted text is returned along with an error, if any.
//
// Parameters:
//
// - alphabet: A string containing the characters of the alphabet.
//
// - encryptedText: The text to be decrypted.
//
// - key: The random sequence of number of positions to shift each character in the plainText.
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
// - randomSequence := otp.RandomSequence("hello", 26)
// - decryptedText, err := caesar.Decrypt("abcdefghijklmnopqrstuvwxyz", randomSequence, 3,true)
func Decrypt(alphabet, encryptedText string, key []int, ignoreUnknown bool) (string, error) {
	if len(encryptedText) != len(key) {
		return "", fmt.Errorf("text and key must have the same length")
	}
	decryptedText := ""
	alphRune := []rune(alphabet)

	for i, char := range encryptedText {
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
		decryptedIndex := (currentIndex - key[i]) % len(alphRune)

		if decryptedIndex < 0 {
			decryptedIndex += len(alphRune)
		}
		decryptedText += string(alphRune[decryptedIndex])
	}

	return decryptedText, nil
}
