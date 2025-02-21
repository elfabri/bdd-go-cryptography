package main

import (
	"fmt"
	"testing"
)

func TestBase8Char(t *testing.T) {
	type testCase struct {
		bits       byte
		expected   string
		shouldFail bool
	}

	tests := []testCase{
		{0b000, "A", false}, // 0000 -> A
		{0b001, "B", false}, // 0001 -> B
		{0b010, "C", false}, // 0010 -> C
		{0b011, "D", false}, // 0011 -> D
		{0b100, "E", false}, // 0100 -> E
		{0b101, "F", false}, // 0101 -> F
		{0b110, "G", false}, // 0110 -> G
		{0b111, "H", false}, // 0111 -> H
	}

	if withSubmit {
		tests = append(tests, []testCase{
			{0b101, "F", false},    // Valid additional case
			{0b1111, "", true},     // 4-bit, out of range
			{0b1000, "", true},     // 1000 is out of bounds
			{0b11111111, "", true}, // Max byte value, out of bounds
		}...)
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		result := base8Char(test.bits)
		if (result == "") != test.shouldFail {
			failCount++
			t.Errorf(`---------------------------------
Inputs:      bits: %04b
Expecting:   %v
Actual:      %v
Fail`, test.bits, test.expected, result)
		} else if result != test.expected && !test.shouldFail {
			failCount++
			t.Errorf(`---------------------------------
Inputs:      bits: %04b
Expecting:   %v
Actual:      %v
Fail`, test.bits, test.expected, result)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:      bits: %04b
Expecting:   %v
Actual:      %v
Pass`, test.bits, test.expected, result)
		}
	}

	fmt.Printf("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

var withSubmit = true
