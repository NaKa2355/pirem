package device

import (
	"context"
	"encoding/json"

	plugin "github.com/NaKa2355/pirem/pkg/plugin/driver"
	"github.com/NaKa2355/pirem/third_party/airer/internal/app/airer/driver"
)

type Device struct {
	d    *driver.Driver
	info *plugin.DeviceInfo
}

type DeviceConfig struct {
	SpiDevFile string `json:"spi_dev_file"`
	BusyPin    int    `json:"busy_pin"`
}

const DriverVersion = "0.1.0"

func (dev *Device) setInfo() error {
	var err error = nil
	dev.info = &plugin.DeviceInfo{}
	firmVersion, err := dev.d.GetVersion()
	if err != nil {
		return err
	}
	dev.info.FirmwareVersion = firmVersion
	dev.info.DriverVersion = DriverVersion
	return nil
}

func NewDevice(jsonConf json.RawMessage) (dev *Device, err error) {
	dev = &Device{}
	conf := DeviceConfig{}
	err = json.Unmarshal(jsonConf, &conf)
	if err != nil {
		return dev, err
	}

	d, err := driver.New(conf.SpiDevFile, conf.BusyPin)
	if err != nil {
		return dev, err
	}
	dev.d = d
	if err := dev.setInfo(); err != nil {
		return dev, err
	}

	return dev, err
}

func (dev *Device) GetInfo(ctx context.Context) (*plugin.DeviceInfo, error) {
	return dev.info, nil
}

func (dev *Device) SendIR(ctx context.Context, irData *plugin.IRData) error {
	return dev.d.SendIr(ctx, convertToDriverIrRawData(irData.PluseNanoSec))
}

func (dev *Device) ReceiveIR(ctx context.Context) (*plugin.IRData, error) {
	irData := &plugin.IRData{}
	data, err := dev.d.ReceiveIr(ctx)
	if err != nil {
		return irData, err
	}
	irData.CarrierFreqKiloHz = 40
	irData.PluseNanoSec = convertToApiIrRawData(data)
	return irData, nil
}

func (dev *Device) Drop() error {
	return dev.d.Close()
}
