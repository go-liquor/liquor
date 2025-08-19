package grpcserver

import (
	"context"
	"fmt"
	"net"

	"github.com/go-liquor/liquor/v3/config"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// RegisterGRPCServer creates and configures a new gRPC server with the provided implementation.
// It uses the uber-fx dependency injection framework to wire up the server components.
//
// Parameters:
//   - implementation: The grpc server implementation struct (e.g.: &adapters.MyGrpcHandler{})
//   - instance: The function to return the grpc implementation (e.g.: adapters.NewGrpcHandler)
//   - register: Function to register the implementation with the gRPC server
//
// Returns:
//   - fx.Option: A configured module containing the gRPC server setup
//
// The server will automatically:
//   - Start on the configured port from config.GrpcPort
//   - Handle graceful shutdown
//   - Log server lifecycle events
func RegisterGRPCServer[T any, A any](implementation T, instance A, register func(imp T, registrar *grpc.Server)) fx.Option {
	return fx.Module("liquor-grpc-server", fx.Provide(instance), fx.Provide(func() *grpc.Server {
		return grpc.NewServer()
	}), fx.Invoke(register),
		fx.Invoke(
			func(cfg *config.Config, logger *zap.Logger, svc *grpc.Server, lc fx.Lifecycle) {
				lc.Append(fx.Hook{
					OnStart: func(ctx context.Context) error {
						logger.Info("starting grpc server", zap.Int64("port", cfg.GetInt64(config.GrpcPort)))
						lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GetInt64(config.GrpcPort)))
						if err != nil {
							logger.Fatal("failed to start tcp grpc", zap.Error(err))
							return err
						}
						go svc.Serve(lis)
						return nil
					},
					OnStop: func(ctx context.Context) error {
						logger.Info("stopping grpc server")
						svc.GracefulStop()
						return nil
					},
				})

			},
		))
}
