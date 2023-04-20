package main

import (
	"context"
	"flag"
	"github.com/sirupsen/logrus"
	"github.com/timickb/link-shortener/internal/config"
	"github.com/timickb/link-shortener/internal/factory"
	"os"
	"strconv"
)

var (
	storageType string
)

func init() {
	flag.StringVar(&storageType, "storage", "postgres", "Must be postgres or memory")
}

func main() {
	flag.Parse()

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

	errChan := make(chan error)
	ctx := context.Background()

	httpServer, err := factory.InitializeHTTPServer(logger, cfg, storageType)
	if err != nil {
		return err
	}

	rpcServer, err := factory.InitializeRPCServer(ctx, logger, cfg, storageType)
	if err != nil {
		return err
	}

	go func(errChan chan<- error) {
		if err := httpServer.Run(); err != nil {
			errChan <- err
		}
	}(errChan)

	go func(errChan chan<- error) {
		if err := rpcServer.Run(); err != nil {
			errChan <- err
		}
	}(errChan)

	for {
		select {
		case err := <-errChan:
			return err
		case <-ctx.Done():
			return ctx.Err()
		}
	}
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
	if os.Getenv("HTTP_PORT") != "" {
		cfg.HTTPPort, _ = strconv.Atoi(os.Getenv("HTTP_PORT"))
	}
	if os.Getenv("RPC_PORT") != "" {
		cfg.RPCPort, _ = strconv.Atoi(os.Getenv("RPC_PORT"))
	}
}
