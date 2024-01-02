package adapter

import (
	"context"

	"github.com/NaKa2355/pirem/pkg/logger"
	"google.golang.org/grpc"
)

func LoggingUnaryServerInterceptor(logger logger.Logger) func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		res, err := handler(ctx, req)
		if err != nil {
			logger.Debug("handled request",
				"request", req,
				"response", res,
				"error", err)
			return res, err
		}
		logger.Debug("handled request",
			"request", req,
			"response", res)
		return res, err
	}
}
