package services

import (
	"context"
)

type GreeterServerServer struct {
	UnimplementedGreeterServer
}

func (g *GreeterServerServer) SayHello(c context.Context, h *HelloRequest) (*HelloReply, error) {
	return &HelloReply{Message: "hello" + h.Name}, nil

}

func (g *GreeterServerServer) SayHelloAgain(c context.Context, h *HelloRequest) (*HelloReply, error) {
	return &HelloReply{Message: "hello again" + h.Name}, nil

}
