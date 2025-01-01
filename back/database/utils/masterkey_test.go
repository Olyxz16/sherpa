package utils

import (
    "testing"
)

func TestPasswordDecoding(t *testing.T) {
    password := "somePassword111"
    hash, _, _, err := HashFromMasterkey(password)
    if err != nil {
        t.Fatalf("Error hashing password : %v", err)
    }
    matches, err := CompareHashAndPassword(password, hash)
    if err != nil {
        t.Fatalf("Error comparing hash and password : %v", err)
    }
    if !matches {
        t.Fatal(`Hash and password do not match`)
    }
}
