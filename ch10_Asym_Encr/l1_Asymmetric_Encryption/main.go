/*
At Passly, we use Elliptic Curve Cryptography (ECC)
for our public-key cryptography needs in production.

Complete the genKeys() function.
Use the ecdsa.GenerateKey function from the standard library,
along with the elliptic.P256() curve.

We won't be checking the values of the keys,
so feel free to use a secure random source
from the crypto/rand package.
*/
package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
)

func genKeys() (pubKey *ecdsa.PublicKey, privKey *ecdsa.PrivateKey, err error) {
	privKey, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
    if err != nil {
        return nil,nil, err
    }

    pubKey = &privKey.PublicKey
    return pubKey, privKey, nil
}
