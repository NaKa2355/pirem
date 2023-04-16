package server

import (
	"fmt"
	"net"
	"os"
	"os/signal"

	apiremv1 "github.com/NaKa2355/irdeck-proto/gen/go/pirem/api/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	server *grpc.Server
}

func New(handler apiremv1.PiRemServiceServer, useReflectiton bool) *Server {
	s := &Server{}

	apiremv1.RegisterPiRemServiceServer(s.server, handler)
	if useReflectiton {
		reflection.Register(s.server)
	}
	return s
}

func (s *Server) Start(domainSocketPath string) error {
	listenPort, err := net.Listen("unix", domainSocketPath)
	if err != nil {
		return err
	}
	go func() {
		defer listenPort.Close()
		err := s.server.Serve(listenPort)
		fmt.Println(err)
	}()
	return nil
}

func (s *Server) WaitSigAndStop(sig ...os.Signal) {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, sig...)
	<-sigCh
	s.server.GracefulStop()
}
