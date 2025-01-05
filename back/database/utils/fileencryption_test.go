package utils

import (
    "testing"
)

func TestFileEncryption(t *testing.T) {
    key := "ighjenoqiknzjiuhankplskinshutgnt"
    content := "filecontent"
    encryptedContent, encodedNonce, err := EncryptFile(key, content)
    if err != nil {
        t.Fatalf("Error encrypting file : %v", err)
    }
    t.Logf("%v %v", encryptedContent, encodedNonce)
    decryptedContent, err := DecryptFile(key, encodedNonce, encryptedContent)
    if err != nil {
        t.Fatalf("Error decrypting file : %v", err)
    }
    if decryptedContent != content {
        t.Fatalf(`Contents don't match
                expected : %v
                Actual : %v`, content, decryptedContent)
    }
}
