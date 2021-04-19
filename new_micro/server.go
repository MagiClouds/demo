package main

import (
	"context"
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	"go_demo/new_micro/proto"
)

type HelloWorld struct{}

func (h *HelloWorld) SayHello(ctx context.Context, in *hello_world.SayRequest, out *hello_world.SayResponse) error {

	out.Answer = "hello world, i am answer"

	return nil
}

func main() {

	srv := service.New(
		service.Name("hello"),
	)

	srv.Handle(new(HelloWorld))

	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
