package boundary

import (
	"context"

	"github.com/NaKa2355/pirem/internal/app/pirem/domain/device"
)

type GetDeivceInput struct {
	DeviceID device.ID
}

type DeviceGetter interface {
	GetDevice(context.Context, GetDeivceInput) (*device.Device, error)
}
