package model

import (
    "context"
	"database/sql"
	"fmt"
	"log"
	"time"

    _ "github.com/jackc/pgx/v5/stdlib"

	"github.com/Olyxz16/sherpa/config"
)

type Service interface {
	Health() bool
}

type service struct {
	db *sql.DB
}

var (
	instance *service
)

func New() Service {
	// Reuse Connection
	if instance != nil {
		return instance
	}
    cfg, err := config.NewDatabaseConfig()
    if err != nil {
        panic("Error loading database config")
    }
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", cfg.User, cfg.Pass, cfg.Host, cfg.Port, cfg.DBName)
	db, err := sql.Open("pgx", connStr)
	if err != nil {
        log.Fatal(err)
    }
    instance = &service{
		db: db,
	}
	return instance
}

func (s *service) Health() bool {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := s.db.PingContext(ctx)
	if err != nil {
		log.Fatal(fmt.Sprintf("db down: %v", err))
	}

    return err == nil
}
