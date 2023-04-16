package daemon

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"syscall"

	devctr "github.com/NaKa2355/pirem/internal/app/pirem/controller/device"
	"github.com/NaKa2355/pirem/internal/app/pirem/controller/repository"
	"github.com/NaKa2355/pirem/internal/app/pirem/controller/web"
	"github.com/NaKa2355/pirem/internal/app/pirem/driver/server"
	entdev "github.com/NaKa2355/pirem/internal/app/pirem/entity/device"
	interactor "github.com/NaKa2355/pirem/internal/app/pirem/usecases/interactor"
)

type Daemon struct {
	srv        *server.Server
	interactor *interactor.Interactor
	config     *Config
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

func New(configPath string) (*Daemon, error) {
	var err error = nil
	d := &Daemon{}
	repo := repository.New()
	d.interactor = interactor.New(repo)
	web := web.New(d.interactor)

	//load config file
	d.config, err = readConf(configPath)
	if err != nil {
		fmt.Println(err)
		return d, err
	}

	//load entdev plugins
	for _, devConf := range d.config.Devices {
		driver, err := devctr.New(devConf.PluginPath, devConf.Config)
		if err != nil {
			errors.Join(err)
			continue
		}

		dev, err := entdev.New(devConf.ID, devConf.Name, driver.Info, driver)
		if err != nil {
			errors.Join(err)
			continue
		}

		err = repo.CreateDevice(dev)
		if err != nil {
			errors.Join(err)
			continue
		}

		fmt.Println(
			MsgLoadedDevice,
			"plugin file path", devConf.PluginPath,
			"entdev name", devConf.Name,
			"entdev id", devConf.ID,
		)
	}

	d.srv = server.New(web, d.config.EnableReflection)
	return d, nil
}

// run until signal comes
func (d *Daemon) Start(domainSocket string) error {
	if err := d.srv.Start(domainSocket); err != nil {
		fmt.Println(MsgFaildStartDaemon, "error", err)
		return err
	}
	fmt.Println(
		MsgStartDaemon,
		"unix domain socket path", domainSocket,
	)

	d.srv.WaitSigAndStop(syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT)
	fmt.Println(MsgShuttingDownDaemon)
	defer d.interactor.Drop()
	fmt.Println(MsgStopDaemon)
	return nil
}
