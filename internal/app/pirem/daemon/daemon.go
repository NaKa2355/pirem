package daemon

import (
	"encoding/json"
	"os"
	"syscall"

	"github.com/NaKa2355/pirem/internal/app/pirem/controller/web"
	server "github.com/NaKa2355/pirem/internal/app/pirem/server"
	interactor "github.com/NaKa2355/pirem/internal/app/pirem/usecases/interactor"
	"github.com/hashicorp/go-hclog"
)

type Daemon struct {
	srv        *server.Server
	interactor *interactor.Interactor
	config     *Config
	logger     hclog.Logger
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

func New(logger hclog.Logger, configPath string) (*Daemon, error) {
	var err error = nil
	d := &Daemon{}
	d.logger = logger
	d.interactor = interactor.New(logger)
	web := web.New(d.interactor)

	//load config file
	d.config, err = readConf(configPath)
	if err != nil {
		logger.Error(MsgFaildLoadConfig, "error", err)
		return d, err
	}

	//load device plugins
	for _, devConf := range d.config.Devices {
		err := d.interactor.AddDevice(
			devConf.ID, devConf.Name, devConf.PluginPath, devConf.Config)
		if err != nil {
			logger.Error(
				MsgFaildLoadDev,
				"error", err,
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

	d.srv = server.New(logger, web, d.config.EnableReflection)
	return d, nil
}

// run until signal comes
func (d *Daemon) Start(domainSocket string) error {
	if err := d.srv.Start(domainSocket); err != nil {
		d.logger.Error(MsgFaildStartDaemon, "error", err)
		return err
	}
	d.logger.Info(
		MsgStartDaemon,
		"unix domain socket path", domainSocket,
	)

	d.srv.WaitSigAndStop(syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT)
	d.logger.Info(MsgShuttingDownDaemon)
	defer d.interactor.Drop()
	d.logger.Info(MsgStopDaemon)
	return nil
}
