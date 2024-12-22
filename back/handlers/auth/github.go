package auth

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/labstack/echo/v4"

	"github.com/Olyxz16/go-vue-template/database"
)


type UserData struct {
    UserID              int         `json:"userId"`
    Username            string      `json:"username"`
    AvatarUrl           string      `json:"avatarUrl"`
    RepoNames           []string    `json:"repositories"`
}

// TODO: properly handle errors
func AuthGithubLogin(c echo.Context) error {
    code := c.QueryParam("code")
    if code == "" {
        slog.Error("Auth : missing code")
        c.QueryParams().Add("autherr", "github")
        return c.Redirect(302, "/")
    }

    auth, cookie, err := exchangeCode(code)
    if err != nil {
        slog.Error(fmt.Sprintf("Auth : %v", err))
        c.QueryParams().Add("autherr", "github")
        return c.Redirect(302, "/")
    }

    data := &UserData{}
    err = getUserName(auth.Access_token, data)
    if err != nil {
        slog.Error(fmt.Sprintf("Auth : %v", err))
        c.QueryParams().Add("autherr", "github")
        return c.Redirect(302, "/")
    }

    auth.UserId = data.UserID
    
    err = database.AddGithubUser(*auth)
    if err != nil {
        slog.Error(fmt.Sprintf("Auth : %v", err))
        c.QueryParams().Add("autherr", "github")
        return c.Redirect(302, "/")
    }
    
    http.SetCookie(c.Response(), cookie)
    return c.Redirect(302, "/")
}



func GetUserData(access_token string) (UserData, error) {
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

    var body map[string]interface{}
    err = json.NewDecoder(resp.Body).Decode(&body)
    if err != nil {
        return err
    }
    
    data.UserID = int(body["id"].(float64))
    data.Username = body["login"].(string)
    data.AvatarUrl = body["avatar_url"].(string)

    return nil
}
func getUserRepos(access_token string, data *UserData) (error) {
    result := make([]string, 0)
    url := "https://api.github.com/user/repos"
    nextPage := func() bool { 
        req, err := http.NewRequest("GET", url, nil)
        if err != nil {
            return false
        }
        req.Header.Add("Accept", "*/*")
        req.Header.Add("Content-Type", "application/vnd.github.v3+json")
        req.Header.Add("User-Agent", "curl/7.64.0")
        req.Header.Add("Authorization", fmt.Sprintf("token %s", access_token))
        req.Header.Add("X-Accepted-GitHub-Permissions", "metadata=read")
        resp, err := http.DefaultClient.Do(req)
        if err != nil {
            return false
        }
        defer resp.Body.Close()

        var body []map[string]interface{}
        err = json.NewDecoder(resp.Body).Decode(&body)
        if err != nil {
            return false
        }

        for _, v := range body {
            name := v["name"].(string)
            result = append(result, name)
        }

        next := resp.Header.Get("Link")
        url = next[strings.Index(next, "<")+1 : strings.Index(next, ">")]
        return url[len(url)-1] != '1'
    }
    
    for nextPage() {
    }
    data.RepoNames = result
    return nil
}

func exchangeCode(code string) (*database.GithubAuth, *http.Cookie, error) {
    params := url.Values{}
    params.Set("client_id", os.Getenv("CLIENT_ID"))
    params.Set("client_secret", os.Getenv("CLIENT_SECRET"))
    params.Set("code", code)
    
    url := "https://github.com/login/oauth/access_token"
    req, err := http.NewRequest("POST", url, strings.NewReader(params.Encode()))
    if err != nil {
        return nil, nil, err
    }
    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    req.Header.Add("Accept", "application/json")
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return nil, nil, err
    }
    
    cookie, err := generateCookie()
    if err != nil {
        return nil, nil, err
    }

    var data map[string]interface{}
    json.NewDecoder(resp.Body).Decode(&data)
    expires_in := data["expires_in"].(float64)
    refresh_expires_in := data["refresh_token_expires_in"].(float64)
    result := &database.GithubAuth {
        Cookie: cookie.Value,
        Access_token: data["access_token"].(string),
        Refresh_token: data["refresh_token"].(string),
        Expires_in: int(expires_in),
        Refresh_expires_in: int(refresh_expires_in),
    }
    return result, cookie, nil
}
