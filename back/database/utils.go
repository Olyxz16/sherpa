package database

import (
	"crypto/rand"
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
