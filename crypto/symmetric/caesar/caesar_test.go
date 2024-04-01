package caesar

import (
	"testing"

	"github.com/nubufi/CrypGo/data"
)

func TestEncrypt(t *testing.T) {
	plainText := "HELLO World"
	alphabet := data.EnglishAlphabet
	shift := 3
	result, _ := Encrypt(alphabet, plainText, shift)

	if result != "khoor zruog" {
		t.Errorf("Expected khoor zruog but got %s", result)
	}
}

func TestDecrypt(t *testing.T) {
	encryptedText := "khoor zruog"
	alphabet := data.EnglishAlphabet
	shift := 3
	result, _ := Decrypt(alphabet, encryptedText, shift)

	if result != "hello world" {
		t.Errorf("Expected hello world but got %s", result)
	}

}
