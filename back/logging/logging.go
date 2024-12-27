package logging

import (
	"log/slog"
	"os"
    "runtime/debug"

	_ "github.com/joho/godotenv/autoload"
)

var (
    debugMode bool
)

func init() {
    _, debugMode = os.LookupEnv("DEBUG")
    var logger *slog.Logger
    if debugMode {
        logger = slog.Default()
    } else {
        logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
    }
    
    slog.SetDefault(logger)
}

func ErrLog(msg string, args ...any) {
    slog.Error(msg, args)
    if debugMode {
        debug.PrintStack()
    }
}
