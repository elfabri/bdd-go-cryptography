/*
The Go standard library has built-in support
for the AES and DES block ciphers

We've been asked by leadership to check
on the block sizes of each algorithm and report back.
Complete the getBlockSize function.

This function accepts a keyLen and cipherType
and returns the block size of the cipher
along with any errors encountered.

The cipherType is an enum of typeAES or typeDES.
Depending on the cipher type,
create a new cipher using:

  - aes.NewCipher or
  - des.NewCipher

The value of the key passed in doesn't matter here,
all that matters is its length.

Return 0 and the error if an error occurs when creating a new cipher.
If no error occurs, return the .BlockSize() of the new cipher and nil.
Return an invalid cipher type error if
the cipherType isn't one of the valid values.
*/
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"errors"
)

func getBlockSize(keyLen, cipherType int) (int, error) {

    var block cipher.Block
    var err error
    switch (cipherType) {
        case typeAES: {
            key := make([]byte, keyLen)
            block, err = aes.NewCipher(key)
        }
        case typeDES: {
            key := make([]byte, keyLen)
            block, err = des.NewCipher(key)
        }
        default: {
            return 0, errors.New("Invalid cipher type")
        }
    }

    if err != nil {
        return 0, err
    }

    return block.BlockSize(), nil
}
