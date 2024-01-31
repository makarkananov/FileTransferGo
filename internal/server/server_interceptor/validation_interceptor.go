package server_interceptor

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ValidationInterceptor returns a unary server interceptor that performs validation on incoming gRPC requests.
func ValidationInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// Check if the request type implements the Validate method
		if v, ok := req.(interface{ Validate() error }); ok {
			// Validate the request and return an error if validation fails
			if err := v.Validate(); err != nil {
				return nil, status.Error(codes.InvalidArgument, err.Error())
			}
		}

		// Call the handler to process the request
		resp, err := handler(ctx, req)

		return resp, err
	}
}
