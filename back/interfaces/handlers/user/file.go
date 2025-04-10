package user

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"github.com/jackc/pgx/v5"

	"github.com/Olyxz16/sherpa/domain/model"
	"github.com/Olyxz16/sherpa/infrastructure/persistence/repository"
)


type saveFileRequest struct {
    Source      string      `json:"source"`
    RepoName    string      `json:"repoName"`
    FileName    string      `json:"fileName"`
    Content     string      `json:"content"`
}

func SaveUserRepoFile(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("user").(*model.User)
    if !ok {
        w.WriteHeader(401)
        render.JSON(w, r, map[string]string {"message": "Unauthorized"})
		return
    }
   
    var sfr saveFileRequest
	err := render.DecodeJSON(r.Body, &sfr)
    if err != nil {
        w.WriteHeader(500)
        render.JSON(w, r, map[string]string {"message": "Bad request"})
		return
    }
    
	fileRepo := repository.NewFileRepository()
	source := model.AuthSource(sfr.Source)
	file := model.CreateFile(user, source, sfr.RepoName, sfr.FileName);

	err = file.Encrypt(sfr.Content)
    if err != nil {
        w.WriteHeader(500)
        render.JSON(w, r, map[string]string {"message": "Internal error"})
		return
    }

	err = fileRepo.Persist(file, r.Context())
    if err != nil {
        w.WriteHeader(500)
        render.JSON(w, r, map[string]string {"message": "Internal error"})
		return
    }

    render.JSON(w, r, map[string]string {"message": "OK"})
}

func FetchUserRepoFile(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("user").(*model.User)
    if !ok {
        w.WriteHeader(401)
        render.JSON(w, r, map[string]string {"message": "Unauthorized"})
		return
    }

    query := r.URL.Query()
    source := query.Get("source")
    if source == "" {
        w.WriteHeader(400)
        render.JSON(w, r, map[string]string {"message": "Missing source"})
		return
    }
    reponame := query.Get("repo")
    if reponame == "" {
        w.WriteHeader(400)
        render.JSON(w, r, map[string]string {"message": "Missing repository name"})
		return
    }
    filename := query.Get("file")
    if filename == "" {
        w.WriteHeader(400)
        render.JSON(w, r, map[string]string {"message": "Missing file name"})
		return
    }
    
	fileRepo := repository.NewFileRepository()
	file, err := fileRepo.Find(user, model.AuthSource(source), reponame, filename, r.Context())
    if err != nil && errors.Is(err, pgx.ErrNoRows) {
        render.JSON(w, r, map[string]string {"content": ""})
		return
	} else if err != nil {
        w.WriteHeader(404)
        render.JSON(w, r, map[string]string {"message": "Missing data"})
		return
    }

	content, err := file.Decrypt()
    if err != nil {
        w.WriteHeader(500)
        render.JSON(w, r, map[string]string {"message": "Internal error"})
		return
    }

    response := make(map[string]interface{})
    response["content"] = content
	render.JSON(w, r, response)
}
