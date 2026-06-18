package test

import (
	"errors"
	"testing"

	"github.com/vishalpatel08/distributed-rate-limiter/internal/limiter"
)

type MockRepository struct{}

func (m *MockRepository) ConsumeToken(clientID string) (bool, int, error) {
	return true, 99, nil
}

func TestAllowSuccess(t *testing.T) {

	mockRepo := &MockRepository{}

	service := limiter.NewService(mockRepo)

	allowed, remaining, err := service.Allow("user123")

	if err != nil {
		t.Fatal(err)
	}

	if !allowed {
		t.Fatal("expected request to be allowed")
	}

	if remaining != 99 {
		t.Fatal("Unxpected remaining tokens")
	}
}

func TestEmptyClientID(t *testing.T) {

	mockRepo := &MockRepository{}

	service := limiter.NewService(mockRepo)

	_, _, err := service.Allow("")

	if err == nil {
		t.Fatal(errors.New("expected error"))
	}
}
