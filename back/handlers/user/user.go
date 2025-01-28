package user

import (
	"encoding/json"

	"github.com/labstack/echo/v4"

	"github.com/Olyxz16/sherpa/database"
	"github.com/Olyxz16/sherpa/handlers/github"
)


type masterkeyRequest struct {
    Masterkey   string      `json:"masterkey"`
}

func FetchUser(c echo.Context) error {
    // handle data source
    source := "github.com"
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

func SetUserMasterkey(c echo.Context) error {
    var mkr masterkeyRequest
    err := c.Bind(mkr)
    if err != nil {
        return c.JSON(400, `{message: "Missing masterkey"}`)
    }

    cookie, err := c.Cookie("session")
    if err != nil {
        return c.JSON(401, `{message: "Missing cookie"}`)
    }
    err = database.SetUserMasterkey(cookie, mkr.Masterkey)      
    if err != nil {
        return c.JSON(500, `{message: "Error"}`)
    }
    return c.JSON(200, `{message: "OK"}`)
}
