package main

import (
	"fmt"
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	type testCase struct {
		plaintext string
		key       string
	}

	tests := []testCase{
		{"Shazam", "Sk7p13"},
		{"I'm lovin it", "mysecurepass"},
	}

	if withSubmit {
		tests = append(tests, []testCase{
			{"Don't tell him I'm in love", "c5f149783abf22a96e9a7bb999"},
		}...)
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		ciphertext := encrypt([]byte(test.plaintext), []byte(test.key))
		decrypted := decrypt(ciphertext, []byte(test.key))
		if string(decrypted) != test.plaintext {
			failCount++
			t.Errorf(`---------------------------------
Inputs:      plaintext: %s, key: %s
Expecting:   decrypted: %s
Actual:      decrypted: %s
Fail`, test.plaintext, test.key, test.plaintext, string(decrypted))
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:      plaintext: %s, key: %s
Expecting:   decrypted: %s
Actual:      decrypted: %s
Pass`, test.plaintext, test.key, test.plaintext, string(decrypted))
		}
	}

	fmt.Printf("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

var withSubmit = true
