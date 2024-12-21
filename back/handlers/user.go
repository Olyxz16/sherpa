package handlers

import (
	"encoding/json"
	"strings"

	"github.com/labstack/echo/v4"

	"github.com/Olyxz16/go-vue-template/database"
	"github.com/Olyxz16/go-vue-template/handlers/auth"
)


func FetchUser(c echo.Context) error {
    authHeader := c.Request().Header.Get("Authorization")
    if authHeader == "" {
        return c.JSON(401, "{message: Unauthorized}")
    }
    sessionCookie := strings.Split(authHeader, " ")[1]
    access_token, err := database.TokenFromCookie(sessionCookie)
    if err != nil {
        return c.JSON(401, "{message: Unauthorized}")
    }
    userData, err := auth.GetUserData(access_token) 
    if err != nil {
        return c.JSON(500, "{message: Internal error}")
    }
    json, err := json.Marshal(userData)
    if err != nil {
        return c.JSON(500, "{message: Internal error}")
    }
    return c.JSON(200, string(json))
}
