package main

import (
	"fmt"
	"github.com/go-kratos/aegis/ratelimit/bbr"
	"github.com/go-kratos/aegis/ratelimit"
)

var limiter = bbr.NewLimiter()

func main()  {
	done, err := limiter.Allow()
		if err != nil {
		fmt.Println(err.Error())
		return
	}
	done(ratelimit.DoneInfo{})
}
