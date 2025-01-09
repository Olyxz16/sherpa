package user

import (
	"encoding/json"

	"github.com/labstack/echo/v4"

	"github.com/Olyxz16/sherpa/database"
)


type saveFileRequest struct {
    Source      string      `json:"source"`
    RepoName    string      `json:"repoName"`
    FileName    string      `json:"fileName"`
    Content     string      `json:"content"`
}

func SaveUserRepoFile(c echo.Context) error {
    cookie, err := c.Cookie("session")
    if err != nil {
        return c.JSON(401, `{message: "Unauthorized"}`)
    }
    
    var sfr saveFileRequest
    err = c.Bind(&sfr)
    if err != nil {
        return c.JSON(500, `{message: "Bad request"}`)
    }
    
    err = database.SaveFile(cookie, sfr.Source, sfr.RepoName, sfr.FileName, sfr.Content)
    if err != nil {
        return c.JSON(500, `{message: "Missing data"}`)
    }
    return c.JSON(200, `{message: "OK"}`)
}

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
    jsonBytes, err := json.Marshal(response)
    if err != nil {
        return c.JSON(500, `{message: "Error fetch file"}`)
    }
    json := string(jsonBytes)
    return c.JSON(200, json)
}
