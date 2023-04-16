package device

import (
	"context"
	"encoding/json"
	"errors"
	"plugin"

	entdev "github.com/NaKa2355/pirem/internal/app/pirem/entity/device"
	ir "github.com/NaKa2355/pirem/internal/app/pirem/entity/ir"
	"github.com/NaKa2355/pirem/pkg/plugin/driver"
)

var _ entdev.Driver = &Device{}

type Device struct {
	Info   entdev.Info
	Driver driver.Driver
}

// デバイスを操作する構造体をプラグインから取得する
func New(pluginPath string, conf json.RawMessage) (*Device, error) {
	dev := &Device{}
	p, err := plugin.Open(pluginPath)
	if err != nil {
		return dev, err
	}

	s, err := p.Lookup("GetDriver")
	if err != nil {
		return dev, err
	}

	GetDriver, ok := s.(func(json.RawMessage) (driver.Driver, error))
	if !ok {
		return dev, errors.New("function type is wrong")
	}

	d, err := GetDriver(conf)
	if err != nil {
		return dev, err
	}

	info, err := d.GetInfo(context.Background())
	if err != nil {
		return dev, err
	}

	dev.Driver = d
	dev.Info = entdev.Info{
		CanSend:         info.CanReceive,
		CanReceive:      info.CanReceive,
		DriverVersion:   info.DriverVersion,
		FirmwareVersion: info.FirmwareVersion,
	}

	return dev, nil
}

func (d *Device) SendIR(ctx context.Context, irData ir.Data) error {
	rawIRData := irData.ConvertToRaw()
	sendData := driver.IRData{
		CarrierFreqKiloHz: rawIRData.CarrierFreqKiloHz,
		PluseNanoSec:      rawIRData.PluseNanoSec,
	}

	return d.Driver.SendIR(ctx, sendData)
}

func (d *Device) ReceiveIR(ctx context.Context) (ir.Data, error) {
	rawIRData := ir.RawData{}
	irData, err := d.Driver.ReceiveIR(ctx)
	if err != nil {
		return rawIRData, err
	}

	rawIRData.CarrierFreqKiloHz = irData.CarrierFreqKiloHz
	rawIRData.PluseNanoSec = irData.PluseNanoSec

	return rawIRData, nil
}

func (d *Device) GetInfo() (entdev.Info, error) {
	return d.Info, nil
}
