package server_interceptor

import (
	"context"
	"filetransfer/internal/logger"
	"fmt"
	"google.golang.org/grpc/status"
	"time"

	"google.golang.org/grpc"
)

// LoggingInterceptor returns a unary server interceptor that logs information about gRPC method calls.
func LoggingInterceptor(logger logger.ServerLogger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		startTime := time.Now()

		// Call the handler to process the request
		resp, err := handler(ctx, req)

		duration := time.Since(startTime)
		logger.Printf("gRPC method %s took %s\n", info.FullMethod, duration)

		// Log and return an error if the handler encounters an error
		if err != nil {
			loggedError := fmt.Errorf("gRPC method %s failed: %v", info.FullMethod, err)
			logger.Printf(loggedError.Error())
			return nil, status.Error(status.Code(err), loggedError.Error())
		}

		return resp, err
	}
}
