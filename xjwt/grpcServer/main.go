package main

import (
	"context"
	"go_demo/xjwt"
	"google.golang.org/grpc"
	"log"
	"net"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"google.golang.org/grpc/metadata"
)

// Claims defines the struct containing the token claims.
type Claims struct {
	jwt.StandardClaims

	// Username defines the identity of the user.
	Username string `json:"username"`
}

type server struct {
	xjwt.UnimplementedXjwtServer
}

func (s *server) CreateXjwt(context.Context, *xjwt.CreateXjwtRequest) (*xjwt.CreateXjwtReply, error) {
	panic("implement me")
}

func (s *server) UpdateXjwt(context.Context, *xjwt.UpdateXjwtRequest) (*xjwt.UpdateXjwtReply, error) {
	panic("implement me")
}

func (s *server) DeleteXjwt(context.Context, *xjwt.DeleteXjwtRequest) (*xjwt.DeleteXjwtReply, error) {
	panic("implement me")
}

func (s *server) GetXjwt(ctx context.Context, r *xjwt.GetXjwtRequest) (*xjwt.GetXjwtReply, error) {
	tokenStr, err := getTokenFromContext(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get token from context error")
	}

	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if token.Header["alg"] != "HS256" {
			return nil, nil
		}
		return []byte("verysecret"), nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "jwt parse error")
	}

	if !token.Valid {
		return nil, nil
	}

	log.Printf("token: %v", token.Claims)
	return &xjwt.GetXjwtReply{}, nil
}

func (s *server) ListXjwt(context.Context, *xjwt.ListXjwtRequest) (*xjwt.ListXjwtReply, error) {
	panic("implement me")
}

// Step1. 从 context 的 metadata 中，取出 token

func getTokenFromContext(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", nil
	}
	// md 的类型是 type MD map[string][]string
	token, ok := md["authorization"]
	if !ok || len(token) == 0 {
		return "", nil
	}
	// 因此，token 是一个字符串数组，我们只用了 token[0]
	return token[0], nil
}

// Step2. 从 token 解析出 jwt 的 claim

// SayHello implements helloworld.GreeterServer

func main() {
	lis, err := net.Listen("tcp", "localhost:1100")
	if err != nil {
		panic(err)
	}

	defer lis.Close()

	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)
	xjwt.RegisterXjwtServer(s, &server{})
	s.Serve(lis)
}
