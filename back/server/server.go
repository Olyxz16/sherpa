package server

import (
    "io/fs"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"

	"github.com/Olyxz16/go-vue-template/database"
)

var (
    host string
    port int
    staticDir fs.FS
    db database.Service
)


func NewServer(fs fs.FS) *http.Server {
    staticDir = fs
    
    db = database.New()
	
    // Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", host, port),
		Handler:      RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
    slog.Info(fmt.Sprintf("Starting server at %s:%d ...", host, port))

	return server
}

func init() {
    var err error
    godotenv.Load()

    host = strings.Trim(os.Getenv("HOST"), " ")
    portStr := strings.Trim(os.Getenv("PORT"), " ")
    port, err = strconv.Atoi(portStr)
    if err != nil {
        panic("Cannot parse port !")
    }
}
