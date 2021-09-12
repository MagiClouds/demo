package main

import (
	"fmt"
	"golang.org/x/time/rate"
	"time"
)

var limiter = rate.NewLimiter(100, 10)

func main() {
	for i := 0; i < 100; i++{
		if limiter.Allow() {
			fmt.Println("allow")
			continue
		}
		fmt.Println("not...")
		time.Sleep(100*time.Millisecond)
	}
}
