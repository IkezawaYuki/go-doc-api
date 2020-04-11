package main

import (
	"github.com/IkezawaYuki/go-dog-api/internal/infrastructure"
	"github.com/IkezawaYuki/go-dog-api/pkg/pb/dog"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	listenPort, err := net.Listen("tcp", ":9998")
	if err != nil {
		log.Fatalf("failed to listen port: %v", err)
	}

	zapLogger := infrastructure.CreateLogger()
	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_zap.UnaryServerInterceptor(zapLogger),
			infrastructure.AccessLogUnaryServerInterceptor(),
			grpc_auth.UnaryServerInterceptor(infrastructure.Authentication),
			infrastructure.AuthorizationUnaryServerInterceptor(),
		)),
	)
	dogService := &infrastructure.DogService{}
	dog.RegisterDogServer(server, dogService)

	reflection.Register(server)

	if err := server.Serve(listenPort); err != nil {
		log.Fatalf("faled to serve: %v", err)
	}
}
