/*
Toy HMAC
At Passly, we use HMACs to authenticate messages
between our internal servers when they need to make
requests to each other over the public internet.

To demonstrate to our Luddite manager why we should
use an open-source crypto library instead of writing
our own HMAC implementation, we decided to write our
own and then prove its inferiority. Hopefully
we don't get fired instead of getting our way.

Complete the hmac function. It should:

1 - Split the key into two halves. The second half should be
	the larger half if key's length is odd

2 - Return the result of
	sha256(keyFirstHalf + sha256(keySecondHalf + message))
	as a string in lowercase hex

Cast strings directly to slices of bytes and
don't use any delimiters when concatenating the data
*/
package main

import (
	"crypto/sha256"
	"fmt"
	"math"
)

func hmac(message, key string) string {
    l := len(key)
    half := int(math.Floor(float64(l/2)))

    k1 := key[:half]
    k2 := key[half:]

    h1 := sha256.New()
    h1.Write([]byte(k2 + message))

    h2 := sha256.New()
    h2.Write([]byte(k1 + string(h1.Sum(nil))))

    return fmt.Sprintf("%x", h2.Sum(nil))
}
