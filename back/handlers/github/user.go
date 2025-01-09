package github

import (
	"fmt"
    "errors"
	"net/http"
	"strings"
	"encoding/json"

    "github.com/Olyxz16/sherpa/logging"
)


type UserData struct {
    PlatformID          int         `json:"userId"`
    Username            string      `json:"username"`
    AvatarUrl           string      `json:"avatarUrl"`
    RepoNames           []string    `json:"repositories"`
}

var (
    InvalidCookieError = errors.New("Cookie is invalid")
)


func GetUserData(access_token string) (UserData, error) {
    userData := UserData{}
    err := getUserName(access_token, &userData)
    if err != nil {
        return UserData{}, err
    }
    err = getUserRepos(access_token, &userData)
    if err != nil {
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

    err = parseUsername(body, data)
    if err != nil {
        return err
    }

    return nil
}

func parseUsername(body map[string]interface{}, data *UserData) (error) {
    var idJson float64
    var usernameJson string
    var avatarUrlJson string
    var ok bool

    if idJson, ok = body["id"].(float64) ; !ok {
        return InvalidCookieError
    }
    id := int(idJson)
    if usernameJson, ok = body["login"].(string) ; !ok {
        return InvalidCookieError
    }
    if avatarUrlJson, ok = body["avatar_url"].(string) ; !ok {
        return InvalidCookieError
    }
    
    data.PlatformID = id
    data.Username = usernameJson
    data.AvatarUrl = avatarUrlJson

    return nil    
}

func getUserRepos(access_token string, data *UserData) (error) {
    result := make([]string, 0)
    url := "https://api.github.com/user/repos"
    var ferr error
    nextPage := func() bool { 
        req, err := http.NewRequest("GET", url, nil)
        if err != nil {
            logging.ErrLog(fmt.Sprintf("GetUserRepos: %v", err.Error()))
            return false
        }
        req.Header.Add("Accept", "*/*")
        req.Header.Add("Content-Type", "application/vnd.github.v3+json")
        req.Header.Add("User-Agent", "curl/7.64.0")
        req.Header.Add("Authorization", fmt.Sprintf("token %s", access_token))
        req.Header.Add("X-Accepted-GitHub-Permissions", "metadata=read")
        resp, err := http.DefaultClient.Do(req)
        if err != nil {
            logging.ErrLog(fmt.Sprintf("GetUserRepos: %v", err.Error()))
            return false
        }
        defer resp.Body.Close()

        var body []map[string]interface{}
        err = json.NewDecoder(resp.Body).Decode(&body)
        if err != nil {
            logging.ErrLog(fmt.Sprintf("GetUserRepos: %v", err.Error()))
            return false
        }

        for _, v := range body {
            name, ok := v["name"].(string)
            if !ok {
                ferr = InvalidCookieError
                return false
            }
            result = append(result, name)
        }

        next := resp.Header.Get("Link")
        url = next[strings.Index(next, "<")+1 : strings.Index(next, ">")]
        return url[len(url)-1] != '1'
    }
    
    for nextPage() {
    }
    if ferr != nil {
        return ferr
    }
    data.RepoNames = result
    return nil
}
