package main

import (
	"context"
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Olyxz16/sherpa/server"
)

//go:embed static/*
var f embed.FS

func main() {
    
    healthCheckFlag := flag.Bool("healthcheck", false, "Performs a health check on the running service")
    flag.Parse()
    if *healthCheckFlag {
        checkHealth()
    }

    f, err := fs.Sub(f, "static")
    if err != nil {
        panic("Static folder static/ missing !") 
    }
    server := server.NewServer(f)

    go func() {
        err := server.ListenAndServe()
        if err != http.ErrServerClosed {
            slog.Error("Server error !")
            os.Exit(1)
        }
    }()

    ch := make(chan os.Signal, 1)
    signal.Notify(
        ch,
        os.Interrupt,
    )

    <-ch
    slog.Error("os.Interrupt - shutting down...")
    go func() {
        <-ch
        slog.Error("os.Kill - terminating...")
    }()

    gracefulCtx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancelShutdown()

    err = server.Shutdown(gracefulCtx)
    if err != nil {
        slog.Error(fmt.Sprintf("Shutdown error : %v", err))
        defer os.Exit(1)
        return
    }
    slog.Info("Gracefully stopped")

    defer os.Exit(0)
    return

}

func checkHealth() {
    port := os.Getenv("PORT")
    url := fmt.Sprintf("http://localhost:%s/health", port)
    _, err := http.Get(url)
    if err != nil {
        os.Exit(1)
    }
    os.Exit(0)
}
