/*
Aside from the message itself,
all we need to perform encryption are
the numbers e and n. Together,
these are the public key.

The security of RSA relies on the fact that
given n, it's really hard to guess
what p and q (the private keys) were.

Remember, n = p * q

The math for encrypting a message with RSA follows this formula:

    ciphertext = me (mod n)

where:
    m = message
    e = public key exponent
    n = public key modulus

Complete the encrypt function.

 - encrypt(m, e, n *big.Int) *big.Int
Return the result of m^e (mod n). Use the .Exp method.
*/
package main

import (
	"math/big"
)

func encrypt(m, e, n *big.Int) *big.Int {
	return new(big.Int).Exp(m, e, n)
}
