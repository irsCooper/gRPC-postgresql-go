package main

import (
	"log/slog"
	"os"

	"github.com/irsCooper/gRPC-postgresql-go/sso/internal/app"
	"github.com/irsCooper/gRPC-postgresql-go/sso/internal/config"
)

const (
	envLocal = "local"
	envDev = "dev"
	envProd = "prod"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info(
		"starting applocation", 
		slog.String("env", cfg.Env),
		slog.Any("cfg", cfg),
		slog.Int("port", cfg.GRPC.Port),
	)

	application := app.New(log, cfg.GRPC.Port, cfg.StoragePath, cfg.TokenTTL)

	log.Info("tututu")

	application.GRPCSrv.MustRun()

	log.Info("tututu")
}


func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
		case envLocal:
			log = slog.New(
				slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
			)
		case envDev:
			log = slog.New(
				slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
			)
		case envProd:
			log = slog.New(
				slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
			)
	}

	return log
}
