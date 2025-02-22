package main

import "math"

// It accepts a number of bits per character
// as input, and returns the size of the alphabet
// that can be represented by that number of bits
func alphabetSize(numBits int) float64 {
    return math.Pow(2, float64(numBits))
}
