package client_interceptor

import (
	"context"
	"filetransfer/internal/logger"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

// ClientLoggingInterceptor returns a gRPC unary client interceptor that logs the duration
// of each gRPC method call and any errors that occur.
func ClientLoggingInterceptor(logger logger.ClientLogger) grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		startTime := time.Now()

		// Invoke the gRPC method
		err := invoker(ctx, method, req, reply, cc, opts...)

		duration := time.Since(startTime)
		logger.Printf("gRPC method %s took %s\n", method, duration)

		// Log any errors that occurred during the gRPC method call
		if err != nil {
			statusErr, ok := status.FromError(err)
			if ok {
				loggedError := fmt.Errorf("gRPC method %s failed: %s", method, statusErr.Message())
				logger.Printf(loggedError.Error())
				return loggedError
			}
		}

		return err
	}
}
