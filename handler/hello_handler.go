package handler

import (
	"context"

	pb "github.com/sugamon/grpc-openapi-example/pb/grpc_openapi_example/v1"
)

type HelloHandler struct {
	pb.UnimplementedHelloServiceServer
}

func NewHelloHandler() pb.HelloServiceServer {
	return &HelloHandler{}
}

func (h *HelloHandler) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Greeting: "Hello World!",
	}, nil
}
