package http

import (
	"fmt"
	"io/fs"
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/Olyxz16/sherpa/config"
)


func NewServer(fs fs.FS, cfg config.ServerConfig) *http.Server {
    // Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Handler:      RegisterRoutes(fs),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
    zap.L().Sugar().Infof("Starting server at %s:%d ...", cfg.Host, cfg.Port)

	return server
}
