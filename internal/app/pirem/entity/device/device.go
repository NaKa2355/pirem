package device

import (
	"context"
	"fmt"
	"time"

	"github.com/NaKa2355/pirem/internal/app/pirem/entity"
	"github.com/NaKa2355/pirem/internal/app/pirem/entity/ir"
)

const SendIrInterval = 200 * time.Millisecond

type Driver interface {
	SendIR(ir.Data) error
	ReceiveIR() (ir.Data, error)
}

type Device struct {
	driver Driver
	Name   string
	ID     string
	info   Info
	mu     chan (struct{})
}

func New(id string, name string, info Info, driver Driver) (*Device, error) {
	dev := &Device{}
	dev.info = info

	dev.driver = driver
	dev.Name = name
	dev.ID = id
	dev.mu = make(chan struct{}, 1)
	return dev, nil
}

func (d *Device) GetDeviceInfo(ctx context.Context) Info {
	return d.info
}

func (d *Device) SendRawIR(ctx context.Context, irData ir.Data) error {
	if !d.info.CanSend {
		return entity.WrapErr(
			entity.CodeNotSupported,
			fmt.Errorf("this device does not support sending"),
		)
	}

	select {
	case d.mu <- struct{}{}:
		defer func() {
			<-d.mu
		}()

		err := d.driver.SendIR(irData)
		if err != nil {
			return err
		}

		//interval time to avoid conflict of data
		time.Sleep(SendIrInterval)
		return nil

	case <-ctx.Done():
		ctx.Deadline()
		return ctx.Err()
	}

}

func (d *Device) ReceiveRawIR(ctx context.Context) (ir.Data, error) {
	var irData ir.Data
	if !d.info.CanReceive {
		return irData, fmt.Errorf("this device does not support receiving")
	}

	select {
	case d.mu <- struct{}{}:
		defer func() {
			<-d.mu
		}()
		return d.driver.ReceiveIR()
	case <-ctx.Done():
		return irData, ctx.Err()
	}
}
