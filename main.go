package main

import (
	"fmt"

	"github.com/nubufi/CrypGo/crypto/symmetric/caesar"
	"github.com/nubufi/CrypGo/data"
)

func main() {
	plainText := "احبك نعمان"
	alphabet := data.ArabicAlphabet

	shift := 3

	encryptedText, _ := caesar.Encrypt(alphabet, plainText, shift)

	fmt.Print("Encrypted text: ")
	fmt.Println(encryptedText)

	decryptedText, _ := caesar.Decrypt(alphabet, encryptedText, shift)

	fmt.Print("Decrypted text: ")
	fmt.Println(decryptedText)
}
