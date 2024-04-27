package viginere

import (
	"fmt"
	"testing"

	"github.com/nubufi/CrypGo/data"
)

func TestEncrypt(t *testing.T) {
	alphabet := data.EnglishAlphabet
	key := "secret"
	tests := []struct {
		plainText     string
		expected      string
		expectedError error
	}{
		{
			plainText:     "this is a message",
			expected:      "llkj ml s qgjwtyi",
			expectedError: nil,
		},
		{
			plainText:     "5invalid",
			expected:      "",
			expectedError: fmt.Errorf("character 5 not found in alphabet"),
		},
	}

	for _, test := range tests {
		encryptedText, err := Encrypt(alphabet, test.plainText, key, false)
		if encryptedText != test.expected {
			t.Errorf("Encrypt(%s, %s, %s) = %s, expected %s", alphabet, test.plainText, key, encryptedText, test.expected)
		}
		if test.expectedError != nil && err == nil {
			t.Errorf("Encrypt(%s, %s, %s) returned error %v, expected %v", alphabet, test.plainText, key, err, test.expectedError)
		}
	}
}

func TestDecrypt(t *testing.T) {
	alphabet := data.EnglishAlphabet
	key := "secret"
	tests := []struct {
		encryptedText string
		expected      string
		expectedError error
	}{
		{
			encryptedText: "llkj ml s qgjwtyi",
			expected:      "this is a message",
			expectedError: nil,
		},
		{
			encryptedText: "5invalid",
			expected:      "",
			expectedError: fmt.Errorf("character 5 not found in alphabet"),
		},
	}

	for _, test := range tests {
		decryptedText, err := Decrypt(alphabet, test.encryptedText, key, false)
		if decryptedText != test.expected {
			t.Errorf("Decrypt(%s, %s, %s) = %s, expected %s", alphabet, test.encryptedText, key, decryptedText, test.expected)
		}
		if test.expectedError != nil && err == nil {
			t.Errorf("Decrypt(%s, %s, %s) returned error %v, expected %v", alphabet, test.encryptedText, key, err, test.expectedError)
		}
	}
}
