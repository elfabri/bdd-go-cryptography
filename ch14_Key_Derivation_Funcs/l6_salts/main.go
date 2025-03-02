/*
They didn't salt their passwords and they're
using SHA-256 instead of a slow KDF. We can
only tackle one problem at a time, so for now,
we've just been asked to salt and re-hash
the passwords. We can migrate from SHA-256 to Bcrypt later.

To do that, we need all of our users to reset
their passwords, because we don't store the plaintext.

Complete the generateSalt and hashPassword functions.

  - generateSalt

Use crypto/rand to generate a random salt
of the specified length. Use rand.Read().

- hashPassword
Append the salt directly to the end
of the password, then hash it
with SHA-256. Use crypto/sha256.
Return the result of the hash.
*/
package main

import (
	"crypto/rand"
	"crypto/sha256"
)

func generateSalt(length int) ([]byte, error) {
    salt := make([]byte, length)
    _, err := rand.Read(salt)
    if err != nil {
        return nil, err
    }
    return salt, nil
}

func hashPassword(password, salt []byte) []byte {
    hasher := sha256.New()
    saltedPass := append(password, salt...)
    hasher.Write(saltedPass)
    return hasher.Sum(nil)
}
