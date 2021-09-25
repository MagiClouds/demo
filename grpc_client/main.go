package main

import (
	"context"
	"fmt"
	"go_demo/grpc_client/services"
	"google.golang.org/grpc"
	"log"
	"time"
)


func clientInterceptor(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	// Logic before invoking the invoker
	start := time.Now()
	// Calls the invoker to execute RPC
	err := invoker(ctx, method, req, reply, cc, opts...)
	// Logic after invoking the invoker
	log.Printf("Invoked RPC method=%s; Duration=%s; Error=%v", method, time.Since(start), err)
	return err
}

var key string="name"

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure(), grpc.WithUnaryInterceptor(clientInterceptor))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

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

	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Millisecond)
	defer cancelFunc()

	ctx = context.WithValue(ctx, "test", "test-value")

	reply, err := client.SayHello(ctx, &services.HelloRequest{Name: "Grpc"})
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(reply.Message)

	replyAgain, err := client.SayHello(ctx, &services.HelloRequest{Name: "timeout111"})
	if err != nil {
		panic("hahha say hello failed")
	}

	fmt.Println(replyAgain.Message)


}
