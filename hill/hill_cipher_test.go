package main

import (
	"testing"
)

func deepEqualSingleDimensionArray(ans []int, expect []int, t *testing.T) {
	if len(ans) != len(expect) {
		t.Errorf("Expected %v, got %v", expect, ans)
		return
	}

	for i := 0; i < len(ans); i++ {
		if ans[i] != expect[i] {
			t.Errorf("Expected %v, got %v", expect, ans)
			return
		}
	}
}

func TestVectorMultiplication(t *testing.T) {
	vec := [][]int{
		{11, 5},
		{3, 4},
	}
	key := []int{
		7, 2,
	}

	ans := vectorMulti(vec, key)
	expect := []int{87, 29}

	deepEqualSingleDimensionArray(ans, expect, t)
}

func testUtil(alphabet string, expect int, t *testing.T) {
	num := alphabetToNumber(alphabet)
	if num != expect {
		t.Errorf("Expected %v, got %v", expect, num)
	}

	numToAlpha := numberToAlphabet(num)
	exp := alphabet
	if numToAlpha != exp {
		t.Errorf("Expected %v, got %v", exp, numToAlpha)
	}
}

func TestAlphabetToNumber(t *testing.T) {
	testUtil("A", 0, t)
	testUtil("B", 1, t)
	testUtil("C", 2, t)
}

func TestVectorizedKey(t *testing.T) {
	key := "GYBNQKURP"
	size := 3
	ans := vectorizeKey(key, size)
	expect := [][]int{
		{6, 24, 1},
		{13, 16, 10},
		{20, 17, 15},
	}

	ansLen := len(ans)
	expectLen := len(expect)

	if ansLen != expectLen {
		t.Errorf("Lenght not matched")
	}

	for i := 0; i < ansLen; i++ {
		for j := 0; j < ansLen; j++ {
			if ans[i][j] != expect[i][j] {
				t.Errorf("Expected %v, got %v", expect, ans)
				return
			}
		}
	}
}

func TestVectorizedPlainText(t *testing.T) {
	plainText := "ABC"
	ans := vectorizePlainText(plainText)
	expect := []int{0, 1, 2}

	deepEqualSingleDimensionArray(ans, expect, t)
}

func TextVectorToPlainText(t *testing.T) {
	vec := []int{0, 1, 2}
	ans := vectorToPlainText(vec)
	expect := "ABC"

	if ans != expect {
		t.Errorf("Expected %v, got %v", expect, ans)
	}
}

func TestEncryption(t *testing.T) {
	size := 2
	key := "DZYR"
	plainText := "MISSISSIPPIK"
	encrypted := Encrypt(plainText, key, size)
	expect := "CIKKGEUWEROY"

	if expect != encrypted {
		t.Errorf("Expected %v, got %v", expect, encrypted)
	}
}
