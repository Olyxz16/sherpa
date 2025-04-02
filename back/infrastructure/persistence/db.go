package persistence

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/Olyxz16/sherpa/config"
)

type Service interface {
	Health() bool
}

type service struct {
	pool *pgxpool.Pool
}

var (
	instance *service
)

func New(cfg config.DatabaseConfig) Service {
	// Reuse Connection
	if instance != nil {
		return instance
	}
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", cfg.User, cfg.Pass, cfg.Host, cfg.Port, cfg.DBName)
	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
        log.Fatal(err)
    }
    instance = &service{
		pool: pool,
	}
	return instance
}

func Instance() (Service, error) {
	if instance == nil {
		return nil, fmt.Errorf("Service has not been initialized.") 
	}
	return instance, nil
}

func Conn() *pgxpool.Pool {
	if instance == nil {
		panic("Service has not been initialized.")	
	}
	return instance.pool
}

func (s *service) Health() bool {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := s.pool.Ping(ctx)
	if err != nil {
		log.Fatal(fmt.Sprintf("db down: %v", err))
	}

    return err == nil
}
