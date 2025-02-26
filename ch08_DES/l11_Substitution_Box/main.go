/*
Complete the sBox() function.
It maps the last 4 bits of a byte,
down to the last 2 bits of a byte.
*/
package main

import "errors"

var table = [][]byte{
    {0b00, 0b10, 0b01, 0b11},
    {0b10, 0b00, 0b11, 0b01},
    {0b01, 0b11, 0b00, 0b10},
    {0b11, 0b01, 0b10, 0b00},
}

// Each column represents the first two bits of input,
// the rows represent the second two bits.

func sBox(b byte) (byte, error) {
    if b > 0b1111 {
        return 0, errors.New("invalid input")
    }

    first := (b >> 2) & 0b11
    last := b & 0b11

    return table[first][last], nil
}
