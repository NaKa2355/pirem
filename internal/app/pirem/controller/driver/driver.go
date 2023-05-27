package driver

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/NaKa2355/pirem/internal/app/pirem/entity"
	"github.com/NaKa2355/pirem/internal/app/pirem/entity/device"
	"github.com/NaKa2355/pirem/internal/app/pirem/entity/ir"
	"github.com/NaKa2355/pirem/internal/app/pirem/usecases/driver"
	"github.com/NaKa2355/pirem/pkg/module/v1"
)

var _ device.Driver = &Driver{}

type Driver struct {
	Info   device.Info
	Driver module.Driver
}

func convertErr(_err error) error {
	if _err == nil {
		return nil
	}
	var code driver.ErrCode
	switch err := _err.(type) {
	case *module.Error:
		switch err.Code {
		case module.CodeBusy:
			code = driver.CodeBusy
		case module.CodeDevice:
			code = driver.CodeInternal
		case module.CodeInvaildInput:
			code = driver.CodeInvaildInput
		case module.CodeTimeout:
			code = driver.CodeTimeout
		default:
			code = driver.CodeUnknown
		}
		return driver.WrapErr(code, err)
	}
	return _err
}

// デバイスを操作する構造体をプラグインから取得する
func New(pluginName string, devConf json.RawMessage, modules map[string]module.Module) (*Driver, error) {
	dev := &Driver{}

	m, ok := modules[pluginName]
	if !ok {
		return dev, entity.WrapErr(
			entity.CodeInvaildInput,
			fmt.Errorf("pluin not found"),
		)
	}

	d, err := m.NewDriver(devConf)
	if err != nil {
		return dev, err
	}

	info, err := d.GetInfo(context.Background())
	if err != nil {
		return dev, convertErr(err)
	}

	dev.Driver = d
	dev.Info = device.Info{
		CanSend:         info.CanReceive,
		CanReceive:      info.CanReceive,
		DriverVersion:   info.DriverVersion,
		FirmwareVersion: info.FirmwareVersion,
	}

	return dev, nil
}

func (d *Driver) SendIR(ctx context.Context, irData ir.Data) error {
	rawIRData := irData.ConvertToRaw()
	sendData := &module.IRData{
		CarrierFreqKiloHz: rawIRData.CarrierFreqKiloHz,
		PluseNanoSec:      rawIRData.PluseNanoSec,
	}

	err := d.Driver.SendIR(ctx, sendData)
	return convertErr(err)
}

func (d *Driver) ReceiveIR(ctx context.Context) (ir.Data, error) {
	rawIRData := &ir.RawData{}
	irData, err := d.Driver.ReceiveIR(ctx)
	if err != nil {
		return rawIRData, convertErr(err)
	}

	rawIRData.CarrierFreqKiloHz = irData.CarrierFreqKiloHz
	rawIRData.PluseNanoSec = irData.PluseNanoSec

	return rawIRData, nil
}

func (d *Driver) GetDeviceInfo() *device.Info {
	return &d.Info
}
