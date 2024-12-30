package database

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var (
    cookie_size = 24
    // 1 year duration | TODO : make actual definition
    cookie_duration = time.Now().AddDate(1 , 0 , 0)
)


func generateUserCookie() (*http.Cookie, error) {
    val, err := randLetterString()
    if err != nil {
        return nil, err
    }
    cookie := cookieFromKey(val)
    return cookie, nil
}

func randLetterString() (string, error) {
    token := make([]byte, cookie_size)
    _, err := rand.Read(token)
    if err != nil {
        return "", err
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
    return val, nil
}

func cookieFromKey(key string) (*http.Cookie) {
    expires := cookie_duration
    fmt.Printf("%v\n", expires)
    return &http.Cookie {
        Name: "session",
        Value: key,
        Path: "/",
        Expires: expires,
    } 
}
func marshalCookie(cookie *http.Cookie) (string, error) {
    jsonStr, err := json.Marshal(cookie.Value)
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
    var cookieStr string
    err = json.Unmarshal(decodedText, &cookieStr)
    cookie := cookieFromKey(cookieStr)
    return cookie, err
}
