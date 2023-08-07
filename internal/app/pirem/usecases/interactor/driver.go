package interactor

import (
	"context"

	"github.com/NaKa2355/pirem/internal/app/pirem/entity/device"
	"github.com/NaKa2355/pirem/internal/app/pirem/entity/ir"
	"github.com/NaKa2355/pirem/pkg/module/v1"
)

type Driver struct {
	Driver module.Driver
	info   *device.Info
}

var _ device.Driver = &Driver{}

func NewDriver(driver module.Driver) (*Driver, error) {
	_info, err := driver.GetInfo(context.Background())
	info := &device.Info{
		CanSend:         _info.CanSend,
		CanReceive:      _info.CanReceive,
		DriverVersion:   _info.DriverVersion,
		FirmwareVersion: _info.FirmwareVersion,
	}

	d := &Driver{
		Driver: driver,
		info:   info,
	}

	return d, err
}

func (d *Driver) SendIR(ctx context.Context, in ir.Data) error {
	irdata := &module.IRData{
		CarrierFreqKiloHz: in.ConvertToRaw().CarrierFreqKiloHz,
		PluseNanoSec:      in.ConvertToRaw().PluseNanoSec,
	}
	return d.Driver.SendIR(ctx, irdata)
}

func (d *Driver) ReceiveIR(ctx context.Context) (ir.Data, error) {
	irdata, err := d.Driver.ReceiveIR(ctx)
	out := &ir.RawData{
		CarrierFreqKiloHz: irdata.CarrierFreqKiloHz,
		PluseNanoSec:      irdata.PluseNanoSec,
	}
	return out, err
}

func (d *Driver) GetDeviceInfo() *device.Info {
	return d.info
}
