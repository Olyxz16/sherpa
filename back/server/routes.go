package server

import (
    "io/fs"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/Olyxz16/go-vue-template/handlers"
	"github.com/Olyxz16/go-vue-template/handlers/user"
	"github.com/Olyxz16/go-vue-template/handlers/github"
)

func RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Recover())
    e.Use(middleware.Logger())
    e.StaticFS("/", staticDir)
    
    /* Static pages */
    e.GET("/", staticHandler)
    e.GET("/login", staticHandler)
    e.GET("/welcome", staticHandler)
    
    /* Api endpoints */
    /* Auth */
    e.POST("/auth/masterkey", user.SetUserMasterkey)
    e.GET("/auth/github/callback", github.AuthGithubLogin)

    e.GET("/user", user.FetchUser)
    e.GET("/file", user.FetchUserRepoFile)
    e.POST("/file", user.SaveUserRepoFile)

    /* Health checks */
    e.GET("/health", handlers.Health)

	return e
}

func staticHandler(c echo.Context) error {
    content, err := fs.ReadFile(staticDir, "index.html")
    if err != nil {
        return err
    }
    err = c.HTMLBlob(200, content)
    if err != nil {
        slog.Warn(err.Error())
        return err
    }
    return nil
}
