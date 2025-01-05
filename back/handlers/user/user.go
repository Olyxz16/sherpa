package user

import (
	"encoding/json"

	"github.com/labstack/echo/v4"

	"github.com/Olyxz16/go-vue-template/database"
	"github.com/Olyxz16/go-vue-template/handlers/github"
)


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
    var body map[string]interface{}
    err := json.NewDecoder(c.Request().Body).Decode(&body)
    if err != nil {
        return c.JSON(401, `{message: "Missing masterkey"}`)
    }

    masterkey, ok := body["masterkey"].(string)
    if !ok {
        return c.JSON(401, `{message: "Missing masterkey"}`)
    }

    cookie, err := c.Cookie("session")
    if err != nil {
        return c.JSON(401, `{message: "Missing cookie"}`)
    }
    err = database.SetUserMasterkey(cookie, masterkey)      
    if err != nil {
        return c.JSON(500, `{message: "Error"}`)
    }
    return c.JSON(200, `{message: "OK"}`)
}
