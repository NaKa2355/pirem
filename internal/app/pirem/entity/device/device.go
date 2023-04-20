/*
赤外線送受信機の機能と排他処理のふるまいを書く
*/

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
	SendIR(context.Context, ir.Data) error
	ReceiveIR(context.Context) (ir.Data, error)
	GetDeviceInfo() *Info
}

type Device struct {
	Name   string
	ID     string
	driver Driver
	mu     chan (struct{})
}

func New(id string, name string, driver Driver) (*Device, error) {
	dev := &Device{}
	dev.driver = driver
	dev.Name = name
	dev.ID = id
	dev.mu = make(chan struct{}, 1)
	return dev, nil
}

func (d *Device) GetDeviceInfo(ctx context.Context) *Info {
	return d.driver.GetDeviceInfo()
}

func (d *Device) SendIR(ctx context.Context, irData ir.Data) error {
	if !d.driver.GetDeviceInfo().CanSend {
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

		err := d.driver.SendIR(ctx, irData)
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

func (d *Device) ReceiveIR(ctx context.Context) (ir.Data, error) {
	var irData ir.Data
	if !d.driver.GetDeviceInfo().CanReceive {
		return irData, entity.WrapErr(
			entity.CodeNotSupported,
			fmt.Errorf("this device does not support sending"),
		)
	}

	select {
	case d.mu <- struct{}{}:
		defer func() {
			<-d.mu
		}()
		return d.driver.ReceiveIR(ctx)
	case <-ctx.Done():
		return irData, ctx.Err()
	}
}
