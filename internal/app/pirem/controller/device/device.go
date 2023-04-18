package device

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"plugin"

	"github.com/NaKa2355/pirem/internal/app/pirem/entity"
	entdev "github.com/NaKa2355/pirem/internal/app/pirem/entity/device"
	ir "github.com/NaKa2355/pirem/internal/app/pirem/entity/ir"
	"github.com/NaKa2355/pirem/pkg/plugin/driver"
)

var _ entdev.Driver = &Device{}

type Device struct {
	Info   entdev.Info
	Driver driver.Driver
}

func convertErr(_err error) error {
	if _err == nil {
		return nil
	}
	var code entity.ErrCode
	switch err := _err.(type) {
	case *driver.Error:
		switch err.Code {
		case driver.CodeBusy:
			code = entity.CodeBusy
		case driver.CodeDevice:
			code = entity.CodeInternal
		case driver.CodeInvaildInput:
			code = entity.CodeInvaildInput
		case driver.CodeTimeout:
			code = entity.CodeTimeout
		}
		return entity.WrapErr(code, err)
	}
	return _err
}

// デバイスを操作する構造体をプラグインから取得する
func New(pluginPath string, conf json.RawMessage) (*Device, error) {
	dev := &Device{}
	p, err := plugin.Open(pluginPath)
	if err != nil {
		return dev, entity.WrapErr(
			entity.CodeInvaildInput,
			fmt.Errorf("faild to open plugin: %w", err),
		)
	}

	s, err := p.Lookup("GetDriver")
	if err != nil {
		return dev, err
	}

	GetDriver, ok := s.(func(json.RawMessage) (driver.Driver, error))
	if !ok {
		return dev, entity.WrapErr(
			entity.CodeInternal,
			errors.New("function type is wrong"),
		)
	}

	d, err := GetDriver(conf)
	if err != nil {
		return dev, err
	}

	info, err := d.GetInfo(context.Background())
	if err != nil {
		return dev, convertErr(err)
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
	sendData := &driver.IRData{
		CarrierFreqKiloHz: rawIRData.CarrierFreqKiloHz,
		PluseNanoSec:      rawIRData.PluseNanoSec,
	}

	err := d.Driver.SendIR(ctx, sendData)
	return convertErr(err)
}

func (d *Device) ReceiveIR(ctx context.Context) (ir.Data, error) {
	rawIRData := ir.RawData{}
	irData, err := d.Driver.ReceiveIR(ctx)
	if err != nil {
		return rawIRData, convertErr(err)
	}

	rawIRData.CarrierFreqKiloHz = irData.CarrierFreqKiloHz
	rawIRData.PluseNanoSec = irData.PluseNanoSec

	return rawIRData, nil
}
