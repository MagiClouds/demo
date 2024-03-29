package services

import (
	"context"
	"fmt"
	"time"
)

type GreeterServerServer struct {
	UnimplementedGreeterServer
}

func (g *GreeterServerServer) SayHello(c context.Context, h *HelloRequest) (*HelloReply, error) {
	time.Sleep(2*time.Millisecond)
	deadline, ok := c.Deadline()
	if ok {
		fmt.Println(deadline)
	}

	return &HelloReply{Message: "hello" + h.Name}, nil
}

func (g *GreeterServerServer) SayHelloAgain(c context.Context, h *HelloRequest) (*HelloReply, error) {
	return &HelloReply{Message: "hello again" + h.Name}, nil
}
