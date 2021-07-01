package main

import (
	"context"
	"fmt"
	"go_demo/xjwt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	dial, err := grpc.Dial("localhost:1100", opts...)
	if err != nil {
		panic(err)
	}

	defer dial.Close()

	client := xjwt.NewXjwtClient(dial)

	var clentOpts []grpc.CallOption

	ctx := context.Background()

	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJsb3JhLWFwcC1zZXJ2ZXIiLCJleHAiOjE2MzU3MjQ4MDAsImlzcyI6ImxvcmEtYXBwLXNlcnZlciIsIm5iZiI6MTYyNTEyMTc5OCwic3ViIjoidXNlciIsInVzZXJuYW1lIjoiYWRtaW4ifQ.BCkjxz6ZqRmEJ8Ya-642bWYSMz-a_Ge5NmD4-D_YMXY")
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJsb3JhLWFwcC1zZXJ2ZXIiLCJleHAiOjE2MzU3MjQ4MDAsImlzcyI6ImxvcmEtYXBwLXNlcnZlciIsIm5iZiI6MTYyNTEyMjkxMCwic3ViIjoidXNlcjIiLCJ1c2VybmFtZSI6ImFkbWluIn0.hm_sIa562fZ8W111MTW-kS3JNI_lzkxmH261WhTz7MI")
	getXjwt, err := client.GetXjwt(ctx, &xjwt.GetXjwtRequest{}, clentOpts...)
	if err != nil {
		panic(err)
	}

	fmt.Println(getXjwt)
}
