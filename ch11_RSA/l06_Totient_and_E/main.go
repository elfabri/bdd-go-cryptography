/*
Complete the getTot and getE functions.

  - getTot(p, q *big.Int) *big.Int

Use the math/big package to calculate (p-1)(q-1)
and return it as a pointer to a big.Int.
This is the "totient" of n, which we
can also call "phi of n", or ϕ(n).

  - getE(tot *big.Int) *big.Int

Use the math/big package to generate a random number
"e" that adheres to the following constraints:

  - e is greater than 1
  - e is less than tot
  - e and tot have a greatest common divisor of 1

The gcd function is provided for you.
It calculates the greatest common divisor of two big ints.

Generate random e values in the range of [2, tot)
until you find one that satisfies the constraints.
Use crand.Int with the globally provided randReader
to generate random big ints. You'll need to do
some manual arithmetic to get the range you want
because crand.Int only generates random numbers
in the range of [0, max)
*/
package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func getTot(p, q *big.Int) *big.Int {
    return p.Mul(
        p.Sub(p, big.NewInt(1)),
        q.Sub(q, big.NewInt(1)),
    )
}

func getE(tot *big.Int) *big.Int {
    for {
        e, err := rand.Int(randReader, tot)
        if err != nil {
            fmt.Printf("Error getting E: %v\n", err)
        }
        if e.Cmp(big.NewInt(2)) == -1 {
            // less than 2
            continue
        }
        if e.Cmp(tot) == 1 {
            // bigger than tot
            continue
        }
        if gcd(e, tot).Cmp(big.NewInt(1)) != 0 {
            // greatest common divisor has to be 1
            continue
        }
        return e
    }
}
