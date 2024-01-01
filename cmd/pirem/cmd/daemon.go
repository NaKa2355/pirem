/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/NaKa2355/pirem/config"
	"github.com/NaKa2355/pirem/internal/app/pirem/infra/db"
	"github.com/NaKa2355/pirem/internal/app/pirem/infra/devices"
	"github.com/NaKa2355/pirem/internal/app/pirem/infra/web"
	"github.com/NaKa2355/pirem/internal/app/pirem/registry"
	"github.com/NaKa2355/pirem/internal/app/pirem/usecases/boundary"
	"github.com/NaKa2355/pirem/internal/app/pirem/usecases/controllers"
	"github.com/NaKa2355/pirem/internal/app/pirem/usecases/gateways"
	"github.com/NaKa2355/pirem/internal/app/pirem/usecases/interactor"
	"github.com/NaKa2355/pirem/pkg/logger"
	"github.com/NaKa2355/pirem/tools/modules"
	"golang.org/x/exp/slog"

	"github.com/spf13/cobra"
)

const ConfigFilePath = "/etc/piremd.json"

// daemonCmd represents the daemon command
var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "execute as daemon",
	Long: `execute as daemon. 
	config file: /etc/piremd.json
	service file: /lib/systemd/system/piremd.service
	socket file: /tmp/pirem.sock`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return startDaemon()
	},
}

func startDaemon() error {
	level := new(slog.LevelVar)
	handler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: level,
	})
	var logger logger.Logger = slog.New(handler)
	config, err := config.ReadConfig(ConfigFilePath)
	if err != nil {
		logger.Error("faild to read config file", "file_path", ConfigFilePath)
	}
	r := registry.NewRegistry(&Handler{
		logger:   logger,
		devsConf: config.Devices,
	})
	r.SolveDependencies()
	r.Start()
	return nil
}

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

	repo, err := db.New(homeDirectory + "/.pirem")
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

func init() {
	rootCmd.AddCommand(daemonCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// daemonCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// daemonCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
