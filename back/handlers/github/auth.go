package github

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/labstack/echo/v4"

    "github.com/Olyxz16/sherpa/logging"
	"github.com/Olyxz16/sherpa/database"
)


func AuthGithubLogin(c echo.Context) error {
    code := c.QueryParam("code")
    if code == "" {
        return c.Redirect(302, "/auth_error")
    }

    platformAuth, err := exchangeCode(code)
    if err != nil {
        return c.Redirect(302, "/auth_error")
    }

    data := &UserData{}
    err = getUserName(platformAuth.Access_token, data)
    if err != nil {
        return c.Redirect(302, "/auth_error")
    }
    
    platformAuth.PlatformId = data.PlatformID
    
    userAuth, isNew, err := database.AuthenticateUser(*platformAuth)
    if err != nil {
        return c.Redirect(302, "/auth_error")
    }
    
    http.SetCookie(c.Response(), userAuth.Cookie)
    if isNew {
        return c.Redirect(302, "/welcome")
    }
    return c.Redirect(302, "/")
}

func exchangeCode(code string) (*database.PlatformUserAuth, error) {
    params := url.Values{}
    params.Set("client_id", os.Getenv("CLIENT_ID"))
    params.Set("client_secret", os.Getenv("CLIENT_SECRET"))
    params.Set("code", code)
    
    url := "https://github.com/login/oauth/access_token"
    req, err := http.NewRequest("POST", url, strings.NewReader(params.Encode()))
    if err != nil {
        logging.ErrLog("[Exchange code] : request")
        return nil, err
    }
    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    req.Header.Add("Accept", "application/json")
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        logging.ErrLog("[Exchange code] : response")
        return nil, err
    }
    
    var data map[string]interface{}
    json.NewDecoder(resp.Body).Decode(&data)
    expires_in := data["expires_in"].(float64)
    refresh_expires_in := data["refresh_token_expires_in"].(float64)
    result := &database.PlatformUserAuth {
        Source: "github.com",
        Access_token: data["access_token"].(string),
        Refresh_token: data["refresh_token"].(string),
        Expires_in: int(expires_in),
        Refresh_expires_in: int(refresh_expires_in),
    }
    return result, nil
}
