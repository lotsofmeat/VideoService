package main 

import (
	"log"
)

type ConnLimiter struct {
	concurrentConn int
	bucket chan int
}

//bucket token algorithm for rate limiter
//bucket has specific amount of token, 1 request occupy 1 token, 
//and release token back to bucket once the request finish
//If all of token has been occupied, frontend will reject user request
//Go uses shared channel instead of shared memory

func NewConnLimiter(cc int) *ConnLimiter {
	return &ConnLimiter {
		concurrentConn: cc,
		bucket: make(chan int, cc),
	}
}

func (cl *ConnLimiter) GetConn() bool {
	if len(cl.bucket) >= cl.concurrentConn {
		log.Printf("Reached the rate limitation.")
		return false
	}

	cl.bucket <- 1
	return true
}

func (cl *ConnLimiter) ReleaseConn() {
	c :=<- cl.bucket
	log.Printf("New connction coming: %d", c)
}
