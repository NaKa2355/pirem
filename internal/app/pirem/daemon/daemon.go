package daemon

import (
	"encoding/json"
	"os"
	"syscall"

	"github.com/NaKa2355/pirem/internal/app/pirem/controller"
	"github.com/NaKa2355/pirem/internal/app/pirem/device"
	"github.com/NaKa2355/pirem/internal/app/pirem/entity"
	server "github.com/NaKa2355/pirem/internal/app/pirem/server"

	"golang.org/x/exp/slog"
)

type Daemon struct {
	srv    *server.Server
	entity *entity.Entity
	config *Config
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

func readConf(filePath string) (*Config, error) {
	config := &Config{}
	config_data, err := os.ReadFile(filePath)
	if err != nil {
		return config, err
	}
	err = json.Unmarshal(config_data, config)
	return config, err
}

func addDevice(devConf DeviceConfig, entity *entity.Entity) error {
	dev, err := device.NewFromPlugin(devConf.ID, devConf.Name, devConf.Config, devConf.PluginPath)
	if err != nil {
		return err
	}
	entity.AddDevice(dev)
	return nil
}

func New(logger *slog.Logger, configPath string) (*Daemon, error) {
	var err error = nil
	d := &Daemon{}
	d.logger = logger
	d.entity = entity.New()

	//load config file
	d.config, err = readConf(configPath)
	if err != nil {
		logger.Error(MsgFaildLoadConfig, err)
		return d, err
	}

	//load device plugins
	for _, devConf := range d.config.Devices {
		err := addDevice(devConf, d.entity)
		if err != nil {
			logger.Error(
				MsgFaildLoadDev,
				err,
				"plugin file path", devConf.PluginPath,
			)
			continue
		}
		logger.Info(
			MsgLoadedDevice,
			"plugin file path", devConf.PluginPath,
			"device name", devConf.Name,
			"device id", devConf.ID,
		)
	}

	controller := controller.New(d.entity)
	d.srv = server.New(logger, controller, d.config.EnableReflection)
	return d, nil
}

// run until signal comes
func (d *Daemon) Start(domainSocket string) error {
	if err := d.srv.Start(domainSocket); err != nil {
		d.logger.Error(MsgFaildStartDaemon, err)
		return err
	}
	d.logger.Info(
		MsgStartDaemon,
		"unix domain socket path", domainSocket,
	)

	d.srv.WaitSigAndStop(syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT)
	d.logger.Info(MsgShuttingDownDaemon)
	d.entity.Drop()
	d.logger.Info(MsgStopDaemon)
	return nil
}
