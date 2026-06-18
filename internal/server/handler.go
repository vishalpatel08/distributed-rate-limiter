package server

import (
	"context"

	"github.com/vishalpatel08/distributed-rate-limiter/internal/limiter"
	"github.com/vishalpatel08/distributed-rate-limiter/internal/server/pb"
)

type Handler struct {
	pb.UnimplementedRateLimiterServer
	service *limiter.Service
}

func NewHandler(service *limiter.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Allow(ctx context.Context, req *pb.AllowRequest) (*pb.AllowResponse, error) {

	allowed, remaining, err := h.service.Allow(req.ClientId)

	if err != nil {
		return nil, err
	}

	return &pb.AllowResponse{
		Allowed:         allowed,
		RemainingTokens: int32(remaining),
	}, nil

}
