package user

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"

	"github.com/Olyxz16/sherpa/model"
	"github.com/Olyxz16/sherpa/handlers/github"
)


type masterkeyRequest struct {
    Masterkey   string      `json:"masterkey"`
}

func FetchUser(w http.ResponseWriter, r *http.Request) {
    // handle data source
    source := "github.com"
    cookie, err :=  r.Cookie("session")
    if err != nil {
        w.WriteHeader(401)
        render.JSON(w, r, map[string]string {"message": "Missing cookie"})
    }
    access_token, err := model.TokenFromCookie(cookie, source)
    if err != nil {
        w.WriteHeader(401)
        render.JSON(w, r, map[string]string {"message": "Unauthorized"})
    }
    userData, err := github.GetUserData(access_token) 
    if err != nil {
        w.WriteHeader(500)
        render.JSON(w, r, map[string]string {"message": "Internal error"})
    }
    json, err := json.Marshal(userData)
    if err != nil {
        w.WriteHeader(500)
        render.JSON(w, r, map[string]string {"message": "Internal error"})
    }
    w.Write(json)
}

func SetUserMasterkey(w http.ResponseWriter, r *http.Request) {
    var mkr masterkeyRequest
    err := render.DecodeJSON(r.Body, &mkr)
    if err != nil {
        w.WriteHeader(400)
        render.JSON(w, r, map[string]string {"message": "Missing masterkey"})
    }

    cookie, err := r.Cookie("session")
    if err != nil {
        w.WriteHeader(401)
        render.JSON(w, r, map[string]string {"message": "Missing cookie"})
    }
    err = model.SetUserMasterkey(cookie, mkr.Masterkey)      
    if err != nil {
        w.WriteHeader(500)
        render.JSON(w, r, map[string]string {"message": "Error"})
    }
    render.JSON(w, r, map[string]string {"message": "OK"})
}
