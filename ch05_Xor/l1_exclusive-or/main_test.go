package main

import (
	"fmt"
	"testing"
)

func TestXor(t *testing.T) {
	type testCase struct {
		lhs      bool
		rhs      bool
		expected bool
	}

	tests := []testCase{
		{true, true, false}, // true XOR true = false
		{true, false, true}, // true XOR false = true
	}

	if withSubmit {
		tests = append(tests, []testCase{
			{false, true, true},   // false XOR true = true
			{false, false, false}, // false XOR false = false
		}...)
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		result := xor(test.lhs, test.rhs)
		if result != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Inputs:      lhs: %v, rhs: %v
Expecting:   %v
Actual:      %v
Fail`, test.lhs, test.rhs, test.expected, result)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:      lhs: %v, rhs: %v
Expecting:   %v
Actual:      %v
Pass`, test.lhs, test.rhs, test.expected, result)
		}
	}

	fmt.Printf("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

var withSubmit = true
