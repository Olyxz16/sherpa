package database

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"net/http"
)

const (
    cookie_size = 24
)

// TODO: Implement expiration
func generateUserCookie() (*http.Cookie, error) {
    token := make([]byte, cookie_size)
    _, err := rand.Read(token)
    if err != nil {
        return nil, err
    }

    for i, v := range token {
        newv := v % 52
        if newv <= 25 {
            token[i] = newv + 'a' 
        } else {
            token[i] = newv -26 + 'A'
        }
    }

    val := string(token)
    cookie := &http.Cookie{ Name: "session", Value: val, Path: "/" }
    return cookie, nil
}

func marshalCookie(cookie *http.Cookie) (string, error) {
    jsonStr, err := json.Marshal(cookie)
    if err != nil {
        return "", err
    }
    encodedText := base64.StdEncoding.EncodeToString([]byte(jsonStr))
    return encodedText, nil
}

func unmarshalCookie(str string) (*http.Cookie, error) {
    decodedText, err := base64.StdEncoding.DecodeString(str)
    if err != nil {
        return nil, err
    }
    cookie := &http.Cookie{}
    err = json.Unmarshal(decodedText, cookie)
    return cookie, err
}
