package boundary

import (
	"context"

	"github.com/NaKa2355/pirem/internal/app/pirem/domain/device"
)

type ListDevicesOutput struct {
	Devices []*device.Device
}

type DevicesLister interface {
	ListDevices(context.Context) ([]*device.Device, error)
}
