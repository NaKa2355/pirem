package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/NaKa2355/pirem/internal/app/pirem/controller/driver"
	"github.com/NaKa2355/pirem/internal/app/pirem/controller/repository"
	"github.com/NaKa2355/pirem/internal/app/pirem/controller/web"
	"github.com/NaKa2355/pirem/internal/app/pirem/entity/device"
	"github.com/NaKa2355/pirem/internal/app/pirem/infrastructure/server"
	ud "github.com/NaKa2355/pirem/internal/app/pirem/usecases/driver"
	"github.com/NaKa2355/pirem/internal/app/pirem/usecases/interactor"
	"golang.org/x/exp/slog"
)

func main() {
	logger := slog.New(slog.Default().Handler())
	repo := repository.New()
	i := interactor.New(repo)
	dev, err := device.New("1", "test", driver.NewMock(nil, nil))
	if err != nil {
		panic(err)
	}
	errDev, _ := device.New("2", "error device", driver.NewMock(
		ud.WrapErr(ud.CodeBusy, fmt.Errorf("error")),
		ud.WrapErr(ud.CodeTimeout, fmt.Errorf("error"))))

	repo.CreateDevice(dev)
	repo.CreateDevice(errDev)
	h := web.New(i)

	port := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	srv := server.New(h, true, logger)

	go func() {
		log.Printf("start gRPC server port: %v", port)
		srv.Start(listener)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	srv.WaitSigAndStop(<-quit)
}
