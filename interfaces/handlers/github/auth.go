package github

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/go-chi/render"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"

	"github.com/Olyxz16/sherpa/domain/model"
	"github.com/Olyxz16/sherpa/infrastructure/jwt"
	"github.com/Olyxz16/sherpa/infrastructure/persistence/repository"
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
		zap.L().DPanic("Error parsing github callback code")
        http.Redirect(w, r, "/", 401)
		return
    }

    token, err := exchangeCode(code)
    if err != nil {
		zap.L().DPanic("Error exchanging code", zap.Error(err))
        http.Redirect(w, r, "/", 401)
		return
    }

    data := &UserData{}
    err = getUserName(token.AccessToken, data)
    if err != nil {
		zap.L().DPanic("Error fetching user data", zap.Error(err))
        http.Redirect(w, r, "/", 401)
		return
    }

	authRepo := repository.NewAuthRepository()
	auth, err := authRepo.Find(data.PlatformID, r.Context())
	isNewUser := err != nil && errors.Is(err, pgx.ErrNoRows)
	if isNewUser {
		// New user
		userRepo := repository.NewUserRepository()
		user := model.CreateUser(data.Username)
		auth = model.NewAuth(
			data.PlatformID,
			user,
			model.Github,
			token.AccessToken,
			token.RefreshToken,
			token.ExpiresIn,
			token.RefExpiresIn,
		)
		defer authRepo.Persist(auth, r.Context())
		defer userRepo.Persist(user, r.Context())
	} else if err != nil {
		zap.L().DPanic("Error fetching user auth", zap.Error(err))
		render.JSON(w, r, map[string]string{"message": "Internal server error"})
		return
	} else {
		auth.Refresh(token.AccessToken, token.RefreshToken, token.ExpiresIn, token.RefExpiresIn)
		defer authRepo.Persist(auth, r.Context())
	}

	cookie, err := jwt.GenerateSessionCookie(auth.GetUser())
	if err != nil {
		zap.L().DPanic("Error generating cookie", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"message": "Internal server error"})	
		return
	}
	http.SetCookie(w, cookie)
	if isNewUser || auth.GetUser().IsNew() {
		http.Redirect(w, r, "/welcome", 302)
	} else {
    	http.Redirect(w, r, "/", 302)
	}
}

func exchangeCode(code string) (*githubAccessToken, error) {
    params := url.Values{}
    params.Set("client_id", os.Getenv("GITHUB_ID"))
    params.Set("client_secret", os.Getenv("GITHUB_SECRET"))
    params.Set("code", code)
    
	url := fmt.Sprintf("https://github.com/login/oauth/access_token?%s", params.Encode())
    req, err := http.NewRequest("POST", url, nil)
    if err != nil {
        zap.L().DPanic("[Echange code] : request", zap.Error(err))
        return nil, err
    }
    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    req.Header.Add("Accept", "application/json")
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        zap.L().DPanic("[Exchange code] : response", zap.Error(err))
        return nil, err
    }
    defer resp.Body.Close()

    var data githubAccessToken
    err = json.NewDecoder(resp.Body).Decode(&data)
    if err != nil {
        return nil, err
    }
    return &data, nil
}
