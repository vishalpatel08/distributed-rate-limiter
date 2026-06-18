package server

import (
	"fmt"
	"net"

	"github.com/vishalpatel08/distributed-rate-limiter/internal/config"
	"github.com/vishalpatel08/distributed-rate-limiter/internal/server/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartGRPCServer(cfg *config.Config, handler *Handler) (*grpc.Server, net.Listener, error) {
	lis, err := net.Listen("tcp", ":"+cfg.GRPCPort)
	if err != nil {
		return nil, nil, err
	}

	grpcServer := grpc.NewServer()

	pb.RegisterRateLimiterServer(grpcServer, handler)
	reflection.Register(grpcServer)

	fmt.Println("gRPC server listening on port", cfg.GRPCPort)
	return grpcServer, lis, nil
}
