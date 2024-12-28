package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/Olyxz16/go-vue-template/handlers"
	"github.com/Olyxz16/go-vue-template/handlers/github"
)

func RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Recover())
    e.Use(middleware.Logger())
    e.Static("/", staticDir)
    
    /* Static pages */
    e.GET("/", handlers.Index)
    e.GET("/login", handlers.Index)
    
    /* Api endpoints */
    /* Auth */
    e.GET("/auth/github/callback", github.AuthGithubLogin)

    e.GET("/user", handlers.FetchUser)

    /* Health checks */
    e.GET("/health", handlers.Health)

	return e
}
