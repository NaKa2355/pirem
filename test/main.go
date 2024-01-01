package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/NaKa2355/pirem/internal/app/pirem/infra/db"
	"github.com/NaKa2355/pirem/internal/app/pirem/infra/devices"
	"github.com/NaKa2355/pirem/internal/app/pirem/infra/web"
	"github.com/NaKa2355/pirem/internal/app/pirem/registry"
	"github.com/NaKa2355/pirem/internal/app/pirem/usecases/boundary"
	"github.com/NaKa2355/pirem/internal/app/pirem/usecases/controllers"
	"github.com/NaKa2355/pirem/internal/app/pirem/usecases/gateways"
	"github.com/NaKa2355/pirem/internal/app/pirem/usecases/interactor"
	"github.com/NaKa2355/pirem/pkg/logger"
	"golang.org/x/exp/slog"
)

type Handler struct {
	logger logger.Logger
}

func (h *Handler) DeviceFactory() (controllers.IRDevice, error) {
	d := devices.NewIRDevices()
	return d, nil
}

func (h *Handler) RepositoryFactory() (gateways.Repository, error) {
	return db.New("./pirem_db")
}

func (h *Handler) BoudaryFactory(device controllers.IRDevice, repo gateways.Repository) (boundary.Boundary, error) {
	return interactor.NewInteractor(repo, device), nil
}

func (h *Handler) EntryPoint(boundary boundary.Boundary) {
	s, err := web.NewUnixDomainServer("./pirem_sock", boundary)
	if err != nil {
		return
	}
	go s.StartListen()
	h.logger.Info(
		"daemon started",
		"unix domain socket path", "./pirem_sock",
	)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	h.logger.Info("shutting down daemon...")
	s.Quit()
	h.logger.Info("daemon stopped")
}

func main() {
	level := new(slog.LevelVar)
	handler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: level,
	})
	r := registry.NewRegistry(&Handler{logger: slog.New(handler)})
	r.SolveDependencies()
	r.Start()
}
