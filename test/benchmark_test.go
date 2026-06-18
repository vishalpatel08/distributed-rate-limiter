package test

import (
	"fmt"
	"testing"

	"github.com/vishalpatel08/distributed-rate-limiter/internal/limiter"
)

func BenchmarkAllow(b *testing.B) {
	mockRepo := &MockRepository{}
	service := limiter.NewService(mockRepo)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			service.Allow(fmt.Sprintf("client_%d", i%100))
			i++
		}
	})
}
