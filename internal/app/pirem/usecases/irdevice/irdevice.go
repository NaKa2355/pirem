package irdevice

import (
	"context"

	"github.com/NaKa2355/pirem/internal/app/pirem/domain/device"
	"github.com/NaKa2355/pirem/internal/app/pirem/domain/irdata"
)

type IRDevice interface {
	ReadDevices(ctx context.Context) ([]*device.Device, error)
	ReadDevice(ctx context.Context, id device.ID) (*device.Device, error)

	SendIR(ctx context.Context, id device.ID, data irdata.IRData) error
	ReceiveIR(tx context.Context, id device.ID) (irdata.IRData, error)
}
