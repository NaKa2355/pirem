package boundary

import (
	"context"

	"github.com/NaKa2355/pirem/internal/app/pirem/domain"
)

type ListDevicesOutput struct {
	Devices []*domain.Device
}

type DevicesLister interface {
	ListDevices(context.Context) ([]*domain.Device, error)
}
