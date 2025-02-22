package main

import (
	"fmt"
	"testing"
)

func TestAlphabetSize(t *testing.T) {
	type testCase struct {
		numBits  int
		expected float64
	}

	tests := []testCase{
		{1, 2},      // 2^1 = 2
		{2, 4},      // 2^2 = 4
		{3, 8},      // 2^3 = 8
		{4, 16},     // 2^4 = 16
		{5, 32},     // 2^5 = 32
		{8, 256},    // 2^8 = 256
		{16, 65536}, // 2^16 = 65536
	}

	if withSubmit {
		tests = append(tests, []testCase{
			{6, 64},  // 2^6 = 64
			{7, 128}, // 2^7 = 128
		}...)
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		result := alphabetSize(test.numBits)
		if result != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Inputs:      numBits: %v
Expecting:   %v
Actual:      %v
Fail`, test.numBits, test.expected, result)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:      numBits: %v
Expecting:   %v
Actual:      %v
Pass`, test.numBits, test.expected, result)
		}
	}

	fmt.Printf("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

var withSubmit = true
