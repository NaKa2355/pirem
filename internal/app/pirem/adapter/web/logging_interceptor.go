package adapter

import (
	"context"

	"google.golang.org/grpc"
)

func LoggingUnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	res, err := handler(ctx, req)
	return res, err
}
