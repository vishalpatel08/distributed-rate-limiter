package server

import (
	"fmt"
	"net"

	"github.com/vishalpatel08/distributed-rate-limiter/internal/config"
	"github.com/vishalpatel08/distributed-rate-limiter/internal/server/pb"
	"google.golang.org/grpc"
)

func StartGRPCServer(cfg *config.Config, handler *Handler) error {
	lis, err := net.Listen("tcp", ":"+cfg.GRPCPort)

	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	pb.RegisterRateLimiterServer(grpcServer, handler)

	fmt.Println("gRPC server listening on port", cfg.GRPCPort)
	return grpcServer.Serve(lis)
}
