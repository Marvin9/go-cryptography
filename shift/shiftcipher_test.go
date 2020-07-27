package main

import "testing"

type input struct {
	plainText  string
	cipherText string
	key        int
}

var inputs = []input{
	input{
		plainText:  "mayur",
		cipherText: "nbzvs",
		key:        1,
	},
	input{
		"mayur",
		"uigcz",
		8,
	},
	input{
		"asdkfjlasjldf",
		"cufmhlnculnfh",
		2,
	},
}

func TestEncryption(t *testing.T) {
	for _, input := range inputs {
		cipherText := Encrypt(input.plainText, input.key)
		if cipherText != input.cipherText {
			t.Errorf("Incorrect encryption of %v, expected %v but got %v.", input.plainText, input.cipherText, cipherText)
		}
	}
}

func TestDecryption(t *testing.T) {
	for _, input := range inputs {
		plainText := Decrypt(input.cipherText, input.key)
		if plainText != input.plainText {
			t.Errorf("Incorrect decryption of %v, expected %v but got %v.", input.cipherText, input.plainText, plainText)
		}
	}
}
