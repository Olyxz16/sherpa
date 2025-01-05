package utils

import (
	"crypto/aes"
	"crypto/cipher"
    "crypto/rand"
	"encoding/base64"
	"io"
)


func EncryptFile(key, content string) (string, string, error) {
    block, err := aes.NewCipher([]byte(key))
    if err != nil {
        return "", "", err
    }
        
    nonce := make([]byte, 12)
    if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
        return "", "", err
    }

    aesgcm, err := cipher.NewGCM(block)
    if err != nil {
        return "", "", err
    }

    encodedTextBytes := aesgcm.Seal(nil, nonce, []byte(content), nil)
    encodedText := base64.StdEncoding.EncodeToString(encodedTextBytes)
    encodedNonce := base64.StdEncoding.EncodeToString(nonce)
    return encodedText, encodedNonce, nil
}

func DecryptFile(key, encodedNonce, encodedContent string) (string, error) {
    content, err := base64.StdEncoding.DecodeString(encodedContent)
    if err != nil {
        return  "", err
    }
    nonce, err := base64.StdEncoding.DecodeString(encodedNonce)
    if err != nil {
        return  "", err
    }

    block, err := aes.NewCipher([]byte(key))
    if err != nil {
        return "", err
    }

    aesgcm, err := cipher.NewGCM(block)
    if err != nil {
        return "", err
    }

    decodedContentBytes, err := aesgcm.Open(nil, nonce, content, nil)
    if err != nil {
        return "", err
    }
    decodedContent := string(decodedContentBytes)
    return decodedContent, nil
}
