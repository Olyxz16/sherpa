package handlers

import (
	"log/slog"
	"os"

	"github.com/labstack/echo/v4"
)

var (
    StaticDir string
)

/* All static pages should handle api mode */

func Index(c echo.Context) error {
    if os.Getenv("API_MODE") == "true" {
        return APIModeResponse(c)
    }
    err := c.File(StaticDir + "/index.html")
    if err != nil {
        slog.Warn(err.Error())
    }
    return err
}
