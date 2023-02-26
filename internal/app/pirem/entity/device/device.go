package entity

import (
	"context"
	"fmt"
	"sync"
	"time"

	apiremv1 "github.com/NaKa2355/pirem/gen/apirem/v1"
	dev_usecases "github.com/NaKa2355/pirem/internal/app/pirem/usecases/device"
)

const SendIrInterval = 200 * time.Millisecond

type Device struct {
	ctrl dev_usecases.DeviceController
	info *apiremv1.DeviceInfo
	mu   sync.Mutex
}

// check the device supports sending infraread
func canSend(serviceType apiremv1.DeviceInfo_ServiceType) bool {
	if serviceType == apiremv1.DeviceInfo_SERVICE_TYPE_SEND || serviceType == apiremv1.DeviceInfo_SERVICE_TYPE_SEND_RECEIVE {
		return true
	}
	return false
}

// check the device supports receiving infraread
func canReceive(serviceType apiremv1.DeviceInfo_ServiceType) bool {
	if serviceType == apiremv1.DeviceInfo_SERVICE_TYPE_RECEIVE || serviceType == apiremv1.DeviceInfo_SERVICE_TYPE_SEND_RECEIVE {
		return true
	}
	return false
}

func New(id string, name string, devCtrl dev_usecases.DeviceController) (*Device, error) {
	dev := &Device{}
	info, err := devCtrl.GetDeviceInfo(context.Background())
	if err != nil {
		return dev, err
	}
	dev.info = info
	dev.ctrl = devCtrl
	dev.info.Name = name
	dev.info.Id = id
	return dev, nil
}

func (d *Device) GetDeviceInfo(ctx context.Context) (*apiremv1.DeviceInfo, error) {
	return d.info, nil
}

func (d *Device) IsBusy(ctx context.Context) (bool, error) {
	return d.ctrl.IsBusy(ctx)
}

func (d *Device) GetDeviceStatus(ctx context.Context) (*apiremv1.DeviceStatus, error) {
	var devStatus *apiremv1.DeviceStatus

	d.mu.Lock()
	defer d.mu.Unlock()

	isBusy, err := d.ctrl.IsBusy(context.Background())
	if err != nil {
		return devStatus, err
	}
	if isBusy {
		return devStatus, fmt.Errorf("device is busy")
	}

	select {
	case <-ctx.Done():
		return devStatus, ctx.Err()
	default:
		return d.ctrl.GetDeviceStatus(ctx)
	}
}

func (d *Device) SendRawIr(ctx context.Context, irData *apiremv1.RawIrData) error {
	if !canSend(d.info.Service) {
		return fmt.Errorf("this device does not support sending")
	}
	d.mu.Lock()
	defer d.mu.Unlock()

	isBusy, err := d.ctrl.IsBusy(context.Background())
	if err != nil {
		return err
	}
	if isBusy {
		return fmt.Errorf("device is busy")
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		err := d.ctrl.SendRawIr(ctx, irData)
		if err != nil {
			return err
		}

		//interval time to avoid conflict of data
		time.Sleep(SendIrInterval)
		return nil
	}
}

func (d *Device) ReceiveRawIr(ctx context.Context) (*apiremv1.RawIrData, error) {
	var irData *apiremv1.RawIrData
	if !canReceive(d.info.Service) {
		return irData, fmt.Errorf("this device does not support receiving")
	}

	d.mu.Lock()
	defer d.mu.Unlock()

	isBusy, err := d.ctrl.IsBusy(context.Background())
	if err != nil {
		return irData, err
	}
	if isBusy {
		return irData, fmt.Errorf("device is busy")
	}

	select {
	case <-ctx.Done():
		return irData, ctx.Err()
	default:
		return d.ctrl.ReceiveRawIr(ctx)
	}
}

func (d *Device) Drop() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.ctrl.Drop()
}
