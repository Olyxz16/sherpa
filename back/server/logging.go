package server

import (
    "os"
    "log/slog"

	_ "github.com/joho/godotenv/autoload"
)

func init() {
    var logger *slog.Logger
    if os.Getenv("DEBUG") == "true" {
        logger = slog.Default()
    } else {
        logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
    }
    slog.SetDefault(logger)
}
