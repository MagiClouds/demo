package baz

import (
	"context"
	"errors"
	"github.com/google/wire"
	"go_demo/xwire/foo"
)

type Bar struct {
	X int
}

// ProvideBar returns a Bar: a negative Foo.
func ProvideBar(foo foo.Foo) Bar {
	return Bar{X: -foo.X}
}

type Baz struct {
	X int
}

// ProvideBaz returns a value if Bar is not zero.
func ProvideBaz(ctx context.Context, bar Bar) (Baz, error) {
	if bar.X == 0 {
		return Baz{}, errors.New("cannot provide baz when bar is zero")
	}
	return Baz{X: bar.X}, nil
}

type Ss struct {
	A string
}

// ProvideTi returns nil if Bar is not zero.
func ProvideTi(ctx context.Context, biz Baz) (Ss, error) {
	if biz.X == 0 {
		return Ss{}, errors.New("cannot provide baz when bar is zero")
	}
	return Ss{"aaa"}, nil
}

var SuperSet = wire.NewSet(foo.ProvideFoo, ProvideBar, ProvideBaz, ProvideTi)