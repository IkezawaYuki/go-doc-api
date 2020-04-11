package infrastructure

import (
	"context"
	health "github.com/IkezawaYuki/go-dog-api/google.golang.org/grpc/health/grpc_health_v1"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SkipAuthHealthServer struct {
	health.HealthServer
}

var _ grpc_auth.ServiceAuthFuncOverride = (*SkipAuthHealthServer)(nil)

func (*SkipAuthHealthServer) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	return ctx, nil
}

func (s *SkipAuthHealthServer) Check(ctx context.Context, req *health.HealthCheckRequest) (*health.HealthCheckResponse, error) {
	return &health.HealthCheckResponse{
		Status: health.HealthCheckResponse_SERVING,
	}, nil
}
func (s *SkipAuthHealthServer) Watch(req *health.HealthCheckRequest, stream health.Health_WatchServer) error {
	return status.Error(codes.Unimplemented, "service watch is not implemented current version.")
}
