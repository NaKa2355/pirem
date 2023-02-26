package entity

import (
	"context"
	"fmt"
	"sync"
	"time"

	apiremv1 "github.com/NaKa2355/pirem/gen/apirem/v1"
	dev_usecases "github.com/NaKa2355/pirem/internal/app/pirem/usecases/device"
	plugin "github.com/hashicorp/go-plugin"
)

const SendIrInterval = 200 * time.Millisecond

type Device struct {
	Information *apiremv1.DeviceInfo
	ctrl        dev_usecases.DeviceController
	mu          sync.Mutex
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

func New(name string, devCtrl dev_usecases.DeviceController, client *plugin.Client) (*Device, error) {
	dev := &Device{}
	info, err := devCtrl.GetDeviceInfo(context.Background())
	if err != nil {
		return dev, err
	}
	dev.Information = info
	dev.ctrl = devCtrl
	return dev, nil
}

func (d *Device) GetDeviceInfo(ctx context.Context) *apiremv1.DeviceInfo {
	return d.Information
}

func (d *Device) IsBusy(ctx context.Context) (bool, error) {
	return d.ctrl.IsBusy(ctx)
}

func (d *Device) GetDeviceStatus(ctx context.Context) (*apiremv1.DeviceStatus, error) {
	var devStatus *apiremv1.DeviceStatus

	d.mu.Lock()
	defer d.mu.Unlock()

	isBusy, err := d.IsBusy(context.Background())
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

func (d *Device) SendIr(ctx context.Context, irData *apiremv1.RawIrData) error {
	if !canSend(d.Information.Service) {
		return fmt.Errorf("this device does not support sending")
	}
	d.mu.Lock()
	defer d.mu.Unlock()

	isBusy, err := d.IsBusy(context.Background())
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

func (d *Device) ReceiveIr(ctx context.Context) (*apiremv1.RawIrData, error) {
	var irData *apiremv1.RawIrData
	if !canReceive(d.Information.Service) {
		return irData, fmt.Errorf("this device does not support receiving")
	}

	d.mu.Lock()
	defer d.mu.Unlock()

	isBusy, err := d.IsBusy(context.Background())
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
