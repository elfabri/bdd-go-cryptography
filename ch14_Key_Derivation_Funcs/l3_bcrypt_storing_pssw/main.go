/*
Each user has a master password that they use
to log into their cloud account. That password
is hashed with Bcrypt before being stored.

Use the golang.org/x/crypto/bcrypt package
to complete the hashPassword() and checkPasswordHash()
functions. You do not need to modify
the function signatures, just implement
the Bcrypt library's API and do
the []byte <-> string conversions.

Use a cost factor of 13
*/
package main

import "golang.org/x/crypto/bcrypt"

func hashPassword(password string) (string, error) {
    h, err := bcrypt.GenerateFromPassword([]byte(password), 13)
    if err != nil {
        return "", err
    }
    return string(h), nil
}

func checkPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    if err != nil {
        return false
    }
    return true
}
