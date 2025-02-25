/*
Complete the encrypt, decrypt, crypt, and getOffsetChar functions.
The encrypt and decrypt functions do what you would expect,
crypt and getOffsetChar are helper functions.

If you can't find the character in the alphabet,
return an empty string.
*/

package main

func encrypt(plaintext string, key int) string {
    return crypt(plaintext, key)
}

func decrypt(ciphertext string, key int) string {
    return crypt(ciphertext, -key)
}

func crypt(text string, key int) string {
    ciphered := ""

    for _, v := range text {
        ciphered += getOffsetChar(v, key)
    }

    return ciphered
}

func getOffsetChar(c rune, offset int) string {
	const alphabet = "abcdefghijklmnopqrstuvwxyz"
    charIdx := offset
    found := false
    for i, v := range alphabet {
        if c == v {
            found = true
            charIdx += i
            break
        }
    }

    if !found {
        return ""
    }

    if charIdx >= len(alphabet) {
        charIdx -= len(alphabet)
    }

    if charIdx < 0 {
        charIdx += len(alphabet)
    }

    return string(alphabet[charIdx])
}
