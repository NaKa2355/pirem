package rest_server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"

	adapter "github.com/NaKa2355/pirem/internal/app/pirem/adapter/rest_api"
	"github.com/NaKa2355/pirem/internal/app/pirem/usecases/boundary"
	"github.com/NaKa2355/pirem/internal/app/pirem/usecases/controllers"
	"github.com/NaKa2355/pirem/pkg/logger"
)

type UnixDomainServer struct {
	path   string
	server http.Server
}

func NewUnixDomainRestServer(domainSocketPath string, boundary boundary.Boundary, logger logger.Logger) controllers.Web {
	handler := adapter.NewRestApiAdapter(boundary)
	return &UnixDomainServer{
		path:   domainSocketPath,
		server: http.Server{Handler: &handler},
	}
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
	s.server.Serve(listener)
}

func (s *UnixDomainServer) Quit() {
	s.server.Shutdown(context.Background())
}
