package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"start/grpc_start/order"
	"time"
)

func createOrder(client order.OrderClient, params *order.OrderParams) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	orderResult, err := client.Add(ctx, params)
	if err != nil {
		log.Fatalf("%v.GetFeatures(_) = _, %v:", client, err)
	}
	log.Printf("the result is %v", orderResult)
}

func main()  {
	conn, err := grpc.Dial("127.0.0.1:8081",  grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := order.NewOrderClient(conn)
	orderParams := &order.OrderParams{
		BuyerID: 10318003,
	}
	createOrder(client, orderParams)
}
