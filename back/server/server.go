package server

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"github.com/Olyxz16/go-vue-template/database"
	"github.com/Olyxz16/go-vue-template/handlers"
)

var (
    host string
    port int
    staticDir string
    db database.Service
)


func NewServer() *http.Server {
	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", host, port),
		Handler:      RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
    slog.Info(fmt.Sprintf("Starting server at %s:%d ...", host, port))

    db = database.New()

	return server
}

func init() {
    var err error

    host = strings.Trim(os.Getenv("HOST"), " ")
    portStr := strings.Trim(os.Getenv("PORT"), " ")
    port, err = strconv.Atoi(portStr)
    if err != nil {
        panic("Cannot parse port !")
    }
    
    staticDir = strings.Trim(os.Getenv("STATIC_DIR"), " ")
    staticDir, err = filepath.Abs(staticDir)
    if err != nil {
        panic("Error parsing filepath")
    }
    isDebugMode := strings.Trim(os.Getenv("DEBUG"), " ")
    if staticDir == "" {
        if isDebugMode == "true" {
            slog.Warn("API MODE IS ACTIVE. Add --staticFilepath flag to serve static file")
        } else {
            panic("STATIC_DIR env is not set ! Cannot serve static files !")
        }
    } else {
        handlers.StaticDir = staticDir
    }
}
