package controller

import (
	"fmt"
	"io/fs"
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/Olyxz16/sherpa/config"
	"github.com/Olyxz16/sherpa/model"
)

var (
    staticDir fs.FS
)


func NewServer(fs fs.FS) *http.Server {
    staticDir = fs
    
    cfg, err := config.NewServerConfig()
    if err != nil {
        panic("Error loading server config")
    }
    model.New()
	
    // Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Handler:      RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
    zap.L().Sugar().Infof("Starting server at %s:%d ...", cfg.Host, cfg.Port)

	return server
}
