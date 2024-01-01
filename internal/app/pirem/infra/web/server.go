package web

import (
	"net"

	pirem "github.com/NaKa2355/pirem/api/gen/go/api/v1"
	adapter "github.com/NaKa2355/pirem/internal/app/pirem/adapter/web"
	"github.com/NaKa2355/pirem/internal/app/pirem/usecases/boundary"
	"github.com/NaKa2355/pirem/internal/app/pirem/usecases/controllers"
	"google.golang.org/grpc"
)

type Server struct {
	s *grpc.Server
	l net.Listener
}

func NewServer(l net.Listener, boundary boundary.Boundary) controllers.Web {
	handler := adapter.NewRequestHandler(boundary)
	s := grpc.NewServer()
	pirem.RegisterPiRemServiceServer(s, handler)
	return &Server{
		s: s,
		l: l,
	}
}

func (s *Server) StartListen() {
	s.s.Serve(s.l)
}

func (s *Server) Quit() {
	s.s.GracefulStop()
}
