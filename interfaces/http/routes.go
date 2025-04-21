package http

import (
	"io/fs"
	"net/http"

    "go.uber.org/zap"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/Olyxz16/sherpa/interfaces/handlers"
	"github.com/Olyxz16/sherpa/interfaces/handlers/github"
	"github.com/Olyxz16/sherpa/interfaces/handlers/user"
    mw "github.com/Olyxz16/sherpa/interfaces/middleware"
)

func RegisterRoutes(fs fs.FS) http.Handler {
    r := chi.NewRouter()
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)
	r.Use(mw.Session)

    staticFS := http.FileServer(http.FS(fs))
    r.Handle("/*", staticFS)
    
    /* Static pages */
	r.Get("/", staticHandler(fs))
    r.Get("/login", staticHandler(fs))
    r.Get("/welcome", staticHandler(fs))
    
	r.Route("/auth", func(r chi.Router) {
		r.Post("/masterkey", user.SetUserMasterkey)
		r.Get("/github/callback", github.AuthGithubLogin)
	})
	
	r.Group(func (r chi.Router) {
		r.Use(mw.Auth)

		r.Get("/user", user.FetchUser)
		r.Get("/file", user.FetchUserRepoFile)
		r.Post("/file", user.SaveUserRepoFile)
	})

    /* Health checks */
    r.Get("/health", handlers.Health)

	return r
}

func staticHandler(staticFS fs.FS) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		content, err := fs.ReadFile(staticFS, "index.html")
		if err != nil {
			return
		}
		_, err = w.Write(content)
		if err != nil {
			zap.L().Warn(err.Error())
			return
		}
	}
}
