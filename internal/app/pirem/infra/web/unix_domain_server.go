package web

import (
	"fmt"
	"net"
	"os"

	pirem "github.com/NaKa2355/pirem-proto/gen/go/api/v1"
	adapter "github.com/NaKa2355/pirem/internal/app/pirem/adapter/web"
	"github.com/NaKa2355/pirem/internal/app/pirem/usecases/boundary"
	"github.com/NaKa2355/pirem/internal/app/pirem/usecases/controllers"
	"github.com/NaKa2355/pirem/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type UnixDomainServer struct {
	s    *grpc.Server
	path string
}

func NewUnixDomainServer(domainSocketPath string, boundary boundary.Boundary, enableReflection bool, logger logger.Logger) (controllers.Web, error) {
	handler := adapter.NewRequestHandler(boundary)
	gs := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			adapter.LoggingUnaryServerInterceptor(logger),
			adapter.ErrorUnaryServerInterceptor,
		),
	)
	pirem.RegisterPiRemServiceServer(gs, handler)
	if enableReflection {
		reflection.Register(gs)
	}
	return &UnixDomainServer{
		s:    gs,
		path: domainSocketPath,
	}, nil
}

func (s *UnixDomainServer) StartListen() {
	listener, err := net.Listen("unix", s.path)
	if err != nil {
		panic(fmt.Errorf("faild to start listener: %w", err))
	}
	err = os.Chmod(s.path, 0770)
	if err != nil {
		os.Remove(s.path)
		panic(fmt.Errorf("faild to get pernission: %w", err))
	}
	s.s.Serve(listener)
}

func (s *UnixDomainServer) Quit() {
	s.s.GracefulStop()
}
