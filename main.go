package main

import (
	"context"
	"net"

	"github.com/Striker87/Basic-API-with-gRPC-and-Protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//рабочий пример

//use:
//http://localhost:8080/add/152/108
//http://localhost:8080/mult/152/108

type server struct{}

func main() {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	proto.RegisterAddServiceServer(srv, &server{})
	reflection.Register(srv) //serialize/unserialize data

	if err = srv.Serve(listener); err != nil {
		panic(err)
	}
}

func (s *server) Add(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a + b

	return &proto.Response{Result: result}, nil
}

func (s *server) Multiply(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a * b

	return &proto.Response{Result: result}, nil
}
