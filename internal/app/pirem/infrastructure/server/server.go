/*
grpcのサーバーをunixドメインソケットで立ち上げる
*/

package server

import (
	"net"
	"os"
	"os/signal"

	apiremv1 "github.com/NaKa2355/irdeck-proto/gen/go/pirem/api/v1"
	"github.com/NaKa2355/pirem/pkg/logger"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	server *grpc.Server
	logger logger.Logger
}

func New(handler apiremv1.PiRemServiceServer, useReflectiton bool, logger logger.Logger) *Server {
	s := &Server{}
	s.logger = logger
	s.server = grpc.NewServer(grpc.UnaryInterceptor(s.UnaryServerInterceptor))
	apiremv1.RegisterPiRemServiceServer(s.server, handler)
	if useReflectiton {
		reflection.Register(s.server)
	}
	return s
}

func (s *Server) Start(listener net.Listener) {
	go func() {
		defer listener.Close()
		s.server.Serve(listener)
	}()
}

func (s *Server) WaitSigAndStop(sig ...os.Signal) {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, sig...)
	<-sigCh
	s.server.GracefulStop()
}
