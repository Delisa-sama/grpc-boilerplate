package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	gw "grpc-boilerplate/api/v1"
)

func run(cfg *Config) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	grpcServerEndpoint := fmt.Sprintf("%s:%s", cfg.Endpoint.Host, cfg.Endpoint.Port)

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := gw.RegisterEchoHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts); err != nil {
		return err
	}

	if err := gw.RegisterMemoHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts); err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	proxyEntrypoint := fmt.Sprintf("%s:%s", cfg.Entrypoint.Host, cfg.Entrypoint.Port)
	return http.ListenAndServe(proxyEntrypoint, mux)
}

func main() {
	cfg, err := GetConfig()
	if err != nil {
		panic("failed to get proxy configuration: " + err.Error())
	}
	defer glog.Flush()

	if err := run(cfg); err != nil {
		glog.Fatal(err)
	}
}
