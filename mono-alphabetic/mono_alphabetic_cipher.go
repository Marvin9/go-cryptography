package main

import (
	"fmt"
	"math/rand"
	"time"
)

const min int = 97  // a
const max int = 123 // z + 1

func randomAlphabate() string {
	rand.Seed(time.Now().UnixNano())
	alph := string(rand.Intn(max-min) + min)
	return alph
}

func generatePermutation() map[string]string {
	var alphabetIncluded = make(map[string]bool)
	var permutation = make(map[string]string)
	tot := 26
	for i := 0; i < tot; i++ {
		for {
			alphabet := randomAlphabate()
			_, ok := alphabetIncluded[alphabet]
			if !ok {
				permutation[string(min+i)] = alphabet
				alphabetIncluded[alphabet] = true
				break
			}
		}
	}
	return permutation
}

func stringify(key map[string]string) string {
	str := ""
	for _, val := range key {
		str += val
	}
	return str
}

func main() {
	var plainText string
	fmt.Printf("Enter plain text: ")
	fmt.Scan(&plainText)
	key := generatePermutation()
	fmt.Printf("Generated key: %v\n", stringify(key))

	cipherText := ""
	for _, char := range plainText {
		cipherText += key[string(char)]
	}
	fmt.Println("Cipher text:", cipherText)
}
