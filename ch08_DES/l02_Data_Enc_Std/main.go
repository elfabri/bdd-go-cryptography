/*
For Cryptanalysis purposes, Passly keeps
a DES implementation around so that our engineers
can practice breaking it. Complete the encrypt function.
It should use the standard library's crypto/des package
to encrypt the plaintext with the given key.

Complete the encrypt function and its helper padMsg function.
The decrypt function is already written for you.

example using aes
https://pkg.go.dev/crypto/cipher#NewCBCEncrypter
*/
package main

import (
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"io"
)

func encrypt(key, plaintext []byte) ([]byte, error) {
    if len(plaintext) % des.BlockSize != 0 {
        plaintext = padMsg(plaintext, des.BlockSize)
    }

    block, err := des.NewCipher(key)
    if err != nil {
        return []byte{}, err
    }

    cipherText := make([]byte, des.BlockSize + len(plaintext))
    iv := cipherText[:des.BlockSize]

    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        return nil, err
    }

    mode := cipher.NewCBCEncrypter(block, iv)
    mode.CryptBlocks(cipherText[des.BlockSize:], plaintext)

    return cipherText, nil
}

func padMsg(plaintext []byte, blockSize int) []byte {
    padLen := blockSize - (len(plaintext) % blockSize)
    pad := make([]byte, padLen)
    plaintext = append(plaintext, pad...)
    return plaintext
}
