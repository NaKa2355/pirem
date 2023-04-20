package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	apiremv1 "github.com/NaKa2355/irdeck-proto/gen/go/pirem/api/v1"
	"github.com/NaKa2355/pirem/internal/app/pirem/controller/driver"
	"github.com/NaKa2355/pirem/internal/app/pirem/controller/repository"
	"github.com/NaKa2355/pirem/internal/app/pirem/controller/web"
	"github.com/NaKa2355/pirem/internal/app/pirem/entity/device"
	"github.com/NaKa2355/pirem/internal/app/pirem/usecases/interactor"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	repo := repository.New()
	i := interactor.New(repo)
	dev, err := device.New("1", "test", driver.NewMock(nil, nil))
	if err != nil {
		panic(err)
	}

	repo.CreateDevice(dev)
	h := web.New(i)

	port := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(web.ErrorUnaryServerInterceptor))
	apiremv1.RegisterPiRemServiceServer(s, h)
	reflection.Register(s)

	go func() {
		log.Printf("start gRPC server port: %v", port)
		s.Serve(listener)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()
}
