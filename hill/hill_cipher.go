package main

import "fmt"

func vectorMulti(vec [][]int, key []int) []int {
	ptSize := len(vec[0])
	keySize := len(key)

	if ptSize != keySize {
		panic("Cannot perform multiplication.")
	}

	multipliedMatrix := make([]int, keySize)

	for i := 0; i < keySize; i++ {
		ans := 0
		for j := 0; j < keySize; j++ {
			ans += vec[i][j] * key[j]
		}
		multipliedMatrix[i] = ans
	}

	return multipliedMatrix
}

func alphabetToNumber(alphabet string) int {
	return int(alphabet[0]) - 65
}

func numberToAlphabet(num int) string {
	return string(num + 65)
}

func vectorizeKey(key string, size int) [][]int {
	if size*size != len(key) {
		panic("Invalid size to make matrix")
	}
	vectorized := make([][]int, size)
	charIterator := 0
	for i := 0; i < size; i++ {
		vectorized[i] = make([]int, size)
		for j := 0; j < size; j++ {
			vectorized[i][j] = alphabetToNumber(string(key[charIterator]))
			charIterator++
		}
	}
	return vectorized
}

func vectorizePlainText(text string) []int {
	textLen := len(text)
	vectorized := make([]int, textLen)
	for i := 0; i < textLen; i++ {
		vectorized[i] = alphabetToNumber(string(text[i]))
	}
	return vectorized
}

func vectorToPlainText(vec []int) string {
	pt := ""
	for i := 0; i < len(vec); i++ {
		pt += numberToAlphabet(vec[i] % 26)
	}
	return pt
}

func Encrypt(pt, key string, size int) string {
	vectorizedKey := vectorizeKey(key, size)

	cipherText := ""
	for i := 0; i < len(pt); i += size {
		plainText := pt[:i+size]
		vectorizedPt := vectorizePlainText(plainText)

		multipliedMatrix := vectorMulti(vectorizedKey, vectorizedPt)

		cipherChunk := vectorToPlainText(multipliedMatrix)

		cipherText += cipherChunk
	}

	return cipherText
}

func main() {
	var plainText string
	var key string
	size := 3
	fmt.Printf("Enter plaintext: ")
	fmt.Scan(&plainText)
	fmt.Printf("Enter key: ")
	fmt.Scan(&key)

	fmt.Println(Encrypt(plainText, key, size))
}
