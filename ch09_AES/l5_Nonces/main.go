/*
Because we use AES-256 at Passly, we're very concerned
about using truly random nonces so we don't accidentally
reuse one. Write the nonceStrength function.
It should return the strength of a nonce as an integer.

We refer to the strength of a nonce as the total number
of possible combinations of bits
that could exist in the nonce.
*/

package main

import "math"

// nonceStrength returns the number of bits of entropy in the nonce.
func nonceStrength(nonce []byte) int {
    // each byte has 8 bits, which can be 1 or 0
	return int(math.Pow(2, float64(len(nonce)) * 8))
}
