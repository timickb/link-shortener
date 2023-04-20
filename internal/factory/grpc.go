package factory

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/timickb/link-shortener/internal/config"
	v1 "github.com/timickb/link-shortener/internal/delivery/grpc/v1"
	"github.com/timickb/link-shortener/internal/interfaces"
	"github.com/timickb/link-shortener/internal/repository/memory"
	"github.com/timickb/link-shortener/internal/repository/postgres"
	"github.com/timickb/link-shortener/internal/usecase/shortener"
)

func InitializeRPCServer(ctx context.Context, log interfaces.Logger, cfg *config.AppConfig, storage string) (*v1.Server, error) {
	if storage == "postgres" {
		connStr := fmt.Sprintf(
			"host=%s user=%s dbname=%s sslmode=%s port=%d password=%s",
			cfg.Postgres.Host,
			cfg.Postgres.User,
			cfg.Postgres.Name,
			cfg.Postgres.SSLMode,
			cfg.Postgres.Port,
			cfg.Postgres.Password)

		db, err := sql.Open("postgres", connStr)
		if err != nil {
			return nil, err
		}

		if err := db.Ping(); err != nil {
			return nil, err
		}

		repo := postgres.New(db)
		service := shortener.New(log, repo)
		server := v1.New(ctx, log, cfg, service)

		return server, nil
	}

	if storage == "memory" {
		repo := memory.New()
		service := shortener.New(log, repo)
		server := v1.New(ctx, log, cfg, service)

		return server, nil
	}

	return nil, errors.New("err invalid storage type")
}
