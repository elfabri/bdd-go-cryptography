/*
Add the following functions and methods to the program:

  - newHasher

  - h.Write

  - h.GetHex

- newHasher
Returns a pointer to a new hasher. Uses sha256.New()
to create a new hash.Hash.

- h.Write
A method on a pointer to a hasher. Uses hash.Write()
to write data to the hasher. It should accept
a string and cast the string to a []byte.
It should pass along the return values,
that is, it returns the number of bytes
written from p (0 <= n <= len(p)) and any error
encountered that caused the write to stop early.

- h.GetHex
A method on a pointer to a hasher. Uses hash.Sum()
to get the hash value of the data written to the hasher.
It should encode the hash value as a lowercase hex string and
return it.
*/
package main

import (
	"crypto/sha256"
	"fmt"
	"hash"
)

type hasher struct {
	hash hash.Hash
}

func newHasher() *hasher {
    h := hasher {
        hash: sha256.New(),
    }
    return &h
}

func (h *hasher) Write(msg string) (int, error) {
    /*
    It should accept a string and cast the string to a []byte.
    It should pass along the return values, that is,
    it returns the number of bytes written
    from p (0 <= n <= len(p)) and any error encountered
    that caused the write to stop early.
    */
    b := []byte(msg)
    return h.hash.Write(b)
}

func (h *hasher) GetHex() string {
    /*
    Uses hash.Sum() to get the hash value
    of the data written to the hasher.
    It should encode the hash value
    as a lowercase hex string and return it.
    */
    return fmt.Sprintf("%x", h.hash.Sum(nil))
}
