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
	"github.com/NaKa2355/pirem/pkg/driver_module/v1"
)

var ErrDeviceNotFound = fmt.Errorf("device not found")

type IRDevices struct {
	mu      *sync.RWMutex
	devices map[string]*DeviceManager
}

var _ controllers.IRDevice = &IRDevices{}

func convertError(err *error) {
	if *err == nil {
		return
	}

	switch _err := (*err).(type) {
	case *driver_module.Error:
		switch _err.Code {
		case driver_module.CodeBusy:
			*err = usecases.WrapError(usecases.CodeBusy, *err)
		case driver_module.CodeInvaildInput:
			*err = usecases.WrapError(usecases.CodeInvaildInput, *err)
		case driver_module.CodeTimeout:
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
		devices: map[string]*DeviceManager{},
	}
}

func (devices *IRDevices) AddDevice(deviceId string, name string, driver driver_module.Device, sendIRInterval time.Duration, mutexLockDeadLine time.Duration) (err error) {
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

		_, canSend := d.device.(driver_module.Sender)
		_, canReceive := d.device.(driver_module.Receiver)

		fetchedDevices = append(fetchedDevices, &domain.Device{
			ID:              id,
			Name:            d.name,
			CanSend:         canSend,
			CanReceive:      canReceive,
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
		return &domain.Device{}, usecases.WrapError(usecases.CodeNotFound, ErrDeviceNotFound)
	}
	info := fetchedDevice.info

	_, canSend := fetchedDevice.device.(driver_module.Sender)
	_, canReceive := fetchedDevice.device.(driver_module.Receiver)

	result = &domain.Device{
		ID:              id,
		CanSend:         canSend,
		CanReceive:      canReceive,
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
		return usecases.WrapError(usecases.CodeNotFound, ErrDeviceNotFound)
	}
	rawIR := data.ConvertToRaw()
	return device.SendIR(ctx, &driver_module.IRData{
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
		return nil, usecases.WrapError(usecases.CodeNotFound, ErrDeviceNotFound)
	}
	d, err := device.ReceiveIR(ctx)
	return adapter.UnMarshalIRData(d), err
}
