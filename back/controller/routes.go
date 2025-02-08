package controller

import (
	"io/fs"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/Olyxz16/sherpa/handlers"
	"github.com/Olyxz16/sherpa/handlers/github"
	"github.com/Olyxz16/sherpa/handlers/user"
)

func RegisterRoutes() http.Handler {
    r := chi.NewRouter()
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)

    staticFS := http.FileServer(http.FS(staticDir))
    r.Handle("/*", staticFS)
    
    /* Static pages */
    r.Get("/", staticHandler)
    r.Get("/login", staticHandler)
    r.Get("/welcome", staticHandler)
    
    /* Api endpoints */
    /* Auth */
    r.Route("/auth", func(r chi.Router) {
        r.Post("/masterkey", user.SetUserMasterkey)
        r.Get("/github/callback", github.AuthGithubLogin)
    })

    r.Get("/user", user.FetchUser)
    r.Get("/file", user.FetchUserRepoFile)
    r.Post("/file", user.SaveUserRepoFile)

    /* Health checks */
    r.Get("/health", handlers.Health)

	return r
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
    content, err := fs.ReadFile(staticDir, "index.html")
    if err != nil {
        return
    }
    _, err = w.Write(content)
    if err != nil {
        slog.Warn(err.Error())
        return
    }
}
