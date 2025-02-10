package adapter

import (
	"context"

	"github.com/NaKa2355/pirem/internal/app/pirem/usecases"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ErrorUnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	res, err := handler(ctx, req)
	var grpcErrCode codes.Code
	if err == nil {
		return res, err
	}
	switch err := err.(type) {
	case *usecases.Error:
		switch err.Code {
		case usecases.CodeAlreadyExists:
			grpcErrCode = codes.AlreadyExists
		case usecases.CodeBusy:
			grpcErrCode = codes.ResourceExhausted
		case usecases.CodeDataBase:
			grpcErrCode = codes.Internal
		case usecases.CodeInvaildInput:
			grpcErrCode = codes.InvalidArgument
		case usecases.CodeNotFound:
			grpcErrCode = codes.NotFound
		case usecases.CodeNotSupported:
			grpcErrCode = codes.Unimplemented
		case usecases.CodeTimeout:
			grpcErrCode = codes.DeadlineExceeded
		default:
			grpcErrCode = codes.Unimplemented
		}
	}
	return res, status.Error(grpcErrCode, err.Error())
}
