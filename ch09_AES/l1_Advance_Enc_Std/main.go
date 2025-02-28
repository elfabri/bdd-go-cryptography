/*
At Passly, the production cipher we use
to encrypt password values is AES in GCM mode!
Complete the decrypt function.

  - Create a new cipher block using the key.

  - Use the cipher block to create a new GCM

  - Use the GCM (which implements the AEAD interface)
    to decrypt the ciphertext using aesgcm.Open

  - Return the plaintext as a []byte

Return any errors that occur without modifying them.
*/
package main

import (
	"crypto/aes"
	"crypto/cipher"
)

func decrypt(key, ciphertext, nonce []byte) (plaintext []byte, err error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return
    }

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
        return
	}

    plaintext, err = aesgcm.Open(plaintext, nonce, ciphertext, nil)
    return
}
