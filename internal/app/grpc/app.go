package grpc

import (
	"fmt"
	authgrpc "github.com/1ommyS/sso-example-go/internal/grpc/auth"
	"google.golang.org/grpc"
	"log/slog"
	"net"
)

type App struct {
	logger     *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

// New create new gRPC server app.
func New(
	log *slog.Logger,
	port int,
) *App {
	gRPCServer := grpc.NewServer()

	authgrpc.Register(gRPCServer)

	return &App{
		logger:     log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}

func (app *App) MustRun() {
	if err := app.Run(); err != nil {
		panic(err)
	}
}

func (app *App) Run() error {
	const op = "grpcapp.Run"

	log := app.logger.With(
		slog.String("op", op),
		slog.Int("port", app.port),
	)

	l, err := net.Listen("tcp", fmt.Sprint(":%d", app.port))

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info(" gRPC server is running", slog.String("addr", l.Addr().String()))

	if err := app.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (app *App) Stop() {
	const op = "grpcapp.Stop"

	app.logger.With(slog.String("op", op))
	app.logger.Info("Stopping gRPC server", slog.Int("port", app.port))

	app.gRPCServer.GracefulStop()
}
