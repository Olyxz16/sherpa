package crypto

import (
	"crypto/aes"
	"crypto/cipher"
    "crypto/rand"
	"encoding/base64"
	"io"
)

func EncryptFile(key []byte, content string) (b64content, b64nonce string, err error) {
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

func DecryptFile(key []byte, b64nonce, b64content string) (content string, err error) {
    encryptedContent, err := base64.StdEncoding.DecodeString(b64content)
    if err != nil {
        return  "", err
    }

    nonce, err := base64.StdEncoding.DecodeString(b64nonce)

    block, err := aes.NewCipher([]byte(key))
    if err != nil {
        return "", err
    }

    aesgcm, err := cipher.NewGCM(block)
    if err != nil {
        return "", err
    }

    contentBytes, err := aesgcm.Open(nil, nonce, encryptedContent, nil)
    if err != nil {
        return "", err
    }
    content = string(contentBytes)
    return content, nil
}
