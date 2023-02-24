package daemon

import (
	"encoding/json"
	"os"
	"pirem/internal/app/pirem/controller"
	"pirem/internal/app/pirem/device"
	"pirem/internal/app/pirem/entity"
	server "pirem/internal/app/pirem/server"
	"plugin"
	"syscall"

	device_controllerv1 "github.com/NaKa2355/pirem_pkg/device_controller"

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

func loadConfig(filePath string) (*Config, error) {
	config := &Config{}
	config_data, err := os.ReadFile(filePath)
	if err != nil {
		return config, err
	}
	err = json.Unmarshal(config_data, config)
	return config, err
}

func getAndAddDevfromConf(devConf DeviceConfig, entity *entity.Entity) error {
	p, err := plugin.Open(devConf.PluginPath)
	if err != nil {
		return err
	}

	getCtrlSym, err := p.Lookup("GetController")
	if err != nil {
		return err
	}

	getController := getCtrlSym.(func(json.RawMessage) (device_controllerv1.DeviceController, error))
	devCtrl, err := getController(devConf.Config)
	if err != nil {
		return err
	}

	dev, err := device.New(devConf.ID, devConf.Name, devCtrl)
	if err != nil {
		return err
	}

	return entity.AddDevice(dev)
}

func New(logger *slog.Logger, configPath string) (*Daemon, error) {
	var err error = nil
	d := &Daemon{}
	d.logger = logger
	d.entity = entity.New()

	//load config file
	d.config, err = loadConfig(configPath)
	if err != nil {
		logger.Error(FaildLoadConfig, err)
		return d, err
	}

	//load device plugins
	for _, devConf := range d.config.Devices {
		err := getAndAddDevfromConf(devConf, d.entity)
		if err != nil {
			logger.Error(
				FaildLoadDev,
				err,
				"plugin file path", devConf.PluginPath,
			)
			continue
		}
		logger.Info(
			LoadDevice,
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
		d.logger.Error(FaildStartDaemon, err)
		return err
	}
	d.logger.Info(
		StartDaemon,
		"unix domain socket path", domainSocket,
	)

	d.srv.WaitSigAndStop(syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT)
	d.logger.Info(ShuttingDownDaemon)
	d.entity.Drop()
	d.logger.Info(StopDaemon)
	return nil
}
