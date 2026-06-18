package main

import (
	"log"

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

	log.Fatal(server.StartGRPCServer(cfg, handler))
}
