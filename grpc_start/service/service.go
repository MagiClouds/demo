package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"start/grpc_start/order"
	"start/grpc_start/util"
)

type Order struct {
}

func (o *Order) Add (ctx context.Context, in *order.OrderParams) (*order.OrderResult, error) {
	return &order.OrderResult{
		OrderID: util.CreateOrder(1),
	}, nil
}

func main()  {
	lis, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	order.RegisterOrderServer(grpcServer, &Order{})
	grpcServer.Serve(lis)
}
