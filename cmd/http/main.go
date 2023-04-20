package main

import (
	"github.com/sirupsen/logrus"
	"github.com/timickb/link-shortener/internal/config"
	"github.com/timickb/link-shortener/internal/factory"
	"os"
	"strconv"
)

func main() {
	logger := logrus.New()
	formatter := &logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	}
	logger.SetFormatter(formatter)

	if err := mainNoExit(logger); err != nil {
		logger.Fatal(err)
	}
}

func mainNoExit(logger *logrus.Logger) error {
	cfg := config.NewDefault()
	fillConfigFromEnv(cfg)

	server, err := factory.InitializeHTTPServerPostgres(logger, cfg)
	if err != nil {
		return err
	}

	if err := server.Run(); err != nil {
		return err
	}

	return nil
}

func fillConfigFromEnv(cfg *config.AppConfig) {
	if os.Getenv("DB_HOST") != "" {
		cfg.Postgres.Host = os.Getenv("DB_HOST")
	}
	if os.Getenv("DB_USER") != "" {
		cfg.Postgres.User = os.Getenv("DB_USER")
	}
	if os.Getenv("DB_NAME") != "" {
		cfg.Postgres.Name = os.Getenv("DB_NAME")
	}
	if os.Getenv("DB_PASSWORD") != "" {
		cfg.Postgres.Password = os.Getenv("DB_PASSWORD")
	}
	if os.Getenv("DB_SSL_MODE") != "" {
		cfg.Postgres.SSLMode = os.Getenv("DB_SSL_MODE")
	}
	if os.Getenv("DB_PORT") != "" {
		cfg.Postgres.Port, _ = strconv.Atoi(os.Getenv("DB_PORT"))
	}
	if os.Getenv("APP_PORT") != "" {
		cfg.AppPort, _ = strconv.Atoi(os.Getenv("APP_PORT"))
	}
}
