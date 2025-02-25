package main

func crypt(plaintext, key []byte) []byte {
    res := make([]byte, len(plaintext))
    for i, v := range plaintext {
        res[i] = v ^ key[i]
    }
    return res
}
