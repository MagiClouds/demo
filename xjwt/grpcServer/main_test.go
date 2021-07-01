package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/metadata"
	"testing"
)

func Test_getTokenFromContext(t *testing.T) {
	ctx := context.Background()
	// create a new context with some metadata
	ctx = metadata.AppendToOutgoingContext(ctx, "k1", "v1", "k1", "v2", "k2", "v3")

	// later, add some more metadata to the context (e.g. in an interceptor)
	ctx = metadata.AppendToOutgoingContext(ctx, "k3", "v4")
	ctx = metadata.AppendToOutgoingContext(ctx, "k3", "v5")

	fmt.Println(ctx)

	// make unary RPC
	//response, err := client.SomeRPC(ctx, someRequest)

	// or make streaming RPC
	//stream, err := client.SomeStreamingRPC(ctx)// create a new context with some metadata
	//ctx := metadata.AppendToOutgoingContext(ctx, "k1", "v1", "k1", "v2", "k2", "v3")

	// later, add some more metadata to the context (e.g. in an interceptor)
	//ctx := metadata.AppendToOutgoingContext(ctx, "k3", "v4")

	// make unary RPC
	//response, err := client.SomeRPC(ctx, someRequest)

	// or make streaming RPC
	//stream, err := client.SomeStreamingRPC(ctx)
}