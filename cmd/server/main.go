package main

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"os/signal"
	"time"

	"go.uber.org/zap"

	"github.com/Olyxz16/sherpa/config"
	"github.com/Olyxz16/sherpa/infrastructure/persistence"
	controller "github.com/Olyxz16/sherpa/interfaces/http"
)

//go:embed static/*
var f embed.FS

func main() {
    
    f, err := fs.Sub(f, "static")
    if err != nil {
        panic("Static folder static/ missing !") 
    }

	config.DefaultLogger()
    server := controller.NewServer(f, config.NewServerConfig())
	_ = persistence.New(config.NewDatabaseConfig())

    go func() {
        err := server.ListenAndServe()
        if err != http.ErrServerClosed {
            zap.L().Info("Server error !")
            os.Exit(1)
        }
    }()

    ch := make(chan os.Signal, 1)
    signal.Notify(
        ch,
        os.Interrupt,
    )

    <-ch
    zap.L().Error("os.Interrupt - shutting down...")
    go func() {
        <-ch
        zap.L().Error("os.Kill - terminating...")
    }()

    gracefulCtx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancelShutdown()

    err = server.Shutdown(gracefulCtx)
    if err != nil {
        zap.L().Error(fmt.Sprintf("Shutdown error : %v", err))
        os.Exit(1)
    }
    zap.L().Info("Gracefully stopped")

    os.Exit(0)
}
