package app

import (
	grpcapp "github.com/1ommyS/sso-example-go/internal/app/grpc"
	"log/slog"
	"time"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(
	log *slog.Logger,
	gRPCPort int,
	storagePath string,
	tokenTTL time.Duration,
) *App {
	// TOOD: init storage

	// TODO: init auth service

	grpcApp := grpcapp.New(
		log,
		gRPCPort,
	)

	return &App{
		GRPCServer: grpcApp,
	}
}
