package main

import "testing"

func testError(expected string, got string, t *testing.T) {
	t.Errorf("Expected %v, got %v", expected, got)
}

func TestRemoveDuplicateFunc(t *testing.T) {
	init := "abcd"
	removeDuplicatesFromString(&init)
	expect := "abcd"
	if init != expect {
		testError(expect, init, t)
	}

	init = "aaaa"
	removeDuplicatesFromString(&init)
	expect = "a"
	if init != expect {
		testError(expect, init, t)
	}

	init = "hidden"
	removeDuplicatesFromString(&init)
	expect = "hiden"
	if init != expect {
		testError(expect, init, t)
	}
}

func TestIncrementCharacter(t *testing.T) {
	char := "a"
	incrementChar(&char)
	expect := "b"
	if char != expect {
		testError(expect, char, t)
	}

	char = "i"
	incrementChar(&char)
	expect = "k"
	if char != expect {
		t.Errorf("Must ignore j, expected %v, got %v", expect, char)
	}

	char = "z"
	incrementChar(&char)
	expect = "a"
	if char != expect {
		t.Errorf("Incrementing z should lead to a, but got %v", char)
	}
}

func TestNormalizePlainText(t *testing.T) {
	text := "abc"
	normalizePlainText(&text)
	expect := "abcx"
	if text != expect {
		testError(expect, text, t)
	}

	text = "abcd"
	normalizePlainText(&text)
	expect = "abcd"
	if text != expect {
		testError(expect, text, t)
	}

	text = "aabbc"
	normalizePlainText(&text)
	expect = "axbxcx"
	if text != expect {
		testError(expect, text, t)
	}

	text = "hammer"
	normalizePlainText(&text)
	expect = "hamxer"
	if text != expect {
		testError(expect, text, t)
	}
}

func TestPlaifair(t *testing.T) {
	pt := "we are discovered save yourself"
	key := "monarchy"
	cipher := Encrypt(pt, key)
	expect := "ugrmkcsxhmufmkbtoxgcmvatluiv"
	if cipher != expect {
		testError(expect, cipher, t)
	}
	pt = Decrypt(cipher, key)
	expect = "wearediscoveredsaveyourselfx"
	if pt != expect {
		testError(expect, pt, t)
	}
}
