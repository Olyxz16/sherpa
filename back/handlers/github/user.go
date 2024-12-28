package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

    "github.com/Olyxz16/go-vue-template/logging"
)


type UserData struct {
    PlatformID          int         `json:"userId"`
    Username            string      `json:"username"`
    AvatarUrl           string      `json:"avatarUrl"`
    RepoNames           []string    `json:"repositories"`
}

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
        logging.ErrLog(fmt.Sprintf("GetUserName: %v", err.Error()))
        return err
    }
    defer resp.Body.Close()

    var body map[string]interface{}
    err = json.NewDecoder(resp.Body).Decode(&body)
    if err != nil {
        logging.ErrLog(fmt.Sprintf("GetUserName: %v", err.Error()))
        return err
    }
    
    data.PlatformID = int(body["id"].(float64))
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
                // Handle parsing error
                continue
            }
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
