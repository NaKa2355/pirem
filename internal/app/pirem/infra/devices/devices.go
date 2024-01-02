package devices

import (
	"context"
	"fmt"
	"sync"
	"time"

	adapter "github.com/NaKa2355/pirem/internal/app/pirem/adapter/device_module"
	"github.com/NaKa2355/pirem/internal/app/pirem/domain"
	"github.com/NaKa2355/pirem/internal/app/pirem/usecases"
	"github.com/NaKa2355/pirem/internal/app/pirem/usecases/controllers"
	"github.com/NaKa2355/pirem/pkg/module/v1"
)

var ErrDeviceNotFound = fmt.Errorf("device not found")

type IRDevices struct {
	mu      *sync.RWMutex
	devices map[string]*DriverManager
}

var _ controllers.IRDevice = &IRDevices{}

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

func (devices *IRDevices) ReadDevices(ctx context.Context) (fetchedDevices []*domain.Device, err error) {
	devices.mu.RLock()
	defer convertError(&err)
	defer devices.mu.RUnlock()
	err = nil
	for id, d := range devices.devices {
		fetchedDevices = append(fetchedDevices, &domain.Device{
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

func (devices *IRDevices) ReadDevice(ctx context.Context, id domain.DeviceID) (result *domain.Device, err error) {
	devices.mu.RLock()
	defer convertError(&err)
	defer devices.mu.RUnlock()
	fetchedDevice, ok := devices.devices[id]
	if !ok {
		return &domain.Device{}, usecases.WrapError(usecases.CodeDeviceNotFound, ErrDeviceNotFound)
	}
	info := fetchedDevice.info

	result = &domain.Device{
		ID:              id,
		CanSend:         info.CanSend,
		CanReceive:      info.CanReceive,
		DriverVersion:   info.DriverVersion,
		FirmwareVersion: info.DriverVersion,
	}
	return
}

func (devices *IRDevices) SendIR(ctx context.Context, deviceID domain.DeviceID, data domain.IRData) (err error) {
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

func (devices *IRDevices) ReceiveIR(ctx context.Context, deviceID domain.DeviceID) (irData domain.IRData, err error) {
	devices.mu.RLock()
	defer convertError(&err)
	defer devices.mu.RUnlock()
	device, ok := devices.devices[deviceID]
	if !ok {
		return nil, usecases.WrapError(usecases.CodeDeviceNotFound, ErrDeviceNotFound)
	}
	d, err := device.ReceiveIR(ctx)
	return adapter.UnMarshalIRData(d), err
}
