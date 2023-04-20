package factory

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/timickb/link-shortener/internal/config"
	v1 "github.com/timickb/link-shortener/internal/delivery/http/v1"
	"github.com/timickb/link-shortener/internal/interfaces"
	"github.com/timickb/link-shortener/internal/repository/memory"
	"github.com/timickb/link-shortener/internal/repository/postgres"
	"github.com/timickb/link-shortener/internal/usecase/shortener"
)

func InitializeHTTPServerPostgres(log interfaces.Logger, cfg *config.AppConfig) (*v1.Server, error) {
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
	server := v1.New(log, cfg, service)

	return server, nil
}

func InitializeHTTPServerMemStore(log interfaces.Logger, cfg *config.AppConfig) (*v1.Server, error) {
	repo := memory.New()
	service := shortener.New(log, repo)
	server := v1.New(log, cfg, service)

	return server, nil
}
