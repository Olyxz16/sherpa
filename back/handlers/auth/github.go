package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/labstack/echo/v4"

	"github.com/Olyxz16/go-vue-template/database"
)


type UserData struct {
    Username            string
    AvatarUrl           string
    RepoNames           []string
}


func AuthGithubLogin(c echo.Context) error {
    code := c.QueryParam("code")
    if code == "" {
        slog.Error("Auth : missing code")
        c.QueryParams().Add("autherr", "github")
        return c.Redirect(302, "/")
    }
    
    auth, err := exchangeCode(code)
    if err != nil {
        slog.Error(fmt.Sprintf("Auth : %v", err))
        c.QueryParams().Add("autherr", "github")
        return c.Redirect(302, "/")
    }
    
    data, err := getUserData(auth.Access_token)
    if err != nil {
        slog.Error(fmt.Sprintf("Auth : %v", err))
        c.QueryParams().Add("autherr", "github")
        return c.Redirect(302, "/")
    }
    fmt.Print(data)
    
    cookie := &http.Cookie{ Name: "session", Value: string(code) }
    http.SetCookie(c.Response(), cookie)
    
    return c.Redirect(302, "/")
}



func getUserData(access_token string) (UserData, error) {
    userData := UserData{}
    err := getUserName(access_token, &userData)
    if err != nil {
        slog.Error(fmt.Sprintf("GetUserData: %v", err.Error()))
        return UserData{}, err
    }
    err = getUserRepos(access_token, &userData)
    if err != nil {
        slog.Error(fmt.Sprintf("GetUserData: %v", err.Error()))
        return UserData{}, err
    }
    return userData, nil
}

func getUserName(access_token string, data *UserData) (error) {
    url := "https://api.github.com/user"
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return err
    }
    req.Header.Add("Accept", "application/json")
    req.Header.Add("Content-Type", "application/json")
    req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", access_token))
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    var result map[string]interface{}
    json.NewDecoder(resp.Body).Decode(&result)

    data.Username = result["login"].(string)
    data.AvatarUrl = result["avatar_url"].(string)
    return nil
}
func getUserRepos(access_token string, data *UserData) (error) {
    url := "https://api.github.com/user/repos"
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return err
    }
    req.Header.Add("Accept", "*/*")
    req.Header.Add("Content-Type", "application/vnd.github.v3+json")
    req.Header.Add("User-Agent", "curl/7.64.0")
    req.Header.Add("Authorization", fmt.Sprintf("token %s", access_token))
    req.Header.Add("X-Accepted-GitHub-Permissions", "metadata=read")
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return err
    }

    var result []map[string]interface{}
    err = json.Unmarshal(body, &result)
    if err != nil {
        return err
    }

    names := []string{}
    for _, repo := range result {
        repoName := repo["name"].(string)
        names = append(names, repoName)
    }
    data.RepoNames = names

    return nil
}

func exchangeCode(code string) (database.GithubAuth, error) {
    params := url.Values{}
    params.Set("client_id", os.Getenv("CLIENT_ID"))
    params.Set("client_secret", os.Getenv("CLIENT_SECRET"))
    params.Set("code", code)
    
    url := "https://github.com/login/oauth/access_token"
    req, err := http.NewRequest("POST", url, strings.NewReader(params.Encode()))
    if err != nil {
        return database.GithubAuth{}, err
    }
    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    req.Header.Add("Accept", "application/json")
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return database.GithubAuth{}, err
    }

    var data map[string]interface{}
    json.NewDecoder(resp.Body).Decode(&data)
    expires_in := data["expires_in"].(float64)
    refresh_expires_in := data["refresh_token_expires_in"].(float64)
    result := database.GithubAuth {
        Access_token: data["access_token"].(string),
        Refresh_token: data["refresh_token"].(string),
        Expires_in: expires_in,
        Refresh_expires_in: refresh_expires_in,
    }
    return result, nil
}
