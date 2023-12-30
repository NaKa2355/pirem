package irdevice

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/NaKa2355/pirem/internal/app/pirem/domain/device"
	"github.com/NaKa2355/pirem/internal/app/pirem/domain/irdata"
	"github.com/NaKa2355/pirem/internal/app/pirem/usecases"
	"github.com/NaKa2355/pirem/internal/app/pirem/usecases/irdevice"
	"github.com/NaKa2355/pirem/pkg/module/v1"
)

var ErrDeviceNotFound = fmt.Errorf("device not found")

type IRData struct {
	irdata *module.IRData
}

func (i *IRData) ConvertToRaw() *irdata.RawData {
	return &irdata.RawData{
		CarrierFreqKiloHz: i.irdata.CarrierFreqKiloHz,
		PluseNanoSec:      i.irdata.PluseNanoSec,
	}
}

type IRDevices struct {
	mu      *sync.RWMutex
	devices map[string]*DriverManager
}

var _ irdevice.IRDevice = &IRDevices{}

func convertError(err *error) {
	if *err == nil {
		return
	}

	switch _err := (*err).(type) {
	case *module.Error:
		switch _err.Code {
		case module.CodeBusy:
			*err = usecases.WrapError(usecases.CodeBusy, *err)
		case module.CodeInvaildInput:
			*err = usecases.WrapError(usecases.CodeInvaildInput, *err)
		case module.CodeTimeout:
			*err = usecases.WrapError(usecases.CodeTimeout, *err)
		default:
			*err = usecases.WrapError(usecases.CodeUnknown, *err)
		}
	default:
		return
	}
}

func NewIRDevices() *IRDevices {
	return &IRDevices{
		mu:      &sync.RWMutex{},
		devices: map[string]*DriverManager{},
	}
}

func (devices *IRDevices) AddDevice(deviceId string, name string, driver module.Driver, sendIRInterval time.Duration, mutexLockDeadLine time.Duration) (err error) {
	devices.mu.Lock()
	defer convertError(&err)
	defer devices.mu.Unlock()
	devices.devices[deviceId], err = newDriverManager(driver, name, sendIRInterval, mutexLockDeadLine)
	return err
}

func (devices *IRDevices) ReadDevices(ctx context.Context) (fetchedDevices []*device.Device, err error) {
	devices.mu.RLock()
	defer convertError(&err)
	defer devices.mu.RUnlock()
	err = nil
	for id, d := range devices.devices {
		fetchedDevices = append(fetchedDevices, &device.Device{
			ID:              id,
			Name:            d.name,
			CanSend:         d.info.CanSend,
			CanReceive:      d.info.CanReceive,
			DriverVersion:   d.info.DriverVersion,
			FirmwareVersion: d.info.FirmwareVersion,
		})
	}
	return
}

func (devices *IRDevices) ReadDevice(ctx context.Context, id device.ID) (result *device.Device, err error) {
	devices.mu.RLock()
	defer convertError(&err)
	defer devices.mu.RUnlock()
	fetchedDevice, ok := devices.devices[id]
	if !ok {
		return nil, usecases.WrapError(usecases.CodeDeviceNotFound, ErrDeviceNotFound)
	}
	info := fetchedDevice.info

	result = &device.Device{
		ID:              id,
		CanSend:         info.CanSend,
		CanReceive:      info.CanReceive,
		DriverVersion:   info.DriverVersion,
		FirmwareVersion: info.DriverVersion,
	}
	return
}

func (devices *IRDevices) SendIR(ctx context.Context, deviceID device.ID, data irdata.IRData) (err error) {
	devices.mu.RLock()
	defer convertError(&err)
	defer devices.mu.RUnlock()
	device, ok := devices.devices[deviceID]
	if !ok {
		return usecases.WrapError(usecases.CodeDeviceNotFound, ErrDeviceNotFound)
	}
	rawIR := data.ConvertToRaw()
	return device.SendIR(ctx, &module.IRData{
		CarrierFreqKiloHz: rawIR.CarrierFreqKiloHz,
		PluseNanoSec:      rawIR.PluseNanoSec,
	})
}

func (devices *IRDevices) ReceiveIR(ctx context.Context, deviceID device.ID) (irdata irdata.IRData, err error) {
	devices.mu.RLock()
	defer convertError(&err)
	defer devices.mu.RUnlock()
	device, ok := devices.devices[deviceID]
	if !ok {
		return nil, usecases.WrapError(usecases.CodeDeviceNotFound, ErrDeviceNotFound)
	}
	llIRdata, err := device.ReceiveIR(ctx)
	return &IRData{irdata: llIRdata}, err
}
