package interactor

import (
	"context"

	"github.com/NaKa2355/pirem/internal/app/pirem/entity/device"
	"github.com/NaKa2355/pirem/internal/app/pirem/entity/ir"
	bdy "github.com/NaKa2355/pirem/internal/app/pirem/usecases/boundary"
	repo "github.com/NaKa2355/pirem/internal/app/pirem/usecases/repository"
)

type Repository interface {
	repo.DeviceCreator
	repo.DeviceReader
	repo.DevicesReader
	repo.DeviceDeleter
}

type Interactor struct {
	devsRepo Repository
}

// get all devices information
func (i *Interactor) getDevicesInfo(ctx context.Context) (out bdy.GetDevicesInfoOutput, err error) {
	devs, err := i.devsRepo.ReadDevices()
	if err != nil {
		return out, err
	}

	out.Devices = make([]bdy.DeviceInfo, 0, len(devs))
	for _, d := range devs {
		info := d.GetDeviceInfo(ctx)
		out.Devices = append(out.Devices, bdy.DeviceInfo{
			ID:              d.ID,
			Name:            d.Name,
			CanSend:         info.CanSend,
			CanReceive:      info.CanReceive,
			DriverVersion:   info.DriverVersion,
			FirmwareVersion: info.FirmwareVersion,
		})
	}
	return out, err
}

// get device information
func (i *Interactor) getDeviceInfo(ctx context.Context, in bdy.GetDeviceInput) (out bdy.GetDeviceInfoOutput, err error) {
	dev, err := i.devsRepo.ReadDevice(in.ID)
	if err != nil {
		return out, err
	}

	info := dev.GetDeviceInfo(ctx)
	out.Device = bdy.DeviceInfo{
		ID:              dev.ID,
		Name:            dev.Name,
		CanSend:         info.CanSend,
		CanReceive:      info.CanReceive,
		DriverVersion:   info.DriverVersion,
		FirmwareVersion: info.FirmwareVersion,
	}

	return out, err
}

func (i *Interactor) sendIR(ctx context.Context, in bdy.SendIRInput) (err error) {
	var irdata ir.Data

	dev, err := i.devsRepo.ReadDevice(in.ID)
	if err != nil {
		return err
	}

	irdata = &ir.RawData{
		CarrierFreqKiloHz: in.IRData.ConvertToRaw().CarrierFreqKiloHz,
		PluseNanoSec:      in.IRData.ConvertToRaw().PluseNanoSec,
	}

	return dev.SendIR(ctx, irdata)
}

func (i *Interactor) receiveIR(ctx context.Context, in bdy.ReceiveIRInput) (out bdy.IRData, err error) {
	device, err := i.devsRepo.ReadDevice(in.ID)
	if err != nil {
		return out, err
	}

	irData, err := device.ReceiveIR(ctx)
	if err != nil {
		return out, err
	}

	rawIRData := irData.ConvertToRaw()
	out = bdy.RawIRData{
		CarrierFreqKiloHz: rawIRData.CarrierFreqKiloHz,
		PluseNanoSec:      rawIRData.PluseNanoSec,
	}

	return out, err
}

func (i *Interactor) addDevice(ctx context.Context, in bdy.AddDeviceInput) error {
	md, err := in.Module.NewDriver(in.Config)
	if err != nil {
		return err
	}

	d, err := NewDriver(md)
	if err != nil {
		return err
	}

	ed, err := device.New(in.ID, in.DeviceName, d)
	if err != nil {
		return err
	}

	return i.devsRepo.CreateDevice(ed)
}
