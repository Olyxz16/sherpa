package utils

import (
    "testing"
)

func TestFileEncryption(t *testing.T) {
    key := []byte("ighjenoqiknzjiuhankplskinshutgnt")
    expected := "filecontent"
    b64content, b64nonce, err := EncryptFile(key, expected)
    if err != nil {
        t.Fatalf("Error encrypting file : %v", err)
    }
    actual, err := DecryptFile(key, b64nonce, b64content)
    if err != nil {
        t.Fatalf("Error decrypting file : %v", err)
    }
    if expected != actual {
        t.Fatalf(`Contents don't match
                expected : %v
                actual : %v`, expected, actual)
    }
}
