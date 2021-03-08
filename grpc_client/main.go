package main

import (
	"context"
	"fmt"
	"go_demo/grpc_client/services"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	ctx := context.Background()

	//prodClient := services.NewProdServiceClient(conn)
	//prodRes, err := prodClient.GetProdStock(ctx, &services.ProdRequest{ProdId:10})
	//if err != nil {
	//	log.Fatal(err)
	//}

	//res, err := prodClient.GetProdStocks(ctx, &services.QuerySize{Size:3})
	//if err != nil {
	//	log.Fatal(err)
	//}

	//for _, i := range res.ProdInfo {
	//	fmt.Println(i)
	//}

	client := services.NewGreeterClient(conn)

	reply, err := client.SayHello(ctx, &services.HelloRequest{Name: "Grpc"})
	if err != nil {
		panic("hahha say hello failed")
	}

	fmt.Println(reply.Message)

	replyAgain, err := client.SayHello(ctx, &services.HelloRequest{Name: "Grpc"})
	if err != nil {
		panic("hahha say hello failed")
	}

	fmt.Println(replyAgain.Message)




}
