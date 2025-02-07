package github

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/Olyxz16/sherpa/database"
	"github.com/Olyxz16/sherpa/logging"
)


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
    
    userAuth, isNew, err := database.AuthenticateUser(*platformAuth)
    if err != nil {
        http.Redirect(w, r, "/auth_error", 302)
    }

    http.SetCookie(w, userAuth.Cookie)
    if isNew {
        http.Redirect(w, r, "/welcome", 302)
    }
    http.Redirect(w, r, "/", 302)
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
