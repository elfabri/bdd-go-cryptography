/*
An iv, or initialization vector, is a random value
that is used to initialize a block cipher.
It is used to ensure that the same plaintext
always encrypts to a different ciphertext.

Update the generateIV function.
It's currently returning unique values
due to a global count variable, but it's not random,
which is what Passly's cipher requires.

Use the math/rand package's rand.Read function
to generate a random iv of the specified byte length.
*/
package main

import "math/rand"

// var count = 0

func generateIV(length int) ([]byte, error) {
	// count++
    iv := make([]byte, length)
    _, err := rand.Read(iv)
    if err != nil {
        return nil, err
    }

	return iv, nil
}
