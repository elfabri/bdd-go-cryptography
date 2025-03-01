/*
Our web systems at Passly still use RSA,
even though we've moved our native desktop
encryption to ECC. Complete the encrypt function.

Use the rsa.EncryptOAEP function to encrypt
the message with the public key. Use nil as the label
because we don't need it. Follow the patterns
in the documentation as well as in the decrypt function
if you're concerned about the other parameters.
*/
package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

func encrypt(pubKey *rsa.PublicKey, msg []byte) ([]byte, error) {
	plaintext, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, pubKey, msg, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}
