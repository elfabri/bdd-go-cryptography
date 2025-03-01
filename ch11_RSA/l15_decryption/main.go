/*
Complete the decrypt function.

decrypt(c, d, n *big.Int) *big.Int
Use the big.Int.Exp function
to raise the encrypted message (c)
to the power of the private key (d)
within (mod n).
*/

package main

import (
	"math/big"
)

func decrypt(c, d, n *big.Int) *big.Int {
	return new(big.Int).Exp( c, d, n )
}
