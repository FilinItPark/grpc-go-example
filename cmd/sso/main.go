package main

import (
	"github.com/1ommyS/sso-example-go/internal/app"
	"github.com/1ommyS/sso-example-go/internal/config"
	"github.com/1ommyS/sso-example-go/internal/lib/logger/handlers/slogpretty"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	logger := setupLogger(cfg.Env)

	logger.Info("Starting application", slog.Any("cfg", cfg))

	application := app.New(logger,
		cfg.GRPC.Port, cfg.StoragePath, cfg.TokenTTL,
	)

	application.GRPCServer.MustRun()

	// TODO: запустить gRPC-сервер приложения
}

func setupLogger(env string) (logger *slog.Logger) {
	switch env {
	case envLocal:
		logger = setupPrettySlog()
	case envDev:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return logger
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
