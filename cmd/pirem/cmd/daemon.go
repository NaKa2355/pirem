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
	"time"

	"github.com/NaKa2355/pirem/cmd/pirem/utils"
	"github.com/NaKa2355/pirem/config"
	"github.com/NaKa2355/pirem/internal/app/pirem/infra/db"
	"github.com/NaKa2355/pirem/internal/app/pirem/infra/devices"
	"github.com/NaKa2355/pirem/internal/app/pirem/infra/grpc_server"
	"github.com/NaKa2355/pirem/internal/app/pirem/infra/rest_server"
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

const configFilePath = "/etc/piremd.json"
const serviceFilePath = "/lib/systemd/system/piremd.service"
const grpcSocketFilePath = utils.GrpcDomainSocketPath
const restSocketFilePath = utils.RestDomainSocketPath
const longDiscribe = "execute as daemon.\nconfig file path(default) " +
	configFilePath +
	"\nservice file path: " + serviceFilePath +
	"\ngrpc api socket file path: " + grpcSocketFilePath +
	"\nrest api socket file path: " + restSocketFilePath

// daemonCmd represents the daemon command
var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "execute as daemon",
	Long:  longDiscribe,
	RunE: func(cmd *cobra.Command, args []string) error {
		configFilePath, err := cmd.Flags().GetString("config")
		if err != nil {
			return err
		}
		return startDaemon(configFilePath)
	},
}

func startDaemon(configFilePath string) error {
	level := new(slog.LevelVar)
	handler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: level,
	})
	var logger logger.Logger = slog.New(handler)
	config, err := config.ReadConfig(configFilePath)
	if err != nil {
		logger.Error("faild to read config file",
			"file_path", configFilePath,
			"error", err,
		)
		return err
	}
	if config.Debug {
		level.Set(slog.LevelDebug)
	}
	r := registry.NewRegistry(&Handler{
		logger: logger,
		config: config,
	})
	err = r.SolveDependencies()
	if err != nil {
		logger.Error("faild to solve dependencies",
			"error", err,
		)
		return err
	}
	r.Start()
	return nil
}

type Handler struct {
	logger logger.Logger
	config *config.Config
}

func (h *Handler) DeviceFactory() (controllers.IRDevice, error) {
	var err error = nil
	devices := devices.NewIRDevices()
	for _, devConf := range h.config.Devices {
		module, ok := modules.Modules[devConf.ModuleName]
		if !ok {
			err = errors.Join(err, fmt.Errorf("module \"%s\" not found", devConf.ModuleName))
			continue
		}

		driver, _err := module.LoadDevice(devConf.Config)
		if _err != nil {
			err = errors.Join(err, _err)
			continue
		}

		_err = devices.AddDevice(
			devConf.ID,
			devConf.Name,
			driver,
			time.Duration(time.Duration(h.config.DeviceSendIRIntervalMs)*time.Millisecond),
			time.Duration(time.Duration(h.config.DeviceMutexLockDeadlineMs)*time.Millisecond),
		)
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
	grpc_s, err := grpc_server.NewUnixDomainGrpcServer(grpcSocketFilePath, boundary, h.config.EnableReflection, h.logger)
	if err != nil {
		return
	}
	rest_s := rest_server.NewUnixDomainRestServer(restSocketFilePath, boundary, h.logger)

	go grpc_s.StartListen()
	go rest_s.StartListen()

	h.logger.Info(
		"daemon started",
		"grpc api unix domain socket path", grpcSocketFilePath,
		"rest api unix domain socket path", restSocketFilePath,
	)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	h.logger.Info("shutting down daemon...")
	grpc_s.Quit()
	rest_s.Quit()
	h.logger.Info("daemon stopped")
}

func init() {
	rootCmd.AddCommand(daemonCmd)
	daemonCmd.Flags().String("config", configFilePath, "config file path")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// daemonCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// daemonCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
