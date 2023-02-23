package device

import (
	"context"
	"fmt"
	"pirem/internal/app/pirem/usecases"
	apiremv1 "pirem/pkg/apirem.v1"
	dev_controller "pirem/pkg/device_controller"
	"sync"
	"time"
)

type Device struct {
	Information *apiremv1.DeviceInfo
	controller  dev_controller.DeviceController
	mu          sync.Mutex

	usecases.DeviceController
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

func New(id string, name string, dev_ctrler dev_controller.DeviceController) (*Device, error) {
	dev := &Device{}
	var err error
	ctx := context.Background()
	dev.controller = dev_ctrler
	dev.Information, err = dev.controller.GetDeviceInfo(ctx)
	dev.Information.Id = id
	dev.Information.Name = name
	return dev, err
}

func (d *Device) GetDeviceInfo(ctx context.Context) *apiremv1.DeviceInfo {
	return d.Information
}

func (d *Device) IsBusy(ctx context.Context) (bool, error) {
	isBusy := !d.mu.TryLock()
	if !isBusy {
		d.mu.Unlock()
	}
	return isBusy, nil
}

func (d *Device) GetDeviceStatus(ctx context.Context) (*apiremv1.DeviceStatus, error) {
	var devStatus *apiremv1.DeviceStatus
	d.mu.Lock()
	defer d.mu.Unlock()
	select {
	case <-ctx.Done():
		return devStatus, ctx.Err()
	default:
		return d.controller.GetDeviceStatus(ctx)
	}
}

func (d *Device) SendIr(ctx context.Context, ir_data *apiremv1.RawIrData) error {
	if !canSend(d.Information.Service) {
		return fmt.Errorf("this device does not support sending")
	}
	d.mu.Lock()
	defer d.mu.Unlock()
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		err := d.controller.SendIr(ctx, ir_data)
		if err != nil {
			return err
		}

		//interval time to avoid conflict of data
		time.Sleep(200 * time.Millisecond)
		return nil
	}
}

func (d *Device) ReceiveIr(ctx context.Context) (*apiremv1.RawIrData, error) {
	var ir_data *apiremv1.RawIrData
	if !canReceive(d.Information.Service) {
		return ir_data, fmt.Errorf("this device does not support receiving")
	}
	d.mu.Lock()
	defer d.mu.Unlock()
	select {
	case <-ctx.Done():
		return ir_data, ctx.Err()
	default:
		return d.controller.ReceiveIr(ctx)
	}
}

func (d *Device) Drop() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	err := d.controller.Drop()
	return err
}
