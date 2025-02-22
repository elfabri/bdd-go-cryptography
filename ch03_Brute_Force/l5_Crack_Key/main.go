package main

import (
	"errors"
	"math"
)

func findKey(encrypted []byte, decrypted string) ([]byte, error) {
    // key is a number from 0 to 2 ^ 24 in bytes
    var k []byte
    var data []byte
    for i := 0; i < int(math.Pow(2, 24)); i++ {
        k = intToBytes(i)
        data = crypt(encrypted, k)
        if string(data) == decrypted {
            return k, nil
        }
    }
    return []byte{}, errors.New("key not found")
}
