// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package main

import (
	"context"
	"go_demo/xwire/baz"

	"github.com/google/wire"
)

func initializeBaz(ctx context.Context) (baz.Ss, error) {
	wire.Build(baz.SuperSet)
	return baz.Ss{}, nil
}