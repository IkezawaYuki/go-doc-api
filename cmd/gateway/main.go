package main

import (
	"context"
	"flag"
	"github.com/IkezawaYuki/golang-grpc-server/google.golang.org/grpc/health/grpc_health_v1"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"net/http"
	"os"
)

func run() error {
	grpcServerEndpoint := "localhost:9998"
	if os.Getenv("DEPLOY_STAGE") == "local" || os.Getenv("DEPLOY_STAGE") == "" {
		grpcServerEndpoint = "localhost:9998"
	}
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := grpc_health_v1.RegisterHealthHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		return err
	}
	return http.ListenAndServe(":8081", mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()
	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
