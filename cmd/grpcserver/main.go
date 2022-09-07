package main

import (
	"log"
	"net"

	"github.com/sugamon/grpc-openapi-example/handler"
	pb "github.com/sugamon/grpc-openapi-example/pb/grpc_openapi_example/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	grpcServer := grpc.NewServer()
	handler := handler.NewHelloHandler()

	// gRPCリフレクションを有効にする
	reflection.Register(grpcServer)

	pb.RegisterHelloServiceServer(grpcServer, handler)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen. %v", err)
	}
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
