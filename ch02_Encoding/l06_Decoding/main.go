package main

import (
	"encoding/hex"
	"strings"
)

func getHexBytes(s string) ([]byte, error) {
    // s = 48:65:6c:6c:6f
    // [001010, 00011, ...]

    encd := strings.Split(s, ":")
    dncd, err := hex.DecodeString(strings.Join(encd, ""))
    if err != nil {
        return []byte{}, err
    }

    return dncd, nil
}
