package main

import (
	"github.com/NaKa2355/pirem/internal/app/pirem/usecases/boundary"
	"github.com/NaKa2355/pirem/internal/app/pirem/usecases/controllers"
	"github.com/NaKa2355/pirem/internal/app/pirem/usecases/gateways"
	"github.com/NaKa2355/pirem/pkg/logger"
	"golang.org/x/exp/slog"
)

type Handler struct{}

func (h *Handler) DeviceFactory() (controllers.IRDevice, error) {

	return
}

func (h *Handler) RepositoryFactory() (gateways.Repository, error) {

}

func (h *Handler) BoudaryFactory(device controllers.IRDevice, repo gateways.Repository) (boundary.Boundary, error) {

}

func (h *Handler) EntryPoint(boundary boundary.Boundary) {

}

func main() {
	var logger logger.Logger = &slog.Logger{}

}
