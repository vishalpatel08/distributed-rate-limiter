package limiter

import (
	"errors"
)

type Repository interface {
	ConsumeToken(clientID string) (bool, int, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Allow(clientID string) (bool, int, error) {
	if clientID == "" {
		return false, 0, errors.New("empty client Id")
	}

	return s.repo.ConsumeToken(clientID)
}
