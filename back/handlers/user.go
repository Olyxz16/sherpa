package handlers

import (
	"encoding/json"

	"github.com/labstack/echo/v4"

	"github.com/Olyxz16/go-vue-template/database"
	"github.com/Olyxz16/go-vue-template/handlers/github"
)


func FetchUser(c echo.Context) error {
    source := c.QueryParam("source")
    if source == "" {
        return c.JSON(401, `{message: Missing source}`)
    }
    cookie, err :=  c.Cookie("session")
    if err != nil {
        return c.JSON(401, `{message: "Missing cookie"}`)
    }
    access_token, err := database.TokenFromCookie(cookie, source)
    if err != nil {
        return c.JSON(401, "{message: Unauthorized}")
    }
    userData, err := github.GetUserData(access_token) 
    if err != nil {
        return c.JSON(500, "{message: Internal error}")
    }
    json, err := json.Marshal(userData)
    if err != nil {
        return c.JSON(500, "{message: Internal error}")
    }
    return c.JSON(200, string(json))
}
