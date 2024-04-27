package viginere

import (
	"fmt"

	"github.com/nubufi/CrypGo/utils"
)

// Encrypt applies the Viginere cipher to the given plainText using the provided key.
// It returns the encrypted text and an error if any character in the plainText or the key is not found in the alphabet.
//
// Parameters:
//
// - alphabet: A string containing the characters of the alphabet.
//
// - plainText: The text to be encrypted.
//
// - key: The key to be used for encryption.
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
// - encryptedText, err := viginere.Encrypt("abcdefghijklmnopqrstuvwxyz", "hello", "secret", true)
func Encrypt(alphabet, plainText string, key string, ignoreUnknown bool) (string, error) {
	encryptedText := ""
	alphRune := []rune(alphabet)

	n := 0
	for _, char := range plainText {
		currentIndex := utils.IndexOfRune(alphRune, char)
		if currentIndex == -1 {
			if ignoreUnknown {
				encryptedText += string(char)
				continue
			}
			return "", fmt.Errorf("character %s not found in alphabet", string(char))
		}
		shift := utils.IndexOfRune(alphRune, rune(key[n%len(key)]))

		if shift == -1 {
			return "", fmt.Errorf("character %s not found in alphabet", string(char))
		}
		encryptedIndex := (currentIndex + shift) % len(alphRune)

		encryptedText += string(alphRune[encryptedIndex])
		n += 1
	}

	return encryptedText, nil
}

// Decrypt takes an alphabet, encrypted text, and a key as input and returns the decrypted text.
// It uses the Viginere cipher algorithm to decrypt the text by shifting each character in the encrypted text
// backwards. The decrypted text is returned along with an error, if any.
//
// Parameters:
//
// - alphabet: A string containing the characters of the alphabet.
//
// - encryptedText: The text to be decrypted.
//
// - key: The key to be used for decryption.
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
// - decryptedText, err := viginere.Decrypt("abcdefghijklmnopqrstuvwxyz", "khoor", "secret",true)
func Decrypt(alphabet, encryptedText string, key string, ignoreUnknown bool) (string, error) {
	decryptedText := ""
	alphRune := []rune(alphabet)

	n := 0
	for _, char := range encryptedText {
		currentIndex := utils.IndexOfRune(alphRune, char)
		if currentIndex == -1 {
			if ignoreUnknown {
				decryptedText += string(char)
				continue
			}
			return "", fmt.Errorf("character %s not found in alphabet", string(char))
		}
		shift := utils.IndexOfRune(alphRune, rune(key[n%len(key)]))
		if shift == -1 {
			return "", fmt.Errorf("character %s not found in alphabet", string(char))
		}

		decryptedIndex := (currentIndex - shift) % len(alphRune)

		if decryptedIndex < 0 {
			decryptedIndex += len(alphRune)
		}
		decryptedText += string(alphRune[decryptedIndex])
		n += 1

	}

	return decryptedText, nil
}
