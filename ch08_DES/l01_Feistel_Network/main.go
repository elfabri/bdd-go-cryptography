/*
For marketing purposes, Passly has decided
to create its own Feistel network.
It will use the Go standard library's SHA-256 hash function
as the round function.
*/
package main

func feistel(msg []byte, roundKeys [][]byte) []byte {
    l := len(msg)
    right := msg[:l]
    left := msg[l:]
    var nextR, nextL []byte
    for i, k := range roundKeys {
        switch {
        case i == 0:
            nextR = xor(left, hash(right, k, l))
            nextL = right
        case i%2 == 0:
            tmp := nextL
            nextL = nextR
            nextR = xor(tmp, hash(nextL, k, l))
        default:
            tmp := nextR
            nextR = nextL
            nextL = xor(tmp, hash(nextR, k, l))
        }
    }
    nextL = append(nextL, nextR...)
    return nextL
}
