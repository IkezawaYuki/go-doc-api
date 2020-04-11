package main

import (
	health "github.com/IkezawaYuki/go-dog-api/google.golang.org/grpc/health/grpc_health_v1"
	"github.com/IkezawaYuki/go-dog-api/internal/infrastructure"
	"github.com/IkezawaYuki/go-dog-api/pkg/pb/dog"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
)

func main() {
	log.Println("grpc server is start...")
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

	healthService := &infrastructure.SkipAuthHealthServer{}
	health.RegisterHealthServer(server, healthService)

	reflection.Register(server)

	go func() {
		if err := server.Serve(listenPort); err != nil {
			log.Fatalf("faled to serve: %v", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping grpc server...")
	server.Stop()
}
