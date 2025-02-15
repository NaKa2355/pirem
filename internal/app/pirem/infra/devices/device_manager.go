package devices

import (
	"context"
	"fmt"
	"time"

	"github.com/NaKa2355/pirem/internal/app/pirem/usecases"
	"github.com/NaKa2355/pirem/pkg/driver_module/v1"
)

var errDeviceBusy = fmt.Errorf("faild to get access admission because of too many requests")
var errNotSupportSending = fmt.Errorf("this device does not support sending")
var errNotSupportReceiving = fmt.Errorf("this device does not support receiving")

type DeviceManager struct {
	device            driver_module.Device
	sendIRInterval    time.Duration
	mutexLockDeadline time.Duration
	mu                chan (struct{})
	name              string
	info              *driver_module.DeviceInfo
}

func newDriverManager(driver driver_module.Device, name string, sendIRInterval time.Duration, mutexLockDeadLine time.Duration) (*DeviceManager, error) {
	info, err := driver.GetInfo(context.Background())
	return &DeviceManager{
		name:              name,
		device:            driver,
		mutexLockDeadline: mutexLockDeadLine,
		sendIRInterval:    sendIRInterval,
		mu:                make(chan struct{}, 1),
		info:              info,
	}, (err)
}

func (d *DeviceManager) GetInfo(ctx context.Context) *driver_module.DeviceInfo {
	return d.info
}

func (d *DeviceManager) SendIR(ctx context.Context, irData *driver_module.IRData) error {
	sender, canSend := d.device.(driver_module.Sender)
	if !canSend {
		return usecases.WrapError(
			usecases.CodeNotSupported,
			errNotSupportSending,
		)
	}
	select {
	case d.mu <- struct{}{}:
		defer func() {
			<-d.mu
		}()
		err := sender.SendIR(ctx, irData)
		//interval time to avoid conflict of data
		time.Sleep(d.sendIRInterval)
		return err

	case <-time.After(d.mutexLockDeadline):
		return usecases.WrapError(
			usecases.CodeBusy,
			errDeviceBusy,
		)

	case <-ctx.Done():
		ctx.Deadline()
		return ctx.Err()
	}

}

func (d *DeviceManager) ReceiveIR(ctx context.Context) (*driver_module.IRData, error) {
	irData := &driver_module.IRData{}
	receiver, canReceive := d.device.(driver_module.Receiver)
	if !canReceive {
		return irData, usecases.WrapError(
			usecases.CodeNotSupported,
			errNotSupportReceiving,
		)
	}

	select {
	case d.mu <- struct{}{}:
		defer func() {
			<-d.mu
		}()
		return receiver.ReceiveIR(ctx)

	case <-time.After(d.mutexLockDeadline):
		return irData, usecases.WrapError(
			usecases.CodeBusy,
			errDeviceBusy,
		)

	case <-ctx.Done():
		return irData, ctx.Err()
	}
}
