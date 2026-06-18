package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/vishalpatel08/distributed-rate-limiter/internal/config"
	"github.com/vishalpatel08/distributed-rate-limiter/internal/limiter"
	"github.com/vishalpatel08/distributed-rate-limiter/internal/server"
	"github.com/vishalpatel08/distributed-rate-limiter/internal/storage"
)

func main() {
	cfg := config.Load()

	redisClient := storage.NewRadisClient(cfg)

	repo := storage.NewRepository(cfg, redisClient)

	service := limiter.NewService(repo)

	handler := server.NewHandler(service)

	grpcServer, lis, err := server.StartGRPCServer(cfg, handler)

	if err != nil {
		log.Fatal(err)
	}

	go func() {
		log.Printf("gRPC server listening on %s", cfg.GRPCPort)

		if err := grpcServer.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(
		quit,
		os.Interrupt,
		syscall.SIGTERM,
	)

	<-quit

	log.Println("Shutting down server...")

	grpcServer.GracefulStop()

	redisClient.Close()

	log.Println("Server stopped")
}
