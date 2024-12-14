package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/Olyxz16/go-vue-template/handlers"
)

func RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Recover())
    e.Use(middleware.Logger())
    e.Static("/", staticDir)
    
    /* Static pages */
    e.GET("/", handlers.Index)
    e.GET("/about", handlers.Index)
    
    /* Api endpoints */

    /* Health checks */
    e.GET("/health", handlers.Health)

	return e
}
