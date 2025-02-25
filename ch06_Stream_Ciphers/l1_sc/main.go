package main

func crypt(textCh, keyCh <-chan byte, result chan<- byte) {
    for char := range textCh {
        result <- char ^ (<-keyCh)
    }
    close(result)
}
