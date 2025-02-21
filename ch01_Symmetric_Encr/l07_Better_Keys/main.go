package main

import (
	"crypto/aes"
	"crypto/cipher"
)

func keyToCipher(key string) (cipher.Block, error) {
	return aes.NewCipher([]byte(key))
    // NewCipher creates and returns a new cipher.Block.
    // The key argument should be the AES key,
    // either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256.
}

