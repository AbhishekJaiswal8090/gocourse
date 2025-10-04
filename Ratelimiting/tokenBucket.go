// package main

// import (
// 	"time"
// 	"fmt"
// )
// // token bucket algom

// type RateLimiter struct {
// 	tokens     chan struct{}
// 	refillTime time.Duration
// }

// func NewRateLimiter(ratelimit int, refillTime time.Duration) *RateLimiter {
// 	rl := &RateLimiter{
// 		tokens:     make(chan struct{}, ratelimit),
// 		refillTime: refillTime,
// 	}
// 	for i := 0; i < ratelimit; i++ {
// 		rl.tokens <- struct{}{}
// 	}
// 	go rl.startRefill()
// 	return rl
// }

// func (rl *RateLimiter) startRefill() {
// 	ticker := time.NewTicker(rl.refillTime)
// 	defer ticker.Stop()
// 	for range ticker.C {
// 		select {
// 		case rl.tokens <- struct{}{}:
// 		default:
// 		}
// 	}
// }

// func (rl *RateLimiter) allow() bool {
// 	select {
// 	case <-rl.tokens:
// 		return true
// 	default:
// 		return false
// 	}
// }



func tokenBucketAlgom() {

	rateLimiter :=NewRateLimiter(5 ,time.Second)
	for range 10{
		if rateLimiter.allow() == true{
			fmt.Println("Request allowed")
		}else{
			fmt.Println("Request Denied")
		}
		time.Sleep(200*time.Millisecond)
	}


}
