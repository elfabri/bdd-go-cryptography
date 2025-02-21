package main

import (
	"fmt"
	"strings"
)

func getHexString(b []byte) string {
    hex := make([]string, len(b))

    for i, v := range b {
        hex[i] = fmt.Sprintf("%02x", v)
    }

    return strings.Join(hex, ":")
}

func getBinaryString(b []byte) string {
    bin := make([]string, len(b))

    for i, v := range b {
        bin[i] = fmt.Sprintf("%08b", v)
    }

    return strings.Join(bin, ":")
}
