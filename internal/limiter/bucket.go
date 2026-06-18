package limiter

import (
	"time"
)

type Bucket struct {
	Tokens     int
	Capacity   int
	RefillRate int
	LastRefill int64
}

func NewBucket(cap, rate int) *Bucket {
	return &Bucket{
		Tokens:     cap,
		Capacity:   cap,
		RefillRate: rate,
		LastRefill: time.Now().Unix(),
	}
}
