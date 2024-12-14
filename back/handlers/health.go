package handlers

import (
	"log/slog"

	"github.com/labstack/echo/v4"

	"github.com/Olyxz16/go-vue-template/database"
)

/*********************/
/* Healthcheck utils */
/*********************/

func Health(c echo.Context) error {
    if !database.New().Health() {
        slog.Error("HEALTHCHECK NOT PASSING !")
        return c.JSON(500, map[string]string {"message": "KO"})
    }
    return c.JSON(200, map[string]string {"message": "OK"})
}
