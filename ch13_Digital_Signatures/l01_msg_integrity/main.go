/*
Message integrity ensures that a message has not been changed.

While sending data through the internet,
there's a chance that information is lost in transit.
If there is data loss, you'll want to know about it
so you can try again.

To accomplish this, we can take a hash of the data,
and send it alongside it.

Once on its destination, you will be able to hash
the data, and compare the it with the one sent alongside
the data.

This hash and check process is called checksum
*/
package main

import (
	"crypto/sha256"
	"fmt"
)

func checksumMatches(message string, checksum string) bool {
    hasher := sha256.New()
    _, err := hasher.Write([]byte(message))
    if err != nil {
        fmt.Print("Could not hash the message\n")
        return false
    }
    return fmt.Sprintf("%x", hasher.Sum(nil)) == checksum
}
