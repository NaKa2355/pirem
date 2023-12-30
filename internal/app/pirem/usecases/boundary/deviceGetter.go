package boundary

import (
	"context"

	"github.com/NaKa2355/pirem/internal/app/pirem/domain"
)

type GetDeivceInput struct {
	DeviceID domain.DeviceID
}

type DeviceGetter interface {
	GetDevice(context.Context, GetDeivceInput) (*domain.Device, error)
}
