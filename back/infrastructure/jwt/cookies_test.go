package jwt

import (
	"testing"
	"time"

	"github.com/Olyxz16/sherpa/domain/model"
)

func TestJWTReadWrite(t *testing.T) {
    user_id := 645135435
    user := model.NewUser(user_id, "", "", "", "")

    cookie, err := GenerateSessionCookie(user)
    if err != nil {
        t.Fatal("Error generating cookie", err)
    }

    actual_id, err := ParseSessionCookie(cookie)
    if err != nil {
        t.Fatal("Error parsing cookie", err)
    }

    if user_id != actual_id {
        t.Fatal("ids don't match", user_id, actual_id)
    }
}

func TestJWTRefresh(t *testing.T) {
    user_id := 645135435
    user := model.NewUser(user_id, "", "", "", "")

    cookie, err := GenerateSessionCookie(user)
    if err != nil {
        t.Fatal("Error generating cookie", err)
    }

    duration := cookie.Expires.Sub(time.Now())
    cookie.Expires = time.Now()
    cookie = RefreshCookie(cookie)
    if !approxEquals(duration, cookieDuration) {
        t.Fatal("Wrong duration", duration) 
    }
}

func TestJWTClean(t *testing.T) {
    cookie := CleanCookie()
    duration := cookie.Expires.Sub(time.Now())
    if !approxEquals(duration, 0) {
        t.Fatal("Wrong duration", duration)
    }
}

func approxEquals(t1, t2 time.Duration) bool {
    return (t1-t2).Abs() < 100*time.Millisecond
}
