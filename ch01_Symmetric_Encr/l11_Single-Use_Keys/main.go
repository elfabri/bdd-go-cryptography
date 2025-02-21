// Note
// We use rand.Read and rand.Seed from the math/rand package
// in this assignment, even though they are deprecated,
// to ensure predictable results for learning purposes.
// In real-world applications, you should use crypto/rand.Read
// for secure randomness and rand.New(rand.NewSource(seed))
// for a local random generator, as these offer
// better security and reliability.

package main

import (
	"fmt"
	"math/rand"
)

func generateRandomKey(length int) (string, error) {
    // length in bytes,
    // and returns a random key of that length,
    // formatted in a hex string.

    b := make([]byte, length)
    _, err := rand.Read(b)
    if err != nil {
        return "", err
    }
    return fmt.Sprintf("%x", b), nil
}
