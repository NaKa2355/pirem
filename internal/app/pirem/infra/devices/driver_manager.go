package devices

import (
	"context"
	"fmt"
	"time"

	"github.com/NaKa2355/pirem/internal/app/pirem/usecases"
	"github.com/NaKa2355/pirem/pkg/module/v1"
)

var errDeviceBusy = fmt.Errorf("faild to get access admission because of too many requests")
var errNotSupportSending = fmt.Errorf("this device does not support sending")
var errNotSupportReceiving = fmt.Errorf("this device does not support receiving")

type DriverManager struct {
	driver            module.Driver
	sendIRInterval    time.Duration
	mutexLockDeadline time.Duration
	mu                chan (struct{})
	name              string
	info              *module.DeviceInfo
}

func newDriverManager(driver module.Driver, name string, sendIRInterval time.Duration, mutexLockDeadLine time.Duration) (*DriverManager, error) {
	info, err := driver.GetInfo(context.Background())
	return &DriverManager{
		name:              name,
		driver:            driver,
		mutexLockDeadline: mutexLockDeadLine,
		sendIRInterval:    sendIRInterval,
		mu:                make(chan struct{}, 1),
		info:              info,
	}, (err)
}

func (d *DriverManager) GetInfo(ctx context.Context) *module.DeviceInfo {
	return d.info
}

func (d *DriverManager) SendIR(ctx context.Context, irData *module.IRData) error {
	if !d.info.CanSend {
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
		err := d.driver.SendIR(ctx, irData)
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

func (d *DriverManager) ReceiveIR(ctx context.Context) (*module.IRData, error) {
	irData := &module.IRData{}
	if !d.info.CanReceive {
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
		return d.driver.ReceiveIR(ctx)

	case <-time.After(d.mutexLockDeadline):
		return irData, usecases.WrapError(
			usecases.CodeBusy,
			errDeviceBusy,
		)

	case <-ctx.Done():
		return irData, ctx.Err()
	}
}
