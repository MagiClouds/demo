package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"start/grpc_server/services"
)

func main() {
	grpcServer := grpc.NewServer()
	services.RegisterProdServiceServer(grpcServer, new(services.ProdService))
	services.RegisterGreeterServer(grpcServer, new(services.GreeterServerServer))

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
