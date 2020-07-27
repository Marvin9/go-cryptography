package main

import "fmt"

// Encrypt - to encrypt plain text
func Encrypt(text string, key int) string {
	cipherText := ""
	for _, char := range text {
		encryptedChar := string(int(char) + key)
		if int(encryptedChar[0]) > int('z') {
			offsetFromA := key - (int('z') - int(char))
			encryptedChar = string(int('a') + offsetFromA - 1)
		}
		cipherText += encryptedChar
	}
	return cipherText
}

// Decrypt - to decrypt cipher text
func Decrypt(text string, key int) string {
	plainText := ""
	for _, char := range text {
		decryptedChar := string(int(char) - key)
		if int(decryptedChar[0]) < int('a') {
			offsetFromZ := key - (int(char) - int('a'))
			decryptedChar = string(int('z') - offsetFromZ + 1)
		}
		plainText += decryptedChar
	}
	return plainText
}

func main() {
	var plainText string
	var key int

	fmt.Printf("Enter plain text: ")
	fmt.Scan(&plainText)
	fmt.Printf("Enter key: ")
	fmt.Scan(&key)

	cipherText := Encrypt(plainText, key)
	fmt.Println("Cipher text:", cipherText)
	fmt.Println("Decrypted text: ", Decrypt(cipherText, key))
}
