package user

import (
	"encoding/json"

	"github.com/labstack/echo/v4"

	"github.com/Olyxz16/go-vue-template/database"
)


func FetchUserRepoFile(c echo.Context) error {
    cookie, err := c.Cookie("session")
    if err != nil {
        return c.JSON(401, `{message: "Unauthorized"}`)
    }
    source := c.QueryParam("source")
    if source == "" {
        return c.JSON(400, `{message: "Missing source"}`)
    }
    repoName := c.QueryParam("repo")
    if repoName == "" {
        return c.JSON(400, `{message: "Missing repository name"}`)
    }
    fileName := c.QueryParam("file")
    if fileName == "" {
        return c.JSON(400, `{message: "Missing file name"}`)
    }
    
    content, err := database.FetchFileContent(cookie, source, repoName, fileName)
    if err != nil {
        // handle different errors
        return c.JSON(500, `{message: "Missing data"}`)
    }

    response := make(map[string]interface{})
    response["content"] = content
    json, err := json.Marshal(response)
    if err != nil {
        return c.JSON(500, `{message: "Error fetch file"}`)
    }
    return c.JSON(200, json)
}
