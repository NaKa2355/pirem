/*
設定ファイルを読み込んで、デーモンを立ち上げる
*/

package daemon

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"net"
	"os"
	"syscall"

	modules "github.com/NaKa2355/pirem/build"
	"github.com/NaKa2355/pirem/internal/app/pirem/controller/repository"
	"github.com/NaKa2355/pirem/internal/app/pirem/controller/web"
	"github.com/NaKa2355/pirem/internal/app/pirem/infrastructure/server"
	"github.com/NaKa2355/pirem/internal/app/pirem/usecases/boundary"
	"github.com/NaKa2355/pirem/internal/app/pirem/usecases/interactor"
	"github.com/NaKa2355/pirem/pkg/logger"

	"golang.org/x/exp/slog"
)

type Daemon struct {
	srv    *server.Server
	logger logger.Logger
}

type DeviceConfig struct {
	Name       string          `json:"name"`
	ID         string          `json:"id"`
	ModuleName string          `json:"module_name"`
	Config     json.RawMessage `json:"config"`
}

type Config struct {
	Devices          []DeviceConfig `json:"devices"`
	EnableReflection bool           `json:"enable_reflection"`
	Debug            bool           `json:"debug"`
}

func (d *Daemon) readConf(filePath string) (*Config, error) {
	config := &Config{}
	config_data, err := os.ReadFile(filePath)
	if err != nil {
		return config, err
	}
	err = json.Unmarshal(config_data, config)
	return config, err
}

func (d *Daemon) loadDevices(interactor *interactor.Interactor, devsConf []DeviceConfig) (err error) {
	for _, devConf := range devsConf {

		module, ok := modules.Modules[devConf.ModuleName]
		if !ok {
			err = errors.Join(err, fmt.Errorf("module \"%s\" not found", devConf.ModuleName))
			continue
		}

		addDevReq := boundary.AddDeviceInput{
			ID:         devConf.ID,
			DeviceName: devConf.Name,
			Config:     devConf.Config,
			Module:     module,
		}

		_err := interactor.AddDevice(context.Background(), addDevReq)
		if _err != nil {
			err = errors.Join(err, _err)
			continue
		}

		d.logger.Info(
			"device loaded",
			"module name", devConf.ModuleName,
			"device name", devConf.Name,
			"device id", devConf.ID,
		)
	}

	return err
}

func New(configPath string) (d *Daemon, err error) {
	d = &Daemon{}

	level := new(slog.LevelVar)
	handler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: level,
	})
	d.logger = slog.New(handler)

	repo := repository.New()
	interactor := interactor.New(repo)
	web := web.New(interactor)

	//load config file
	config, err := d.readConf(configPath)
	if err != nil {
		d.logger.Error("faild to read config file",
			"config file path", configPath,
			"error", err.Error())
		return d, err
	}

	if config.Debug {
		level.Set(slog.LevelDebug)
	}

	err = d.loadDevices(interactor, config.Devices)
	if err != nil {
		d.logger.Error("faild to load plugin(s)",
			"error", err.Error())
	}

	d.srv = server.New(web, config.EnableReflection, d.logger)
	return d, nil
}

// run until signal comes
func (d *Daemon) Start(domainSocket string) error {
	listener, err := net.Listen("unix", domainSocket)
	if err != nil {
		d.logger.Error("faild to make a socket", "error", err)
		return err
	}

	err = os.Chmod(domainSocket, 0770)
	if err != nil {
		d.logger.Error("faild to change permisson", "error", err)
		return err
	}

	d.srv.Start(listener)

	d.logger.Info(
		"daemon started",
		"unix domain socket path", domainSocket,
	)

	d.srv.WaitSigAndStop(syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT)
	d.logger.Info("shutting down daemon...")
	d.logger.Info("stopped daemon")
	return nil
}
