package web

import (
	"fmt"
	"net"
	"os"

	pirem "github.com/NaKa2355/pirem/api/gen/go/api/v1"
	adapter "github.com/NaKa2355/pirem/internal/app/pirem/adapter/web"
	"github.com/NaKa2355/pirem/internal/app/pirem/usecases/boundary"
	"github.com/NaKa2355/pirem/internal/app/pirem/usecases/controllers"
	"google.golang.org/grpc"
)

type UnixDomainServer struct {
	s    *grpc.Server
	path string
}

func NewUnixDomainServer(domainSocketPath string, boundary boundary.Boundary) (controllers.Web, error) {
	handler := adapter.NewRequestHandler(boundary)
	gs := grpc.NewServer()
	pirem.RegisterPiRemServiceServer(gs, handler)
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
