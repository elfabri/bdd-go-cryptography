package main

func deriveRoundKey(masterKey [4]byte, roundNumber int) [4]byte {
    bNum := byte(roundNumber)
    roundK := [4]byte{}

    for i, v := range masterKey {
        roundK[i] = v ^ bNum
    }

    return roundK
}
