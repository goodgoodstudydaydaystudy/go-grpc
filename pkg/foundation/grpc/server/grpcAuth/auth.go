/**
 * Author: Orange
 * Date: 20-02-06
 */
package grpcAuth

import (
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	errUnauthenticated = status.Error(codes.Unauthenticated, "Not login.")
	errTokenExpired    = status.Error(codes.PermissionDenied, "Token is expired.")
	errInternal        = status.Error(codes.Internal, "Internal error")
)

func UnaryServerInterceptor(authFunc grpc_auth.AuthFunc) grpc.UnaryServerInterceptor {
	return grpc_auth.UnaryServerInterceptor(authFunc)
}

type authFuncBuilder struct {
	fullMethodNameException map[string]bool
}

func defaultBuilder() *authFuncBuilder {
	return &authFuncBuilder{
		fullMethodNameException: make(map[string]bool),
	}
}

func NewAuthFuncBuilder() *authFuncBuilder {
	return defaultBuilder()
}

func (builder *authFuncBuilder) WithFullMethodException(
	fullMethodNames ...string) *authFuncBuilder {
	for _, method := range fullMethodNames {
		builder.fullMethodNameException[method] = true
	}
	return builder
}
