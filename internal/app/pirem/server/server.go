package server

import (
	"net"
	"os"
	"os/signal"

	apiremv1 "github.com/NaKa2355/irdeck-proto/gen/go/pirem/api/v1"
	"github.com/hashicorp/go-hclog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	logger hclog.Logger
	server *grpc.Server
}

func New(logger hclog.Logger, handler apiremv1.PiRemServiceServer, useReflectiton bool) *Server {
	piremServer := &Server{}
	piremServer.logger = logger
	piremServer.server = grpc.NewServer(grpc.UnaryInterceptor(ErrorWrapping))
	apiremv1.RegisterPiRemServiceServer(piremServer.server, handler)
	if useReflectiton {
		reflection.Register(piremServer.server)
	}
	return piremServer
}

func (s *Server) Start(domainSocketPath string) error {
	listenPort, err := net.Listen("unix", domainSocketPath)
	if err != nil {
		return err
	}
	go func() {
		defer listenPort.Close()
		err := s.server.Serve(listenPort)
		if err != nil {
			s.logger.Error(
				"server error",
				"error", err,
			)
		}
	}()
	return nil
}

func (s *Server) WaitSigAndStop(sig ...os.Signal) {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, sig...)
	<-sigCh
	s.server.GracefulStop()
}
