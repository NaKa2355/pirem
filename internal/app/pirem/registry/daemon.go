/*
設定ファイルを読み込んで、デーモンを立ち上げる
*/

package registry

import (
	"errors"
	"fmt"
	"os/signal"

	"os"
	"syscall"

	"github.com/NaKa2355/pirem/config"
	dataAccess "github.com/NaKa2355/pirem/internal/app/pirem/infra/db"
	"github.com/NaKa2355/pirem/internal/app/pirem/infra/devices"
	"github.com/NaKa2355/pirem/internal/app/pirem/infra/web"
	"github.com/NaKa2355/pirem/internal/app/pirem/usecases/boundary"
	"github.com/NaKa2355/pirem/internal/app/pirem/usecases/controllers"
	"github.com/NaKa2355/pirem/internal/app/pirem/usecases/gateways"
	"github.com/NaKa2355/pirem/internal/app/pirem/usecases/interactor"
	"github.com/NaKa2355/pirem/pkg/logger"
	"github.com/NaKa2355/pirem/tools/modules"
	"golang.org/x/exp/slog"
)

type Handler struct {
	logger   logger.Logger
	devsConf []config.DeviceConfig
}

func (h *Handler) DeviceFactory() (controllers.IRDevice, error) {
	var err error = nil
	devices := devices.NewIRDevices()
	for _, devConf := range h.devsConf {
		module, ok := modules.Modules[devConf.ModuleName]
		if !ok {
			err = errors.Join(err, fmt.Errorf("module \"%s\" not found", devConf.ModuleName))
			continue
		}

		driver, _err := module.NewDriver(devConf.Config)
		if _err != nil {
			err = errors.Join(err, _err)
			continue
		}

		_err = devices.AddDevice(devConf.ID, devConf.Name, driver, 100, 5000)
		if _err != nil {
			err = errors.Join(err, _err)
			continue
		}
		h.logger.Info(
			"device loaded",
			"module name", devConf.ModuleName,
			"device name", devConf.Name,
			"device id", devConf.ID,
		)
	}
	return devices, err
}

func (h *Handler) RepositoryFactory() (gateways.Repository, error) {
	homeDirectory, err := os.UserHomeDir()
	if err != nil {
		panic(fmt.Errorf("faild to get home directory: %w", err))
	}

	repo, err := dataAccess.New(homeDirectory + "/.pirem")
	if err != nil {
		panic(fmt.Errorf("faild to create repository: %w", err))
	}
	return repo, err
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

func NewDaemon(confFile string) *Registry {
	level := new(slog.LevelVar)
	handler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: level,
	})
	var logger logger.Logger = slog.New(handler)
	config, err := config.ReadConfig(confFile)
	if err != nil {
		logger.Error("faild to read config file", "file_path", confFile)
	}
	r := NewRegistry(&Handler{
		logger:   logger,
		devsConf: config.Devices,
	})
	r.SolveDependencies()
	return r
}
