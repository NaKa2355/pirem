/*
設定ファイルを読み込んで、デーモンを立ち上げる
*/

package daemon

import (
	"encoding/json"
	"errors"
	"os"
	"syscall"

	"github.com/NaKa2355/pirem/internal/app/pirem/controller/driver"
	"github.com/NaKa2355/pirem/internal/app/pirem/controller/repository"
	"github.com/NaKa2355/pirem/internal/app/pirem/controller/web"
	"github.com/NaKa2355/pirem/internal/app/pirem/driver/server"
	"github.com/NaKa2355/pirem/internal/app/pirem/entity/device"
	"github.com/NaKa2355/pirem/internal/app/pirem/usecases/interactor"
	"golang.org/x/exp/slog"
)

type Daemon struct {
	srv    *server.Server
	logger *slog.Logger
}

type DeviceConfig struct {
	Name       string          `json:"name"`
	ID         string          `json:"id"`
	PluginPath string          `json:"plugin_path"`
	Config     json.RawMessage `json:"config"`
}

type Config struct {
	Devices          []DeviceConfig `json:"devices"`
	EnableReflection bool           `json:"enable_reflection"`
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

func (d *Daemon) loadDevices(repo *repository.Repository, devsConf []DeviceConfig) (err error) {
	for _, devConf := range devsConf {
		drv, _err := driver.New(devConf.PluginPath, devConf.Config)
		if _err != nil {
			errors.Join(err, _err)
			continue
		}

		dev, _err := device.New(devConf.ID, devConf.Name, drv)
		if _err != nil {
			errors.Join(err, _err)
			continue
		}

		_err = repo.CreateDevice(dev)
		if _err != nil {
			errors.Join(err, _err)
			continue
		}

		d.logger.Info(
			"device loaded",
			"plugin file path", devConf.PluginPath,
			"device name", devConf.Name,
			"device id", devConf.ID,
		)
	}

	return err
}

func New(configPath string) (*Daemon, error) {
	var err error = nil
	d := &Daemon{}
	d.logger = slog.New(slog.Default().Handler())
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

	err = d.loadDevices(repo, config.Devices)
	if err != nil {
		d.logger.Error("faild to load plugin(s)",
			"error", err.Error())
	}

	d.srv = server.New(web, config.EnableReflection)
	return d, nil
}

// run until signal comes
func (d *Daemon) Start(domainSocket string) error {
	if err := d.srv.Start(domainSocket); err != nil {
		d.logger.Error("faild to start daemon", "error", err)
		return err
	}

	d.logger.Info(
		"daemon started",
		"unix domain socket path", domainSocket,
	)

	d.logger.Info("shutting down daemon...")
	d.srv.WaitSigAndStop(syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT)
	d.logger.Info("stopped daemon")
	return nil
}
