package jwt

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/Olyxz16/sherpa/domain/model"
	"github.com/golang-jwt/jwt/v5"
)

type claims struct {
	Uid	int `json:"uid"`
	jwt.RegisteredClaims
}

var (
	cookieDuration = 2 * time.Minute
	jwtKey = []byte(os.Getenv("JWT_KEY"))
)

func GenerateSessionCookie(user *model.User) (*http.Cookie, error) {
	expirationTime := time.Now().Add(cookieDuration)
	claims := &claims {
		Uid: user.GetID(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return nil, err
	}

	cookie := &http.Cookie{
		Name: "session",
		Value: tokenString,
		Expires: expirationTime,
		Path: "/",
		HttpOnly: false,
		Secure: true,
		SameSite: http.SameSiteStrictMode,
	}

	return cookie, nil
}

func ParseSessionCookie(cookie *http.Cookie) (int, error) {
	tokenString := cookie.Value
	claims := &claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		return jwtKey, nil
	})
	if err != nil {
		return 0, err
	}
	if !token.Valid {
		return 0, errors.New("Invalid token")
	}

	return claims.Uid, nil
}

func RefreshCookie(cookie *http.Cookie) (*http.Cookie) {
	cookie.Expires = time.Now().Add(cookieDuration)
	cookie.Path = "/"
	return cookie
}

func CleanCookie() (*http.Cookie) {
	return &http.Cookie{
		Name: "session",
		Value: "",
		Path: "/",
		Expires: time.Now(),
	}
}
