package main

import "math/bits"

func hash(input []byte) [4]byte {
    var final [4]byte
    gg := make([]byte, len(input))

    for i, b := range input {
        // Rotate bits left 3 bits
        // b = (b << 3) | (b >> (8 - 3))

        // Shift bits left 2 bits
        // b <<= 2

        // Rotate left 3 bits using RotateLeft8 and
        // shift left 2 bits
        gg[i] = bits.RotateLeft8(uint8(b), 3) << 2

        // XOR with the corresponding byte in the final array
        final[i%4] ^= gg[i]
    }

    return final
}
