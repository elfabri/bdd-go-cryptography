package main

func padWithZeros(block []byte, desiredSize int) []byte {
    block_s := len(block)
    offset := desiredSize - block_s

    zeroed_block := make([]byte, offset)
    block = append(block, zeroed_block...)
    return block
}
