/*
Passly wants to be an issuer of JWTs.
Our massive egos demand that we push clients
into using "Sign in with Passly".

Complete the createECDSAMessage function.

- createECDSAMessage()
Hash the message using SHA-256
Create a signature of the hashed message
using the private key and the SignASN1 function.

Return a new token in the following format:

	MESSAGE.signature

Where MESSAGE is the original message,
and signature is the signature of the hashed message in lowercase hex.

Keep in mind, this isn't a full JWT,
it's an arbitrary message and a signature.
*/
package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

func createECDSAMessage(message string, privateKey *ecdsa.PrivateKey) (string, error) {
    h := sha256.New()
    _, err := h.Write([]byte(message))
    if err != nil {
        return "", err
    }
    signature, err := ecdsa.SignASN1(rand.Reader, privateKey, h.Sum(nil))
    if err != nil {
        return "", err
    }

    return fmt.Sprintf("%s.%x", message, signature), nil
}
