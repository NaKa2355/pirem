package server

import (
	"context"
	"errors"
	pirem_err "github.com/NaKa2355/pirem_pkg/error"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
)

func ErrorWrapping(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	res, err := handler(ctx, req)
	if err == nil {
		return res, err
	}

	if errors.Is(err, pirem_err.ErrDeviceNotFound) {
		err = status.Error(codes.NotFound, err.Error())
		return res, err
	}
	if errors.Is(err, pirem_err.ErrInvaildArgument) {
		err = status.Error(codes.InvalidArgument, err.Error())
		return res, err
	}

	if errors.Is(err, pirem_err.ErrDeviceInternal) {
		err = status.Error(codes.Internal, err.Error())
		return res, err
	}

	return res, err
}
