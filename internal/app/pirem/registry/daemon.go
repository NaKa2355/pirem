/*
設定ファイルを読み込んで、デーモンを立ち上げる
*/

package registry

import (
	"encoding/json"
	"errors"
	"fmt"
	"os/signal"

	"net"
	"os"
	"syscall"

	dataAccess "github.com/NaKa2355/pirem/internal/app/pirem/infra/db"
	"github.com/NaKa2355/pirem/internal/app/pirem/infra/irdevice"
	"github.com/NaKa2355/pirem/internal/app/pirem/infra/web"
	"github.com/NaKa2355/pirem/internal/app/pirem/usecases/interactor"
	"github.com/NaKa2355/pirem/pkg/logger"
	"github.com/NaKa2355/pirem/tools/modules"

	"golang.org/x/exp/slog"
)

type Daemon struct {
	i      *interactor.Interactor
	logger logger.Logger
}

type DeviceConfig struct {
	Name       string          `json:"name"`
	ID         string          `json:"id"`
	ModuleName string          `json:"module_name"`
	Config     json.RawMessage `json:"config"`
}

type Config struct {
	Devices                 []DeviceConfig `json:"devices"`
	DeviceMutexLockDeadline int            `json:"device_mutex_lock_deadline"`
	EnableReflection        bool           `json:"enable_reflection"`
	Debug                   bool           `json:"debug"`
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

func (d *Daemon) loadDevices(devices *irdevice.IRDevices, devsConf []DeviceConfig) (err error) {
	for _, devConf := range devsConf {

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

	//load config file
	config, err := d.readConf(configPath)
	if err != nil {
		d.logger.Error("faild to read config file",
			"config file path", configPath,
			"error", err.Error())
		return d, err
	}

	homeDirectory, err := os.UserHomeDir()
	if err != nil {
		panic(fmt.Errorf("faild to get home directory: %w", err))
	}

	repo, err := dataAccess.New(homeDirectory + "/.pirem")
	if err != nil {
		panic(fmt.Errorf("faild to create repository: %w", err))
	}

	devices := irdevice.NewIRDevices()
	err = d.loadDevices(devices, config.Devices)
	if err != nil {
		d.logger.Error("faild to load plugin(s)",
			"error", err.Error())
	}

	d.i = interactor.NewInteractor(repo, devices)
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

	server := web.NewServer(listener, d.i)

	go server.StartListen()
	d.logger.Info(
		"daemon started",
		"unix domain socket path", domainSocket,
	)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	server.Quit()
	d.logger.Info("shutting down daemon...")
	d.logger.Info("stopped daemon")
	return nil
}
