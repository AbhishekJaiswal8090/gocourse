package main

import (
	"time"
)

type RateLimiter struct {
	tokens     chan struct{}
	refillTime time.Duration
}

func NewRateLimiter(limit int, refillTime time.Duration) *RateLimiter {
	rl := RateLimiter{
		tokens:     make(chan struct{}, limit),
		refillTime: refillTime,
	}

	for i := 0; i < limit; i++ {
		rl.tokens <- struct{}{}
	}
	go rl.StartRefill(refillTime)
	return &rl
}

func (rl *RateLimiter) StartRefill(refillTime time.Duration) {
	ticker := time.NewTicker(refillTime)
	defer ticker.Stop()
	for range ticker.C {
		select {
		case rl.tokens <- struct{}{}:
		default:
		}

	}
}

func (rl *RateLimiter) allow()bool{
	select{
	case <- rl.tokens:
		return true
	default:
		return false
	}
}




func main() {}
