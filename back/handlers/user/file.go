package user

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"

	"github.com/Olyxz16/sherpa/model"
)


type saveFileRequest struct {
    Source      string      `json:"source"`
    RepoName    string      `json:"repoName"`
    FileName    string      `json:"fileName"`
    Content     string      `json:"content"`
}

func SaveUserRepoFile(w http.ResponseWriter, r *http.Request) {
    cookie, err := r.Cookie("session")
    if err != nil {
        w.WriteHeader(401)
        render.JSON(w, r, map[string]string {"message": "Unauthorized"})
    }
    
    var sfr saveFileRequest
    err = render.DecodeJSON(r.Body, &sfr)
    if err != nil {
        w.WriteHeader(500)
        render.JSON(w, r, map[string]string {"message": "Bad request"})
    }
    
    err = model.SaveFile(cookie, sfr.Source, sfr.RepoName, sfr.FileName, sfr.Content)
    if err != nil {
        w.WriteHeader(500)
        render.JSON(w, r, map[string]string {"message": "Missing data"})
    }
    render.JSON(w, r, map[string]string {"message": "OK"})
}

func FetchUserRepoFile(w http.ResponseWriter, r *http.Request) {
    cookie, err := r.Cookie("session")
    if err != nil {
        w.WriteHeader(401)
        render.JSON(w, r, map[string]string {"message": "Unauthorized"})
    }
    query := r.URL.Query()
    source := query.Get("source")
    if source == "" {
        w.WriteHeader(400)
        render.JSON(w, r, map[string]string {"message": "Missing source"})
    }
    repoName := query.Get("repo")
    if repoName == "" {
        w.WriteHeader(400)
        render.JSON(w, r, map[string]string {"message": "Missing repository name"})
    }
    fileName := query.Get("file")
    if fileName == "" {
        w.WriteHeader(400)
        render.JSON(w, r, map[string]string {"message": "Missing file name"})
    }
    
    content, err := model.FetchFileContent(cookie, source, repoName, fileName)
    if err != nil {
        // handle different errors
        w.WriteHeader(500)
        render.JSON(w, r, map[string]string {"message": "Missing data"})
    }

    response := make(map[string]interface{})
    response["content"] = content
    json, err := json.Marshal(response)
    if err != nil {
        w.WriteHeader(500)
        render.JSON(w, r, map[string]string {"message": "Error fetch file"})
    }
    w.Write(json)
}
