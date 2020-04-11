package infrastructure

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var dogMethodList = map[string][]string{
	"/Dog/FindSmartDog": {"FindSmartDog"},
}

type User struct {
	permissions []string
}

func findUser(id string) *User {
	switch id {
	case "1":
		return &User{permissions: []string{"FindSmartDog"}}
	case "2":
		return &User{permissions: []string{"CreateDog"}}
	}
	return &User{}
}

func AuthorizationUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		if info.FullMethod == "/grpc.health.v1.Health/Check" {
			return handler(ctx, req)
		}
		token, ok := GetAuthToken(ctx)
		if ok {
			user := findUser(token.Subject)
			if canAccessToMethod(info.FullMethod, user) {
				return handler(ctx, req)
			}
		}
		return nil, status.Errorf(
			codes.PermissionDenied,
			"could not access to specified method",
		)
	}
}

func canAccessToMethod(method string, user *User) bool {
	r, ok := dogMethodList[method]
	if !ok {
		return false
	}
	permissions := map[string]bool{}
	for _, p := range user.permissions {
		permissions[p] = true
	}
	for _, p := range r {
		if !permissions[p] {
			return false
		}
	}
	return true
}
