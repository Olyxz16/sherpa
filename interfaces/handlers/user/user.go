package user

import (
	"net/http"

	"github.com/go-chi/render"
	"go.uber.org/zap"

	"github.com/Olyxz16/sherpa/domain/model"
	"github.com/Olyxz16/sherpa/infrastructure/persistence/repository"
	"github.com/Olyxz16/sherpa/interfaces/handlers/github"
)


type masterkeyRequest struct {
    Masterkey   string      `json:"masterkey"`
}

func FetchUser(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("user").(*model.User)
    if !ok {
		zap.L().Error("Wrong context")
        w.WriteHeader(401)
        render.JSON(w, r, map[string]string {"message": "Unauthorized"})
		return
    }

    /*query := r.URL.Query()
    source := query.Get("source")
    if source == "" {
        w.WriteHeader(400)
        render.JSON(w, r, map[string]string {"message": "Missing source"})
		return
    }*/
	source := "github.com"

	authRepo := repository.NewAuthRepository()
	auth, err := authRepo.FindByUser(user, model.AuthSource(source), r.Context())
    if err != nil {
        w.WriteHeader(500)
        render.JSON(w, r, map[string]string {"message": "Internal error"})
		return
    }

    userData, err := github.GetUserData(auth.GetAccessToken()) 
    if err != nil {
        w.WriteHeader(500)
        render.JSON(w, r, map[string]string {"message": "Internal error"})
		return
    }

    render.JSON(w, r, userData)
}

func SetUserMasterkey(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("user").(*model.User)
    if !ok {
		zap.L().Error("Wrong context")
        w.WriteHeader(401)
        render.JSON(w, r, map[string]string {"message": "Unauthorized"})
		return
    }

    var mkr masterkeyRequest
    err := render.DecodeJSON(r.Body, &mkr)
    if err != nil {
        w.WriteHeader(400)
        render.JSON(w, r, map[string]string {"message": "Missing masterkey"})
		return
    }
	
	userRepo := repository.NewUserRepository()
	err = user.SetUserMasterkey(mkr.Masterkey)
    if err != nil {
        w.WriteHeader(500)
        render.JSON(w, r, map[string]string {"message": "Internal error"})
		return
    }

	err = userRepo.Persist(user, r.Context())
	if err != nil {
		zap.L().DPanic("Error persisting user", zap.Error(err))
	}
	
    render.JSON(w, r, map[string]string {"message": "OK"})
}
