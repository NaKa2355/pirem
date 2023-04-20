package web

import (
	"context"

	bdy "github.com/NaKa2355/pirem/internal/app/pirem/usecases/boundary"
	"google.golang.org/grpc"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func errWrapper(err error) error {
	if err == nil {
		return err
	}

	var grpcErrCode codes.Code
	switch _err := err.(type) {
	case *bdy.Error:
		switch _err.Code {
		case bdy.CodeTimeout:
			grpcErrCode = codes.Canceled
		case bdy.CodeNotExist:
			grpcErrCode = codes.NotFound
		case bdy.CodeAlreadyExists:
			grpcErrCode = codes.AlreadyExists
		case bdy.CodeBusy:
			grpcErrCode = codes.ResourceExhausted
		case bdy.CodeDevice:
			grpcErrCode = codes.FailedPrecondition
		case bdy.CodeInternal:
			grpcErrCode = codes.Internal
		case bdy.CodeInvaildInput:
			grpcErrCode = codes.InvalidArgument
		case bdy.CodeNotSupported:
			grpcErrCode = codes.Unimplemented
		default:
			grpcErrCode = codes.Unknown
		}
	}
	return status.Errorf(grpcErrCode, err.Error())
}

func ErrorUnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	res, err := handler(ctx, req)
	return res, errWrapper(err)
}
