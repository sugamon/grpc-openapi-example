package main

import (
	"context"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	gw "github.com/sugamon/grpc-openapi-example/pb/grpc_openapi_example/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// docker-composeで立ち上げたswagger-uiとgrpc-serverを指定
const grpcServerEndpoint = "grpc-server:50051"
const swaggerEndpoint = "http://swagger-ui:8080"

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	grpcGateway := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := gw.RegisterHelloServiceHandlerFromEndpoint(ctx, grpcGateway, grpcServerEndpoint, opts)
	if err != nil {
		log.Fatalf("faild to register. %v", err)
	}

	swaggerEndpoint, err := url.Parse(swaggerEndpoint)
	if err != nil {
		log.Fatalf("faild to parse. %v", err)
	}
	swaggerProxy := httputil.NewSingleHostReverseProxy(swaggerEndpoint)

	mux := http.NewServeMux()
	mux.Handle("/swagger/", swaggerProxy)
	mux.Handle("/", grpcGateway)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("error occurred. %v", err)
	}
}
