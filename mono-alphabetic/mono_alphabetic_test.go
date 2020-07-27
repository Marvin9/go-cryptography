package main

import "testing"

func TestPermutation(t *testing.T) {
	var keys = make(map[string]bool)
	checkLen := generatePermutation()
	if len(checkLen) != 26 {
		t.Errorf("Expected length of permutation %v to be %v, but got %v", checkLen, 26, len(checkLen))
	}

	for i := 0; i < 20; i++ {
		perm := stringify(generatePermutation())
		_, ok := keys[perm]
		if ok {
			t.Errorf("Permutation %v repeated in small range.", perm)
		}
		keys[perm] = true
	}
}
