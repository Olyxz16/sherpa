package github

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/Olyxz16/sherpa/model"
	"github.com/Olyxz16/sherpa/logging"
)

type githubAccessToken struct {
    AccessToken     string  `json:"access_token"` 
    RefreshToken    string  `json:"refresh_token"`
    ExpiresIn       int     `json:"expires_in"`
    RefExpiresIn    int     `json:"refresh_token_expires_in"`
}

func AuthGithubLogin(w http.ResponseWriter, r *http.Request) {
    code := r.URL.Query().Get("code")
    if code == "" {
        http.Redirect(w, r, "/auth_error", 302)
    }

    platformAuth, err := exchangeCode(code)
    if err != nil {
        http.Redirect(w, r, "/auth_error", 302)
    }

    data := UserData{}
    err = getUserName(platformAuth.Access_token, &data)
    if err != nil {
        http.Redirect(w, r, "/auth_error", 302)
    }
    
    platformAuth.PlatformId = data.PlatformID
    
    userAuth, isNew, err := model.AuthenticateUser(*platformAuth)
    if err != nil {
        http.Redirect(w, r, "/auth_error", 302)
    }

    http.SetCookie(w, userAuth.Cookie)
    if isNew {
        http.Redirect(w, r, "/welcome", 302)
    }
    http.Redirect(w, r, "/", 302)
}

func exchangeCode(code string) (*model.PlatformUserAuth, error) {
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
    defer resp.Body.Close()

    var data githubAccessToken
    err = json.NewDecoder(resp.Body).Decode(&data)
    if err != nil {
        return nil, err
    }
    result := &model.PlatformUserAuth {
        Source: "github.com",
        Access_token: data.AccessToken,
        Refresh_token: data.RefreshToken,
        Expires_in: data.ExpiresIn,
        Refresh_expires_in: data.RefExpiresIn,
    }
    return result, nil
}
