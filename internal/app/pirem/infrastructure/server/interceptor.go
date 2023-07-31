package server

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

	var code codes.Code
	switch _err := err.(type) {
	case *bdy.Error:
		switch _err.Code {
		case bdy.CodeTimeout:
			code = codes.DeadlineExceeded
		case bdy.CodeNotExist:
			code = codes.NotFound
		case bdy.CodeAlreadyExists:
			code = codes.AlreadyExists
		case bdy.CodeBusy:
			code = codes.ResourceExhausted
		case bdy.CodeDevice:
			code = codes.FailedPrecondition
		case bdy.CodeInternal:
			code = codes.Internal
		case bdy.CodeInvaildInput:
			code = codes.InvalidArgument
		case bdy.CodeNotSupported:
			code = codes.Unimplemented
		default:
			code = codes.Unknown
		}
	}
	return status.Errorf(code, err.Error())
}

func (s *Server) UnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	res, err := handler(ctx, req)
	s.logger.Debug("grpc called", "method", info.FullMethod, "req", req, "res", res)
	return res, errWrapper(err)
}
