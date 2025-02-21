package main

func base8Char(bits byte) string {
	const base8Alphabet = "ABCDEFGH"
    idx := int(bits)

    if idx > 7 {
        return ""
    }

    return string(base8Alphabet[idx])
}
